<template>
    <div class="user">
        <table>
            <thead>
                <tr>
                    <th>用户ID</th>
                    <th>用户名</th>
                    <th>用户密码</th>
                    <th>用户邮箱</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="user in users">
                    <th>{{ user.uid }}</th>
                    <th>{{ user.username }}</th>
                    <th>{{ user.password }}</th>
                    <th>{{ user.email }}</th>
                </tr>
            </tbody>
        </table>
        <button @click="getUsers">获取用户列表信息</button>
    </div>
</template>
<script lang="ts">
import axios from 'axios'
export default {
    name: "User",
    data() {
        return {
            users: null,
        }
    },
    methods: {
        getUsers() {
    // api 请求地址
            var api = "http://localhost:8080/admin/user/list";
            // 发起 get 请求
            axios.get(api)
            // 响应值为 response
                .then((response) => {
            // 判断响应值
                    if (response.status == 200 && response.data.code == 1) {
                        this.users = response.data.data;
                    } else {
                        alert("用户列表获取失败！")
                    }
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
    }
}
</script>
<style>
table {
    border-collapse: collapse;
    width: 100%;
}

th,
td {
    border: 1px solid black;
    padding: 8px;
    text-align: left;
}

th {
    background-color: #f2f2f2;
}
</style>