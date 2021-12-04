import { getterTree, mutationTree } from "typed-vuex"
import { Type, Gender, Characteristic, Description, Ability, PokemonQuery } from "@/graphql/generated"
import { QueryType } from "@/shared/graphql"

export type PokemonQueryType<Model extends keyof Omit<PokemonQuery, "__typename">, Typename extends PokemonQuery[Model]["__typename"]> = QueryType<PokemonQuery[Model], Typename>

export const state = () => ({
  nationalNo: 0,
  name: "",
  imageURL: "",
  species: "",
  weight: "",
  height: "",
  genders: [] as Gender[],
  types: [] as Type[],
  characteristics: [] as Characteristic[],
  description: {
    text: "",
    series: ""
  } as Description,
  ability: {
    heart: 0,
    attack: 0,
    defense: 0,
    specialAttack: 0,
    specialDefense: 0,
    speed: 0
  } as Ability,
  pageInfo: {
    prevId: 0,
    nextId: 0,
    hasPrev: false,
    hasNext: false
  } as PokemonQueryType<"pageInfo", "PageInfo">,
  evolutions: [] as PokemonQueryType<"evolutions", "Evolutions">["pokemons"]
})

export const abilityMaxStatus = {
  heart: 250,
  attack: 250,
  defense: 250,
  specialAttack: 250,
  specialDefense: 250,
  speed: 250
}

export const nationalNoToText = (value: number): string => {
  return `No.${value.toString().padStart(3, "0")}`
}

export const getters = getterTree(state, {
  nationalNoText(state): string {
    return nationalNoToText(state.nationalNo)
  }
})

export const mutations = mutationTree(state, {
  save(state, params: {
    pokemon: PokemonQueryType<"pokemon", "Pokemon">,
    pageInfo: PokemonQueryType<"pageInfo", "PageInfo">,
    evolutions: PokemonQueryType<"evolutions", "Evolutions">
  }): void {
    state.nationalNo = params.pokemon.nationalNo
    state.name = params.pokemon.name
    state.species = params.pokemon.species
    state.height = params.pokemon.height
    state.weight = params.pokemon.weight
    state.imageURL = params.pokemon.imageURL
    state.genders = params.pokemon.genders
    state.types = params.pokemon.types
    state.characteristics = params.pokemon.characteristics
    state.description = params.pokemon.description
    state.ability = params.pokemon.ability
    state.pageInfo = params.pageInfo
    state.evolutions = params.evolutions.pokemons
  }
})
