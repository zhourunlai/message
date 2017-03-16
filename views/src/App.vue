<script>
import { actions } from './store';

export default {
    vuex: {
        actions: actions
    },
    created () {
        this.initData();
    }
}
</script>

<template>
<div id="xiaorun">
    <!-- <transition name="fade" mode="out-in" appear>
        <keep-alive> -->
          <router-view></router-view>
        <!-- </keep-alive>
      </transition> -->
</div>
</template>

<style lang="less" scoped>
</style>
