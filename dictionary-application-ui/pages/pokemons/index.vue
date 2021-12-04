<template>
  <div>
    <app-header>
      <site-logo />
    </app-header>

    <pokemon-list />

    <!-- <app-footer /> -->
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "nuxt-property-decorator"
import { PokemonsDocument, PokemonsQuery, PokemonsQueryVariables } from "@/graphql/generated"
import { HttpStatusCode } from "@/shared/http"
import { PokemonsQueryType } from "@/store/pokemonConnection"
import Header from "@/components/singletons/Header.vue"
import Logo from "@/components/singletons/Logo.vue"
import PPP from "@/components/tightly-coupled/pokemons/index/PokemonList.vue"
import Footer from "@/components/singletons/Footer.vue"

@Component({
  components: {
    "app-header": Header,
    "site-logo": Logo,
    "pokemon-list": PPP,
    "app-footer": Footer
  },
  async fetch({ app, error }) {
    let response

    const initial = 64

    try {
      response = await app.apolloProvider!!.defaultClient.query<PokemonsQuery, PokemonsQueryVariables>({
        query: PokemonsDocument,
        variables: { after: null, first: initial }
      })
    } catch (e) {
      console.error(e)

      return error({
        statusCode: HttpStatusCode.INTERNAL_SERVER_ERROR
      })
    }

    switch (response.data.pokemons.__typename) {
      case "PokemonConnection":
        return app.$accessor.pokemonConnection.save(
          response.data.pokemons as PokemonsQueryType<"pokemons", "PokemonConnection">
        )
      case "PokemonNotFound":
        return error({
          statusCode: HttpStatusCode.NOT_FOUND
        })
      case "IllegalArguments":
        return error({
          statusCode: HttpStatusCode.INTERNAL_SERVER_ERROR
        })
      default:
        return error({
          statusCode: HttpStatusCode.INTERNAL_SERVER_ERROR
        })
    }
  }
})
export default class PokemonList extends Vue {
}
</script>
