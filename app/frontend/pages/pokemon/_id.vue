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

    <pokemon-details
      :species="state.species"
      :height="state.height"
      :weight="state.weight"
      :types="state.types"
      :characteristics="state.characteristics"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, useContext, useFetch } from "@nuxtjs/composition-api"
import { useQuery } from "@vue/apollo-composable"
import { PokemonDocument, PokemonQuery, PokemonQueryVariables, Characteristic } from "@/graphql/generated/client"
import { PokemonGender, PokemonType } from "@/components/tightly-coupled/pokemon/_id/types"
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
      species: string,
      weight: string,
      height: string,
      types: PokemonType[],
      characteristics: Characteristic[],
      genders: PokemonGender[]
    }>({
      nationalNo: "",
      name: "",
      imageURL: "",
      species: "",
      weight: "",
      height: "",
      genders: [],
      types: [],
      characteristics: []
    })

    const { route, error } = useContext()

    const pokemonId = parseInt(route.value.params.id)

    if (isNaN(pokemonId)) {
      error({ statusCode: 404 })

      return { state }
    }

    useFetch(async () => {
      const variables: PokemonQueryVariables = {
        pokemonId: pokemonId
      }

      const { onError, onResult } = useQuery<PokemonQuery>(PokemonDocument, variables)

      await new Promise<void>((resolve, reject) => {
        onResult(result => {
          if (result.loading || result.error) {
            return
          }

          const format = (nationalNo: number) => ("000" + nationalNo.toString()).slice(-3)

          const nationalNo = format(result.data.pokemon.nationalNo)
          const genders: PokemonGender[] = result.data.pokemon.genders.map(gender => {
            return {
              name: gender.name,
              iconURL: `/image/${gender.iconName}`
            }
          })
          const types: PokemonType[] = result.data.pokemon.types.map(type => {
            return {
              name: type.name,
              iconURL: `/image/${type.iconName}`
            }
          })

          state.nationalNo = `No.${nationalNo}`
          state.name = result.data.pokemon.name
          state.species = result.data.pokemon.species
          state.height = result.data.pokemon.height
          state.weight = result.data.pokemon.weight
          state.imageURL = `/image/${result.data.pokemon.imageName}`
          state.genders = genders
          state.types = types
          state.characteristics = result.data.pokemon.characteristics

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