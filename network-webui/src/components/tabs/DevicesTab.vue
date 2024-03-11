<template>
  <v-window-item :value="value">
    <feed>
      <post />
    </feed>
  </v-window-item>
</template>

<script>
export default {
  name: 'DevicesTab',
  props: {
    value: {
      type: Object,
      required: true
    }
  },
}
</script>
