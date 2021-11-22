<template>
  <div>
    <app-header>
      <logo />
    </app-header>

    <app-footer />
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "nuxt-property-decorator"
import Header from "@/components/singletons/Header.vue"
import Logo from "@/components/singletons/Logo.vue"
import Footer from "@/components/singletons/Footer.vue"

@Component({
  components: {
    "app-header": Header,
    "logo": Logo,
    "app-footer": Footer
  }
})
export default class PokemonList extends Vue {
}
</script>
