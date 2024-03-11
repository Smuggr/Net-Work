import { reactive } from 'vue';
import { defineStore } from 'pinia';

export const useAppStore = defineStore('app', () => {
  const states = reactive({
    isDrawerToggled: false,
    isLoginDialogToggled: true,
    isLoggedIn: false,
    isLoading: false,
  });

  return {
    states
  }
})