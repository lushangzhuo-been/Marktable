<template>
    <div>
        <!-- 标题 -->
        <div class="title-block">
            <b class="customer-rule-icon"></b>
            自定义规则
        </div>
        <!-- tabs -->
        <div class="tabs-block">
            <el-tabs
                class="basic-ui add-customer-tabs"
                v-model="activeName"
                @tab-click="tabChange"
            >
                <el-tab-pane
                    v-for="(item, index) in tabsList"
                    :label="item.label"
                    :name="item.name"
                    :key="index"
                >
                    <span slot="label">
                        {{ item.label }}
                    </span>
                </el-tab-pane>
            </el-tabs>
        </div>
        <div v-if="activeName === '规则列表'" class="filter-block">
            <div class="filter-list">
                <filter-trigger
                    v-model="rule_type_select"
                    class="filter-item"
                ></filter-trigger>
                <filter-action v-if="false" class="filter-item"></filter-action>
                <filter-enable
                    v-model="is_enable_select"
                    class="filter-item"
                ></filter-enable>
                <FilterCreator
                    v-model="created_by_select"
                    class="filter-item"
                ></FilterCreator>
            </div>
            <!-- add rule -->
            <div class="add-rule">
                <el-button type="primary" size="small" @click="addRule">
                    +添加规则
                </el-button>
            </div>
        </div>
        <!-- table -->
        <n-table
            class="table-block"
            :col="tableCol"
            :data="tableData"
            @interactive="interactive"
        >
            <!-- 创建人 -->
            <!-- SinglePersonColumn -->
            <template slot="creator" slot-scope="{ column }">
                <single-person-column :column="column"></single-person-column>
            </template>
            <template slot="is_enable" slot-scope="{ column }">
                <enable-column
                    :column="column"
                    @enable-change="enableChange"
                ></enable-column>
            </template>
            <!-- 操作列 -->
            <template slot="operation" slot-scope="{ column }">
                <operation-column
                    :column="column"
                    @list-item-operation="listItemOperation"
                ></operation-column>
            </template>
            <!-- 触发类型 -->
            <template slot="rule_type" slot-scope="{ column }">
                <RuleTypeColumn :column="column"></RuleTypeColumn>
            </template>
            <!-- 执行结果 -->
            <template slot="status" slot-scope="{ column }">
                <status-column :column="column"></status-column>
            </template>
            <!-- enable_status -->
            <template slot="enable_status" slot-scope="{ column }">
                <enable-status-column :column="column"></enable-status-column>
            </template>
        </n-table>
        <el-pagination
            v-show="table.count > 20"
            class="basic-ui"
            @current-change="currentChange"
            layout="total, prev, pager, next"
            :current-page="table.pageNum"
            :page-size="table.pageSize"
            :total="table.count"
        >
        </el-pagination>
        <add-rule-dialog
            ref="AddRuleDialog"
            @refresh="refreshList"
        ></add-rule-dialog>
        <log-detail-dialog ref="LogDetailDialog"></log-detail-dialog>
    </div>
</template>

<script>
import { baseMixin } from "@/mixin.js";
import api from "@/common/api/module/progress_rule";
import DataHandle from "./data_handle";
import NTable from "@/pages/common/tables/common_table/table";
import AddRuleDialog from "@/pages/progress/customer_rule/components/add_rule_dialog";
import LogDetailDialog from "@/pages/progress/customer_rule/components/log_detail/log_detail_dialog";
import EnableColumn from "@/pages/progress/customer_rule/components/enable_column";
import SinglePersonColumn from "@/pages/progress/customer_rule/components/single_person_column";
import OperationColumn from "@/pages/progress/customer_rule/components/operation_column";
import EnableStatusColumn from "@/pages/progress/customer_rule/components/enable_status_column";
import StatusColumn from "@/pages/progress/customer_rule/components/status_column";
import RuleTypeColumn from "@/pages/progress/customer_rule/components/rule_type_column";
import FilterTrigger from "./components/filter_trigger";
import FilterAction from "./components/filter_action";
import FilterEnable from "./components/filter_enable";
import FilterCreator from "./components/filter_creator";
export default {
    mixins: [baseMixin],
    components: {
        NTable,
        AddRuleDialog,
        LogDetailDialog,
        EnableColumn,
        SinglePersonColumn,
        OperationColumn,
        FilterTrigger,
        FilterAction,
        FilterEnable,
        FilterCreator,
        EnableStatusColumn,
        StatusColumn,
        RuleTypeColumn
    },
    data() {
        return {
            table: {
                count: 120,
                pageSize: 20,
                pageNum: 1
            },
            activeName: "规则列表",
            tabsList: [
                {
                    name: "规则列表",
                    label: "规则列表"
                },
                {
                    name: "执行日志",
                    label: "执行日志"
                }
            ],
            tableCol: DataHandle.ruleListCol,
            tableData: [],
            rule_type_select: "",
            is_enable_select: "",
            created_by_select: ""
        };
    },
    watch: {
        curProgress: {
            handler(id) {
                this.activeName = "规则列表";
                this.tableCol = DataHandle.ruleListCol;
                this.table.pageNum = 1;
                this.getRuleList();
            }
        },
        rule_type_select: {
            handler() {
                this.table.pageNum = 1;
                this.getRuleList();
            }
        },
        is_enable_select: {
            handler() {
                this.table.pageNum = 1;
                this.getRuleList();
            }
        },
        created_by_select: {
            handler() {
                this.table.pageNum = 1;
                this.getRuleList();
            }
        }
    },
    mounted() {
        this.tabChange();
    },
    methods: {
        interactive(row, col) {
            this.$refs.LogDetailDialog.openDialog(row);
        },
        tabChange() {
            this.tableData = [];
            this.table.pageNum = 1;
            if (this.activeName === "规则列表") {
                this.tableCol = DataHandle.ruleListCol;
                this.table.pageNum = 1;
                this.getRuleList();
            }
            if (this.activeName === "执行日志") {
                this.tableCol = DataHandle.logListCol;
                let param = {
                    // 可以入顶部筛选的参数
                };
                this.getLogList(param);
            }
        },
        enableChange(val, col, row) {
            let params = {
                id: row.id,
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                is_enable: val ? "yes" : "no"
            };
            api.ruleEnable(params).then((res) => {
                if (res.resultCode === 200) {
                    this.getRuleList();
                }
            });
        },
        listItemOperation(col, row, type) {
            let params = {
                id: row.id,
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress
            };
            if (type === "edit" || type === "copy") {
                api.getRuleDetail(params).then((res) => {
                    if (res.resultCode === 200) {
                        let detail = res.data;
                        this.$refs.AddRuleDialog.openDialog(type, detail);
                    }
                });
            }
            if (type === "delete") {
                api.ruleDelete(params).then((res) => {
                    if (res.resultCode === 200) {
                        this.table.pageNum = 1;
                        this.getRuleList();
                    }
                });
            }
            if (type === "log") {
                this.activeName = "执行日志";
                this.tableCol = DataHandle.logListCol;
                this.table.pageNum = 1;
                let param = {
                    rule_id: row.id
                };
                this.getLogList(param);
            }
        },
        refreshList() {
            this.table.pageNum = 1;
            this.getRuleList();
        },
        getLogList(param) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                page_size: this.table.pageSize,
                page_num: this.table.pageNum,
                ...param
            };
            api.getLogList(params).then((res) => {
                if (res.resultCode === 200 && res.data.list) {
                    this.tableData = res.data.list;
                    this.table.count = res.data.cnt;
                } else {
                    this.tableData = [];
                    this.table.count = 0;
                }
            });
        },
        getRuleList() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                page_size: this.table.pageSize,
                page_num: this.table.pageNum,
                rule_type_select: this.rule_type_select,
                is_enable_select: this.is_enable_select,
                created_by_select: this.created_by_select
            };
            api.getRuleList(params).then((res) => {
                if (res.resultCode === 200 && res.data.list) {
                    this.tableData = res.data.list;
                    this.table.count = res.data.cnt;
                } else {
                    this.tableData = [];
                    this.table.count = 0;
                }
            });
        },
        currentChange(page) {
            this.table.pageNum = page;
            if (this.activeName === "规则列表") {
                this.getRuleList();
            }
            if (this.activeName === "执行日志") {
                let param = {
                    // 可以入顶部筛选的参数
                };
                this.getLogList(param);
            }
        },
        addRule() {
            this.$refs.AddRuleDialog.openDialog("add");
        }
    }
};
</script>

<style lang="scss" scoped>
::v-deep.add-customer-tabs.el-tabs {
    .el-tabs__item {
        font-weight: 600;
        font-size: 16px;
        color: #82889e;
        &.is-active,
        &:hover {
            color: #409eff;
        }
    }
}
.title-block {
    height: 24px;
    line-height: 24px;
    font-weight: 500;
    font-size: 20px;
    color: #171e31;
    .customer-rule-icon {
        display: inline-block;
        width: 24px;
        height: 24px;
        background-image: url(@/assets/image/common/customer_rule.svg);
        background-size: 100% 100%;
        vertical-align: text-bottom;
    }
}
.tabs-block {
    margin-top: 16px;
}
.filter-block {
    display: flex;
    justify-content: space-between;
    .filter-list {
        .filter-item {
            display: inline-block;
            margin: 6px 16px 6px 0;
        }
    }
    .add-rule {
    }
}
.table-block {
    margin-top: 16px;
}
</style>
