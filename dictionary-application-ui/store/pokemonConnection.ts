import { mutationTree } from "typed-vuex"
import { PokemonsQuery } from "@/graphql/generated"
import { QueryType } from "@/shared/graphql"

type K = keyof Omit<PokemonsQuery, "__typename">
export type PokemonsQueryType<Model extends K, Typename extends PokemonsQuery[Model]["__typename"]> = QueryType<PokemonsQuery[Model], Typename>

const state = () => ({
  nextID: 0,
  pokemons: [] as PokemonsQueryType<"pokemons", "PokemonConnection">["items"]
})

export const mutations = mutationTree(state, {
  save(state, params: {
    nextID: number,
    pokemons: PokemonsQueryType<"pokemons", "PokemonConnection">["items"]
  }): void {
    state.nextID = params.nextID
    state.pokemons = params.pokemons
  }
})
