<template>
  <v-container
    class="feed"
  >
    <v-col>
      <slot />
    </v-col>
  </v-container>
</template>

<script>
export default {

};
</script>

<style scoped>
.feed {
  max-width: 50%;
  min-width: 480px;
}
</style>