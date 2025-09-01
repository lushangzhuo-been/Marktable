<template>
    <el-dialog
        :visible.sync="dialogVisible"
        class="preview-file basic-ui"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        @close="closeDialog"
    >
        <div slot="title">
            <div class="title-line">
                <span v-overflow class="file-name">{{
                    curPreviewFile.file_name
                }}</span>
                <span class="file-size">{{
                    curPreviewFile.file_size_desc
                }}</span>
            </div>
            <div class="title-line">
                <span class="creator">{{ curPreviewFile.creator }}</span>
                <span class="date">{{ curPreviewFile.created_at }}</span>
            </div>
        </div>
        <div class="preview-main">
            <!-- 预览模块 -->
            <preview
                class="preview"
                :style="{ width: mainWidth }"
                :curFileItem="curFileItem"
                :detailId="detailId"
                :curPreviewFile="curPreviewFile"
                :isClickClose="isClickClose"
            ></preview>
            <!-- 操作模块 -->
            <switch-version
                ref="switchVersion"
                class="version"
                v-show="versionSelected"
                :detailId="detailId"
                :fileAuth="fileAuth"
                :curFileItem="curFileItem"
                @check-current-version="checkCurrentVersion"
                @pre-view-file="preViewFile"
            >
                <b
                    slot="collapse"
                    class="coll-icon"
                    @click="versionSelected = false"
                ></b>
            </switch-version>
            <!-- 版本 -->
            <div class="right">
                <div
                    class="item-btn"
                    :class="{ actived: versionSelected }"
                    @click="versionSelected = !versionSelected"
                >
                    <img
                        :src="
                            versionSelected
                                ? require(`@/assets/image/file_upload_preview/preview_dialog/version_active.png`)
                                : require(`@/assets/image/file_upload_preview/preview_dialog/version.png`)
                        "
                        alt=""
                        width="24px"
                        height="24px"
                    />
                    <span>版本</span>
                </div>
                <div class="item-btn" @click="deleteAllVerion" v-if="fileAuth">
                    <img
                        :src="require(`@/assets/image/common/delete.svg`)"
                        alt=""
                        width="24px"
                        height="24px"
                    />
                    <span>删除</span>
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
    </el-dialog>
</template>
<script>
import _ from "lodash";
import Preview from "./preview.vue";
import SwitchVersion from "./switch_version.vue";
import api from "@/common/api/module/progress";
export default {
    components: {
        Preview,
        SwitchVersion
    },
    computed: {
        mainWidth() {
            return this.versionSelected
                ? "calc(100% - 80px - 260px)"
                : "calc(100% - 80px)";
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        }
    },
    data() {
        return {
            isClickClose: false,
            dialogVisible: false, //弹窗显示与否
            versionSelected: false, // 版本操作展开与否
            curFileItem: {}, // 当前打开得文件
            detailId: "",
            fileAuth: false,
            deleteVisible: false,
            newCurVersion: {}, //每次设置完新的都传过来
            curPreviewFile: {} //预览得文件
        };
    },
    methods: {
        // 每次设置新版本都当当前版本传过来
        checkCurrentVersion(obj) {
            this.newCurVersion = _.cloneDeep(obj);
        },
        // 选中得要预览得版本
        preViewFile(obj) {
            this.curPreviewFile = _.cloneDeep(obj);
        },
        // 打开预览弹窗
        show(item, detailId, fileAuth) {
            this.isClickClose = false;
            this.detailId = detailId;
            this.fileAuth = fileAuth;
            this.curFileItem = _.cloneDeep(item);
            // this.curPreviewFile = _.cloneDeep(item); // 第一次刚打开时候预览得是当前版本
            this.dialogVisible = true;
        },
        // 关闭预览弹窗
        closeDialog() {
            this.isClickClose = true;
            this.curFileItem = {};
            this.dialogVisible = false;
            this.versionSelected = false;
            this.deleteVisible = false;
            this.newCurVersion = {}; //每次设置完新的都传过来
            this.curPreviewFile = {}; //预览得文件
            this.$emit("closeDialog");
        },
        // 右侧删除全部
        deleteAllVerion() {
            this.deleteVisible = true;
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
                id: this.newCurVersion.id,
                group_id: this.newCurVersion.group_id,
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
                    this.closeDialog();
                }
            });
        }
    }
};
</script>
<style lang="scss" scoped>
.preview-file.basic-ui ::v-deep .el-dialog .el-dialog__header {
    padding: 12px 32px;
    .title-line {
        display: flex;
        align-items: center;
        width: 100%;
        height: 29px;
        line-height: 33px;
        .file-name {
            max-width: calc(100% - 170px);
            font-weight: 700;
            overflow: hidden;
            white-space: nowrap;
            text-overflow: ellipsis;
        }
        .file-size {
            font-size: 14px;
            color: #82889e;
            margin-left: 18px;
        }
        .creator,
        .date {
            font-size: 14px;
            color: #82889e;
            margin-right: 12px;
        }
    }
}

.preview-main {
    height: 100%;
    display: flex;
    width: 100%;
    .right {
        // width: 80px;
        height: 100%;
        box-sizing: border-box;
        background-color: #f5f6f8;
        // overflow-y: auto;
        padding-top: 10px;
        text-align: center;
        .item-btn {
            width: 64px;
            height: 64px;
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            margin: 0 8px 8px;
            border-radius: 4px;
            cursor: pointer;
            &.actived {
                color: #1890ff;
                background-color: #e9f0f8;
            }
            &:hover {
                background-color: #e9f0f8;
            }
        }
    }
    .version {
        position: relative;
        width: 300px;
        box-sizing: border-box;
        padding: 14px 20px;
        background-color: #fafafb;
        border-right: 1px solid #e6e9f0;
        .coll-icon {
            position: absolute;
            top: 13px;
            left: 0;
            display: inline-block;
            width: 12px;
            height: 24px;
            cursor: pointer;
            background-image: url(~@/assets/image/file_upload_preview/preview_dialog/collapse.png);
            background-size: 100% 100%;
            &:hover {
                background-image: url(~@/assets/image/file_upload_preview/preview_dialog/collapse_active.png);
                background-size: 100% 100%;
            }
        }
    }
    .preview {
        box-sizing: border-box;
        overflow-y: auto;
        padding: 20px 20px 0 20px;
        width: calc(100% - 382px);
    }
}
</style>
