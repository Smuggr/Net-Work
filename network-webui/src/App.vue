<template>
  <v-app>
    <app-bar v-model:isLoading="states.isLoading" @toggle-drawer="states.isDrawerToggled = !states.isDrawerToggled;" title="Smuggr Network" />

    <login-dialog v-model="states.isLoginDialogToggled" />

    <side-bar v-model="states.isDrawerToggled">
      <template v-slot:primary>
        <side-bar-button :destination="Destinations.HOME" title="Home" icon="mdi-home" @button-click="handleSideBarButtonClick" />

        <template v-if="states.isLoggedIn">
          <side-bar-button :destination="Destinations.DASHBOARD" title="Dashboard" icon="mdi-view-dashboard" @button-click="handleSideBarButtonClick" />
        </template>
      </template>


      <template v-slot:secondary>
        <template v-if="states.isLoggedIn">
          <side-bar-button :destination="Destinations.MY_ACCOUNT" size="large" title="My Account" icon="mdi-account" @button-click="handleSideBarButtonClick"/>
          <side-bar-button :destination="Destinations.LOG_OUT" size="large" title="Log Out" icon="mdi-logout" @button-click="handleSideBarButtonClick"/>
        </template>
        <template v-else>
          <side-bar-button :destination="Destinations.LOG_IN" size="large" title="Log In" icon="mdi-login" @button-click="handleSideBarButtonClick" />
        </template>

        <side-bar-button :destination="Destinations.SETTINGS" size="large" title="Settings" icon="mdi-cog" @button-click="handleSideBarButtonClick" />
        <side-bar-button :destination="Destinations.ABOUT" size="large" title="About" icon="mdi-information" @button-click="handleSideBarButtonClick" />
      </template>
    </side-bar>
    
    <v-main>
      <tabs :value="CurrentTab" @update:value="handleTabChange">
        <template v-slot:content>
          <home-tab :value="Tabs.HOME"/>
          <dashboard-tab :value="Tabs.DASHBOARD"/>
          <about-tab :value="Tabs.ABOUT"/>
        </template>
      </tabs>
    </v-main>
  </v-app>
</template>

<script>
import { Destinations, Tabs, CurrentDestination, CurrentTab, handleSideBarButtonClick, handleTabChange, states } from './navigationHandler';

export default {
  methods: {
    handleSideBarButtonClick,
    handleTabChange,
  },
  setup() {
    return {
      Destinations,
      CurrentDestination,
      Tabs,
      CurrentTab,
      states,
    };
  },
};
</script>
