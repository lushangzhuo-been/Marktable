<template>
    <div style="width: calc(100% - 60px)">
        <div
            class="preview-box-left"
            :class="currentTab === 'table' ? 'active' : ''"
            style="margin-bottom: 16px"
            @click="clickLeft('table')"
        >
            <div style="margin-bottom: 16px; font-size: 12px">
                <!-- {{ `${tab === "card" ? "卡片" : "列表"}效果` }} -->
                列表效果
            </div>
            <div class="create-page-index">
                <p-table
                    ref="ProgressTable"
                    :data="tableData"
                    :col="tableCol"
                ></p-table>
            </div>
            <!-- <div class="create-page-index"> -->
            <!-- v-show="tab === 'card'" -->
            <!-- <card-view></card-view> -->
            <!-- <table-view v-show="tab !== 'card'"></table-view> -->
            <!-- v-show="tab !== 'card'" -->
            <!-- </div> -->
        </div>
        <div
            class="preview-box-left"
            :class="currentTab === 'card' ? 'active' : ''"
            @click="clickLeft('card')"
        >
            <div style="margin-bottom: 16px; font-size: 12px">
                <!-- {{ `${tab === "card" ? "卡片" : "列表"}效果` }} -->
                卡片效果
            </div>
            <!-- <div class="create-page-index" style="margin-bottom: 16px">
                <p-table
                    ref="ProgressTable"
                    :data="tableData"
                    :col="tableCol"
                ></p-table>
            </div> -->
            <div class="create-page-index">
                <!-- v-show="tab === 'card'" -->
                <card-view></card-view>
                <!-- <table-view v-show="tab !== 'card'"></table-view> -->
                <!-- v-show="tab !== 'card'" -->
            </div>
        </div>
    </div>
</template>
<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import { mapState } from "vuex";
import PTable from "@/components/table/p_table_preview/progress_table";
import CardView from "./components/card.vue";

const operateCol = [
    {
        prop: "operation",
        label: "操作",
        fixed: "right",
        align: "center",
        mode: "operation",
        columnName: "operation"
    }
];

export default {
    props: {
        tab: {
            type: String,
            default: "table"
        }
    },
    components: {
        // TableView,
        PTable,
        CardView
    },
    data() {
        return {
            tableData: [],
            tableCol: [],
            currentTab: "table"
        };
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        }
    },
    // watch: {
    //     tab: {
    //         handler() {
    //             this.fetchData();
    //         },
    //         immediate: true
    //     }
    // },
    mounted() {
        this.$emit("tab-change", this.currentTab);
    },
    methods: {
        clickLeft(tab) {
            this.currentTab = tab;
            this.$emit("tab-change", this.currentTab);
        },
        fetchData() {
            if (this.tab) {
                this.tab === "card" ? this.fetchCardData() : this.getViewInfo();
            }
        },
        fetchCardData() {},
        fetchTableCol() {},
        fetchTableData() {},
        getViewInfo() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: -1
            };
            api.getViewInfo(params).then((res) => {
                if (res.resultCode === 200) {
                    let viewInfo = res.data;
                    // 请求表头， 要入参viewInfo.columns , 本身就是字符串
                    let paramsColumn = viewInfo.columns;
                    this.getColConfig(paramsColumn);
                    this.fetchTableData(paramsColumn);
                    // 请求数据 第一次手动调列表数据,  之后每次filter变动时自动调(仅当切换到列表时调用)
                }
            });
        },
        getColConfig(columns) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                // view_id: this.viewMode === "board" ? -1 : this.viewId,
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
                    let apiCol = [];
                    res.data.forEach((col) => {
                        let obj = {
                            ...col,
                            columnName: col.field_key
                        };
                        apiCol.push(obj);
                    });
                    // this.tableCol = [];
                    // let backupCol = _.cloneDeep(operateCol);
                    // this.tableCol = res.data.concat(backupCol);
                    this.dataCol = _.cloneDeep(apiCol);
                    this.tableCol = apiCol.concat(operateCol);
                } else {
                    this.tableCol = [];
                    this.dataCol = [];
                }
                this.tableData = this.getTableData();
            });
        },
        getTableData() {
            if (!this.dataCol || !this.dataCol.length) return [];
            let obj = {};
            this.dataCol.forEach((item) => {
                if (item.mode === "text_com" && item.field_key === "title") {
                    obj[item.field_key] = "文档详情编辑";
                }
                if (item.mode === "number_com") {
                    obj[item.field_key] = "4567";
                }
                if (
                    item.mode === "select_com" ||
                    item.mode === "multi_select_com"
                ) {
                    obj[item.field_key] = ["样式问题"];
                }
                if (item.mode === "person_com") {
                    obj[item.field_key] = [];
                    obj[item.field_key][0] = {
                        avatar: "",
                        email: "360826018@qq.com",
                        full_name: "张成",
                        id: 44,
                        phone: "18501951295",
                        username: "张成"
                    };
                }
                if (item.mode === "date_com") {
                    obj[item.field_key] = "2025-03-30";
                }
                if (item.mode === "time_com") {
                    obj[item.field_key] = "2025-03-30 19:53:26";
                }
                if (item.mode === "status_com") {
                    obj[item.field_key] = {
                        id: 276,
                        ws_id: 80,
                        tmpl_id: 262,
                        name: "未开始",
                        color: "rgb(24, 144, 255)",
                        mode: "first"
                    };
                }
            });
            return [obj];
        }
    }
};
</script>
<style lang="scss" scoped>
.preview-box-left {
    box-sizing: border-box;
    max-height: calc(100% - 100px);
    min-height: 200px;
    width: 100%;
    padding: 24px;
    background: #ffffff;
    box-shadow: 0px 3px 10px 1px rgba(0, 0, 0, 0.1);
    border-radius: 4px 4px 4px 4px;
    &.active {
        border: 1px dashed var(--PRIMARY_COLOR);
    }
    &:hover {
        border: 1px dashed var(--PRIMARY_COLOR);
    }
}
</style>
