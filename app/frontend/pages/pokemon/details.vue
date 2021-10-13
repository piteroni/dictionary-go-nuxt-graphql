<template>
  <div class="site">
    <app-header>
      <h1 class="site-logo">
        <img src="~/assets/image/logo.svg" alt="site-logo">
      </h1>
    </app-header>

    <heading />

    <app-details />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@nuxtjs/composition-api"
import Header from "@/components/singleton/Header.vue"
import Heading from "@/components/tightly-coupled/pokemon/details/Heading.vue"
import Details from "@/components/tightly-coupled/pokemon/details/Details.vue"

export default defineComponent({
  components: {
    "app-header": Header,
    "heading": Heading,
    "app-details": Details
  },
})
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