<template>
    <div class="limiter-content">
        <!-- 添加按钮 -->
        <div class="add-rule-block">
            <el-button
                size="small"
                class="basic-ui"
                type="primary"
                @click="addLimter"
                >+添加规则</el-button
            >
        </div>
        <!-- 表格 -->
        <n-table class="table-content" :col="limiterCol" :data="limiterData">
            <template slot="operation" slot-scope="{ column }">
                <el-table-column
                    show-overflow-tooltip
                    :width="200"
                    :label="column.name"
                >
                    <template slot-scope="scope">
                        <div>
                            <el-button
                                type="text"
                                size="small"
                                @click="openLimitEdit(scope.row)"
                            >
                                编辑
                            </el-button>
                            <el-button
                                type="text"
                                size="small"
                                @click="deleteLimit(scope.row)"
                            >
                                删除
                            </el-button>
                        </div>
                    </template>
                </el-table-column>
            </template>
        </n-table>
        <AddLimterRuleDialog
            ref="AddLimterRuleDialog"
            :rowInfo="rowInfo"
            :btnInfo="btnInfo"
            @refresh-list="refreshList"
        ></AddLimterRuleDialog>
    </div>
</template>

<script>
import _ from "lodash";
import NewSelectItem from "@/pages/progress/filter/filter_item/new_select";
import PersonItem from "@/pages/progress/filter/filter_item/person";
import MulPeople from "@/pages/progress/progress_setting/filed/common/mul_people";
import AddUserAndRole from "@/pages/progress/progress_setting/filed/common/add_user_and_role";
import AddLimterRuleDialog from "./add_limter_rule_dialog.vue";
import NTable from "@/pages/common/tables/common_table/table";
import HandleData from "./handle_data";
import { baseMixin } from "@/mixin.js";
import api from "@/common/api/module/progress_setting";
export default {
    mixins: [baseMixin],
    components: {
        NewSelectItem,
        PersonItem,
        MulPeople,
        AddUserAndRole,
        NTable,
        AddLimterRuleDialog,
    },
    props: {
        rowInfo: {
            type: Object,
            default: () => {},
        },
        btnInfo: {
            type: Object,
            default: () => {},
        },
    },
    data() {
        return {
            userList: [],
            userInfoList: [],
            roleList: [],
            limiterCol: HandleData.limiterCol,
            limiterData: [
                {
                    type: "子任务",
                    desc: "一份描述",
                },
            ],
        };
    },
    watch: {},
    mounted() {
        this.getLimitList();
    },
    methods: {
        refreshList() {
            this.getLimitList();
        },
        openLimitEdit(row) {
            this.$refs.AddLimterRuleDialog.openDialog(row, "edit");
        },
        deleteLimit(row) {
            let params = {
                ws_id: row.ws_id,
                tmpl_id: row.tmpl_id,
                node_id: row.node_id,
                button_id: row.button_id,
                id: row.id,
            };
            api.deleteLimitInfo(params).then((res) => {
                if (res.resultCode === 200 && res.data) {
                    this.getLimitList();
                }
            });
        },
        getLimitList() {
            let params = {
                ws_id: this.btnInfo.ws_id,
                tmpl_id: this.btnInfo.tmpl_id,
                node_id: this.btnInfo.node_id,
                button_id: this.btnInfo.id,
            };
            api.getLimitList(params).then((res) => {
                if (res.resultCode === 200 && res.data) {
                    this.limiterData = res.data;
                } else {
                    this.limiterData = [];
                }
            });
        },
        addLimter() {
            this.$refs.AddLimterRuleDialog.openDialog();
        },
    },
};
</script>

<style lang="scss" scoped>
.limiter-content {
    height: 100%;
    // position: relative;
    .add-rule-block {
        display: flex;
        justify-content: end;
        margin-bottom: 16px;
    }
    .footer-group {
        box-sizing: border-box;
        position: absolute;
        width: 100%;
        bottom: 0;
        left: 0;
        padding: 20px 32px;
        // background-color: aqua;
        display: flex;
        justify-content: end;
        box-shadow: 0px -3px 8px 1px rgba(0, 0, 0, 0.1);
    }
}
.could-edit-block {
    margin-top: 20px;
    padding: 20px 0;
    border-top: 1px dotted #e6e9f0;
    .could-edit-title {
        margin-bottom: 12px;
    }
    .third-person {
        margin-top: 10px;
        padding: 16px 20px;
        background: #fafafb;
        border-radius: 4px;
        .third-person-title {
            margin-bottom: 12px;
        }
    }
}
.basic-ui.el-button.add-condition {
    color: #1890ff;
    border: 1px solid #1890ff;
    &:hover {
        border: 1px solid #1890ff;
    }
}
.add-condition-icon {
    display: inline-block;
    width: 10px;
    height: 10px;
    background-size: 100% 100%;
    background-image: url(@/assets/image/progress/add-card-active.png);
}
.operation-box {
    display: inline-block;
    width: 18px;
    height: 18px;
    background-size: 100% 100%;
    vertical-align: text-bottom;
    &.add {
        background-image: url(@/assets/image/progress/add.png);
    }
}
.filter-list {
    padding-left: 72px;
    margin-top: 16px;
    // max-height: 540px;
    // overflow-y: auto;
    .filter-list-item {
        height: 32px;
        line-height: 32px;
        display: flex;
        margin-bottom: 12px;
        .delete-box {
            display: inline-block;
            width: 18px;
            height: 18px;
            background-size: 100% 100%;
            background-image: url(@/assets/image/progress/delete.png);
            vertical-align: middle;
            cursor: pointer;
            margin: 7px;
            &:hover {
                background-image: url(@/assets/image/progress/delete-active.png);
            }
        }
    }
}
</style>
