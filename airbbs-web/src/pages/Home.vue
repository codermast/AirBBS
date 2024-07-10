<script setup lang="ts">

import Sidebar from "@/components/Sidebar.vue";
</script>

<template>
  <div class="article-container">

    <n-grid :cols="9" :x-gap="20">
      <n-gi :span="1">
      </n-gi>
      <n-gi :span="5">
        <router-view></router-view>
      </n-gi>
      <n-gi :span="2">
        <Sidebar></Sidebar>
      </n-gi>
      <n-gi :span="1">
      </n-gi>
    </n-grid>
  </div>
</template>

<style scoped>

.article-container{
}

</style>