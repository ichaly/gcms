import {createApp} from 'vue'
import router from "../../router"
import store from '../../store'
import App from './App.vue'
import ElementPlus from 'element-plus'

const app = createApp(App)
app.use(router)
app.use(store)
app.use(ElementPlus, {zIndex: 3000, size: 'normal'})
app.mount('#app')