<template>
    <div class="mark-table-dashboard">
        <div class="top-operation">
            <div>
                <!-- <div class="filter-box" v-show="isFilterIconShow">
                    <b class="filter-icon"></b>
                    筛选
                </div> -->
                <dashboard-filters
                    v-show="isFilterIconShow"
                    ref="FilterBlock"
                    class="operation-btn"
                    :class="{
                        'view-had-filter': filterData.filterParams.length
                    }"
                    :filterRelation="filterData.filterRelation"
                    :viewFilter="viewFilter"
                    :curTabItems="curTabItems"
                    :viewFilterStandby="filterData.viewFilterStandby"
                    @save-filter="saveFilter"
                    @filter-search="filterSearch"
                    @get-view-filter="getViewFilter"
                ></dashboard-filters>
            </div>
            <div class="operation-right">
                <el-button
                    type="primary"
                    class="basic-ui btnsss"
                    size="small"
                    @click="addBoard"
                    title="新增"
                >
                    +新增</el-button
                >
            </div>
        </div>
        <!-- dashboard盒子内容 -->
        <dashboard
            ref="dashBoard"
            :curTabItems="curTabItemsObj"
            :filterData="filterData"
            :curViewInfo="curViewInfo"
            @layout-change="layoutChange"
        ></dashboard>
        <!-- 新增box -->
        <add-board
            ref="addBoardDialog"
            :curTabItems="curTabItemsObj"
            :filterData="filterData"
            @on-confirm="confirmAddBoard"
        ></add-board>
    </div>
</template>

<script>
import Dashboard from "./dashboard_item/index.vue";
import AddBoard from "./common/add_board_dialog/index.vue";
import api_progress from "@/common/api/module/progress";
import api from "@/common/api/module/progress_dashboard";
import dataHandle from "./data_handle.js";
import _ from "lodash";
import DashboardFilters from "./common/filter/index";
export default {
    components: {
        Dashboard,
        AddBoard,
        DashboardFilters
    },
    props: {
        curTabItems: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            isFilterIconShow: false,
            orderMapping: ["", "asc", "desc"],
            viewFilter: [], // 当前选中视图，保存的过滤条件，
            curTabItemsObj: {},
            filterData: {
                filterParams: "",
                viewFilterStandby: [],
                filterRelation: "filter_and"
            },
            curViewInfo: {}
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
        // 监听是否点击到dashboard
        curTabItems: {
            handler(curTab) {
                // 点击tab时候触发
                if (
                    curTab &&
                    Object.keys(curTab).length &&
                    curTab.mode === "board"
                ) {
                    this.fetchViewInfo();
                }
            },
            deep: true,
            immediate: true
        }
    },
    methods: {
        fetchViewInfo() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: this.curTabItems.id
            };
            api_progress.getViewInfo(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    this.curViewInfo = res.data;
                    this.filterForCheck(this.curViewInfo);
                } else {
                    this.curViewInfo = {};
                }
            });
        },
        //
        filterForCheck(curTab) {
            // filterRelation
            this.filterData.filterRelation = curTab.lor || "filter_and";
            // filterParams
            if (curTab.filter) {
                this.filterData.filterParams = JSON.parse(curTab.filter);
            } else {
                this.filterData.filterParams = [];
            }
            // viewFilterStandby
            if (curTab.format_filter && curTab.format_filter !== "null") {
                this.filterData.viewFilterStandby = JSON.parse(
                    curTab.format_filter
                );
            } else {
                this.filterData.viewFilterStandby = [];
            }
            this.viewFilter = _.cloneDeep(this.filterData.viewFilterStandby);
            let tmpTab = {
                ..._.cloneDeep(curTab),
                ...this.filterParams
            };
            if (this.$refs.dashBoard) {
                this.$refs.dashBoard.count = 0;
            } else {
                this.$nextTick(() => {
                    this.$refs.dashBoard.count = 0;
                });
            }
            this.count = 0;
            this.curTabItemsObj = tmpTab;
            delete tmpTab.viewFilterStandby;
            delete tmpTab.filterRelation;
        },
        // 确认筛选
        filterSearch(params, viewFilter, filterRelation) {
            this.filterData = {
                filterParams: params,
                viewFilterStandby: viewFilter,
                filterRelation: filterRelation
            };
            let tmpTab = _.cloneDeep(this.curTabItemsObj);
            tmpTab.filterParams = params;
            tmpTab.viewFilterStandby = viewFilter;
            tmpTab.filterRelation = filterRelation;
            this.$nextTick(() => {
                this.$refs.dashBoard.count++;
            });
            delete tmpTab.viewFilterStandby;
            delete tmpTab.filterRelation;
            this.curTabItemsObj = tmpTab;
        },
        // 保存筛选
        saveFilter(obj, order) {
            this.saveFilterGroup = obj;
            // 新视图打开dialog确认明细 否则直接更新
            this.updateViewFilter(order);
        },
        // 更新过滤
        updateViewFilter(order) {
            // 更新视图
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: Number(this.curProgress),
                id: this.curTabItems.id, //视图id
                filter: JSON.stringify(this.saveFilterGroup),
                lor: order
            };
            api_progress.updateViewFilter(params).then((res) => {
                if (this.$route.name !== "progress") return;
                this.$refs.FilterBlock.closePopver();
                this.$refs.dashBoard.fetchDashboardList();
                this.$emit("dashboard-save-filter");
            });
        },
        getViewFilter() {
            this.viewFilter = _.cloneDeep(this.filterData.viewFilterStandby);
        },
        layoutChange(flag) {
            if (!flag) {
                this.$refs.FilterBlock && this.$refs.FilterBlock.closePopver();
            }
            this.isFilterIconShow = flag;
        },
        // 新增图弹窗
        addBoard() {
            this.$refs.addBoardDialog.openDialog();
        },
        // 确认新增图
        confirmAddBoard(form) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: Number(this.curProgress),
                view_id: this.curTabItems.id,
                userid: this.userInfo.id,
                filter: "",
                ...form
            };
            params.mode = dataHandle.chartTypeMpping[form.chartType] || "";
            params.show_empty = params.show_empty ? "yes" : "no";
            params.order = params.order === "not" ? "" : params.order;
            delete params.chartType;
            api.addDashboardItem(params).then((res) => {
                if (res && res.resultCode === 200 && res.data && res.data.id) {
                    let chartData = res.data.dashboard_data
                        ? res.data.dashboard_data.board_statistics || []
                        : [];
                    this.$refs.dashBoard.addChart(
                        params,
                        chartData,
                        res.data.id
                    );
                    this.$refs.addBoardDialog.stepsTwoCancel();
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.mark-table-dashboard {
    .top-operation {
        display: flex;
        justify-content: space-between;
        align-items: center;
        .filter-box {
            height: 20px;
            cursor: pointer;
            .filter-icon {
                display: inline-block;
                width: 20px;
                height: 20px;
                background-size: 100% 100%;
                background-image: url(@/assets/image/progress/filter.png);
                vertical-align: middle;
                &:hover {
                    background-size: 100% 100%;
                    background-image: url(@/assets/image/progress/filter-active.png);
                }
            }
            &:hover {
                color: var(--PRIMARY_COLOR);
                .filter-icon {
                    background-size: 100% 100%;
                    background-image: url(@/assets/image/progress/filter-active.png);
                }
            }
        }
    }
}
</style>
