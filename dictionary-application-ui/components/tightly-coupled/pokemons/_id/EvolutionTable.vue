<template>
  <div>
    <div class="table-container">
      <div class="caption flex justify-center">
        <img style="margin: 0 auto;" width="400px" src="~/assets/image/ttl_evolution.svg" alt="evolution-caption">
      </div>

      <div class="evolution-table">
        <div v-if="doneLoad" class="flex justify-center">
          <div v-for="(pokemon, key) in evolutions" :key="key" class="mr-2">
            <div class="flex items-center">
              <img class="pokemon mb-2" width="290px" height="290px" :src="pokemon.imageURL" :alt="pokemon.name">

              <div class="ml-2" style="height: 40px;">
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
              <type
                v-for="(type, typeKey) in pokemon.types"
                :key="typeKey"
                :icon-u-r-l="type.iconURL"
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
import { PokemonQuery } from "@/graphql/generated/client"
import Type from "@/components/basic/Type.vue"
import { nationalNoToText } from "~/store/pokemonDataset"

@Component({
  components: {
    "type": Type
  }
})
export default class EvolutionTable extends Vue {
  public get doneLoad(): boolean {
    return this.$accessor.pokemonDataset.nationalNo !== 0
  }

  public get name(): string {
    return this.$accessor.pokemonDataset.name
  }

  // public get canEvolution(): boolean {
  //   return this.$accessor.pokemonDataset.ca
  // }

  public get nationalNo(): string {
    return this.$accessor.pokemonDataset.nationalNoText
  }

  public get evolutions() {
    return this.$accessor.pokemonDataset.evolutions.map(p => {
      return {...p, nationalToText: nationalNoToText(p.nationalNo)}
    })
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
.evolution-table {
  padding: 80px 20px 30px;
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
.evolution-allow {
  width: 0;
  height: 0;
  border-style: solid;
  border-width: 20px 0 20px 20px;
  border-color: transparent transparent transparent #d8dfe6;
}
</style>
