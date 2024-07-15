<script setup lang="ts">

import GolangLogo from "@/icons/GolangLogo.vue";
import Edit from "@/icons/Edit.vue";
import ArrowRight from "@/icons/ArrowRight.vue";
import { MdCatalog } from "md-editor-v3";
import { ref } from "vue";

const openGolangNotes = () => {
  window.open('https://www.golangnotes.com', '_blank')
}

const id = ref("air-bbs-md-id")
</script>

<template>
  <n-grid cols="1" y-gap="10">
    <n-gi :span="1">
      <n-card>
        <n-grid cols="5">
          <n-gi :span="1">
            <n-image
                style="border: #dddddd 1px solid"
                width="100%"
                src="https://cdn.learnku.com/uploads/avatars/82441_1621694272.jpeg!/both/100x100"
            >
            </n-image>
          </n-gi>
          <n-gi :span="2">
            <div style="margin-left: 10px;line-height: 26px">
              友人
            </div>
            <div style="margin-left: 10px;line-height: 26px">

              CoderMast
            </div>
          </n-gi>
          <n-gi :span="2">
            <div style="height: 100%;">
              <n-button style="height: 25px" type="info">发帖</n-button>
              <n-button style="height: 25px;margin-top: 5px" type="success">发帖</n-button>
            </div>
          </n-gi>
        </n-grid>
      </n-card>
    </n-gi>
    <n-gi :span="1">
      <MdCatalog :editor-id="id"></MdCatalog>
    </n-gi>
    <n-gi :span="1">
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
    </n-gi>
    <n-gi :span="1">
      <n-card
          :segmented="true"
      >
        <template #header>
          友情链接
        </template>
        <n-image preview-disabled width="200px" style="display: flex;align-items: center;justify-content: center;"
                 src="https://cdn.learnku.com/assets/images/friends/ruby-china.png"></n-image>
        <n-image preview-disabled width="200px" style="display: flex;align-items: center;justify-content: center;"
                 src="https://cdn.learnku.com/uploads/banners/lHLqvDd0TQZD7CKdmguG.png"></n-image>
        <n-image preview-disabled width="200px" style="display: flex;align-items: center;justify-content: center;"
                 src="https://cdn.learnku.com/uploads/banners/21xSS4vxAXrQYNpDi6wA.png!large"></n-image>
        <n-image preview-disabled width="200px" style="display: flex;align-items: center;justify-content: center;"
                 src="https://cdn.learnku.com/uploads/banners/c4LriLLt6ZDYA8vvZPdf.png"></n-image>
      </n-card>
    </n-gi>
  </n-grid>
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