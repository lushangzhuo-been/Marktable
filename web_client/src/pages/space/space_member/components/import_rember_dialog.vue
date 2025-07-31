<template>
    <el-dialog
        title="导入成员"
        :visible.sync="dialogVisible"
        class="add-rember-dialog basic-ui"
        append-to-body
        :close-on-click-modal="false"
        :close-on-press-escape="false"
    >
        <el-upload
            class="upload-demo"
            drag
            action="https://jsonplaceholder.typicode.com/posts/"
            multiple
        >
            <i class="el-icon-upload"></i>
            <div class="el-upload__text">
                将文件拖到此处，或<em>点击上传</em>
            </div>
            <div class="format">仅支持xlsx、XLSX</div>
        </el-upload>
        <div></div>
        <span slot="footer" class="dialog-footer">
            <el-button size="small" @click="cancel">取 消</el-button>
            <el-button size="small" type="primary" @click="onConfirm"
                >确 定</el-button
            >
        </span>
    </el-dialog>
</template>

<script>
import api from "@/common/api/module/space";
export default {
    data() {
        return {
            dialogVisible: false
        };
    },
    methods: {
        openDialog() {
            this.dialogVisible = true;
        },
        cancel() {
            this.dialogVisible = false;
        },
        onConfirm() {
            this.$emit("on-confirm", this.spaceRemberForm);
        }
    }
};
</script>

<style lang="scss">
.add-rember-dialog {
    .el-dialog {
        width: 490px;
        border-radius: 4px;
        .el-dialog__header {
            padding: 24px 32px;
            font-size: 16px;
            color: #2f384c;
        }
        .el-dialog__body {
            padding: 0 32px;
        }
        .el-dialog__footer {
            padding: 0 32px 32px;
        }
        .el-dialog__headerbtn {
            top: 24px;
            right: 32px;
        }
        .el-upload {
            width: 100%;
        }
        .el-upload-dragger {
            width: 100%;
            background-color: #fbfbfb;
            .format {
                font-size: 12px;
                color: #828793;
                margin-top: 6px;
            }
        }
    }
}
</style>
<style lang="scss" scoped>
.add-rember-dialog {
    .width100 {
        width: 100%;
    }
}
</style>
