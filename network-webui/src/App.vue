<template>
  <v-app>
    <app-bar :loading="isLoading" @toggle-drawer="toggleDrawer" title="Smuggr Network" />
    
    <side-bar v-model="drawer">
      <template v-slot:primary>
        <side-bar-button title="Home" icon="mdi-home" />

        <template v-if="isLoggedIn">
          <side-bar-button title="Dashboard" icon="mdi-view-dashboard" />
        </template>
      </template>


      <template v-slot:secondary>
        <template v-if="isLoggedIn">
          <side-bar-button size="large" title="My Account" icon="mdi-account" />
          <side-bar-button size="large" title="Log Out" icon="mdi-logout" />
        </template>
        <template v-else>
          <side-bar-button size="large" title="Log In" icon="mdi-login" />
        </template>

        <side-bar-button size="large" title="Settings" icon="mdi-cog" />
        <side-bar-button size="large" title="About" icon="mdi-information" />
      </template>
    </side-bar>

    <v-main>
      <feed :posts="postsData">
        <post />
      </feed>
    </v-main>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      drawer: true,
      isLoggedIn: true,
      isLoading: true,
    };
  },
  methods: {
    toggleDrawer() {
      this.drawer = !this.drawer;
    },
  },
};
</script>
