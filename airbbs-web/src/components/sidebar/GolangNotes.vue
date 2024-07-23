<script setup lang="ts">

import GolangLogo from "@/icons/GolangLogo.vue";
import ArrowRight from "@/icons/ArrowRight.vue";
import Edit from "@/icons/Edit.vue";

const openGolangNotes = () => {
  window.open('https://www.golangnotes.com', '_blank')
}
</script>

<template>
  <n-card
      :segmented="{
              content: true,
              footer: 'soft'
          }"
  >
    <template #header>
      <div class="setting-userinfo-header">
        <n-icon size="28" :component="GolangLogo">
        </n-icon>
        <div style="margin-left: 5px">Go 技术论坛</div>
      </div>
    </template>
    Go（又称 Golang）是 Google 开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。Go 被誉为是未来的服务器端编程语言。

    <template #footer>
      <div class="action-button-list">
        <n-button
            class="action-button-item"
            type="success"
            @click="openGolangNotes"
        >
          <template #icon>
            <n-icon :component="ArrowRight"></n-icon>
          </template>
          Go 全栈指南
        </n-button>
        <n-button class="action-button-item" type="info">
          <template #icon>
            <n-icon :component="Edit"></n-icon>
          </template>
          发贴
        </n-button>
      </div>
    </template>
  </n-card>

</template>

<style scoped>
.setting-userinfo-header {
  display: flex;
  align-items: center;
}

.action-button-list {
  display: flex;
  margin: auto;
}

.action-button-item {
  width: 50%;
}

.action-button-item:not(:first-child) {
  margin-left: 10px;
}
</style>