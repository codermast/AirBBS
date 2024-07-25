<script setup lang="ts">
import { computed, onMounted, ref, watch } from "vue";
import { useRouter } from "vue-router";
import { getArticleListPage } from "@/api/article";
import type { ArticleListPage, ArticlePageRequest } from "@/models/article";
import { useMessage } from "naive-ui";
import Hot from "@/icons/Hot.vue";
import MarkdownIt from "@/components/markdown/MarkdownIt.vue";

const router = useRouter()
const message = useMessage()

let articleListPage = ref<ArticleListPage>();

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
    // 更新文章信息
    articleListPage.value = response.data.data;
    console.log("articleListPage: ", articleListPage.value)
    // 更新文章页码信息

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
watch(() => articlePageRequest.value.pageSize, async (newValue) => {
  console.log(`count 变为 ${ newValue }`);
  await fetchArticleList()
})

// 获取 Markdown 中的第一张图片
function getFirstImageUrl(markdownText: string): string {
  // 匹配 Markdown 图片链接的正则表达式
  const regex = /!\[(.*?)\]\((.*?)\)/g;

  // 使用正则表达式匹配第一个图片链接
  const match = regex.exec(markdownText);

  if (match && match.length >= 3) {
    return match[2]; // 返回图片的 URL
  }

  return "https://bing.img.run/rand_1366x768.php?" + Math.random(); // 没有匹配到图片链接
}

// 清除 Markdown 中的特殊字符
function cleanMarkdownText(markdownText: string): string {
  // Regular expressions to match various Markdown elements
  const imageRegex = /!\[.*?\]\(.*?\)/g;         // Markdown image syntax ![alt](url)
  const headerRegex = /^#+\s*/gm;               // Markdown headers (lines starting with #)
  const linkRegex = /\[([^\]]+)\]\([^\)]+\)/g;  // Markdown hyperlinks [text](url)
  const boldItalicRegex = /(\*\*|__)(.*?)\1/g;  // Markdown bold and italic (**bold**, *italic*)

  // Replace images, headers, links, bold and italic with an empty string
  let result = markdownText.replace(imageRegex, '')
      .replace(headerRegex, '')
      .replace(linkRegex, (_, text) => text)
      .replace(boldItalicRegex, (_, text) => text);

  // Remove all newline characters and trim whitespace
  result = result.replace(/\n/g, '').trim();

  return result;
}

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
          >
            <template #header>
              <div class="article-header">
                <n-icon :component="Hot"></n-icon>
                <div class="article-header-title">{{ article.title }}</div>
              </div>
            </template>

            <template #default>
              <n-grid cols="4" x-gap="10px">

                <n-gi :span="1">
                  <n-image width="100%" :src="getFirstImageUrl(article.content)"></n-image>
                </n-gi>

                <n-gi :span="3">
                  <n-ellipsis :line-clamp="4" style="text-indent: 2em" :tooltip="false">

                    <MarkdownIt :content="cleanMarkdownText(article.content)" ></MarkdownIt>
                  </n-ellipsis>
                </n-gi>
              </n-grid>

            </template>
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

.article-header {
  display: flex;
  justify-content: left;
  align-items: center;
}

.article-header-title {
  margin-left: 5px;
}

</style>