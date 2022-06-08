import { createRouter, createWebHashHistory } from 'vue-router'


const routes = [
  //首页
  {
    path: '/',
    name: 'Index',
    component: ()=>import('../views/Index.vue')
  },
  //登录页
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue')
  },
  //注册页
  {
    path: '/register',
    name: 'Register',
    component: ()=>import('../views/Register.vue')
  },
  //布局页
  {
    path: '/layout', 
    name:'Layout',
    component:()=>import('../views/Layout.vue')
  },

]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

//导入nprogress 
import Nprogress from 'nprogress'
//导入nprogress样式
import 'nprogress/nprogress.css'

//定义路由导航前置守卫
router.beforeEach((to,from,next)=>{
  Nprogress.start();
  next()
})
//定义路由导航后置守卫
router.afterEach((to,from)=>{
  Nprogress.done();

})

export default router
