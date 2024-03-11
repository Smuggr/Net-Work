<template>
  <v-window
    v-model="selectedTabName"
    direction="horizontal"
    elevation="0"
    mandatory
  >
    <slot name="content" />
  </v-window>
</template>

<script>
export default {
  name: 'Tabs',
  props: {
    value: {
      type: Object,
      required: true,
    },
  },
  computed: {
    selectedTabName: {
      get() {
        console.log('getting tab name ', this.value);
        return this.value;
      },
      set(newValue) {
        if (this.value && newValue !== this.value) {
          console.log('updating tab name ', newValue);
          this.$emit('update:value', newValue);
        }
      }
    }
  },
  mounted() {
    console.log('Tabs component received value prop:', this.value);
  }
};
</script>