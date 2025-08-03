<template>
    <div style="width: 100%; height: 100%">
        <div :ref="id" :option="option" style="width: 100%; height: 100%"></div>
    </div>
</template>
<script>
import _ from "lodash";
import $ from "jquery";
export default {
    props: {
        id: {
            type: String,
            default: ""
        },
        option: {
            type: Object,
            default: () => {}
        },
        needToggleY: {
            type: Boolean,
            default: false
        },
        leftSort: {
            type: Boolean,
            default: false
        }
    },
    data() {
        return {
            newOptions: {},
            myChart: null
        };
    },
    watch: {
        option: {
            handler(obj) {
                if (obj && Object.keys(obj).length) {
                    this.newOptions = _.cloneDeep(obj);
                    this.$nextTick(() => {
                        this.drawChart(obj);
                    });
                }
            },
            deep: true,
            immediate: true
        }
    },
    methods: {
        drawChart(val) {
            if (val && val.series) {
                if (this.myChart) {
                    this.myChart.dispose();
                }
                let that = this;
                var extension = (myChart) => {
                    var id = document.getElementById("extension");
                    if (!id) {
                        var div =
                            "<div id='extension' class=\"display-none\"></div>";
                        $("html").append(div);
                    }
                    myChart.on("mouseover", (params) => {
                        if (
                            params.componentType === "xAxis" &&
                            params.targetType === "axisLabel"
                        ) {
                            $("#extension")
                                .css({
                                    position: "absolute",
                                    top: "-10px",
                                    color: "#333",
                                    border: "1px solid #ddd",
                                    "border-radius": "2px",
                                    "background-color": "#fff",
                                    "font-size": "12px",
                                    padding: "2px 4px",
                                    display: "inline",
                                    "z-index": 9999
                                })
                                .text(params.value);
                            $("html").mousemove((event) => {
                                var xx = event.pageX - 30;
                                var yy = event.pageY + 20;
                                $("#extension").css("top", yy).css("left", xx);
                            });
                        } else if (
                            params.componentType === "xAxis" &&
                            params.targetType === "axisName"
                        ) {
                            $("#extension")
                                .css({
                                    position: "absolute",
                                    top: "-10px",
                                    color: "#333",
                                    border: "1px solid #ddd",
                                    "border-radius": "2px",
                                    "background-color": "#fff",
                                    "font-size": "12px",
                                    padding: "2px 4px",
                                    display: "inline"
                                })
                                .text(params.name);
                            $("html").mousemove((event) => {
                                var xx = event.pageX - 30;
                                var yy = event.pageY + 20;
                                $("#extension").css("top", yy).css("left", xx);
                            });
                        } else {
                            $("#extension").css("display", "none");
                        }
                    });
                    //  移入事件如果被多次触发，则hover的时候无法显示全部
                    myChart.on("mouseout", function (params) {
                        $("#extension").css("display", "none");
                    });
                    myChart.on("legendselectchanged", (params) => {
                        that.$emit("legend-change", params);
                    });
                    myChart.on("click", (params) => {
                        that.$emit("bar-line-click", params, this.option);
                    });
                };
                var chartDom = this.$refs[this.id];
                if (chartDom) {
                    this.myChart = this.$echarts.init(chartDom);
                    this.myChart.setOption(val);
                    extension(this.myChart);
                    this.addResizeListenet();
                }
            }
        },
        addResizeListenet() {
            this.removeResizeListener();
            window.addEventListener("resize", this.listener);
        },
        removeResizeListener() {
            if (this.listener) {
                window.removeEventListener("resize", this.listener);
            }
        },
        listener() {
            this.myChart && this.myChart.resize();
        },
        exportImg() {
            if (!this.myChart) return;
            const src = this.myChart.getDataURL({
                pixelRatio: 2,
                backgroundColor: "#fff"
            });
            const a = document.createElement("a");
            a.href = src;
            a.download = "chart-img";
            a.click();
        }
    }
};
</script>
<style></style>
