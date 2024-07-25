<script setup lang="ts">

import GolangNotes from "@/components/sidebar/GolangNotes.vue";
import FriendLink from "@/components/sidebar/FriendLink.vue";
</script>

<template>
  <n-grid cols="1" y-gap="10">
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