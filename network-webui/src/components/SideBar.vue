<template>
  <v-navigation-drawer
    v-model="drawer"
    app
    temporary
    class="sidebar"
  >
    <v-list color="transparent" class="text-center">
      <slot name="primary"></slot>
    </v-list>

    <template v-slot:append>
      <v-divider inset></v-divider>
      <slot name="secondary"></slot>
    </template>
  </v-navigation-drawer>
</template>

<script>
export default {
  props: {
    value: Boolean
  },
  computed: {
    drawer: {
      get() {
        return this.value;
      },
      set(val) {
        this.$emit('input', val);
      }
    }
  }
};
</script>

<style scoped>
.sidebar {
  width: 100%;
}
</style>
