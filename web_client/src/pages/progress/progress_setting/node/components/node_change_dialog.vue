<template>
    <el-dialog
        title="设置"
        :visible.sync="dialogVisible"
        class="node-change-dialog"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        @close="cancel"
    >
        <div class="setting">
            <div class="radio-group">
                <el-radio-group v-model="radio">
                    <el-radio class="first-ntd" :label="1">显示当前值</el-radio>
                    <br class="br" />
                    <el-radio :label="2">清空当前值</el-radio>
                </el-radio-group>
            </div>
            <div class="data-range">
                <div class="data-range-title">
                    <el-checkbox v-model="checked"
                        >重新选择数据范围</el-checkbox
                    >
                </div>
                <div>请选择</div>
            </div>
        </div>
        <span slot="footer" class="dialog-footer">
            <el-button size="small" @click="cancel">取 消</el-button>
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
import _ from "lodash";
export default {
    components: {},
    computed: {
        getLabelWidth() {
            return this.editFiledForm.mode === "person_com" ? "72px" : "60px";
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        }
    },
    data() {
        return {
            dialogVisible: false,
            radio: 1,
            checked: [1]
        };
    },
    methods: {
        cancel() {
            this.dialogVisible = false;
        },
        onConfirm() {
            this.dialogVisible = false;
        },
        openDialog(row) {
            this.dialogVisible = true;
        }
    }
};
</script>
<style lang="scss">
.setting {
    .radio-group {
        padding: 20px 0 16px;
        .el-radio {
            &.first-ntd {
                margin-bottom: 16px;
            }
        }
    }
    .data-range {
        padding: 20px 0 24px;
        border-top: 1px dotted #e6e9f0;
        .data-range-title {
            margin-bottom: 12px;
        }
    }
}
.node-change-dialog {
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
            border-top: 1px solid #e6e9f0;
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
