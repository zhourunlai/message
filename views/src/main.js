import 'babel-polyfill';

import Vue from 'vue';
import App from './App';
import store from './store';
import router from './router';
// import resource from './resource';

Vue.config.devtools = true;

new Vue({
    el: 'body',
    components: { App },
    store: store,
    router: router,
    //resource: resource,
});

router.push('/signup');
