<template>
  <div>
    <div class="table-container">
      <div class="caption flex justify-center">
        <img class="caption-image mx-auto" src="~/assets/image/ttl_evolution.svg" alt="evolution-caption">
      </div>

      <div class="evolution-table">
        <div v-if="doneLoad" class="flex justify-center">
          <div v-for="(pokemon, key) in evolutions" :key="key" @click="() => showPokemon(pokemon.id)" class="mr-2">
            <div class="flex items-center">
              <img class="pokemon mb-2" :src="pokemon.imageURL" :alt="pokemon.name">

              <div class="h-10 ml-2">
                <div v-if="pokemon.canEvolution" class="evolution-allow" />
              </div>
            </div>

            <p class="national-no">
              {{ pokemon.nationalToText }}
            </p>

            <p class="pokemon-name mb-1">
              {{ pokemon.name }}
            </p>

            <div class="flex">
              <pokemon-type
                v-for="(type, typeKey) in pokemon.types"
                :key="typeKey"
                :icon-url="type.iconURL"
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
import { Vue, Component } from "nuxt-property-decorator"
import { nationalNoToText } from "@/store/pokemonDataset"
import PokemonType from "@/components/basic/PokemonType.vue"

@Component({
  components: {
    "pokemon-type": PokemonType
  }
})
export default class EvolutionTable extends Vue {
  public get doneLoad(): boolean {
    return this.$accessor.pokemonDataset.nationalNo !== 0
  }

  public get name(): string {
    return this.$accessor.pokemonDataset.name
  }

  public get nationalNo(): string {
    return this.$accessor.pokemonDataset.nationalNoText
  }

  public get evolutions() {
    return this.$accessor.pokemonDataset.evolutions.map(pokemon => {
      return {
        ...pokemon,
        nationalToText: nationalNoToText(pokemon.nationalNo)
      }
    })
  }

  public showPokemon(pokemonId: number): void {
    this.$router.push(`/pokemons/${pokemonId}`)
  }
}
</script>

<style scoped>
.table-container {
  position: relative;
  width: 1100px;
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
.caption-image {
  width: 25rem;
}
.evolution-table {
  padding: 80px 20px 30px;
}
.pokemon {
  width: 18.125rem;
  border: solid 1.5px #d9d9d9;
  border-radius: 5px;
  cursor: pointer;
  transition: .3s;
}
.pokemon:hover {
  opacity: 0.8;
}
.national-no {
  font-size: 16px;
  font-weight: 700;
}
.pokemon-name {
  font-size: 22px;
  font-weight: 700;
}
.evolution-allow {
  width: 0;
  height: 0;
  border-style: solid;
  border-width: 20px 0 20px 20px;
  border-color: transparent transparent transparent #d8dfe6;
}
</style>
