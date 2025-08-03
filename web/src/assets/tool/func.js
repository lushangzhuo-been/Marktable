export const kSeq = (num) => {
    if (num) {
        let str = num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
        return str;
    } else if (num === 0) {
        return num;
    } else {
        return "--";
    }
};
export const emptySpace = (param) => {
    if (param) {
        return param;
    } else if (param === 0) {
        return 0;
    } else {
        return "--";
    }
};
// rgb颜色转化为rgba
export const rgbToRgba = (color, opacity, defaultColor) => {
    let rgbaAttr = color.match(/[\d.]+/g);
    if (rgbaAttr && rgbaAttr.length >= 3) {
        let r, g, b;
        r = rgbaAttr[0];
        g = rgbaAttr[1];
        b = rgbaAttr[2];
        return "rgba(" + r + "," + g + "," + b + "," + opacity + ")";
    } else {
        return defaultColor || "#fff";
    }
}

