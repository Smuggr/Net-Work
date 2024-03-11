<template>
  <v-window-item :value="value">
    <feed>
      <post />
    </feed>
  </v-window-item>
</template>

<script>
export default {
  name: 'AboutTab',
  props: {
    value: {
      type: Object,
      required: true
    }
  },
}
</script>
