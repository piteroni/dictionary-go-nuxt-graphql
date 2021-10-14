<template>
  <div class="heading-container px-12 w-full flex flex-wrap content-center justify-between">
    <div class="nav flex flex-wrap content-center">
      <div class="nav-button flex flex-wrap justify-center content-center">
        <img width="12px" src="~/assets/image/prev.png" alt="prev">
      </div>
    </div>

    <div class="pokemon-heading flex flex-wrap content-center">
      <div class="fixed-aria">
        <img v-if="state.imageURL !== ''" height="338px" width="338px" :src="state.imageURL" alt="image-of-pokemon">
      </div>

      <div class="flex flex-wrap content-center">
        <div class="pokemon-abstract">
          <div class="national-no">
            {{ state.nationalNo }}
          </div>

          <div class="pokemon-name">
            {{ state.name }}
          </div>

          <div class="flex mt-6">
            <img
              class="mr-2"
              v-for="(gender, key) in state.genders"
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
import gql from "graphql-tag"
import { reactive, defineComponent } from "@nuxtjs/composition-api"
import { useQuery } from "@vue/apollo-composable"

export type PokemonGender = {
  name: string
  iconURL: string
}

export default defineComponent({
  setup() {
    const query = gql`
      query pokemon($pokemonId: Int!) {
        pokemon(pokemonId: $pokemonId) {
          nationalNo,
          name,
          imageName,
          genders {
            name,
            iconName
          }
        }
      }
    `

    type R = {
      pokemon: {
        nationalNo: number,
        name: string,
        imageName: string,
        genders: {name: string, iconName: string}[]
      }
    }

    const state = reactive<{
      nationalNo: string,
      name: string,
      imageURL: string,
      genders: PokemonGender[]
    }>({
      nationalNo: "",
      name: "",
      imageURL: "",
      genders: []
    })
    
    const { onResult } = useQuery<R>(query, {
      pokemonId: 1
    })

    if (process.client) {
      const format = (nationalNo: number) => ("000" + nationalNo.toString()).slice(-3)

      onResult(result => {
        if (!result.loading && !result.error) {
          const nationalNo = format(result.data.pokemon.nationalNo)

          const genders: PokemonGender[] = result.data.pokemon.genders.map(gender => {
            return {
              name: gender.name,
              iconURL: `/_nuxt/assets/image/${gender.iconName}`
            }
          })

          state.nationalNo = `No.${nationalNo}`
          state.name = result.data.pokemon.name
          state.imageURL = `/_nuxt/assets/image/${result.data.pokemon.imageName}`
          state.genders = genders
        }
      })
    }

    return { state }
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