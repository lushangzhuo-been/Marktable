export default {
    data1: [{
        name: 'poppy',
        value: 1
    }, {
        name: 'poppy222',
        value: 3
    }, {
        name: 'poppy3322',
        value: 3
    }, {
        name: 'poppy2252',
        value: 7
    }],
    data2: [{
        name: 'poppy',
        value: 1,
        group_dimension: [{
            name: '111',
            value: 1
        }]
    }, {
        name: 'poppy222',
        value: 3,
        group_dimension: [{
            name: '111',
            value: 1,
        }, {
            name: '222',
            value: 2
        }]
    },
    {
        name: 'poppy333',
        value: 2,
        group_dimension: [{
            name: '333',
            value: 1,
        }]
    }],
    // data1: {
    //     name: ['poppy', 'poppy123'],
    //     value: [1, 2],
    // },
    // data2: {
    //     name: ['poppy', 'poppy123'],
    //     value: {
    //         name: '已完成',
    //         data: [1, 2],
    //     },
    //     value2: {
    //         name: '处理中',
    //         data: [2, 1],
    //     },
    //     value3: {
    //         name: '已关闭',
    //         data: [2, 0],
    //     },
    // },
    chartTypeMpping: {
        基础柱形图: "base_histogram",
        分组柱形图: "group_histogram",
        堆叠柱形图: "stack_histogram",
        基础折线图: "base_line",
        分组折线图: "group_line",
        基础饼图: "base_pie",
        环形饼图: "donut_pie"
    },
    chartTypeParent: {
        donut_pie: ['环形饼图', '饼图'],
        base_pie: ['基础饼图', '饼图'],
        group_line: ['分组折线图', '折线图'],
        base_line: ['基础折线图', '折线图'],
        stack_histogram: ['堆叠柱形图', '柱形图'],
        group_histogram: ['分组柱形图', '柱形图'],
        base_histogram: ['基础柱形图', '柱形图'],
    },
    chartsTypeArr: [
        {
            title: "柱形图",
            typeArr: [
                {
                    name: "基础柱形图",
                    value: 'base_histogram',
                    img: require(`@/assets/image/progress_dashboard/bar_1.png`),
                    active: true
                },
                {
                    name: "分组柱形图",
                    value: 'group_histogram',
                    img: require(`@/assets/image/progress_dashboard/bar_2.png`),
                    active: false
                },
                {
                    name: "堆叠柱形图",
                    value: 'stack_histogram',
                    img: require(`@/assets/image/progress_dashboard/bar_3.png`),
                    active: false
                }
            ]
        },
        {
            title: "饼图",
            typeArr: [
                {
                    name: "基础饼图",
                    img: require(`@/assets/image/progress_dashboard/pie_1.png`),
                    active: false
                },
                {
                    name: "环形饼图",
                    img: require(`@/assets/image/progress_dashboard/pie_2.png`),
                    active: false
                }
            ]
        },
        {
            title: "折线图",
            typeArr: [
                {
                    name: "基础折线图",
                    img: require(`@/assets/image/progress_dashboard/line_1.png`),
                    active: false
                },
                {
                    name: "分组折线图",
                    img: require(`@/assets/image/progress_dashboard/line_2.png`),
                    active: false
                }
            ]
        }
    ],
    options1: {
        grid: {
            top: 20,
            right: 20,
            left: 30,
            bottom: 20
        },
        xAxis: {
            type: 'category',
            data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
        },
        yAxis: {
            type: 'value'
        },
        series: [
            {
                data: [120, 200, 150, 80, 70, 110, 130],
                type: 'bar'
            }
        ]
    },
    options2: {
        xAxis: {
            type: 'category',
            data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
        },
        yAxis: {
            type: 'value'
        },
        series: [
            {
                data: [120, 200, 150, 80, 70, 110, 130],
                type: 'line'
            }
        ]
    },
    options3: {
        // title: {
        //     text: 'Referer of a Website',
        //     subtext: 'Fake Data',
        //     left: 'center'
        // },
        tooltip: {
            trigger: 'item'
        },
        // legend: {
        //     orient: 'vertical',
        //     left: 'right'
        // },
        series: [
            {
                name: 'Access From',
                type: 'pie',
                radius: '50%',
                data: [
                    { value: 1048, name: 'Search' },
                    { value: 735, name: 'Direct' },
                    { value: 580, name: 'Email' },
                    { value: 484, name: 'Union' },
                    { value: 300, name: 'Video' }
                ],
                emphasis: {
                    itemStyle: {
                        shadowBlur: 10,
                        shadowOffsetX: 0,
                        shadowColor: 'rgba(0, 0, 0, 0.5)'
                    }
                }
            }
        ]
    },
    options11: {
        grid: {
            top: 20,
            right: 5,
            left: 30,
            bottom: 20
        },
        xAxis: {
            type: 'category',
            data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
        },
        yAxis: {
            type: 'value'
        },
        series: [
            {
                data: [120, 200, 150, 80, 70, 110, 130],
                type: 'bar'
            }
        ]
    },
    options22: {
        xAxis: {
            type: 'category',
            data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
        },
        yAxis: {
            type: 'value'
        },
        series: [
            {
                data: [120, 200, 150, 80, 70, 110, 130],
                type: 'line'
            }
        ]
    },
    options33: {
        // title: {
        //     text: 'Referer of a Website',
        //     subtext: 'Fake Data',
        //     left: 'center'
        // },
        tooltip: {
            trigger: 'item'
        },
        // legend: {
        //     orient: 'vertical',
        //     left: 'right'
        // },
        series: [
            {
                name: 'Access From',
                type: 'pie',
                radius: '50%',
                data: [
                    { value: 1048, name: 'Search' },
                    { value: 735, name: 'Direct' },
                    { value: 580, name: 'Email' },
                    { value: 484, name: 'Union' },
                    { value: 300, name: 'Video' }
                ],
                emphasis: {
                    itemStyle: {
                        shadowBlur: 10,
                        shadowOffsetX: 0,
                        shadowColor: 'rgba(0, 0, 0, 0.5)'
                    }
                }
            }
        ]
    },
    layout: [
        // {
        //     x: 0,
        //     y: 4,
        //     w: 3,
        //     h: 4,
        //     i: "0",
        //     moved: false,
        //     rename: false,
        //     // type: "pie",
        //     parent: '饼图',
        //     type: '基础饼图',
        //     title: "饼图",
        //     data: {
        //         name: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
        //         value1: [120, 200, 150, 80, 70, 110, 130]
        //     }
        // },
        // {
        //     x: 0,
        //     y: 1,
        //     w: 4,
        //     h: 3,
        //     i: "1",
        //     moved: false,
        //     rename: false,
        //     // type: "bar",
        //     parent: '柱形图',
        //     type: '基础柱形图',
        //     title: "处理人数据",
        //     // mode: ''
        //     data: {
        //         name: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
        //         value1: [120, 200, 150, 80, 70, 110, 130]
        //     }
        // },
        // {
        //     x: 4,
        //     y: 0,
        //     w: 4,
        //     h: 4,
        //     i: "2",
        //     moved: false,
        //     rename: false,
        //     type: "line",
        //     title: "折线图",
        //     data: {
        //         name: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
        //         value1: [120, 200, 150, 80, 70, 110, 130]
        //     }
        // },
        // {
        //     x: 0,
        //     y: 0,
        //     w: 8,
        //     h: 4,
        //     i: "3",
        //     moved: false,
        //     rename: false,
        //     type: "number",
        //     title: "数字"
        // }
    ]
}