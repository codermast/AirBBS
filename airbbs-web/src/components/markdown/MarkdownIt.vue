<script setup lang="ts">
import MarkdownIt from 'markdown-it';
import { ref, watch } from "vue";
import 'highlight.js/styles/github.css';
import markdownItHighlightjs from 'markdown-it-highlightjs';
import hljs from "highlight.js"; // 根据需要选择代码高亮主题


const props = defineProps({
  content: {
    type: String,
    required: true
  }
});


const md: any = new MarkdownIt({
  html: true, // 允许在 Markdown 中使用 HTML 标签
  highlight: function (str, lang) {
    if (lang && hljs.getLanguage(lang)) {
      try {
        return '<pre style="hljs"><code>' +
            hljs.highlight(lang, str, true).value +
            '</code></pre>';
      } catch (__) {
      }
    }
    return '<pre class="hljs"><code>' + md.utils.escapeHtml(str) + '</code></pre>';
  }
}).use(markdownItHighlightjs);

const renderedHtml = ref(md.render(props.content));


watch(() => props.content, (newContent) => {
  renderedHtml.value = md.render(newContent);
});

</script>

<template>
  <div>
    <div v-html="renderedHtml"></div>
  </div>
</template>

<style scoped>

</style>