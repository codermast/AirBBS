<script setup lang="ts">
import User from "@/pages/User.vue";
import { ref } from "vue";
import { Glasses, GlassesOutline, Key, LockClosedOutline, Mail } from '@vicons/ionicons5'
import ArrowRight from "@/icons/ArrowRight.vue";
import { registerUser } from "@/api/user";
import type { UserInfo } from "@/models/user"

let user = ref<UserInfo>({
  id: "",
  username: "",
  password: "",
  mail: "",
  tel: "",
  nickname: "",
  admin: ""
});

function register() {
  console.log("register",user.value)
  registerUser(user.value)
}
</script>

<template>
  <n-card
      class="register-card"
      :segmented="{
      content: 'soft',
      footer: 'soft'
    }"
  >
    <template #header>
      <div class="register-card-title">
        <n-icon size="20" :component="User"></n-icon>
        <div class="user-register-header-title">用户注册</div>
      </div>


    </template>


    <n-form
        label-placement="left"
        class="register-form"
    >
      <n-grid :cols="1">
        <n-gi :span="1">
          <n-form-item>
            <n-input v-model:value="user.username" placeholder="请输入您的用户名">
              <template #prefix>
                <n-icon :component="User"/>
              </template>
            </n-input>
          </n-form-item>
        </n-gi>

        <n-gi :span="1">
          <n-form-item>
            <n-input
                type="password"
                v-model:value="user.password"
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
          <n-form-item>
            <n-input
                type="password"
                v-model:value="user.password"
                placeholder="确认密码"
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
          <n-form-item>
            <n-input-group>
              <n-input
                  v-model:value="user.mail"
                  placeholder="邮箱"
              >
                <template #prefix>
                  <n-icon :component="Mail"></n-icon>
                </template>
              </n-input>

              <n-button>发送验证码</n-button>
            </n-input-group>
          </n-form-item>
        </n-gi>

        <n-gi :span="1">
          <n-form-item>
            <n-input
                show-count

                :maxlength="6"
                placeholder="请输入6位验证码"

            >
              <template #prefix>
                <n-icon :component="Key"></n-icon>
              </template>
            </n-input>
          </n-form-item>
        </n-gi>

        <n-gi :span="1">
          <n-grid cols="3">
            <n-gi :span="1">
              <n-button
                  type="success"
                  icon-placement="right"
                  @click="register"
              >
                <template #icon>
                  <n-icon :component="ArrowRight"></n-icon>
                </template>
                注册
              </n-button>
            </n-gi>

            <n-gi :span="1">

            </n-gi>

            <n-gi :span="1">
              <div class="register-tip-info">
                已有账号? 去登录
              </div>
            </n-gi>
          </n-grid>
        </n-gi>

      </n-grid>
    </n-form>

  </n-card>
</template>

<style scoped>
.register-card {
  margin: auto;
  width: 30vw;
  border-radius: 5px;
}

.register-card-title {
  color: #646b6f;
  display: flex;
  align-items: center;
}

.register-form {
  display: flex;
  align-items: center; /* 垂直居中 */
}

.register-tip-info {
  float: right;
}

.user-register-header-title {
  margin-left: 10px;
}
</style>