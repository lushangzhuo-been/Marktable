<template>
    <el-dialog
        title="删除"
        :visible.sync="dialogVisible"
        append-to-body
        class="add-trigger-dialog basic-ui"
    >
        <div>{{ `确定要删除图表"${chartMsg.title || "--"}"吗？` }}</div>
        <span slot="footer" class="dialog-footer">
            <el-button size="small" class="basic-ui" @click="cancel"
                >取消</el-button
            >
            <el-button
                size="small"
                class="basic-ui"
                type="primary"
                @click="onConfirm"
            >
                确定删除
            </el-button>
        </span>
    </el-dialog>
</template>

<script>
export default {
    data() {
        return {
            dialogVisible: false,
            chartMsg: {}
        };
    },
    methods: {
        openDialog(item) {
            this.dialogVisible = true;
            this.chartMsg = item;
        },
        cancel() {
            this.dialogVisible = false;
        },
        onConfirm() {
            this.$emit("delete", this.chartMsg);
            this.dialogVisible = false;
        }
    }
};
</script>

<style lang="scss" scoped>
// .add-trigger-dialog {
//     .el-dialog {
//         width: 520px;
//         border-radius: 4px;
//         .el-dialog__header {
//             padding: 24px 32px;
//             font-size: 16px;
//             color: #2f384c;
//         }
//         .el-dialog__body {
//             padding: 0 32px;
//         }
//         .el-dialog__footer {
//             padding: 0 32px 32px;
//         }
//         .el-dialog__headerbtn {
//             top: 24px;
//             right: 32px;
//         }
//         .el-form-item {
//             margin-bottom: 24px;
//         }
//         .el-dialog__footer {
//             text-align: right;
//         }
//     }
// }
</style>
<style lang="scss" scoped>
.add-trigger-dialog {
    .width100 {
        width: 100%;
    }
}
</style>
