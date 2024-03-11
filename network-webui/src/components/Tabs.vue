<template>
  <v-card>
    <v-tabs
      :value="selectedTab"
      @update:value="updateSelectedTab"
      align-tabs="center"
      color="deep-purple-accent-4"
      hide-slider
    >
      <v-tab v-for="tab in tabs" :key="tab.id" :value="tab.id">{{ tab.name }}</v-tab>
    </v-tabs>
    <v-window v-model="selectedTab">
      <v-window-item v-for="tab in tabs" :key="tab.id" :value="tab.id">
        {{ tab.content }}
      </v-window-item>
    </v-window>
  </v-card>
</template>

<script>
export default {
  props: {
    value: Number, // Change to match your prop name and type
    tabs: Array // Assuming you have an array of tabs with 'id' and 'name' properties
  },
  data() {
    return {
      selectedTab: this.value // Initialize selectedTab with the prop value
    };
  },
  methods: {
    updateSelectedTab(newValue) {
      this.selectedTab = newValue;
      this.$emit('update:value', newValue); // Emit event to update prop value
    }
  }
};
</script>
