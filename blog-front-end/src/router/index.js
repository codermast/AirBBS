// router/index.ts or router/index.js
import { createRouter, createWebHistory } from 'vue-router';

import { User } from './src/components/User.vue'
const routes = [
    {
        path: '/',
        name: 'Home',
        component: User,
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
