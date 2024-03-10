import { reactive } from "vue";

export const Destinations = {
  HOME: 'home',
  DASHBOARD: 'dashboard',
  MY_ACCOUNT: 'my-account',
  LOG_OUT: 'log-out',
  LOG_IN: 'log-in',
  SETTINGS: 'settings',
  ABOUT: 'about',
};

export let states = reactive({
  isDrawerToggled: false,
  isLoginDialogToggled: true,
  isLoggedIn: false,
  isLoading: true,
});

export let currentDestination = Destinations.HOME;

const handleHomeTraversal = () => {

};

const handleDashboardTraversal = () => {
  // Code for handling dashboard traversal
};

const handleMyAccountTraversal = () => {
  // Code for handling my account traversal
};

const handleLogOutTraversal = () => {
  // Code for handling log out traversal
};

const handleLogInTraversal = () => {
  states.isLoginDialogToggled = true;
  states.isDrawerToggled = false;
  console.log(states);
};

const handleSettingsTraversal = () => {
  // Code for handling settings traversal
};

const handleAboutTraversal = () => {
  // Code for handling about traversal
};

const handleTraverse = {
  [Destinations.HOME]: handleHomeTraversal,
  [Destinations.DASHBOARD]: handleDashboardTraversal,
  [Destinations.MY_ACCOUNT]: handleMyAccountTraversal,
  [Destinations.LOG_OUT]: handleLogOutTraversal,
  [Destinations.LOG_IN]: handleLogInTraversal,
  [Destinations.SETTINGS]: handleSettingsTraversal,
  [Destinations.ABOUT]: handleAboutTraversal,
};

export const handleSideBarButtonClick = (button) => {
  console.log(button.title + ' button clicked');

  if (handleTraverse.hasOwnProperty(button.destination)) {
    console.log('traversing to ', button.destination);
    handleTraverse[button.destination]();
  } else {
    console.log(button.destination, ' destination not found');
  }
};