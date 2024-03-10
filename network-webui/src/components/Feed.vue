<template>
  <v-container class="bg-surface-variant mb-6">
    <v-col>
      <slot />
    </v-col>
  </v-container>
</template>

<script>
export default {
  props: {
    posts: {
      type: Array,
      required: true
    },
  },
};
</script>