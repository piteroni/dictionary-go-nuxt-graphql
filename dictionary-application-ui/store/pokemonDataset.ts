import { getterTree, mutationTree } from "typed-vuex"
import { Type, Gender, Characteristic, Description, Ability, LinkInfo, PokemonQuery, Pokemon } from "@/graphql/generated"
import { QueryType } from "@/shared/graphql"

export type PokemonQueryType<Typename extends PokemonQuery["pokemon"]["__typename"]> = QueryType<PokemonQuery["pokemon"], Typename>

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
  linkInfo: {
    prevId: 0,
    nextId: 0,
    hasPrev: false,
    hasNext: false
  } as LinkInfo,
  evolutions: [] as PokemonQueryType<"Pokemon">["evolutions"]
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
  save(state, params: PokemonQueryType<"Pokemon">): void {
    state.nationalNo = params.nationalNo
    state.name = params.name
    state.species = params.species
    state.height = params.height
    state.weight = params.weight
    state.imageURL = params.imageURL
    state.genders = params.genders
    state.types = params.types
    state.characteristics = params.characteristics
    state.description = params.description
    state.ability = params.ability
    state.linkInfo = params.linkInfo
    state.evolutions = params.evolutions
  }
})
