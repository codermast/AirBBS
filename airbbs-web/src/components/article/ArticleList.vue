<script setup lang="ts">
import { onMounted, ref } from "vue";

import PageNum from "@/components/article/Pagination.vue";
import { useRouter } from "vue-router";
import { getArticlePublicList } from "@/api/article";

const router = useRouter()

let articleList = ref();

onMounted(async () => {

  let response = await getArticlePublicList()

  if (response.status == 200) {
    articleList.value = response.data.data;
    console.log("articleList: ", articleList)
  }
})

// 文章被点击
function articleClick(articleID: any) {
  // 跳转到指定文章详情
  router.push(`/articles/${ articleID }`);
}

</script>

<template>
  <n-grid :cols="1" y-gap="10">
    <n-gi :span="1">
      <n-grid :cols="1" y-gap="15">
        <n-gi :span="1" v-for="article in articleList">
          <n-card
              class="articleInfo"
              hoverable
              @click="articleClick(article.id)"
              :title="article.title">


            <n-ellipsis :line-clamp="5" style="text-indent: 2em" :tooltip="false">
              {{ article.content }}
            </n-ellipsis>
          </n-card>
        </n-gi>
      </n-grid>
    </n-gi>
    <n-gi :span="1" style="margin:auto">
      <PageNum></PageNum>
    </n-gi>
  </n-grid>
</template>

<style scoped>

.articleInfo {
  height: 200px;
  cursor: pointer;
}

.articleInfo:not(:first-child) {
  margin-top: 20px;
}

</style>