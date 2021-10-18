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
      :description="state.description"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, useContext, useFetch } from "@nuxtjs/composition-api"
import { useQuery } from "@vue/apollo-composable"
import { PokemonDocument, PokemonQuery, PokemonQueryVariables, Type, Gender, Characteristic, Description } from "@/graphql/generated/client"
import Header from "@/components/Header.vue"
import PokemonHeading from "@/composable/pokemons/_id/PokemonHeading.vue"
import PokemonDetails from "@/composable/pokemons/_id/PokemonDetails.vue"

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
      types: Type[],
      characteristics: Characteristic[],
      genders: Gender[],
      description: Description
    }>({
      nationalNo: "",
      name: "",
      imageURL: "",
      species: "",
      weight: "",
      height: "",
      genders: [],
      types: [],
      characteristics: [],
      description: { text: "", series: "" }
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

          state.nationalNo = `No.${nationalNo}`
          state.name = result.data.pokemon.name
          state.species = result.data.pokemon.species
          state.height = result.data.pokemon.height
          state.weight = result.data.pokemon.weight
          state.imageURL = result.data.pokemon.imageURL
          state.genders = result.data.pokemon.genders
          state.types = result.data.pokemon.types
          state.characteristics = result.data.pokemon.characteristics
          state.description = result.data.pokemon.description

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