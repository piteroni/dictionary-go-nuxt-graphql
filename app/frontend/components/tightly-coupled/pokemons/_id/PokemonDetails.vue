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
      <div class="details-container ability">
        <div class="flex flex-wrap content-center justify-center mb-8">
          <div class="details-key w-20">
            HP
          </div>

          <div class="details-value flex ml-8">
            <div class="ability-value" :class="value" v-for="(value, key) in heartGauge" :key="key">
            </div>
          </div>
        </div>

        <div class="flex flex-wrap content-center justify-center mb-8">
          <div class="details-key w-20">
            こうげき
          </div>

          <div class="details-value flex ml-8">
            <div class="ability-value" :class="value" v-for="(value, key) in attackGauge" :key="key">
            </div>
          </div>
        </div>

        <div class="flex flex-wrap content-center justify-center mb-8">
          <div class="details-key w-20">
            ぼうぎょ
          </div>

          <div class="details-value flex ml-8">
            <div class="ability-value" :class="value" v-for="(value, key) in defenseGauge" :key="key">
            </div>
          </div>
        </div>

        <div class="flex flex-wrap content-center justify-center mb-8">
          <div class="details-key w-20">
            とくこう
          </div>

          <div class="details-value flex ml-8">
            <div class="ability-value" :class="value" v-for="(value, key) in specialAttackGauge" :key="key">
            </div>
          </div>
        </div>

        <div class="flex flex-wrap content-center justify-center mb-8">
          <div class="details-key w-20">
            とくぼう
          </div>

          <div class="details-value flex ml-8">
            <div class="ability-value" :class="value" v-for="(value, key) in specialDefenseGauge" :key="key">
            </div>
          </div>
        </div>

        <div class="flex flex-wrap content-center justify-center">
          <div class="details-key w-20">
            すばやさ
          </div>

          <div class="details-value flex ml-8">
            <div class="ability-value" :class="value" v-for="(value, key) in speedGauge" :key="key">
            </div>
          </div>
        </div>
      </div>

      <div v-if="description.text !== ''" class="details-container pokemon-description mt-8">
        {{ description.text }} （{{ description.series }}）
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, inject, readonly } from "@nuxtjs/composition-api"
import { pokemonDetailsKey, abilityMaxStatus } from "@/composables/pokemonDetails"
import { Ability } from "@/graphql/generated/client"

export default defineComponent({
  setup() {
    const { pokemon } = inject(pokemonDetailsKey)!!

    const evaluation = 15

    const scaleAbility = (value: number, maxValue: number, evaluation: number): number => {
      return Math.round(evaluation * (value / maxValue))
    }

    const evaluateAbility = (ability: Ability): Ability => {
      return {
        heart: scaleAbility(ability.heart, abilityMaxStatus.heart, evaluation),
        attack: scaleAbility(ability.attack, abilityMaxStatus.attack, evaluation),
        defense: scaleAbility(ability.defense, abilityMaxStatus.defense, evaluation),
        specialAttack: scaleAbility(ability.specialAttack, abilityMaxStatus.specialAttack, evaluation),
        specialDefense: scaleAbility(ability.specialDefense, abilityMaxStatus.specialDefense, evaluation),
        speed: scaleAbility(ability.speed, abilityMaxStatus.speed, evaluation),
      }
    }

    const valueState = readonly({
      on: "on",
      off: "off"
    })

    const generateAbilityValue = (value: number, evaluation: number): string[] => {
      const gauge = []

      if (value < 0 || value > evaluation) {
        throw Error(`incorrect ability value, value = ${value}`)
      }

      for (let i = 0; i < evaluation; i++) {
        if (value !== 0) {
          gauge.push(valueState.on)
          value--
        } else  {
          gauge.push(valueState.off)
        }
      }

      return gauge
    }

    const ability = evaluateAbility(pokemon.ability)

    const heartGauge = generateAbilityValue(ability.heart, evaluation)
    const attackGauge = generateAbilityValue(ability.attack, evaluation)
    const defenseGauge = generateAbilityValue(ability.defense, evaluation)
    const specialAttackGauge = generateAbilityValue(ability.specialAttack, evaluation)
    const specialDefenseGauge = generateAbilityValue(ability.specialDefense, evaluation)
    const speedGauge = generateAbilityValue(ability.speed, evaluation)

    return {
      species: pokemon.species,
      weight: pokemon.weight,
      height: pokemon.height,
      types: pokemon.types,
      characteristics: pokemon.characteristics,
      description: pokemon.description,
      heartGauge,
      attackGauge,
      defenseGauge,
      specialAttackGauge,
      specialDefenseGauge,
      speedGauge,
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
