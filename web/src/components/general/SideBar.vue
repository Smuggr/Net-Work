<template>
  <v-navigation-drawer
    app
    elevation="10"
    scrim="false"
    class="sidebar"
    
  >
    <template v-slot:prepend>
      <v-list color="transparent" class="text-center">
        <slot name="primary"></slot>
      </v-list>
    </template>

    <template v-slot:append>
      <v-divider />
      
      <v-list color="transparent" class="text-center">
        <slot name="secondary"></slot>
      </v-list>
    </template>
  </v-navigation-drawer>
</template>

<script>
export default {
  name: 'SideBar',
}
</script>

<style scoped>
.sidebar {
  height: 100%;
  overflow-y: auto;
}
</style>