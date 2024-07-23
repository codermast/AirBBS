<script setup lang="ts">
import {
  Balloon as BalloonIcon,
  Chatbox as ChatBoxIcon,
  CodeSlash as CodeSlashIcon,
  HelpCircle as HelpCircleIcon,
  Home as HomeIcon,
  LogInOutline,
  LogOutOutline as LogoutIcon,
  Notifications as NotificationsIcon,
  Search,
} from '@vicons/ionicons5'
import MarkPen from "@/icons/Edit.vue";
import { type Component, computed, h, onMounted, ref } from 'vue'
import { NIcon, useMessage } from "naive-ui"
import { useRouter } from "vue-router";
import UserPlus from "@/icons/UserPlus.vue";
import About from "@/icons/AboutOutline.vue";
import { useUserStore } from "@/stores/userStore";
import Hot from "@/icons/Hot.vue";
import type { User } from "@/models/user";

const router = useRouter()
const message = useMessage()
const userStore = useUserStore()

let isLogin = computed(() => userStore.userLoginStatus);

let notifyColor = ref("#c03f53")

let userInfo = ref()

// 通知被点击
function notifyClick() {
  console.log("notify")
}

// user 被点击
function userClick() {
    router.push({ name : "SettingUserInfo" })
}

function notifyOver() {
  notifyColor.value = "green"
}

function notifyLeave() {
  notifyColor.value = "#c03f53"
}

let userOptions = ref([
  {
    label: '个人主页',
    key: "home",
    icon: renderIcon(HomeIcon),
    props: {
      onClick: () => {
        router.push({name: "SettingUserInfo"})
      }
    }
  },
  {
    label: '内容中心',
    key: 'content',
    icon: renderIcon(Hot),
    props: {
      onClick: () => {
        router.push({name: "ArticleCenter"})
      }
    }
  },
  {
    label: '通知中心',
    key: 'notification',
    icon: renderIcon(NotificationsIcon)
  },
  {
    label: '退出登录',
    key: 'logout',
    icon: renderIcon(LogoutIcon),
    props: {
      onClick: () => {
        // 执行退出操作，清空登录记录信息
        userStore.logout();
        message.success('退出成功!')
        router.push({name: "Index"})
      }
    }
  }
])

let editOptions = ref([
  {
    label: '新建文章',
    key: 'profile',
    icon: renderIcon(CodeSlashIcon),
    props: {
      onClick: () => {
        router.push({name: "ArticleCreate"})
      }
    }
  },
  {
    label: '发起讨论',
    key: 'editProfile',
    icon: renderIcon(ChatBoxIcon)
  },
  {
    label: '提个问题',
    key: 'deleteProfile',
    icon: renderIcon(HelpCircleIcon)
  },
  {
    label: '分享动态',
    key: 'status',
    icon: renderIcon(BalloonIcon),
    props: {
      onClick: () => {
        router.push({name: "Blink"})
      }
    }
  }

])

function renderIcon(icon: Component) {
  return () => {
    return h(NIcon, null, {
      default: () => h(icon)
    })
  }
}

onMounted(async () => {

  // 当用户状态为已登录时，获取用户信息
  if (userStore.userLoginStatus) {
    userInfo.value = userStore.userInfo
    console.log("nav-userInfo : ",userInfo)
  }

})


</script>

<template>
  <div class="navbar">
    <n-grid :cols="8">
      <n-gi :span="1">
      </n-gi>
      <n-gi :span="1" class="navbar-brand">
        <n-image
            @click="router.push({name: 'Index'})"
            style="cursor: pointer;"
            preview-disabled
            src="https://cdn.learnku.com//uploads/communities/sNljssWWQoW6J88O9G37.png!/both/44x44"
        ></n-image>
      </n-gi>
      <n-gi :span="2" class="navbar-items">
        <div class="navbar-item" @click="router.push({name : 'Index'})">
          <n-icon :component="HomeIcon" size="18px"></n-icon>
          <div class="navbar-item-title">
            首页
          </div>
        </div>

        <div class="navbar-item" @click="router.push({name : 'Index'})">

          <n-icon :component="ChatBoxIcon" size="18px"></n-icon>
          <div class="navbar-item-title">
            论坛
          </div>
        </div>

        <div class="navbar-item" @click="router.push({name : 'Blink'})">

          <n-icon :component="BalloonIcon" size="18px"></n-icon>
          <div class="navbar-item-title">
            动态
          </div>
        </div>
        <div class="navbar-item" @click="router.push({name : 'Index'})">

          <n-icon :component="About" size="18px"></n-icon>
          <div class="navbar-item-title">
            关于
          </div>
        </div>

      </n-gi>
      <n-gi :span="1" class="navbar-search">
        <n-input round placeholder="全站搜索">
          <template #suffix>
            <n-icon :component="Search"/>
          </template>
        </n-input>
      </n-gi>

      <!-- 未登录展示 -->
      <n-gi
          v-if="!isLogin"
          :span="1"
          class="navbar-user"
      >

        <n-button
            :bordered="false"
            @click="router.push({name: 'Register'})"
        >
          <template #icon>
            <n-icon :component="UserPlus"></n-icon>
          </template>
          注册
        </n-button>

        <n-button
            :bordered="false"
            @click="router.push({name: 'Login'})"
        >
          <template #icon>
            <n-icon :component="LogInOutline"></n-icon>
          </template>
          登录
        </n-button>
      </n-gi>

      <!--  登录展示  -->
      <n-gi
          v-if="isLogin"
          :span="1"
          class="navbar-user"
      >
        <n-dropdown
            trigger="hover"
            :options="editOptions"
            :show-arrow="true"
        >
          <div class="navbar-edit">
            <n-icon size="30px" color="#848484">
              <mark-pen/>
            </n-icon>
          </div>
        </n-dropdown>

        <n-dropdown
            trigger="hover"
            :options="userOptions"
            :show-arrow="true"
        >
          <div class="navbar-user-info" @click="userClick">
            <n-badge @mouseover.stop="notifyOver"
                     @mouseout.stop="notifyLeave"
                     :color="notifyColor"
                     :max="10"
                     :value="20"
                     @click.stop="notifyClick">
              <n-avatar
                  @mouseover.stop=""
                  @mouseout.stop=""
                  @click.stop="userClick"
                  round
                  size="medium"
                  color="white"
                  :src="userInfo?.photo"
              />

            </n-badge>
            <div class="navbar-username">{{ userInfo?.nickname != "" ? userInfo?.nickname : userInfo?.username }}</div>
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

.navbar-items a {
  text-decoration: none;
}

.navbar-item {
  display: flex;
  justify-content: center; /* 在主轴上水平居中 */
  align-items: center;
  height: 60px;
  padding: 0 10px;
  color: #777777;
  cursor: pointer;

}

.navbar-item:hover {
  background-color: #f7f7f7;
  color: #0d0d0d;

}

.navbar-search {
  width: 100%;
  display: flex;
  align-items: center; /* 垂直居中 */
}


.navbar-user {
  display: flex;
  align-items: center;
  width: 100%;
  margin-left: 30px;
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

.navbar-item-title {
  margin-left: 5px;
}
</style>