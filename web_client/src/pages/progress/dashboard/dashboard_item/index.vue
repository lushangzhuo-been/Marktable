<template>
    <div class="dashboard-content">
        <grid-layout
            v-if="isShowDashboard"
            :layout.sync="layout"
            :col-num="colNum"
            :row-height="90"
            :is-draggable="true"
            :is-resizable="true"
            :is-mirrored="false"
            :vertical-compact="true"
            :margin="[10, 10]"
            :use-css-transforms="true"
        >
            <grid-item
                v-for="item in layout"
                :x="item.x"
                :y="item.y"
                :w="item.w"
                :h="item.h"
                :i="item.i"
                :key="item.i"
                :minH="3"
                :minW="3"
                @move="moveEvent"
                @moved="movedEvent"
                @resize="resizeEvent"
                @resized="resizedEvent"
                @container-resized="containerResizedEvent"
                @layout-updated="layoutUpdatedEvent"
                drag-allow-from=".drag-icon"
            >
                <div
                    class="grid-item-block"
                    @mouseover="mouseChange('over', item)"
                    @mouseout="mouseChange('out', item)"
                >
                    <!-- 标题栏 -->
                    <div class="head">
                        <!-- 拖拽-标题 -->
                        <div class="left">
                            <b
                                class="drag-icon"
                                v-show="isMouseHover && curItem.i === item.i"
                            ></b>
                            <b class="filter-icon"></b>
                            <div t-overflow class="title">
                                <span> {{ item.title }}</span>
                            </div>
                        </div>
                        <!-- 操作 -->
                        <div class="right">
                            <!-- 展示图表弹窗 -->
                            <b
                                class="opertion-item-box refresh-chart"
                                @click="refreshChart(item)"
                            ></b>
                            <!-- 展示图表弹窗 -->
                            <b
                                class="opertion-item-box full-screen"
                                @click="openFullScreenDialog(item)"
                            ></b>
                            <!-- 三个点更多设置 -->
                            <el-dropdown
                                trigger="click"
                                class="more-setting"
                                placement="bottom-start"
                                @command="
                                    (command) => gridOperation(command, item)
                                "
                            >
                                <b
                                    class="more-operation"
                                    :class="
                                        isMouseHover && curItem.i === item.i
                                            ? 'box-over'
                                            : 'box-out'
                                    "
                                ></b>
                                <el-dropdown-menu slot="dropdown">
                                    <el-dropdown-item
                                        class="basic-ui"
                                        command="edit"
                                    >
                                        <b class="opertion-item-box edit"></b>
                                        编辑
                                    </el-dropdown-item>
                                    <el-dropdown-item
                                        class="basic-ui"
                                        command="delete"
                                    >
                                        <b class="opertion-item-box delete"></b>
                                        删除
                                    </el-dropdown-item>
                                    <el-dropdown-item
                                        class="basic-ui"
                                        command="copy"
                                    >
                                        <b class="opertion-item-box copy"></b>
                                        复制
                                    </el-dropdown-item>
                                    <el-dropdown-item
                                        class="basic-ui"
                                        command="download"
                                    >
                                        <b
                                            class="opertion-item-box download"
                                        ></b>
                                        下载
                                    </el-dropdown-item>
                                </el-dropdown-menu>
                            </el-dropdown>
                        </div>
                    </div>
                    <div class="content">
                        <dashboard-item-chart
                            :ref="`chart${item.i}`"
                            :item="item"
                            :data="item.data"
                            @on-chart-click="chartClick"
                        ></dashboard-item-chart>
                    </div>
                </div>
            </grid-item>
        </grid-layout>
        <div
            v-else
            class="no-dashboard-data"
            v-loading="loading"
            element-loading-text="加载中"
            element-loading-background="rgba(255, 255, 255)"
        >
            <div v-show="imgShow" class="empty-box">
                <div class="img"></div>
                <div class="text">
                    <div>使用多个图标可以组成可视化看板</div>
                    <div>如需添加图表，请点击右上角【+新增】按钮</div>
                </div>
            </div>
        </div>
        <!-- 确认删除dialog -->
        <delete-board-dialog
            ref="DeleteBoardDialog"
            @on-confirm="confirmDeleteItem"
        ></delete-board-dialog>
        <!-- 全屏dialog -->
        <chart-table-dialog
            ref="chartTableDialog"
            :filterData="filterData"
            :curViewInfo="curViewInfo"
            @dashboard-refresh="dashboardRefresh"
        ></chart-table-dialog>
        <!-- 编辑&新增dialog -->
        <edit-board
            ref="editBoardDialog"
            :curTabItems="curTabItems"
            :filterData="filterData"
            @on-confirm="confirmEditBoard"
        ></edit-board>
    </div>
</template>

<script>
import _ from "lodash";
import VueGridLayout from "vue-grid-layout";
import api from "@/common/api/module/progress_dashboard";
import DashboardItemChart from "./dashboard_chart.vue";
import EditBoard from "../common/add_board_dialog/index.vue";
import DeleteBoardDialog from "../common/delete_board_dialog";
import ChartTableDialog from "../common/full_screen_dashboard_dialog/index.vue";
import dataHandle from "../data_handle";
export default {
    components: {
        GridLayout: VueGridLayout.GridLayout,
        GridItem: VueGridLayout.GridItem,
        EditBoard,
        DashboardItemChart,
        // PieChart,
        // BarChart,ss
        // LineChart,
        // Number,
        ChartTableDialog,
        DeleteBoardDialog
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
        },
        curViewInfo: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            layout: [],
            colNum: 12,
            index: 100,
            isShowDashboard: false,
            loading: true,
            imgShow: false,
            count: 0,
            isMouseHover: false,
            curItem: {}
        };
    },
    computed: {
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
    watch: {
        curTabItems: {
            handler(curTab) {
                // 点击tab时候触发
                if (
                    curTab &&
                    Object.keys(curTab).length &&
                    curTab.mode === "board"
                ) {
                    if (this.count === 0) {
                        this.fetchDashboardConfig();
                        this.fetchDashboardList();
                    } else {
                        this.updateChart();
                    }
                    this.count++;
                }
            },
            deep: true,
            immediate: true
        },
        layout: {
            handler(arr) {
                if (arr && arr.length) {
                    this.isShowDashboard = true;
                } else {
                    this.isShowDashboard = false;
                }
                this.$emit("layout-change", this.isShowDashboard);
            },
            deep: true,
            immediate: true
        }
    },

    methods: {
        chartClick(chart, option, item) {
            this.$refs.chartTableDialog.openDialog(item, true, chart, option);
        },
        mouseChange(type, item) {
            this.curItem = _.cloneDeep(item);
            // this.$set(item, "isMouseHover", type === "over" ? true : false);
            this.isMouseHover = type === "over" ? true : false;
        },
        // 拖拽以后
        layoutUpdatedEvent(newLayout) {},
        rename(item) {
            this.$set(item, "rename", true);
            this.$nextTick(() => {
                this.$refs[`rename${item.i}`][0].focus();
            });
        },
        // 获取dashboard的config数据
        fetchDashboardConfig() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress
            };
            this.$store
                .dispatch("fetchDashboardConfig", params)
                .then((res) => {});
        },
        updateChart() {
            this.layout.forEach((item) => {
                this.fetchCharts(item);
            });
        },
        // 获取dashboard列表数据
        fetchDashboardList() {
            this.layout = [];
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                view_id: this.curTabItems.id
            };

            api.getDashboardList(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    if (res.data.length) {
                        // this.layout = res.data;
                        let tmpLayout = _.cloneDeep(res.data);
                        tmpLayout.forEach((item) => {
                            item.i = item.id;
                            delete item.id;
                        });
                        this.layout = tmpLayout;
                        this.layout.forEach((item) => {
                            this.fetchCharts(item);
                        });
                    }
                }
                setTimeout(() => {
                    this.imgShow = true;
                }, 50);
                this.loading = false;
            });
        },
        fetchCharts(item, isCloseFullScreen = false) {
            let filterParams = this.filterData.filterParams;
            let params = {
                board_id: item.i,
                filter:
                    filterParams && filterParams.length
                        ? JSON.stringify(filterParams)
                        : "",
                lor: this.filterData.filterRelation
            };
            api.getChart(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    res.data.board_statistics &&
                    res.data.board_statistics.length
                ) {
                    if (isCloseFullScreen) {
                        this.layout.forEach((sub) => {
                            if (sub.i === item.i) {
                                this.$set(sub, "x_cn", res.data.x_cn);
                                this.$set(sub, "group_cn", res.data.group_cn);
                                this.$set(
                                    sub,
                                    "data",
                                    res.data.board_statistics
                                );
                            }
                        });
                    } else {
                        this.$set(item, "group_cn", res.data.group_cn);
                        this.$set(item, "x_cn", res.data.x_cn);
                        this.$set(item, "data", res.data.board_statistics);
                    }
                } else {
                    if (isCloseFullScreen) {
                        this.layout.forEach((sub) => {
                            if (sub.i === item.i) {
                                this.$set(sub, "data", []);
                            }
                        });
                    } else {
                        this.$set(item, "data", []);
                    }
                }
            });
        },
        refreshChart(item) {
            this.fetchCharts(item, true);
        },
        openFullScreenDialog(item) {
            this.$refs.chartTableDialog.openDialog(item);
        },
        addChart(params, data, id) {
            let obj = {
                x: (this.layout.length * 4) % (this.colNum || 12),
                y: this.layout.length + (this.colNum || 12), // puts it at the bottom
                w: 4,
                h: 4,
                i: id,
                ...params,
                data
            };
            this.layout.push(obj);
            this.setResizeDashboard(); //新增完后将全部图表的位置存储到后端
            this.fetchCharts(obj);
        },
        setResizeDashboard() {
            let tmpLayout = _.cloneDeep(this.layout);
            tmpLayout.forEach((item) => {
                item.id = item.i;
                delete item.i;
            });
            api.refresLocation({ locations: tmpLayout }).then((res) => {});
        },
        resizeEvent() {},
        resizedEvent(i, newH, newW, newHPx, newWPx) {
            this.$nextTick(() => {
                // 调用echarts重绘
                if (this.$refs["chart" + i][0]) {
                    this.$refs["chart" + i][0].chartResize();
                }
            });
            this.setResizeDashboard();
        },
        moveEvent() {
            // this.setResizeDashboard();
            // 拿起时，获取dom最大高度/每行高度，可以获得当前最大行数， 用于配置阴影
        },
        movedEvent() {
            this.setResizeDashboard();
        },
        containerResizedEvent() {},
        // 操作
        gridOperation(command, item) {
            let mapping = {
                edit: "gridItemEdit", // 改名
                delete: "girdItemDelete", //删除
                copy: "girdItemCopy", //删除
                download: "gridItemDowndload" //下载
            };
            this[`${mapping[command]}`](item);
        },
        renameInputBlur(item) {
            item.rename = false;
        },
        // 设置编辑图表
        gridItemEdit(item) {
            this.curItem = item;
            this.$refs.editBoardDialog.openDialog(item);
        },
        // 确认新增图
        confirmEditBoard(form) {
            let tmpForm = _.cloneDeep(form);
            tmpForm.mode = dataHandle.chartTypeMpping[form.chartType] || "";
            tmpForm.show_empty = tmpForm.show_empty ? "yes" : "no";
            tmpForm.order = tmpForm.order === "not" ? "" : tmpForm.order;
            // 确认新增图
            let params = {
                id: form.i,
                ws_id: this.curSpace.id,
                tmpl_id: Number(this.curProgress),
                view_id: this.curTabItems.id,
                userid: this.userInfo.id,
                filter: "",
                ...tmpForm
            };
            delete params.chartType;
            api.modifyDashboardItem(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.$refs.editBoardDialog.stepsTwoCancel();
                    this.layout.forEach((sub) => {
                        if (sub.i === form.i) {
                            Object.assign(sub, tmpForm);
                            this.fetchCharts(sub);
                        }
                    });
                }
            });
        },
        // 删除
        girdItemDelete(item) {
            this.$refs.DeleteBoardDialog.openDialog(item);
        },
        // 确认删除
        confirmDeleteItem(item) {
            api.deleteDashboardItem({ board_id: item.i }).then((res) => {
                if (res && res.resultCode === 200) {
                    // 1、提交新位置给后端
                    // 2.重新获取list
                    this.$message({
                        showClose: true,
                        message: "删除成功",
                        type: "success"
                    });
                    this.layout.splice(
                        this.layout.findIndex((sub) => sub.i === item.i),
                        1
                    );
                    this.setResizeDashboard();
                }
            });
        },
        // 复制
        girdItemCopy(item) {
            let params = {
                id: item.i,
                w: item.w,
                h: item.h,
                x: (this.layout.length * 4) % (this.colNum || 12),
                y: this.layout.length + (this.colNum || 12)
            };
            let copyItem = _.cloneDeep(item);
            api.copyDashboardItem(params).then((res) => {
                if (res && res.resultCode === 200 && res.data && res.data.id) {
                    copyItem.i = res.data.id;
                    copyItem.x = res.data.x;
                    copyItem.y = res.data.y;
                    copyItem.title = res.data.title;
                    this.layout.push(copyItem);
                    this.$message({
                        showClose: true,
                        message: "复制成功",
                        type: "success"
                    });
                    this.setResizeDashboard();
                }
            });
        },
        // 下载
        gridItemDowndload(item) {
            this.$refs["chart" + item.i][0].exportImg();
        },
        // 关闭全屏刷新当前item
        dashboardRefresh(item) {
            this.fetchCharts(item, true);
        },
        // item新增确定位置
        newItemPositionad(item, itemId) {
            // 初始化元素
            let newItem = {
                ...item,
                i: itemId,
                x: 0,
                y: 0,
                w: item.w,
                h: item.h
            };
            // 确定边界
            let Ys = [],
                maxX = 0,
                maxY = 0,
                edgeX = 0,
                edgeY = 0;
            this.layout.forEach((item) => {
                Ys.push(item.y + item.h);
            });
            maxY = (Ys.length && Math.max.apply(null, Ys)) || 1;
            edgeX = 12;
            edgeY = maxY;

            // 使用二维数组生成地图
            let gridMap = new Array();
            for (let x = 0; x < edgeX; x++) {
                gridMap[x] = new Array();
                for (let y = 0; y < edgeY; y++) {
                    gridMap[x][y] = 0;
                }
            }

            // 标记占位
            this.layout.forEach((item) => {
                // 将layout中卡片所占区域标记为1
                for (let x = item.x; x < item.x + item.w; x++) {
                    for (let y = item.y; y < item.y + item.h; y++) {
                        gridMap[x][y] = 1;
                    }
                }
            });
            // 遍历地图，申请位置
            for (let y = 0; y < edgeY; y++) {
                for (let x = 0; x < edgeX; x++) {
                    // 申请所需空间
                    if (edgeX - x >= item.w && edgeY - y >= item.h) {
                        let itemSignArr = [];
                        for (let a = x; a < x + item.w; a++) {
                            for (let b = y; b < y + item.h; b++) {
                                itemSignArr.push(gridMap[x][y]);
                            }
                        }
                        if (itemSignArr.indexOf(1) < 0) {
                            newItem.x = x;
                            newItem.y = y;
                            return newItem;
                        }
                    }
                }
            }
            // 无满足条件
            newItem.x = 0;
            newItem.y = edgeY + 1;
            return newItem;
        }
    }
};
</script>

<style lang="scss" scoped>
// 放在外面是因为popover也使用
.opertion-item-box {
    display: inline-block;
    width: 20px;
    height: 20px;
    vertical-align: middle;
    background-size: 100% 100%;
    margin-right: 8px;
    cursor: pointer;
    &.edit {
        background-image: url(@/assets/image/common/edit.svg);
    }
    &.delete {
        background-image: url(@/assets/image/common/delete.svg);
    }
    &.copy {
        background-image: url(@/assets/image/progress_column/copy.svg);
    }
    &.download {
        background-image: url(@/assets/image/common/download.svg);
    }
    &.refresh-chart {
        background-image: url(@/assets/image/progress_dashboard/refresh_chart.png);
    }
    &.full-screen {
        background-image: url(@/assets/image/common/full_screen.png);
    }
}
.dashboard-content {
    background-color: #ffffff;

    ::v-deep .vue-grid-layout {
        margin-top: 12px;
        background-color: #f5f5f5;
    }
    .vue-grid-item {
        box-sizing: border-box;
        background-color: #ffffff;
        touch-action: none;
        border: 1px solid #e6e9f0;
        border-radius: 2px;
        // box-shadow: 0px 2px 4px 1px rgba(0, 0, 0, 0.06);
    }

    // 内容结构
    .grid-item-block {
        display: flex;
        flex-direction: column;
        height: 100%;
        width: 100%;
        .head {
            display: flex;
            align-items: center;
            justify-content: space-between;
            height: 40px;
            padding: 0 10px;
            // background: rgba($color: #f5f6f8, $alpha: 0.5);
            background-color: #fff;
            border-radius: 8px 8px 0 0;
            border-bottom: 1px solid #e6e9f0;
            .left {
                position: relative;
                display: flex;
                align-items: center;
                width: calc(100% - 78px);
                .drag-icon {
                    position: absolute;
                    left: -8px;
                    display: inline-block;
                    width: 24px;
                    height: 24px;
                    background-image: url(@/assets/image/common/icon_dragtable_move.png);
                    background-size: 100% 100%;
                }
                .title {
                    width: calc(100% - 40px);
                    height: 32px;
                    line-height: 32px;
                    margin-left: 8px;
                    span {
                        max-width: 100%;
                        text-overflow: ellipsis;
                        overflow: hidden;
                        white-space: nowrap;
                        display: inline-block;
                        padding: 0 8px;
                        border-radius: 4px;
                        // &:hover {
                        //     background-color: #efefef;
                        // }
                    }
                }
                .title-input {
                    max-width: 150px;
                }
            }
            .right {
                display: flex;
                justify-content: flex-end;
                align-items: center;
                .more-setting {
                    line-height: 1;
                    height: auto;
                }
                .more-operation {
                    display: inline-block;
                    width: 20px;
                    height: 20px;
                    cursor: pointer;
                    background-size: 100% 100%;
                    background-image: url(@/assets/image/common/more-operation.png);
                    &.box-out {
                        background-image: url(@/assets/image/common/more-operation.png);
                    }
                    &.box-over {
                        background-image: url(@/assets/image/common/space_more_default.png);
                    }

                    &:hover {
                        background-image: url(@/assets/image/common/space_more_active.png);
                    }
                }
            }
        }

        .content {
            flex: 1;
            height: calc(100% - 42px);
            width: calc(100% - 32px);
            margin: 16px;
        }
    }

    .rename-input.el-input {
        width: 160px;
    }
    .no-dashboard-data {
        box-sizing: border-box;
        text-align: center;
        height: 570px;
        padding: 150px 0 100px;
        .empty-box {
            margin: auto;
            width: 558px;
            .img {
                width: 100%;
                height: 390px;
                background-image: url(@/assets/image/progress_dashboard/no_dashboard.png);
                background-size: 100% 100%;
            }
            .text {
                color: #a6abbc;
                font-size: 14px;
                line-height: 24px;
                margin-top: 32px;
            }
        }
    }
}
</style>
<style lang="scss">
.vue-grid-item.vue-grid-placeholder {
    background: #cccccc;
}
</style>
