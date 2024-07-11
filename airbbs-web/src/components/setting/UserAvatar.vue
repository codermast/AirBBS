<script setup lang="ts">

import Edit from "@/icons/Edit.vue";
import { ArchiveOutline as Archive } from "@vicons/ionicons5";
import { type UploadFileInfo, useMessage } from 'naive-ui'

const message = useMessage()

// 图片上传之前校验图片格式，这里是头像上传，故只支持单文件
function beforeUpload(file: UploadFileInfo) {
  if (file.file?.type !== 'image/png' && file.file?.type !== 'image/jpg') {
    message.error('只能上传png格式的图片文件，请重新上传')
    return false
  }
  return true
}
</script>

<template>
  <n-card>
    <template #header>
      <div class="setting-user-avatar-header">
        <n-icon :size="20" :component="Edit"></n-icon>
        <div class="setting-user-avatar-header-title">
          修改头像
        </div>
      </div>
    </template>

    <n-form
        label-placement="left"
        class="edit-user-avatar-form"
    >
      <n-grid :cols="1" y-gap="10px">
        <n-gi :span="1">
          <n-alert title="请注意！" type="warning">
            请勿上传涉及色情、政治等违法头像，情节严重者将会被禁言处理。
          </n-alert>
        </n-gi>
        <n-gi :span="1">
          <n-image

              src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
              class="edit-user-avatar-image"
          />

        </n-gi>
        <n-gi :span="1">
          <n-grid :cols="3" x-gap="10px">
            <n-gi :span="2">
              <n-form-item
                  label-width="70"
                  label="头像上传"
                  path="inputValue"
              >
                <n-upload
                    :multiple="false"
                    directory-dnd
                    action="https://www.mocky.io/v2/5e4bafc63100007100d8b70f"
                    :max="1"
                    :show-file-list="true"
                    accept="image/png,image/jpeg"
                    @before-upload="beforeUpload"
                >
                  <n-upload-dragger>
                    <div style="margin-bottom: 12px">
                      <n-icon size="48" :depth="3">
                        <Archive/>
                      </n-icon>
                    </div>
                    <n-text style="font-size: 16px">
                      点击或者拖动图片到该区域来上传
                    </n-text>
                    <n-p depth="3" style="margin: 8px 0 0 0">
                      请勿上传涉及色情、政治等违法头像，情节严重者将会被禁言处理。
                    </n-p>
                  </n-upload-dragger>
                </n-upload>
              </n-form-item>
            </n-gi>
          </n-grid>
        </n-gi>


        <n-gi :span="1">
          <div class="action-button-list">
            <n-button
                class="action-button-item"
                type="primary"
            >上传头像
            </n-button>

          </div>
        </n-gi>
      </n-grid>


    </n-form>
  </n-card>

</template>

<style scoped>
.setting-user-avatar-header {
  display: flex;
  align-items: center;
}

.setting-user-avatar-header-title {
  margin-left: 10px;
}

.edit-user-avatar-form {
  display: flex;
  align-items: center; /* 垂直居中 */
}


.action-button-item {
  margin-left: 10px;
}

.edit-user-avatar-image {
  border: 1px solid #eee;
  width: 300px;
  height: 300px;
}
</style>