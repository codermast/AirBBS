<script setup lang="ts">

import BlinkLeftSidebar from "@/layouts/sidebar/BlinkLeftSidebar.vue";
import BlinkRightSidebar from "@/layouts/sidebar/BlinkRightSidebar.vue";
import BlinkMain from "@/components/BlinkMain.vue";
</script>

<template>
  <n-grid cols="8" x-gap="10" style="width: 80vw;margin: 0 auto;">
    <n-gi :span="1">
      <BlinkLeftSidebar></BlinkLeftSidebar>
    </n-gi>
    <n-gi :span="4">
      <BlinkMain></BlinkMain>
    </n-gi>
    <n-gi :span="2">
      <BlinkRightSidebar></BlinkRightSidebar>
    </n-gi>
  </n-grid>
</template>

<style scoped>

</style>