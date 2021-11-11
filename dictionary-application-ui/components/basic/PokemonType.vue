<template>
  <div>
    <div class="mr-3">
      <img class="mx-auto" height="36px" width="36px" :src="iconUrl" :alt="name">
      <div class="font-semibold text-xs flex justify-center">
        <span>
          {{ name }}
        </span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from "nuxt-property-decorator"

@Component
export default class PokemonType extends Vue {
  @Prop({ required: true, type: String })
  public iconUrl!: string

  @Prop({ required: true, type: String })
  public name!: string
}
</script>
