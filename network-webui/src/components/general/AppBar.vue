<template>
  <v-app-bar elevation="10" prominent>
    <v-app-bar-nav-icon @click.stop="$emit('toggle-drawer')"></v-app-bar-nav-icon>
    <v-app-bar-title>{{ title }}</v-app-bar-title>
    
    <v-progress-linear
        :active="isLoading"
        :indeterminate="true"
        location="bottom"
        color="deep-purple-accent-4"
        absolute
        bottom
    ></v-progress-linear>
  </v-app-bar>
</template>

<script>
export default {
  name: 'AppBar',
  props: {
    title: {
      type: String,
      required: true,
    },
    isLoading: {
      type: Boolean,
      default: false,
      required: true,
    },
  },
}
</script>
