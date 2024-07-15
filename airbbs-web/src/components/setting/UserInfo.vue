<script setup lang="ts">
import { onMounted, ref } from 'vue'
import Edit from "@/icons/Edit.vue";
import { getUserById, updateUserInfo } from "@/api/user"

import { useStatusStore } from "@/stores/statusStore";
import { useMessage } from "naive-ui";
import type { UserQueryInfo } from "@/models/user";

const statusStore = useStatusStore();
const message = useMessage()

let userInfo = ref<UserQueryInfo>({
  id: "",
  github: "",
  introduce: "",
  mail: "",
  nickname: "",
  sex: "",
  tel: "",
  username: ""
})


let sexOptions = ref([
  {
    label: '男',
    value: "male"
  },
  {
    label: '女',
    value: "female"
  }
])

// 初始数据
let initialUserInfo = ref<UserQueryInfo>({
  id: "",
  github: "",
  introduce: "",
  mail: "",
  nickname: "",
  sex: "",
  tel: "",
  username: ""
})


onMounted(async () => {
  // 获取用户信息
  let response = await getUserById(statusStore.userLoginId)

  if (response.status == 200) {
    // message.success("查询成功！")
    console.log(response)
    userInfo.value = response.data.data

    if (response.data.data.sex) {
      userInfo.value.sex = "male"
    } else {
      userInfo.value.sex = "female"
    }

    // 保存初始数据
    initialUserInfo = ref<UserQueryInfo>({ ...userInfo.value })
  } else {
    message.error("用户信息获取失败！")
  }
})

// 保存用户信息
async function saveUserInfo() {
  // 转换用户信息状态
  const submitUserInfo = {...userInfo.value, sex: userInfo.value.sex === "male"}
  console.log("123", submitUserInfo)
  // 保存数据
  let response = await updateUserInfo(submitUserInfo)
  console.log(response)
  if (response.status == 200) {
    message.success("更新成功！")
  } else {
    8
    message.error(response.data.message)
  }
}

// 重置用户信息
function resetUserInfo() {
  userInfo.value = initialUserInfo.value
  message.info("重置成功！")
}
</script>

<template>
  <n-card>
    <template #header>
      <div class="setting-userinfo-header">
        <n-icon :size="20" :component="Edit"></n-icon>
        <div class="setting-userinfo-header-title">
          修改资料
        </div>
      </div>
    </template>

    <n-form
        :model="userInfo"
        label-placement="left"
        class="edit-userinfo-form"
    >
      <n-grid :cols="1">
        <n-gi :span="1">
          <n-grid :cols="3" x-gap="10px">
            <n-gi :span="2">
              <n-form-item
                  label-width="70"
                  label="用户名"
                  path="inputValue"
              >
                <n-input
                    v-model:value="userInfo.username"
                    placeholder="username"
                />
              </n-form-item>
            </n-gi>

            <n-gi :span="1" style="line-height: 20px;align-items: center">
              <p>用户名将作为你登录的账号</p>
            </n-gi>
          </n-grid>
        </n-gi>
        <n-gi :span="1">
          <n-grid :cols="3" x-gap="10px">
            <n-gi :span="2">
              <n-form-item
                  label-width="70"
                  label="昵称"
                  path="inputValue"
              >
                <n-input
                    v-model:value="userInfo.nickname"
                    placeholder="nickname"/>
              </n-form-item>
            </n-gi>

            <n-gi :span="1" style="line-height: 20px;align-items: center">
              <p>昵称将作为名称展示在任何需要的地方，如果为空则展示用户名</p>
            </n-gi>
          </n-grid>
        </n-gi>
        <n-gi :span="1">
          <n-grid :cols="3">
            <n-gi :span="2">
              <n-form-item
                  label-width="70"
                  label="性别"
                  path="inputValue"
              >
                <n-select
                    v-model:value="userInfo.sex"
                    :options="sexOptions"
                />
              </n-form-item>
            </n-gi>

            <n-gi :span="1">

            </n-gi>
          </n-grid>
        </n-gi>
        <n-gi :span="1">
          <n-grid :cols="3">
            <n-gi :span="2">
              <n-form-item
                  label-width="70"
                  label="Github"
                  path="inputValue"
              >
                <n-input
                    v-model:value="userInfo.github"
                    placeholder="Github Username"/>
              </n-form-item>
            </n-gi>

            <n-gi :span="1">

            </n-gi>
          </n-grid>
        </n-gi>

        <n-gi :span="1">
          <n-grid :cols="3">
            <n-gi :span="2">
              <n-form-item
                  label-width="70"
                  label="邮箱"
                  path="inputValue"
              >
                <n-input
                    v-model:value="userInfo.mail"
                    placeholder="Email"/>
              </n-form-item>
            </n-gi>

            <n-gi :span="1">

            </n-gi>
          </n-grid>
        </n-gi>
        <n-gi :span="1">
          <n-grid :cols="3" x-gap="10px">
            <n-gi :span="2">
              <n-form-item
                  label-width="70"
                  label="个人简介"
                  path="inputValue"
              >
                <n-input
                    v-model:value="userInfo.introduce"
                    type="textarea"
                    placeholder="Introduce Yourself"
                />
              </n-form-item>
            </n-gi>

            <n-gi :span="1" style="line-height: 20px;align-items: center">
              <p>请一句话介绍你自己，大部分情况下会在你的头像和名字旁边显示</p>
            </n-gi>
          </n-grid>
        </n-gi>
        <n-gi :span="1">
          <div class="action-button-list">
            <n-button
                @click="saveUserInfo"
                class="action-button-item"
                type="primary"
            >保存修改
            </n-button>

            <n-button
                class="action-button-item"
                @click="resetUserInfo"
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

.edit-userinfo-form {
  display: flex;
  align-items: center; /* 垂直居中 */
}


.action-button-item {
  margin-left: 10px;
}
</style>