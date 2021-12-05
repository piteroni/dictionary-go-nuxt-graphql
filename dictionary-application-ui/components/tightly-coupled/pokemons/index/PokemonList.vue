<template>
  <div class="pt-14 bg-white">
    <div class="flex flex-wrap px-8">
      <div
        v-for="pokemon, key in pokemons" :key="key"
        class="cursor-pointer mb-12"
        @click="() => moveToDetails(pokemon.id)"
      >
        <img
          class="pokemon-image-border transition duration-300 hover:opacity-80 mr-2"
          :src="pokemon.imageURL"
          :alt="pokemon.name"
        >

        <p class="mt-1" style="max-width: 150px;">
          {{ pokemon.name }}
        </p>
      </div>
    </div>

    <app-footer v-if="isFetchAll" />
    <connection-loading v-else />
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "nuxt-property-decorator"
import ConnectionLoading from "./ConnectionLoading.vue"
import { fetchCount, PokemonsQueryType } from "@/store/pokemonConnection"
import { PokemonsDocument, PokemonsQuery, PokemonsQueryVariables } from "@/graphql/generated"
import { HttpStatusCode } from "@/shared/http"
import Footer from "@/components/singletons/Footer.vue"

@Component({
  components: {
    "connection-loading": ConnectionLoading,
    "app-footer": Footer
  },
})
export default class PokemenList extends Vue {
  public isFetchAll = false

  private isFetch = false

  public get pokemons() {
    return this.$accessor.pokemonConnection.pokemons
  }

  public mounted() {
    window.addEventListener("scroll", this.handleScroll)
  }

  public destroyed() {
    window.removeEventListener("scroll", this.handleScroll)
  }

  public moveToDetails(pokemonId: number): void {
    this.$router.push(`/pokemons/${pokemonId}`)
  }

  private async handleScroll(): Promise<void> {
    if (this.isFetchAll) {
      return
    }

    if (this.isFetch) {
      return
    }

    const bottom = document.body.clientHeight - window.innerHeight

    if (bottom <= window.scrollY) {
      this.isFetch = true
      await this.fetchConnection()
      this.isFetch = false
    }
  }

  private async fetchConnection() {
    const nextId = this.$accessor.pokemonConnection.nextID

    let response

    try {
      response = await this.$apollo.query<PokemonsQuery, PokemonsQueryVariables>({
        query: PokemonsDocument,
        variables: { after: nextId, first: fetchCount }
      })
    } catch { return }

    switch (response.data.pokemons.__typename) {
      case "PokemonConnection":
        return this.$accessor.pokemonConnection.accumulate(
          response.data.pokemons as PokemonsQueryType<"pokemons", "PokemonConnection">
        )
      case "PokemonNotFound":
        this.isFetchAll = true
        return
      case "IllegalArguments":
        return this.$nuxt.error({
          statusCode: HttpStatusCode.INTERNAL_SERVER_ERROR
        })
      default:
        return this.$nuxt.error({
          statusCode: HttpStatusCode.INTERNAL_SERVER_ERROR
        })
    }
  }
}
</script>

<style scoped>
.pokemon-image-border {
  width: 10.25rem;
  border: solid 1.5px #e8e8e8;
  border-radius: 5px;
}
</style>
