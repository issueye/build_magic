import { createApp } from 'vue'
import './assets/css/tailwind.css'

import App from './App.vue'
import router from './router'
import { install } from './install'

import './assets/css/style.css'
import './assets/css/vars.scss'

const app = createApp(App)
app.use(router)
install(app)
app.mount('#app')
