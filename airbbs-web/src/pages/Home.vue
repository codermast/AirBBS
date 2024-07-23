<script setup lang="ts">


import emitter from '@/utils/emitter'
import { useLoadingBar } from "naive-ui";
import Article from "@/icons/Article.vue";
import ArticleSidebar from "@/layouts/sidebar/ArticleSidebar.vue";

const loadingBar = useLoadingBar();


// 绑定事件
emitter.on('loadingBarStart', () => {
  loadingBar.start();
})

// 绑定事件
emitter.on('loadingBarError', () => {
  loadingBar.error();
})

// 绑定事件
emitter.on('loadingBarFinish', () => {
  loadingBar.finish();
})

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
        <ArticleSidebar></ArticleSidebar>
      </n-gi>
      <n-gi :span="1">
      </n-gi>
    </n-grid>
  </div>
</template>

<style scoped>

.article-container {
}

</style>