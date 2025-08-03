<template>
    <el-dialog
        :title="title"
        class="basic-ui add-board-dialog"
        :class="steps === 2 ? 'steps2' : ''"
        :visible.sync="dialogVisible"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        @close="stepsOnecancel"
    >
        <div class="add-board-body">
            <!-- v-show可以获取到切换图表之前输入的数据 -->
            <div v-show="steps === 1" class="steps-one">
                <div v-for="(item, index) in chartsTypeArr" :key="index">
                    <div class="sub-title">{{ item.title }}</div>
                    <div class="sub-content">
                        <div
                            :span="12"
                            v-for="(sub, ind) in item.typeArr"
                            :key="ind"
                            :class="['item', `${sub.active ? 'active' : ''}`]"
                            @click="itemSelected(item, sub)"
                        >
                            <div class="sub-name">{{ sub.name }}</div>
                            <img :src="sub.img" width="258px" height="132px" />
                        </div>
                    </div>
                </div>
            </div>
            <div v-show="steps === 2" class="steps-two">
                <div :class="['left', isShowChart ? '' : 'no-data']">
                    <bar-chart
                        v-if="isShowChart"
                        ref="barForm"
                        :option="option"
                        :id="chartId"
                        style="height: 100%"
                    ></bar-chart>
                    <div v-else>
                        <no-data
                            :isTextShow="true"
                            text="暂无内容"
                            mainText="请选择右侧数据定制图表"
                            :loading="loading"
                            :imgShow="imgShow"
                        >
                        </no-data>
                        <!-- 暂无内容
                        <div>请选择右侧数据定制图表</div> -->
                    </div>
                </div>
                <div class="right">
                    <bar-form
                        v-if="dialogVisible"
                        ref="chartForm"
                        :curBoardItem="curBoardItem"
                        :actCurChart="actCurChart"
                        :actCurParentChart="actCurParentChart"
                        @form-change="formChange"
                        @returnStepsOne="returnStepsOne"
                    ></bar-form>
                </div>
            </div>
        </div>
        <span slot="footer" class="dialog-footer">
            <div v-show="steps === 1">
                <el-button v-if="!isReturn" size="small" @click="stepsOnecancel"
                    >取消</el-button
                >
                <el-button v-else size="small" @click="returnSteps2"
                    >返回</el-button
                >
                <el-button
                    size="small"
                    type="primary"
                    :disabled="!isShowNextBtn"
                    @click="nextSteps"
                    >下一步</el-button
                >
            </div>
            <div v-show="steps === 2">
                <el-button size="small" @click="stepsTwoCancel">取消</el-button>
                <el-button size="small" type="primary" @click="stepTwoConfirm"
                    >提交</el-button
                >
            </div>
        </span>
    </el-dialog>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress_dashboard";
import service from "../../service";
import dataHandle from "../../data_handle";
import BarForm from "./bar_form.vue";
import BarChart from "./bar_chart.vue";
import NoData from "@/pages/common/no_data.vue";

export default {
    components: {
        BarForm,
        BarChart,
        NoData
    },
    props: {
        curTabItems: {
            type: Object,
            default: () => {}
        },
        filterData: {
            type: Object,
            default: () => {
                return {
                    filterParams: "",
                    viewFilterStandby: null,
                    filterRelation: "filter_and"
                };
            }
        }
    },
    data() {
        return {
            steps: 1, // 1即 选择图类型页面，2即选择后编辑页面
            dialogVisible: false,
            currentType: "number",
            currentTypeName: "数值",
            title: "新增图表",
            chartsTypeArr: _.cloneDeep(dataHandle.chartsTypeArr),
            curChart: {
                name: "基础柱形图",
                img: require(`@/assets/image/progress_dashboard/bar_1.png`),
                active: true
            },
            curParentChart: {
                title: "柱形图"
            },
            actCurChart: {},
            actCurParentChart: {
                title: "柱形图"
            },
            isReturn: false,
            formMapping: {
                柱形图: "barForm",
                折线图: "lineForm",
                饼图: "pieForm"
            },
            option: {},
            isShowChart: false,
            chartMapping: {
                柱形图: dataHandle.options1,
                折线图: dataHandle.options2,
                饼图: dataHandle.options3
            },
            curBoardItem: {},
            isEdit: false,
            loading: false,
            imgShow: true,
            chartId: "new-chart"
        };
    },
    computed: {
        isShowNextBtn() {
            return this.curChart && Object.keys(this.curChart).length;
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        },
        progressTree() {
            return this.$store.state.progressTree;
        },
        userInfo() {
            return this.$store.getters.userInfo;
        }
    },
    watch: {},
    methods: {
        formChange(form, valid) {
            if (!valid) {
                this.isShowChart = false;
                this.option = {};
                return;
            }
            if (form && Object.keys(form).length) {
                let filterParams = this.filterData.filterParams;
                let params = {
                    ws_id: this.curSpace.id,
                    tmpl_id: Number(this.curProgress),
                    view_id: this.curTabItems.id,
                    userid: this.userInfo.id,
                    filter:
                        filterParams && filterParams.length
                            ? JSON.stringify(filterParams)
                            : "",
                    lor: this.filterData.filterRelation,
                    ...form
                };
                params.mode = dataHandle.chartTypeMpping[form.chartType] || "";
                params.show_empty = params.show_empty ? "yes" : "no";
                params.order = params.order === "not" ? "" : params.order;
                delete params.chartType;
                //调接口获取预览图
                api.getPrviewChart(params).then((res) => {
                    if (
                        res &&
                        res.resultCode === 200 &&
                        res.data &&
                        res.data.board_statistics &&
                        res.data.board_statistics.length
                    ) {
                        let obj = {
                            ...params,
                            data: res.data.board_statistics
                        };
                        this.isShowChart = true;
                        if (params.mode.indexOf("pie") > -1) {
                            // 饼图
                            this.option = service.getPieOPtions(obj);
                        } else if (
                            params.mode === "base_histogram" ||
                            params.mode === "base_line"
                        ) {
                            // 其他
                            this.option = service.getOtherChartOptions(obj);
                            this.option.x_cn = res.data.x_cn;
                        } else {
                            // 分组和堆叠
                            this.option = service.getGroupBarLineOptions(obj);
                            this.option.x_cn = res.data.x_cn;
                            this.option.group_cn = res.data.group_cn;
                        }
                    } else {
                        this.isShowChart = false;
                        this.option = {};
                    }
                });
            } else {
                // 2.否则置空
                this.isShowChart = false;
                this.option = {};
            }
        },
        changeType(item) {
            if (item.type === "welcome") {
                return;
            }
            this.currentType = item.type;
            this.currentTypeName = item.name;
        },
        itemSelected(curParentChart, curChart) {
            this.curParentChart = curParentChart;
            this.curChart = curChart;
            this.chartsTypeArr.forEach((item) => {
                if (curParentChart.title === item.title) {
                    item.typeArr.forEach((sub) => {
                        if (curChart.name === sub.name) {
                            this.$set(sub, "active", true);
                        } else {
                            this.$set(sub, "active", false);
                        }
                    });
                } else {
                    item.typeArr.forEach((sub) => {
                        this.$set(sub, "active", false);
                    });
                }
            });
        },
        openDialog(item) {
            this.dialogVisible = true;
            this.isShowChart = false;
            if (item && Object.keys(item).length) {
                this.curBoardItem = _.cloneDeep(item);
                this.isEdit = true;
                this.editGridItem(item);
                this.chartId = "edit-chart";
            } else {
                this.isEdit = false;
                this.curBoardItem = {};
                this.steps = 1;
                this.chartId = "new-chart";
            }
        },
        editGridItem(item) {
            this.itemSelected(
                {
                    title: dataHandle.chartTypeParent[item.mode][1]
                },
                {
                    name: dataHandle.chartTypeParent[item.mode][0],
                    active: true
                }
            );
            this.actCurChart = {
                name: dataHandle.chartTypeParent[item.mode][0],
                active: true
            };
            this.actCurParentChart = {
                title: dataHandle.chartTypeParent[item.mode][1]
            };
            this.steps = 2;
            this.title = "编辑图表";
            this.setFormChartType();
        },
        fetchcurBoardItemData() {
            api.getModifyDashboardItem();
        },
        //  steps1界面的取消
        stepsOnecancel() {
            this.curChart = {
                name: "基础柱形图",
                img: require(`@/assets/image/progress_dashboard/bar_1.png`),
                active: true
            };
            this.actCurParentChart = {
                title: "柱形图"
            };
            this.$refs["chartForm"].initData();
            this.steps = 1;
            this.isShowChart = false;
            this.option = {};
            this.chartsTypeArr = _.cloneDeep(dataHandle.chartsTypeArr);
            this.dialogVisible = false;
        },
        //  steps1界面的返回， 如果是返回按钮，说明是在第二步切换回steps1的
        returnSteps2() {
            this.steps = 2;
            this.isReturn = false;
            this.itemSelected(this.actCurParentChart, this.actCurChart);
            this.setFormChartType();
        },
        setFormChartType() {
            this.$nextTick(() => {
                if (
                    this.actCurParentChart &&
                    Object.keys(this.actCurParentChart).length &&
                    this.$refs["chartForm"]
                ) {
                    this.$refs["chartForm"].chartType = this.actCurChart.name;
                }
                if (this.isEdit) {
                    this.$refs["chartForm"].title = this.curBoardItem.title;
                    this.$refs["chartForm"].setFormData();
                }
                this.$refs["chartForm"].$refs.form.clearValidate();
            });
        },
        // steps1界面的下一步
        nextSteps() {
            this.actCurChart = _.cloneDeep(this.curChart);
            this.actCurParentChart = _.cloneDeep(this.curParentChart);
            this.steps = 2;
            this.title = "编辑图表";
            this.setFormChartType();
        },
        // form表单
        returnStepsOne() {
            this.curChart = _.cloneDeep(this.actCurChart);
            this.curParentChart = _.cloneDeep(this.actCurParentChart);
            this.isReturn = true;
            this.steps = 1;
        },
        //  steps2界面的取消
        stepsTwoCancel() {
            this.curChart = {
                name: "基础柱形图",
                img: require(`@/assets/image/progress_dashboard/bar_1.png`),
                active: true
            };
            this.actCurParentChart = {
                title: "柱形图"
            };
            this.$refs["chartForm"].initData();
            this.steps = 1;
            this.isShowChart = false;
            this.option = {};
            this.chartsTypeArr = _.cloneDeep(dataHandle.chartsTypeArr);
            this.dialogVisible = false;
        },
        // steps2界面的确定，同时给后端传递form表单所有数据请求接口关闭弹窗
        stepTwoConfirm() {
            this.$refs["chartForm"].registerRules.title[0].required = true;
            this.$refs["chartForm"].validate((valid, data) => {
                if (valid) {
                    this.$emit(
                        "on-confirm",
                        { ...this.curBoardItem, ...data },
                        this.isEdit
                    );
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.add-board-body {
    display: flex;
    align-items: stretch;
    height: 100%;
    width: calc(100% + 31px);
    margin-right: -31px;
    .steps-one {
        margin-top: 20px;
        box-sizing: border-box;
        .sub-title {
            font-size: 16px;
            margin-bottom: 12px;
        }
        .sub-content {
            display: flex;
            .item {
                box-sizing: border-box;
                width: 300px;
                height: 200px;
                padding: 16px 20px;
                border: 1px solid #e6e9f0;
                border-radius: 8px;
                margin: 0 16px 20px 0;
                .sub-name {
                    line-height: 20px;
                    height: 20px;
                    margin-bottom: 14px;
                }
                img {
                    height: calc(100% - 34px);
                    background-size: 100% 100%;
                }
                &:nth-child(3n) {
                    margin-right: 0;
                }
                &:hover {
                    border: 1px solid #1890ff;
                }
                &.active {
                    background-color: #f4faff;
                    border: 1px solid #1890ff;
                }
            }
        }
    }
    .steps-two {
        display: flex;
        align-items: stretch;
        height: 100%;
        width: 100%;
        .left {
            width: calc(100% - 380px);
            // height: 100%;
            padding: 16px 0;
            &.no-data {
                text-align: center;
                margin: auto 0;
            }
            ::v-deep .chart-box {
                position: relative;
                top: 50%;
                left: 50%;
                transform: translate(-50%, -50%);
            }
        }
        .right {
            width: 400px;
            height: 100%;
            overflow-y: auto;
            box-sizing: border-box;
            border-left: 1px solid #e6e9f0;
            margin-left: 20px;
            padding: 16px 32px 16px 16px;
            background-color: #f5f6f8;
        }
    }
}
</style>
<style lang="scss">
.basic-ui.add-board-dialog .el-dialog {
    width: auto;
    min-width: 200px;
    height: calc(100% - 120px);
    max-height: 800px;
    .el-dialog__header {
        padding: 24px 32px 16px;
    }
    .el-dialog__body {
        height: calc(100% - 124px); // 减去头部和尾部
        max-height: calc(100% - 124px); // 减去头部和尾部
        border-top: 1px solid #e6e9f0;
        border-bottom: 1px solid #e6e9f0;
    }
    .el-dialog__footer {
        background-color: #fff;
    }
}
.steps2.basic-ui.add-board-dialog .el-dialog {
    width: calc(100% - 164px);
    height: calc(100% - 120px);
    max-height: 800px;
    max-width: 1100px;

    .el-dialog__body {
        height: calc(100% - 124px); // 减去头部和尾部
        max-height: calc(100% - 124px); // 减去头部和尾部
        overflow: hidden;
    }
}
</style>
