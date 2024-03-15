<template>
  <v-window-item :value="value">
  
  </v-window-item>
</template>

<script>
import { ref } from 'vue';

export default {
  name: 'SettingsTab',
  props: {
    value: {
      type: Object,
      required: true
    }
  }
}
</script>
