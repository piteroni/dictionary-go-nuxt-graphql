import { readonly, reactive, computed, InjectionKey } from "@nuxtjs/composition-api"
import { useQuery } from "@vue/apollo-composable"
import { Characteristic, Description, Gender, Type, PokemonQuery, PokemonDocument, PokemonQueryVariables, Ability, LinkInfo } from "@/graphql/generated/client"

type State = {
  nationalNo: number,
  name: string,
  imageURL: string,
  species: string,
  weight: string,
  height: string,
  types: Type[],
  characteristics: Characteristic[],
  genders: Gender[],
  description: Description,
  ability: Ability,
  linkInfo: LinkInfo,
  evolutions: PokemonQuery["pokemon"]["evolutions"]
}

const initialState: State = {
  nationalNo: 0,
  name: "",
  imageURL: "",
  species: "",
  weight: "",
  height: "",
  genders: [],
  types: [],
  characteristics: [],
  description: {
    text: "",
    series: ""
  },
  ability: {
    heart: 0,
    attack: 0,
    defense: 0,
    specialAttack: 0,
    specialDefense: 0,
    speed: 0
  },
  linkInfo: {
    prevNationalNo: 0,
    nextNationalNo: 0,
    hasPrev: false,
    hasNext: false
  },
  evolutions: []
}

export const abilityMaxStatus = readonly({
  heart: 250,
  attack: 250,
  defense: 250,
  specialAttack: 250,
  specialDefense: 250,
  speed: 250
})

const fetch = (state: State) => (variables: PokemonQueryVariables) => {
  const { onError, onResult } = useQuery<PokemonQuery>(PokemonDocument, variables)

  return new Promise<void>((resolve, reject) => {
    onResult(result => {
      if (result.loading || result.error) {
        return
      }

      state.nationalNo = result.data.pokemon.nationalNo
      state.name = result.data.pokemon.name
      state.species = result.data.pokemon.species
      state.height = result.data.pokemon.height
      state.weight = result.data.pokemon.weight
      state.imageURL = result.data.pokemon.imageURL
      state.genders = result.data.pokemon.genders
      state.types = result.data.pokemon.types
      state.characteristics = result.data.pokemon.characteristics
      state.description = result.data.pokemon.description
      state.ability = result.data.pokemon.ability
      state.linkInfo = result.data.pokemon.linkInfo
      state.evolutions = result.data.pokemon.evolutions

      resolve()
    })

    onError(e => reject(e))
  })
}

export function usePokemonDetails(initial = initialState) {
  const state: State = reactive(initial)

  return {
    pokemon: computed(() => state),
    nationalNo: computed(() => state.nationalNo),
    name: computed(() => state.name),
    imageURL: computed(() => state.imageURL),
    species: computed(() => state.species),
    height: computed(() => state.height),
    weight: computed(() => state.weight),
    genders: computed(() => state.genders),
    types: computed(() => state.types),
    characteristics: computed(() => state.characteristics),
    description: computed(() => state.description),
    ability: computed(() => state.ability),
    linkInfo: computed(() => state.linkInfo),
    evolutions: computed(() => state.evolutions),
    fetch: fetch(state),
  }
}

export type PokemonDetailsComposition = ReturnType<typeof usePokemonDetails>

export const pokemonDetailsKey: InjectionKey<PokemonDetailsComposition> = Symbol("pokemon-details")
