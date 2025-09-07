<template>
    <el-dialog
        :title="title"
        :visible.sync="dialogVisible"
        class="basic-ui remove-folder-progress-dialog"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        @close="cancel"
    >
        <div slot="title" class="dialog-title">
            <i class="el-icon-warning"></i>
            <span>{{ title }}</span>
        </div>
        <el-form
            ref="removeFolderProgress"
            class="basic-ui"
            :model="removeFolderProgress"
            :rules="rule"
            label-width="52px"
        >
            <el-form-item prop="name" label="名称">
                <el-input
                    class="basic-ui"
                    v-model="removeFolderProgress.name"
                    type="text"
                    size="small"
                    :placeholder="placeholder"
                ></el-input>
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
export default {
    data() {
        return {
            title: "",
            placeholder: "请输入名称",
            dialogVisible: false,
            removeFolderProgress: {
                name: ""
            },
            objData: {},
            rule: {
                name: [
                    {
                        required: true,
                        message: "",
                        trigger: "blur"
                    }
                ]
            }
        };
    },
    methods: {
        openDialog(obj) {
            if (!obj || !obj.data) return;
            this.objData = obj;
            this.$nextTick(() => {
                this.$refs.removeFolderProgress.clearValidate();
            });
            this.title =
                obj.data.impl_table_name === "ws_file"
                    ? "删除文件夹"
                    : "删除流程";
            this.placeholder =
                obj.data.impl_table_name === "ws_file"
                    ? "请输入要删除的文件夹名称"
                    : "请输入要删除的流程名称";
            this.dialogVisible = true;
            this.rule = {
                name: [
                    {
                        required: true,
                        message: this.placeholder,
                        trigger: "blur"
                    },
                    {
                        message: "",
                        trigger: "blur",
                        validator: (rule, value, callback) => {
                            if (
                                this.removeFolderProgress.name !==
                                this.objData.data.name
                            ) {
                                callback(
                                    new Error(
                                        this.objData.data &&
                                        this.objData.data.impl_table_name ===
                                            "ws_file"
                                            ? "请输入正确的文件夹名称"
                                            : "请输入正确的流程名称"
                                    )
                                );
                            } else {
                                callback();
                            }
                        }
                    }
                ]
            };
        },
        cancel() {
            this.removeFolderProgress = {
                name: ""
            };
            this.dialogVisible = false;
        },
        onConfirm() {
            this.$refs.removeFolderProgress.validate((flag) => {
                if (flag) {
                    this.$emit(
                        "on-confirm",
                        this.objData,
                        this.removeFolderProgress
                    );
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.el-form.basic-ui {
    margin: 6px 0;
}
</style>
