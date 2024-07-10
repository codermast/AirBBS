<script lang="ts" setup>
import Navbar from "@/components/Navbar.vue";
import Footer from "@/components/Footer.vue";
</script>

<template>
  <n-message-provider>
    <n-grid :cols="1" :y-gap="20">
      <n-gi :span="1">
        <Navbar></Navbar>
      </n-gi>
      <n-gi :span="1">
        <router-view></router-view>
      </n-gi>
      <n-gi :span="1">
        <Footer></Footer>
      </n-gi>
    </n-grid>
  </n-message-provider>
</template>

<style scoped>

</style>