<template>
  <v-card>
    <v-tabs
      v-model="selectedTab"
      align-tabs="center"
      color="deep-purple-accent-4"
      hide-slider
    >
      <slot name="tabs"></slot>
    </v-tabs>
    <v-window v-model="selectedTab">
      <slot name="content"></slot>
    </v-window>
  </v-card>
</template>

<script>
export default {
  props: {
    value: String,
  },
  data() {
    return {
      selectedTab: this.value
    };
  },
  watch: {
    value(newValue) {
      this.selectedTab = newValue;
    },
    selectedTab(newTab) {
      this.$emit('update:value', newTab);
    },
  },
};
</script>
