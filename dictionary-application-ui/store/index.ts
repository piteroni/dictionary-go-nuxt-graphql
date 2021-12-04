import { getAccessorType } from "typed-vuex"
import * as pokemonDataset from "@/store/pokemonDataset"
import * as pokemonConnection from "@/store/pokemonConnection"

export const accessorType = getAccessorType({
  modules: {
    pokemonDataset,
    pokemonConnection
  }
})
