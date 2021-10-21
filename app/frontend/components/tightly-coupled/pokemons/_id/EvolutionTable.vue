<template>
  <div>
    <div class="table-container">
      <div class="caption flex justify-center">
        <img style="margin: 0 auto;" width="400px" src="~/assets/image/ttl_evolution.svg" alt="evolution-caption">
      </div>

      <div class="evolution-table">
        <div class="flex" v-if="isPresent">
          <div class="mr-8" v-for="(pokemon, key) in evolutions" :key="key">
            <img class="pokemon mb-2" width="290px" height="290px" :src="pokemon.imageURL" :alt="pokemon.name">

            <p class="national-no">
              {{ formatNationalNo(pokemon.nationalNo) }}
            </p>

            <p class="pokemon-name mb-1">
              {{ pokemon.name }}
            </p>

            <div class="flex">
              <type
                v-for="(type, key) in pokemon.types"
                :key="key"
                :iconURL="type.iconURL"
                :name="type.name"
              />
            </div>
          </div>
        </div>

        <div v-else>
          しんかしない
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, inject } from "@nuxtjs/composition-api"
import { pokemonDetailsKey } from "@/composables/pokemonDetails"
import Type from "@/components/basic/Type.vue"

export default defineComponent({
  components: {
    "type": Type
  },
  setup() {
    const pokemon = inject(pokemonDetailsKey)!!

    const formatNationalNo = (nationalNo: number) => {
      const formated = ("000" + nationalNo.toString()).slice(-3)

      return `No.${formated}`
    }

    const isPresent = computed(() => pokemon.evolutions.value.length !== 0)

    return {
      evolutions: pokemon.evolutions,
      formatNationalNo,
      isPresent
    }
  }
})
</script>

<style scoped>
.table-container {
  position: relative;
  width: 1000px;
  height: 520px;
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

.evolution-table {
  padding: 80px 50px 30px;
}

.pokemon {
  border: solid 1.5px #d9d9d9;
  border-radius: 5px;
}

.national-no {
  font-size: 16px;
  font-weight: 700;
}

.pokemon-name {
  font-size: 22px;
  font-weight: 700;
}
</style>
