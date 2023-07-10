import {createApp} from 'vue'
import router from "../router"
import store from '../store'
import App from './index.vue'
import ElementPlus from 'element-plus'
import "~/styles/index.scss"
const app = createApp(App)
app.use(router)
app.use(store)
app.use(ElementPlus, {zIndex: 3000, size: 'normal'})
app.mount('#app')