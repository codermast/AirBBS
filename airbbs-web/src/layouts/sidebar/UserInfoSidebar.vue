<script setup lang="ts">
import { ref } from "vue"
import { ImageSharp as ImageSharpIcon, Menu as MenuIcon,LockClosed as LockClosedIcon } from '@vicons/ionicons5'
import Random from '@/icons/Random.vue'
import Phone from "@/icons/Phone.vue";

let menuIconSize = ref(20)
</script>

<template>
  <n-list hoverable clickable bordered>
    <n-list-item>
      <div class="setting-userinfo-item">
        <n-icon :size="menuIconSize" :component="MenuIcon"></n-icon>
        <div class="setting-userinfo-item-title">
          修改资料
        </div>
      </div>
    </n-list-item>
    <n-list-item>
      <div class="setting-userinfo-item">
        <n-icon
            :size="menuIconSize"
            :component="ImageSharpIcon"
        ></n-icon>
        <div class="setting-userinfo-item-title">
          修改头像
        </div>
      </div>
    </n-list-item>
    <n-list-item>
      <div class="setting-userinfo-item">
        <n-icon
            :size="menuIconSize"
            :component="Random"
        ></n-icon>
        <div class="setting-userinfo-item-title">
          社交账号
        </div>
      </div>
    </n-list-item>

    <n-list-item>
      <div class="setting-userinfo-item">
        <n-icon
            :size="menuIconSize"
            :component="Phone"
        ></n-icon>
        <div class="setting-userinfo-item-title">
          绑定手机
        </div>
      </div>
    </n-list-item>

    <n-list-item>
      <div class="setting-userinfo-item">
        <n-icon
            :size="menuIconSize"
            :component="LockClosedIcon"
        ></n-icon>
        <div class="setting-userinfo-item-title">
          修改密码
        </div>
      </div>
    </n-list-item>
  </n-list>
</template>

<style scoped>
.setting-userinfo-item {
  display: flex;
  align-items: center;
}

.setting-userinfo-item-title {
  margin-left: 5px;
}
</style>