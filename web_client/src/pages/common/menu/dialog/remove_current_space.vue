<template>
    <el-dialog
        :title="title"
        :visible.sync="dialogVisible"
        class="basic-ui remove-space-dialog"
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
            ref="delSpaceForm"
            class="basic-ui"
            :model="delSpaceForm"
            :rules="rule"
            label-width="52px"
        >
            <el-form-item prop="name" label="名称">
                <el-input
                    class="basic-ui"
                    v-model="delSpaceForm.name"
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
            delSpaceForm: {
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
            if (!obj) return;
            this.objData = obj;
            this.$nextTick(() => {
                this.$refs.delSpaceForm.clearValidate();
            });
            this.title = "删除空间";
            this.placeholder = "请输入要删除的空间名称";
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
                                this.delSpaceForm.name !==
                                this.objData.spaceName
                            ) {
                                callback("请输入正确的空间名称");
                            } else {
                                callback();
                            }
                        }
                    }
                ]
            };
        },
        cancel() {
            this.delSpaceForm = {
                name: ""
            };
            this.dialogVisible = false;
        },
        onConfirm() {
            this.$refs.delSpaceForm.validate((flag) => {
                if (flag) {
                    this.$emit("on-confirm", this.objData);
                }
            });
        }
    }
};
</script>

<style lang="scss">
.remove-space-dialog {
    .dialog-title {
        display: flex;
        align-items: center;
    }
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
