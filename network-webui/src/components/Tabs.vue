<template>
  <v-card>
    <v-tabs
      v-model="selectedTabName"
      align-tabs="center"
      color="deep-purple-accent-4"
      hide-slider
    >
      <slot name="tabs"></slot>
    </v-tabs>
    <v-window v-model="selectedTabName">
      <slot name="content"></slot>
    </v-window>
  </v-card>
</template>

<script>
export default {
  props: {
    value: Object,
  },
  computed: {
    selectedTabName: {
      get() {
        return this.value;
      },
      set(newValue) {
        if (this.value && newValue !== this.value) {
          this.$emit('update:value', { ...this.value, tabName: newValue });
        }
      }
    }
  },
};
</script>
