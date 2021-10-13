<template>
  <div class="heading-container px-12 w-full flex flex-wrap content-center justify-between">
    <div class="nav flex flex-wrap content-center">
      <div class="nav-button flex flex-wrap justify-center content-center">
        <img width="12px" src="~/assets/image/prev.png" alt="prev">
      </div>
    </div>

    <div class="pokemon-heading flex flex-wrap content-center">
      <div>
        <img height="338px" width="338px" :src="pokemonImageURL" alt="image-of-pokemon">
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

    <div class="nav flex flex-wrap content-center">
      <div class="nav-button flex flex-wrap justify-center content-center">
        <img width="12px" src="~/assets/image/next.png" alt="next">
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@nuxtjs/composition-api"

export type PokemonGender = {
  name: string
  iconURL: string
}

export default defineComponent({
  setup() {
    const nationalNo = "No.001"
    const name = "フシギダネ"

    const pokemonImageName = "afa02eaba4c39820fc57f4e8abaeea80.png"
    const pokemonImageURL = require(`~/assets/image/${pokemonImageName}`)

    const genders: PokemonGender[] = [
      {
        name: "male",
        iconURL: require(`~/assets/image/icon_male.svg`)
      },
      {
        name: "female",
        iconURL: require(`~/assets/image/icon_female.svg`)
      }
    ]

    return {
      nationalNo,
      name,
      pokemonImageURL,
      genders
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
</style>