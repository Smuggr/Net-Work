<template>
  <v-window-item :value="value">
    <feed>
      <post />
    </feed>
  </v-window-item>
</template>

<script>
export default {
  name: 'HomeTab',
  props: {
    value: {
      type: Object,
      required: true
    }
  },
}
</script>
