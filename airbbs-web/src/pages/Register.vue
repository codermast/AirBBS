<script setup lang="ts">
import User from "@/icons/User.vue";
import { ref } from "vue";
import { Glasses, GlassesOutline, Key, LockClosedOutline, Mail } from '@vicons/ionicons5'
import ArrowRight from "@/icons/ArrowRight.vue";
import { registerUser } from "@/api/user";
import type { FormInst, FormItemRule } from 'naive-ui'
import { useMessage } from "naive-ui"
import { sendRegisterAuthCode } from "@/api/auth";

const message = useMessage()
const formRef = ref<FormInst | null>(null)

let user = ref({
  id: '',
  username: '',
  password: '',
  mail: '',
  tel: '',
  nickname: '',
  code: '',
  confirmPassword : ""
});

// 邮箱是否输入正确
let isEmailValid = ref(false);

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
  confirmPassword: [
    {
      required: true,
      message: '请再次输入密码',
      trigger: ['input', 'blur']
    },
    {
      validator: validatePasswordStartWith,
      message: '两次密码输入不一致',
      trigger: ['input']
    },
    {
      validator: validatePasswordSame,
      message: '两次密码输入不一致',
      trigger: ['blur', 'password-input']
    }
  ],
  mail: [
    {
      required: true,
      message: '请输入邮箱地址',
      trigger: ['input',"blur"]
    },
    {
      validator: validateMailFormat,
      message: '邮箱格式错误',
      trigger: ["blur",'mail-input']
    },]
})


function validatePasswordStartWith(rule: FormItemRule, value: string): boolean {
  return (
      !!user.value.password &&
      user.value.password.startsWith(value) &&
      user.value.password.length >= value.length
  )
}

function validatePasswordSame(rule: FormItemRule, value: string): boolean {
  return value === user.value.password
}

function validateMailFormat(rule: FormItemRule, value: string): boolean {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  isEmailValid.value = emailRegex.test(user.value.mail);
  // console.log("验证结果", isEmailValid.value)
  return isEmailValid.value;
}


// 发送验证码
async function sendAuthCodeToMail(event: Event) {
  // 如果邮箱验证未通过，则阻止向后传递
  if (!isEmailValid.value) {
    // console.log("验证码发送被阻止！")
    event.stopImmediatePropagation()
    return
  }

  let response = await sendRegisterAuthCode(user.value.mail)
  if (response.status === 200) {
    message.success(response.data.message)
  } else {
    message.error(response.data.message)
  }
}


// 注册
async function register() {
  let response = await registerUser(user.value)
  console.log(response)
  if (response.status == 200) {
    message.success(response.data.message)
  } else {
    message.error(response.data.message)
  }
}



function handleValidateClick(e: MouseEvent) {
  e.preventDefault()
  formRef.value?.validate((errors) => {
    if (!errors) {
      // 表单校验成功，则进行注册
      register()
    }
  })
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
        :model="user"
        :rules="rules"
        label-placement="left"
        class="register-form"
        ref="formRef"
    >
      <n-grid :cols="1">
        <n-gi :span="1">

          <n-form-item path="username">
            <n-input
                v-model:value="user.username"
                placeholder="用户名"
            >
              <template #prefix>
                <n-icon :component="User"/>
              </template>
            </n-input>
          </n-form-item>
        </n-gi>

        <n-gi :span="1">
          <n-form-item path="password">
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
          <n-form-item path="confirmPassword">
            <n-input
                :disabled="!user.password"
                type="password"
                v-model:value="user.confirmPassword"
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
          <n-form-item path="mail">
            <n-input-group>
              <n-input
                  v-model:value="user.mail"
                  placeholder="邮箱"
              >
                <template #prefix>
                  <n-icon :component="Mail"></n-icon>
                </template>
              </n-input>

              <n-button
                  @click="sendAuthCodeToMail"
              >发送验证码
              </n-button>
            </n-input-group>
          </n-form-item>
        </n-gi>

        <n-gi :span="1">
          <n-form-item>
            <n-input
                v-model:value="user.code"
                :disabled="!isEmailValid"
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
                  @click="handleValidateClick"
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
                <router-link :to="{ name : 'Login' }" class="router-link">已有账号? 去登录</router-link>
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
  text-decoration: none;
  float: right;
}

.user-register-header-title {
  margin-left: 10px;
}
</style>
