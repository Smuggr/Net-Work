<template>
  <v-app>
    <app-bar :isLoading="appStore.$state.isLoading" @toggle-drawer="appStore.setIsDrawerToggled(!appStore.$state.isDrawerToggled)" title="Smuggr Network" />

    <login-dialog v-model="appStore.$state.isLoginDialogToggled" />

    <side-bar v-model="appStore.$state.isDrawerToggled">
      <template v-slot:primary>
        <side-bar-button :destination="Destinations.HOME" @button-click="handleSideBarButtonClick" />

        <template v-if="appStore.$state.isLoggedIn">
          <side-bar-button :destination="Destinations.DASHBOARD" @button-click="handleSideBarButtonClick" />
        </template>
      </template>

      <template v-slot:secondary>
        <template v-if="appStore.$state.isLoggedIn">
          <side-bar-button :destination="Destinations.MY_PROFILE" size="large" @button-click="handleSideBarButtonClick"/>
          <side-bar-button :destination="Destinations.LOG_OUT" size="large" @button-click="handleSideBarButtonClick"/>
        </template>
        <template v-else>
          <side-bar-button :destination="Destinations.LOG_IN" size="large" @button-click="handleSideBarButtonClick" />
        </template>

        <side-bar-button :destination="Destinations.SETTINGS" size="large" @button-click="handleSideBarButtonClick" />
        <side-bar-button :destination="Destinations.ABOUT" size="large" @button-click="handleSideBarButtonClick" />
      </template>
    </side-bar>
    
    <v-main>
      <tabs :value="appStore.$state.currentTab" @update:value="handleTabChange">
        <template v-slot:content>
          <dashboard-tabs :value="Tabs.DASHBOARD" :childValue="appStore.$state.currentDashboardTab" @update:childValue="handleDashboardTabChange">
            <template v-slot:buttons>
              <dashboard-tab-button :tab="DashboardTabs.DEVICES" />
              <dashboard-tab-button :tab="DashboardTabs.PLUGINS" />
              <dashboard-tab-button :tab="DashboardTabs.USERS" />
            </template>

            <template v-slot:content>
              <devices-tab :tab="DashboardTabs.DEVICES" />
              <plugins-tab :tab="DashboardTabs.PLUGINS" />
              <users-tab :tab="DashboardTabs.USERS" />
            </template>
          </dashboard-tabs>

          <home-tab :value="Tabs.HOME"/>
          <about-tab :value="Tabs.ABOUT"/>
          <my-profile-tab :value="Tabs.MY_PROFILE" />
          <settings-tab :value="Tabs.SETTINGS" />
        </template>
      </tabs>
    </v-main>
  </v-app>
</template>

<script>
import {
  handleSideBarButtonClick,
  handleDashboardTabChange,
  handleTabChange,

  Destinations,
  Tabs,
  DashboardTabs,
} from './navigationHandler';

import { useAppStore } from './stores/app';

export default {
  methods: {
    handleSideBarButtonClick,
    handleDashboardTabChange,
    handleTabChange,
  },
  setup() {
    const appStore = useAppStore();

    return {
      Destinations,
      Tabs,
      DashboardTabs,
      appStore,
    };
  },
};
</script>
