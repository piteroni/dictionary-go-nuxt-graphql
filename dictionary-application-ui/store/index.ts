import { getAccessorType } from "typed-vuex"
import * as pokemonDataset from "@/store/pokemonDataset"

export const accessorType = getAccessorType({
  modules: {
    pokemonDataset
  }
})
