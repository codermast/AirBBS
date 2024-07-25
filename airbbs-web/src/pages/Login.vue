<script setup lang="ts">
import User from "@/icons/User.vue";
import { Glasses, GlassesOutline, LockClosedOutline, LogoGithub, LogoWechat } from '@vicons/ionicons5'
import LogoQQ from '@/icons/QQ.vue'
import ArrowRight from "@/icons/ArrowRight.vue";
import { getUserById, loginUser } from "@/api/user";
import { ref } from 'vue'
import type { UserLoginInfo } from "@/models/user";
import { useRouter } from "vue-router";
import { useMessage,useLoadingBar } from "naive-ui"
import { HeaderAuthTokenName } from "@/config";
import { useUserStore } from "@/stores/userStore";

const userStore = useUserStore()
const router = useRouter();
const message = useMessage();
const loadingBar = useLoadingBar()

let user = ref<UserLoginInfo>({
  username: '',
  password: ''
})

let rules = ref({
  username: {
    required: true,
    message: '请输入用户名',
    trigger: ['input', 'blur']
  },
  password: {
    required: true,
    message: '请输入密码',
    trigger: ['input', 'blur']
  },
})

// 登录执行方法
async function login() {
  let response = await loginUser(user.value)
  console.log("loginResponse",response)
  if (response.status == 200) {
    // 登录成功，更新登录信息
    userStore.JWTToken = response.data.data.token
    userStore.userLoginStatus = true

    let userInfoResponse = await getUserById(response.data.data.user_id)
    console.log("userInfoResponse",userInfoResponse)
    if (userInfoResponse.status == 200){
      userStore.userInfo = userInfoResponse.data.data
      console.log("userInfo:",userStore.userInfo)
    }

    message.success("登录成功！")

    // 路由跳转
    router.push({name: "Index"})

  }else{
    message.error(response.data.message)
  }
}

</script>

<template>
  <n-card
      class="login-card"
      :segmented="{
      content: 'soft',
      footer: 'soft'
    }"
  >
    <template #header>
      <div class="login-card-title">
        <n-icon size="20" :component="User"></n-icon>
        <div class="user-login-header-title">用户登录</div>
      </div>
    </template>

    <n-form
        :rules="rules"
        :model="user"
        label-placement="left"
        class="login-form"
    >
      <n-grid :cols="1">
        <n-gi :span="1">
          <n-form-item path="username">
            <n-input
                v-model:value="user.username"
                placeholder="邮箱/用户名/手机号">
              <template #prefix>
                <n-icon :component="User"/>
              </template>
            </n-input>
          </n-form-item>
        </n-gi>

        <n-gi :span="1">
          <n-form-item path="password">
            <n-input
                v-model:value="user.password"
                type="password"
                placeholder="密码"
                show-password-on="click"
            >
              <template #prefix>
                <n-icon :component="LockClosedOutline"/>
              </template>

              <template #password-visible-icon>
                <n-icon :size="16" :component="GlassesOutline"/>
              </template>
              <template #password-invisible-icon>
                <n-icon :size="16" :component="Glasses"/>
              </template>
            </n-input>
          </n-form-item>
        </n-gi>

        <n-gi :span="1">
          <n-grid cols="3">
            <n-gi :span="1">
              <n-button
                  @click="login"
                  type="success"
                  icon-placement="right"
              >
                <template #icon>
                  <n-icon :component="ArrowRight"></n-icon>
                </template>
                登录
              </n-button>
            </n-gi>

            <n-gi :span="1">

            </n-gi>

            <n-gi :span="1">
              <div class="login-tip-info">
                <router-link :to="{ name : 'ResetPassword' }" class="router-link">忘记密码</router-link>

                ? or

                <router-link :to="{ name : 'Register' }" class="router-link">注册</router-link>

              </div>
            </n-gi>
          </n-grid>
        </n-gi>

        <n-gi :span="1">
          <n-divider title-placement="center" class="divider-login">
            第三方登录
          </n-divider>
        </n-gi>

        <n-gi :span="1">
          <n-grid :cols="1" y-gap="10px">
            <n-gi :span="1">
              <n-button class="login-third-button">
                <template #icon>
                  <n-icon :component="LogoWechat">

                  </n-icon>
                </template>
                微信登录
              </n-button>
            </n-gi>

            <n-gi :span="1">
              <n-button class="login-third-button">
                <template #icon>
                  <n-icon :component="LogoQQ">

                  </n-icon>
                </template>
                QQ登录
              </n-button>
            </n-gi>

            <n-gi :span="1">
              <n-button class="login-third-button">
                <template #icon>
                  <n-icon :component="LogoGithub">

                  </n-icon>
                </template>

                Github登录
              </n-button>
            </n-gi>

          </n-grid>
        </n-gi>
      </n-grid>
    </n-form>

  </n-card>
</template>

<style scoped>
.login-card {
  margin: auto;
  width: 30vw;
  border-radius: 5px;
}

.login-card-title {
  color: #646b6f;
  display: flex;
  align-items: center;
}

.login-form {
  display: flex;
  align-items: center; /* 垂直居中 */
}

.divider-login {
  font-size: .85714286rem;
  color: #a5a5a5;
  font-weight: 400;
}

.login-third-button {
  width: 100%;
}

.login-tip-info {
  float: right;
}

.user-login-header-title {
  margin-left: 10px;
}
</style>