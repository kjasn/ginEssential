import Vue from 'vue';
import VueRouter from 'vue-router';
import HomeView from '../views/HomeView.vue';
// import RegisterView from '../views/register/RegisterView.vue';
// import LoginView from '../views/login/LoginView.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue'),
  },
  {
    path: '/register',
    name: 'register',
    // component: () => import(RegisterView),
    // component: RegisterView,
    // lazy load
    component: () => import('../views/register/RegisterView.vue'),
  },
  {
    path: '/login',
    name: 'login',
    // component: () => import(LoginView),
    // component: LoginView,
    component: () => import('../views/login/LoginView.vue'),
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
