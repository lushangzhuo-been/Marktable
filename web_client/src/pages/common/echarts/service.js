import { COLOR_ECHART } from "@/assets/tool/color_list";
// x轴相关配置
export const xAxis = {
    triggerEvent: true,
    // 坐标轴线设置
    axisLine: {
        show: true,
        lineStyle: {
            color: "#82889E",
            width: 0.5
        }
    },
    // 坐标轴刻度相关设置
    axisTick: {
        show: true,
        lineStyle: {
            color: 'red',
        }
    },
    // 坐标轴刻度标签的相关设置。
    axisLabel: {
        show: true,
        // interval: 'auto',
        // rotate: 0,
        margin: 8, // label与轴线之间的距离
        color: "#40485F",
        interval: 0,
        rotate: 45, //代表逆时针旋转45度
        formatter: function (params) {
            return params.length > 4 ? params.substring(0, 4) + '...' : params
        }
    }
}
// y轴相关配置
export const yAxis = {
    // 坐标轴线设置
    axisLine: {
        show: true,
        lineStyle: {
            color: "#82889E",
            width: 0.5
        }
    },
    // 坐标轴刻度相关设置
    axisTick: {
        show: true,
        lineStyle: {
            color: 'red',
        }
    },
    // 坐标轴刻度标签的相关设置。
    axisLabel: {
        show: true,
        interval: 'auto',
        rotate: 0,
        margin: 8, // label与轴线之间的距离
    }
}
export default {
    // 折线、柱图options
    barLineOptions: {
        color: COLOR_ECHART,
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
            appendToBody: true,
            axisPointer: {
                type: "shadow",
                shadowStyle: {
                    color: "rgb(245, 246, 248)"
                },
                z: -1
            },
            formatter: function (params) {
                let title = "";
                let result = ``;
                params.forEach((item) => {
                    if (item.name) {
                        title = ` <div class="title">
                                      ${item.axisValue}
                                  </div>`;
                        result += `<div class="box-line">
                                        <b class="circle" style="background-color: ${item.color}"></b>
                                        <span class="name" title="${item.name}">
                                            ${item.name}
                                        </span>
                                        <span class="value">${item.value}</span>个
                                    </div>`;
                    }
                });

                return ` <div class="dashboard-echarts-bar-line-tooltips">
                            <div class="overflow">
                                ${result}
                            </div>
                        </div>`;
            }
        },
        grid: {
            top: 20,
            right: 20,
            left: 30,
            bottom: 20,
            backgroundColor: 'red'
        },
        legend: {
            itemGap: 10,
            itemWidth: 14,
            itemHeight: 14,
            inactiveBorderColor: '#fff',
            inactiveBorderWidth: 0,
            data: []
        },
        xAxis: {
            triggerEvent: true,
            type: 'category',
            data: [],
            position: "bottom",
            axisTick: {
                show: false
            },
            axisLabel: {
                color: "#40485F",
                interval: 0,
                rotate: 45, //代表逆时针旋转45度
                formatter: function (params) {
                    return params.length > 4 ? params.substring(0, 4) + '...' : params
                }
            },
            axisLine: {
                lineStyle: {
                    color: "#82889E",
                    width: 0.5
                }
            },

            nameTextStyle: {
                color: '#40485F',
            },
            // nameRotate: 90, // y轴name旋转90度 使其垂直
            // nameGap: 50,  // y轴name与横纵坐标轴线的间距
            // nameLocation: "middle", // y轴name处于y轴的什么位置
        },
        yAxis: {
            type: 'value',
            splitNumber: 6,
            axisLabel: {
                color: '#40485F',
            },
            axisLine: {
                show: true,
                lineStyle: {
                    color: "#82889E",
                    width: 0.5
                }
            },
            splitLine: {
                show: true,
                lineStyle: {
                    color: ['#CDD5E6'],
                    width: 0.5
                }
            },
            // name: '',
            nameTextStyle: {
                color: '#40485F',
            },
            // nameRotate: 90, // y轴name旋转90度 使其垂直
            // nameGap: 50,  // y轴name与横纵坐标轴线的间距
            // nameLocation: "middle", // y轴name处于y轴的什么位置
        },
        series: []
    },
    // 饼图options
    pieOptions: {

        color: COLOR_ECHART,
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
            trigger: "item",
            appendToBody: true,
            axisPointer: {
                type: "shadow",
                shadowStyle: {
                    color: "rgb(245, 246, 248)"
                },
                z: -1
            },
            formatter: function (params) {
                let result = ``;
                if (params.name) {
                    result += `<div class="box-line" >
                                    <b class="circle" style="background-color: ${params.color}"></b>
                                    <span class="name" title="${params.name}">
                                        ${params.name}
                                    </span>
                                    <span class="value">${params.value}</span>个
                                </div>`;
                }

                return ` <div class="dashboard-echarts-pie-tooltips">
                                ${result}
                        </div>`;
            }
        },
        legend: {
            inactiveBorderColor: '#fff',
            inactiveBorderWidth: 0,
            itemGap: 10,
            itemWidth: 14,
            itemHeight: 14,
            padding: [0, 10, 0, 10],
            type: "scroll",
            // right: 0,
            // top: '5%',
            pageIconSize: [10, 10],
            textStyle: {
                overflow: "truncate"
            },
            tooltip: {
                show: true,
                formatter: function (p) {
                    if (p.name.length > 10) {
                        return `<div style="font-size: 12px; max-width: 300px;white-space: normal; padding: 4px 8px;">
                                    ${p.name}
                                <div>`;
                    }
                    return ``
                },
                textStyle: {
                    color: "#171e31",
                },
                borderColor: "#f5f5f5",
                backgroundColor: "#fff",
                borderWidth: 1,
                padding: [0],

            },
            formatter: function (name) {
                return (name.length > 10 ? (name.slice(0, 10) + "...") : name);
            }
        },

        series: [
            {
                name: '',
                type: 'pie',
                radius: '70%',
                center: ['50%', '50%'],
                data: [],

                labelLine: {
                    show: true
                },
                label: {
                    position: 'outer',
                    alignTo: 'none',
                    bleedMargin: 5,
                    show: true,
                    formatter(p) {
                        return `${p.name} ${p.percent === 0 || p.percent ? p.percent + '%' : ''}`
                    },
                },
                itemStyle: {
                    borderWidth: 0.5,
                    borderColor: "#fff",
                },
                emphasis: {
                    label: {
                        show: true,
                        formatter(p) {
                            return `${p.name} ${p.percent === 0 || p.percent ? p.percent + '%' : ''}`
                        },
                    },
                }
            }
        ]
    },
    // 环形饼图options
    bountPieOptions: {
        color: COLOR_ECHART,
        grid: {
            right: 34,
            backgroundColor: 'red'
        },
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
            trigger: "item",
            appendToBody: true,
            axisPointer: {
                type: "shadow",
                shadowStyle: {
                    color: "rgb(245, 246, 248)"
                },
                z: -1
            },
            formatter: function (params) {
                let result = ``;
                if (params.name) {
                    result += `<div class="box-line" >
                                    <b class="circle" style="background-color: ${params.color}"></b>
                                    <span class="name" title="${params.name}">
                                        ${params.name}
                                    </span>
                                    <span class="value">${params.value}</span>个
                                </div>`;
                }

                return ` <div class="dashboard-echarts-pie-tooltips">
                                ${result}
                        </div>`;
            }
        },
        legend: {
            itemGap: 10,
            itemWidth: 14,
            itemHeight: 14,
            padding: [0, 10, 0, 10],
            type: "scroll",
            inactiveBorderColor: '#fff',
            inactiveBorderWidth: 0,
            // right: 0,
            // top: '5%',
            pageIconSize: [10, 10],
            textStyle: {
                overflow: "truncate"
            },
            tooltip: {
                show: true,
                formatter: function (p) {
                    if (p.name.length > 10) {
                        return `<div style="font-size: 12px; max-width: 300px;white-space: normal; padding: 4px 8px;">
                                    ${p.name}
                                <div>`;
                    }
                    return ``
                },
                textStyle: {
                    color: "#171e31",
                },
                borderColor: "#f5f5f5",
                backgroundColor: "#fff",
                borderWidth: 1,
                padding: [0],

            },
            formatter: function (name) {
                return (name.length > 10 ? (name.slice(0, 10) + "...") : name);
            }
        },
        // 侧边legend待优化，先放顶部
        // legend: {
        //     width: 30,
        //     itemGap: 10,
        //     itemWidth: 14,
        //     itemHeight: 14,
        //     type: 'scroll',
        //     orient: 'vertical',
        //     top: 'center',
        //     right: 5,
        //     bottom: 20,
        //     pageIconSize: [10, 10],
        //     textStyle: {
        //         overflow: 'truncate',
        //     },
        //     tooltip: {
        //         show: true,
        //         formatter: function (p) {
        //             return `<div style="font-size: 12px; max-width: 100px; white-space: normal;">
        //                       ${p.name.length > 10 ? p.name : ""}
        //                    <div>`;
        //         },
        //         borderColor: '#f5f5f5',
        //         borderWidth: 1,
        //         padding: [0, 2],
        //     },
        //     formatter: function (name) {
        //         return (name.length > 10 ? (name.slice(0, 10) + "...") : name);
        //     }
        // },
        series: [
            {
                name: '',
                type: 'pie',
                radius: ['50%', '70%'],
                center: ['50%', '50%'],
                data: [],
                label: {
                    show: true,
                    formatter(p) {
                        return `{a|20}\n\r{b|合计}`
                    },
                    rich: {
                        a: {
                            backgroundColor: '#fff',
                            padding: [5, 10, 1, 10],
                            fontSize: '20px'
                        },
                        b: {
                            backgroundColor: '#fff',
                            padding: [1, 10, 5, 10]
                        },
                    },
                    position: 'center'
                },
                // labelLine: {
                //     show: false
                // },
                // label: {
                //     position: 'outer',
                //     alignTo: 'none',
                //     bleedMargin: 5
                // },
                itemStyle: {
                    borderWidth: 1,
                    borderColor: "#fff",
                },
                emphasis: {
                    label: {
                        show: true,
                        formatter(p) {
                            return `{a|${p.value}}\n\r{b|${p.name.length > 4 ? (p.name.slice(0, 4) + "...") : p.name}}`
                        },
                        rich: {
                            a: {
                                backgroundColor: '#fff',
                                padding: [5, 10, 1, 10],
                                fontSize: '20px',

                            },
                            b: {
                                backgroundColor: '#fff',
                                padding: [1, 10, 5, 10]
                            },
                        },
                    },
                }
            }
        ]
    }
}