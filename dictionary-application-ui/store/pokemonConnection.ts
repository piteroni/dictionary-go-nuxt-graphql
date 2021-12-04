import { mutationTree } from "typed-vuex"
import { PokemonsQuery } from "@/graphql/generated"
import { QueryType } from "@/shared/graphql"

export type PokemonsQueryType<Model extends keyof Omit<PokemonsQuery, "__typename">, Typename extends PokemonsQuery[Model]["__typename"]> = QueryType<PokemonsQuery[Model], Typename>

export const state = () => ({
  nextID: 0,
  pokemons: [] as PokemonsQueryType<"pokemons", "PokemonConnection">["items"]
})

export const mutations = mutationTree(state, {
  save(state, params: PokemonsQueryType<"pokemons", "PokemonConnection">): void {
    state.nextID = params.nextID
    state.pokemons = params.items
  }
})
