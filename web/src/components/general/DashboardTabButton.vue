<template>
  <v-tab :value="tab">
    <v-icon>{{ tab.icon }}</v-icon>
    {{ tab.title }}
  </v-tab>
</template>

<script>
export default {
  props: {
    tab: {
      type: Object,
      required: true
    }
  }
};
</script>