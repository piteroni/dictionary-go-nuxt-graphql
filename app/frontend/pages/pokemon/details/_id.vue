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
import { defineComponent, reactive } from "@nuxtjs/composition-api"
import { useQuery } from "@vue/apollo-composable"
import { PokemonDocument, PokemonQuery, PokemonQueryVariables } from "@/graphql/generated/client"
import { PokemonGender } from "~/components/tightly-coupled/pokemon/details/_id/types"
import Header from "@/components/singleton/Header.vue"
import PokemonHeading from "@/components/tightly-coupled/pokemon/details/_id/PokemonHeading.vue"
import PokemonDetails from "@/components/tightly-coupled/pokemon/details/_id/PokemonDetails.vue"

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

    const variables: PokemonQueryVariables = {
      pokemonId: 1
    }

    const { onResult } = useQuery<PokemonQuery>(PokemonDocument, variables)

    if (process.client) {
      const format = (nationalNo: number) => ("000" + nationalNo.toString()).slice(-3)
      onResult(result => {
        if (!result.loading && !result.error) {
          const nationalNo = format(result.data.pokemon.nationalNo)
          const genders: PokemonGender[] = result.data.pokemon.genders.map(gender => {
            return {
              name: gender.name,
              iconURL: `/image/${gender.iconName}`
            }
          })

          state.nationalNo = `No.${nationalNo}`
          state.name = result.data.pokemon.name
          state.imageURL = `/image/${result.data.pokemon.imageName}`
          state.genders = genders
        }
      })
    }

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