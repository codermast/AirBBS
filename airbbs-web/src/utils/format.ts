// 将数据量转换为可读性更好的格式（个、k、m）
export function formatNumber(Total: string | undefined): string {
    if (Total == undefined){
        return "0";
    }

    let num = Number(Total);

    if (num < 1000) {
        return `${num}`;
    } else if (num < 1000000) {
        return `${(num / 1000).toFixed(1)} k`;
    } else {
        return `${(num / 1000000).toFixed(1)} m`;
    }
}