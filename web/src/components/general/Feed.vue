<template>
  <v-container class="feed">
    <v-col>
      <slot />
    </v-col>
  </v-container>
</template>

<script>
export default {
  name: 'Feed',
}
</script>

<style scoped>
.feed {
  max-width: 70%;
  min-width: 480px;
}
</style>