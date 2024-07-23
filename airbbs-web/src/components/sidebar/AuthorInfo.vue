<script setup lang="ts">

import { Add } from "@vicons/ionicons5";
import Message from "@/icons/Message.vue";
import { formatNumber } from "@/utils/format";
import { ref } from "vue";
import type { AuthorInfo } from "@/models/article";
import emitter from "@/utils/emitter";

const authorInfo = ref<AuthorInfo | null>(null)

emitter.on("sendArticleAuthorInfo", (articleAuthorInfo: AuthorInfo) => {
  authorInfo.value = articleAuthorInfo;
})
</script>

<template>
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
                  :src="authorInfo?.photo"
              />
            </div>
          </n-gi>
          <n-gi :span="3">
            <n-grid :cols="1" y-gap="5px">

              <n-gi span="1" style="height: 20px">
                <b>{{ authorInfo?.nickname == "" ? authorInfo?.username : authorInfo?.nickname }}</b>
              </n-gi>

              <n-gi span="1" style="height: 20px">
                {{ authorInfo?.introduce }}
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

</template>

<style scoped>

</style>