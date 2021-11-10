<template>
  <div>
    <div class="mr-3">
      <img class="mx-auto" height="36px" width="36px" :src="iconURL" :alt="name">
      <div class="font-semibold flex justify-center" style="font-size: 5px;">
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
export default class Type extends Vue {
  @Prop({ required: true, type: String })
  public iconURL!: string

  @Prop({ required: true, type: String })
  public name!: string
}
</script>
