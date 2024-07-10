<script setup lang="ts">
import { onMounted, ref } from "vue";
import axios from "@/api/axios"
import { useRoute } from "vue-router";

const route = useRoute()

let article = ref("");

onMounted(() => {
  let articleID = route.params.articleID
  console.log("articleID", articleID)



  axios.get("/articles/" + articleID).then((res) => {
    console.log(res)
    if (res.status == 200) {
        article.value = res.data.data;
        console.log("article: ", article)
    }
  })
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

    {{ article.content }}

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

</style>