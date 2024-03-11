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
};

export let states = reactive({
  isDrawerToggled: false,
  isLoginDialogToggled: false,
  isLoggedIn: true,
  isLoading: false,
});

export const CurrentDestination = ref(Destinations.HOME);
export const CurrentTabName = ref(Destinations.HOME.tabName);

const handleHomeTraversal = () => {
  CurrentDestination.value = Destinations.HOME;
  CurrentTabName.value = CurrentDestination.value.tabName;
};

const handleDashboardTraversal = () => {
  CurrentDestination.value = Destinations.DASHBOARD;
  CurrentTabName.value = CurrentDestination.value.tabName;
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
  CurrentDestination.value = Destinations.ABOUT;
  CurrentTabName.value = CurrentDestination.value.tabName;
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
  CurrentTabName.value = newValue;
};

export const handleSideBarButtonClick = (button) => {
  console.log(button.title + ' button clicked');

  if (handleTraverse.hasOwnProperty(button.destination.tabName)) {
    console.log('traversing to ', button.destination.tabName);
    handleTraverse[button.destination.tabName]();

    CurrentDestination.value = button.destination;
  } else {
    console.log(button.destination, ' destination not found');
  }
};