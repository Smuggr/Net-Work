export const Destinations = {
  HOME: 'home',
  DASHBOARD: 'dashboard',
  MY_ACCOUNT: 'my-account',
  LOG_OUT: 'log-out',
  LOG_IN: 'log-in',
  SETTINGS: 'settings',
  ABOUT: 'about',
};

export const toggleDrawer = (drawer) => {
  return !drawer;
};

export const handleSideBarButtonClick = (buttonName) => {
  console.log(buttonName + ' button clicked');
};