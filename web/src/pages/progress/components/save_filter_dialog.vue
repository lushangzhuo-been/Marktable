<template>
    <el-dialog
        :visible.sync="dialogVisible"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        class="basic-ui"
    >
        <div slot="title" class="title">
            <b class="save-box"></b>
            保存为新视图
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
                确定
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
            }
        };
    },
    watch: {
        dialogVisible: {
            handler(boolean) {
                if (!boolean) {
                    this.form = {
                        name: ""
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
                    this.$emit("create-view", this.form);
                    this.dialogVisible = false;
                }
            });
        }
    }
};
</script>
<style lang="scss" scoped>
.title {
    font-weight: 500;
    font-size: 16px;
    .save-box {
        display: inline-block;
        width: 20px;
        height: 20px;
        background-size: 100% 100%;
        vertical-align: middle;
        background-image: url(@/assets/image/progress/save.svg);
    }
}
</style>
