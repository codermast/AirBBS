<script setup lang="ts">
import { onMounted, ref } from 'vue';
import Vditor from 'vditor';
import 'vditor/dist/index.css';

const vditor = ref<Vditor | null>(null);

// 接收 options 参数
const props = defineProps<{
  options: IOptions;
}>();

// 声明事件
const emit = defineEmits(['vditorInstance']);

onMounted(() => {
  // 创建 Vditor 实例，传入 options 对象
  vditor.value = new Vditor('vditor', props.options);
  emit("vditorInstance", vditor.value);
});

function test() {
  console.log('HTML :', vditor.value?.getHTML());
  console.log('Value :', vditor.value?.getValue());
}

</script>

<template>

  <div id="vditor"/>

  <button @click="test">test</button>
</template>

<style scoped>
</style>

