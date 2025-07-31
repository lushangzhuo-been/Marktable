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
                // this.drawChart()
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
                    // color:['red','yellow'],
                    data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
                },
                yAxis: {
                    type: "value"
                },
                series: [
                    {
                        data: [120, 200, 150, 80, 70, 110, 130],
                        color: ["#65D13A"],
                        type: "bar"
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
