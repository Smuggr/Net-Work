<template>
  <v-dialog v-model="isVisible" max-width="500">
    <template v-slot:default="{ isActive }">
      <v-card title="Dialog">
        <v-card-text>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn
            text="Close Dialog"
            @click="isActive.value = false"
          ></v-btn>
        </v-card-actions>
      </v-card>
    </template>
  </v-dialog>
</template>

<script>
export default {
  name: 'LoginDialog',
  props: {
    value: {
      type: Boolean,
      default: false,
    },
  },
  computed: {
    isVisible: {
      get() {
        return this.value;
      },
      set(val) {
        this.$emit('update:value', val);
      },
    },
  },
  methods: {
    closeDialog() {
      this.$emit('update:value', false);
    },
  },
};
</script>
