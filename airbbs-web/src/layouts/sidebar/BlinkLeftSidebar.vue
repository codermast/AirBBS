<script setup lang="ts">

import Hot from "@/icons/Hot.vue";
import City from "@/icons/City.vue";
import User from "@/icons/User.vue";
import Lightning from "@/icons/Lightning.vue";
import { Balloon } from "@vicons/ionicons5";

let iconSize = 15
</script>

<template>
  <n-card :bordered="false">

    <n-list hoverable clickable>
      <n-list-item>
        <div class="class-list-item">
          <n-icon :size="iconSize"  :component="Hot"></n-icon>
          <div class="class-list-item-title">
            热门
          </div>
        </div>
      </n-list-item>
      <n-list-item>
        <div class="class-list-item">
          <n-icon :size="iconSize" :component="Balloon"></n-icon>
          <div class="class-list-item-title">
            最新
          </div>
        </div>
      </n-list-item>
      <n-list-item>
        <div class="class-list-item">
          <n-icon  :size="iconSize" :component="User"></n-icon>
          <div class="class-list-item-title">
            我的

          </div>
        </div>
      </n-list-item>
      <n-list-item>
        <div class="class-list-item">
          <n-icon  :size="iconSize" :component="City"></n-icon>
          <div class="class-list-item-title">
            同城
          </div>
        </div>
      </n-list-item>
    </n-list>
  </n-card>
</template>

<style scoped>

.class-list-item {
  display: flex;
  align-items: center;
  justify-content: center;
}

.class-list-item-title {
  margin-left: 2px;
}
</style>