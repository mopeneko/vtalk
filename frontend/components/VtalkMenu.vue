<template>
  <v-navigation-drawer
    v-model="isEnabled"
    app
  />
</template>

<script lang=ts>
import { Vue, Component, Prop } from "vue-property-decorator";

@Component
export default class VtalkMenu extends Vue {
  @Prop({ type: Boolean, required: true }) readonly isEnabled!: Boolean;
}
</script>
