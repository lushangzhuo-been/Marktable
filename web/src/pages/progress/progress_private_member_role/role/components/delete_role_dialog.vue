<template>
    <el-dialog
        title="删除"
        :visible.sync="dialogVisible"
        class="basic-ui"
        append-to-body
        @close="cancel"
    >
        <div slot="title" class="dialog-title">
            <i class="el-icon-warning"></i>
            <span>删除</span>
        </div>
        确定要删除角色"{{ curRow.name || "--" }}"吗
        <span slot="footer" class="dialog-footer">
            <el-button size="small" @click="cancel">取 消</el-button>
            <el-button size="small" type="primary" @click="onConfirm"
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
            curRow: {},
            dialogVisible: false
        };
    },
    methods: {
        openDialog(row) {
            this.curRow = _.cloneDeep(row);
            this.dialogVisible = true;
        },
        cancel() {
            this.dialogVisible = false;
        },
        onConfirm() {
            // 删除角色
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                id: this.curRow.id
            };
            api.delRole(params).then((res) => {
                if (res && res.resultCode === 200) {
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
