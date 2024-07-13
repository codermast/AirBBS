import { createApp } from 'vue'

import "./main.css"
import App from './App.vue'

// 导入自定义的 Pinia 实例
import pinia from '@/stores/pinia'; // 导入定义的 Pinia 实例

// 注册路由
import router from './router'

let app = createApp(App);

// 使用 Pinia
app.use(pinia)
app.use(router)
app.mount('#app')
