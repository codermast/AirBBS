import { ref } from 'vue'

// 文章类别选项
export let articleTypeOptions = ref([
	{
		label: "草稿",
		value: 0,
	},
	{
		label: '已发布',
		value: 1
	},
	{
		label: '下架',
		value: 2
	},
	{
		label: "错误",
		value: 3,
	}
])