<!-- 分页组件 -->
<script setup lang="ts">
import { ref } from 'vue'

let page = ref(1);

</script>

<template>
  <n-space vertical>
    <n-pagination v-model:page="page" :page-count="4" show-quick-jumper>
      <template #goto>
        跳转至
      </template>
    </n-pagination>
  </n-space>
</template>


<style scoped>

</style>
