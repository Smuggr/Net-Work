<template>
  <v-app>
    <app-bar @toggle-drawer="toggleDrawer" title="Smuggr Network" />
    <side-bar v-model="drawer">
      <template v-slot:primary>
        <side-bar-button title="Home" icon="mdi-home" />
        <side-bar-button title="Dashboard" icon="mdi-view-dashboard" />
      </template>

      <template v-slot:secondary>
        <side-bar-button size="large" title="My Account" icon="mdi-account" />
        <side-bar-button size="large" title="Log Out" icon="mdi-logout" />
        <side-bar-button size="large" title="About" icon="mdi-information" />
      </template>
    </side-bar>

    <v-main>
      <v-container fluid>
        <v-layout justify-center align-center>
          <v-flex>
            <v-card>
              <v-card-text>Main Content</v-card-text>
            </v-card>
          </v-flex>
        </v-layout>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      drawer: false,
    };
  },
  methods: {
    toggleDrawer() {
      this.drawer = !this.drawer;
    },
  },
};
</script>
