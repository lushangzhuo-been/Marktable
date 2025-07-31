<template>
    <el-dialog
        title="添加角色"
        :visible.sync="dialogVisible"
        class="add-rember-dialog basic-ui"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        @close="cancel"
    >
        <el-form
            ref="addRoleForm"
            class="basic-ui"
            :model="addRoleForm"
            :rules="rule"
            label-width="84px"
            size="small"
        >
            <el-form-item prop="name" label="名称:">
                <el-input
                    class="basic-ui"
                    v-model="addRoleForm.name"
                    type="text"
                    size="small"
                    placeholder="请输入角色名称"
                ></el-input>
            </el-form-item>
            <el-form-item prop="desc" label="说明:">
                <el-input
                    class="basic-ui"
                    v-model="addRoleForm.desc"
                    type="textarea"
                    :autosize="{ minRows: 2, maxRows: 6 }"
                    size="small"
                    placeholder="请输入角色说明"
                ></el-input>
            </el-form-item>
            <el-form-item prop="view_permission" label="浏览权限:">
                <el-radio-group
                    class="basic-ui view-permission"
                    v-model="addRoleForm.view_permission"
                >
                    <el-radio label="view_all">能浏览所有条目</el-radio>
                    <el-radio label="view_only_me"
                        >仅能浏览自己创建的或者指派给自己的条目</el-radio
                    >
                </el-radio-group>
            </el-form-item>
            <el-form-item prop="edit_permission" label="编辑权限:">
                <div class="edit-permit-desc">
                    默认权限,用户只能操作自己有关的条目(条目的创建人,当前处理人),删除权限只有创建人可以操作
                </div>
                <!-- <el-radio-group
                    class="basic-ui right"
                    v-model="editRoleForm.edit_permission"
                >
                    <el-radio label="edit_all">能编辑能浏览到的条目</el-radio>
                    <el-radio label="edit_only_progress">仅能更新进展</el-radio>
                    <el-radio label="edit_none">无编辑权限</el-radio>
                </el-radio-group> -->
                <el-checkbox-group
                    class="basic-ui edit-permission"
                    v-model="addRoleForm.edit_permission"
                >
                    <el-checkbox label="create">创建</el-checkbox>
                    <el-checkbox label="export_list">导出列表</el-checkbox>
                </el-checkbox-group>
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
export default {
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        }
    },
    data() {
        return {
            dialogVisible: false,
            addRoleForm: {
                name: "",
                desc: "",
                view_permission: "view_only_me",
                edit_permission: []
            },

            rule: {
                name: [
                    {
                        required: true,
                        message: "请输入角色名称",
                        trigger: "blur"
                    }
                ],
                desc: [
                    {
                        required: true,
                        message: "请输入角色说明",
                        trigger: "blur"
                    }
                ]
            }
        };
    },
    methods: {
        openDialog() {
            this.dialogVisible = true;
            this.$nextTick(() => {
                this.$refs.addRoleForm.clearValidate();
            });
        },
        cancel() {
            this.$refs.addRoleForm.clearValidate();
            this.addRoleForm = {
                name: "",
                desc: "",
                view_permission: "view_only_me",
                edit_permission: []
            };
            this.dialogVisible = false;
        },
        onConfirm() {
            this.$refs.addRoleForm.validate((flag) => {
                let tmpForm = _.cloneDeep(this.addRoleForm);
                for (let key in tmpForm) {
                    if (key === "edit_permission") {
                        tmpForm["edit_permission"] =
                            tmpForm["edit_permission"].join(",");
                    }
                }
                if (flag) {
                    let params = {
                        ws_id: this.curSpace.id,
                        tmpl_id: this.$route.params.id,
                        ...tmpForm
                    };
                    api.createRole(params).then((res) => {
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

<style lang="scss" scoped>
.add-rember-dialog {
    .width100 {
        width: 100%;
    }
    .view-permission.el-radio-group .el-radio,
    .edit-permission.el-checkbox-group .el-checkbox {
        display: block;
        margin-right: 0;
        line-height: 2;
    }

    .edit-permit-desc {
        padding: 8px 14px;
        margin-bottom: 12px;
        line-height: 18px;
        font-size: 12px;
        color: #82889e;
        background-color: #f8f8f8;
    }
}
</style>
