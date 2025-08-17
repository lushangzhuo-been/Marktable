<template>
    <el-dialog
        title="添加文件夹"
        :visible.sync="dialogVisible"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        class="basic-ui add-folder-dialog"
        @close="cancel"
    >
        <el-form
            ref="newFolderForm"
            class="basic-ui"
            :model="newFolderForm"
            :rules="rule"
            label-width="52px"
        >
            <el-form-item prop="name" label="名称">
                <el-input
                    class="basic-ui"
                    v-model="newFolderForm.name"
                    type="text"
                    maxlength="25"
                    show-word-limit
                    size="small"
                    placeholder="请输入文件夹名称"
                ></el-input>
            </el-form-item>
            <el-form-item label="描述">
                <el-input
                    class="basic-ui fixed-height-dialog"
                    v-model="newFolderForm.desc"
                    type="textarea"
                    size="small"
                    placeholder="请输入文件夹描述"
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
import api from "@/common/api/module/progress_setting";
export default {
    data() {
        return {
            curSpace: {},
            dialogVisible: false,
            newFolderForm: {
                name: "",
                desc: ""
            },
            rule: {
                name: [
                    {
                        required: true,
                        message: "请输入文件夹名称",
                        trigger: "blur"
                    }
                ]
            }
        };
    },
    methods: {
        openDialog(obj) {
            this.curSpace = obj;
            this.dialogVisible = true;
        },
        cancel() {
            this.$refs.newFolderForm.clearValidate();
            this.newFolderForm = {
                name: "",
                desc: ""
            };
            this.dialogVisible = false;
        },
        onConfirm() {
            this.$refs.newFolderForm.validate((flag) => {
                let params = {
                    ws_id: this.curSpace.id,
                    ...this.newFolderForm
                };
                api.addFolder(params).then((res) => {
                    if (!res) return;
                    if (res.resultCode === 200) {
                        this.$emit("on-confirm");
                    } else if (
                        res.resultCode === 401 ||
                        res.resultCode === 403
                    ) {
                        this.dialogVisible = false;
                    }
                });
            });
        }
    }
};
</script>

<style lang="scss">
.add-folder-dialog {
    .el-dialog {
        border-radius: 4px;
        .el-dialog__header {
            padding: 24px 32px;
            font-size: 16px;
            color: #2f384c;
        }
        .el-dialog__body {
            padding: 0 32px;
        }
        .el-dialog__headerbtn {
            top: 24px;
            right: 32px;
        }
        .el-dialog__footer {
            padding: 12px 32px 16px;
        }
    }
}
</style>
