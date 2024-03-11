<template>
  <v-window-item :value="tab">
    Balls Devices
  </v-window-item>
</template>


<script>
export default {
  name: 'DevicesTab',
  props: {
    tab: {
      type: Object,
      required: true,
    }
  },
}
</script>
