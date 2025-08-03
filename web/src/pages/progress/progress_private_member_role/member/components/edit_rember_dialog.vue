<template>
    <el-dialog
        title="编辑"
        :visible.sync="dialogVisible"
        class="edit-rember-dialog basic-ui"
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
            <!-- <el-form-item prop="name" label="名称:">
                <el-select
                    v-model="privateRemberForm.name"
                    clearable
                    size="small"
                    multiple
                    filterable
                    popper-class="private-member-name"
                    class="basic-ui width100"
                    placeholder="请选择/输入名称"
                >
                    <el-checkbox-group v-model="privateRemberForm.name">
                        <el-option
                            v-for="item in nameOptions"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        >
                            <span class="format-role-option">
                                <span>
                                    <img
                                        :src="item.icon"
                                        alt=""
                                        width="20px"
                                        height="20px"
                                    />
                                    <span>{{ item.label }}</span>
                                </span>
                                <el-checkbox
                                    @click.native.prevent
                                    :label="item.label"
                                ></el-checkbox>
                            </span>
                        </el-option>
                    </el-checkbox-group>
                </el-select>
            </el-form-item> -->
            <el-form-item prop="role" label="角色:">
                <el-select
                    v-model="privateRemberForm.role"
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
import { PRIMARY_COLOR_LIST } from "@/assets/tool/color_list";
export default {
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        progressRoleList() {
            let list = this.$store.state.progressRoleList;
            if (list && list.length) {
                return list.map((item) => {
                    return {
                        label: item.name,
                        value: item.id
                    };
                });
            }
            return [];
        }
    },
    data() {
        return {
            PRIMARY_COLOR: PRIMARY_COLOR_LIST.PRIMARY_COLOR,
            dialogVisible: false,
            privateRemberForm: {
                role: "",
                status: ""
            },
            roleOptions: [
                {
                    label: "管理员",
                    value: "admin"
                },
                {
                    label: "普通",
                    value: "general"
                }
            ],
            nameOptions: [],
            rule: {
                number: [
                    {
                        required: true,
                        message: "请输入账号",
                        trigger: "blur"
                    }
                ],
                name: [
                    {
                        required: true,
                        message: "请输入名称",
                        trigger: "blur"
                    }
                ],
                role: [
                    {
                        required: true,
                        message: "请选择角色",
                        trigger: "change"
                    }
                ]
            }
        };
    },
    methods: {
        switchChange() {},
        openDialog(row) {
            this.curRow = _.cloneDeep(row);
            for (let key in this.privateRemberForm) {
                // if (key === 'status') {
                //     this.privateRemberForm[key] =
                //         row[key] === 'ok' ? true : false
                // }
                if (key === "role") {
                    this.privateRemberForm[key] = row["role_id"] || "";
                } else {
                    this.privateRemberForm[key] = row[key] || "";
                }
            }
            this.dialogVisible = true;
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
        cancel() {
            this.dialogVisible = false;
            this.privateRemberForm = { role: "", status: "" };
            this.$nextTick(() => {
                this.$refs.privateRemberForm.clearValidate();
            });
        },
        onConfirm() {
            this.$refs.privateRemberForm.validate((flag) => {
                if (flag) {
                    const tmpForm = _.cloneDeep(this.privateRemberForm);
                    // tmpForm.status = this.privateRemberForm.status
                    //     ? 'ok'
                    //     : 'forbid'
                    let params = {
                        ws_id: this.curSpace.id,
                        tmpl_id: this.$route.params.id,
                        id: this.curRow.id,
                        status: tmpForm.status,
                        role_id: tmpForm.role
                    };
                    api.updateRember(params).then((res) => {
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
.edit-rember-dialog {
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
</style>
<style lang="scss" scoped>
.edit-rember-dialog {
    .width100 {
        width: 100%;
    }
}
</style>
