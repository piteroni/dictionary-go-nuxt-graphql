<template>
  <div class="heading-container px-12 w-full flex flex-wrap content-center justify-between">
    <div class="w-8 nav flex flex-wrap content-center">
      <div v-if="hasPrev" class="nav-button flex flex-wrap justify-center content-center" @click="moveToPrev">
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
              v-for="(gender, key) in genders"
              :key="key"
              class="mr-2"
              :src="gender.iconURL"
              :alt="gender.name"
            >
          </div>
        </div>
      </div>
    </div>

    <div class="w-8 nav flex flex-wrap content-center">
      <div v-if="hasNext" class="nav-button flex flex-wrap justify-center content-center" @click="moveToNext">
        <img width="12px" src="~/assets/image/next.png" alt="next">
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "nuxt-property-decorator"
import { Gender } from "@/graphql/generated/client"

@Component
export default class PokemonHeading extends Vue {
  public get nationalNo(): string {
    return this.$accessor.pokemonDataset.nationalNoText
  }

  public get name(): string {
    return this.$accessor.pokemonDataset.name
  }

  public get genders(): Gender[] {
    return this.$accessor.pokemonDataset.genders
  }

  public get imageURL(): string {
    return this.$accessor.pokemonDataset.imageURL
  }

  public get hasPrev(): boolean {
    return this.$accessor.pokemonDataset.linkInfo.hasPrev
  }

  public get hasNext(): boolean {
    return this.$accessor.pokemonDataset.linkInfo.hasNext
  }

  public moveToPrev() {
    if (!this.$accessor.pokemonDataset.linkInfo.hasPrev) {
      return
    }

    const id = this.$accessor.pokemonDataset.linkInfo.prevNationalNo

    this.$router.push(`/pokemons/${id}`)
  }

  public moveToNext() {
    if (!this.$accessor.pokemonDataset.linkInfo.hasNext) {
      return
    }

    const id = this.$accessor.pokemonDataset.linkInfo.nextNationalNo

    this.$router.push(`/pokemons/${id}`)
  }
}
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
