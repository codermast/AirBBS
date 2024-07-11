<script setup lang="ts">

import UserInfoSidebar from "@/layouts/sidebar/UserInfoSidebar.vue";
</script>

<template>
<n-grid :cols="17" x-gap="10px">
  <n-gi :span="2"></n-gi>
  <n-gi :span="3">
    <UserInfoSidebar></UserInfoSidebar>
  </n-gi>

  <n-gi :span="10">
    <router-view></router-view>
  </n-gi>


  <n-gi :span="2"></n-gi>

</n-grid>
</template>

<style scoped>

</style>