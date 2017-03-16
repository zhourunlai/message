import Vue from 'vue';
import VueRouter from 'vue-router';

import Signup from 'components/signup';
import Main from 'components/main';

Vue.use(VueRouter);

const routes = [
    {
        path: '/signup',
        name: 'signup',
        component: Signup,
    },
    {
        path: '/main',
        name: 'main',
        component: Main,
    },
];

const router = new VueRouter({
    // mode: 'history',
    routers: routes
});

export default router;
