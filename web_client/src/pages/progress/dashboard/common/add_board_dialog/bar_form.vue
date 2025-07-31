<template>
    <div class="dialog-chart-form">
        <el-form
            class="basic-ui"
            ref="form"
            size="small"
            :model="form"
            :rules="registerRules"
        >
            <!-- 全部显示 -->
            <el-form-item>
                <div slot="label" class="label">图表类型：</div>
                <el-input
                    class="basic-ui switch"
                    v-model="chartType"
                    type="text"
                    @click.native="returnStepsOne"
                >
                    <i slot="suffix" class="switch-img"></i>
                </el-input>
            </el-form-item>
            <!-- 全部显示 -->
            <el-form-item>
                <div slot="label" class="label">视图标题：</div>
                <el-input
                    class="basic-ui"
                    v-model="title"
                    type="text"
                    placeholder="请输入"
                    @input="titleChange"
                ></el-input>
                <div
                    v-if="isShowError"
                    style="
                        color: #f56c6c;
                        font-size: 12px;
                        line-height: 1;
                        padding-top: 4px;
                        position: absolute;
                        top: 100%;
                        left: 0;
                    "
                >
                    请输入视图标题
                </div>
            </el-form-item>
            <!-- x轴 柱形图、折线图显示 -->
            <el-form-item v-if="isShowXY" prop="x_axis">
                <div slot="label">X轴：</div>
                <el-select
                    class="basic-ui"
                    v-model="form.x_axis"
                    placeholder="请选择"
                    @change="(value) => selectChange(value, 'x_axis')"
                >
                    <el-option
                        v-for="item in xAxisOptions"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    >
                        <span style="display: flex; align-items: center">
                            <img
                                :src="
                                    item.mode === 'blank_com'
                                        ? typeIconMapping[
                                              `${item.mode}_${item.value}`
                                          ]
                                        : typeIconMapping[item.mode]
                                "
                                alt=""
                                width="20px"
                                height="20px"
                            />
                            <span style="margin-left: 8px" v-overflow>{{
                                item.label
                            }}</span>
                        </span>
                    </el-option>
                </el-select>
            </el-form-item>
            <!-- x轴年月日 如果选中的日期时间类，显示下面 -->
            <el-form-item
                v-if="
                    isShowXY &&
                    (curFormXAxis.mode === 'time_com' ||
                        curFormXAxis.mode === 'date_com')
                "
                class="date-select"
            >
                <el-select
                    class="basic-ui"
                    v-model="form.axis_date_mode"
                    placeholder="请选择"
                >
                    <el-option
                        v-for="item in dashboardConfig['date_mode']"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    >
                    </el-option>
                </el-select>
            </el-form-item>
            <!-- y轴 柱形图、折线图显示 如果x轴选中的日期时间类，显示下面 -->
            <el-form-item v-if="isShowXY" prop="y_axis">
                <div slot="label">Y轴：</div>
                <el-select
                    class="basic-ui"
                    v-model="form.y_axis"
                    placeholder="请选择"
                >
                    <el-option
                        v-for="item in yAxisOptions"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    >
                    </el-option>
                </el-select>
            </el-form-item>
            <!--  分组柱形图、分组折线图、堆叠柱形图、饼图 -->
            <el-form-item v-if="isShowClassify" prop="group_axis">
                <div slot="label">分组：</div>
                <div>
                    <el-select
                        class="basic-ui"
                        v-model="form.group_axis"
                        placeholder="请选择"
                        @change="(value) => selectChange(value, 'group_axis')"
                    >
                        <el-option
                            v-for="item in classifyOptions"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        >
                            <span style="display: flex; align-items: center">
                                <img
                                    :src="
                                        item.mode === 'blank_com'
                                            ? typeIconMapping[
                                                  `${item.mode}_${item.value}`
                                              ]
                                            : typeIconMapping[item.mode]
                                    "
                                    alt=""
                                    width="20px"
                                    height="20px"
                                />
                                <span style="margin-left: 8px" v-overflow>{{
                                    item.label
                                }}</span>
                            </span>
                        </el-option>
                    </el-select>
                </div>
            </el-form-item>
            <!-- 分组柱形图、分组折线图、堆叠柱形图、饼图， 如果分组选中的日期时间类，显示下面 -->
            <el-form-item
                v-if="
                    isShowClassify &&
                    (curFormClassify.mode === 'time_com' ||
                        curFormClassify.mode === 'date_com')
                "
                class="date-select"
            >
                <el-select
                    class="basic-ui"
                    v-model="form.group_date_mode"
                    placeholder="请选择"
                >
                    <el-option
                        v-for="item in dashboardConfig['date_mode']"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    >
                    </el-option>
                </el-select>
            </el-form-item>
            <!-- 全部显示 -->
            <el-form-item class="yes-or-no">
                <div slot="label">包含空单元格：</div>
                <el-switch
                    v-model="form.show_empty"
                    :active-color="PRIMARY_COLOR"
                    inactive-color="#e6e9f0"
                >
                </el-switch>
            </el-form-item>
            <!-- 全部显示 -->
            <el-form-item prop="order">
                <div slot="label">数据排序：</div>
                <el-radio-group class="basic-ui order" v-model="form.order">
                    <el-radio label="not">不排序</el-radio>
                    <el-radio label="asc">升序</el-radio>
                    <el-radio label="desc">降序</el-radio>
                </el-radio-group>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
import _ from "lodash";
import { PRIMARY_COLOR_LIST } from "@/assets/tool/color_list";
import dataHandle from "@/pages/progress/progress_setting/filed/data_handle";
export default {
    props: {
        actCurChart: {
            type: Object,
            default: () => {}
        },
        actCurParentChart: {
            type: Object,
            default: () => {}
        },
        curBoardItem: {
            type: Object,
            default: () => {}
        }
    },
    components: {},
    data() {
        return {
            // 类型iconyings
            isShowError: false,
            typeIconMapping: _.cloneDeep(dataHandle.typeIconMapping),
            PRIMARY_COLOR: PRIMARY_COLOR_LIST.PRIMARY_COLOR,
            chartType: "",
            title: "",
            form: {
                // title: "",
                x_axis: "",
                axis_date_mode: "year",
                y_axis: "count",
                show_empty: false,
                order: "not",
                group_axis: "",
                group_date_mode: "year"
            },
            curFormXAxis: {},
            curFormClassify: {},
            registerRules: {
                chartType: [
                    {
                        required: true,
                        message: "请选择图表类型",
                        trigger: "change"
                    }
                ],
                title: [
                    {
                        required: false,
                        message: "请输入",
                        trigger: ["blur", "change"]
                    }
                ],
                x_axis: [
                    {
                        required: true,
                        message: "请选择",
                        trigger: "change"
                    }
                ],
                y_axis: [
                    {
                        required: true,
                        message: "请选择",
                        trigger: "change"
                    }
                ],
                group_axis: [
                    {
                        required: true,
                        message: "请选择",
                        trigger: "change"
                    }
                ],
                order: [
                    {
                        required: true,
                        message: "请选择",
                        trigger: "change"
                    }
                ],
                show_empty: [
                    {
                        required: true,
                        message: "请选择",
                        trigger: "change"
                    }
                ]
            },
            xAxisOptions: [],
            classifyOptions: []
        };
    },
    computed: {
        isShowClassify() {
            return (
                this.actCurChart.name === "分组柱形图" ||
                this.actCurChart.name === "分组折线图" ||
                this.actCurChart.name === "堆叠柱形图" ||
                this.actCurParentChart.title === "饼图"
            );
        },
        isShowXY() {
            return (
                this.actCurParentChart.title === "柱形图" ||
                this.actCurParentChart.title === "折线图"
            );
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
        },
        dashboardConfig() {
            return this.$store.state.dashboard.config;
        }
    },
    actCurChart() {},
    watch: {
        dashboardConfig: {
            handler(obj) {
                if (obj && Object.keys(obj).length) {
                    this.classifyOptions = _.cloneDeep(obj.field_list);
                    this.xAxisOptions = _.cloneDeep(obj.field_list);
                    this.yAxisOptions = _.cloneDeep(obj.statistics_func);
                } else {
                    this.classifyOptions = [];
                    this.xAxisOptions = [];
                    this.yAxisOptions = [];
                }
            },
            deep: true,
            immediate: true
        },
        form: {
            handler(obj) {
                let tmpObj = _.cloneDeep(obj);
                // this.setOptionsDiasbled(tmpObj);
                this.validForm();
            },
            deep: true
        },
        chartType(newVal, oldVal) {
            if (newVal && oldVal) {
                if (newVal === "基础饼图" || newVal === "环形饼图") {
                    this.classifyOptions = _.cloneDeep(
                        this.dashboardConfig.field_list
                    );
                } else if (newVal === "基础折线图" || newVal === "基础柱形图") {
                    this.xAxisOptions = _.cloneDeep(
                        this.dashboardConfig.field_list
                    );
                } else {
                    // this.setOptionsDiasbled(this.form);
                }
                this.validForm();
                this.$nextTick(() => {
                    this.$refs.form.clearValidate();
                });
            }
        }
    },
    methods: {
        // 标题编辑时
        titleChange(title) {
            if (title) {
                this.isShowError = false;
            } else {
                this.isShowError = true;
            }
        },
        setOptionsDiasbled(obj) {
            if (obj.x_axis) {
                this.classifyOptions.forEach((item) => {
                    this.$set(item, "disabled", item.value === obj.x_axis);
                });
            }
            if (obj.group_axis) {
                this.xAxisOptions.forEach((item) => {
                    this.$set(item, "disabled", item.value === obj.group_axis);
                });
            }
        },
        initData() {
            this.PRIMARY_COLOR = PRIMARY_COLOR_LIST.PRIMARY_COLOR;
            // this.chartType = "";
            this.title = "";
            this.form = {
                // title: "",
                x_axis: "",
                axis_date_mode: "year",
                y_axis: "count",
                show_empty: false,
                order: "not",
                group_axis: "",
                group_date_mode: "year"
            };
            this.curFormXAxis = {};
            this.curFormClassify = {};
            this.registerRules = {
                chartType: [
                    {
                        required: true,
                        message: "请选择活动资源",
                        trigger: "change"
                    }
                ],
                title: [
                    {
                        required: true,
                        message: "请输入",
                        trigger: ["blur", "change"]
                    }
                ],
                x_axis: [
                    {
                        required: true,
                        message: "请选择",
                        trigger: "change"
                    }
                ],
                y_axis: [
                    {
                        required: true,
                        message: "请选择",
                        trigger: "change"
                    }
                ],
                group_axis: [
                    {
                        required: true,
                        message: "请选择",
                        trigger: "change"
                    }
                ],
                order: [
                    {
                        required: true,
                        message: "请选择",
                        trigger: "change"
                    }
                ]
            };
        },
        selectChange(value, type) {
            if (type === "x_axis") {
                this.form.axis_date_mode = "year";
                this.xAxisOptions.forEach((item) => {
                    if (item.value === value) {
                        this.curFormXAxis = item;
                    }
                });
            }
            if (type === "group_axis") {
                this.form.group_date_mode = "year";
                this.classifyOptions.forEach((item) => {
                    if (item.value === value) {
                        this.curFormClassify = item;
                    }
                });
            }
        },
        returnStepsOne() {
            this.$emit("returnStepsOne");
        },
        validForm() {
            this.isShowError = false;
            this.$refs.form.validate((valid) => {
                let data = this.handleFormData({
                    ...this.form,
                    title: this.title,
                    chartType: this.chartType,
                    xMode: this.curFormXAxis.mode
                });
                // this.$emit("form-change", data);
                this.$emit("form-change", data, valid);
                // else {
                //     // 判断是否只有title为空，如果是，可以画图
                //     let arr = [];
                //     for (let key in this.form) {
                //         if (!this.form[key]) {
                //             arr.push(key);
                //         }
                //     }
                //     if (arr.index === 1 && arr[0] === "title") {
                //         this.$emit("form-change", data);
                //     } else {
                //         this.$emit("form-change", {});
                //     }
                // }
            });
        },
        // 仅在编辑完图表后，确认新增时调用
        validate(callback) {
            this.$refs.form.validate((valid) => {
                let data = this.handleFormData({
                    ...this.form,
                    title: this.title,
                    chartType: this.chartType,
                    xMode: this.curFormXAxis.mode
                });
                if (!this.title) {
                    this.isShowError = true;
                } else {
                    callback(valid, data);
                }
            });
        },
        handleFormData(data) {
            let tmpData = _.cloneDeep(data);
            if (this.actCurParentChart.title === "饼图") {
                delete tmpData.x_axis;
                delete tmpData.y_axis;
            } else if (
                this.actCurChart.name === "基础柱形图" ||
                this.actCurChart.name === "基础折线图"
            ) {
                delete tmpData.group_axis;
            }
            if (
                this.curFormXAxis.mode !== "time_com" &&
                this.curFormXAxis.mode !== "date_com"
            ) {
                delete tmpData.axis_date_mode;
                delete tmpData.curFormXAxis;
            }
            if (
                this.curFormClassify.mode !== "time_com" &&
                this.curFormClassify.mode !== "date_com"
            ) {
                delete tmpData.group_date_mode;
                delete tmpData.curFormClassify;
            }
            return tmpData;
        },
        // 编辑时回填form表单
        setFormData() {
            if (this.isShowClassify) {
                this.form.group_axis = this.curBoardItem.group_axis;
                this.selectChange(this.form.group_axis, "group_axis");
                if (
                    this.curFormClassify.mode === "time_com" ||
                    this.curFormClassify.mode === "date_com"
                ) {
                    this.form.group_date_mode =
                        this.curBoardItem.group_date_mode;
                }
            }
            if (this.isShowXY) {
                this.form.x_axis = this.curBoardItem.x_axis;
                this.selectChange(this.form.x_axis, "x_axis");
                if (
                    this.curFormXAxis.mode === "time_com" ||
                    this.curFormXAxis.mode === "date_com"
                ) {
                    this.form.axis_date_mode = this.curBoardItem.axis_date_mode;
                }
            }

            this.form.show_empty =
                this.curBoardItem.show_empty === "no" ? false : true;
            this.form.order = !this.curBoardItem.order
                ? "not"
                : this.curBoardItem.order;
            this.form.y_axis = "count";
        }
    }
};
</script>

<style lang="scss" scoped>
.dialog-chart-form ::v-deep {
    .order.el-radio-group .el-radio {
        display: block;
        margin-right: 0;
        line-height: 32px;
    }
    .switch {
        .switch-img {
            position: relative;
            top: 4px;
            width: 20px;
            height: 20px;
            background-image: url(@/assets/image/progress_dashboard/switch_chart.png);
            background-size: 100% 100%;
            cursor: pointer;
        }
        .el-input__inner {
            cursor: pointer;
        }
    }
    .yes-or-no .el-form-item__content {
        text-align: right;
    }
    .date-select {
        margin-top: -10px;
    }
    .el-select {
        display: block;
    }
    .el-select > .el-input {
        display: inline-block;
    }
    .el-form-item__label {
        position: relative;
        &::before {
            position: absolute;
            left: 0;
        }
        div {
            margin-left: 10px;
        }
        div.label {
            margin-left: 0;
            &::before {
                content: "*";
                color: #f56c6c;
                margin-right: 4px;
            }
        }
    }
}
</style>
