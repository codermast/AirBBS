<script setup lang="ts">
import { ref } from 'vue'
import Check from "@/icons/Check.vue";
import { PaperPlane, SaveOutline } from "@vicons/ionicons5";
import File from "@/icons/File.vue";
import Title from "@/icons/Title.vue";
import MdEditor from "@/components/markdown/MdEditor.vue";
import { createArticle } from "@/api/article";
import type { Article } from "@/models/article";
import { useMessage } from "naive-ui"
import { articleTypeOptions } from "@/utils/options";

const message = useMessage()

let article = ref<Article>({
  id: "",
  title: "",
  content: "Hello Air BBS!",
  status: 1,
  author: ""
})

// 文章发布
async function saveArticle() {
  // 在这里可以将 Markdown 内容保存到数据库或进行其他操作
  let response = await createArticle(article.value);

  if (response.status == 200) {
    message.success("发布成功！")
  } else {
    message.error(response.data.message)
  }
}

// 保存草稿
async function saveDraft() {
  // 先置为草稿状态
  article.value.status = 0
  // 在这里可以将 Markdown 内容保存到数据库或进行其他操作
  let response = await createArticle(article.value);

  if (response.status == 200) {
    message.success("保存成功！")
  } else {
    message.error(response.data.message)
  }
}

</script>

<template>
  <n-card style="width: 80vw;margin: auto">
    <template #header>
      <div class="article-post-header">
        <n-icon :component="Check"></n-icon>
        <div class="article-post-header-title">文章发布</div>
      </div>
    </template>
    <n-grid :cols="1" y-gap="10px">
      <n-gi :span="1">
        <div class="article-title">
          <n-input
              v-model:value="article.title"
              placeholder="请输入文章标题"
          >
            <template #prefix>
              <n-icon :component="Title"></n-icon>
            </template>
          </n-input>
        </div>
      </n-gi>
      <n-gi :span="1">
        <div class="editor-container">
          <MdEditor v-model="article.content"></MdEditor>
        </div>
      </n-gi>

      <n-gi :span="1" class="article-post-higher-list">
        <n-collapse default-expanded-names="1" accordion>
          <n-collapse-item title="高级设置" name="1">
            <n-form
                label-placement="left"
            >

              <n-form-item label="文章标签" path="inputValue">


                <n-input placeholder="多个标签以英文逗号隔开">

                </n-input>
              </n-form-item>

              <n-form-item label="文章分类" path="inputValue">


                <n-select
                    placeholder="请选择文章分类"
                />
              </n-form-item>

              <n-form-item label="文章状态" path="inputValue">
                <n-select
                    v-model:value="article.status"
                    :options="articleTypeOptions"
                    placeholder="请选择文章状态">

                </n-select>
              </n-form-item>
            </n-form>
          </n-collapse-item>
        </n-collapse>
      </n-gi>

      <n-gi :span="1" class="article-post-action-list">
        <div class="article-post-action-item">
          <n-button
              type="success"
              @click="saveArticle"
          >
            <template #icon>
              <n-icon :component="PaperPlane"></n-icon>
            </template>
            发布文章
          </n-button>
        </div>

        <div class="article-post-action-item">
          <n-button
              @click="saveDraft"
              type="tertiary">
            <template #icon>
              <n-icon :component="SaveOutline"></n-icon>
            </template>
            保存草稿
          </n-button>
        </div>

        <div class="article-post-markdown-info-wrap">
          <div class="article-post-markdown-info">
            <n-icon :component="File"></n-icon>
            <div class="article-post-markdown-info-title">
              编辑器使用指南
            </div>
          </div>
        </div>
      </n-gi>
      <n-gi :span="1">

      </n-gi>
      <n-gi :span="1">

      </n-gi>
    </n-grid>

  </n-card>
</template>

<style scoped>

.article-post-header {
  display: flex;
  align-items: center;
  justify-content: center;
}

.article-post-header-title {
  margin-left: 5px;
}

.article-post-higher-list {
  padding: 20px;
  border-radius: .28571429rem;
  background-color: #f8f8f9;
  align-items: center;
  box-shadow: inset 0 0 0 1px #d3e0e9, 0 0 0 0 transparent;
}


.article-post-action-list {
  display: flex;
  height: 70px;
  border-radius: .28571429rem;
  background-color: #f8f8f9;
  align-items: center;
  box-shadow: inset 0 0 0 1px #d3e0e9, 0 0 0 0 transparent;
}


.article-post-action-item {
  margin-left: 15px;
}

.article-post-markdown-info-wrap {
  margin-left: auto; /* 将最后一个元素推到右侧 */
  margin-right: 20px;
}

.article-post-markdown-info {
  display: flex;
  align-items: center;
}

.article-post-markdown-info-title {
  margin-left: 5px;
}

.article-title {
}
</style>