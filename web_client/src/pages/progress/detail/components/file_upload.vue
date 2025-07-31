<template>
    <!-- 附件模块 -->
    <div>
        <el-upload
            class="progress-upload"
            :class="{
                'disabled-permission': fileAuth.edit !== 'yes'
            }"
            action=""
            drag
            :auto-upload="false"
            :show-file-list="false"
            :on-change="fileChange"
            :disabled="fileAuth.edit !== 'yes'"
        >
            <i class="el-icon-upload"></i>
            <div class="el-upload__text">
                将文件拖到此处，或<em>点击上传</em>
            </div>
            <div style="font-size: 12px; color: #82889e">
                <!-- 支持xls/xlsx、doc、ppt格式 -->
                支持{{ typeList ? typeList : "xls/xlsx、doc、ppt" }}格式
            </div>
        </el-upload>
        <div class="file-content">
            <!-- 上传的文件 -->
            <div class="file-item" v-show="isUploading">
                <img
                    class="file-box"
                    :src="getFileIcon(currentUploadingFile)"
                    alt=""
                />
                <div class="msg-content">
                    <div class="title-operation">
                        <div class="msg-content-item">
                            <div v-overflow class="name" @click="preview(item)">
                                {{ currentUploadingFile.file_name }}
                            </div>
                        </div>
                    </div>
                    <div class="creator-msg">
                        <el-progress
                            style="margin-top: 4px"
                            :percentage="progressPercent"
                            :stroke-width="4"
                            status="success"
                            color="#1890FF"
                        >
                        </el-progress>
                    </div>
                </div>
            </div>
            <div
                class="file-item"
                v-for="(item, index) in fileList"
                :key="index"
            >
                <!-- <b
                    @click="preview(item)"
                    class="file-box"
                    :class="getFileTypeClass(item.fileType)"
                ></b> -->
                <img
                    class="file-box"
                    :src="getFileIcon(item)"
                    alt=""
                    @click="preview(item)"
                />
                <div class="msg-content">
                    <div class="title-operation">
                        <div class="msg-content-item">
                            <tip-one :text="item.file_name" position="top">
                                <div class="name" @click="preview(item)">
                                    {{ item.file_name }}
                                </div>
                            </tip-one>
                            <div v-overflow class="size">
                                {{ item.file_size_desc }}
                            </div>
                        </div>

                        <div class="operation">
                            <b
                                class="operation-box pre-view"
                                @click="preview(item)"
                            ></b>
                            <b
                                v-if="fileAuth.edit === 'yes'"
                                class="operation-box download"
                                @click="downloadFile(item)"
                            ></b>
                            <b
                                v-if="fileAuth.delete === 'yes'"
                                class="operation-box delete"
                                @click="deleteFile(item)"
                            ></b>
                        </div>
                    </div>
                    <div class="creator-msg">
                        <span class="ver"> V{{ item.version_code }} </span>
                        <span class="creator">{{ item.creator }}</span>
                        <span class="time">{{ item.created_at }}</span>
                    </div>
                </div>
            </div>
        </div>
        <!-- 删除弹窗 -->
        <el-dialog
            :visible.sync="deleteVisible"
            class="basic-ui"
            append-to-body
        >
            <div slot="title" class="dialog-title">
                <i class="el-icon-warning"></i>
                <span>删除</span>
            </div>
            <span>此操作会删除该文件下所有版本，确定删除吗？</span>
            <span slot="footer" class="dialog-footer">
                <el-button class="basic-ui" size="small" @click="cancel"
                    >取 消</el-button
                >
                <el-button
                    class="basic-ui"
                    size="small"
                    type="primary"
                    @click="onConfirmDeleteAll"
                    >确 定</el-button
                >
            </span>
        </el-dialog>
        <!-- 预览弹窗 -->
        <file-preview
            ref="previewDialog"
            @closeDialog="closePreviewDialog"
        ></file-preview>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import TipOne from "@/pages/common/tooltip_one_line.vue";
import FilePreview from "./file_preview/index.vue";
import { imgHost } from "@/assets/tool/const";
export default {
    props: {
        detailId: {
            type: String,
            default: ""
        },
        permissionEdit: {
            type: Boolean,
            default: false
        }
    },
    components: {
        TipOne,
        FilePreview
    },
    data() {
        return {
            fileList: [],
            fileAuth: {},
            typeList: "",
            fileUrl: "",
            deleteVisible: false,
            curDeleteFile: {},
            currentUploadingFile: {},
            isUploading: false,
            progressPercent: 0
        };
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        },
        getFileIcon() {
            return (item) => {
                if (!Object.keys(item).length) {
                    return require("@/assets/image/file_type/others.svg");
                }
                let fileType = this.splitAndGetLast(item.file_name, ".");
                if (fileType === "doc" || fileType === "docx") {
                    return require("@/assets/image/file_type/doc.svg");
                } else if (fileType === "xls" || fileType === "xlsx") {
                    return require("@/assets/image/file_type/excel.svg");
                } else if (fileType === "ppt" || fileType === "pptx") {
                    return require("@/assets/image/file_type/ppt.svg");
                } else if (fileType === "xml") {
                    return require("@/assets/image/file_type/xml.svg");
                } else if (
                    fileType === "png" ||
                    fileType === "jpg" ||
                    fileType === "jpeg" ||
                    fileType === "gif"
                ) {
                    return `${imgHost}${item.relative_path}`;
                } else {
                    return require("@/assets/image/file_type/others.svg");
                }
            };
        }
    },
    watch: {},

    methods: {
        // 文件预览
        preview(item) {
            this.$refs.previewDialog.show(item, this.detailId, this.fileAuth);
        },
        closePreviewDialog() {
            this.getFileList();
        },
        fileChange(file) {
            this.currentUploadingFile.file_name = file.name;
            this.isUploading = true;
            // 将file 入参上传
            this.uploadFile(file);
        },
        downloadTypeCheck() {
            api.downloadTypeCheck().then((res) => {
                if (res && res.resultCode === 200) {
                    this.typeList = res.data;
                }
            });
        },
        // 获取文件权限
        getFileAuth() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: this.detailId
            };
            api.getFileAuth(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.fileAuth = res.data;
                } else {
                    this.fileAuth = {};
                }
            });
        },
        // 获取文件列表
        getFileList() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                issue_id: this.detailId,
                group_id: -1,
                is_current_version: 1 //
            };
            api.getFileList(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.fileList = res.data;
                    // 遍历获取文件类型 并转为小写
                    if (this.fileList && this.fileList.length) {
                        this.filesCount = this.fileList.length;
                        this.fileList.forEach((file) => {
                            file.fileType = this.splitAndGetLast(
                                file.file_name,
                                "."
                            );
                        });
                    } else {
                        this.fileList = [];
                        this.filesCount = 0;
                    }
                } else {
                    this.fileList = [];
                    this.filesCount = 0;
                }
                this.$emit("transmit-file-count", this.filesCount);
            });
        },
        // 上传
        uploadFile(file) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                issue_id: this.detailId,
                tmpl_file: file.raw,
                group_id: -1,
                is_current_version: 1
            };
            let formData = new FormData();
            Object.keys(params).forEach((key) => {
                formData.append(key, params[key]);
            });
            const uploadChange = (progressEvent) => {
                this.progressPercent = Math.ceil(
                    (progressEvent.loaded * 100) / progressEvent.total
                );
            };
            api.uploadFile(formData, uploadChange).then((res) => {
                if (res && res.resultCode === 200) {
                    this.getFileList();
                    this.isUploading = false;
                    this.progressPercent = 0;
                    this.currentUploadingFile = {};
                } else {
                    this.isUploading = false;
                    this.progressPercent = 0;
                    this.currentUploadingFile = {};
                }
            });
        },

        // 下载
        downloadFile(file) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                issue_id: this.detailId,
                id: file.id,
                download_file_type: "original_name"
            };
            api.downloadFile(params);
        },
        //  取消全部删除弹窗
        cancel() {
            this.deleteVisible = false;
        },
        // 确认删除全部弹窗
        onConfirmDeleteAll() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                issue_id: this.detailId,
                id: this.curDeleteFile.id,
                group_id: -1,
                is_current_version: 1
            };
            api.deleteFile(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.$message({
                        showClose: true,
                        message: "删除成功",
                        type: "success"
                    });
                    this.cancel();
                    this.getFileList();
                }
            });
        },
        // 删除
        deleteFile(file) {
            this.curDeleteFile = _.cloneDeep(file);
            this.deleteVisible = true;
        },
        splitAndGetLast(str, separator) {
            // 使用split方法按separator分割字符串，然后取最后一个元素
            return str.split(separator).slice(-1)[0].toLowerCase();
        },
        getFileTypeClass(type) {
            if (type === "doc") {
                return "doc";
            } else if (type === "xls" || type === "xlsx") {
                return "excel";
            } else if (type === "jpg") {
                return "jpg";
            } else if (type === "ppt" || type === "pptx") {
                return "ppt";
            } else if (type === "xml") {
                return "xml";
            } else {
                return "others";
            }
        }
    }
};
</script>

<style lang="scss" scoped>
// 上传block
::v-deep.progress-upload {
    .el-upload {
        width: 100%;
        .el-upload-dragger {
            width: 100%;
        }
    }
    &.disabled-permission {
        .el-upload {
            cursor: not-allowed;
            .el-upload-dragger {
                cursor: not-allowed;
            }
        }
    }
}
// 列表block
.file-content {
    margin-top: 16px;
    .file-item {
        box-sizing: border-box;
        height: 60px;
        display: flex;
        padding: 12px 8px;
        background: #f5f6f8;
        border-radius: 4px;
        margin-bottom: 8px;
        img.file-box {
            display: inline-block;
            width: 32px;
            height: 32px;
            margin: auto 8px auto 0;
            // background-size: 100% 100%;
            cursor: pointer;
            // &.doc {
            //     background-image: url("~@/assets/image/file_type/doc.svg");
            // }
            // &.excel {
            //     background-image: url("~@/assets/image/file_type/excel.svg");
            // }
            // &.jpg {
            //     background-image: url("~@/assets/image/file_type/jpg.svg");
            // }
            // &.ppt {
            //     background-image: url("~@/assets/image/file_type/ppt.svg");
            // }
            // &.xml {
            //     background-image: url("~@/assets/image/file_type/xml.svg");
            // }
            // &.others {
            //     background-image: url("~@/assets/image/file_type/others.svg");
            // }
        }
        .msg-content {
            width: calc(100% - 44px);
            .title-operation {
                height: 18px;
                display: flex;
                justify-content: space-between;
                .msg-content-item {
                    height: 18px;
                    width: calc(100% - 70px);
                    line-height: 18px;
                    font-weight: 400;
                    font-size: 12px;
                    display: flex;
                    align-items: center;
                    .name {
                        color: #171e31;
                        max-width: calc(100% - 95px);
                        overflow: hidden;
                        white-space: nowrap;
                        text-overflow: ellipsis;
                        cursor: pointer;
                        &:hover {
                            color: #1890ff;
                        }
                    }
                    .size {
                        margin-left: 20px;
                        color: #82889e;
                        max-width: 60px;
                        overflow: hidden;
                        white-space: nowrap;
                        text-overflow: ellipsis;
                    }
                }

                .operation {
                    .operation-box {
                        display: inline-block;
                        width: 16px;
                        height: 16px;
                        background-size: 100% 100%;
                        cursor: pointer;
                        vertical-align: inherit;
                        &.pre-view {
                            margin-right: 8px;
                            background-image: url("~@/assets/image/file_upload_preview/preview.png");
                            &:hover {
                                background-image: url("~@/assets/image/file_upload_preview/preview_active.png");
                            }
                        }
                        &.download {
                            background-image: url("~@/assets/image/common/download.svg");
                            &:hover {
                                background-image: url("~@/assets/image/common/download-active.svg");
                            }
                        }
                        &.delete {
                            margin-left: 8px;
                            background-image: url("~@/assets/image/common/delete.svg");
                            &:hover {
                                background-image: url("~@/assets/image/common/delete-active.svg");
                            }
                        }
                    }
                }
            }
            .creator-msg {
                height: 22px;
                line-height: 22px;
                font-weight: 400;
                font-size: 12px;
                color: #a6abbc;
                ::v-deep .el-progress {
                    margin-top: 4px;
                    .el-progress-bar {
                        padding-right: 10px;
                    }
                    .el-progress__text {
                        display: none;
                    }
                }
                .ver {
                    display: inline-block;
                    height: 18px;
                    line-height: 18px;
                    color: #0c6c50;
                    background-color: #e8eefa;
                    padding: 0 4px;
                    margin-right: 8px;
                    border-radius: 4px;
                }
                .creator {
                    margin-right: 12px;
                }
            }
        }
    }
}
</style>
<style lang="scss">
.preview-file.basic-ui .el-dialog {
    width: 90%;
    height: 90%;
    .el-dialog__body {
        height: calc(100% - 73px - 16px);
        max-height: none;
        padding: 0 0 16px;
        overflow: inherit;
        border-top: 1px solid #e6e9f0;
    }
}
</style>
