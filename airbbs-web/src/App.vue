<script lang="ts" setup>
import Navbar from "@/components/Navbar.vue";
import Footer from "@/components/Footer.vue";

import { computed } from 'vue';

const themeOverrides = computed(() => ({
  common: {
    primaryColorHover: '#3498db', // 这里设置你想要的 hover 颜色
  }
}));
</script>

<template>
  <n-config-provider :theme-overrides="themeOverrides">
    <n-message-provider>
      <n-grid :cols="1" :y-gap="20">
        <n-gi :span="1">
          <Navbar></Navbar>
        </n-gi>
        <n-gi :span="1">
          <router-view></router-view>
        </n-gi>
        <n-gi :span="1">
          <Footer></Footer>
        </n-gi>
      </n-grid>
    </n-message-provider>
  </n-config-provider>
</template>

<style scoped>

</style>