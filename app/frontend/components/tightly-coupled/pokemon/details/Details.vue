<template>
  <div class="bg-white">
    <div class="flex flex-wrap p-12 justify-center">
      <div class="details-container mr-8">
        <div class="flex flex-wrap content-center">
          <div class="details-key">
            分類：
          </div>

          <div class="details-value">
            {{ species }}
          </div>
        </div>
      </div>
      <div class="details-container">
      </div>

      <div class="details-container pokemon-description mt-8">
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent} from "@nuxtjs/composition-api"

export default defineComponent({
  setup() {
    const species = "たねポケモン"

    return {
      species
    }
  }
})
</script>

<style scoped>
.details-container {
  width: 480px;
  padding: 40px 60px;
  border: 3px solid rgb(204, 204, 204);
  border-radius: 7px;
}

.details-key {
  font-size: 20px;
  font-weight: 700;
}

.details-value {
  font-size: 20px;
  font-weight: 200;
}

.pokemon-description {
  width: 994px;
}
</style>