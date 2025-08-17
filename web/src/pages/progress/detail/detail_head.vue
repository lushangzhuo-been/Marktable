<template>
    <!-- 详情及单路由 共用form部分抽取 -->
    <div>
        <!-- 标题及按钮 -->
        <div v-if="formLabelShow.length" class="drawer-title">
            <div class="left">
                <!-- <el-button
                    v-if="jumpBack"
                    class="basic-ui back-btn"
                    size="small"
                    @click="backJump"
                >
                    <b class="back-box"></b>
                    返回
                </el-button> -->
                <b v-if="jumpBack" class="back-box" @click="backJump"></b>
                <!-- 标题 -->
                <div class="title">
                    <title-item
                        :formItem="getConfigItem(formLabelShow, 'title')"
                        :formData="formDataShow"
                        @edit-form-item="editFormItem"
                    ></title-item>
                </div>
            </div>
            <div class="right">
                <el-button
                    v-if="false"
                    class="basic-ui head-btn"
                    size="small"
                    @click="toUrging"
                >
                    <b class="urging"></b> 催办
                </el-button>
                <!-- 拓展更多操作 -->
                <div v-if="moreOperation" class="more-operation">
                    <el-dropdown
                        placement="bottom-end"
                        trigger="click"
                        class="operation-block-dropdown basic-ui"
                        width="200"
                        @command="moreOperationAction"
                        @visible-change="moreDropShow"
                    >
                        <el-dropdown-menu slot="dropdown">
                            <el-dropdown-item
                                class="basic-ui height32"
                                command="share"
                            >
                                <div>
                                    <b class="share-box"></b>
                                    复制链接
                                </div>
                            </el-dropdown-item>
                            <el-dropdown-item
                                class="basic-ui height32"
                                :class="hasDeleteAuth ? '' : 'disabled'"
                                command="delete"
                            >
                                <div>
                                    <b class="delete-box"></b>
                                    删除
                                </div>
                            </el-dropdown-item>
                        </el-dropdown-menu>
                        <b
                            class="tabs-partition"
                            :class="{ active: moreOperating }"
                        ></b>
                    </el-dropdown>
                </div>
                <!-- 遍历操作按钮 -->
                <!-- <el-button
                    v-for="(btnItem, btnIndex) in formDataShow.buttons"
                    :key="btnIndex"
                    class="basic-ui head-btn"
                    size="small"
                    type="primary"
                    :disabled="!permissionEdit"
                    @click="operationNode(btnItem)"
                >
                    {{ btnItem.name }}
                </el-button> -->
                <!-- 分享 -->
                <!-- <div v-if="moreOperation" class="share-block">
                    <el-tooltip
                        class="item"
                        effect="dark"
                        placement="top-start"
                        :popper-class="`tips-popper`"
                    >
                        <div slot="content">复制任务链接</div>
                        <b class="share-box" @click="shareTaskDetail"></b>
                    </el-tooltip>
                </div> -->
            </div>
        </div>
        <!-- 其他信息 -->
        <div v-if="formLabelShow.length" class="other-msg">
            <div style="display: flex; margin-left: 12px">
                <!-- 状态 -->
                <div
                    class="msg-item status-item"
                    style="max-width: calc(100% - 404px)"
                >
                    <span class="msg-item-title">状态：</span>
                    <template
                        v-if="
                            formDataShow.status &&
                            Object.keys(formDataShow.status).length
                        "
                    >
                        <status-item
                            :formItem="getConfigItem(formLabelShow, 'status')"
                            :formData="formDataShow"
                            @edit-form-item="editFormItem"
                            @status-change-info="statusChangeInfo"
                        ></status-item>
                    </template>
                    <span v-else>--</span>
                </div>
                <!-- 处理人 -->
                <div class="msg-item handler" style="width: 50%">
                    <span class="msg-item-title"
                        >{{
                            getConfigItem(formLabelShow, "handler").name
                        }}：</span
                    >
                    <!-- 分单多人 -->
                    <single-people-item
                        class="title-handler"
                        v-if="
                            getConfigItem(formLabelShow, 'handler').expr ===
                            'single'
                        "
                        :formItem="getConfigItem(formLabelShow, 'handler')"
                        :formData="formDataShow"
                        @edit-form-item="editFormItem"
                    ></single-people-item>
                    <multiple-people-item
                        class="title-handler"
                        v-if="
                            getConfigItem(formLabelShow, 'handler').expr ===
                            'multiple'
                        "
                        :formItem="getConfigItem(formLabelShow, 'handler')"
                        :formData="formDataShow"
                        @edit-form-item="editFormItem"
                    ></multiple-people-item>
                </div>
            </div>
            <div class="row first">
                <!-- 创建人 -->
                <div class="msg-item creator">
                    <span class="msg-item-title">创建人：</span>
                    <span
                        v-if="
                            formDataShow.creator && formDataShow.creator.length
                        "
                        class="msg-item-content"
                    >
                        <user-message :userMessage="formDataShow.creator[0]">
                            <span>
                                <el-avatar
                                    class="progress-avatar"
                                    icon="el-icon-user-solid"
                                    size="small"
                                    :src="
                                        getAvatar(
                                            formDataShow.creator[0].avatar
                                        )
                                    "
                                ></el-avatar>
                                {{ formDataShow.creator[0].full_name || "--" }}
                            </span>
                        </user-message>
                    </span>
                    <span v-else>--</span>
                </div>
                <!-- 创建时间 -->
                <div class="msg-item">
                    <span class="msg-item-title">创建时间：</span>
                    <span class="msg-item-content">
                        {{ formDataShow.created_at || "--" }}
                    </span>
                </div>
                <!-- 最后修改时间 -->
                <div class="msg-item">
                    <span class="msg-item-title">最后修改时间：</span>
                    <span class="msg-item-content">
                        {{ formDataShow.updated_at || "--" }}
                    </span>
                </div>
            </div>
        </div>
        <node-operation-dialog
            ref="NodeOperationDialog"
            @confirm-next-state="confirmNextState"
            :nodeOperationConfig="nodeOperationConfig"
            :nodeOperationData="nodeOperationData"
        ></node-operation-dialog>
        <!-- 删除表格行提醒 -->
        <delete-table-row-dialog
            ref="deleteTableRowDialog"
            @cancel="cancelDeleteTableRow"
            @on-confirm="comfirmDeleteTableRow"
        ></delete-table-row-dialog>
    </div>
</template>

<script>
import _ from "lodash";
import { rgbToRgba } from "@/assets/tool/func";
import { FieldType } from "@/assets/tool/const";
import TitleItem from "./form_item/title_item/index.vue";
import SinglePeopleItem from "./form_item/single_people_item/index2.vue";
import MultiplePeopleItem from "./form_item/multiple_people_item/index2.vue";
import StatusItem from "./form_item/status_item/index.vue";
import UserMessage from "@/components/user_message_tip";
import { imgHost } from "@/assets/tool/const";
import api from "@/common/api/module/progress";
import NodeOperationDialog from "./components/node_operation_dialog";
// 删除表格行提醒弹窗
import DeleteTableRowDialog from "@/pages/progress/components/delete_table_row_dialog.vue";
export default {
    components: {
        TitleItem,
        SinglePeopleItem,
        MultiplePeopleItem,
        UserMessage,
        NodeOperationDialog,
        StatusItem,
        DeleteTableRowDialog
    },
    props: {
        wsId: {
            type: [String, Number],
            default: ""
        },
        tempId: {
            type: [String, Number],
            default: ""
        },
        detailId: {
            type: [String, Number],
            default: ""
        },
        formLabelShow: {
            type: Array,
            default: () => []
        },
        formDataShow: {
            type: Object,
            default: () => {}
        },
        permissionEdit: {
            type: Boolean,
            default: false
        },
        moreOperation: {
            type: Boolean,
            default: false
        },
        jumpBack: {
            type: Boolean,
            default: false
        }
        // optionsList: {
        //     type: Array,
        //     default: () => []
        // }
    },
    data() {
        return {
            nodeOperationConfig: [],
            nodeOperationData: {},
            moreOperating: false,
            searchValue: "",
            optionsList: [],
            tmpOptionsList: [],
            hasDeleteAuth: false
        };
    },
    watch: {
        searchValue(str) {
            if (str.trim()) {
                this.optionsList = this.tmpOptionsList.filter((item) => {
                    return (
                        item.end_status_name
                            .toLocaleUpperCase()
                            .indexOf(str.toLocaleUpperCase()) > -1
                    );
                });
            } else {
                this.optionsList = this.tmpOptionsList;
            }
        }
    },
    methods: {
        rgbToRgba(color, opacity) {
            return rgbToRgba(color, opacity);
        },
        getStatusList() {
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId,
                id: this.detailId
            };
            api.getStatusMenusList(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    res.data.length
                ) {
                    this.optionsList = res.data;
                    this.tmpOptionsList = _.cloneDeep(this.optionsList);
                } else {
                    this.optionsList = [];
                    this.tmpOptionsList = [];
                }
            });
        },
        selectStauts(item) {
            document.body.click();
        },
        confirmNextState(btn, form) {
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId,
                step_id: btn.id,
                id: this.detailId,
                // 解构被编辑的参数
                ...form
            };
            // return;
            api.switchStepAction(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.$emit("refresh-form");
                    this.$refs.NodeOperationDialog.closeDialog();
                }
            });
        },
        // 操作流程按钮
        operationNode(btn) {
            // this.$emit("operation-node", btn);
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId,
                step_id: btn.id,
                id: this.detailId
            };
            api.getBtnScreen(params).then((res) => {
                if (res && res.resultCode === 200) {
                    if (!res.data || (res.data && !res.data.length)) {
                        // 直接放行
                        this.switchStepAction(btn);
                    } else {
                        this.nodeOperationConfig = res.data;
                        // 获取节点配置数据
                        this.$refs.NodeOperationDialog.openDialog(btn);
                        document.body.click();
                        this.nodeOperationData = this.formDataShow;
                        this.getNodeOperationData(btn);
                    }
                }
            });
        },
        switchStepAction(btn) {
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId,
                step_id: btn.id,
                id: this.detailId
            };
            api.switchStepAction(params).then((res) => {
                if (res && res.resultCode === 200) {
                    document.body.click();
                    this.$emit("refresh-form");
                }
            });
        },
        getNodeOperationData(btn) {
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId,
                step_id: btn.id,
                id: this.detailId
            };
            api.getFormData(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    this.nodeOperationData = res.data;
                }
            });
        },
        // 其他小功能
        shareTaskDetail() {
            let jumpSite = window.location.origin;
            jumpSite += `/#/task/detail/${this.wsId}/${this.tempId}/${this.detailId}`;
            let tempTextarea = document.createElement("textarea");
            // 设置textarea的value为当前网址
            tempTextarea.value = jumpSite;
            // 将textarea添加到DOM中
            document.body.appendChild(tempTextarea);
            // 选中textarea中的文本
            tempTextarea.select();
            // 复制选中的文本
            document.execCommand("copy");
            // 移除临时的textarea元素
            document.body.removeChild(tempTextarea);
            this.$message({
                showClose: true,
                message: "链接已复制",
                type: "success"
            });
        },
        backJump() {
            this.$router.push({
                path: `/progress/${this.wsId}/${this.tempId}`
            });
        },
        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        },
        getConfigItem(arr, key) {
            let obj = _.find(arr, { field_key: key });
            return obj;
        },
        getType(type, expr) {
            if (type && !expr) {
                return {
                    "background-image": `url(${FieldType[type]})`
                };
            } else {
                return {
                    "background-image": `url(${FieldType[type + "_" + expr]})`
                };
            }
        },
        // 提交编辑
        editFormItem(res, key, mode) {
            this.$emit("edit-form-item", res, key, mode);
        },
        statusChangeInfo(info) {
            this.$emit("status-change-info", info);
        },
        moreOperationAction(command) {
            if (command === "share") {
                this.shareTaskDetail();
            } else if (command === "delete") {
                if (!this.hasDeleteAuth) return;
                this.deleteItem();
            }
        },
        deleteItem() {
            this.$refs.deleteTableRowDialog.openDialog(this.formDataShow);
        },
        // 取消删除表格行
        cancelDeleteTableRow() {
            document.body.click();
        },
        // 确认删除表格行调接口
        comfirmDeleteTableRow() {
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId,
                ids: this.formDataShow._id
            };
            api.deleteProgressItems(params).then((res) => {
                if (this.$route.name !== "progress") return;
                if (res && res.resultCode === 200) {
                    this.$refs.deleteTableRowDialog.cancel();
                    this.$message({
                        showClose: true,
                        message: "删除成功",
                        type: "success"
                    });
                    document.body.click();
                    this.$emit("delete-row");
                }
            });
        },
        moreDropShow(boolean) {
            if (boolean) {
                this.fetAuthDelete();
                this.moreOperating = true;
            } else {
                this.moreOperating = false;
            }
        },
        fetAuthDelete(row) {
            // 获取进展权限
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: row._id,
                auth_mode: "delete"
            };
            api.getUserAuth(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.hasDeleteAuth = res.data;
                } else {
                    this.hasDeleteAuth = false;
                }
                this.$set(row, "isActived", true);
            });
        }
    }
};
</script>

<style lang="scss" scoped>
@import "@/components/table/p_table/style.scss";
@import "@/pages/progress/detail/draw_and_single_route.scss";
.more-operation {
    display: inline-block;
    margin-right: 16px;
    .tabs-partition {
        display: inline-block;
        width: 24px;
        height: 24px;
        vertical-align: sub;
        background-image: url(@/assets/image/progress/operation.png);
        background-size: 100% 100%;
        &:hover {
            background-image: url(@/assets/image/progress/operation-active.png);
        }
        &.active {
            background-image: url(@/assets/image/progress/operation-active.png);
        }
    }
}
.head-btn {
    vertical-align: text-top;
}
</style>
<style lang="scss">
@import "@/pages/progress/detail/draw_and_single_route_drill.scss"; // @import "@/components/table/p_table/drill_style.scss";
</style>
