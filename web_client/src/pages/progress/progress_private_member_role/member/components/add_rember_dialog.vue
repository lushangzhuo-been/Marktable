<template>
    <el-dialog
        title="添加"
        :visible.sync="dialogVisible"
        class="add-progress-rember-dialog basic-ui"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        @close="cancel"
    >
        <el-form
            ref="privateRemberForm"
            class="basic-ui"
            :model="privateRemberForm"
            :rules="rule"
            label-width="60px"
            size="small"
        >
            <el-form-item prop="userid_list" label="名称:">
                <el-select
                    v-model="privateRemberForm.userid_list"
                    clearable
                    size="small"
                    multiple
                    remote
                    filterable
                    popper-class="private-member-userid_list"
                    class="basic-ui width100"
                    placeholder="请选择/输入名称"
                    :remote-method="userNameRemote"
                    @focus="userNameBlur"
                >
                    <el-checkbox-group v-model="privateRemberForm.userid_list">
                        <el-option
                            v-for="item in nameOptions"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        >
                            <span class="format-role-option">
                                <span>
                                    <el-avatar
                                        icon="el-icon-user-solid"
                                        size="small"
                                        :src="getAvatar(item.icon)"
                                    ></el-avatar>
                                    <span>{{ item.label }}</span>
                                </span>
                                <el-checkbox
                                    @click.native.prevent
                                    :label="item.value"
                                ></el-checkbox>
                            </span>
                        </el-option>
                    </el-checkbox-group>
                </el-select>
            </el-form-item>
            <el-form-item prop="role_id" label="角色:">
                <el-select
                    v-model="privateRemberForm.role_id"
                    filterable
                    size="small"
                    class="basic-ui width100"
                    placeholder="请选择角色"
                    :filter-method="roleFilter"
                >
                    <el-option
                        v-for="item in progressRoleList"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    >
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item prop="status" label="状态:">
                <!-- <el-switch
                    v-model="privateRemberForm.status"
                    :active-color="PRIMARY_COLOR"
                    inactive-color="#E6E9F0"
                    size="small"
                >
                </el-switch> -->
                <el-radio-group
                    class="basic-ui right"
                    v-model="privateRemberForm.status"
                >
                    <el-radio label="ok">开启</el-radio>
                    <el-radio label="forbid">禁用</el-radio>
                </el-radio-group>
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
            <el-button class="basic-ui" size="small" @click="cancel"
                >取 消</el-button
            >
            <el-button
                class="basic-ui"
                size="small"
                type="primary"
                @click="onConfirm"
                >确 定</el-button
            >
        </span>
    </el-dialog>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress_setting";
import api_common from "@/common/api";
import { PRIMARY_COLOR_LIST } from "@/assets/tool/color_list";
import { imgHost } from "@/assets/tool/const";
export default {
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        getAvatar() {
            return (avatar) => {
                if (avatar) {
                    return `${imgHost}${avatar}`;
                }
                return require(`@/assets/image/common/default_avatar_big.png`);
            };
        }
    },
    data() {
        return {
            PRIMARY_COLOR: PRIMARY_COLOR_LIST.PRIMARY_COLOR,
            dialogVisible: false,
            privateRemberForm: {
                userid_list: [],
                role_id: "",
                status: "ok"
            },
            progressRoleList: [],

            nameOptions: [],
            rule: {
                userid_list: [
                    {
                        required: true,
                        message: "请输入名称",
                        trigger: "blur"
                    }
                ],
                role_id: [
                    {
                        required: true,
                        message: "请选择角色",
                        trigger: "change"
                    }
                ],
                email: [
                    {
                        required: true,
                        message: "请选择邮箱",
                        trigger: "change"
                    }
                ]
            }
        };
    },
    methods: {
        userNameRemote(val) {
            this.fetchAllUserList(val);
        },
        userNameBlur() {
            if (!this.nameOptions.length) {
                this.fetchAllUserList();
            }
        },
        fetchAllUserList(val = "") {
            let params = {
                key: val,
                page_size: 10,
                page: 1,
                ws_id: this.curSpace.id
            };
            api_common.getCurProgressMemberList(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.nameOptions = this.nameListFormat(res.data.list || []);
                } else {
                    this.nameOptions = [];
                }
            });
        },
        nameListFormat(arr) {
            if (arr && arr.length) {
                return arr.map((item) => {
                    return {
                        label: item.username,
                        value: item.id,
                        text: item.full_name,
                        icon: item.avatar
                    };
                });
            }
            return [];
        },
        roleFilter(val) {
            this.getRoleList(val);
        },
        getRoleList(val = "") {
            let params = {
                key: val,
                page_size: 10,
                page: 1,
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id
            };
            api.getRoleList(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.progressRoleList = this.getFormat(res.data.list || []);
                } else {
                    this.progressRoleList = [];
                }
            });
        },
        getFormat(arr) {
            if (arr && arr.length) {
                return arr.map((item) => {
                    return {
                        label: item.name,
                        value: item.id,
                        text: item.name
                    };
                });
            }
        },
        openDialog() {
            this.dialogVisible = true;
            this.fetchAllUserList();
            this.getRoleList();
            this.$nextTick(() => {
                this.$refs.privateRemberForm.clearValidate();
            });
        },
        cancel() {
            this.privateRemberForm = {
                userid_list: [],
                role_id: "",
                status: "ok"
            };
            this.$refs.privateRemberForm.clearValidate();
            this.dialogVisible = false;
        },
        onConfirm() {
            this.$refs.privateRemberForm.validate((flag) => {
                if (flag) {
                    const tmpForm = _.cloneDeep(this.privateRemberForm);
                    tmpForm.userid_list =
                        this.privateRemberForm.userid_list.join(",");
                    // tmpForm.status = this.privateRemberForm.status
                    //     ? 'ok'
                    //     : 'forbid'
                    let params = {
                        ws_id: this.curSpace.id,
                        tmpl_id: this.$route.params.id,
                        ...tmpForm
                    };
                    api.addRember(params).then((res) => {
                        if (res && res.resultCode === 200) {
                            this.$emit("on-confirm");
                        }
                    });
                }
            });
        }
    }
};
</script>

<style lang="scss">
.add-progress-rember-dialog {
    .el-dialog {
        width: 490px;
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
        // .el-icon-close:before {
        //     content: url('~@/assets/image/common/close.png');
        // }
        .el-form-item {
            margin-bottom: 24px;
        }
        .el-dialog__footer {
            text-align: right;
        }
    }
}
.private-member-userid_list {
    .el-checkbox-group .el-checkbox__label {
        display: none;
    }
}
.private-member-userid_list.el-select-dropdown.el-popper
    .el-scrollbar
    .el-select-dropdown__item.selected {
    color: var(--FONT_COLOR);
    font-weight: 400;
}
.format-role-option {
    display: flex;
    justify-content: space-between;
    align-items: center;
    .el-avatar {
        vertical-align: middle;
        margin-right: 10px;
        height: 20px;
        width: 20px;
        line-height: 20px;
        border-radius: 50%;
    }
}
</style>
<style lang="scss" scoped>
.add-progress-rember-dialog {
    .width100 {
        width: 100%;
    }
}
</style>
