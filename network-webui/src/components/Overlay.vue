<template>
  <div class="text-center">
    <v-btn
      color="error"
      @click="overlay = !overlay"
    >
      Show Overlay
    </v-btn>

    <v-overlay v-model="overlay"></v-overlay>
  </div>
</template>


<script>
export default {
  data: () => ({
    overlay: false,
  }),
};
</script>