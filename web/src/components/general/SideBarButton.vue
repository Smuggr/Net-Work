<template>
  <v-list-item>
    <v-btn
      :size="size"
      :prepend-icon="destination.icon"
      block
      variant="text"
      rounded="xl"
      class="justify-start"
      @click="handleButtonClick"
    >
      {{ destination.title }}
    </v-btn>
  </v-list-item>
</template>

<script>
export default {
  name: 'SideBarButton',
  props: {
    size: {
      type: String,
      default: 'x-large',
      validator: (value) => ['x-small', 'small', 'medium', 'large', 'x-large'].includes(value),
    },
    destination: {
      type: Object,
      required: true,
    },
  },
  methods: {
    handleButtonClick() {
      this.$emit('button-click', this);
    },
  },
};
</script>