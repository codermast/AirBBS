<script setup lang="ts">

import { ref } from "vue"
import { MdEditor, NormalToolbar, type ToolbarNames } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import GolangLogoBlack from "@/icons/GolangLogoBlack.vue";

let content = ref("")
const toolbars: ToolbarNames[] = ['bold',
  'underline',
  'italic',
  'strikeThrough',
  '-',
  'title',
  'sub',
  'sup',
  'quote',
  'unorderedList',
  'orderedList',
  'task',
  '-',
  'codeRow',
  'code',
  'link',
  'image',
  'table',
  'mermaid',
  'katex',
  0,
  1,
  2,
  3,
  '-',
  'revoke',
  'next',
  'save',
  '=',
  'prettier',
  'pageFullscreen',
  'fullscreen',
  'preview',
  'previewOnly',
  'htmlPreview',
  'catalog',
];

const handler = () => {
  console.log('NormalToolbar clicked!');
};
</script>

<template>
  <MdEditor
      v-model="content"
      :toolbars="toolbars">
<!--    <template #defToolbars>-->
<!--      <NormalToolbar title="mark" @onClick="handler">-->
<!--        <template #trigger>-->
<!--          <div class="md-editor-icon">-->
<!--            <n-icon :component="GolangLogoBlack" size="24" ></n-icon>-->
<!--          </div>-->
<!--        </template>-->
<!--      </NormalToolbar>-->
<!--    </template>-->
  </MdEditor>

</template>

<style scoped>

</style>