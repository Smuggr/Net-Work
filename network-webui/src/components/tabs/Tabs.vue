<template>
  <v-card>
    <v-tabs
      v-model="selectedTabName"
      align-tabs="center"
      color="deep-purple-accent-4"
      elevation="0"
      mandatory
    >
      <slot name="tabs" />
    </v-tabs>

    <v-window
      v-model="selectedTabName"
      direction="horizontal"
      elevation="0"
      mandatory
    >
      <slot name="content" />
    </v-window>
  </v-card>
</template>

<script>
export default {
  name: 'Tabs',
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
