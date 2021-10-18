<template>
  <div class="bg-white">
    <div class="flex flex-wrap p-12 justify-center">
      <div class="details-container mr-8">
        <div class="flex flex-wrap content-center mb-8">
          <div class="details-key">
            分類：
          </div>

          <div class="details-value">
            {{ species }}
          </div>
        </div>

        <div class="flex flex-wrap content-center mb-8">
          <div class="details-key">
            タイプ：
          </div>

          <div class="details-value">
            <div class="flex">
              <div class="mr-6" v-for="(type, key) in types" :key="key">
                <img class="mx-auto" height="40px" width="40px" :src="type.iconURL" :alt="type.name">
                <div class="type-icon">
                  {{ type.name }}
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="flex mb-8">
          <div class="details-key">
            高さ：
          </div>

          <div class="details-value mr-6">
            {{ height }}
          </div>

          <div class="details-key">
            重さ：
          </div>

          <div class="details-value">
            {{ weight }}
          </div>
        </div>

        <div class="flex flex-wrap content-center mb-8">
          <div class="details-key">
            とくせい：
          </div>

          <div class="details-value">
            <div class="mr-4" v-for="(characteristic, key) in characteristics" :key="key">
              {{ characteristic.name }}
            </div>
          </div>
        </div>
      </div>
      <div class="details-container">
      </div>

      <div v-if="description.text !== ''" class="details-container pokemon-description mt-8">
        {{ description.text }} （{{ description.series }}）
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@nuxtjs/composition-api"
import { Characteristic, Description } from "~/graphql/generated/client"
import { PokemonType } from "./types"

type Props = {
  species: string,
  weight: string,
  height: string,
  types: PokemonType[],
  characteristics: Characteristic[],
  description: Description
}

export default defineComponent<Props, Props>({
  props: {
    species: {
      type: String,
      required: true
    },
    height: {
      type: String,
      required: true
    },
    weight: {
      type: String,
      required: true
    },
    types: {
      type: Array,
      required: true,
    },
    characteristics: {
      type: Array,
      required: true,
    },
    description: {
      type: Object,
      required: true
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
  font-size: 20px;
  font-weight: 200;
  text-align: center;
}

.type-icon {
  font-size: 5px;
  font-weight: 600;
}
</style>