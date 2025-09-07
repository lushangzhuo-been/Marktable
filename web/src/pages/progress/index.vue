<template>
    <div v-if="noProgressAuth">
        <no-permissions></no-permissions>
    </div>
    <div v-else>
        <!-- 标题及管理员信息 -->
        <div class="title-administrators">
            <div class="title">{{ progressInfo.tmpl.name || "--" }}</div>
        </div>
        <!-- 流程描述 -->
        <div
            v-show="progressInfo.tmpl.desc"
            class="progress-desc"
            id="ProgressDesc"
            ref="ProgressDesc"
            :class="{ 'show-more': showMore, 'single-col': !showMore }"
        >
            <div
                class="desc-content"
                :class="{ 'show-more': showMore, 'single-col': !showMore }"
                id="DescContent"
                ref="DescContent"
            >
                {{ progressInfo.tmpl.desc }}
                <div
                    v-if="expandBtnShow && showMore"
                    class="expand-btn close"
                    @click="showMore = !showMore"
                >
                    {{ showMore ? "收起" : "展开" }}
                </div>
            </div>
            <div
                v-if="expandBtnShow && !showMore"
                class="expand-btn"
                @click="showMore = !showMore"
            >
                {{ showMore ? "收起" : "展开" }}
            </div>
        </div>
        <card-tabs
            :customerClass="'progress-card-001'"
            :tabsList="tabsList"
            :currentTab="currentTab"
            @tab-change="tabChange"
            @add-tab="addTab"
            @lock-tab="lockTab"
            @unlock-tab="unlockTab"
            @delete-view="deleteView"
        ></card-tabs>
        <!-- 列表 -->
        <div v-if="currentTabType === 'list' || currentTabType === 'card'">
            <!-- 操作与新增 -->
            <div class="operation-block">
                <div class="operation-left">
                    <filters
                        ref="FilterBlock"
                        class="operation-btn"
                        :class="{ 'view-had-filter': filterParams.length }"
                        :viewFilter="viewFilter"
                        :viewFilterStandby="viewFilterStandby"
                        :filterRelation="filterRelation"
                        :currentTab="currentTab"
                        @save-filter="saveFilter"
                        @filter-search="filterSearch"
                        @get-view-filter="getViewFilter"
                    ></filters>
                    <sorts
                        ref="SortBlock"
                        class="operation-btn"
                        :class="{ 'view-had-sort': sortOrder.length }"
                        :currentTab="currentTab"
                        :sortOrder="sortOrder"
                        @save-sort="saveSort"
                        @sort-search="sortSearch"
                    >
                    </sorts>
                    <show-hide
                        ref="HiddenShowBlock"
                        class="operation-btn"
                        :hiddenShowColumnsList="hiddenShowColumnsList"
                        :currentTab="currentTab"
                        @save-hidden-show="saveHiddenShow"
                        @check-change="checkChange"
                    ></show-hide>
                    <group-by
                        ref="GroupByBlock"
                        class="operation-btn"
                        :currentTab="currentTab"
                        :groupByOption="groupByOption"
                        :cardFilterDown="cardFilterDown"
                        @confim-group-by="confimGroupBy"
                        @check-field-search="checkFieldSearch"
                        @group-by-save-view="groupBySaveView"
                    ></group-by>
                </div>
                <div class="operation-right">
                    <el-button
                        class="basic-ui add-btn"
                        size="small"
                        type="primary"
                        :disabled="!createPrmission"
                        @click="addProgress"
                    >
                        <b class="add"></b>新增
                    </el-button>
                    <span class="split-line">|</span>
                    <el-popover
                        placement="bottom-end"
                        trigger="click"
                        popper-class="progress-more-operate-poppover"
                        :visible-arrow="false"
                    >
                        <div
                            class="pop-item"
                            :class="exportPrmission ? '' : 'disabled'"
                            @click="handleOpreate('export')"
                        >
                            <img
                                :src="
                                    require(`@/assets/image/common/export.svg`)
                                "
                                alt=""
                                width="18px"
                                height="18px"
                            />
                            导出
                        </div>
                        <b
                            slot="reference"
                            class="progress-more-operate-icon"
                            width="18px"
                            height="18px"
                        >
                        </b>
                    </el-popover>
                </div>
            </div>
            <div
                v-show="currentTabType === 'list'"
                v-loading="tableLoading"
                class="table-loading"
                element-loading-text="加载中"
                element-loading-background="rgba(255, 255, 255)"
            >
                <!-- 维度枚举 -->
                <div
                    class="gourp-by-with-table"
                    :class="{
                        'had-group-by-temp':
                            groupByFieldInfo && groupByFieldInfo.field_key
                    }"
                >
                    <group-by-temp
                        v-if="groupByFieldInfo && groupByFieldInfo.field_key"
                        ref="GroupByTemp"
                        :groupByFieldInfo="groupByFieldInfo"
                        :groupByList="groupByList"
                        :cardFilterDown="cardFilterDown"
                        :currentTab="currentTab"
                        @check-field-search="checkFieldSearch"
                    ></group-by-temp>
                    <p-table
                        ref="ProgressTable"
                        class="table-content"
                        :data="tableData"
                        :col="tableCol"
                        @open-html-col="openHtmlCol"
                        @open-detail="openDetail"
                        @multiple-selection="multipleSelection"
                        @refresh="refreshCurrentPage"
                        @copy-link="copyLink"
                        @delete-row="deleteRow"
                        @change-column-width="changeColumnWidth"
                    ></p-table>
                </div>
                <el-pagination
                    v-show="table.count > 20"
                    class="basic-ui"
                    @current-change="currentChange"
                    layout="total, prev, pager, next"
                    :current-page="table.page_num"
                    :page-size="20"
                    :total="table.count"
                >
                </el-pagination>
            </div>

            <!-- 看板 -->
            <div v-show="currentTabType === 'card'" class="kanban-view-block">
                <card-board
                    ref="KanBan"
                    class="kanban-block"
                    :currentTab="currentTab"
                    :curTabItems="curTabItems"
                    :filterParams="filterParams"
                    :sortOrder="sortOrder"
                    :filterRelation="filterRelation"
                    :cardFilterDown="cardFilterDown"
                ></card-board>
            </div>
        </div>
        <!-- 仪表盘 -->
        <div v-if="currentTabType === 'board'">
            <dashboard
                :curTabItems="curTabItems"
                :filterParams="filterParams"
                @dashboard-save-filter="refreshView"
            ></dashboard>
        </div>
        <add-progress
            ref="AddProgress"
            :formLabel="createCol"
            @refresh="refreshForAdd"
        ></add-progress>
        <save-filter ref="SaveFilter" @create-view="createView"></save-filter>
        <!-- html富文本编辑器 -->
        <html-col-dialog
            ref="htmlColDialog"
            @edit-form-item="editFormItem"
        ></html-col-dialog>
        <detail-drawer
            ref="DetailDrawer"
            @refresh="refreshCurrentPage"
        ></detail-drawer>
        <!-- 删除表格行提醒 -->
        <delete-table-row-dialog
            ref="deleteTableRowDialog"
            @cancel="cancelDeleteTableRow"
            @on-confirm="comfirmDeleteTableRow"
        ></delete-table-row-dialog>
    </div>
</template>
<script>
import dashApi from "@/common/api/module/progress_dashboard";
import NoPermissions from "../common/no_permissions_progress.vue";
import CardTabs from "@/components/tabs/card_tabs/index";
import PTable from "@/components/table/p_table/progress_table";
import ShowHide from "./show_hide/index";
import Filters from "./filter/index_v2";
import DetailDrawer from "./detail/index";
import HtmlColDialog from "./components/html_col_dialog.vue";
import AddProgress from "./add_progress/index";
import Dashboard from "./dashboard/index";
import SaveFilter from "./components/save_filter_dialog.vue";
import _ from "lodash";
import api from "@/common/api/module/progress";
import HandleData from "./data_handle";
import { imgHost } from "@/assets/tool/const";
import Sorts from "./sort/index";
import GroupBy from "./group_by/index";
import GroupByTemp from "@/pages/progress/card_board/group_by_temp";
import CardBoard from "./card_board/index.vue";
import UserMessage from "@/components/user_message_tip";
import DataHandle from "@/pages/progress/card_board/data_handle.js";
// 删除表格行提醒弹窗
import DeleteTableRowDialog from "./components/delete_table_row_dialog.vue";
export default {
    components: {
        CardTabs,
        PTable,
        ShowHide,
        Filters,
        DetailDrawer,
        HtmlColDialog,
        Dashboard,
        AddProgress,
        SaveFilter,
        NoPermissions,
        Sorts,
        GroupBy,
        GroupByTemp,
        CardBoard,
        UserMessage,
        DeleteTableRowDialog
    },

    data() {
        return {
            title: "流程table",
            createCol: [],
            tableCol: [],
            tableData: [],
            table: {
                count: 0,
                page_num: 1,
                page_size: 20
            },
            formData: {},
            currentTabType: "",
            currentTab: "",
            tabsList: [],
            saveFilterGroup: [], //保存新视图的条件
            sortOrder: [], // 保存新视图排序
            saveHiddenShowColumns: [], //保存新视图的显示隐藏列
            tableMultipleSelection: [], //表格多选用于删除
            hiddenShowColumnsList: [], //全量显示隐藏列带check结果
            hideShowCol: [],
            hadSelectShowCol: [],
            filterParams: [],
            backupParams: [], // 备份filter参数，  前后不一致再调接口
            sortParams: [],
            viewFilter: [], // 当前选中视图，保存的过滤条件
            viewFilterStandby: [], // 打开过滤时，赋给viewFilter
            viewId: null, // 当前打开的视图id
            viewMode: "list",
            progressInfo: {
                admin_list: [],
                tmpl: {}
            },
            homeView: {
                id: -1,
                name: "主列表",
                mode: "list",
                filter: ""
            },
            noProgressAuth: true, // 默认无权限
            exportPrmission: false,
            createPrmission: false,
            curTabItems: {},
            tableLoading: false, // 表格加载 1、模板切换，分页，过滤
            filterRelation: "filter_and",
            expandBtnShow: false,
            showMore: false,
            listView: "list", // 列表的视图 切换卡片
            groupByOption: [], // 分组字段列表
            // cardsFieldShow: [], // 卡片显示的字段
            cardFilterDown: "", // 卡片xy轴入参信息， 用来从viewinfo接口取回显信息回显用
            groupByFieldInfo: {}, //分组字段信息
            groupByList: [],
            groupByFieldEnumInfo: {}, // 分组枚举值详情信息
            curProgress: ""
        };
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        // curProgress() {
        //     return this.$route.params.id;
        // },
        progressTree() {
            return this.$store.state.progressTree;
        },
        userInfo() {
            return this.$store.getters.userInfo;
        }
    },
    watch: {
        progressTree: {
            handler(arr) {
                if (
                    this.progressTree &&
                    Object.keys(this.progressTree).length
                ) {
                    // 获取查看权限
                    this.fetchAuthView();
                }
            },
            deep: true,
            immediate: true
        },
        $route: {
            // 流程切换
            handler(CurId, PreId) {
                if (
                    this.progressTree &&
                    Object.keys(this.progressTree).length
                ) {
                    // 获取查看权限
                    this.fetchAuthView();
                }
            }
        }
    },
    created() {},
    methods: {
        fetchAuthView() {
            this.curProgress = this.$route.params.id;
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: "all",
                auth_mode: "see"
            };
            api.getUserAuth(params).then((res) => {
                if (this.$route.name !== "progress") return;
                if (res && res.resultCode === 200) {
                    this.noProgressAuth = !res.data;
                    if (!this.noProgressAuth) {
                        // 有权限才重置
                        // 重置配置
                        this.tableCol = [];
                        this.tableData = [];
                        this.tableLoading = true;
                        this.table.page_num = 1;
                        this.currentTab = "";
                        this.viewFilter = [];
                        this.viewFilterStandby = [];
                        this.filterParams = [];
                        this.backupParams = []; // 备份filter参数，  前后不一致再调接口
                        this.expandBtnShow = false;
                        this.showMore = false;
                        this.filterRelation = "filter_and";
                        // 获取导出权限
                        this.fetchAuthExport();
                        // 获取新增权限
                        this.fetchAuthAdd();
                        this.getProgressInfo();
                        this.getView();
                        if (this.$refs.ProgressTable) {
                            this.$refs.ProgressTable.refreshTableKey();
                        }
                    }
                } else {
                    this.noProgressAuth = true;
                }
            });
        },
        fetchAuthExport() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: "all",
                auth_mode: "export"
            };
            api.getUserAuth(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.exportPrmission = res.data;
                } else {
                    this.exportPrmission = false;
                }
            });
        },
        fetchAuthAdd() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: "all",
                auth_mode: "create"
            };
            api.getUserAuth(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.createPrmission = res.data;
                } else {
                    this.createPrmission = false;
                }
            });
        },
        resursionTree(tree, cur = {}) {
            tree.forEach((sub) => {
                if (sub.id == Number(this.curProgress)) {
                    cur = sub;
                }
                if (sub.children && sub.children.length) {
                    sub.children.forEach((s) => {
                        if (s.id == Number(this.curProgress)) {
                            cur = s;
                        }
                    });
                }
            });
            return cur;
        },
        // 操作列复制连接
        copyLink(row) {
            this.$emit("copy-link", row);
        },
        // 操作列删除行
        deleteRow(row) {
            // 打开删除表格行的提醒弹窗
            this.$refs.deleteTableRowDialog.openDialog(row);
        },
        // 本地以流程为单位存储列宽
        changeColumnWidth(field, width) {
            // 读取当前流程的列宽缓存，  如果没有就创建一个新的，
            if (localStorage.getItem(`progress-${this.curProgress}`)) {
                let configArrJSON = localStorage.getItem(
                    `progress-${this.curProgress}`
                );
                let configArr = JSON.parse(configArrJSON);
                // 是创建还是编辑
                let targetConfig = _.find(configArr, { field: field });
                if (targetConfig) {
                    targetConfig.width = width;
                } else {
                    let obj = {
                        field,
                        width
                    };
                    configArr.push(obj);
                }
                localStorage.setItem(
                    `progress-${this.curProgress}`,
                    JSON.stringify(configArr)
                );
            } else {
                let configArr = [];
                let obj = {
                    field,
                    width
                };
                configArr.push(obj);
                localStorage.setItem(
                    `progress-${this.curProgress}`,
                    JSON.stringify(configArr)
                );
            }
        },
        getDescConfig() {
            if (this.noProgressAuth) return;
            // ProgressDesc DescContent
            let el = document.getElementById("DescContent");
            const childNodesLength = el.childNodes.length;
            // 使用 range 代替 scrollWidth 来判断文本是否溢出，这样做是为了解决潜在的 bug。
            const range = document.createRange();
            // 设置 range 的起点
            range.setStart(el, 0);
            // 设置 range 的终点，因为起终点都在同一个节点上，所以设置终点偏移量以选中节点的内容
            range.setEnd(el, childNodesLength);
            // 获取节点的内容的宽度
            const rangeWidth = range.getBoundingClientRect().width;
            // 获取el盒模型
            // 获取el左右pandding
            if (rangeWidth + 4 > el.offsetWidth) {
                // 文字超出显示按钮
                this.expandBtnShow = true;
            } else {
                this.expandBtnShow = false;
            }
        },
        openShareTask() {
            let row = {
                _id: this.$route.query._id
            };
            this.$refs["DetailDrawer"].openDrawer("tmpl_detail", row);
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
                    this.refreshCurrentPage();
                }
            });
        },
        openHtmlCol(row, col) {
            this.$refs["htmlColDialog"].openDialog(row, col);
        },
        getProgressInfo() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress
            };
            api.getProgressInfo(params).then((res) => {
                if (this.$route.name !== "progress") return;
                if (res && res.resultCode === 200 && res.data) {
                    this.progressInfo = res.data;
                    // 新建 导出权限
                    // if (
                    //     this.progressInfo.permission &&
                    //     this.progressInfo.permission.edit_permission
                    // ) {
                    //     let editPermission = [];
                    //     editPermission =
                    //         this.progressInfo.permission.edit_permission.split(
                    //             ","
                    //         );
                    //     this.createPrmission =
                    //         editPermission.indexOf("create") > -1;
                    //     this.exportPrmission =
                    //         editPermission.indexOf("export_list") > -1;
                    // } else {
                    //     this.createPrmission = false;
                    //     this.exportPrmission = false;
                    // }
                } else {
                    this.progressInfo = {
                        admin_list: [],
                        tmpl: {}
                    };
                    // this.createPrmission = false;
                    // this.exportPrmission = false;
                }
                this.$nextTick(() => {
                    this.getDescConfig();
                });
                // 检查路由中是否有指定任务_id,直接打开详情
                // if (this.$route.query && this.$route.query._id) {
                //     this.openShareTask();
                // }
            });
        },
        // 获取视图
        getView() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress
            };
            api.getView(params).then((res) => {
                if (this.$route.name !== "progress") return;
                if (res && res.resultCode === 200) {
                    this.tabsList = res.data || [];
                    // this.tabsList.unshift(this.homeView);
                    this.tabsList.forEach((item) => {
                        item.jsonId = JSON.stringify(item.id);
                    });
                    this.$nextTick(() => {
                        if (this.tabsList && this.tabsList instanceof Array) {
                            // 定位打开的视图
                            this.anchorPointLock();
                        }
                    });
                }
            });
        },
        // 移除视图
        deleteView() {
            this.currentTab = "";
            this.getView();
        },
        refreshView(order) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress
            };
            api.getView(params).then((res) => {
                if (this.$route.name !== "progress") return;
                if (res && res.resultCode === 200) {
                    this.tabsList = res.data || [];
                    this.tabsList.forEach((item) => {
                        item.jsonId = JSON.stringify(item.id);
                    });
                    this.$nextTick(() => {
                        // order === add , 则定位到最后一个tab
                        if (order === "add") {
                            this.anchorLastTab();
                        }
                    });
                }
            });
        },
        getColConfig(columns) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                module: "list",
                columns: columns
            };
            return api.getProgressColConfig(params).then((res) => {
                if (this.$route.name !== "progress") return;
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    if (this.currentTabType === "list") {
                        let apiCol = [];
                        res.data.forEach((col) => {
                            let obj = {
                                ...col,
                                columnName: col.field_key
                            };
                            apiCol.push(obj);
                        });
                        let basicColCol = _.cloneDeep(
                            HandleData.progressBasicCol
                        );
                        this.tableCol = apiCol.concat(basicColCol);
                    }
                } else {
                    this.tableCol = [];
                }
            });
        },
        // 过滤条件变化
        filterSearch(params, viewFilter, filterRelation) {
            this.filterParams = params;
            this.viewFilterStandby = viewFilter;
            // 需要同时比较 新旧params 和  新旧 filterRelation
            if (
                JSON.stringify(this.backupParams) !== JSON.stringify(params) ||
                this.filterRelation !== filterRelation
            ) {
                // 判断是列表还是卡片
                if (this.currentTabType === "list") {
                    this.filterRelation = filterRelation;
                    this.table.page_num = 1;
                    this.getListData();
                    this.backupParams = _.cloneDeep(params);
                    this.justRefreshEnumInfoNum(this.groupByFieldInfo);
                }
            }
        },
        // 排序条件变化
        sortSearch(sortOrder) {
            this.sortOrder = sortOrder;
            // 查询
            if (this.currentTabType === "list") {
                this.table.page_num = 1;
                this.getListData();
            }
            if (this.currentTabType === "card") {
                this.$refs.KanBan.getAllStatusCardsList();
            }
        },
        // 显示隐藏逻辑变化
        checkChange(columns) {
            // 按columns请求新列头
            let paramsColumns = columns.length ? JSON.stringify(columns) : "";
            if (this.currentTabType === "list") {
                this.getColConfig(paramsColumns);
            }
            if (this.currentTabType === "card") {
                this.$refs.KanBan.getColConfig(paramsColumns);
            }
        },
        getListData: _.debounce(function () {
            if (this.noProgressAuth) return;
            let filterDown = {
                x_axis: "", // 状态
                x_value: "",
                group_axis:
                    this.groupByFieldEnumInfo.filed_mode === "all"
                        ? ""
                        : this.groupByFieldInfo.field_key || "", // 维度
                group_value:
                    this.groupByFieldEnumInfo.filed_mode === "all"
                        ? ""
                        : parseInt(this.groupByFieldEnumInfo.user_id) ||
                          parseInt(this.groupByFieldEnumInfo.status_id) ||
                          this.groupByFieldEnumInfo.name ||
                          "" // 维度值  角色取user_id
            };
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                filter: JSON.stringify(this.filterParams),
                sort_order: JSON.stringify(this.sortOrder),
                lor: this.filterRelation,
                page_num: this.table.page_num,
                page_size: 20,
                filter_down: JSON.stringify(filterDown)
            };
            api.getProgressListData(params).then((res) => {
                if (this.$route.name !== "progress") return;
                if (res && res.resultCode === 200 && res.data) {
                    this.tableData = res.data.list || [];
                    this.table.count = res.data.cnt || 0;
                } else {
                    this.tableData = [];
                    this.table.count = 0;
                }
                this.tableLoading = false;
            });
        }),
        // 更新列表
        refreshPage: _.debounce(function () {
            if (this.currentTabType === "list") {
                this.table.page_num = 1;
                // 添加事件  不用下传
                this.getListData();
            }
        }),
        refreshCurrentPage() {
            if (this.groupByFieldInfo) {
                // 刷新当前页。并刷新枚举数量
                this.justRefreshEnumInfoNum(this.groupByFieldInfo);
            } else {
                this.getListData();
            }
        },
        refreshForAdd() {
            if (this.currentTabType === "list") {
                // 刷新枚举信息与列表
                this.table.page_num = 1;
                if (this.groupByFieldInfo) {
                    this.refreshEnumInfo(); // 是否有分组信息等 调刷新接口
                } else {
                    this.getListData();
                }
            }
            if (this.currentTabType === "card") {
                this.$refs.KanBan.refreshForAdd();
            }
        },
        // 刷新枚举信息，新增 删除  编辑 时调用
        refreshEnumInfo() {
            if (this.groupByFieldInfo) {
                this.confimGroupBy(this.groupByFieldInfo);
            }
        },
        gePersonList() {
            let params = {
                ws_id: "",
                tmpl_id: 29,
                page_size: 10,
                page: 1
            };
            api.gePersonList(params).then((res) => {});
        },
        // 定位打开的视图
        anchorPointLock() {
            this.viewFilter = [];
            this.filterRelation = "filter_and";
            if (this.tabsList.length) {
                this.tabsList.forEach((item) => {
                    if (item.pin === "yes") {
                        this.currentTab = item.jsonId;
                        this.viewId = item.id;
                        this.viewMode = item.mode;
                    }
                });
                if (!this.currentTab) {
                    // 进入主列表
                    this.currentTab = this.tabsList[0].jsonId;
                    this.viewId = this.tabsList[0].id;
                    this.viewMode = this.tabsList[0].mode;
                    this.viewFilterStandby = [];
                    this.filterParams = [];
                    this.backupParams = [];
                    this.filterRelation = "filter_and";
                }
            }
        },
        // 定位新建的最新视图
        anchorLastTab() {
            this.viewFilter = [];
            if (this.tabsList.length) {
                let lastTabRank = this.tabsList.length - 1;
                this.currentTab = this.tabsList[lastTabRank].jsonId;
                this.viewId = this.tabsList[lastTabRank].id;
                this.viewMode = this.tabsList[lastTabRank].mode;
            }
        },
        // tab切换
        tabChange(tab) {
            if (!tab) {
                // 首次进入视图tab为undefind
                return;
            }
            this.tableLoading = true;
            this.groupByFieldInfo = {};
            this.currentTab = tab.jsonId;
            this.currentTabType = tab.mode;
            this.curTabItems = _.cloneDeep(tab);
            // board的信息组件内查询
            if (tab.mode === "board") {
                return;
            }
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: tab.id
            };
            // 表格和卡片的条件配置
            api.getViewInfo(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    this.currentTab == tab.id
                ) {
                    let viewInfo = res.data;
                    viewInfo.jsonId = JSON.stringify(viewInfo.id);
                    this.viewFilter = [];
                    this.currentTab = viewInfo.jsonId;
                    this.currentTabType = viewInfo.mode;
                    // 筛选
                    if (viewInfo.filter) {
                        this.saveFilterGroup = JSON.parse(viewInfo.filter);
                        this.filterParams = JSON.parse(viewInfo.filter);
                        this.backupParams = _.cloneDeep(
                            JSON.parse(viewInfo.filter)
                        );
                    } else {
                        // 切回主列表强制清空
                        this.filterParams = [];
                        this.backupParams = [];
                    }
                    this.filterRelation = viewInfo.lor || "filter_and";
                    if (viewInfo.format_filter) {
                        this.viewFilterStandby = JSON.parse(
                            viewInfo.format_filter
                        );
                    } else {
                        this.viewFilterStandby = [];
                    }
                    // 排序
                    if (viewInfo.sort_order) {
                        this.sortOrder = JSON.parse(viewInfo.sort_order);
                    } else {
                        this.sortOrder = [];
                    }
                    // 显示隐藏
                    if (viewInfo.columns) {
                        this.saveHiddenShowColumns = viewInfo.columns;
                        let hiddenShowColumnsList = JSON.parse(
                            viewInfo.columns
                        );
                        this.hiddenShowColumnsList = hiddenShowColumnsList;
                    } else {
                        this.saveHiddenShowColumns = "";
                        this.hiddenShowColumnsList = [];
                    }
                    // 分组信息
                    if (viewInfo.card_group) {
                        this.cardFilterDown = JSON.parse(viewInfo.card_group);
                    } else {
                        this.cardFilterDown = "";
                    }
                    // 请求表头， 要入参viewInfo.columns , 本身就是字符串
                    let paramsColumn = viewInfo.columns;
                    if (tab.mode === "card" || tab.mode === "list") {
                        let params = {
                            ws_id: this.curSpace.id,
                            tmpl_id: this.curProgress,
                            id: tab.id
                        };
                        // 卡片分组信息
                        api.getCardGroupByList(params).then((cardRes) => {
                            if (
                                cardRes &&
                                cardRes.resultCode === 200 &&
                                this.currentTab == tab.id
                            ) {
                                this.groupByOption = _.cloneDeep(
                                    cardRes.data || []
                                );
                            } else {
                                this.groupByOption = _.cloneDeep([]);
                            }
                            if (tab.mode === "list") {
                                this.getColConfig(paramsColumn);
                            }
                            if (tab.mode === "card") {
                                // 调配置类接口
                                // 手动移除状态项
                                _.remove(this.groupByOption, {
                                    field_key: "status"
                                });
                                this.$nextTick(() => {
                                    this.$refs.KanBan.getConfigTypeInfo(
                                        paramsColumn
                                    );
                                });
                            }
                        });
                    }
                }
            });
        },
        // 切换枚举值第一次请求
        checkFieldSearch: _.debounce(function (info) {
            this.groupByFieldEnumInfo = info;
            this.refreshPage();
        }, 200),
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
                    this.refreshCurrentPage();
                }
            });
        },
        multipleSelection(arr) {
            this.tableMultipleSelection = arr;
        },
        // 创建新视图
        createView(form) {
            let groupCard;
            if (this.groupByFieldInfo.id) {
                groupCard = {
                    group_axis: this.groupByFieldInfo.field_key,
                    group_value:
                        parseInt(this.groupByFieldEnumInfo.user_id) ||
                        parseInt(this.groupByFieldEnumInfo.status_id) ||
                        this.groupByFieldEnumInfo.name ||
                        "" // 维度值  角色取user_id
                };
            } else {
                // 无分组
                groupCard = "";
            }
            // 通过筛选创建视图
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                mode: "list",
                name: form.name,
                filter: this.saveFilterGroup.length
                    ? JSON.stringify(this.saveFilterGroup)
                    : "",
                lor: this.filterRelation,
                sort_order: this.sortOrder.length
                    ? JSON.stringify(this.sortOrder)
                    : "",
                columns: this.saveHiddenShowColumns,
                card_group: JSON.stringify(groupCard)
            };
            api.createView(params).then((res) => {
                if (res.resultCode === 200) {
                    if (this.$route.name !== "progress") return;
                    this.refreshView("add");
                }
            });
        },
        // 更新过滤
        updateViewFilter() {
            if (this.noProgressAuth) return;
            // 更新视图
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: this.currentTab, //视图id
                filter: this.saveFilterGroup.length
                    ? JSON.stringify(this.saveFilterGroup)
                    : "",
                lor: this.filterRelation
            };
            api.updateViewFilter(params).then((res) => {
                if (this.$route.name !== "progress") return;
                this.$refs.FilterBlock.closePopver();
            });
        },
        // 更新排序
        updateSortOrder() {
            if (this.noProgressAuth) return;
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: this.currentTab, //视图id
                sort_order: this.sortOrder.length
                    ? JSON.stringify(this.sortOrder)
                    : ""
            };
            api.updateSortOrder(params).then((res) => {
                if (this.$route.name !== "progress") return;
                this.$refs.SortBlock.closePopver();
            });
        },
        // 更新显示隐藏
        updateColumnHieddenShow() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: this.currentTab, //视图id
                columns: this.saveHiddenShowColumns
            };
            api.updateViewColumns(params).then((res) => {
                if (this.$route.name !== "progress") return;
                this.$refs.HiddenShowBlock.closePopver();
                let paramsColumn = this.saveHiddenShowColumns;
                this.getColConfig(paramsColumn);
            });
        },

        addTab(type) {
            let viewMap = {
                list: "新增列表",
                board: "仪表盘",
                card: "看板"
            };
            // tab新建
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                mode: type,
                name: viewMap[type],
                filter: "",
                lor: "filter_and",
                sort_order: "",
                columns: ""
            };
            if (type === "card") {
                params.card_group = {};
            }
            api.createView(params).then((res) => {
                if (res.resultCode === 200) {
                    if (this.$route.name !== "progress") return;
                    this.refreshView("add");
                }
            });
        },
        lockTab(tab) {
            this.tabsList.forEach((item) => {
                if (item.id === tab.id) {
                    this.$set(item, "lock", true);
                } else {
                    this.$set(item, "lock", false);
                }
            });
            this.getView();
        },
        unlockTab(tab) {
            // 解除锁定
            this.tabsList.forEach((item) => {
                if (item.id === tab.id) {
                    this.$set(item, "lock", false);
                }
            });
            this.getView();
        },
        openDetail(scope, type) {
            this.formData = scope.row;
            this.$refs["DetailDrawer"].openDrawer(type, scope.row);
        },
        addProgress() {
            if (!this.createPrmission) return;
            // 掉新增config接口，打开dialog
            this.getAddProgressConfig();
            this.$refs.AddProgress.openDialog();
        },
        handleOpreate(command) {
            // 导出表格
            if (command === "export") {
                if (!this.exportPrmission) return;
                let params = {
                    ws_id: this.curSpace.id,
                    tmpl_id: this.curProgress,
                    filter: JSON.stringify(this.filterParams),
                    sort_order: JSON.stringify(this.sortOrder),
                    lor: this.filterRelation,
                    export: "yes"
                };
                api.exportProgress(params);
            }
        },
        getAddProgressConfig() {
            if (this.noProgressAuth) return;
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
                }
            });
        },
        getViewFilter() {
            this.viewFilter = _.cloneDeep(this.viewFilterStandby);
        },
        // 保存过滤
        saveFilter(obj, filterRelation, saveOrder) {
            this.saveFilterGroup = obj;
            this.filterRelation = filterRelation;
            // 新视图打开dialog确认明细 否则直接更新
            if (saveOrder) {
                if (this.currentTab === "-1") {
                    this.$refs.SaveFilter.openDialog();
                } else {
                    this.updateViewFilter();
                }
            }
        },
        // 保存排序
        saveSort(obj) {
            this.sortOrder = obj;
            // 新视图打开dialog确认明细 否则直接更新
            if (this.currentTab === "-1") {
                this.$refs.SaveFilter.openDialog();
            } else {
                this.updateSortOrder();
            }
        },
        // 保存显示隐藏
        saveHiddenShow(columns) {
            this.saveHiddenShowColumns = JSON.stringify(columns);
            if (this.currentTab === "-1") {
                this.$refs.SaveFilter.openDialog();
            } else {
                // 更新视图显示隐藏列
                this.updateColumnHieddenShow();
            }
        },
        // 获取字段分组信息
        confimGroupBy(groupInfo, order) {
            if (this.currentTabType === "list") {
                this.groupByFieldInfo = groupInfo;
                // 获取分组字段信息
                // 如果是无分组 需要直接调卡片列表 否则调枚举值列表
                if (groupInfo.id) {
                    let param = {
                        filter: this.filterParams.length
                            ? JSON.stringify(this.filterParams)
                            : "",
                        lor: "filter_and",
                        mode: "base_histogram",
                        order: "desc",
                        title: "title",
                        tmpl_id: parseInt(this.curProgress),
                        userid: this.userInfo.id,
                        view_id: parseInt(this.currentTab),
                        ws_id: this.curSpace.id,
                        show_empty: "yes", //枚举值中返回空值
                        xMode: groupInfo.mode, // 分组字段类型
                        x_axis: groupInfo.field_key, // 分组字段
                        y_axis: "count" //  写死
                    };
                    let selectAll = {
                        filed_mode: "all",
                        name: "全部"
                    };
                    dashApi.getPrviewChart(param).then((res) => {
                        if (res.resultCode === 200) {
                            this.groupByList = res.data.board_statistics;
                            this.groupByList.unshift(selectAll);
                        } else {
                            this.groupByList = [];
                            this.groupByList.unshift(selectAll);
                        }
                    });
                } else {
                    // 无分组
                    // 枚举字段为空
                    this.groupByList = [];
                    this.checkFieldSearch(DataHandle.noGroupInfo);
                }
            }
            if (this.currentTabType === "card") {
                this.$nextTick(() => {
                    if (!order) {
                        this.cardFilterDown = "";
                    }
                    this.$refs.KanBan.confimGroupBy(groupInfo, order);
                });
            }
        },
        // 只刷新数量
        justRefreshEnumInfoNum(groupInfo) {
            if (this.currentTabType === "list") {
                this.groupByFieldInfo = groupInfo;
                this.getListData();
                // 获取分组字段信息
                // 如果是无分组 需要直接调卡片列表 否则调枚举值列表
                if (groupInfo.id) {
                    let param = {
                        filter: this.filterParams.length
                            ? JSON.stringify(this.filterParams)
                            : "",
                        lor: "filter_and",
                        mode: "base_histogram",
                        order: "desc",
                        title: "title",
                        tmpl_id: parseInt(this.curProgress),
                        userid: this.userInfo.id,
                        view_id: parseInt(this.currentTab),
                        ws_id: this.curSpace.id,
                        show_empty: "yes", //枚举值中返回空值
                        xMode: groupInfo.mode, // 分组字段类型
                        x_axis: groupInfo.field_key, // 分组字段
                        y_axis: "count" //  写死
                    };
                    let selectAll = {
                        filed_mode: "all",
                        name: "全部"
                    };
                    dashApi.getPrviewChart(param).then((res) => {
                        let arr = [];
                        // 刷新枚举，但不能
                        if (res.resultCode === 200) {
                            arr = res.data.board_statistics;
                            arr.unshift(selectAll);
                        } else {
                            arr = [];
                            arr.unshift(selectAll);
                        }
                        // 判断当前枚举值是否还有效果
                        this.$refs.GroupByTemp.justRefreshEnumInfoNum(arr);
                    });
                }
            }
        },
        groupBySaveView() {
            if (this.currentTabType === "list") {
                if (this.currentTab === "-1") {
                    this.$refs.SaveFilter.openDialog();
                } else {
                    // 更新视图
                    this.saveGroupView();
                }
            }
            if (this.currentTabType === "card") {
                this.$refs.KanBan.saveGroupView();
            }
        },
        saveGroupView() {
            let groupCard;
            if (this.groupByFieldInfo.id) {
                groupCard = {
                    group_axis: this.groupByFieldInfo.field_key,
                    group_value:
                        parseInt(this.groupByFieldEnumInfo.user_id) ||
                        parseInt(this.groupByFieldEnumInfo.status_id) ||
                        this.groupByFieldEnumInfo.name ||
                        "" // 维度值  角色取user_id
                };
            } else {
                // 无分组
                groupCard = "";
            }
            // 更新卡片分组
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: this.currentTab, //视图id
                card_group: JSON.stringify(groupCard)
            };
            api.updateViewCardGroup(params);
        },
        currentChange(page) {
            this.tableLoading = true;
            this.table.page_num = page;
            this.getListData();
        },
        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        }
    }
};
</script>

<style lang="scss" scoped>
.table-loading {
    min-height: calc(100vh - 400px);
    .gourp-by-with-table {
        &.had-group-by-temp {
            display: flex;
            // min-height: calc(100vh - 220px);
            min-height: 600px;
        }
    }
    .group-by-list {
        width: 200px;
        ~ .table-content {
            width: calc(100% - 200px);
        }
        &.close {
            width: 40px;
            ~ .table-content {
                width: calc(100% - 40px);
            }
        }
    }
}
.title-administrators {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    .title {
        margin-right: 12px;
        font-weight: 500;
        font-size: 20px;
        color: #171e31;
        overflow: hidden;
        text-overflow: ellipsis;
    }
    .administrators {
        display: flex;
        align-items: start;
        margin-top: 4px;
        font-weight: 400;
        font-size: 12px;
        color: #40485f;
        .member-label {
            display: inline-block;
            width: 74px;
            margin-right: 6px;
        }
        .member-value {
            display: inline-block;
            flex: 1;
        }
        .admin-tag {
            margin-right: 8px;
        }
    }
}
.operation-block {
    display: flex;
    justify-content: space-between;
    margin-bottom: 12px;
    .operation-left {
        line-height: 32px;
    }
    .operation-right {
        display: flex;
        align-items: center;
        .add-btn {
            height: 28px !important;
            line-height: 28px !important;
        }
        b.add {
            display: inline-block;
            width: 16px;
            height: 16px;
            margin-right: 2px;
            vertical-align: text-bottom;
            background-image: url(@/assets/image/common/add_white.png);
            background-size: 100% 100%;
        }
        .split-line {
            font-size: 12px;
            color: #bfc1c6;
            margin: 0 12px 0 16px;
        }
        .progress-more-operate-icon {
            display: inline-block;
            width: 18px;
            height: 18px;
            cursor: pointer;
            vertical-align: middle;
            background-size: 100% 100%;
            background-image: url(@/assets/image/common/progress_more_operate.svg);
            &.actived,
            &.hover {
                background-image: url(@/assets/image/common/progress_more_operate.svg);
            }
        }
    }
}

.progress-desc {
    margin-top: 6px;
    font-weight: 400;
    font-size: 12px;
    color: #a6abbc;
    line-height: 20px;
    margin-bottom: 12px;
    height: 20px;
    display: flex;
    &.single-col {
        .desc-content {
            &.single-col {
                text-overflow: ellipsis;
                white-space: nowrap;
                overflow: hidden;
            }
        }
    }
    &.show-more {
        height: auto;
    }
    .desc-content {
        flex: 1;
    }
    .expand-btn {
        cursor: pointer;
        width: 24px;
        padding-left: 12px;
        color: #1890ff;
        &.close {
            display: inline-block;
        }
    }
}
.operation-btn {
    display: inline-block;
    margin-right: 16px;
    vertical-align: middle;
}
.basic-ui.delete-btn {
    margin-right: 16px;
    vertical-align: bottom;
    .delete-box {
        display: inline-block;
        width: 18px;
        height: 18px;
        background-size: 100% 100%;
        background-image: url(@/assets/image/common/delete.svg);
        vertical-align: middle;
        position: relative;
        top: -2px;
        cursor: pointer;
    }
    &:hover {
        .delete-box {
            display: inline-block;
            width: 18px;
            height: 18px;
            background-size: 100% 100%;
            background-image: url(@/assets/image/common/delete-active.svg);
            vertical-align: middle;
            position: relative;
            top: -2px;
            cursor: pointer;
        }
    }
}
// 骨架图
.table-flex {
    display: flex;
}
.table-flex .table-flex_div {
    width: 30%;
    padding: 0 10px;
}
.btn-group ::v-deep {
    .el-button--small {
        height: 32px;
        line-height: 32px;
        padding: 0 15px;
        padding: 0;
        font-size: 14px;
    }
    .el-dropdown__caret-button {
        padding: 0 5px;
    }
}
</style>
<style lang="scss">
// 新增右侧-流程更多操作
// .el-popover.el-popper.progress-more-operate-poppover[x-placement^="bottom"] {
//     margin-top: 2px;
// }
// .el-popover.el-popper.progress-more-operate-poppover[x-placement^="top"] {
//     margin-bottom: 2px;
// }
.el-popover.el-popper.progress-more-operate-poppover {
    min-width: 20px;
    padding: 8px;
    border-radius: 4px;
    background-color: #fff;
    box-shadow: 2px 2px 8px 1px rgba(47, 56, 76, 0.3);
    .pop-item {
        display: flex;
        align-items: center;
        height: 32px;
        padding: 0 8px;
        border-radius: 4px;
        cursor: pointer;
        &.disabled {
            opacity: 0.5;
            cursor: not-allowed;
            img {
                cursor: not-allowed;
            }
            &:hover {
                background-color: #fff;
                opacity: 0.5;
            }
        }
        &.his {
            width: 106px;
        }
        img {
            margin-right: 4px;
        }
        &:hover {
            background-color: #ecf5ff;
        }
    }
}
.el-popover.el-popper.progress-operate-col[x-placement^="bottom"] {
    margin-top: 2px;
}
.el-popover.el-popper.progress-operate-col[x-placement^="top"] {
    margin-bottom: 2px;
}
.el-popover.el-popper.progress-operate-col {
    min-width: 20px;
    padding: 8px;
    border-radius: 4px;
    background-color: #fff;
    box-shadow: 2px 2px 8px 1px rgba(47, 56, 76, 0.3);
    .pop-item {
        display: flex;
        align-items: center;
        height: 32px;
        padding: 0 8px;
        border-radius: 4px;
        cursor: pointer;
        &.disabled {
            opacity: 0.5;
            cursor: not-allowed;
            img {
                cursor: not-allowed;
            }
            &:hover {
                background-color: #fff;
                opacity: 0.5;
            }
        }
        &.his {
            width: 106px;
        }
        img {
            margin-right: 4px;
        }
        &:hover {
            background-color: #ecf5ff;
        }
    }
}
</style>
