<script setup lang="ts">

import Share from "@/icons/Share.vue";
import Message from "@/icons/Message.vue";
import Star from "@/icons/Star.vue";
import RedPacket from "@/icons/RedPacket.vue";

import { ref } from "vue"
import { createBlink } from "@/api/blink";
import type { BlinkCreateRequest } from "@/models/ro/blink";
import { useMessage } from "naive-ui"

const message = useMessage()

let theme = ["success", "warning", "danger", "info"];

let randomTheme = () => {
  let randomIndex = Math.floor(Math.random() * 4)
  return theme[randomIndex];
}

let blinkContent = ref<BlinkCreateRequest>({
  content:""
})

// 发布 Blink
async function postBlink() {
  if (blinkContent.value.content == "") {
      message.error("请先输入动态信息再发布！")
  } else {
    let response = await createBlink(blinkContent.value);
    if (response.status == 200) {
      message.success("发布成功！")
    }else{
      message.error(response.data.message)
    }
  }
}
</script>

<template>
  <n-grid :cols="1" y-gap="10">
    <n-gi :span="1">
      <n-card
      >
        <n-input
            v-model:value="blinkContent.content"
            type="textarea"
            placeholder="有什么新的观点，说来看看~"
        />

        <template #footer>
          <n-button @click="postBlink">发布</n-button>
        </template>
      </n-card>
    </n-gi>
    <n-gi :span="1">
      <n-grid cols="1" y-gap="0">
        <n-gi :span="1" v-for="blink in 10">
          <n-card

              style="border-radius: 0"
          >
            <template #header>
              <n-grid :cols="9" x-gap="10">
                <n-gi :span="1">

                  <n-avatar
                      round
                      color="white"
                      :size="40"
                      src="https://cdn.learnku.com/uploads/avatars/82441_1621694272.jpeg!/both/100x100"
                  />

                </n-gi>
                <n-gi :span="1">

                  <n-gradient-text :type="randomTheme()">
                    友人
                  </n-gradient-text>

                  <span style="color: #a5a5a5;font-size: 12px">15 天前</span>
                </n-gi>

              </n-grid>
            </template>

            <template #header-extra>
              <n-button>关注</n-button>
            </template>
            {{ blink }}

            <template #footer>
              <n-grid :cols="6">
                <n-gi :span="1"></n-gi>
                <n-gi :span="1">

                  <n-popover trigger="hover">
                    <template #trigger>
                      <n-button :bordered="false" type="default" style="width: 90%">
                        <template #icon>
                          <n-icon :component="Share"></n-icon>
                        </template>
                        分享
                      </n-button>
                    </template>
                    <div style="display: flex">
                      <n-grid :cols="1">
                        <n-gi :span="1">
                          <n-qr-code :value="blink"/>
                        </n-gi>
                        <n-gi :span="1">

                          <div style="display: flex;align-items: center;justify-content: center">扫码分享查看</div>
                        </n-gi>
                      </n-grid>
                    </div>
                  </n-popover>

                </n-gi>
                <n-gi :span="1">

                  <n-button :bordered="false" type="default" style="width: 90%">
                    <template #icon>
                      <n-icon :component="Message"></n-icon>
                    </template>
                    评论
                  </n-button>
                </n-gi>
                <n-gi :span="1">
                  <n-button :bordered="false" type="default" style="width: 90%">
                    <template #icon>
                      <n-icon :component="Star"></n-icon>
                    </template>
                    点赞

                  </n-button>
                </n-gi>
                <n-gi :span="1">
                  <n-button :bordered="false" type="default" style="width: 90%">
                    <template #icon>
                      <n-icon :component="RedPacket"></n-icon>
                    </template>
                    打赏

                  </n-button>
                </n-gi>
                <n-gi :span="1">
                  <n-button :bordered="false" type="default" style="width: 90%">

                    ...

                  </n-button>
                </n-gi>
              </n-grid>
            </template>
          </n-card>
        </n-gi>
      </n-grid>
    </n-gi>
    <n-gi :span="1"></n-gi>
  </n-grid>

</template>

<style scoped>

</style>