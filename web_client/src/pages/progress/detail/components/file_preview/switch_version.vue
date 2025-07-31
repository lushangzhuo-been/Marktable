<template>
    <div>
        <slot name="collapse"></slot>
        <div class="title">版本</div>
        <div class="cur-title flex-between">
            <span>当前版本</span>
            <el-upload
                v-if="fileAuth.edit === 'yes'"
                class="upload-demo"
                action=""
                :auto-upload="false"
                :show-file-list="false"
                :on-change="fileChange"
            >
                <span class="upload">
                    <b class="upload-img" @click="preview(item)"></b
                    >上传版本</span
                >
            </el-upload>
        </div>
        <!-- 当前版本 -->
        <div
            class="ver-box"
            :class="{ actived: curVerObj.actived }"
            @click.capture="previewCurrent"
        >
            <div class="flex">
                <img
                    :src="getFileIcon(curVerObj)"
                    alt=""
                    width="48px"
                    height="48px"
                />
                <div class="sub-right">
                    <div class="line flex-between">
                        <span class="name-size">
                            <span v-overflow class="name">{{
                                curVerObj.file_name
                            }}</span>
                            <span v-overflow class="size">{{
                                curVerObj.file_size_desc
                            }}</span>
                        </span>
                        <el-popover
                            placement="bottom-end"
                            trigger="click"
                            popper-class="opera-box-popover"
                            :visible-arrow="false"
                        >
                            <div
                                class="pop-item"
                                @click="downLoadCurrent(curVerObj)"
                            >
                                <img
                                    :src="
                                        require(`@/assets/image/common/download.svg`)
                                    "
                                    alt=""
                                    width="16px"
                                    height="16px"
                                />
                                下载
                            </div>
                            <!-- 阻止外层冒泡得写法 -->
                            <template #reference v-if="fileAuth.edit === 'yes'">
                                <b class="more-icon" @click.stop></b>
                            </template>
                        </el-popover>
                    </div>
                    <div class="line">
                        <span v-overflow class="creator">{{
                            curVerObj.creator
                        }}</span>
                        <span v-overflow class="date">{{
                            curVerObj.created_at
                        }}</span>
                    </div>
                    <div v-show="curVerObj.version_code">
                        <span class="ver"> V{{ curVerObj.version_code }} </span>
                    </div>
                </div>
            </div>
        </div>
        <div v-show="isHistoryVerShow" class="his-title">历史版本</div>
        <!-- 历史版本数组 -->
        <div
            v-show="isHistoryVerShow"
            id="hisBox"
            class="his-box"
            @scroll="handleScroll"
        >
            <div
                class="version-box"
                v-for="(item, index) in versionArr"
                :key="index"
            >
                <div
                    class="ver-box"
                    :class="{ actived: item.actived }"
                    @click="previewVersion(item)"
                >
                    <div class="flex">
                        <img
                            :src="getFileIcon(item)"
                            alt=""
                            width="48px"
                            height="48px"
                        />
                        <div class="sub-right">
                            <div class="line flex-between">
                                <span class="name-size">
                                    <span v-overflow class="name">{{
                                        item.file_name
                                    }}</span>
                                    <span v-overflow class="size">{{
                                        item.file_size_desc
                                    }}</span>
                                </span>
                                <el-popover
                                    placement="bottom-end"
                                    trigger="click"
                                    popper-class="opera-box-popover"
                                    :visible-arrow="false"
                                >
                                    <div
                                        v-if="fileAuth.edit === 'yes'"
                                        class="pop-item his"
                                        @click="
                                            operatHisVersion(item, 'setting')
                                        "
                                    >
                                        <img
                                            :src="
                                                require(`@/assets/image/file_upload_preview/preview_dialog/setting.png`)
                                            "
                                            alt=""
                                            width="16px"
                                            height="16px"
                                        />
                                        设置当前版本
                                    </div>
                                    <div
                                        v-if="fileAuth.edit === 'yes'"
                                        class="pop-item his"
                                        @click="
                                            operatHisVersion(item, 'download')
                                        "
                                    >
                                        <img
                                            :src="
                                                require(`@/assets/image/common/download.svg`)
                                            "
                                            alt=""
                                            width="16px"
                                            height="16px"
                                        />
                                        下载
                                    </div>
                                    <div
                                        v-if="fileAuth.delete === 'yes'"
                                        class="pop-item his"
                                        @click="
                                            operatHisVersion(item, 'delete')
                                        "
                                    >
                                        <img
                                            :src="
                                                require(`@/assets/image/common/delete.svg`)
                                            "
                                            alt=""
                                            width="16px"
                                            height="16px"
                                        />
                                        删除
                                    </div>
                                    <b
                                        v-if="
                                            fileAuth.edit === 'yes' ||
                                            fileAuth.delete === 'yes'
                                        "
                                        slot="reference"
                                        class="more-icon"
                                        @click.stop
                                    ></b>
                                </el-popover>
                            </div>
                            <div class="line">
                                <span v-overflow class="creator">{{
                                    item.creator
                                }}</span>
                                <span v-overflow class="date">{{
                                    item.created_at
                                }}</span>
                            </div>
                            <div v-show="item.version_code">
                                <span class="ver">
                                    V{{ item.version_code }}
                                </span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <!-- 删除历史版本弹窗 -->
        <el-dialog
            :visible.sync="deleteVisible"
            class="basic-ui"
            append-to-body
        >
            <div slot="title" class="dialog-title">
                <i class="el-icon-warning"></i>
                <span>删除</span>
            </div>
            <span
                >确定要删除{{
                    curDeleteFile.version_code
                        ? `V${curDeleteFile.version_code}`
                        : ""
                }}版本{{
                    curDeleteFile.file_name
                        ? `"${curDeleteFile.file_name}"`
                        : ""
                }}吗？</span
            >
            <span slot="footer" class="dialog-footer">
                <el-button
                    class="basic-ui"
                    size="small"
                    @click="cancelDelVersion"
                    >取 消</el-button
                >
                <el-button
                    class="basic-ui"
                    size="small"
                    type="primary"
                    @click="onConfirmDelVersion"
                    >确 定</el-button
                >
            </span>
        </el-dialog>
    </div>
</template>
<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import { imgHost } from "@/assets/tool/const";
export default {
    props: {
        curFileItem: {
            type: Object,
            default: () => {}
        },
        detailId: {
            type: String,
            default: ""
        },
        fileAuth: {
            type: Boolean,
            default: false
        }
    },
    data() {
        return {
            curVerObj: {},
            versionArr: [],
            isHistoryVerShow: false,
            deleteVisible: false,
            curDeleteFile: {},
            previewFile: {} // 当前要预览得文件
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
    watch: {
        curFileItem: {
            handler(obj) {
                if (obj && Object.keys(obj).length) {
                    this.getCurPreview(true); // true为打开弹窗后第一次获取当前文件
                    this.getHistoryList();
                }
            },
            immediate: true
        }
    },
    methods: {
        // 滚动事件,关闭打开得操作popover
        handleScroll(event) {
            // 在滚动时执行的操作
            document.body.click();
        },
        splitAndGetLast(str, separator) {
            if (!str || !separator) return;
            // 使用split方法按separator分割字符串，然后取最后一个元素
            return str.split(separator).slice(-1)[0].toLowerCase();
        },
        fileChange(file) {
            // 将file 入参上传
            this.uploadFile(file);
        },
        // 上传
        uploadFile(file) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                issue_id: this.detailId,
                tmpl_file: file.raw,
                group_id: this.curFileItem.group_id,
                is_current_version: 1
            };
            let formData = new FormData();
            Object.keys(params).forEach((key) => {
                formData.append(key, params[key]);
            });
            api.uploadFile(formData).then((res) => {
                if (res && res.resultCode === 200) {
                    this.getCurPreview(true);
                    this.getHistoryList();
                }
            });
        },
        // 获取当前文件
        getCurPreview(flag = false) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                issue_id: this.detailId,
                group_id: this.curFileItem.group_id,
                is_current_version: 1
            };
            api.getFileList(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    res.data.length
                ) {
                    this.curVerObj = res.data[0];
                    this.$emit("check-current-version", this.curVerObj);
                    if (flag) {
                        // 第一次打开弹窗或者上传文件时，预览当前得
                        this.$set(this.curVerObj, "actived", true);
                        // this.previewFile = this.curVerObj;
                        this.$emit("pre-view-file", this.curVerObj);
                    }
                } else {
                    this.curVerObj = {};
                }
            });
        },
        // 获取历史文件列表
        getHistoryList() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                issue_id: this.detailId,
                group_id: this.curFileItem.group_id,
                is_current_version: 0
            };
            api.getFileList(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    res.data.length
                ) {
                    this.versionArr = res.data;
                    this.isHistoryVerShow = true;
                } else {
                    this.versionArr = [];
                    this.isHistoryVerShow = false;
                }
            });
        },
        // 预览选中得历史文件
        previewVersion(item) {
            if (item.actived) return;
            this.versionArr.forEach((sub) => {
                if (item.id === sub.id) {
                    this.$set(sub, "actived", true);
                    // this.previewFile = item;
                    this.$emit("pre-view-file", item);
                } else {
                    this.$set(sub, "actived", false);
                }
            });
            this.curVerObj.actived = false;
        },
        // 预览当前文件，当选中历史文件后当前文件不被选中，然后可点击
        previewCurrent() {
            if (this.curVerObj.actived) return;
            this.curVerObj.actived = true;
            // this.previewFile = this.curVerObj;
            this.$emit("pre-view-file", this.curVerObj);
            this.versionArr.forEach((sub) => {
                this.$set(sub, "actived", false);
            });
        },
        // 设置历史文件为当前文件
        setCurrent(file) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                issue_id: this.detailId,
                group_id: this.curFileItem.group_id,
                is_current_version: 1,
                id: file.id
            };
            api.setCurrentFile(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.getCurPreview(true);
                    this.getHistoryList();
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
        // 删除
        deleteVersionFile(file) {
            this.curDeleteFile = _.cloneDeep(file);
            this.deleteVisible = true; // 打开确认弹窗
        },
        // 取消删除
        cancelDelVersion() {
            this.deleteVisible = false;
        },
        // 确认删除版本
        onConfirmDelVersion() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                issue_id: this.detailId,
                id: this.curDeleteFile.id,
                group_id: -1
            };
            let isPrevew = this.curDeleteFile.actived;
            api.deleteFile(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.$message({
                        showClose: true,
                        message: "删除成功",
                        type: "success"
                    });
                    this.cancelDelVersion();
                    if (isPrevew) {
                        // 如果删除得这条正好是正在预览得，删除成功后，将预览当前版本, 并且设置当前版本得选中样式
                        this.$set(this.curVerObj, "actived", true);
                        this.$emit("pre-view-file", this.curVerObj);
                        // this.$emit("check-current-version", this.curVerObj);
                    }
                    this.getHistoryList();
                    this.curDeleteFile = {}; //删除成功后清空
                }
            });
        },
        // 点击下载当前得图片
        downLoadCurrent(item) {
            document.body.click();
            this.downloadFile(item);
        },
        // 点击操作历史版本功能按钮
        operatHisVersion(item, type) {
            let mapping = {
                setting: "setCurrent",
                download: "downloadFile",
                delete: "deleteVersionFile"
            };
            // 调接口
            // this.$set(item, "visible", false);
            document.body.click();
            this[mapping[type]](item);
        }
    }
};
</script>
<style lang="scss" scoped>
.title {
    font-size: 16px;
    font-weight: 700;
    margin-bottom: 16px;
}
.flex-between {
    display: flex;
    justify-content: space-between;
}
.cur-title {
    height: 32px;
    line-height: 32px;
    .upload {
        color: #1890ff;
        display: flex;
        align-items: center;
        cursor: pointer;
        .upload-img {
            display: inline-block;
            width: 16px;
            height: 16px;
            background-size: 100% 100%;
            vertical-align: inherit;
            background-image: url("~@/assets/image/common/upload.png");
        }
    }
}
.his-title {
    font-weight: 700;
    height: 32px;
    line-height: 32px;
}
.his-box {
    overflow-y: auto;
    height: calc(100% - 183px - 10px);
    padding-right: 14px;
    margin-right: -20px;
}

.more-icon {
    display: inline-block;
    height: 16px;
    width: 16px;
    background-image: url(~@/assets/image/file_upload_preview/preview_dialog/more.png);
    background-size: 100% 100%;
    cursor: pointer;
    &:hover {
        background-image: url(~@/assets/image/file_upload_preview/preview_dialog/more_active.png);
        background-size: 100% 100%;
    }
}

.ver-box {
    box-sizing: border-box;
    height: 82px;
    border-radius: 8px;
    border: 1px solid #e6e9f0;
    margin-bottom: 23px;
    padding: 8px;
    cursor: pointer;
    .flex {
        display: flex;
        align-items: center;
        img {
            margin-right: 8px;
        }
        .sub-right {
            width: calc(100% - 56px);
            font-size: 12px;
        }
        .line {
            height: 23px;
            .name-size {
                display: inline-block;
                width: calc(100% - 20px);
            }
            .name {
                display: inline-block;
                max-width: calc(69.5% - 4px);
                margin-right: 4px;
                overflow: hidden;
                white-space: nowrap;
                text-overflow: ellipsis;
            }
            .size {
                display: inline-block;
                width: 30%;
                color: #a6abbc;
                overflow: hidden;
                white-space: nowrap;
                text-overflow: ellipsis;
            }
        }
        .creator {
            display: inline-block;
            max-width: 56px;
            margin-right: 8px;
            overflow: hidden;
            white-space: nowrap;
            text-overflow: ellipsis;
            font-size: 12px;
            color: #a6abbc;
        }
        .date {
            display: inline-block;
            width: calc(100% - 64px);
            overflow: hidden;
            white-space: nowrap;
            text-overflow: ellipsis;
            font-size: 12px;
            color: #a6abbc;
        }
        .ver {
            display: inline-block;
            height: 18px;
            line-height: 18px;
            color: #0c6c50;
            background-color: #e8eefa;
            padding: 0 4px;
            border-radius: 4px;
        }
    }

    &.actived {
        border: 1px solid #1890ff;
        background-color: #f0f8ff;
    }
    &:hover {
        border: 1px solid #1890ff;
    }
}
</style>
<style lang="scss">
.opera-box-popover.el-popper[x-placement^="bottom"] {
    margin-top: 2px;
}
.opera-box-popover.el-popper[x-placement^="top"] {
    margin-bottom: 2px;
}
.opera-box-popover.el-popover {
    min-width: 20px;
    padding: 8px;
    border-radius: 4px;
    background-color: #fff;
    box-shadow: 2px 2px 8px 1px rgba(47, 56, 76, 0.1);
    .pop-item {
        display: flex;
        align-items: center;
        height: 32px;
        width: 48px;
        padding: 0 8px;
        border-radius: 4px;
        cursor: pointer;
        &.his {
            width: 106px;
        }
        img {
            margin-right: 4px;
        }
        &:hover {
            background-color: #ecf5ff;
        }
    }
}
</style>
