<template>
  <div>
    <div class="table-container">
      <div class="caption flex justify-center">
        <img style="margin: 0 auto;" width="400px" src="~/assets/image/ttl_evolution.svg" alt="evolution-caption">
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@nuxtjs/composition-api"

export default defineComponent({

})
</script>

<style scoped>
.table-container {
  position: relative;
  width: 1000px;
  height: 500px;
  margin: 90px auto;
  background-color: #fff;
  border-radius: 7px;
}

.caption {
  position: absolute;
  top: -40px;
  left: 0;
  right: 0;
  margin: auto;
}
</style>
