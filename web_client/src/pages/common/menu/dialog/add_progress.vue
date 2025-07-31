<template>
    <el-dialog
        title="添加流程"
        class="basic-ui add-progress-dialog"
        :visible.sync="dialogVisible"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        @close="cancel"
    >
        <el-form
            ref="newProgressForm"
            class="basic-ui"
            :model="newProgressForm"
            :rules="rule"
            label-width="52px"
        >
            <el-form-item prop="name" label="名称">
                <el-input
                    class="basic-ui"
                    v-model="newProgressForm.name"
                    type="text"
                    size="small"
                    maxlength="25"
                    show-word-limit
                    placeholder="请输入流程名称"
                ></el-input>
            </el-form-item>
            <el-form-item prop="desc" label="描述">
                <el-input
                    class="basic-ui fixed-height-dialog"
                    v-model="newProgressForm.desc"
                    type="textarea"
                    size="small"
                    placeholder="请输入流程描述"
                ></el-input>
            </el-form-item>
            <el-form-item label="类型">
                <el-radio v-model="newProgressForm.mode" label="public"
                    >公有流程</el-radio
                >
                <el-radio v-model="newProgressForm.mode" label="private"
                    >私有流程</el-radio
                >
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
            newProgressForm: {
                name: "",
                desc: "",
                mode: "public"
            },
            rule: {
                name: [
                    {
                        required: true,
                        message: "请输入空间名称",
                        trigger: "blur"
                    }
                ],
                desc: [
                    {
                        required: true,
                        message: "请输入空间描述",
                        trigger: "blur"
                    }
                ]
            },
            curNodeData: {}
        };
    },
    methods: {
        openDialog(obj, nodeData) {
            this.curSpace = obj;
            if (nodeData && Object.keys(nodeData).length) {
                this.curNodeData = nodeData;
            } else {
                this.curNodeData = {};
            }
            this.dialogVisible = true;
        },
        cancel() {
            this.$refs.newProgressForm.clearValidate();
            this.newProgressForm = {
                name: "",
                desc: "",
                mode: "public"
            };
            this.dialogVisible = false;
        },
        onConfirm() {
            this.$refs.newProgressForm.validate((flag) => {
                if (flag) {
                    let impl_table_name = this.curNodeData.impl_table_name;
                    let params = {
                        ws_id: this.curSpace.id,
                        ...this.newProgressForm,
                        ws_file_id:
                            impl_table_name === "ws_file"
                                ? this.curNodeData.id
                                : 0 // 如果是给文件夹中添加流程为this.curNodeData.id， 否则传0
                    };

                    api.addProgress(params).then((res) => {
                        if (!res) return;
                        if (res.resultCode === 200) {
                            this.$emit("add-progress-success-refresh", {
                                ...params,
                                impl_table_name
                            });
                        } else if (
                            res.resultCode === 401 ||
                            res.resultCode === 403
                        ) {
                            this.dialogVisible = false;
                        }
                    });
                }
            });
        }
    }
};
</script>

<style lang="scss">
.add-progress-dialog {
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
    }
}
</style>
