import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const now = new Date();
const store = new Vuex.Store({
    state: {
        // 获取 username
        // url = window.location.href;
        // var username = url.split("user=");
        // username = username[1];

        // 当前用户
        user: {
            name: 'xiaorun',
            img: 'dist/images/1.jpg'
        },
        // 会话列表
        sessions: [
            {
                id: 1,
                user: {
                    name: 'test1',
                    img: 'dist/images/2.png'
                },
                messages: [
                    {
                        content: 'Hi, 我是test1',
                        date: now
                    }, {
                        content: '看什么看，赶快写 Golang',
                        date: now
                    }
                ]
            },
            {
                id: 2,
                user: {
                    name: 'test2',
                    img: 'dist/images/3.png'
                },
                messages: [
                    {
                        content: 'Hi, 我是test2',
                        date: now
                    },
                ]
            }

            // this.$http.get('/v1/user/xiaorun/contacts', {
            //     params: {
            //         name: this.username,
            //         img: '',
            //
            //     }
            // }).then((response) => {
            //     transition.next({
            //         data: response.json().data,
            //     })
            // }, (response) => {
            //     // error
            // })

        ],
        // 当前会话
        currentSessionId: 1,
        // 过滤只包含这个key的会话
        filterKey: ''
    },
    mutations: {
        INIT_DATA (state) {
            let data = localStorage.getItem('vue-chat-session');
            if (data) {
                state.sessions = JSON.parse(data);
            }
        },
        // 发送消息
        SEND_MESSAGE ({ sessions, currentSessionId }, content) {
            let session = sessions.find(item => item.id === currentSessionId);
            session.messages.push({
                content: content,
                date: new Date(),
                self: true
            });
        },
        // 选择会话
        SELECT_SESSION (state, id) {
            state.currentSessionId = id;
        } ,
        // 搜索用户
        SET_FILTER_KEY (state, value) {
            state.filterKey = value;
        }
    }
});

store.watch(
    (state) => state.sessions,
    (val) => {
        console.log('CHANGE: ', val);
        localStorage.setItem('vue-message-session', JSON.stringify(val));
    },
    {
        deep: true
    }
);

export default store;

export const actions = {
    initData: ({ dispatch }) => dispatch('INIT_DATA'),
    sendMessage: ({ dispatch }, content) => dispatch('SEND_MESSAGE', content),
    selectSession: ({ dispatch }, id) => dispatch('SELECT_SESSION', id),
    search: ({ dispatch }, value) => dispatch('SET_FILTER_KEY', value)
};
