<template>
  <div>
    <app-header>
      <logo />
    </app-header>

    <pokemon-heading />

    <pokemon-details />

    <evolution-table />

    <app-footer />
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "nuxt-property-decorator"
import { PokemonDocument, PokemonQuery, PokemonQueryVariables } from "@/graphql/generated"
import { PokemonQueryType } from "@/store/pokemonDataset"
import { HttpStatusCode } from "@/shared/http"
import Header from "@/components/singletons/Header.vue"
import Logo from "@/components/singletons/Logo.vue"
import PokemonHeading from "@/components/tightly-coupled/pokemons/_id/PokemonHeading.vue"
import Details from "@/components/tightly-coupled/pokemons/_id/PokemonDetails.vue"
import EvolutionTable from "@/components/tightly-coupled/pokemons/_id/EvolutionTable.vue"
import Footer from "@/components/singletons/Footer.vue"

@Component({
  components: {
    "app-header": Header,
    "logo": Logo,
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

    let response

    try {
      response = await this.$apollo.query<PokemonQuery, PokemonQueryVariables>({
        query: PokemonDocument,
        variables: { pokemonId }
      })
    } catch (e) {
      console.error(e)

      return this.$nuxt.error({
        statusCode: HttpStatusCode.INTERNAL_SERVER_ERROR
      })
    }

    if (response.data.pokemon.__typename == "PokemonNotFound") {
      return this.$nuxt.error({ statusCode: HttpStatusCode.NOT_FOUND })
    }

    this.$accessor.pokemonDataset.save({
      pokemon: response.data.pokemon as PokemonQueryType<"pokemon", "Pokemon">,
      evolutions: response.data.evolutions as PokemonQueryType<"evolutions", "Evolutions">,
      pageInfo: response.data.pageInfo as PokemonQueryType<"pageInfo", "PageInfo">
    })
  }
}
</script>

<style scoped>
.site-logo {
  width: 192px;
}
</style>
