
//buffer转文字
export async function readText(buffer, fileType) {
    return new Promise((resolve, reject) => {
        const viewBuf = new Uint8Array(buffer);
        let str = "";
        for (var index in viewBuf) {
            str += String.fromCharCode(viewBuf[index]);
            // @ts-ignore
            if (index >= 1000) {
                //考虑到效率，只取前1000个用于判断字符集
                break;
            }
        }
        // 判断编码方式的插件，使用上有问题，先写死了编码方式
        //   var codepage = jschardet.detect(str.substring(0, 1000)).encoding;
        let code = fileType == 'text/csv' ? 'GB2312' : 'UTF-8'
        const reader = new FileReader();
        reader.onload = loadEvent => resolve(loadEvent.target.result);
        reader.onerror = e => reject(e);
        reader.readAsText(new Blob([buffer]), code);
    });
}
//文件转buffer
export async function readBuffer(file) {
    return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onload = loadEvent => resolve(loadEvent.target.result);
        reader.onerror = e => reject(e);
        reader.readAsArrayBuffer(file);
    });
}
