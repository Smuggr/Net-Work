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
import DashboardTab from './components/tabs/DashboardTab.vue'
import AboutTab from './components/tabs/AboutTab.vue'


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
app.component('dashboard-tab', DashboardTab)
app.component('about-tab', AboutTab)

registerPlugins(app)

app.mount('#app')
