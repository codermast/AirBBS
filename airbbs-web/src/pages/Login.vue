<script setup lang="ts">
import User from "@/pages/User.vue";
import { Glasses, GlassesOutline, LockClosedOutline, LogoGithub, LogoWechat } from '@vicons/ionicons5'
import LogoQQ from '@/icons/LogoQQ.vue'
import ArrowRight from "@/icons/ArrowRight.vue";



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
        label-placement="left"
        class="login-form"
    >
      <n-grid :cols="1">
        <n-gi :span="1">
          <n-form-item>
            <n-input placeholder="邮箱/用户名/手机号">
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
                忘记密码? or 注册
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