<script setup lang="ts">

import { LogoGithub, LogoWechat } from "@vicons/ionicons5";

import { h, ref } from "vue"
import { NIcon } from "naive-ui";
import Random from "@/icons/Random.vue";
import Check from "@/icons/Check.vue";
import Times from "@/icons/Times.vue";
import ArrowRight from "@/icons/ArrowRight.vue";
import WeChatGZH from "@/icons/WeChatGZH.vue";
import QQ from "@/icons/QQ.vue";
import FeiShu from "@/icons/FeiShu.vue";
import DingDing from "@/icons/DingDing.vue";


let columns = ref([
  {
    title: '第三方平台',
    key: 'threePart',
    render(row: Social) {
      return h('div', {style: 'display: flex;align-items: center;'}, [
        h(NIcon, {size: 20, component: row.icon}), // 使用 h 函数创建 NIcon 组件，并传递 props
        h('div', {style: 'margin-left: 5px;'}, row.threePart)]);
    }
  },
  {
    title: '状态',
    key: 'status',
    render(row: Social) {

      return h('div', {style: 'display: flex;align-items: center;'},
          [
            h(NIcon, {size: 20, component: row.status ? Check : Times}),
            // 使用 h 函数创建 NIcon 组件，并传递 props
            h('div', {style: 'margin-left: 5px;'}, row.status ? "已绑定" : "未绑定")
          ]);
    }
  },
  {
    title: '操作',
    key: 'operation',
    render(row: Social) {
      return h('div', {style: 'display: flex;align-items: center;'},
          [
            h(NIcon, {size: 20, component: row.status ? Times : ArrowRight}), // 使用 h 函数创建 NIcon 组件，并传递 props
            h(
                'div',
                {style: 'margin-left: 5px;'},
                [h("a", {
                  href: row.status ? row.bindUrl : row.removeUrl,
                  target: "_blank",
                  style : "color : black;text-decoration : none"
                }, row.status ? "解除绑定" : "前往绑定")],
            )
          ]);
    }
  }])

type Social = {
  threePart: string,
  status: boolean,
  icon: any,
  bindUrl: string,
  removeUrl: string,
}

let data = ref<Social[]>([
  {
    threePart: "微信",
    status: true,
    icon: LogoWechat,
    bindUrl: "string",
    removeUrl: "string",
  },
  {
    threePart: "微信公众号",
    status: true,
    icon: WeChatGZH,
    bindUrl: "string",
    removeUrl: "string",
  },
  {
    threePart: "Github",
    status: true,
    icon: LogoGithub,
    bindUrl: "string",
    removeUrl: "string",
  },
  {
    threePart: "QQ",
    status: false,
    icon: QQ,
    bindUrl: "string",
    removeUrl: "string",
  },
  {
    threePart: "飞书",
    status: true,
    icon: FeiShu,
    bindUrl: "string",
    removeUrl: "string",
  },
  {
    threePart: "钉钉",
    status: true,
    icon: DingDing,
    bindUrl: "string",
    removeUrl: "string",
  },
])
</script>

<template>
  <n-card
      class="social-card"
  >
    <template #header>
      <div class="setting-user-social-header">
        <n-icon :size="20" :component="Random"></n-icon>
        <div class="setting-user-social-header-title">
          社交账号
        </div>
      </div>
    </template>

    <n-form
        label-placement="left"
        class="edit-user-social-form"
    >
      <n-grid :cols="1" y-gap="10px">
        <n-gi :span="1">
          <n-alert type="warning">
            绑定多个第三方账号，允许你使用任意一个第三方账号登录同一个社区账号。
          </n-alert>
        </n-gi>

        <n-gi :span="1">
          <n-data-table
              :columns="columns"
              :data="data"
              :bordered="false"
          />

        </n-gi>

      </n-grid>


    </n-form>
  </n-card>

</template>

<style scoped>

.setting-user-social-header {
  display: flex;
  align-items: center;
}

.setting-user-social-header-title {
  margin-left: 10px;
}

.edit-user-social-form {
  display: flex;
  align-items: center; /* 垂直居中 */
}

</style>