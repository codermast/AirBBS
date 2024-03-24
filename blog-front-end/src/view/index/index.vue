<!-- HTMl -->
<template>
    <div class="app">

        <h2>你好，欢迎来到我的个人博客！</h2>
        <h2>下面是博客模块</h2>

        <table>
            <thead>
                <tr>
                    <th>用户模块</th>
                    <th>文章模块</th>
                    <th>系统模块</th>
                    <th></th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                <tr>1</tr>
                <tr>2</tr>
                <tr>3</tr>
                <tr>4</tr>
                <tr>5</tr>
            </tbody>
        </table>

    </div>

</template>

<!-- JS 或者 TS -->
<script lang="ts">

import User from '../../components/User.vue'

export default {
    name: 'App', // 组件名称
    components: { User } // 注册组件
}

</script>

<!-- CSS 样式 -->
<style>
.app {
    background-color: #ddd;
    box-shadow: 0 0 10px;
    border-radius: 10px;
    padding: 20px;
    text-align: center;
}

* {
    text-align: center;
}

td {
    width: 20vw;
}
</style>