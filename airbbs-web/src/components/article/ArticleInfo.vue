<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import MarkdownIt from "@/components/markdown/MarkdownIt.vue";
import { getArticleById } from "@/api/article";

const route = useRoute()

type Article = {
  title: string
  content: string
  author: string
}

let article = ref<Article>({title: "", content: "", author: ""});

onMounted(async () => {
  let articleID = route.params.articleID
  console.log("articleID", articleID)
  console.log(typeof articleID)

  let articleRet = await getArticleById(articleID as string)

  if (articleRet.status == 200) {
    article.value = articleRet.data.data;
    console.log("article: ", article)
  }
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

      {{ article.title }}
      <n-breadcrumb class="bread-crumbs-nav">

        <n-breadcrumb-item>
          Friend
        </n-breadcrumb-item>
        <n-breadcrumb-item>95</n-breadcrumb-item>
        <n-breadcrumb-item>2 周前</n-breadcrumb-item>
        <n-breadcrumb-item>1 分钟前</n-breadcrumb-item>
      </n-breadcrumb>

    </template>


    <template #header-extra>
      作者：友人
    </template>

    <div class="article-info-content">
      <MarkdownIt :content="article.content"/>
    </div>
    <template #footer>
      <blockquote style="font-size: 0.9em;">
        本作品采用<a href="https://learnku.com/docs/guide/cc4.0/6589">《CC 协议》</a>，转载必须注明作者和本文链接
      </blockquote>
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