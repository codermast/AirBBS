<script setup lang="ts">

import { ref } from "vue"
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';


let content = ref("")
</script>

<template>
  <MdEditor v-model="content" />
</template>

<style scoped>

</style>