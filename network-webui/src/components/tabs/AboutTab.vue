<template>
  <v-window-item :value="value">
    <br/>
    <div class="text-center">
      <h1>Source Code</h1>
    </div>

    <v-card
      append-icon="mdi-open-in-new"
      class="mx-auto my-8"
      href="https://github.com/Smuggr/Network"
      max-width="420"
      prepend-icon="mdi-github"
      rel="noopener"
      target="_blank"
      title="Smuggr Network on GitHub"
    />
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
