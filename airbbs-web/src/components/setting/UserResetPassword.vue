<script setup lang="ts">
import User from "@/icons/User.vue";
import { Glasses, GlassesOutline, Key, LockClosedOutline } from '@vicons/ionicons5'
import ArrowRight from "@/icons/ArrowRight.vue";
import { useRouter } from "vue-router";
import type { FormItemRule, StepsProps } from 'naive-ui'
import { useMessage } from "naive-ui"
import { useUserStore } from "@/stores/userStore";
import ForgetPassword from "@/icons/ForgetPassword.vue";
import { ref } from "vue"
import { matchResetAuthCode, sendResetAuthCode } from "@/api/auth";
import { resetUserPassword } from "@/api/user";

const userStore = useUserStore()
const router = useRouter();
const message = useMessage();

let userReset = ref({
  account: "",
  password: "",
  confirmPassword: "",
  code: "",
})


let isSend = ref(false)


const current = ref<number>(1)
const currentStatus = ref<StepsProps['status']>('process')


let rules = {
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
  code: {
    required: true,
    message: "输入验证码",
    trigger: ['input', 'blur']
  },

}


// 输入时校验
function validatePasswordStartWith(rule: FormItemRule, value: string): boolean {
  return (
      !!userReset.value.password &&
      userReset.value.password.startsWith(value) &&
      userReset.value.password.length >= value.length
  )
}

function validatePasswordSame(rule: FormItemRule, value: string): boolean {
  return value === userReset.value.password
}


// 发送验证码
async function sendResetCode() {
  currentStatus.value = "process"
  let response = await sendResetAuthCode(userReset.value.account)

  if (response.status === 200) {
    message.success(response.data.message)
    isSend.value = true
  } else {
    message.error(response.data.message)
  }
}

// 重置密码
async function resetPassword() {
  let response = await resetUserPassword({
    account: userReset.value.account,
    password: userReset.value.password,
    code: userReset.value.code,
  })

  if (response.status === 200) {
    message.success(response.data.message)
    current.value += 1
  } else {
    message.error(response.data.message)
  }
}


async function oneStepNext() {
  // 1. 判断验证码是否合法
  if (userReset.value.code.length != 6) {
    message.error("验证码格式错误！")
    userReset.value.code = ""
  }else{
    // 格式正确，则向后端判断验证码是否正确

    let response = await matchResetAuthCode(userReset.value.account,userReset.value.code)

    if (response.status === 200) {
      message.success("验证码正确！")
      current.value += 1
    }else{
      message.error(response.data.message)
      currentStatus.value = "error"
    }
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
        <n-icon size="20" :component="ForgetPassword"></n-icon>
        <div class="user-login-header-title">修改密码</div>
      </div>
    </template>

    <n-gi :span="1">

      <n-steps v-model:current="current as number" :status="currentStatus">
        <n-step
            title="校验"
        />
        <n-step
            title="重置"
            disabled
        />
        <n-step
            title="登录"
            disabled
        />
      </n-steps>
    </n-gi>

    <!-- 第一步 -->
    <n-form
        label-placement="left"
        class="login-form"
        v-if="current == 1"
    >
      <n-grid :cols="1">

        <n-gi :span="1" style="margin-bottom: 15px">
          <n-alert type="warning">
            请输入您的邮箱/用户名/手机号，我们将会发送验证码。
          </n-alert>
        </n-gi>


        <n-gi :span="1">

          <n-form-item>
            <n-input-group>
              <n-input
                  v-model:value="userReset.account"
                  placeholder="邮箱/用户名/手机号">
                <template #prefix>
                  <n-icon :component="User"/>
                </template>
              </n-input>

              <n-popover trigger="hover">
                <template #trigger>
                  <n-button
                      :disabled="!userReset.account"
                      @click="sendResetCode"
                  >发送验证码
                  </n-button>
                </template>
                <span>请先输入邮箱/用户名/手机号！</span>
              </n-popover>
            </n-input-group>
          </n-form-item>


          <n-form-item>
            <n-input
                v-model:value="userReset.code"
                :disabled="!isSend"
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
          <n-button
              @click="oneStepNext"
              style="width: 100%"
              type="tertiary"
              icon-placement="right"
          >
            <template #icon>
              <n-icon :component="ArrowRight"></n-icon>
            </template>
            下一步
          </n-button>
        </n-gi>
      </n-grid>
    </n-form>

    <!-- 第二步 -->
    <n-form
        :model="userReset"
        :rules="rules"
        label-placement="left"
        class="login-form"
        v-if="current == 2"
    >
      <n-grid :cols="1">

        <n-gi :span="1" style="margin-bottom: 15px">
          <n-alert type="info">
            请输入新的登录密码。
          </n-alert>
        </n-gi>


        <n-gi :span="1">
          <n-form-item path="password">
            <n-input
                v-model:value="userReset.password"
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
          <n-form-item path="confirmPassword">
            <n-input
                v-model:value="userReset.confirmPassword"
                type="password"
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
          <n-button
              @click="resetPassword"
              style="width: 100%"
              type="success"
              icon-placement="right"
          >
            <template #icon>
              <n-icon :component="ArrowRight"></n-icon>
            </template>
            重置密码
          </n-button>
        </n-gi>
      </n-grid>
    </n-form>

    <!-- 第三步 -->
    <n-form
        label-placement="left"
        class="login-form"
        v-if="current == 3"
    >
      <n-grid :cols="1">

        <n-gi :span="1" style="margin-bottom: 15px">
          <n-alert type="success">
            恭喜你已经成功重置密码，请使用新密码重新登录！
          </n-alert>
        </n-gi>


        <n-gi :span="1">
          <n-button
              @click="router.push({ name : 'Login'} )"
              style="width: 100%"
              type="success"
              icon-placement="right"
          >
            <template #icon>
              <n-icon :component="ArrowRight"></n-icon>
            </template>
            立即登录
          </n-button>
        </n-gi>
      </n-grid>
    </n-form>

  </n-card>
</template>

<style scoped>
.login-card {
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
  margin-top: 20px;
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