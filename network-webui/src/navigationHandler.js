import { reactive, ref } from "vue";

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

export const states = reactive({
  isDrawerToggled: false,
  isLoginDialogToggled: false,
  isLoggedIn: true,
  isLoading: false,
});

export const CurrentDestination = ref(Destinations.HOME);
export const CurrentTab = ref(Tabs.HOME);
export const CurrentDashboardTab = ref(DashboardTabs.DEVICES);

const handleHomeTraversal = () => {
  CurrentTab.value = Tabs.HOME;
};

const handleDashboardTraversal = () => {
  CurrentTab.value = Tabs.DASHBOARD;
};

const handleLogInTraversal = () => {
  states.isLoginDialogToggled = true;
  states.isDrawerToggled = false;
};

const handleAboutTraversal = () => {
  CurrentTab.value = Tabs.ABOUT;
};

const handleTraverse = {
  [Destinations.HOME.tabName]: handleHomeTraversal,
  [Destinations.DASHBOARD.tabName]: handleDashboardTraversal,
  [Destinations.LOG_IN.tabName]: handleLogInTraversal,
  [Destinations.ABOUT.tabName]: handleAboutTraversal,
};

export const handleTabChange = (newValue) => {
  if (newValue == null || !Object.values(Tabs).includes(newValue)) {
    return;
  }
  CurrentTab.value = newValue;
};

export const handleDashboardTabChange = (newValue) => {
  if (newValue == null || !Object.values(DashboardTabs).includes(newValue)) {
    return;
  }
  CurrentDashboardTab.value = newValue;
};

export const handleSideBarButtonClick = (button) => {
  if (handleTraverse.hasOwnProperty(button.destination.tabName)) {
    handleTraverse[button.destination.tabName]();
    CurrentDestination.value = button.destination;
  }
};
