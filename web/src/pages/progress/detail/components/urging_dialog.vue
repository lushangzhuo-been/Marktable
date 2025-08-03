<template>
    <el-dialog
        :title="title"
        :visible.sync="dialogVisible"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        class="basic-ui"
    >
        <!-- add-trigger-dialog  -->
        <el-form
            ref="form"
            class="basic-ui"
            :model="form"
            :rules="rule"
            label-width="84px"
            size="small"
        >
            <el-form-item prop="recipient" label="收件人:">
                <el-select
                    class="basic-ui width100"
                    size="small"
                    multiple
                    v-model="form.recipient"
                    placeholder="请选择"
                >
                    <el-option
                        v-for="item in recipientList"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    >
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item prop="urgingTitle" label="催办标题:">
                <el-input
                    class="basic-ui"
                    size="small"
                    placeholder="请输入催办标题"
                    v-model="form.suggest"
                >
                </el-input>
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
            <el-button size="small" @click="cancel">取消</el-button>
            <el-button size="small" type="primary" @click="onConfirm">
                确定
            </el-button>
        </span>
    </el-dialog>
</template>

<script>
export default {
    data() {
        return {
            title: "发送催办邮件",
            dialogVisible: false,
            form: {
                recipient: [],
                urgingTitle: ""
            },
            rule: {
                recipient: [
                    {
                        required: true,
                        message: "请选择类型"
                    }
                ]
            },
            recipientList: [
                {
                    label: "郭嘉",
                    value: "郭嘉"
                },
                {
                    label: "诸葛亮",
                    value: "诸葛亮"
                },
                {
                    label: "司马懿",
                    value: "司马懿"
                },
                {
                    label: "周瑜",
                    value: "周瑜"
                }
            ]
        };
    },
    watch: {
        dialogVisible: {
            handler(boolean) {
                if (!boolean) {
                    this.form = {
                        recipient: [],
                        urgingTitle: ""
                    };
                }
            }
        }
    },
    methods: {
        openDialog(type, row) {
            this.dialogVisible = true;
            this.$nextTick(() => {
                if (this.$refs.form) {
                    this.$refs.form.clearValidate();
                }
            });
        },
        cancel() {
            this.$refs.form.clearValidate();
            this.dialogVisible = false;
        },
        onConfirm() {
            this.$refs.form.validate((flag) => {
                if (flag) {
                    this.dialogVisible = false;
                    // this.$emit('on-confirm', this.form)
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.add-trigger-dialog {
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
<style lang="scss" scoped>
.width100 {
    width: 100%;
}
</style>
