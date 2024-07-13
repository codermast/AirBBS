<script setup lang="ts">
import MarkdownIt from 'markdown-it';
import { ref, watch } from "vue";
import 'highlight.js/styles/github.css';
import HighLightJs from 'markdown-it-highlightjs';
import hljs from "highlight.js"; // 根据需要选择代码高亮主题
import { full as emoji } from 'markdown-it-emoji'
import MdContainer from 'markdown-it-container'

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
    return '<pre class="hljs" style="background-color: #1b1c1d;"><code>' + md.utils.escapeHtml(str) + '</code></pre>';
  }
})

md.use(HighLightJs);
md.use(emoji)
md.use(MdContainer, 'custom-container', {
  // 定义自定义容器的处理函数
  render(tokens: any, idx: any) {
    if (tokens[idx].nesting === 1) {
      // 开始标签
      return `<div class="custom-container">\n`;
    } else {
      // 结束标签
      return `</div>\n`;
    }
  }
})

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

.hljs {
  background-color: #1b1c1d;
}
.custom-container {
  border: 1px solid #ccc;
  padding: 10px;
  margin-bottom: 10px;
}
</style>