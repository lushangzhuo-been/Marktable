<template>
    <el-dialog
        title="新建空间"
        :visible.sync="dialogVisible"
        class="add-space-dialog basic-ui"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        @close="cancel"
    >
        <el-form
            ref="newSpaceForm"
            class="basic-ui"
            :model="newSpaceForm"
            :rules="rule"
            size="small"
            label-width="52px"
        >
            <el-form-item prop="name" label="名称">
                <el-input
                    class="basic-ui"
                    v-model="newSpaceForm.name"
                    type="text"
                    placeholder="请输入空间名称"
                ></el-input>
            </el-form-item>
            <el-form-item prop="desc" label="描述">
                <el-input
                    class="basic-ui fixed-height-dialog"
                    v-model="newSpaceForm.desc"
                    type="textarea"
                    placeholder="请输入空间描述"
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
            dialogVisible: false,
            newSpaceForm: {
                name: "",
                desc: ""
            },
            rule: {
                name: [
                    {
                        required: true,
                        message: "请输入空间名称",
                        trigger: "blur"
                    }
                ],
                desc: [
                    {
                        required: true,
                        message: "请输入空间描述",
                        trigger: "blur"
                    }
                ]
            }
        };
    },
    methods: {
        openDialog() {
            this.dialogVisible = true;
        },
        cancel() {
            this.$refs.newSpaceForm.clearValidate();
            this.newSpaceForm = {
                name: "",
                desc: ""
            };
            this.dialogVisible = false;
        },
        onConfirm() {
            this.$refs.newSpaceForm.validate((flag) => {
                if (flag) {
                    this.$emit("on-confirm", this.newSpaceForm);
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.el-form.basic-ui {
    margin: 6px 0;
}
</style>
