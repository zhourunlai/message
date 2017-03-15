<script>
import { actions } from '../store';

export default {
    vuex: {
        actions: actions,
    },
};
</script>

<template>
<div class="signup">
    <input type="text" />
    <input type="password" />
    <button>sign up</button>
</div>
</template>

<style lang="less" scoped>
.signup {
    height: 160px;
    border-top: solid 1px #ddd;
}
</style>
