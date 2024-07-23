<script setup lang="ts">

import Phone from "@/icons/Phone.vue";
import { Key } from "@vicons/ionicons5";
import { useUserStore } from "@/stores/userStore";
import { ref } from "vue";
import { type FormItemRule, useMessage } from "naive-ui";

const userStore = useUserStore();
const message = useMessage()

let userInfo = userStore.userInfo

let resetTelRequest = ref({
  id: userInfo.id,
  oldTel: userInfo.tel,
  newTel: "",
  code: ""
});

let isSendAuthCode = ref(false)


let formRef = ref();

let rules = ref({
  newTel: [
    {
      required: true,
      message: '请输入新的手机号！',
      trigger: ['input', "blur"]
    },
    {
      validator: validatePhoneNumber,
      message: "手机号格式错误！",
      trigger: ["blur", "tel-input"]
    },
    {
      validator: validatePhoneNumberSame,
      message: "新手机号不能与原手机号重复！",
      trigger: ["blur"]
    }
  ],
})

// 手机号格式校验
function validatePhoneNumber(rule: FormItemRule, value: string): boolean {
  const phoneRegex = /^1[3-9]\d{9}$/;
  return phoneRegex.test(value);
}

// 手机号相同校验
function validatePhoneNumberSame(rule: FormItemRule, value: string) {
  return value != resetTelRequest.value.oldTel
}

// 发送更换手机验证码
function sendChangeTelAuthCode(e: MouseEvent) {
  e.preventDefault()
  formRef.value?.validate((errors: boolean) => {
    if (!errors) {

      // 格式校验成功后，发送验证码
      // TODO:发送手机验证码
      message.success('发送成功！')
      isSendAuthCode.value = true
    } else {
      message.error('手机号填写有误，请核对后重试！')
    }
  })
}
</script>

<template>
  <n-card>
    <template #header>
      <div class="setting-userinfo-header">
        <n-icon :size="30" :component="Phone"></n-icon>
        <div class="setting-userinfo-header-title">
          绑定手机
        </div>
      </div>
    </template>

    <n-form
        ref="formRef"
        :rules="rules"
        :model="resetTelRequest"
        label-placement="left"
        class="edit-userinfo-form"
    >
      <n-grid :cols="1">
        <n-gi :span="1">
          <n-form-item>
            <n-input
                v-model:value="resetTelRequest.oldTel"
                disabled
                placeholder="已绑定手机号"
            >
              <template #prefix>
                <n-icon :component="Phone"></n-icon>
              </template>
            </n-input>
          </n-form-item>
        </n-gi>

        <n-gi :span="1">
          <n-form-item path="newTel">
            <n-input-group>
              <n-input
                  v-model:value="resetTelRequest.newTel"
                  maxlength="11"
                  placeholder="新手机号"
              >
                <template #prefix>
                  <n-icon :component="Phone"></n-icon>
                </template>
              </n-input>

              <n-button
                  @click="sendChangeTelAuthCode"
              >发送验证码
              </n-button>
            </n-input-group>
          </n-form-item>
        </n-gi>

        <n-gi :span="1">
          <n-form-item>
            <n-input
                v-model:value="resetTelRequest.code"
                :disabled="!isSendAuthCode"
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
          <div class="action-button-list">
            <n-button
                class="action-button-item"
                type="primary"
            >保存修改
            </n-button>

            <n-button
                class="action-button-item"
            >重置
            </n-button>
          </div>
        </n-gi>
      </n-grid>


    </n-form>
  </n-card>

</template>

<style scoped>
.setting-userinfo-header {
  display: flex;
  align-items: center;
}

.setting-userinfo-header-title {
  margin-left: 10px;
}
</style>