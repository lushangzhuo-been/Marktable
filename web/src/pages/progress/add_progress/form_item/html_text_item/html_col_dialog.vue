<template>
    <el-dialog
        :title="title"
        :visible.sync="dialogVisible"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        class="col-html-dialog"
        :class="isFullScreen ? 'fullscreen' : 'basic-ui'"
        @close="cancel"
    >
        <el-form
            ref="htmlColForm"
            class="basic-ui"
            :model="htmlColForm"
            size="small"
            label-width="0"
        >
            <el-form-item prop="content" :rules="rule">
                <Editor
                    v-if="dialogVisible"
                    class="editor"
                    ref="editorVue"
                    :height="editorHeight"
                    :content="htmlColForm.content"
                    :rowObj="rowObj"
                    @changeData="onChange"
                    @fullScreen="fullScreen"
                    @unFullScreen="unFullScreen"
                />
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
            <el-button
                size="small"
                class="basic-ui"
                :disabled="isBtnDisabled"
                @click="cancel"
                >取 消</el-button
            >
            <el-button
                size="small"
                type="primary"
                class="basic-ui"
                :disabled="isBtnDisabled"
                @click="onConfirm"
                >确 定</el-button
            >
        </span>
    </el-dialog>
</template>

<script>
import Editor from "@/pages/common/wang_editor.vue";
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
            rule: {},
            isFullScreen: false,
            screenHeight: document.body.clientHeight,
            editorHeight: document.body.clientHeight - 260
        };
    },
    computed: {
        title() {
            return this.colObj.name;
        },
        isBtnDisabled() {
            return this.colObj.can_modify === "yes" ? false : true;
        }
    },
    watch: {
        screenHeight: {
            handler(val) {
                if (val > 2200) {
                    this.editorHeight = 1900;
                } else {
                    this.editorHeight = val - 260;
                }
            }
        }
    },
    mounted() {
        this.resise();
        window.addEventListener("resize", this.resise);
    },
    methods: {
        resise() {
            const that = this;
            window.screenHeight = document.body.clientHeight;
            that.screenHeight = window.screenHeight;
        },
        fullScreen() {
            this.isFullScreen = true;
        },
        unFullScreen() {
            this.isFullScreen = false;
        },
        onChange(html) {
            if (html === "<p><br></p>" || !html) {
                this.htmlColForm.content = "";
                return;
            }
            this.htmlColForm.content = html;
        },
        openDialog(row, col) {
            this.colObj = col;
            this.rowObj = row;
            this.resise();
            this.htmlColForm.content = row[col.field_key];
            if (col.required === "yes") {
                this.rule = {
                    required: true,
                    message: "请输入内容",
                    trigger: "change"
                };
            } else {
                this.rule = {};
            }

            this.dialogVisible = true;
            this.$nextTick(() => {
                this.$refs.htmlColForm.clearValidate();
            });
        },
        cancel() {
            this.$nextTick(() => {
                this.$refs.htmlColForm.clearValidate();
            });
            this.dialogVisible = false;
        },
        onConfirm() {
            this.$refs.htmlColForm.validate((flag) => {
                if (flag) {
                    // this.$emit('refresh', this.htmlColForm);
                    this.cancel();
                    this.$emit("edit-form-item", this.htmlColForm.content);
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
            max-height: 2300px !important;
            padding: 0 32px;
        }
        .el-dialog__footer {
            padding: 16px 32px 32px;
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
.col-html-dialog.fullscreen {
    .el-dialog {
        width: 100% !important;
        height: 100% !important;
        border-radius: 0;
    }
}
</style>
