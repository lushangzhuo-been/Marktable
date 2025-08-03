<template>
    <el-dialog
        :title="title"
        :visible.sync="dialogVisible"
        class="col-html-dialog basic-ui"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        @close="cancel"
    >
        <el-form
            ref="htmlColForm"
            class="basic-ui"
            :model="htmlColForm"
            :rules="rule"
            size="small"
            label-width="0"
        >
            <el-form-item prop="content">
                <div>
                    <Editor
                        ref="kindeditor"
                        :id="rowObj._id"
                        :minHeight="300"
                        :html="htmlColForm.content"
                        @input="getContent"
                    />
                </div>
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
import Editor from "@/pages/common/kind_editor/index.vue";
export default {
    components: {
        Editor
    },
    data() {
        return {
            colObj: {},
            rowObj: {},
            dialogVisible: false,
            htmlColForm: {
                content: ""
            },
            placeholder: "请输入",
            rule: {
                content: [
                    {
                        required: true,
                        message: "请输入内容",
                        trigger: "blur"
                    }
                ]
            }
        };
    },
    computed: {
        title() {
            return this.colObj.name;
        }
    },
    methods: {
        getContent(content) {
            this.htmlColForm.content = content;
        },
        openDialog(row, col) {
            this.colObj = col;
            this.rowObj = row;
            this.htmlColForm.content = row[col.field_key];
            this.dialogVisible = true;
        },
        cancel() {
            this.$refs.htmlColForm.clearValidate();
            this.htmlColForm = {
                content: ""
            };
            this.dialogVisible = false;
        },
        onConfirm() {
            this.$refs.htmlColForm.validate((flag) => {
                if (flag) {
                    // this.$emit('refresh', this.htmlColForm);
                    this.$emit(
                        "edit-form-item",
                        JSON.stringify(this.htmlColForm.content),
                        this.colObj.field_key,
                        this.rowObj._id,
                        this.colObj.mode
                    );
                }
            });
        }
    }
};
</script>

<style lang="scss">
.col-html-dialog {
    .el-dialog {
        width: 90% !important;
        border-radius: 4px;
        .el-dialog__header {
            padding: 24px 32px;
            font-size: 16px;
            color: #2f384c;
        }
        .el-dialog__body {
            padding: 0 32px;
        }
        .el-dialog__headerbtn {
            top: 24px;
            right: 32px;
        }
        .el-form-item {
            margin-bottom: 24px;
        }
    }
}
</style>
