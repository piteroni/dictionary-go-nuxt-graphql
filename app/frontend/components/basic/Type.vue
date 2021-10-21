<template>
  <div>
    <div class="mr-3">
      <img class="mx-auto" height="36px" width="36px" :src="iconURL" :alt="name">
      <div class="font-semibold flex justify-center" style="font-size: 5px;">
        <span>
          {{ name }}
        </span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@nuxtjs/composition-api"

export default defineComponent({
  props: {
    iconURL: {
      type: String,
      required: true
    },
    name: {
      type: String,
      required: true
    }
  }
})
</script>
