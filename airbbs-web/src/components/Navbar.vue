<script setup lang="ts">
import {
  LogOutOutline as LogoutIcon,
  Pencil as EditIcon,
  PersonCircleOutline as UserIcon,
  Search
} from '@vicons/ionicons5'
import MarkPen from "@/icons/Edit.vue";
import type { Component } from 'vue'
import { h, ref } from "vue"
import { NIcon, useMessage } from "naive-ui"
import { useRouter } from "vue-router";

const router = useRouter()
const message = useMessage()

let notifyColor = ref("#c03f53")

// 通知被点击
function notifyClick() {
  console.log("notify")
}

// user 被点击
function userClick() {
  console.log("user")
}

function notifyOver() {
  notifyColor.value = "green"
}

function notifyLeavef() {
  notifyColor.value = "#c03f53"
}

let userOptions = ref([
  {
    label: '滨海湾金沙，新加坡',
    key: 'marina bay sands',
    disabled: true
  },
  {
    label: '布朗酒店，伦敦',
    key: "brown's hotel, london"
  },
  {
    label: '亚特兰蒂斯巴哈马，拿骚',
    key: 'atlantis nahamas, nassau'
  },
  {
    label: '比佛利山庄酒店，洛杉矶',
    key: 'the beverly hills hotel, los angeles'
  }
])

let editOptions = ref([
  {
    label: '新建文章',
    key: 'profile',
    icon: renderIcon(UserIcon)
  },
  {
    label: '发起讨论',
    key: 'editProfile',
    icon: renderIcon(EditIcon)
  },
  {
    label: '退出登录',
    key: 'logout',
    icon: renderIcon(LogoutIcon)
  }
])

function renderIcon(icon: Component) {
  return () => {
    return h(NIcon, null, {
      default: () => h(icon)
    })
  }
}

function handleSelect(key: string | number) {
  message.info(String(key))
}

function logoClick() {
  router.push({ name : "Index"})
}

</script>

<template>
  <div class="navbar">
    <n-grid :cols="8">
      <n-gi :span="1">
      </n-gi>
      <n-gi :span="1" class="navbar-brand">
          <n-image
              @click="logoClick"
              style="cursor: pointer;"
              preview-disabled
              src="https://cdn.learnku.com//uploads/communities/sNljssWWQoW6J88O9G37.png!/both/44x44"
          ></n-image>
      </n-gi>
      <n-gi :span="2" class="navbar-items">
        <div class="navbar-item"><router-link :to="{ name : 'Index' }">首页</router-link></div>
        <div class="navbar-item">About</div>
        <div class="navbar-item">Contact</div>
      </n-gi>
      <n-gi :span="1" class="navbar-search">
        <n-input round placeholder="搜索">
          <template #suffix>
            <n-icon :component="Search"/>
          </template>
        </n-input>
      </n-gi>

      <n-gi :span="1" class="navbar-user">
        <n-dropdown trigger="hover" :options="editOptions">
          <div class="navbar-edit">
            <n-icon size="30px" color="#848484">
              <mark-pen/>
            </n-icon>
          </div>
        </n-dropdown>

        <n-dropdown trigger="hover" :options="userOptions" @select="handleSelect">
          <div class="navbar-user-info" @click="userClick">
            <n-badge @mouseover.stop="notifyOver"
                     @mouseout.stop="notifyLeavef"
                     :color="notifyColor"
                     :max="10"
                     :value="20" @click.stop="notifyClick">
              <n-avatar
                  @mouseover.stop=""
                  @mouseout.stop=""
                  @click.stop="userClick"
                  round
                  size="medium"
                  src="https://cdn.learnku.com/uploads/avatars/82441_1621694272.jpeg!/both/100x100"
              />

            </n-badge>
            <div class="navbar-username">codermast</div>
          </div>
        </n-dropdown>
      </n-gi>

      <n-gi :span="1">
      </n-gi>

    </n-grid>
  </div>

</template>

<style scoped>

.navbar {
  width: 100vw;
  height: 60px;
  color: black;
  background-color: #fff;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, .1), 0 1px 2px 0 rgba(0, 0, 0, .06);
  display: flex;
  align-items: center; /* 垂直居中导航栏项目 */
}

.navbar-brand {
  display: flex;
  width: 44px;
  margin-left: 2rem;
  justify-content: center; /* 水平居中 */
}

.navbar-items {
  display: flex;
}

.navbar-item {
  display: flex;
  justify-content: center; /* 在主轴上水平居中 */
  align-items: center;
  height: 60px;
  padding: 0 10px;
}

.navbar-item:hover {
  background-color: #f7f7f7;
}

.navbar-search {
  width: 100%;
  display: flex;
  align-items: center; /* 垂直居中 */
}


.navbar-user {
  display: flex;
  align-items: center;
}

.navbar-edit {
  margin-left: 20px;
  color: #848484;
  display: flex;
  justify-content: center; /* 在主轴上水平居中 */
  align-items: center;
  height: 60px;
  padding: 0 10px;
}

.navbar-edit:hover {
  background-color: #f7f7f7;
}


.navbar-user-info {
  display: flex;
  width: 200px;
  color: #848484;
  justify-content: center; /* 在主轴上水平居中 */
  align-items: center;
  height: 60px;
  padding: 0 10px;
  cursor: pointer;

}

.navbar-user-info:hover {
  background-color: #f7f7f7;
}

.navbar-username {
  margin-left: 5px;
  display: flex;
  text-align: center;
  align-items: center;
}

.navbar-notify:hover {
  background-color: #848484;
}
</style>