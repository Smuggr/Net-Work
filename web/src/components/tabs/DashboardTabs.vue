<template>
  <v-window-item :value="value" v-model="selectedSubTabName">
    <v-tabs
      v-model="selectedSubTabName"
      align-tabs="center"
      color="deeep-purple-accent-4"
      mandatory
      stacked
    >
      <slot name="buttons" />
    </v-tabs>

    <v-window
      v-model="selectedSubTabName"
      direction="horizontal"
      elevation="0"
      mandatory
    >
      <slot name="content" />
    </v-window>
  </v-window-item>
</template>

<script>
export default {
  name: 'DashboardTab',
  props: {
    value: {
      type: Object,
      required: true,
    },
    childValue: {
      type: Object,
      required: true,
    },
  },
  computed: {
    selectedSubTabName: {
      get() {
        console.log('getting dashboard tab name ', this.value);
        return this.childValue;
      },
      set(newValue) {
        if (this.childValue && newValue !== this.childValue) {
          console.log('updating dashboard tab name ', newValue);
          this.$emit('update:childValue', newValue );
        }
      },
    },
  },
  mounted() {
    console.log('DashboardTab component received value and childValue props:', this.value, this.childValue);
  }
};
</script>

<style scoped>
/* Your scoped styles here */
</style>
