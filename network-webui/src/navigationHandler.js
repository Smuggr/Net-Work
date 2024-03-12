import { useAppStore } from './stores/app';
import { useAuthStore } from './stores/auth';

import { getDevices } from './apiHandler';

export const Destinations = {
  HOME: {
    title: 'Home',
    icon: 'mdi-home',
    tabName: 'home',
  },
  DASHBOARD: {
    title: 'Dashboard',
    icon: 'mdi-view-dashboard',
    tabName: 'dashboard',
  },
  MY_PROFILE: {
    title: 'My Profile',
    icon: 'mdi-account',
    tabName: 'my-profile',
  },
  LOG_OUT: {
    title: 'Log Out',
    icon: 'mdi-logout',
    tabName: 'log-out',
  },
  LOG_IN: {
    title: 'Log In',
    icon: 'mdi-login',
    tabName: 'log-in',
  },
  SETTINGS: {
    title: 'Settings',
    icon: 'mdi-cog',
    tabName: 'settings',
  },
  ABOUT: {
    title: 'About',
    icon: 'mdi-information',
    tabName: 'about',
  },
  DEVICES: {
    title: 'Devices',
    icon: 'mdi-memory',
    tabName: 'devices',
  },
  USERS: {
    title: 'Users',
    icon: 'mdi-account-group',
    tabName: 'users',
  },
  PLUGINS: {
    title: 'Plugins',
    icon: 'mdi-puzzle',
    tabName: 'plugins',
  },
};

export const Tabs = {
  HOME: Destinations.HOME,
  DASHBOARD: Destinations.DASHBOARD,
  MY_PROFILE: Destinations.MY_PROFILE,
  SETTINGS: Destinations.SETTINGS,
  ABOUT: Destinations.ABOUT,
};

export const DashboardTabs = {
  DEVICES: Destinations.DEVICES,
  USERS: Destinations.USERS,
  PLUGINS: Destinations.PLUGINS,
};

const handleHomeTraversal = () => {
  const appStore = useAppStore();
  appStore.setCurrentTab(Tabs.HOME);
  getDevices();
};

const handleDashboardTraversal = () => {
  const appStore = useAppStore();
  appStore.setCurrentTab(Tabs.DASHBOARD);
};

const handleLogInTraversal = () => {
  const appStore = useAppStore();

  appStore.setIsLoginDialogToggled(true);
  appStore.setIsDrawerToggled(false);
};

const handleLogOutTraversal = () => {
  const appStore = useAppStore();
  const authStore = useAuthStore();

  appStore.setIsLoggedIn(false);
  appStore.setIsLoginDialogToggled(false);
  appStore.setIsDrawerToggled(false);
  appStore.setCurrentTab(Tabs.HOME);

  authStore.clearJWTToken();
};

const handleMyProfileTraversal = () => {
  const appStore = useAppStore();
  appStore.setCurrentTab(Tabs.MY_PROFILE);
}

const handleSettingsTraversal = () => {
  const appStore = useAppStore();
  appStore.setCurrentTab(Tabs.SETTINGS);
};

const handleAboutTraversal = () => {
  const appStore = useAppStore();
  appStore.setCurrentTab(Tabs.ABOUT);
};

const handleTraverse = {
  [Destinations.HOME.tabName]: handleHomeTraversal,
  [Destinations.DASHBOARD.tabName]: handleDashboardTraversal,
  [Destinations.LOG_IN.tabName]: handleLogInTraversal,
  [Destinations.LOG_OUT.tabName]: handleLogOutTraversal,
  [Destinations.MY_PROFILE.tabName]: handleMyProfileTraversal,
  [Destinations.SETTINGS.tabName]: handleSettingsTraversal,
  [Destinations.ABOUT.tabName]: handleAboutTraversal,
};

export const handleTabChange = (newValue) => {
  const appStore = useAppStore();

  if (newValue == null || !Object.values(Tabs).includes(newValue)) {
    return;
  }
  
  appStore.setCurrentTab(newValue);
};

export const handleDashboardTabChange = (newValue) => {
  const appStore = useAppStore();

  if (newValue == null || !Object.values(DashboardTabs).includes(newValue)) {
    return;
  }

  appStore.setCurrentDashboardTab(newValue);
};

export const handleSideBarButtonClick = (button) => {
  if (handleTraverse.hasOwnProperty(button.destination.tabName)) {
    handleTraverse[button.destination.tabName]();
  }
};
