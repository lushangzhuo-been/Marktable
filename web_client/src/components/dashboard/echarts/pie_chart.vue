<template>
    <!-- <div class="chart" ref="chart" :id="'chart'"> -->
    <div class="chart" ref="chart" :id="item.i"></div>
</template>

<script>
export default {
    props: {
        item: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            myChart: null
        };
    },
    mounted() {
        setTimeout(() => {
            this.drawChart();
        }, 200);
        window.addEventListener("resize", () => {
            if (this.myChart) {
                this.myChart.resize();
            }
        });
    },
    methods: {
        redraw() {
            setTimeout(() => {
                this.drawChart();
            }, 200);
        },
        drawChart() {
            if (this.myChart) {
                this.myChart.dispose();
            }
            let chartDom = document.getElementById(this.item.i);
            this.myChart = this.$echarts.init(chartDom);
            let option;
            option = {
                series: [
                    {
                        name: "Access From",
                        type: "pie",
                        radius: ["40%", "70%"],
                        avoidLabelOverlap: false,
                        data: [
                            { value: 1048, name: "Search Engine" },
                            { value: 735, name: "Direct" },
                            { value: 580, name: "Email" },
                            { value: 484, name: "Union Ads" },
                            { value: 300, name: "Video Ads" }
                        ],
                        label: {
                            normal: {
                                show: true,
                                position: "outside",
                                // formatter: '{b}: {c} ({d}%)',
                                labelLine: {
                                    show: true,
                                    length: 20,
                                    length2: 30,
                                    smooth: 0.2
                                }
                            }
                        },
                        itemStyle: {
                            // borderRadius: 10,
                            borderColor: "#fff",
                            borderWidth: 2
                        }
                    }
                ]
            };
            this.myChart.setOption(option);
        }
    }
};
</script>

<style lang="scss" scoped>
::v-deep.chart {
    width: 100%;
    height: 100%;

    div {
        width: 100% !important;
        height: 100% !important;

        canvas {
            width: 100% !important;
            height: 100% !important;
        }
    }
}
</style>
