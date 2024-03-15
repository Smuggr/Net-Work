<template>
  <v-window-item :value="value">
    <v-card-title class="centered">
      <v-row align="center" justify="center" class="icon-container">
        <v-avatar color="surface-variant" size="180">
          <img src="@/assets/image.png" alt="Image" elevation="10" />
        </v-avatar>
      </v-row>
    </v-card-title>
    <v-card-text class="centered">
      <div class="username">{{ appStore.$state.currentUser.username }}</div>
      <div class="login">{{ appStore.$state.currentUser.login }}</div>
    </v-card-text>
</v-window-item>
</template>

<script>
import { useAppStore } from '@/stores/app';

export default {
  name: 'MyProfileTab',
  props: {
    value: {
      type: Object,
      required: true
    },
  },
  setup() {
    const appStore = useAppStore();

    return {
      appStore,
    };
  }
}
</script>

<style scoped>
.centered {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.icon-container {
  margin-top: 32px;
  margin-bottom: 4px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.username {
  font-size: 24px;
  font-weight: bold;
}

.login {
  font-size: 18px;
  font-style: italic;
}
</style>
