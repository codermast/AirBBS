<script setup lang="ts">

import GolangLogo from "@/icons/GolangLogo.vue";
import Edit from "@/icons/Edit.vue";
import ArrowRight from "@/icons/ArrowRight.vue";
import { MdCatalog } from "md-editor-v3";
import { onMounted, ref } from "vue";
import { Add } from "@vicons/ionicons5";
import Message from "@/icons/Message.vue";
import emitter from "@/utils/emitter";
import type { AuthorInfo } from "@/models/article";
import { formatNumber } from "@/utils/format";

const openGolangNotes = () => {
  window.open('https://www.golangnotes.com', '_blank')
}

const id = ref("air-bbs-md-id")

const authorInfo = ref<AuthorInfo | null>(null)



emitter.on("sendArticleAuthorInfo", (articleAuthorInfo : AuthorInfo) => {
  console.log(articleAuthorInfo)
  authorInfo.value = articleAuthorInfo;
})



</script>

<template>
  <n-grid cols="1" y-gap="10">
    <n-gi :span="1">
      <n-card
          :segmented="{
               footer: true,
          }"
      >
        <n-grid :cols="1">
          <n-gi :span="1">

            <n-grid :cols="4" x-gap="10">

              <n-gi :span="1">

                <div style="width: 100%;">
                  <n-avatar
                      style="height: 100%;display: flex;margin: auto"
                      size="40"
                      color="white"
                      src="http://localhost:8080/uploads/logo.png"
                  />
                </div>
              </n-gi>
              <n-gi :span="3">
                <n-grid :cols="1" y-gap="5px">

                  <n-gi span="1" style="height: 20px">
                    <b>{{ authorInfo?.nickname == "" ? authorInfo?.username : authorInfo?.nickname }}</b>
                  </n-gi>

                  <n-gi span="1" style="height: 20px">
                    一个很酷的 Golang 程序员
                  </n-gi>
                </n-grid>
              </n-gi>

            </n-grid>


          </n-gi>

          <n-gi :span="1">
            <n-grid :cols="3">
              <n-gi :span="1">
                <div style="display:flex; justify-content: center;">
                  <n-button text>
                    <n-grid cols="1" y-gap="8px">
                      <n-gi span="1">
                        {{ authorInfo?.articleTotal }}
                      </n-gi>

                      <n-gi span="1">
                        文章

                      </n-gi>
                    </n-grid>

                  </n-button>
                </div>
              </n-gi>

              <n-gi :span="1">
                <div style="display:flex; justify-content: center;">
                  <n-button text>

                    <n-grid cols="1" y-gap="8px">
                      <n-gi span="1">
                        {{ formatNumber(authorInfo?.viewTotal) }}

                      </n-gi>

                      <n-gi span="1">
                        <div>阅读</div>

                      </n-gi>
                    </n-grid>

                  </n-button>
                </div>
              </n-gi>

              <n-gi :span="1">
                <div style="display:flex; justify-content: center;">
                  <n-button text>

                    <n-grid cols="1" y-gap="8px">
                      <n-gi span="1">
                        {{ authorInfo?.fansTotal }}

                      </n-gi>

                      <n-gi span="1">
                        粉丝


                      </n-gi>
                    </n-grid>

                  </n-button>
                </div>
              </n-gi>
            </n-grid>
          </n-gi>


        </n-grid>


        <template #footer>
          <n-grid cols="1" y-gap="10px">
            <n-gi span="1">
              <n-button style="width: 100%">
                <template #icon>
                  <n-icon :component="Add"></n-icon>
                </template>
                关注
              </n-button>
            </n-gi>

            <n-gi span="1">
              <n-button style="width: 100%">
                <template #icon>
                  <n-icon :component="Message"></n-icon>
                </template>

                私信
              </n-button>
            </n-gi>
          </n-grid>
        </template>
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