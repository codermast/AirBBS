<script lang="ts" setup>
import Navbar from "@/layouts/Navbar.vue";
import Footer from "@/layouts/Footer.vue";
import { useLoadingBar } from "naive-ui";

import { computed } from 'vue';
import { dateZhCN, zhCN } from 'naive-ui'

const themeOverrides = computed(() => ({
  common: {
    primaryColorHover: '#3498db', // 这里设置你想要的 hover 颜色
  }
}));




</script>

<template>
  <n-config-provider :theme-overrides="themeOverrides" :locale="zhCN" :date-locale="dateZhCN">
    <n-loading-bar-provider>
      <n-message-provider>
        <n-grid :cols="1" :y-gap="20" class="container">
          <n-gi :span="1" class="header">
            <Navbar></Navbar>
          </n-gi>
          <n-gi :span="1" class="main">
            <router-view></router-view>
          </n-gi>
          <n-gi :span="1">
            <Footer></Footer>
          </n-gi>
        </n-grid>
      </n-message-provider>
    </n-loading-bar-provider>
  </n-config-provider>
</template>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.main {
  flex-grow: 1;
}
</style>