<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { getArticleById, getAuthorInfo } from "@/api/article";
import { MdPreview } from "md-editor-v3";
import type { AuthorInfo } from "@/models/article";
import emitter from "@/utils/emitter";
import User from "@/icons/User.vue";
import { Time } from "@vicons/ionicons5";
import View from "@/icons/View.vue";
import { formatDate, timeAgo } from "@/utils/date";

const route = useRoute()


let loading = ref(true);


let article = ref();
let articleUrl = ref(window.location.href);
let authorInfo = ref<AuthorInfo>()

onMounted(async () => {
  let articleID = route.params.articleID
  console.log("articleID", articleID)
  console.log(typeof articleID)

  let articleRet = await getArticleById(articleID as string)

  if (articleRet.status == 200) {
    article.value = articleRet.data.data;
    console.log("article: ", article)
    loading.value = false
  }


  // 获取文章作者信息
  let authorRet = await getAuthorInfo(article.value.author)
  if (authorRet.status == 200) {
    authorInfo.value = authorRet.data.data
    console.log(authorInfo)
  }

  // 发送文章作者信息到侧边栏组件
  emitter.emit("sendArticleAuthorInfo", authorInfo.value)
})


</script>

<template>
  <n-card
      :segmented="{
      content: true,
      footer: 'soft'
    }"
  >
    <template #header>
      <n-skeleton v-if="loading" text width="60%"/>
      <div v-else>

        <h1>{{ article.title }}</h1>
        <n-breadcrumb class="bread-crumbs-nav">

          <n-breadcrumb-item>
            <n-icon :component="User" size="15"></n-icon>
            {{ authorInfo?.nickname == "" ? authorInfo?.username : authorInfo?.nickname }}
          </n-breadcrumb-item>
          <n-breadcrumb-item>
            <n-icon :component="Time" size="15"></n-icon>
            {{ timeAgo(article.publish_time) }}发布
          </n-breadcrumb-item>

          <n-breadcrumb-item>{{ timeAgo(article.update_time) }}来过</n-breadcrumb-item>

          <n-breadcrumb-item>
            <n-icon :component="View" size="15"></n-icon>
            {{ article.views }}
          </n-breadcrumb-item>

        </n-breadcrumb>
      </div>

    </template>

    <n-skeleton v-if="loading" text :repeat="6"/>

    <div class="article-info-content" v-else>
      <MdPreview
          v-model="article.content"
          preview-theme="github"
          :code-foldable="false"
          :show-code-row-number="false"
      ></MdPreview>
    </div>
    <template #footer>
      <n-skeleton v-if="loading" text width="60%"/>

      <div style="font-size: 0.9em;" v-else>

        <n-alert title="Default 类型" type="success">
          <li>本文地址：<a target="_blank" :href="articleUrl">{{ articleUrl }}</a></li>
          <li>转载必须注明作者和本文链接</li>
        </n-alert>
      </div>

    </template>

  </n-card>


</template>

<style scoped>
.bread-crumbs-nav {
  display: flex;
  height: 30px;
}

.article-info-content {
  margin: 0 20px;
}

</style>