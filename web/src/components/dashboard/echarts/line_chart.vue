<template>
    <!-- <div class="chart" ref="chart" :id="'chart'"> -->
    <div class="chart" ref="chart" :id="item.i"></div>
</template>

<script>
import { PRIMARY_COLOR_LIST } from "@/assets/tool/color_list";
export default {
    props: {
        item: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            myChart: null,
            PRIMARY_COLOR: PRIMARY_COLOR_LIST.PRIMARY_COLOR
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
                xAxis: {
                    type: "category",
                    data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
                },
                yAxis: {
                    type: "value"
                },
                series: [
                    {
                        data: [150, 230, 224, 218, 135, 147, 260],
                        color: "#65D13A",
                        type: "line"
                    },
                    {
                        data: [250, 330, 324, 318, 235, 247, 360],
                        color: this.PRIMARY_COLOR,
                        type: "line"
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
