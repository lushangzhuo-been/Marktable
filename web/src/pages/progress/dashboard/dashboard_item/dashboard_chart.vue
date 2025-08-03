<template>
    <bar-chart-common
        v-if="item.isShowChart"
        ref="chart"
        class="chart"
        :id="`${item.i}`"
        :option="options"
        @bar-line-click="chartClick"
    ></bar-chart-common>
    <!-- <div v-else>暂无图表</div> -->
    <no-data
        v-else
        :isTextShow="false"
        :loading="loading"
        :imgShow="imgShow"
    ></no-data>
</template>

<script>
import _ from "lodash";
import dataHandle from "../data_handle";
import NoData from "@/pages/common/no_data.vue";
import BarChartCommon from "@/pages/common/echarts/basic_chart.vue";
// import service from "@/pages/common/echarts/service.js";
import service from "../service.js";

export default {
    components: {
        NoData,
        BarChartCommon
    },
    props: {
        item: {
            type: Object,
            default: () => {}
        },
        data: {
            type: Array,
            default: () => []
        }
    },
    data() {
        return {
            isShowChart: false,
            options: {},
            imgShow: false,
            loading: true
        };
    },
    watch: {
        // item: {
        //     handler(obj) {
        //         if (obj && obj.data && Object.keys(obj.data).length) {
        //             // this.isShowChart = true;
        //             this.$set(obj, "isShowChart", true);
        //             setTimeout(() => {
        //                 this.options = this.getOptions(obj);
        //             }, 20);
        //             this.$nextTick(() => {
        //                 this.chartResize();
        //             });
        //         } else {
        //             // this.isShowChart = false;
        //             this.$set(obj, "isShowChart", false);
        //             this.options = {};
        //         }
        //         this.loading = false;

        //         setTimeout(() => {
        //             this.imgShow = true;
        //         }, 50);
        //     },
        //     deep: true,
        //     immediate: true
        // },
        data: {
            handler(arr) {
                if (arr && Object.keys(arr).length) {
                    // this.isShowChart = true;
                    this.$set(this.item, "isShowChart", true);
                    setTimeout(() => {
                        this.options = this.getOptions(this.item);
                    }, 20);
                    this.$nextTick(() => {
                        this.chartResize();
                    });
                } else {
                    // this.isShowChart = false;
                    this.$set(this.item, "isShowChart", false);
                    this.options = {};
                }
                this.loading = false;

                setTimeout(() => {
                    this.imgShow = true;
                }, 50);
            },
            deep: true,
            immediate: true
        }
    },
    methods: {
        chartClick(chart, option) {
            this.$emit("on-chart-click", chart, option, this.item);
        },
        getOptions(obj) {
            // 图处理
            if (obj.mode.indexOf("pie") > -1) {
                // 饼图
                return service.getPieOPtions(obj);
            } else if (
                obj.mode === "base_histogram" ||
                obj.mode === "base_line"
            ) {
                // 其他
                return service.getOtherChartOptions(obj);
            } else {
                // 分组和堆叠
                return service.getGroupBarLineOptions(obj);
            }
        },
        chartResize() {
            this.$refs.chart && this.$refs.chart.listener();
        },
        exportImg() {
            this.$refs.chart && this.$refs.chart.exportImg();
        }
    }
};
</script>

<style lang="scss" scoped>
.chart {
    height: 100%;
    width: 100%；;
}
</style>
