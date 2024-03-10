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

    <v-divider />

    <v-list color="transparent" class="text-center">
      <slot name="secondary"></slot>
    </v-list>
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