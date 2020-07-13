<template>
  <header>
    <v-app-bar
      app
      hide-on-scroll
    >
      <v-app-bar-nav-icon @click="isEnabled = !isEnabled" />
      <v-toolbar-title>ブイトーーク！</v-toolbar-title>
    </v-app-bar>

    <vtalk-menu :is-enabled="isEnabled" />
  </header>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";

@Component
export default class VtalkHeader extends Vue {
  isEnabled: Boolean = false;
}
</script>
