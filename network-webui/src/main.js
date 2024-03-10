// Plugins
import { registerPlugins } from '@/plugins'

// Components
import App from './App.vue'
import AppBar from './components/AppBar.vue'
import SideBar from './components/SideBar.vue'
import SideBarButton from './components/SideBarButton.vue'

// Composables
import { createApp } from 'vue'

const app = createApp(App)
app.component('app-bar', AppBar)
app.component('side-bar', SideBar)
app.component('side-bar-button', SideBarButton)

registerPlugins(app)

app.mount('#app')
