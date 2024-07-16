<!-- 文章管理列表 -->
<script setup lang="ts">
import { type Component, h, onMounted, ref } from "vue";
import { deleteArticle, getArticleAllList, getArticlePublicList, updateArticleListStatus } from "@/api/article";
import type { DataTableRowKey } from 'naive-ui'
import { NButton, NIcon, NTag, useMessage } from "naive-ui";
import type { Article } from "@/models/article";
import { RouterLink, useRouter } from "vue-router"
import ArticleIcon from "@/icons/Article.vue"
import Draft from "@/icons/Draft.vue";
import Hidden from "@/icons/Hidden.vue";
import Cancel from "@/icons/Cancel.vue";
import AroundCheck from "@/icons/ArroundCheck.vue";
import Multiple from "@/icons/Multiple.vue";

const message = useMessage();
const router = useRouter()

let articleList = ref<Article[]>();

onMounted(async () => {
  await getArticleListInfo()
})

async function getArticleListInfo() {
  // 获取文章列表
  let response = await getArticleAllList();

  if (response.status == 200) {
    console.log("response: ", response);
    articleList.value = response.data.data;
    console.log("articleList: ", articleList.value)
  }else{
    message.error(response.data.message)
    router.push({ name : "Index"})
  }
}

// 删除指定 id 的 Article
const deleteArticleById = (articleId: string) => {
  articleList.value = articleList.value?.filter(article => article.id !== articleId);
};

let selectedRowKeysRef = ref<string[]>([]);

const rowKey = (row: Article) => row.id;

// 处理选中行数据
const handleCheck = (rowKeys: DataTableRowKey[]) => {
  if (articleList.value) {
    selectedRowKeysRef.value = articleList.value?.filter(row => rowKeys.includes(row.id)).map(row => row.id);
  } else {
    console.error('articleList is undefined or null');
    selectedRowKeysRef.value = []; // 可以选择在这里初始化为空数组
  }
};

const columns = [
  {
    type: 'selection',
  },
  {
    title: '编号',
    key: 'id'
  },
  {
    title: '标题',
    key: 'title',
    render(article: Article) {
      return h(RouterLink, {
        to: {name: "ArticleInfo", params: {articleID: article.id}},
        className: "router-link",
      }, {
        default: () => article.title
      })
    }
  },
  {
    title: '作者',
    key: 'author',
  },
  {
    title: '状态',
    key: 'state',
    render(article: Article) {
      return h(
          NTag,
          {
            style: {
              marginRight: '6px'
            },
            type: getStatusType(article.status),
            bordered: false
          },
          {
            default: () => getStatusName(article.status)
          }
      )
    }
  },
  {
    title: '操作',
    key: 'actions',
    render(article: Article) {
      return h('div', {style: 'display: flex;align-items: center;margin-left : 10px;'}, [
        h(
            NButton,
            {
              strong: true,
              tertiary: true,
              onClick: () => {
                router.push({
                  name: "ArticleModify",
                  query: {
                    id: article.id,
                  }
                })
              }
            },
            {default: () => '编辑'}
        ),
        h(
            NButton,
            {
              strong: true,
              type: "error",
              onClick: async () => {
                // 执行文章删除操作
                let response = await deleteArticle(article.id)

                if (response.status === 200) {
                  message.success("删除成功！")
                  deleteArticleById(article.id)
                } else {
                  message.error(response.data.message)
                }

              }
            },
            {default: () => '删除'}
        )
      ])


    }
  }
]

function getStatusType(status: number) {
  switch (status) {
    case 0:
      return "info"
    case 1:
      return "success"
    case 2:
      return "warning"
    default:
      return "error"
  }
}

function getStatusName(status: number) {
  switch (status) {
    case 0:
      return "草稿"
    case 1:
      return "已发布"
    case 2:
      return "隐藏"
    default:
      return "禁用"
  }
}

let pagination = {
  pageSize: 10
}

function renderIcon(icon: Component) {
  return () => {
    return h(NIcon, null, {
      default: () => h(icon)
    })
  }
}


let options = ref([
  {
    label: '发布',
    key: 'publish',
    icon: renderIcon(AroundCheck),
    props: {
      onClick: async () => {
        let response = await updateArticleListStatus(selectedRowKeysRef.value, 1)

        if (response.status == 200) {
          message.success("批量发布成功！")
        } else {
          message.error("批量发布失败！")
        }
        await getArticleListInfo()
      }
    }
  },
  {
    label: '转草稿',
    key: 'draft',
    icon: renderIcon(Draft),
    props: {
      onClick: async () => {
        let response = await updateArticleListStatus(selectedRowKeysRef.value, 0)

        if (response.status == 200) {
          message.success("批量转草稿成功！")
        } else {
          message.error("批量转草稿失败！")
        }
        await getArticleListInfo()
      }
    }
  },
  {
    label: '隐藏',
    key: 'hidden',
    icon: renderIcon(Hidden),
    props: {
      onClick: async () => {
        let response = await updateArticleListStatus(selectedRowKeysRef.value, 2)

        if (response.status == 200) {
          message.success("批量隐藏成功！")
        } else {
          message.error("批量隐藏失败！")
        }
        await getArticleListInfo()
      }
    }
  },
  {
    label: '禁用',
    key: 'shelves',
    icon: renderIcon(Cancel),
    props: {
      onClick: async () => {
        let response = await updateArticleListStatus(selectedRowKeysRef.value, 3)

        if (response.status == 200) {
          message.success("批量禁用成功！")
        } else {
          message.error("批量禁用失败！")
        }
        await getArticleListInfo()
      }
    }
  },
])

</script>
<template>
  <n-grid :cols="1" y-gap="10" style="width: 95%;margin: auto">
    <n-gi :span="1">
      <div class="content-center-head-title">
        <n-icon :component="ArticleIcon" size="30"></n-icon>
        <h2 style="margin-left: 5px">文章中心</h2>
      </div>
    </n-gi>
    <n-gi :span="1">
      <div class="content-center-action-list">
        <n-button class="content-center-action-item">
          <template #icon>
            <n-icon :component="ArticleIcon"></n-icon>
          </template>
          文章发布
        </n-button>
        <n-dropdown trigger="hover" :options="options">
          <n-button class="content-center-action-item">
            <template #icon>
              <n-icon :component="Multiple"></n-icon>
            </template>
            批量操作
          </n-button>
        </n-dropdown>

      </div>
    </n-gi>
    <n-gi :span="1">
      <n-data-table
          max-height="70vh"
          :columns="columns"
          :data="articleList"
          :pagination="pagination"
          :row-key="rowKey"
          @update:checked-row-keys="handleCheck"
      />
    </n-gi>
  </n-grid>

</template>

<style scoped>
.content-center-head-title {
  display: flex;
  align-items: center;
  justify-content: center;
}

.content-center-action-list {
  display: flex;
}


.content-center-action-item:not(:first-child) {
  margin-left: 5px;
}
</style>