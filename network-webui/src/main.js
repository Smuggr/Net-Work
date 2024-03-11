// Plugins
import { registerPlugins } from '@/plugins'

// Components
import App from './App.vue'
import AppBar from './components/AppBar.vue'
import SideBar from './components/SideBar.vue'
import SideBarButton from './components/SideBarButton.vue'
import Feed from './components/Feed.vue'
import Post from './components/Post.vue'
import LoginDialog from './components/LoginDialog.vue'
import Tabs from './components/Tabs.vue'

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

registerPlugins(app)

app.mount('#app')
