<script setup lang="ts">

import Golang from "@/icons/Golang.vue";
</script>

<template>
  <div class="footer">
    <n-grid :cols="1" y-gap="10">
      <n-gi :span="1">
        <n-grid :cols="10" >
          <n-gi :span="1"></n-gi>
          <n-gi :span="2">
            <n-card
                :bordered="false"
                class="card-container"
                content-style="color: #ffffff;"
            >
              <template #header>
                <div class="card-header">
                  <n-icon size="40">
                    <golang/>
                  </n-icon>
                  <div class="card-header-title">
                    关于 AirBlog
                  </div>
                </div>
              </template>

              <p style="text-indent: 2em">LearnKu 是终身编程者的修道场做最专业、严肃的技术论坛</p>

              <a>LearnKu 诞生的故事</a>
            </n-card>
          </n-gi>
          <n-gi :span="2">
            <n-card
                :bordered="false"
                class="card-container"
                content-style="color: #ffffff;"
            >
              <template #header>
                <div class="card-header">
                  <div class="card-header-title">
                    资源推荐
                  </div>
                </div>
              </template>
              <div class="article-list">
                <p><a href="https://www.golangnotes.com/">《社区使用指南》</a></p>
                <p><a href="https://www.golangnotes.com/">《文档撰写指南》</a></p>
                <p><a href="https://www.golangnotes.com/">《LearnKu 社区规范》</a></p>
                <p><a href="https://www.golangnotes.com/">《提问的智慧》</a></p>
              </div>

            </n-card>
          </n-gi>
          <n-gi :span="2">

            <n-card
                :bordered="false"
                class="card-container"
                content-style="color: #ffffff;"
            >
              <template #header>
                <div class="card-header">
                  <div class="card-header-title">
                    服务提供商
                  </div>
                </div>
              </template>

              <n-image-group class="card-image-list">
                <n-space>
                  <a href="#">

                    <n-image
                        preview-disabled
                        width="100"
                        src="https://cdn.learnku.com/uploads/banners/bQawWl3vT5dc2lYx5JZ7.png!large"></n-image>
                  </a>

                  <a href="#">

                    <n-image
                        preview-disabled
                        width="100"
                        src="https://cdn.learnku.com/uploads/banners/XPtLlZmIN1cQbLuDFEON.png!large"></n-image>
                  </a>

                  <a href="#">
                    <n-image
                        preview-disabled
                        width="100"
                        src="https://cdn.learnku.com/uploads/banners/JpTCK6OKYBIrBIWdtob8.png!large"></n-image>
                  </a>
                </n-space>
              </n-image-group>


            </n-card>
          </n-gi>
          <n-gi :span="2">

            <n-card
                :bordered="false"
                class="card-container"
                content-style="color: #ffffff;"
            >
              <template #header>
                <div class="card-header">
                  <div class="card-header-title">
                    其他信息
                  </div>
                </div>
              </template>

            </n-card>
          </n-gi>
          <n-gi :span="1"></n-gi>
        </n-grid>

      </n-gi>
      <n-gi :span="1">
        <n-grid :cols="10">
          <n-gi :span="1"></n-gi>
          <n-gi :span="8">
            <n-divider/>
          </n-gi>
          <n-gi :span="1"></n-gi>
        </n-grid>
      </n-gi>

      <n-gi :span="1">
        <p style="text-align: center">粤ICP备18099781号-6 | 粤公网安备 44030502004330号 | 违法和不良信息举报</p>
        <p style="text-align: center">由 友人 设计和编码 ❤</p>
      </n-gi>
    </n-grid>
  </div>
</template>

<style scoped>
.footer {
  color: #ffffff;
  padding: 10px 0;
  background-color: #1b1c1d;
}

.card-header {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center; /* 水平居中 */
  color: #ffffff;
}

.card-header-title {
  margin-left: 10px;
}

.card-container {
  background-color: transparent;
  height: 100%;
  padding: 10px;
}

.card-container:not(:last-child) {
  border-radius: 0;
  border-right: 1px solid #ffffff;
}


.article-list {
  align-items: center;
  text-align: center;
}

.article-list a {
  text-decoration: none;
  color: #848484;
}

.article-list a:hover {
  text-decoration: underline;
  color: #ffffff;
}


.card-image-list {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;
}


</style>