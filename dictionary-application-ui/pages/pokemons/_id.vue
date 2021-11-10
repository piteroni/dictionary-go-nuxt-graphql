<template>
  <div class="site">
    <app-header>
      <h1>
        <img class="site-logo" src="~/assets/image/logo.svg" alt="site-logo">
      </h1>
    </app-header>

    <pokemon-heading />

    <pokemon-details />

    <evolution-table />

    <app-footer />
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "nuxt-property-decorator"
import { HttpStatusCode } from "@/shared/http"
import Header from "@/components/singletons/Header.vue"
import PokemonHeading from "@/components/tightly-coupled/pokemons/_id/PokemonHeading.vue"
import Details from "@/components/tightly-coupled/pokemons/_id/PokemonDetails.vue"
import EvolutionTable from "@/components/tightly-coupled/pokemons/_id/EvolutionTable.vue"
import Footer from "@/components/singletons/Footer.vue"

@Component({
  components: {
    "app-header": Header,
    "pokemon-heading": PokemonHeading,
    "pokemon-details": Details,
    "evolution-table": EvolutionTable,
    "app-footer": Footer
  },
  validate({ params }) {
    return /^\d+$/.test(params.id)
  },
})
export default class PokemonDetails extends Vue {
  public async fetch(): Promise<void> {
    const pokemonId = parseInt(this.$route.params.id)

    try {
      await this.$accessor.pokemonDataset.fetch({ pokemonId })
    } catch (e) {
      console.error(e)
      return this.$nuxt.error({
        statusCode: HttpStatusCode.INTERNAL_SERVER_ERROR
      })
    }
  }
}
</script>

<style scoped>
.site {
  background-color: rgb(242, 242, 242);
  background-image: url('~/assets/image/pattern.svg');
  background-repeat: repeat;
}
.site-logo {
  width: 192px;
}
</style>
