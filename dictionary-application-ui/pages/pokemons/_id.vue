<template>
  <div class="site">
    <app-header>
      <h1 class="site-logo">
        <img src="~/assets/image/logo.svg" alt="site-logo">
      </h1>
    </app-header>

    <pokemon-heading />

    <pokemon-details />

    <evolution-table />
  </div>
</template>

<script lang="ts">
import { defineComponent, useContext, useFetch, provide, inject } from "@nuxtjs/composition-api"
import { pokemonDetailsKey, usePokemonDetails } from "@/composables/pokemonDetails"
import Header from "@/components/singletons/Header.vue"
import PokemonHeading from "@/components/tightly-coupled/pokemons/_id/PokemonHeading.vue"
import PokemonDetails from "@/components/tightly-coupled/pokemons/_id/PokemonDetails.vue"
import EvolutionTable from "@/components/tightly-coupled/pokemons/_id/EvolutionTable.vue"

export default defineComponent({
  components: {
    "app-header": Header,
    "pokemon-heading": PokemonHeading,
    "pokemon-details": PokemonDetails,
    "evolution-table": EvolutionTable,
  },
  setup() {
    provide(pokemonDetailsKey, usePokemonDetails())

    const { pokemon, fetch } = inject(pokemonDetailsKey)!!

    const { route, error } = useContext()

    const pokemonId = parseInt(route.value.params.id)

    if (isNaN(pokemonId)) {
      error({ statusCode: 404 })

      return { pokemon }
    }

    useFetch(async () => {
      try {
        await fetch({ pokemonId })
      } catch (e) {
        console.error(e)
        error({ statusCode: 404 })
      }
    })

    return { pokemon }
  }
})
</script>

<style scoped>
.site {
  background-color: rgb(242, 242, 242);
  background-image: url('~/assets/image/pattern.svg');
  background-repeat: repeat;
  padding-bottom: 50px;
}

.site-logo {
  width: 192px;
}
</style>
