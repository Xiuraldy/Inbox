import './assets/main.css'

import { createApp } from 'vue'

import App from './App.vue'
import router from './router'

const app = createApp(App) // Crea la aplicación Vue usando App.vue

app.use(router) // Para navegación entre páginas

app.mount('#app') // Monta la aplicación en el elemento con ID #app en index.html
