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
              <pokemon-type
                v-for="(type, key) in types"
                :key="key"
                :icon-url="type.iconURL"
                :name="type.name"
              />
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
            <div v-for="(characteristic, key) in characteristics" :key="key" class="mr-4">
              {{ characteristic.name }}
            </div>
          </div>
        </div>
      </div>

      <div class="details-container ability">
        <div class="flex flex-wrap content-center justify-center mb-8">
          <div class="details-key w-20">
            HP
          </div>

          <div class="details-value flex ml-8">
            <div v-for="(value, key) in heartGauge" :key="key" class="ability-value" :class="value" />
          </div>
        </div>

        <div class="flex flex-wrap content-center justify-center mb-8">
          <div class="details-key w-20">
            こうげき
          </div>

          <div class="details-value flex ml-8">
            <div v-for="(value, key) in attackGauge" :key="key" class="ability-value" :class="value" />
          </div>
        </div>

        <div class="flex flex-wrap content-center justify-center mb-8">
          <div class="details-key w-20">
            ぼうぎょ
          </div>

          <div class="details-value flex ml-8">
            <div v-for="(value, key) in defenseGauge" :key="key" class="ability-value" :class="value" />
          </div>
        </div>

        <div class="flex flex-wrap content-center justify-center mb-8">
          <div class="details-key w-20">
            とくこう
          </div>

          <div class="details-value flex ml-8">
            <div v-for="(value, key) in specialAttackGauge" :key="key" class="ability-value" :class="value" />
          </div>
        </div>

        <div class="flex flex-wrap content-center justify-center mb-8">
          <div class="details-key w-20">
            とくぼう
          </div>

          <div class="details-value flex ml-8">
            <div v-for="(value, key) in specialDefenseGauge" :key="key" class="ability-value" :class="value" />
          </div>
        </div>

        <div class="flex flex-wrap content-center justify-center">
          <div class="details-key w-20">
            すばやさ
          </div>

          <div class="details-value flex ml-8">
            <div v-for="(value, key) in speedGauge" :key="key" class="ability-value" :class="value" />
          </div>
        </div>
      </div>

      <div class="details-container pokemon-description mt-8">
        {{ descriptionText }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "nuxt-property-decorator"
import { abilityMaxStatus } from "@/store/pokemonDataset"
import { Characteristic, Type } from "@/graphql/generated"
import PokemonType from "@/components/basic/PokemonType.vue"

@Component({
  components: {
    "pokemon-type": PokemonType
  }
})
export default class PokemonDetails extends Vue {
  private evaluation = 15

  private valueState = {
    on: "on",
    off: "off"
  }

  public get species(): string {
    return this.$accessor.pokemonDataset.species
  }

  public get weight(): string {
    return this.$accessor.pokemonDataset.weight
  }

  public get height(): string {
    return this.$accessor.pokemonDataset.height
  }

  public get types(): Type[] {
    return this.$accessor.pokemonDataset.types
  }

  public get characteristics(): Characteristic[] {
    return this.$accessor.pokemonDataset.characteristics
  }

  public get descriptionText(): string {
    if (!this.$accessor.pokemonDataset.description.text) {
      return ""
    }

    return `${this.$accessor.pokemonDataset.description.text}（${this.$accessor.pokemonDataset.description.series}）`
  }

  public get heartGauge(): string[] {
    const ability = this.scaleAbility(this.$accessor.pokemonDataset.ability.heart, abilityMaxStatus.heart, this.evaluation)

    return this.generateAbilityValue(ability, this.evaluation)
  }

  public get attackGauge(): string[] {
    const ability = this.scaleAbility(this.$accessor.pokemonDataset.ability.attack, abilityMaxStatus.attack, this.evaluation)

    return this.generateAbilityValue(ability, this.evaluation)
  }

  public get defenseGauge(): string[] {
    const ability = this.scaleAbility(this.$accessor.pokemonDataset.ability.defense, abilityMaxStatus.defense, this.evaluation)

    return this.generateAbilityValue(ability, this.evaluation)
  }

  public get specialAttackGauge(): string[] {
    const ability = this.scaleAbility(this.$accessor.pokemonDataset.ability.specialAttack, abilityMaxStatus.specialAttack, this.evaluation)

    return this.generateAbilityValue(ability, this.evaluation)
  }

  public get specialDefenseGauge(): string[] {
    const ability = this.scaleAbility(this.$accessor.pokemonDataset.ability.specialDefense, abilityMaxStatus.specialDefense, this.evaluation)

    return this.generateAbilityValue(ability, this.evaluation)
  }

  public get speedGauge(): string[] {
    const ability = this.scaleAbility(this.$accessor.pokemonDataset.ability.speed, abilityMaxStatus.speed, this.evaluation)

    return this.generateAbilityValue(ability, this.evaluation)
  }

  private scaleAbility(value: number, maxValue: number, evaluation: number): number {
    return Math.round(evaluation * (value / maxValue))
  }

  private generateAbilityValue(value: number, evaluation: number): string[] {
    const gauge = []

    if (value < 0 || value > evaluation) {
      throw new Error(`incorrect ability value, value = ${value}`)
    }

    for (let i = 0; i < evaluation; i++) {
      if (value !== 0) {
        gauge.push(this.valueState.on)
        value--
      } else {
        gauge.push(this.valueState.off)
      }
    }

    return gauge
  }
}
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
}
.pokemon-description {
  width: 1014px;
  font-size: 20px;
  font-weight: 200;
}
.type-icon {
  font-size: 5px;
  font-weight: 600;
}
.ability {
  padding: 30px 25px;
  width: 500px;
}
.ability-value {
  border-radius: 12px;
  margin-right: 5px;
  height: 35px;
  width: 15px;
}
.ability-value.on {
  background-color: #fc0;
}
.ability-value.off {
  background-color: #f2f2f2;
}
</style>
