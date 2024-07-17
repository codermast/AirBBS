<script setup lang="ts">

import Golang from "@/icons/Golang.vue";
import { Link, LogoGithub, LogoWechat } from "@vicons/ionicons5";
import GolangLogo from "@/icons/GolangLogo.vue";
import Hot from "@/icons/Hot.vue";
</script>

<template>
  <div class="footer">
    <n-grid :cols="1" y-gap="10">
      <n-gi :span="1">
        <n-grid :cols="8">
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
                    关于 Air BBS
                  </div>
                </div>
              </template>

              <p style="text-indent: 2em">Air BBS 是一套基于 Golang 语言、Gin 框架的前后端分离的论坛管理系统。
                前端采用 Vue + NaiveUI 进行开发。
                Air BBS 目前已开源至 Github。
              </p>


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
                  <n-icon size="40" :component="Hot">
                  </n-icon>
                  <div class="card-header-title">
                    资源推荐

                  </div>
                </div>
              </template>
              <div class="article-list">
                <p><a href="https://www.golangnotes.com/golang/core/">《Golang 核心》</a></p>
                <p><a href="https://www.golangnotes.com/backend-tech/cloud-native/docker/">《Docker 指南》</a></p>
                <p><a href="https://www.golangnotes.com/backend-tech/general/git/">《Git 教程》</a></p>
                <p><a href="https://www.golangnotes.com/golang/web/gin/">《Gin 框架》</a></p>
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

                <div class="card-header" style="display: flex;align-items: center;justify-content: center;">
                  <n-icon size="40" :component="Link">

                  </n-icon>
                  <div class="card-header-title">
                    资源推荐

                  </div>
                </div>
              </template>

              <n-grid cols="1">
                <n-gi :span="1">
                  <div style="display: flex;align-items: center;justify-content: center;">
                    <n-image
                        height="50"
                        preview-disabled
                        src="static/images/golangnotes-logo.png"></n-image>
                  </div>

                </n-gi>

                <n-gi :span="1">
                  <div style="display: flex;align-items: center;justify-content: center;">

                    <n-image
                        height="50"

                        preview-disabled
                        src="static/images/codermast-logo.png"></n-image>

                  </div>

                </n-gi>
              </n-grid>


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
        <p style="text-align: center">
          <n-a href="https://github.com/codermast/AirBBS">
            <n-icon :component="LogoGithub"></n-icon>
            Github
          </n-a>
          |
          <n-a href="https://www.golangnotes.com">
            <n-icon :component="GolangLogo"></n-icon>
            GolangNotes
          </n-a>

          |

          <n-tooltip trigger="hover">
            <template #trigger>
              <n-a href="https://github.com/codermast/AirBBS">
                <n-icon :component="LogoWechat"></n-icon>
                codermast888
              </n-a>
            </template>
            <n-image height="300" src="static/images/codermast888.jpg"></n-image>
          </n-tooltip>

        </p>
        <p style="text-align: center">由 友人 设计和编码 ❤</p>
      </n-gi>
    </n-grid>
  </div>
</template>

<style scoped>
.footer {
  color: #ffffff;
  padding: 10px;
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


</style>