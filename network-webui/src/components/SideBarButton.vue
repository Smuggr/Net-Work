<template>
  <v-list-item>
    <v-btn
      block
      :size="size"
      :prepend-icon="icon"
      variant="text"
      rounded="xl"
      :class="{ 'justify-start': order === 'top', 'justify-end': order === 'bottom' }"
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
    order: {
      type: String,
      default: 'top',
      validator: (value) => ['top', 'bottom'].includes(value),
    },
  },
};
</script>
