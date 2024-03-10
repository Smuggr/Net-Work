<template>
  <v-list-item>
    <v-btn
      :size="size"
      :prepend-icon="icon"
      block
      variant="text"
      rounded="xl"
      class="justify-start"
      @click="handleButtonClick"
    >
      {{ title }}
    </v-btn>
  </v-list-item>
</template>

<script>
export default {
  props: {
    size: {
      type: String,
      default: 'x-large',
      validator: (value) => ['x-small', 'small', 'medium', 'large', 'x-large'].includes(value),
    },
    title: {
      type: String,
      required: true,
    },
    icon: {
      type: String,
      required: true,
    },
    destination: {
      type: String,
      required: true,
    },
  },
  methods: {
    handleButtonClick() {
      this.$emit('button-click', this.destination);
    },
  },
};
</script>