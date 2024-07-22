<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { getArticleById, getAuthorInfo } from "@/api/article";
import { MdPreview } from "md-editor-v3";
import type { AuthorInfo } from "@/models/article";
import emitter from "@/utils/emitter";

const route = useRoute()

type Article = {
  title: string
  content: string
  author: string
}

let loading = ref(true);


let article = ref<Article>({title: "", content: "", author: ""});
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
  if (authorRet.status == 200){
    authorInfo.value = authorRet.data.data
    console.log(authorInfo)
  }

  // 发送文章作者信息到侧边栏组件
  emitter.emit("sendArticleAuthorInfo",authorInfo.value)
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
            Friend
          </n-breadcrumb-item>
          <n-breadcrumb-item>95</n-breadcrumb-item>
          <n-breadcrumb-item>2 周前</n-breadcrumb-item>
          <n-breadcrumb-item>1 分钟前</n-breadcrumb-item>
        </n-breadcrumb>
      </div>

    </template>

    <template #header-extra>
      <n-skeleton v-if="loading" text :repeat="6"/>

      <div v-else>
        作者：{{ authorInfo?.nickname == "" ? authorInfo?.username : authorInfo?.nickname }}
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
          <li>本文地址：<a href="{{ articleUrl }}">{{ articleUrl }}</a></li>
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