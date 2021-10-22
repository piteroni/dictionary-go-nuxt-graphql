<template>
  <div class="heading-container px-12 w-full flex flex-wrap content-center justify-between">
    <div style="width: 30px" class="nav flex flex-wrap content-center">
      <div v-if="hasPrev" @click="moveToPrev" class="nav-button flex flex-wrap justify-center content-center">
        <img width="12px" src="~/assets/image/prev.png" alt="prev">
      </div>
    </div>

    <div class="pokemon-heading flex flex-wrap content-center">
      <div class="fixed-aria">
        <img v-show="imageURL !== ''" height="338px" width="338px" :src="imageURL" alt="image-of-pokemon">
      </div>

      <div class="flex flex-wrap content-center">
        <div class="pokemon-abstract">
          <div class="national-no">
            {{ nationalNo }}
          </div>

          <div class="pokemon-name">
            {{ name }}
          </div>

          <div class="flex mt-6">
            <img
              class="mr-2"
              v-for="(gender, key) in genders"
              :key="key"
              :src="gender.iconURL"
              :alt="gender.name"
            >
          </div>
        </div>
      </div>
    </div>

    <div style="width: 30px" class="nav flex flex-wrap content-center">
      <div v-if="hasNext" @click="moveToNext" class="nav-button flex flex-wrap justify-center content-center">
        <img width="12px" src="~/assets/image/next.png" alt="next">
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, inject, useRouter } from "@nuxtjs/composition-api"
import { pokemonDetailsKey } from "@/composables/pokemonDetails"

export default defineComponent({
  setup() {
    const pokemon = inject(pokemonDetailsKey)!!

    const router = useRouter()

    const formatNationalNo = (nationalNo: number) => {
      const formated = ("000" + nationalNo.toString()).slice(-3)

      return `No.${formated}`
    }

    const nationalNo = computed(() => formatNationalNo(pokemon.nationalNo.value))

    function moveToPrev() {
      if (!pokemon.linkInfo.value.hasPrev) {
        return
      }

      router.push(`/pokemons/${pokemon.linkInfo.value.prevNationalNo}`)
    }

    function moveToNext() {
      if (!pokemon.linkInfo.value.hasNext) {
        return
      }

      router.push(`/pokemons/${pokemon.linkInfo.value.nextNationalNo}`)
    }

    const hasPrev = computed(() => pokemon.linkInfo.value.hasPrev)
    const hasNext = computed(() => pokemon.linkInfo.value.hasNext)

    return {
      nationalNo: nationalNo,
      name: pokemon.name,
      imageURL: pokemon.imageURL,
      genders: pokemon.genders,
      hasPrev,
      hasNext,
      moveToPrev,
      moveToNext,
    }
  }
})
</script>

<style scoped>
.heading-container {
  height: 360px;
}

.nav {
  height: 360px;
}

.nav-button {
  height: 130px;
  cursor: pointer;
  padding-left: 8px;
  padding-right: 8px;
  border-radius: 2px;
  background-color: rgb(255, 255, 255);
  border: 2px solid rgb(204, 204, 204);
}

.pokemon-heading {
  height: 360px;
  padding-left: 160px;
  padding-right: 80px;
}

.pokemon-abstract {
  padding-top: 30px;
  padding-left: 30px;
  margin-left: 36px;
  background-color: white;
  height: 220px;
  width: 480px;
  border-radius: 10px;
  box-shadow: 8px 8px 0 #d9d9d9;
}

.national-no {
  font-size: 20px;
  font-weight: 700;
}

.pokemon-name {
  font-size: 28px;
  font-weight: 700;
}

.fixed-aria {
  height: 338px;
  width: 338px;
}
</style>
