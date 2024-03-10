<template>
  <v-card
    class="post"
    elevation="10"
  >
    <v-img
      height="250"
      src="https://cdn.vuetifyjs.com/images/cards/cooking.png"
      cover
    ></v-img>

    <v-card-item>
      <v-card-title>V1.0.0 Released</v-card-title>

      <v-card-subtitle>
        <span class="me-1">Smuggr</span>

        <v-icon
          icon="mdi-crown"
          size="small"
        ></v-icon>

      </v-card-subtitle>
    </v-card-item>

    <v-card-text>
      <div>Network V1.0.0 is now available on github!</div>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  props: {
    post: {
      type: Object,
      required: true
    },
  },
};
</script>

<style scoped>
.post {
  max-width: 100%;
  max-height: 380px;
  margin-bottom: 64px;
}
</style>