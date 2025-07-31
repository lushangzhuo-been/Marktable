<template>
    <el-dialog
        title="编辑成员"
        :visible.sync="dialogVisible"
        class="edit-rember-dialog basic-ui"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        @close="cancel"
    >
        <el-form
            ref="spaceRemberForm"
            :model="spaceRemberForm"
            :rules="rule"
            label-width="60px"
            size="small"
            @close="cancel"
        >
            <el-form-item prop="role" label="角色:">
                <el-select
                    v-model="spaceRemberForm.role"
                    class="basic-ui width100"
                    placeholder="请选择角色"
                >
                    <el-option
                        v-for="item in roleOptions"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    >
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item prop="status" label="状态:">
                <!-- <el-switch
                    v-model="spaceRemberForm.status"
                    :active-color="PRIMARY_COLOR"
                    inactive-color="#E6E9F0"
                    size="small"
                >
                </el-switch> -->
                <el-radio-group
                    class="basic-ui right"
                    v-model="spaceRemberForm.status"
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
import api from "@/common/api/module/space";
import { PRIMARY_COLOR_LIST } from "@/assets/tool/color_list";
export default {
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        }
    },
    data() {
        return {
            PRIMARY_COLOR: PRIMARY_COLOR_LIST.PRIMARY_COLOR,
            dialogVisible: false,
            spaceRemberForm: {
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
            rule: {
                role: [
                    {
                        required: true,
                        message: "请选择角色",
                        trigger: "change"
                    }
                ]
            },
            mapping: {
                admin: "管理员",
                general: "普通"
            }
        };
    },
    methods: {
        openDialog(row) {
            this.tmpForm = _.cloneDeep(row);
            for (let key in this.tmpForm) {
                this.tmpForm[key] = row[key] || "";
            }
            this.spaceRemberForm = _.cloneDeep(this.tmpForm);
            this.dialogVisible = true;
        },
        cancel() {
            this.spaceRemberForm = {
                role: "",
                status: ""
            };
            this.$refs.spaceRemberForm.clearValidate();
            this.dialogVisible = false;
        },
        onConfirm() {
            this.$refs.spaceRemberForm.validate((flag) => {
                if (flag) {
                    let params = {
                        ws_id: this.curSpace.id,
                        id: this.spaceRemberForm.id,
                        status: this.spaceRemberForm.status,
                        role: this.spaceRemberForm.role
                    };
                    api.updateWorkspaceMember(params).then((res) => {
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
