<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import { useRouter } from "vue-router";
import { getArticleListPage } from "@/api/article";
import type { ArticleListPage, ArticlePageRequest } from "@/models/article";
import { useMessage } from "naive-ui";

const router = useRouter()
const message = useMessage()

let articleListPage = ref<ArticleListPage>();

// 文章数量
let articleCount = ref(0);

let articlePageRequest = ref<ArticlePageRequest>({
  pageNumber: 1,
  pageSize: 10,
})

onMounted(async () => {
  // 获取文章列表
  await fetchArticleList()
})

async function fetchArticleList() {

  let response = await getArticleListPage(articlePageRequest.value)
  console.log(response)
  if (response.status == 200) {
    articleListPage.value = response.data.data;
    console.log("articleListPage: ", articleListPage.value)
  } else {
    message.error(response.data.message)
  }
}


// 文章被点击
function articleClick(articleID: any) {
  // 跳转到指定文章详情
  router.push(`/articles/${ articleID }`);
}

// 监听页码变动
watch(() => articlePageRequest.value.pageNumber, async (newValue) => {
  console.log(`count 变为 ${ newValue }`);
  await fetchArticleList()
})

// 监听页面大小变动
watch(() => articlePageRequest.value.pageSize, (newValue) => {
  console.log(`count 变为 ${ newValue }`);
})


</script>

<template>
  <n-grid :cols="1" y-gap="10">
    <n-gi :span="1">
      <n-grid :cols="1" y-gap="15">
        <n-gi :span="1" v-for="article in articleListPage?.articles">
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
      <n-pagination
          v-model:page="articlePageRequest.pageNumber"
          v-model:page-size="articlePageRequest.pageSize"
          :page-sizes="[10, 20, 30, 40]"
          :page-count="articleListPage?.pageCount"
          show-size-picker
      />

      <!--          show-quick-jumper-->
      <!--          -->
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