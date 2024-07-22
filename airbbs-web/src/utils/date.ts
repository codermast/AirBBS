// 格式化日期为指定格式
export function formatDate(dateStr: string) {
    const date = new Date(dateStr)

    // 获取年、月、日
    const year = date.getFullYear();
    const month = date.getMonth() + 1; // 月份从0开始，所以要加1
    const day = date.getDate();

    return `${year}-${month.toString().padStart(2, '0')}-${day.toString().padStart(2, '0')}`
}

export function timeAgo(dateStr: string): string {
    const date = new Date(dateStr);
    const now = new Date();

    // 计算时间差（单位：毫秒）
    const diff = now.getTime() - date.getTime();

    // 转换为秒
    const seconds = Math.floor(diff / 1000);

    // 转换为分钟
    const minutes = Math.floor(seconds / 60);

    // 转换为小时
    const hours = Math.floor(minutes / 60);

    // 转换为天数
    const days = Math.floor(hours / 24);

    // 转换为月数
    const months = Math.floor(days / 30);

    // 转换为年数
    const years = Math.floor(months / 12);

    // 根据时间差大小返回相应的描述
    if (years > 0) {
        return `${years} 年前`;
    } else if (months > 0) {
        return `${months} 个月前`;
    } else if (days > 0) {
        return `${days} 天前`;
    } else if (hours > 0) {
        return `${hours} 小时前`;
    } else if (minutes > 0) {
        return `${minutes} 分钟前`;
    } else {
        return `${seconds} 秒前`;
    }
}