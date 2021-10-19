import { readonly, reactive, InjectionKey } from "@vue/composition-api"
import { Characteristic, Description, Gender, Type, PokemonQuery, PokemonDocument, PokemonQueryVariables, Ability } from "@/graphql/generated/client"
import { useQuery } from "@vue/apollo-composable"

type State = {
  nationalNo: string,
  name: string,
  imageURL: string,
  species: string,
  weight: string,
  height: string,
  types: Type[],
  characteristics: Characteristic[],
  genders: Gender[],
  description: Description,
  ability: Ability
}

const initialState: State = {
  nationalNo: "",
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
  }
}

const formatNationalNo = (nationalNo: number) => {
  const formated = ("000" + nationalNo.toString()).slice(-3)

  return `No.${formated}`
}

const evaluateAbility = (ability: Ability): Ability => {
  const evaluation = 15

  return {
    heart: scaleAbility(ability.heart, abilityMaxStatus.heart, evaluation),
    attack: scaleAbility(ability.attack, abilityMaxStatus.attack, evaluation),
    defense: scaleAbility(ability.defense, abilityMaxStatus.defense, evaluation),
    specialAttack: scaleAbility(ability.specialAttack, abilityMaxStatus.specialAttack, evaluation),
    specialDefense: scaleAbility(ability.specialDefense, abilityMaxStatus.specialDefense, evaluation),
    speed: scaleAbility(ability.speed, abilityMaxStatus.speed, evaluation),
  }
}

const abilityMaxStatus = readonly({
  heart: 250,
  attack: 250,
  defense: 250,
  specialAttack: 250,
  specialDefense: 250,
  speed: 250
})

const scaleAbility = (value: number, maxValue: number, evaluation: number): number => {
  return Math.round(evaluation * (value / maxValue))
}

const fetch = (state: State) => async (pokemonId: number) => {
  const variables: PokemonQueryVariables = {
    pokemonId
  }

  const { onError, onResult } = useQuery<PokemonQuery>(PokemonDocument, variables)

  return new Promise<void>((resolve, reject) => {
    onResult(result => {
      if (result.loading || result.error) {
        return
      }

      state.name = result.data.pokemon.name
      state.species = result.data.pokemon.species
      state.height = result.data.pokemon.height
      state.weight = result.data.pokemon.weight
      state.imageURL = result.data.pokemon.imageURL
      state.genders = result.data.pokemon.genders
      state.types = result.data.pokemon.types
      state.characteristics = result.data.pokemon.characteristics
      state.description = result.data.pokemon.description

      state.nationalNo = formatNationalNo(result.data.pokemon.nationalNo)
      state.ability = evaluateAbility(result.data.pokemon.ability)

      resolve()
    })

    onError(e => reject(e))
  })
}

export function usePokemonDetails(initial = initialState) {
  const state: State = reactive(initial)

  return {
    pokemon: readonly(state),
    fetch: fetch(state),
  }
}

export type PokemonDetailsComposition = ReturnType<typeof usePokemonDetails>

export const pokemonDetailsKey: InjectionKey<PokemonDetailsComposition> = Symbol("pokemon-details")
