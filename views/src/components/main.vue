<script>
import { actions } from '../store';

import Info from './info';
import List from './list';
import Text from './text';
import Chat from './chat';

export default {
    components: { Info, List, Text, Chat },
    vuex: {
        actions: actions
    },
    created () {
        this.initData();
    }
}
</script>

<template>
<div id="app">
    <div class="sidebar">
        <info></info>
        <list></list>
    </div>
    <div class="main">
        <chat></chat>
        <text></text>
    </div>
</div>
</template>

<style lang="less" scoped>
#app {
    margin: 20px auto;
    width: 800px;
    height: 600px;

    overflow: hidden;
    border-radius: 3px;

    .sidebar, .main {
        height: 100%;
    }
    .sidebar {
        float: left;
        width: 200px;
        color: #f4f4f4;
        background-color: #2e3238;
    }
    .main {
        position: relative;
        overflow: hidden;
        background-color: #eee;
    }
    .text {
        position: absolute;
        width: 100%;
        bottom: 0;
        left: 0;
    }
    .message {
        height: ~'calc(100% - 160px)';
    }
}
</style>
