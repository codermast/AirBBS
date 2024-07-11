<script setup lang="ts">
import { onMounted, ref } from 'vue';
import Vditor from 'vditor';
import 'vditor/dist/index.css';

const vditor = ref<Vditor | null>(null);

// 接收 options 参数
const props = defineProps<{
  options: IOptions;
}>();

onMounted(() => {
  vditor.value = new Vditor('vditor',
    props.options,
  );
});
</script>

<template>
  <div id="vditor"/>
</template>

<style scoped>
</style>

