import { reactive, ref } from 'vue';
import { useAppStore } from './stores/app';
import { authenticateUser } from './apiHandler';

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
  MY_ACCOUNT: {
    title: 'My Account',
    icon: 'mdi-account',
    tabName: 'account',
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
  ABOUT: Destinations.ABOUT,
};

export const DashboardTabs = {
  DEVICES: Destinations.DEVICES,
  USERS: Destinations.USERS,
  PLUGINS: Destinations.PLUGINS,
};

const handleHomeTraversal = () => {
  const store = useAppStore();
  store.setCurrentTab(Tabs.HOME);

  console.log(store);
  authenticateUser('administrator', 'Password123$');
};

const handleDashboardTraversal = () => {
  const store = useAppStore();
  store.setCurrentTab(Tabs.DASHBOARD);
};

const handleLogInTraversal = () => {
  const store = useAppStore();
  console.log('balls', typeof (store.setIsLoginDialogToggled));

  store.setIsLoginDialogToggled(true);
  store.setIsDrawerToggled(false);

  console.log(store.$state.isDrawerToggled);
};

const handleAboutTraversal = () => {
  const store = useAppStore();
  store.setCurrentTab(Tabs.ABOUT);
};

const handleTraverse = {
  [Destinations.HOME.tabName]: handleHomeTraversal,
  [Destinations.DASHBOARD.tabName]: handleDashboardTraversal,
  [Destinations.LOG_IN.tabName]: handleLogInTraversal,
  [Destinations.ABOUT.tabName]: handleAboutTraversal,
};

export const handleTabChange = (newValue) => {
  const store = useAppStore();

  if (newValue == null || !Object.values(Tabs).includes(newValue)) {
    return;
  }
  
  store.setCurrentTab(newValue);
};

export const handleDashboardTabChange = (newValue) => {
  const store = useAppStore();

  if (newValue == null || !Object.values(DashboardTabs).includes(newValue)) {
    return;
  }

  store.setCurrentDashboardTab(newValue);
};

export const handleSideBarButtonClick = (button) => {
  if (handleTraverse.hasOwnProperty(button.destination.tabName)) {
    handleTraverse[button.destination.tabName]();
  }
};
