<script setup lang="ts">

import Edit from "@/icons/Edit.vue";
import Random from "@/icons/Random.vue";
</script>

<template>
  <n-card>
    <template #header>
      <div class="setting-user-social-header">
        <n-icon :size="20" :component="Random"></n-icon>
        <div class="setting-user-social-header-title">
          社交账号
        </div>
      </div>
    </template>

    <n-form
        label-placement="left"
        class="edit-user-social-form"
    >
      <n-grid :cols="1" y-gap="10px">
        <n-gi :span="1">
          <n-alert title="请注意！" type="warning">
            绑定多个第三方账号，允许你使用任意一个第三方账号登录同一个社区账号。
          </n-alert>
        </n-gi>

        <n-gi :span="1">
          <n-image

              src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
              class="edit-user-social-image"
          />

        </n-gi>

      </n-grid>


    </n-form>
  </n-card>

</template>

<style scoped>
.setting-user-social-header {
  display: flex;
  align-items: center;
}

.setting-user-social-header-title {
  margin-left: 10px;
}

.edit-user-social-form {
  display: flex;
  align-items: center; /* 垂直居中 */
}


.action-button-item {
  margin-left: 10px;
}

.edit-user-social-image {
  border: 1px solid #eee;
  width: 300px;
  height: 300px;
}
</style>