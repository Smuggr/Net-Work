import { reactive, ref } from "vue";

export const Destinations = {
  HOME: {
    title: 'Home',
    icon: 'home-icon',
    tabName: 'home',
  },
  DASHBOARD: {
    title: 'Dashboard',
    icon: 'dashboard-icon',
    tabName: 'dashboard',
  },
  MY_ACCOUNT: {
    title: 'My Account',
    icon: 'account-icon',
    tabName: 'account',
  },
  LOG_OUT: {
    title: 'Log Out',
    icon: 'logout-icon',
    tabName: 'log-out',
  },
  LOG_IN: {
    title: 'Log In',
    icon: 'login-icon',
    tabName: 'log-in',
  },
  SETTINGS: {
    title: 'Settings',
    icon: 'settings-icon',
    tabName: 'settings',
  },
  ABOUT: {
    title: 'About',
    icon: 'about-icon',
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
}

export const DashboardTabs = {
  DEVICES: Destinations.DEVICES,
  USERS: Destinations.USERS,
  PLUGINS: Destinations.PLUGINS,
}

export let states = reactive({
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

const handleMyAccountTraversal = () => {

};

const handleLogOutTraversal = () => {

};

const handleLogInTraversal = () => {
  states.isLoginDialogToggled = true;
  states.isDrawerToggled = false;
  console.log(states);
};

const handleSettingsTraversal = () => {

};

const handleAboutTraversal = () => {
  CurrentTab.value = Tabs.ABOUT;
};

const handleTraverse = {
  [Destinations.HOME.tabName]: handleHomeTraversal,
  [Destinations.DASHBOARD.tabName]: handleDashboardTraversal,
  [Destinations.MY_ACCOUNT.tabName]: handleMyAccountTraversal,
  [Destinations.LOG_OUT.tabName]: handleLogOutTraversal,
  [Destinations.LOG_IN.tabName]: handleLogInTraversal,
  [Destinations.SETTINGS.tabName]: handleSettingsTraversal,
  [Destinations.ABOUT.tabName]: handleAboutTraversal,
};

export const handleTabChange = (newValue) => {
  console.log(newValue, ' tab changed');

  if (newValue == null || !Object.values(Tabs).includes(newValue)) {
    return;
  }

  console.log(newValue, ' good')
  CurrentTab.value = newValue;
};

export const handleDashboardTabChange = (newValue) => {
  console.log(newValue, ' tab changed');
    if (newValue == null || !Object.values(DashboardTabs).includes(newValue)) {
    return;
  }

  console.log(newValue, ' good')
  CurrentDashboardTab.value = newValue;
};

export const handleSideBarButtonClick = (button) => {
  console.log(button.title + ' button clicked');

  if (handleTraverse.hasOwnProperty(button.destination.tabName)) {
    console.log('traversing to ', button.destination.tabName);
    handleTraverse[button.destination.tabName]();

    console.log(button.destination);
    CurrentDestination.value = button.destination;
  } else {
    console.log(button.destination, ' destination not found');
  }
};