<template>
    <div class="progress-setting-filed">
        <!-- <el-alert :title="remindMessage" type="warning" show-icon> </el-alert> -->
        <div class="btns">
            <el-button
                size="small"
                class="basic-ui"
                type="primary"
                @click="addNode"
                >+新增步骤</el-button
            >
        </div>
        <n-table :data="tableData" :col="tableCol">
            <!-- 名称 -->
            <template slot="name" slot-scope="{ column }">
                <el-table-column
                    show-overflow-tooltip
                    :label="column.name"
                    :min-width="column.minWidth"
                >
                    <template slot-scope="scope">
                        {{ getColValues(scope.row, column) }}
                    </template>
                </el-table-column>
            </template>
            <!-- 类型 -->
            <template slot="mode" slot-scope="{ column }">
                <el-table-column
                    show-overflow-tooltip
                    :label="column.name"
                    :min-width="column.minWidth"
                >
                    <template slot-scope="scope">
                        {{ getColValues(scope.row, column) }}
                    </template>
                </el-table-column>
            </template>
            <!-- 状态 -->
            <template slot="status_text" slot-scope="{ column }">
                <el-table-column
                    :show-overflow-tooltip="false"
                    :label="column.name"
                    :min-width="column.minWidth"
                >
                    <template slot-scope="scope">
                        <div class="status-flex">
                            <b
                                class="circle"
                                :style="{
                                    backgroundColor: scope.row.node.status_color
                                }"
                                v-if="
                                    scope.row.node &&
                                    scope.row.node.status_color
                                "
                            ></b>
                            <tip-one :text="getColValues(scope.row, column)">
                                <span class="status-text">
                                    {{ getColValues(scope.row, column) }}
                                </span>
                            </tip-one>
                        </div>
                    </template>
                </el-table-column>
            </template>
            <!-- 编辑按钮 -->
            <template slot="buttons" slot-scope="{ column }">
                <el-table-column
                    show-overflow-tooltip
                    :min-width="column.minWidth"
                    :resizable="false"
                    :label="column.name"
                >
                    <template slot-scope="scope">
                        <template
                            v-if="scope.row.buttons && scope.row.buttons.length"
                        >
                            <!-- 做遍历 -->
                            <div
                                v-for="(item, index) in scope.row.buttons"
                                :key="index"
                                class="btns-item"
                            >
                                {{ `${index + 1}.` }}
                                <span
                                    class="name"
                                    @click="openConvertDetail(item, scope.row)"
                                    >{{ item.name || "--" }}</span
                                >
                                <span style="margin-left: 10px"
                                    >→
                                    <span style="margin: 0 24px 0 10px">{{
                                        item.target_node_name || "--"
                                    }}</span>
                                </span>
                                <el-button
                                    size="mini"
                                    class="basic-ui"
                                    @click="editConvertBtn(item, scope.row)"
                                    >编辑转换按钮</el-button
                                >
                                <el-button
                                    size="mini"
                                    class="basic-ui"
                                    @click="deleteConvertBtn(item, scope.row)"
                                    >删除转换按钮</el-button
                                >
                            </div>
                        </template>
                        <template v-else>
                            <span>--</span>
                        </template>
                    </template>
                </el-table-column>
            </template>
            <!-- 操作 -->
            <template slot="operation" slot-scope="{ column }">
                <el-table-column
                    show-overflow-tooltip
                    :width="column.width"
                    :label="column.name"
                >
                    <template slot-scope="scope">
                        <div>
                            <el-button
                                @click="addConvertBtn(column, scope.row)"
                                type="text"
                                size="small"
                            >
                                添加转换按钮
                            </el-button>
                            <el-button
                                @click="editNode(column, scope.row)"
                                type="text"
                                size="small"
                            >
                                编辑
                            </el-button>
                            <el-button
                                @click="deleteNode(column, scope.row)"
                                type="text"
                                size="small"
                            >
                                删除
                            </el-button>
                        </div>
                    </template>
                </el-table-column>
            </template>
        </n-table>
        <node-drawer ref="NodeDrawer"></node-drawer>
        <add-node-dialog
            ref="AddNodeDialog"
            @on-confirm="refreshNodeList('update-node')"
        ></add-node-dialog>
        <delete-node-dialog
            ref="DeleteNodeDialog"
            @on-confirm="refreshNodeList('delete-node')"
        ></delete-node-dialog>
        <add-convert-btn-dialog
            ref="AddConvertBtnDialog"
            @on-confirm="refreshNodeList('update-btn')"
        ></add-convert-btn-dialog>
        <delete-convert-btn-dialog
            ref="DeleteConvertBtnDialog"
            @on-confirm="refreshNodeList('delete-btn')"
        ></delete-convert-btn-dialog>
    </div>
</template>

<script>
import _ from "lodash";
import TEXT from "@/assets/tool/text";
import api from "@/common/api/module/progress_setting";
import NTable from "@/pages/common/tables/common_table/table";
import NodeDrawer from "./components/node_drawer";
import AddNodeDialog from "./components/node_add_dialog";
import DeleteNodeDialog from "./components/node_delete_dialog";
import AddConvertBtnDialog from "./components/convert_add_btn_dialog";
import DeleteConvertBtnDialog from "./components/convert_delete_btn_dialog";
import TipOne from "@/pages/common/tooltip_one_line.vue";

export default {
    components: {
        NTable,
        NodeDrawer,
        AddNodeDialog,
        DeleteNodeDialog,
        AddConvertBtnDialog,
        DeleteConvertBtnDialog,
        TipOne
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        getColValues() {
            return (row, column) => {
                let value = "--";
                if (row && row.node) {
                    if (column.prop === "mode") {
                        this.modeRadioList.forEach((item) => {
                            if (item.value === row.node[column.prop]) {
                                value = item.label;
                            }
                        });
                    } else {
                        value = row.node[column.prop];
                    }
                    return value;
                } else {
                    return value;
                }
            };
        }
    },
    data() {
        return {
            remindMessage: TEXT.PROGRESS_SETTING_NODE,
            modeRadioList: [],
            curIndex: "",
            tableCol: [
                {
                    prop: "name",
                    name: "步骤",
                    minWidth: 90,
                    type: "slot"
                },
                {
                    prop: "mode",
                    name: "类型",
                    minWidth: 90,
                    type: "slot"
                },
                {
                    prop: "status_text",
                    name: "状态",
                    minWidth: 125,
                    type: "slot"
                },
                {
                    prop: "buttons",
                    name: "转换按钮",
                    minWidth: 410,
                    type: "slot"
                },
                {
                    prop: "operation",
                    name: "操作",
                    type: "slot",
                    width: 170
                }
            ],
            tableData: [
                // {
                //     name: '未开始',
                //     type: '首节点',
                //     status: '待开始',
                //     buttons: [
                //         {
                //             name: '下发开始',
                //             targetNode: '进行中',
                //         },
                //         {
                //             name: '重置',
                //             targetNode: '待开始',
                //         },
                //     ],
                //     operation: '编辑',
                // },
                // {
                //     name: '进行中',
                //     type: '中间节点',
                //     status: '进行中',
                //     buttons: [
                //         {
                //             name: '开发完成',
                //             targetNode: '转测',
                //         },
                //     ],
                //     operation: '编辑',
                // },
            ]
        };
    },
    watch: {
        $route() {
            this.fetchNodeList();
        }
    },
    mounted() {
        this.fetchNodeConfig();
    },
    methods: {
        fetchNodeConfig() {
            api.getNodeConfig().then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.modeRadioList = _.cloneDeep(
                        res.data.mode_config || []
                    );
                    this.fetchNodeList();
                } else {
                    this.modeRadioList = [];
                }
            });
        },
        fetchNodeList() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id
            };
            api.getNodeList(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    res.data.length
                ) {
                    this.tableData = res.data;
                } else {
                    this.tableData = [];
                }
            });
        },
        refreshNodeList(type) {
            let mapping = {
                "update-node": "AddNodeDialog",
                "delete-node": "DeleteNodeDialog",
                "update-btn": "AddConvertBtnDialog",
                "delete-btn": "DeleteConvertBtnDialog"
            };
            this.$refs[mapping[type]].cancel();
            this.fetchNodeList();
        },
        addNode() {
            // 新增节点
            this.$refs.AddNodeDialog.openDialog("add", this.modeRadioList);
        },
        editNode(col, row) {
            // 编辑节点
            this.$refs.AddNodeDialog.openDialog(
                "edit",
                this.modeRadioList,
                row
            );
        },
        deleteNode(col, row) {
            // 删除节点
            this.$refs.DeleteNodeDialog.openDialog(row);
        },
        addConvertBtn(item, row) {
            // 添加转换按钮
            this.$refs.AddConvertBtnDialog.openDialog(
                item,
                row,
                "add",
                this.tableData
            );
        },
        editConvertBtn(item, row) {
            // 编辑转换按钮
            this.$refs.AddConvertBtnDialog.openDialog(
                item,
                row,
                "edit",
                this.tableData
            );
        },
        deleteConvertBtn(item, row) {
            // 删除转换按钮
            this.$refs.DeleteConvertBtnDialog.openDialog(item, row);
        },
        openConvertDetail(item, row) {
            // 打开状态转换按钮设置详情状态
            this.$refs.NodeDrawer.openDrawer(item, row);
        }
    }
};
</script>

<style lang="scss" scoped>
.progress-setting-filed {
    .el-alert {
        margin-bottom: 16px;
    }
    .btns {
        text-align: right;
        margin-bottom: 16px;
    }
    .circle {
        display: inline-block;
        width: 8px;
        height: 8px;
        margin-right: 8px;
        border-radius: 50%;
    }
    ::v-deep .el-table {
        th.el-table__cell > .cell {
            font-weight: 400;
        }
        th.el-table__cell {
            background-color: #f5f6f8;
        }
        .el-table__cell {
            box-sizing: border-box;
            padding: 0;
            height: 40px;
            line-height: 39px;
        }
    }
    ::v-deep .el-button--mini {
        height: 24px;
        line-height: 24px;
        padding: 0 8px;
    }
    ::v-deep .el-button + .basic-ui.el-button,
    ::v-deep .el-checkbox.is-bordered + .basic-ui.el-checkbox.is-bordered {
        margin-left: 8px;
    }
    .btns-item {
        height: 32px;
        line-height: 32px;
        display: flex;
        align-items: center;
        .name {
            color: var(--PRIMARY_COLOR);
            text-decoration: underline;
            cursor: pointer;
        }
    }
    .status-flex {
        display: flex;
        align-items: center;
        .status-text {
            display: inline-block;
            width: calc(100% - 16px);
            overflow: hidden;
            white-space: nowrap;
            text-overflow: ellipsis;
        }
    }
}
</style>
