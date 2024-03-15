<template>
  <v-window-item :value="tab">

  </v-window-item>
</template>


<script>
export default {
  name: 'UsersTab',
  props: {
    tab: {
      type: Object,
      required: true,
    }
  },
}
</script>
