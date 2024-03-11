// Plugins
import { registerPlugins } from '@/plugins'

// Components
import App from './App.vue'

import AppBar from './components/general/AppBar.vue'
import SideBar from './components/general/SideBar.vue'
import SideBarButton from './components/general/SideBarButton.vue'
import Feed from './components/general/Feed.vue'
import Post from './components/general/Post.vue'

import LoginDialog from './components/dialogs/LoginDialog.vue'

import Tabs from './components/tabs/Tabs.vue'
import HomeTab from './components/tabs/HomeTab.vue'
import AboutTab from './components/tabs/AboutTab.vue'

import DashboardTabs from './components/tabs/DashboardTabs.vue'
import DashboardTabButton from './components/general/DashboardTabButton.vue'

import DevicesTab from './components/tabs/DevicesTab.vue'
import PluginsTab from './components/tabs/PluginsTab.vue'
import UsersTab from './components/tabs/UsersTab.vue'

// Composables
import { createApp } from 'vue'

const app = createApp(App)
app.component('app-bar', AppBar)
app.component('side-bar', SideBar)
app.component('side-bar-button', SideBarButton)

app.component('feed', Feed)
app.component('post', Post)
app.component('login-dialog', LoginDialog)
app.component('tabs', Tabs)
app.component('home-tab', HomeTab)
app.component('about-tab', AboutTab)

app.component('dashboard-tabs', DashboardTabs)
app.component('dashboard-tab-button', DashboardTabButton)

app.component('devices-tab', DevicesTab)
app.component('plugins-tab', PluginsTab)
app.component('users-tab', UsersTab)

registerPlugins(app)

app.mount('#app')
