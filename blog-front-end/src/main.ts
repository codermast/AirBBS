// 引入 createApp 用于创建应用
import { createApp } from 'vue'
import Index from './view/index/index.vue';
import router from './router'; // 导入路由实例

const app = createApp(Index);
app.use(router); // 使用路由实例
app.mount('#app');
