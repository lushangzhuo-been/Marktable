<template>
    <el-dialog
        title="删除"
        :visible.sync="dialogVisible"
        append-to-body
        class="basic-ui"
    >
        <div slot="title" class="dialog-title">
            <i class="el-icon-warning"></i>
            <span>删除</span>
        </div>
        确定要删除标题为"{{ curRow.title || "--" }}"的任务吗?
        <span slot="footer" class="dialog-footer">
            <el-button size="small" @click="cancel">取 消</el-button>
            <el-button size="small" type="primary" @click="onConfirm"
                >确 定</el-button
            >
        </span>
    </el-dialog>
</template>

<script>
export default {
    data() {
        return {
            curRow: {},
            dialogVisible: false
        };
    },
    methods: {
        openDialog(row) {
            this.curRow = row;
            this.dialogVisible = true;
        },
        cancel() {
            this.dialogVisible = false;
            this.curRow = {};
            this.$emit("cancel");
        },
        onConfirm() {
            this.$emit("on-confirm", this.curRow);
        }
    }
};
</script>

<style lang="scss" scoped>
.width100 {
    width: 100%;
}
.dialog-title {
    display: flex;
    align-items: center;
}
</style>
