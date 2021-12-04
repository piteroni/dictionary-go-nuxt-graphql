<template>
  <div class="flex h-full px-8 py-14 bg-white">
    <div v-for="pokemon, key in pokemons" :key="key">
      <div>
        <img class="pokemon-image-border mr-2" :src="pokemon.imageURL" :alt="pokemon.name">
        <p class="mt-1">
          {{ pokemon.name }}
        </p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "nuxt-property-decorator"
import { PokemonsQueryType } from "@/store/pokemonConnection"

@Component
export default class PokemenList extends Vue {
  public get pokemons(): PokemonsQueryType<"pokemons", "PokemonConnection">["items"] {
    return this.$accessor.pokemonConnection.pokemons
  }
}
</script>

<style scoped>
.pokemon-image-border {
  width: 10.5rem;
  border: solid 1.5px #e8e8e8;
  border-radius: 5px;
  cursor: pointer;
  transition: .3s;
}
</style>
