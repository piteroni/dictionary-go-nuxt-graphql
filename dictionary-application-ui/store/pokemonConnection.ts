import { mutationTree } from "typed-vuex"
import { PokemonsQuery } from "@/graphql/generated"
import { QueryType } from "@/shared/graphql"

export type PokemonsQueryType<Model extends keyof Omit<PokemonsQuery, "__typename">, Typename extends PokemonsQuery[Model]["__typename"]> = QueryType<PokemonsQuery[Model], Typename>

export const fetchCount = 64

export const state = () => ({
  endCursor: "",
  hasNext: false,
  pokemons: [] as PokemonsQueryType<"pokemons", "PokemonConnection">["items"]
})

export const mutations = mutationTree(state, {
  save(state, params: PokemonsQueryType<"pokemons", "PokemonConnection">): void {
    state.endCursor = params.endCursor
    state.hasNext = params.hasNext
    state.pokemons = params.items
  },
  accumulate(state, params: PokemonsQueryType<"pokemons", "PokemonConnection">): void {
    state.endCursor = params.endCursor
    state.hasNext = params.hasNext
    state.pokemons.push(...params.items)
  }
})
