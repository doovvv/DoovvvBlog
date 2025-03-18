import { createRouter, createWebHistory } from 'vue-router'
import Login from '../components/Login.vue'
import Admin from '../components/Admin.vue'
import index from '../components/index.vue'
import addart from '../components/article/addart.vue'
import artlist from '../components/article/artlist.vue'
import catelist from '../components/category/catelist.vue'
import userlist from '../components/user/userlist.vue'
const routes = [
    {
        path: '/login',
        name: 'Login',
        component: Login
    },
    {
        path: '/admin',
        name: 'Admin',
        component: Admin,
        children: [
            {
                path: '/admin/index',
                name: 'index',
                component: index
            },
            {
                path: '/admin/addart',
                name: 'addart',
                component: addart
            },
            {
                path: '/admin/artlist',
                name: 'artlist',
                component: artlist
            },
            {
                path: '/admin/catelist',
                name: 'catelist',
                component: catelist
            },
            {
                path: '/admin/userlist',
                name: 'userlist',
                component: userlist
            }

        ]
    }

]

const router = createRouter({
    history: createWebHistory(), // HTML5 模式
    routes
})
router.beforeEach((to, from, next) => {
    // 检查用户是否已登录
    const isLoggedIn = localStorage.getItem('token')
    if (to.path === '/login') {
        return next()
    }
    if (to.path === '/admin' && !isLoggedIn) {
        return next('/login')
    }
    else {
        return next()
    }
})
export default router
