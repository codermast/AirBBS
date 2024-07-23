<script setup lang="ts">

import { ref } from "vue";
import GolangNotes from "@/components/sidebar/GolangNotes.vue";
import FriendLink from "@/components/sidebar/FriendLink.vue";
import AuthorInfo from "@/components/sidebar/AuthorInfo.vue";
import { useRoute } from "vue-router"

const route = useRoute();

</script>

<template>
  <n-grid cols="1" y-gap="10">
    <n-gi :span="1">
      <AuthorInfo ></AuthorInfo>
    </n-gi>

    <n-gi :span="1">
      <GolangNotes></GolangNotes>
    </n-gi>
    <n-gi :span="1">
      <FriendLink></FriendLink>
    </n-gi>
  </n-grid>
</template>

<style scoped>

</style>