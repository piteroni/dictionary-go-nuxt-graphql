<template>
  <div>
    <slot />
  </div>
</template>

<script>
import { defineComponent, provide } from "@nuxtjs/composition-api"
import { pokemonDetailsKey, usePokemonDetails } from "./pokemonDetails"

export default defineComponent({
  setup() {
    provide(pokemonDetailsKey, usePokemonDetails())
    return {}
  }
})
</script>
