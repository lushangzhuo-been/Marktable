<template>
    <el-dialog
        v-if="dialogVisible"
        title="只读规则"
        :visible.sync="dialogVisible"
        class="only-read-dialog"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :destroy-on-close="true"
        append-to-body
        @close="cancel"
    >
        <div class="only-read-block">
            <div class="block-title">只读状态:</div>
            <div class="block-input">
                <mul-select
                    ref="MulSelect"
                    :option="statusOption"
                    :statusList="statusList"
                    @checked-arr="confirmCheck"
                ></mul-select>
            </div>
        </div>
        <div class="could-edit-block">
            <div>
                <!-- 用户&角色 -->
                <AddUserAndRole
                    ref="AddUserAndRole"
                    :rowInfo="rowInfo"
                    :userList="userList"
                    :userInfoList="userInfoList"
                    :roleCheckList="roleList"
                    @userlist-change="userlistChange"
                    @rolelist-change="rolelistChange"
                ></AddUserAndRole>
            </div>
        </div>
        <span slot="footer" class="dialog-footer">
            <el-button size="small" @click="cancel">取 消</el-button>
            <el-button
                class="basic-ui"
                size="small"
                type="primary"
                @click="onConfirm"
                >确 定</el-button
            >
            <!-- 确 定按钮的 :disabled="!couldConfirm" -->
        </span>
    </el-dialog>
</template>

<script>
import _ from "lodash";
import MulSelect from "./mul_select.vue";
import AddUserAndRole from "@/pages/progress/progress_setting/filed/common/add_user_and_role";
import api from "@/common/api/module/progress_setting";
export default {
    components: {
        MulSelect,
        AddUserAndRole
    },
    computed: {
        getLabelWidth() {
            return this.editFiledForm.mode === "person_com" ? "72px" : "60px";
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        }
    },
    data() {
        return {
            dialogVisible: false,
            statusOption: [],
            statusSelect: [], // 已选状态
            peopleOption: [],
            rowInfo: {},
            readOnlyRule: {},
            statusList: [],
            userList: [],
            userInfoList: [],
            roleList: [],
            couldConfirm: false,
            userListParams: [],
            roleListParams: [],
            ids: []
        };
    },
    watch: {
        dialogVisible: {
            handler(Boolean) {
                if (!Boolean) {
                    this.statusList = [];
                    this.userList = [];
                    this.userInfoList = [];
                    this.roleList = [];
                    this.couldConfirm = false;
                }
            }
        }
    },
    methods: {
        userlistChange(arr) {
            this.userListParams = arr || [];
        },
        rolelistChange(arr) {
            this.roleListParams = arr || [];
        },
        confirmCheck(arr) {
            this.ids = arr;
            if (arr.length) {
                this.couldConfirm = true;
            } else {
                this.couldConfirm = false;
            }
        },
        cancel() {
            this.dialogVisible = false;
        },
        onConfirm() {
            // let userList = this.$refs.AddUserAndRole.selectArr;
            // let issueRoleList = this.$refs.AddUserAndRole.checkArr;
            let params = {
                ws_id: this.rowInfo.ws_id,
                tmpl_id: this.rowInfo.tmpl_id,
                id: this.rowInfo.id,
                status_list: this.ids.length ? JSON.stringify(this.ids) : "",
                user_list: this.userListParams.length
                    ? JSON.stringify(this.userListParams)
                    : "",
                issue_role_list: this.roleListParams.length
                    ? JSON.stringify(this.roleListParams)
                    : ""
            };
            api.updateRuleRoleList(params).then((res) => {
                if (res.resultCode === 200) {
                    this.dialogVisible = false;
                    this.$emit("refresh-table");
                }
            });
        },
        fetchFieldInfo() {
            let params = {
                ws_id: this.rowInfo.ws_id,
                tmpl_id: this.rowInfo.tmpl_id,
                id: this.rowInfo.id
            };
            api.getReadOnlyRule(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.statusList = res.data.status_list_format || [];
                    this.userList = res.data.user_list || [];
                    this.userInfoList = res.data.user_list_format || [];
                    this.roleList = res.data.issue_role_list || [];
                }
            });
        },
        openDialog(row) {
            this.rowInfo = row;
            this.dialogVisible = true;
            this.getRuleStatusList();
        },
        getRuleStatusList() {
            let params = {
                ws_id: this.rowInfo.ws_id,
                tmpl_id: this.rowInfo.tmpl_id
            };
            api.getRuleStatusList(params).then((res) => {
                if (res.resultCode === 200) {
                    this.statusOption = res.data || [];
                    this.fetchFieldInfo();
                } else {
                    this.statusOption = [];
                }
            });
        }
    }
};
</script>
<style lang="scss">
.only-read-block {
    border-radius: 4px;
    background: #fafafb;
    padding: 16px 20px;
    .block-title {
        margin-bottom: 12px;
    }
}
.could-edit-block {
    margin-top: 20px;
    padding: 20px 0;
    border-top: 1px dotted #e6e9f0;
}
.only-read-dialog {
    .el-dialog {
        width: 520px;
        border-radius: 4px;
        .el-dialog__header {
            padding: 24px 32px;
            font-size: 16px;
            color: #2f384c;
        }
        .el-dialog__body {
            padding: 0 32px;
        }
        .el-dialog__footer {
            padding: 0 32px 32px;
        }
        .el-dialog__headerbtn {
            top: 24px;
            right: 32px;
        }
        .el-form-item {
            margin-bottom: 24px;
        }
        .el-dialog__footer {
            text-align: right;
        }
    }
}
</style>
