<script setup lang="ts">

import Phone from "@/icons/Phone.vue";
import { Key } from "@vicons/ionicons5";
import { useStatusStore } from "@/stores/statusStore";
import { onMounted, ref } from "vue";
import { getUserById } from "@/api/user";
import { useMessage } from "naive-ui";

const statusStore = useStatusStore();
const message = useMessage()

let userId = statusStore.userLoginId

let userInfo = ref();

let tel = ref("")
let code = ref("")

let isSendAuthCode = ref(false)

let resetTelRequest = ref({
  id : userId,
  tel : tel.value,
  code : code.value
})

onMounted(async () => {
  let response = await getUserById(userId)
  console.log(response)

  if (response.status === 200) {
    userInfo.value = response.data.data
  } else {
    message.error("登录用户异常，请重新登录！")

    localStorage.removeItem("userLoginId")
  }
})


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
        label-placement="left"
        class="edit-userinfo-form"
    >
      <n-grid :cols="1">
        <n-gi :span="1">
          <n-form-item >
            <n-input
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
          <n-form-item >
            <n-input-group>
              <n-input
                  v-model:value="tel"
                  placeholder="手机号"
              >
                <template #prefix>
                  <n-icon :component="Phone"></n-icon>
                </template>
              </n-input>

              <n-button
                  @click=""
              >发送验证码
              </n-button>
            </n-input-group>
          </n-form-item>
        </n-gi>

        <n-gi :span="1">
          <n-form-item>
            <n-input
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