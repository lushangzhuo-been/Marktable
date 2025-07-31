<template>
    <el-dialog
        title="删除"
        :visible.sync="dialogVisible"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        class="basic-ui"
    >
        <div class="msg">
            若确定要删除该视图，请在下方输入框内输入视图名称。
        </div>
        <el-form
            class="basic-ui"
            size="small"
            ref="form"
            :model="form"
            :rules="rule"
            label-width="84px"
        >
            <el-form-item prop="name" label="视图名称:">
                <el-input
                    class="basic-ui"
                    v-model="form.name"
                    placeholder="请输入视图名称"
                ></el-input>
            </el-form-item>
        </el-form>
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
            form: {
                name: ""
            },
            rule: {
                name: [
                    {
                        required: true,
                        message: "请输入视图名称"
                    }
                ]
            },
            tabInfo: {}
        };
    },
    watch: {
        dialogVisible: {
            handler(boolean) {
                if (!boolean) {
                    this.$refs.form.clearValidate();

                    // this.form = {
                    //     name: "",
                    // };
                }
            }
        }
    },
    methods: {
        openDialog(tab) {
            this.dialogVisible = true;
            this.tabInfo = tab;
            this.form.name = "";
            this.$nextTick(() => {
                if (this.$refs.form) {
                    this.$refs.form.clearValidate();
                }
            });
        },
        cancel() {
            this.dialogVisible = false;
        },
        onConfirm() {
            if (this.form.name === this.tabInfo.name) {
                this.$emit("delete-view", this.tabInfo);
                this.dialogVisible = false;
            } else {
                this.$message({
                    showClose: true,
                    message: "请输入正确名称",
                    type: "warning"
                });
            }
        }
    }
};
</script>

<style lang="scss" scoped>
.msg {
    font-weight: 400;
    font-size: 16px;
    color: #171e31;
    height: 18px;
    line-height: 18px;
    margin-bottom: 24px;
}
</style>
