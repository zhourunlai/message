import 'babel-polyfill';

import Vue from 'vue';
import App from './App';
import store from './store';
import router from './router';

Vue.config.devtools = true;

new Vue({
    el: 'body',
    components: { App },
    store: store,
    router: router
});
