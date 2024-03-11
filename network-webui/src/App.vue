<template>
  <v-app>
    <app-bar v-model:isLoading="states.isLoading" @toggle-drawer="states.isDrawerToggled = !states.isDrawerToggled;" title="Smuggr Network" />

    <login-dialog v-model="states.isLoginDialogToggled" />

    <side-bar v-model="states.isDrawerToggled">
      <template v-slot:primary>
        <side-bar-button :destination="Destinations.HOME" @button-click="handleSideBarButtonClick" />

        <template v-if="states.isLoggedIn">
          <side-bar-button :destination="Destinations.DASHBOARD" @button-click="handleSideBarButtonClick" />
        </template>
      </template>


      <template v-slot:secondary>
        <template v-if="states.isLoggedIn">
          <side-bar-button :destination="Destinations.MY_ACCOUNT" size="large" @button-click="handleSideBarButtonClick"/>
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
      <tabs :value="CurrentTab" @update:value="handleTabChange">
        <template v-slot:content>


          <dashboard-tabs :value="Tabs.DASHBOARD" :childValue="CurrentDashboardTab" @update:childValue="handleDashboardTabChange">
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
        </template>
      </tabs>
    </v-main>
  </v-app>
</template>

<!-- <script setup>
import { useAppStore } from './stores/app';

const store = useAppStore();
</script> -->

<script>
import {
  Destinations,
  Tabs,
  DashboardTabs,
  CurrentTab,
  CurrentDashboardTab,
  handleSideBarButtonClick,
  handleDashboardTabChange,
  handleTabChange,
  states
} from './navigationHandler';

export default {
  methods: {
    handleSideBarButtonClick,
    handleDashboardTabChange,
    handleTabChange,
  },
  setup() {
    return {
      Destinations,
      Tabs,
      DashboardTabs,
      CurrentTab,
      CurrentDashboardTab,
      states,
    };
  },
};
</script>
