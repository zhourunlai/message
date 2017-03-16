<script>
import { actions } from '../store';

export default {
    vuex: {
        actions: actions,
        data:{
            // this.$http.get('/v1/user/signup', {
            //     params: {
            //         username: this.username,
            //         password: this.password,
            //
            //     }
            // }).then((response) => {
            //     transition.next({
            //         data: response.json().data,
            //     })
            // }, (response) => {
            //     // error
            // })
        }
    },
};
</script>

<template>
<div class="signup">
    <input type="text" name="username" />
    <input type="password" name="password" />
    <button>sign up</button>
</div>
</template>

<style lang="less" scoped>
.signup {
    height: 160px;
    border-top: solid 1px #ddd;
}
</style>
