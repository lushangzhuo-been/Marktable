<template>
    <el-dialog
        title="删除"
        :visible.sync="dialogVisible"
        append-to-body
        class="add-trigger-dialog basic-ui"
    >
        <div>{{ `确定要删除图表"${curItem.title || "--"}"吗？` }}</div>
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
            curItem: {}
        };
    },
    methods: {
        openDialog(item) {
            this.dialogVisible = true;
            this.curItem = item;
        },
        cancel() {
            this.dialogVisible = false;
        },
        onConfirm() {
            this.$emit("on-confirm", this.curItem);
            this.dialogVisible = false;
        }
    }
};
</script>

<style lang="scss" scoped>
.add-trigger-dialog {
    .width100 {
        width: 100%;
    }
}
</style>
