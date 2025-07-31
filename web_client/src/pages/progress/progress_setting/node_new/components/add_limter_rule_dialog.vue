<template>
    <el-dialog
        v-if="dialogVisible"
        title="添加规则"
        :visible.sync="dialogVisible"
        class="add-limiter-dialog"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :destroy-on-close="true"
        append-to-body
        @close="cancel"
    >
        <div>
            <span>描述：</span>
            <!-- <el-input v-model="name" placeholder="请输入描述"></el-input> -->
            <el-input
                class="basic-ui name-input"
                size="small"
                v-model="name"
                type="text"
                placeholder="请输入描述"
            ></el-input>
        </div>
        <div class="add-filter-block">
            <!-- <div class="block-title">编辑条件:
                <span>添加条件</span>
            </div> -->
            <!-- 条件 -->
            <!-- <LimterFilter @filter-params="filterParams"></LimterFilter> -->
        </div>
        <div class="could-edit-block">
            <div>
                <!-- 用户&角色 -->
                <AddUserAndRole
                    ref="AddUserAndRole"
                    :rowInfo="btnInfo"
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
                :disabled="confirmDisabled"
                @click="onConfirm"
                >确 定</el-button
            >
        </span>
    </el-dialog>
</template>

<script>
import _ from "lodash";
import AddUserAndRole from "@/pages/progress/progress_setting/filed/common/add_user_and_role";
import api from "@/common/api/module/progress_setting";
import LimterFilter from "./limit_filter.vue";
import { baseMixin } from "@/mixin.js";
export default {
    mixins: [baseMixin],
    components: {
        AddUserAndRole,
        LimterFilter
    },
    props: {
        rowInfo: {
            type: Object,
            default: () => {}
        },
        btnInfo: {
            type: Object,
            default: () => {}
        },
        curId: {
            type: [String, Number],
            default: ""
        }
    },
    data() {
        return {
            type: "",
            name: "",
            dialogVisible: false,
            readOnlyRule: {},
            userList: [],
            userInfoList: [],
            roleList: [],
            couldConfirm: false,
            limitInfo: {},
            userListParams: [],
            roleListParams: []
        };
    },
    computed: {
        confirmDisabled() {
            if (
                this.name &&
                (this.userListParams.length || this.roleListParams.length)
            ) {
                return false;
            } else {
                return true;
            }
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        }
    },
    watch: {
        dialogVisible: {
            handler(Boolean) {
                if (!Boolean) {
                    this.name = "";
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
        filterParams(params, view, relation) {},
        confirmCheck(arr) {
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
            let params = {
                name: this.name,
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                step_id: this.curId,
                mode: "permission", //子任务 权限
                user_list: this.userListParams.length
                    ? JSON.stringify(this.userListParams)
                    : "",
                issue_role_list: this.roleListParams.length
                    ? JSON.stringify(this.roleListParams)
                    : ""
            };
            if (this.type === "add") {
                api.createLimit(params).then((res) => {
                    if (res.resultCode === 200 && res.data) {
                        this.$emit("refresh-list");
                        this.dialogVisible = false;
                    }
                });
            } else {
                params.id = this.limitInfo.id;
                api.updateLimitInfo(params).then((res) => {
                    if (res.resultCode === 200 && res.data) {
                        this.$emit("refresh-list");
                        this.dialogVisible = false;
                    }
                });
            }
        },
        openDialog(row, type) {
            this.dialogVisible = true;
            if (type === "edit") {
                this.type = "edit";
                this.limitInfo = row;
                let params = {
                    ws_id: this.curSpace.id,
                    tmpl_id: this.$route.params.id,
                    step_id: this.curId,
                    id: row.id
                };
                api.getLimitInfo(params).then((res) => {
                    if (res.resultCode === 200 && res.data) {
                        this.name = res.data.name;
                        this.userList = res.data.user_list || [];
                        this.userInfoList = res.data.user_list_format || [];
                        this.roleList = res.data.issue_role_list || [];
                    }
                });
            } else {
                this.type = "add";
            }
        }
    }
};
</script>
<style lang="scss" scoped>
.basic-ui.name-input {
    width: calc(100% - 44px);
}
.could-edit-block {
    padding: 20px 0;
}
.add-limiter-dialog {
    .el-dialog {
        width: 660px;
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
