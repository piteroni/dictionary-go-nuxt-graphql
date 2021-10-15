<template>
  <div class="site">
    <app-header>
      <h1 class="site-logo">
        <img src="~/assets/image/logo.svg" alt="site-logo">
      </h1>
    </app-header>

    <pokemon-heading
      :nationalNo="state.nationalNo"
      :name="state.name"
      :imageURL="state.imageURL"
      :genders="state.genders"
    />

    <pokemon-details />
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, useContext, useFetch } from "@nuxtjs/composition-api"
import { useQuery } from "@vue/apollo-composable"
import { PokemonDocument, PokemonQuery, PokemonQueryVariables } from "@/graphql/generated/client"
import { PokemonGender } from "@/components/tightly-coupled/pokemon/_id/types"
import Header from "@/components/singleton/Header.vue"
import PokemonHeading from "@/components/tightly-coupled/pokemon/_id/PokemonHeading.vue"
import PokemonDetails from "@/components/tightly-coupled/pokemon/_id/PokemonDetails.vue"

export default defineComponent({
  components: {
    "app-header": Header,
    "pokemon-heading": PokemonHeading,
    "pokemon-details": PokemonDetails
  },
  setup() {
    const state = reactive<{
      nationalNo: string,
      name: string,
      imageURL: string,
      genders: PokemonGender[]
    }>({
      nationalNo: "",
      name: "",
      imageURL: "",
      genders: []
    })

    const { route, error, redirect } = useContext()

    const pokemonId = parseInt(route.value.params.id)

    if (isNaN(pokemonId)) {
      redirect("/404")

      return { state }
    }

    const variables: PokemonQueryVariables = {
      pokemonId: pokemonId
    }

    useFetch(async () => {
      const { onError, onResult } = useQuery<PokemonQuery>(PokemonDocument, variables)

      await new Promise<void>((resolve, reject) => {
        onResult(result => {
          if (result.loading || result.error) {
            return
          }

          const format = (nationalNo: number) => ("000" + nationalNo.toString()).slice(-3)

          const pokemonDetails = result.data.pokemon

          const nationalNo = format(pokemonDetails.nationalNo)
          const genders: PokemonGender[] = pokemonDetails.genders.map(gender => {
            return {
              name: gender.name,
              iconURL: `/image/${gender.iconName}`
            }
          })

          state.nationalNo = `No.${nationalNo}`
          state.name = pokemonDetails.name
          state.imageURL = `/image/${pokemonDetails.imageName}`
          state.genders = genders

          resolve()
        })

        onError(e => {
          error({ statusCode: 404 })
          reject(e)
        })
      })
    })
      
    return { state }
  }
})
</script>

<style scoped>
.site {
  background-color: rgb(242, 242, 242);
  background-image: url('~/assets/image/pattern.svg');
  background-repeat: repeat;
}

.site-logo {
  width: 192px;
}
</style>