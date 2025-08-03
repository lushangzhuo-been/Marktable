<template>
    <el-dialog
        title="删除"
        append-to-body
        :visible.sync="dialogVisible"
        class="basic-ui"
    >
        <div slot="title" class="dialog-title">
            <i class="el-icon-warning"></i>
            <span>删除</span>
        </div>
        <div>确定要删除转换按钮"{{ curItem.name || "--" }}"吗？</div>
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
                确定
            </el-button>
        </span>
    </el-dialog>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress_setting";
export default {
    data() {
        return {
            curItem: {},
            curRow: {},
            dialogVisible: false
        };
    },
    props: {
        curId: {
            type: [String, Number],
            default: ""
        }
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        }
    },
    methods: {
        openDialog(item, row) {
            this.curItem = _.cloneDeep(item);
            this.curRow = _.cloneDeep(row);
            this.dialogVisible = true;
        },
        cancel() {
            this.dialogVisible = false;
        },
        onConfirm() {
            if (!this.curItem || !this.curItem.id) return;
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                step_id: this.curId,
                id: this.curItem.id
            };
            api.deleteBtn(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.$message.success("删除转化按钮成功");
                    this.$emit("on-confirm");
                }
            });
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
