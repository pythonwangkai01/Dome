import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

// createApp(App).use(store).use(router).mount('#app')
const app = createApp(App)
app.use(store).use(router).use(ElementPlus)
app.mount('#app')