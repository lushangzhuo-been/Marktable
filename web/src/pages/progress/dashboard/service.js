import { COLOR_ECHART } from "@/assets/tool/color_list";
import service from "@/pages/common/echarts/service.js";
import _ from "lodash";
// 基础柱形图、折线图
export const getOtherChartOptions2 = (obj, isDown = false) => {
    let tmpOptions = _.cloneDeep(service.barLineOptions);
    let data = [];
    obj.data.forEach((item) => {
        tmpOptions.xAxis.data.push({
            value: item.name,
            id: item.user_id || item.status_id || item.document_id || "",
            mode: item.filed_mode
        });
        tmpOptions.legend.data.push(item.name);
        data.push(item.value || item.count);
    });
    tmpOptions.series.push({
        name: "数量",
        data,
        // barWidth: data.length <= 1 ? '80' : "30",
        // barGap: '20%',
        barMaxWidth: 40,
        itemStyle: {
            borderWidth: 0.5,
            borderColor: "#fff",
            borderRadius: 2
        },
        cursor: isDown ? "default" : "pointer",
        type: obj.mode.indexOf("line") > -1 ? "line" : "bar"
    });
    tmpOptions.legend.show = false;
    // if (tmpOptions.xAxis.data.length > 4) {
    //     tmpOptions.dataZoom = [
    //         {
    //             type: "slider", //给x轴设置滚动条
    //             show: true,
    //             bottom: 10,
    //             height: 15,
    //             showDetail: false,
    //             startValue: 0, //滚动条的起始位置
    //             endValue: 4 //滚动条的截止位置（按比例分割你的柱状图x轴长度）
    //         },
    //         {
    //             type: "inside", //设置鼠标滚轮缩放
    //             disabled: false,
    //             height: 15,
    //             startValue: 0,
    //             endValue: 4
    //         }
    //     ]
    //     tmpOptions.grid.bottom = 50
    // } else {
    //     tmpOptions.grid.bottom = 30
    // }
    let xData = tmpOptions.xAxis.data;
    if (
        (xData.length && xData[0].mode === "time_com") ||
        (xData.length && xData[0].mode === "date_com")
    ) {
        if (obj.axis_date_mode === "day") {
            tmpOptions.grid.bottom = 70;
        } else if (obj.axis_date_mode === "month") {
            tmpOptions.grid.bottom = 55;
        } else {
            tmpOptions.grid.bottom = 50;
        }
    } else {
        tmpOptions.grid.bottom = 50;
    }
    tmpOptions.yAxis.name = "计数";
    tmpOptions.grid.top = 40;
    tmpOptions.grid.right = 54;
    tmpOptions.xAxis.name = obj.x_cn;
    tmpOptions.xAxis.nameTruncate = {
        maxWidth: 40,
        ellipsis: "..."
    };
    if (isDown && tmpOptions.series.length) {
        tmpOptions.series[0].color = obj.color;
    }
    tmpOptions.xAxis.axisLabel.formatter = function (params) {
        if (
            (xData.length && xData[0].mode === "time_com") ||
            (xData.length && xData[0].mode === "date_com")
        ) {
            if (obj.axis_date_mode === "day") {
                return params.length > 10
                    ? params.substring(0, 10) + "..."
                    : params;
            } else if (obj.axis_date_mode === "month") {
                return params.length > 7
                    ? params.substring(0, 7) + "..."
                    : params;
            } else {
                return params.length > 4
                    ? params.substring(0, 4) + "..."
                    : params;
            }
        } else {
            return params.length > 4 ? params.substring(0, 4) + "..." : params;
        }
    };
    return tmpOptions;
};
// 饼图和环形图
export const getPieOPtions2 = (obj, isDown = false) => {
    let isDountPie = obj.mode === "donut_pie"; // 环形饼图
    let tmpOptions = {};

    if (isDountPie) {
        tmpOptions = _.cloneDeep(service.bountPieOptions);
        tmpOptions.series[0].data = _.cloneDeep(obj.data);
        tmpOptions.series[0].label.formatter = (p) => {
            return `{a|${obj.data.reduce(
                (total, item) => total + item.value || 0,
                0
            )}}\n\r{b|总计}`;
        };
    } else {
        tmpOptions = _.cloneDeep(service.pieOptions);
        tmpOptions.series[0].data = _.cloneDeep(obj.data);
    }
    tmpOptions.color = COLOR_ECHART;
    if (obj.data.length <= 1) {
        tmpOptions.legend.show = false;
        tmpOptions.series[0].center = ["50%", "50%"];
    } else {
        tmpOptions.series[0].center = ["50%", "57%"];
    }
    tmpOptions.cursor = isDown ? "default" : "pointer";
    if (isDown && tmpOptions.series.length) {
        tmpOptions.series[0].color = obj.color;
    }
    return tmpOptions;
};
// 分组和堆叠折线柱形图
export const getGroupBarLineOptions2 = (obj, isDown = false) => {
    let valueData = obj.data;
    let legenddata = [];
    let group_dimension = [];
    let xAxisName = [];
    valueData.forEach((item) => {
        if (item.group_dimension && item.group_dimension.length) {
            group_dimension.push(item.group_dimension.length);
        }
        xAxisName.push(item.name);
    });
    let xAxis = [];
    let xAxisitem = {
        triggerEvent: true,
        type: "category",
        data: "",
        position: "bottom",
        axisTick: {
            show: false
        },
        axisLabel: {
            color: "#40485F",
            interval: 0,
            rotate: 45 //代表逆时针旋转45度
        },
        axisLine: {
            show: false,
            lineStyle: {
                color: "#82889E",
                width: 0.5
            }
        }
    };
    var emphasisStyle = {
        itemStyle: {
            shadowBlur: 10,
            shadowColor: "rgba(0,0,0,0.3)"
        }
    };
    let xAxisitemdata = [];

    for (let i = 0; i < valueData.length; i++) {
        xAxisitemdata = [];
        for (let i = 0; i < valueData.length; i++) {
            xAxisitemdata.push("");
        }
        xAxisitemdata[i] = {
            value: valueData[i].name,
            id:
                valueData[i].user_id ||
                valueData[i].status_id ||
                valueData[i].document_id ||
                "",
            mode: valueData[i].filed_mode
        };
        xAxisitem.data = JSON.parse(JSON.stringify(xAxisitemdata));
        // xAxisitem.axisLine.show = i === 0 ? true : false
        xAxisitem.axisLine.show = i === 0 ? true : false;
        xAxis.push(JSON.parse(JSON.stringify(xAxisitem)));
    }
    let series = [];
    let data2222 = [];
    valueData.forEach((item) => {
        item.group_dimension.forEach((sub) => {
            data2222.push(sub.name);
        });
    });
    let legenddata1 = Array.from(new Set(data2222));
    let colorMapping = {};
    legenddata = legenddata1.map((item, index) => {
        colorMapping[item] = COLOR_ECHART[index];
        return {
            name: item,
            itemStyle: {
                color: COLOR_ECHART[index]
            }
        };
    });
    for (let j = 0; j < valueData.length; j++) {
        for (let i = 0; i < valueData[j].group_dimension.length; i++) {
            let newArrList = [];

            for (let a = 0; a < j; a++) {
                newArrList.push({});
            }
            newArrList[j] = valueData[j].group_dimension[i];
            for (let k = 0; k < legenddata.length; k++) {
                newArrList.forEach((item) => {
                    item.color = colorMapping[item.name];
                });
                let seriesitem = {
                    type: obj.mode.indexOf("line") > -1 ? "line" : "bar",
                    name: newArrList[j].name,
                    color: colorMapping[newArrList[j].name],
                    // barWidth: "15",
                    barMaxWidth: 40,
                    // barGap: '20%',
                    xAxisIndex: j,
                    itemStyle: {
                        borderWidth: 0.5,
                        borderColor: "#fff",
                        borderRadius: 2
                    },
                    cursor: isDown ? "default" : "pointer",
                    data: newArrList
                };
                if (obj.mode === "stack_histogram") {
                    seriesitem.stack = "ad";
                }
                let a = j % legenddata.length;
                if (valueData[j].group_dimension[i].value !== 0 && k === a) {
                    series.push(JSON.parse(JSON.stringify(seriesitem)));
                }
            }
        }
    }
    let series1 = JSON.parse(JSON.stringify(series));
    let isShowDataZoom = false;
    if (
        Array.from(new Set(series1.map((item) => item.xAxisIndex))).length >
            4 &&
        obj.mode === "stack_histogram"
    ) {
        isShowDataZoom = true;
    }
    let option = {
        tooltip: {
            enterable: true,
            backgroundColor: "rgba(0, 0, 0, 0.65)",
            confine: true,
            textStyle: {
                color: "#fff",
                fontSize: 14
            },
            padding: [0],
            borderWidth: 0,
            trigger: "axis",
            axisPointer: {
                type: "shadow",
                shadowStyle: {
                    color: "rgb(245, 246, 248)"
                },
                z: -1
            },
            appendToBody: true,
            formatter: function (params) {
                let title = "";
                let group = "";
                let result = "";
                params.forEach((item) => {
                    if (item.name) {
                        title = ` <div class="title">
                                        ${item.axisValue}
                                      </div> 
                                `;
                        if (obj.group_cn) {
                            group = `<div class="group">
                                        分组: ${obj.group_cn}
                                      </div>`;
                        }
                        result += `<div class="box-line">
                                        <b class="circle" style="background-color: ${item.color}"></b>
                                        <span class="name" title="${item.name}">
                                            ${item.name}
                                        </span>
                                        <span class="value">${item.value}</span>个
                                    </div>`;
                    }
                });

                if (!params[0].name) {
                    return "";
                }
                return ` <div class="dashboard-echarts-bar-line-tooltips">
                            <div class="overflow">
                                ${title}
                                ${group}
                                ${result}
                            </div>
                        </div>`;
            }
        },
        color: COLOR_ECHART,
        legend: {
            show: legenddata1 && legenddata1.length <= 1 ? false : true,
            itemGap: 10,
            itemWidth: 14,
            itemHeight: obj.mode === "group_line" ? 8 : 14,
            padding: [0, 10, 0, 10],
            type: "scroll",
            // right: 0,
            inactiveBorderColor: "#fff",
            inactiveBorderWidth: 0,
            data: legenddata,
            pageIconSize: [10, 10],
            textStyle: {
                overflow: "truncate"
            },
            tooltip: {
                show: true,
                formatter: function (p) {
                    if (p.name && p.name.length > 10) {
                        return `<div style="font-size: 12px;
                                    max-width: 300px;
                                    white-space: normal;
                                    padding: 0px 4px;
                                    border-radius: 2px;
                                    border: 1px solid #ddd;">
                                        ${p.name}
                                    <div>`;
                    } else {
                        return ``;
                    }
                },
                textStyle: {
                    color: "#171e31"
                },
                backgroundColor: "#fff",
                borderWidth: 0,
                padding: [0]
            },
            formatter: function (name) {
                return name.length > 10 ? name.slice(0, 10) + "..." : name;
            }
        },
        grid: {
            left: 40,
            right: 54,
            top: 40,
            bottom: 50
            // bottom: isShowDataZoom ? 50 : 30
        },
        xAxis: xAxis,
        yAxis: [
            {
                type: "value",
                name: "计数",
                nameGap: 10,
                splitNumber: 6,
                // y轴分割线
                splitLine: {
                    show: true,
                    lineStyle: {
                        color: ["#CDD5E6"],
                        width: 0.5
                    }
                },
                // 刻度标签名称
                axisLabel: {
                    color: "#40485F"
                },
                // 坐标轴轴线
                axisLine: {
                    show: true,
                    lineStyle: {
                        color: "#82889E",
                        width: 0.5
                    }
                },
                nameTextStyle: {
                    color: "#40485F"
                    // padding: [0, 0, 12, 0]
                }
            }
        ],
        series: series1
    };
    let mode = "";
    option.xAxis.forEach((item, index) => {
        if (index === 0 && item.data.length) {
            let hasData = item.data.filter((sub) => {
                return sub && sub.mode;
            });
            if (hasData.length) {
                mode = hasData[0].mode;
            }
        }
        item.axisLabel.formatter = function (params) {
            if (params && (mode === "time_com" || mode === "date_com")) {
                if (obj.axis_date_mode === "day") {
                    return params.length > 10
                        ? params.substring(0, 10) + "..."
                        : params;
                } else if (obj.axis_date_mode === "month") {
                    return params.length > 7
                        ? params.substring(0, 7) + "..."
                        : params;
                } else {
                    return params.length > 4
                        ? params.substring(0, 4) + "..."
                        : params;
                }
            } else {
                return params.length > 4
                    ? params.substring(0, 4) + "..."
                    : params;
            }
        };
        // item.axisLabel.formatter = function (params) {
        //     var newParamsName = "";// 最终拼接成的字符串
        //     var paramsNameNumber = params.length;// 实际标签的个数
        //     var provideNumber = 5;// 每行能显示的字的个数
        //     var rowNumber = Math.ceil(paramsNameNumber / provideNumber);// 换行的话，需要显示几行，向上取整
        //     /**
        //      * 判断标签的个数是否大于规定的个数， 如果大于，则进行换行处理 如果不大于，即等于或小于，就返回原标签
        //      */
        //     // 条件等同于rowNumber>1
        //     if (paramsNameNumber > provideNumber) {
        //         /** 循环每一行,p表示行 */
        //         for (var p = 0; p < rowNumber; p++) {
        //             var tempStr = "";// 表示每一次截取的字符串
        //             var start = p * provideNumber;// 开始截取的位置
        //             var end = start + provideNumber;// 结束截取的位置
        //             // 此处特殊处理最后一行的索引值
        //             if (p == rowNumber - 1) {
        //                 // 最后一次不换行
        //                 tempStr = params.substring(start, paramsNameNumber);
        //             } else {
        //                 // 每一次拼接字符串并换行
        //                 tempStr = params.substring(start, end) + "\n";
        //             }
        //             newParamsName += tempStr;// 最终拼成的字符串
        //         }

        //     } else {
        //         // 将旧标签的值赋给新标签
        //         newParamsName = params;
        //     }
        //     //将最终的字符串返回
        //     return newParamsName
        // }
        // item.label = {
        //     show: true,
        //     textStyle: {
        //         overflow: "truncate"
        //     },
        //     tooltip: {
        //         show: true,
        //         formatter: function (p) {
        //             return `<div style="font-size: 12px; max-width: 100px;white-space: normal;">
        //                           ${p.name.length > 6 ? p.name : ""}
        //                        <div>`;
        //         },
        //         borderColor: "#f5f5f5",
        //         borderWidth: 1,
        //         padding: [0, 2],

        //     },
        //     formatter: function (name) {
        //         return name.length > 5 ? name.slice(0, 5) + "..." : name;
        //     }
        // }
    });
    option.xAxis[0].name = obj.x_cn;
    option.xAxis[0].nameTruncate = {
        maxWidth: 40,
        ellipsis: "..."
    };
    option.xAxis[0].nameTextStyle = {
        color: "#40485F"
    };
    // if (isShowDataZoom) {
    //     option.dataZoom = [
    //         {
    //             type: "slider", //给x轴设置滚动条
    //             show: isShowDataZoom, //flase直接隐藏图形
    //             xAxisIndex: [...valueData.map((item, index) => index)],
    //             bottom: 10,
    //             height: 15,
    //             showDetail: false,
    //             startValue: 0, //滚动条的起始位置
    //             endValue: 4 //滚动条的截止位置（按比例分割你的柱状图x轴长度）
    //         },
    //         {
    //             type: "inside", //设置鼠标滚轮缩放
    //             disabled: false,
    //             height: 15,
    //             xAxisIndex: [...valueData.map((item, index) => index)],
    //             startValue: 0,
    //             endValue: 4
    //             // height: 15,
    //         }
    //     ]
    // }
    if (isDown && option.series.length) {
        option.series[0].color = obj.color;
    }
    if (mode === "time_com" || mode === "date_com") {
        if (obj.axis_date_mode === "day") {
            option.grid.bottom = 70;
        } else if (obj.axis_date_mode === "month") {
            option.grid.bottom = 55;
        } else {
            option.grid.bottom = 50;
        }
    } else {
        option.grid.bottom = 50;
    }
    return option;
};

export default {
    // 饼图和环形图
    getPieOPtions: (obj, isDown) => {
        return getPieOPtions2(obj, isDown);
    },
    // 分组和堆叠折线柱形图
    getGroupBarLineOptions: (obj, isDown) => {
        return getGroupBarLineOptions2(obj, isDown);
    },
    // 基础柱形图、折线图
    getOtherChartOptions: (obj, isDown) => {
        return getOtherChartOptions2(obj, isDown);
    }
};
