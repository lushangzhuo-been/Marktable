<template>
    <el-dialog
        :title="title"
        class="basic-ui full-screen-dialog"
        :visible.sync="dialogVisible"
        :close-on-press-escape="false"
        append-to-body
        :before-close="cancel"
    >
        <div class="top-filter">
            <b v-if="!isReturnBtnShow"></b>
            <el-button
                v-if="isReturnBtnShow"
                class="basic-ui"
                size="mini"
                @click="returnBtn"
                >返回</el-button
            >
            <span class="right">
                <b
                    :class="[
                        'switch-icon',
                        isShowTable ? 'table-show' : 'table-hide'
                    ]"
                    @click="showTable"
                ></b>
            </span>
        </div>
        <div class="content">
            <div class="charts" :style="{ height: echartsHeight }">
                <basic-charts
                    v-show="isShowChart"
                    ref="chart"
                    id="chart"
                    :option="option"
                    @bar-line-click="chartClick"
                ></basic-charts>
                <no-data
                    v-show="!isShowChart"
                    :isTextShow="false"
                    :loading="loading"
                    :imgShow="imgShow"
                ></no-data>
            </div>
            <div class="table" v-show="isShowTable">
                <div class="line"></div>
                <div class="operation">
                    <el-dropdown
                        split-button
                        type="primary"
                        size="small"
                        class="btn-group"
                        @click="addProgress"
                        @command="rightBtn"
                    >
                        <b class="add"></b>新增
                        <el-dropdown-menu slot="dropdown" class="operation-btn">
                            <el-dropdown-item command="export">
                                <b class="export-box"></b>导出</el-dropdown-item
                            >
                        </el-dropdown-menu>
                    </el-dropdown>
                </div>
                <p-table
                    ref="dialogProgressTable"
                    :data="sheet.tableData"
                    :col="sheet.tableCol"
                    @open-html-col="openHtmlCol"
                    @open-detail="openTableDetail"
                    @multiple-selection="multipleSelection"
                    @refresh="refresh"
                    @delete-row="deleteRow"
                ></p-table>
                <pagination
                    v-show="sheet.total > 10"
                    :total="sheet.total"
                    :size="sheet.page_size"
                    @pageSizeChange="pageSizeChange"
                    @curPageChange="currentChange"
                ></pagination>
            </div>
        </div>
        <detail-drawer ref="DetailDrawer" @refresh="refresh"></detail-drawer>
        <add-progress
            ref="AddProgress"
            :formLabel="createCol"
            :formData="addProgressFormData"
            @refresh="refresh"
        ></add-progress>
        <!-- html富文本编辑器 -->
        <html-col-dialog
            ref="htmlColDialog"
            @edit-form-item="editFormItem"
        ></html-col-dialog>
        <!-- 删除表格行提醒 -->
        <delete-table-row-dialog
            ref="deleteTableRowDialog"
            @cancel="cancelDeleteTableRow"
            @on-confirm="comfirmDeleteTableRow"
        ></delete-table-row-dialog>
    </el-dialog>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import api_chart from "@/common/api/module/progress_dashboard";
import dataHandle from "../../data_handle.js";
import service from "../../service.js";
import BasicCharts from "@/pages/common/echarts/basic_chart.vue";
import PTable from "@/components/table/p_table/progress_table";
import Pagination from "@/pages/common/pagination.vue";
import DetailDrawer from "@/pages/progress/detail/index";
import AddProgress from "@/pages/progress/add_progress/index";
import HtmlColDialog from "@/pages/progress/components/html_col_dialog.vue";
import NoData from "@/pages/common/no_data.vue";
import DeleteTableRowDialog from "@/pages/progress/components/delete_table_row_dialog.vue";
export default {
    components: {
        BasicCharts,
        PTable,
        Pagination,
        DetailDrawer,
        AddProgress,
        HtmlColDialog,
        NoData,
        DeleteTableRowDialog
    },
    props: {
        filterData: {
            type: Object,
            default: () => {
                return {
                    filterParams: "",
                    viewFilterStandby: null,
                    filterRelation: "filter_and"
                };
            }
        },
        curViewInfo: {
            type: Object,
            default: () => {}
        }
    },
    computed: {
        // 当前空间
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        // 当前流程
        curProgress() {
            return this.$route.params.id;
        },
        // 当前的整个流程树
        progressTree() {
            return this.$store.state.progressTree;
        },
        // 分组或堆叠图
        isGroupChart() {
            return (
                this.curItem.mode === "group_histogram" ||
                this.curItem.mode === "stack_histogram" ||
                this.curItem.mode === "group_line"
            );
        },
        // 是否饼图或者环形图
        isPie() {
            return (
                this.curItem.mode === "base_pie" ||
                this.curItem.mode === "donut_pie"
            );
        },
        // 基础折线、柱形图
        isBasic() {
            return (
                this.curItem.mode === "base_histogram" ||
                this.curItem.mode === "base_line"
            );
        },
        isDown() {
            return this.filter_down && Object.keys(this.filter_down).length
                ? true
                : false;
        }
    },
    data() {
        return {
            isShowChart: false,
            loading: true,
            imgShow: false,
            tableMultipleSelection: [], //表格多选用于删除
            createCol: [],
            addProgressFormData: {},
            dialogVisible: false,
            title: "",
            option: {},
            isShowTable: true,
            echartsHeight: "300px",
            curItem: {},
            sheet: {
                tableCol: [],
                tableData: [],
                total: 0,
                page: 1,
                page_size: 10
            },
            isReturnBtnShow: false,
            filter_down: {},
            color: "#1890ff"
        };
    },

    watch: {
        isShowTable() {
            this.echartsHeight = this.isShowTable ? "300px" : "100%";
            this.$nextTick(() => {
                this.$refs.chart.listener();
            });
        }
    },
    methods: {
        // 导出表格
        rightBtn(command) {
            if (command === "export") {
                let filterParams = this.filterData.filterParams;
                let params = {
                    ws_id: this.curSpace.id,
                    tmpl_id: this.curProgress,
                    filter:
                        filterParams && filterParams.length
                            ? JSON.stringify(filterParams)
                            : "",
                    lor: this.filterData.filterRelation,
                    page_num: this.sheet.page,
                    page_size: this.sheet.page_size,
                    filter_down:
                        this.filter_down && Object.keys(this.filter_down).length
                            ? JSON.stringify(this.filter_down)
                            : "",
                    export: "yes"
                };
                api.exportProgress(params);
            }
        },
        // 操作列删除行
        deleteRow(row) {
            // 打开删除表格行的提醒弹窗
            this.$refs.deleteTableRowDialog.openDialog(row);
            // this.$emit("delete-row", row);
        },
        // 取消删除表格行
        cancelDeleteTableRow() {
            // this.tableMultipleSelection = [];
        },
        // 确认删除表格行调接口
        comfirmDeleteTableRow(row) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                ids: row._id
                // ids: idArr.join(",")
            };
            api.deleteProgressItems(params).then((res) => {
                if (this.$route.name !== "progress") return;
                if (res && res.resultCode === 200) {
                    this.$refs.deleteTableRowDialog.cancel();
                    this.$message({
                        showClose: true,
                        message: "删除成功",
                        type: "success"
                    });
                    document.body.click();
                    this.getChartData();
                    this.getTableData();
                }
            });
        },
        returnBtn() {
            this.isReturnBtnShow = false;
            this.filter_down = {};
            this.getChartData();
            this.getTableData();
        },
        chartClick(chart, option) {
            this.color = chart.color;
            if (this.isDown) return; // 本身已是下钻后，就不再下钻
            let xValue = {};
            if (this.isGroupChart) {
                xValue = option.xAxis[chart.dataIndex].data[chart.dataIndex];
            } else if (this.isBasic) {
                xValue = option.xAxis.data[chart.dataIndex];
            }
            this.isReturnBtnShow = true;

            if (!this.isPie) {
                this.filter_down = {
                    x_axis: this.curItem.x_axis,
                    axis_date_mode: this.curItem.axis_date_mode,
                    x_value: xValue.id || xValue.value
                };
            } else {
                this.filter_down = {};
            }
            if (this.isGroupChart || this.isPie) {
                this.filter_down.group_axis = this.curItem.group_axis;
                this.filter_down.group_date_mode = this.curItem.group_date_mode;
                this.filter_down.group_value =
                    chart.data.user_id ||
                    chart.data.status_id ||
                    chart.data.document_id ||
                    chart.data.name;
            }
            this.getChartData();
            this.getTableData();
        },
        openDialog(item, out = false, chart, option) {
            this.curItem = _.cloneDeep(item);
            this.title = item.title;
            this.dialogVisible = true;
            if (out) {
                // 外层下钻
                this.$nextTick(() => {
                    this.getTableCol().then(() => {
                        this.chartClick(chart, option);
                    });
                });
            } else {
                this.$nextTick(() => {
                    this.getChartData();
                    this.getTableCol().then(() => {
                        this.getTableData();
                    });
                });
            }
        },
        initData() {
            this.isReturnBtnShow = false;
            this.tableMultipleSelection = []; //表格多选用于删除
            this.createCol = [];
            this.addProgressFormData = {};
            this.dialogVisible = false;
            this.title = "";
            this.option = {};
            this.filter_down = {};
            this.isShowTable = true;
            this.echartsHeight = "350px";
            this.sheet = {
                tableCol: [],
                tableData: [],
                total: 0,
                page: 1,
                page_size: 10
            };
        },
        openTableDetail(scope, type) {
            this.formData = scope.row;
            this.$refs["DetailDrawer"].openDrawer(type, scope.row);
        },
        refresh() {
            // 刷新表格
            this.getTableData();
            this.getChartData();
        },
        cancel() {
            this.dialogVisible = false;
            this.initData();
            this.$emit("dashboard-refresh", this.curItem);
        },
        exportTable() {},
        // newRow() {},
        addProgress() {
            // 掉新增config接口，打开dialog
            this.getAddProgressConfig();
            this.$refs.AddProgress.openDialog();
        },
        getAddProgressConfig() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                module: "create"
            };
            api.getProgressColConfig(params).then((res) => {
                if (this.$route.name !== "progress") return;
                if (
                    res &&
                    res.resultCode === 200 &&
                    Object.keys(res.data).length
                ) {
                    this.createCol = res.data;
                } else {
                    this.createCol = [];
                }
            });
        },
        // 右上角显示隐藏表格icon
        showTable() {
            this.isShowTable = !this.isShowTable;
        },
        getChartData() {
            let filterParams = this.filterData.filterParams;
            let params = {
                board_id: this.curItem.i,
                filter:
                    filterParams && filterParams.length
                        ? JSON.stringify(filterParams)
                        : "",
                lor: this.filterData.filterRelation,
                filter_down:
                    this.filter_down && Object.keys(this.filter_down).length
                        ? JSON.stringify(this.filter_down)
                        : ""
            };
            // 获取图表数据
            api_chart.getChart(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    res.data.board_statistics &&
                    res.data.board_statistics.length
                ) {
                    this.$set(this.curItem, "data", res.data.board_statistics);
                    this.getOptions(this.curItem, res.data);
                } else {
                    this.option = {};
                    this.isShowChart = false;
                }
                this.loading = false;
                setTimeout(() => {
                    this.imgShow = true;
                }, 50);
            });
        },
        getOptions(obj, res) {
            if (this.isDown) {
                obj.color = this.color;
            }
            // 图处理
            if (obj.mode.indexOf("pie") > -1) {
                // 饼图
                this.option = service.getPieOPtions(obj, this.isDown);
            } else if (
                obj.mode === "base_histogram" ||
                obj.mode === "base_line"
            ) {
                // 其他
                this.option = service.getOtherChartOptions(obj, this.isDown);
                this.option.x_cn = res.x_cn;
            } else {
                // 分组和堆叠
                this.option = service.getGroupBarLineOptions(obj, this.isDown);
                this.option.x_cn = res.x_cn;
                this.option.group_cn = res.group_cn;
            }
            this.isShowChart = true;
        },
        // 获取表格列
        getTableCol() {
            this.sheetCol = [];
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                columns: this.curViewInfo.columns || "",
                module: "list"
            };
            return api.getProgressColConfig(params).then((res) => {
                if (this.$route.name !== "progress") return;
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    let backupCol = _.cloneDeep([
                        // {
                        //     mode: "selection"
                        // }
                        {
                            prop: "operation",
                            label: "操作",
                            fixed: "right",
                            align: "center",
                            width: 80,
                            mode: "operation"
                        }
                    ]);
                    this.sheet.tableCol = backupCol.concat(res.data);
                } else {
                    this.sheet.tableCol = [];
                }
            });
        },
        // 获取表格数据
        getTableData() {
            let filterParams = this.filterData.filterParams;
            let params = {
                board_id: this.curItem.i,
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                filter:
                    filterParams && filterParams.length
                        ? JSON.stringify(filterParams)
                        : "",
                lor: this.filterData.filterRelation,
                page_num: this.sheet.page,
                page_size: this.sheet.page_size,
                filter_down:
                    this.filter_down && Object.keys(this.filter_down).length
                        ? JSON.stringify(this.filter_down)
                        : ""
            };
            api.getProgressListData(params).then((res) => {
                if (this.$route.name !== "progress") return;
                if (res && res.resultCode === 200 && res.data) {
                    this.sheet.tableData = res.data.list || [];
                    this.sheet.total = res.data.cnt || 0;
                } else {
                    this.sheet.tableData = [];
                    this.sheet.total = 0;
                }
                // if (this.$refs.dialogProgressTable) {
                //     this.$refs.dialogProgressTable.refreshTableKey();
                // }
            });
        },
        currentChange(page) {
            this.sheet.page = page;
            this.getTableData();
        },
        pageSizeChange(size) {
            this.sheet.page_size = size;
            this.sheet.page = 1;
            this.getTableData();
        },
        // 选中的表格的数据
        multipleSelection(arr) {
            this.tableMultipleSelection = _.cloneDeep(arr);
        },
        // 删除
        // deleteRow() {
        //     if (!this.tableMultipleSelection.length) return;
        //     // 删除表格项
        //     // 遍历取_id
        //     let idArr = [];
        //     this.tableMultipleSelection.forEach((item) => {
        //         idArr.push(item._id);
        //     });
        //     let params = {
        //         ws_id: this.curSpace.id,
        //         tmpl_id: this.curProgress,
        //         ids: idArr.join(",")
        //     };
        //     api.deleteProgressItems(params).then((res) => {
        //         if (this.$route.name !== "progress") return;
        //         if (res && res.resultCode === 200) {
        //             this.$message({ showClose: true,
        //                 message: "删除成功",
        //                 type: "success"
        //             });
        //             this.getTableData();
        //         }
        //     });
        // },
        // 打开html列
        openHtmlCol(row, col) {
            this.$refs["htmlColDialog"].openDialog(row, col);
        },
        // 更新表格内容
        editFormItem(res, key, id, mode) {
            let file_key = key;
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: id,
                field_key: key,
                [file_key]: res
            };
            api.updateProgress(params).then((res) => {
                if (this.$route.name !== "progress") return;
                if (res && res.resultCode === 200) {
                    this.$refs["htmlColDialog"].cancel();
                    this.refresh();
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
// .chart-table-body {
// width: 100%;
.top-filter {
    position: relative;
    z-index: 999;
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 40px;
    width: 100%;
    box-sizing: border-box;
    // background-color: #f5f6f8;
    background-color: #fff;
    box-shadow: 0px 2px 4px 1px rgba(0, 0, 0, 0.06);
    padding: 0 32px;
    .right {
        display: flex;
        align-items: center;
        .switch-icon {
            display: inline-block;
            width: 20px;
            height: 20px;
            cursor: pointer;
        }
        .table-show {
            background-size: 100% 100%;
            background-image: url(@/assets/image/progress_dashboard/table_show.png);
        }
        .table-hide {
            background-size: 100% 100%;
            background-image: url(@/assets/image/progress_dashboard/table_hide.png);
        }
    }
}

.content {
    box-sizing: border-box;
    position: relative;
    height: calc(100% - 52px);
    padding: 16px 20px;
    /* margin-bottom: 20px; */
    overflow-y: auto;
    background-color: #f5f5f5;
    box-shadow: 0px 2px 4px 1px rgba(0, 0, 0, 0.06);
    .return-btn {
        position: absolute;
        top: 10px;
        left: 20px;
    }
    .charts {
        box-sizing: border-box;
        padding: 16px 12px 8px 8px;
        // margin-bottom: 20px;
        background: #ffffff;
        // box-shadow: 0px 2px 4px 1px rgba(0, 0, 0, 0.06);
        border-radius: 2px 2px 2px 2px;
        border: 1px solid #e6e9f0;
    }
    .table {
        // min-height: 300px;
        padding: 20px;
        margin-top: 12px;
        background: #ffffff;
        // box-shadow: 0px 2px 4px 1px rgba(0, 0, 0, 0.06);
        border-radius: 2px 2px 2px 2px;
        border: 1px solid #e6e9f0;
        // .line {
        //     height: 1px;
        //     margin: 0 0 20px;
        //     background-color: rgba(230, 233, 240, 1);
        // }
        .operation {
            text-align: right;
            margin-bottom: 16px;
        }
    }
}
.btn-group ::v-deep {
    .el-button--small {
        height: 32px;
        line-height: 32px;
        padding: 0 15px;
        font-size: 14px;
    }
    .el-dropdown__caret-button {
        padding: 0 5px;
    }
}
.btn-group .add {
    display: inline-block;
    width: 16px;
    height: 16px;
    margin-right: 2px;
    vertical-align: text-bottom;
    background-image: url(@/assets/image/common/add_white.png);
    background-size: 100% 100%;
}
.export-box {
    display: inline-block;
    width: 16px;
    height: 16px;
    margin-right: 2px;
    vertical-align: text-bottom;
    background-image: url(@/assets/image/common/export.svg);
    background-size: 100% 100%;
}
.el-dropdown-menu.el-popper.operation-btn.el-dropdown-menu--small {
    padding: 0 !important;
    margin-right: 0;
    margin-top: 2px;
    .el-dropdown-menu__item {
        display: inline-block;
        height: 32px;
        line-height: 32px;
        box-sizing: border-box;
        width: 97px;
        padding: 0 12px;
        color: #171e31;
        font-size: 14px;
        margin: 4px !important;
        cursor: pointer;
    }
}
</style>
<style lang="scss">
.basic-ui.full-screen-dialog .el-dialog {
    width: calc(100% - 164px);
    height: calc(100% - 120px);
    max-height: 900px;
    max-width: 1300px;
    .el-dialog__header {
        padding: 24px 32px 16px;
    }
    .el-dialog__body {
        height: calc(100% - 64px);
        max-height: calc(100% - 64px);
        border-top: 1px solid #e6e9f0;
        padding: 0;
        overflow: inherit;
    }
    .el-dialog__footer {
        box-shadow: 0px -3px 8px 1px rgba(0, 0, 0, 0.1);
    }
}
.steps2.basic-ui.full-screen-dialog .el-dialog {
    .el-dialog__body {
        .chart-table-body {
            margin-right: -31px;
        }
    }
}
</style>
