<template>
    <div>
        <el-drawer
            :visible.sync="drawer"
            destroy-on-close
            direction="rtl"
            :size="'70%'"
            custom-class="progress-detail-drawer"
            :append-to-body="true"
            :wrapperClosable="false"
        >
            <div slot="title" class="title-slot">
                <DetailHead
                    moreOperation
                    :wsId="wsId"
                    :tempId="tempId"
                    :detailId="detailId"
                    :formLabelShow="formLabelShow"
                    :formDataShow="formDataShow"
                    :permissionEdit="permissionEdit"
                    :statusConfigList="statusConfigList"
                    @edit-form-item="editFormItem"
                    @status-change-info="statusChangeInfo"
                    @refresh-form="getFormData"
                    @delete-row="deleteRow"
                ></DetailHead>
            </div>
            <!-- 切换tab-详情/进展/日志 -->
            <div class="page-body">
                <!-- 左tabs -->
                <div class="body-left">
                    <el-tabs class="tabs detail-tabs" v-model="leftActiveName">
                        <el-tab-pane
                            v-for="(item, index) in leftTabsList"
                            :label="item.tab_cn"
                            :name="item.tab"
                            :key="index"
                        >
                            <span slot="label">
                                {{ item.tab_cn
                                }}<span v-show="item.tab === 'file'">{{
                                    "(" + fileCount + ")"
                                }}</span
                                ><span v-show="item.tab === 'sub_tmpl'">{{
                                    "(" + subTaskCount + ")"
                                }}</span>
                            </span>
                        </el-tab-pane>
                    </el-tabs>
                    <!-- 详情/进展/日志 -->
                    <div class="block-content">
                        <div v-show="leftActiveName === 'tmpl_detail'">
                            <!-- form详情 -->
                            <DetailForm
                                :wsId="wsId"
                                :tempId="tempId"
                                :detailId="detailId"
                                :formLabelShow="formLabelShow"
                                :formDataShow="formDataShow"
                                @edit-form-item="editFormItem"
                            ></DetailForm>
                        </div>
                        <div v-show="leftActiveName === 'file'">
                            <file-upload
                                ref="FileUpload"
                                :detailId="detailId"
                                :permissionEdit="permissionEdit"
                                @transmit-file-count="transmitFileCount"
                            ></file-upload>
                        </div>
                        <div v-show="leftActiveName === 'sub_tmpl'">
                            <sub-task-content
                                ref="subTask"
                                :detailId="detailId"
                                :subTaskConfig="subTaskConfig"
                                @refresh-sub-task-count="getSubTaskCount"
                            ></sub-task-content>
                        </div>
                    </div>
                </div>
                <!-- 右进展 -->
                <div class="body-right">
                    <!-- 标题 -->
                    <el-tabs
                        class="tabs detail-tabs progress"
                        v-model="rightActiveName"
                    >
                        <el-tab-pane
                            v-for="(item, index) in rightTabsList"
                            :label="item.label"
                            :name="item.name"
                            :key="index"
                        >
                            <span slot="label">
                                {{ item.label
                                }}<span v-show="item.name === 'progress'">{{
                                    "(" + evolveCount + ")"
                                }}</span>
                            </span>
                        </el-tab-pane>
                    </el-tabs>
                    <div
                        class="evolve-content"
                        v-if="rightActiveName === 'progress'"
                    >
                        <evolve
                            ref="Evolve"
                            :detailId="detailId"
                            @transmit-evolve-count="transmitEvolveCount"
                        ></evolve>
                    </div>
                    <div class="log-content" v-if="rightActiveName === 'log'">
                        <log
                            ref="Logs"
                            :logs="logs"
                            :count="logsCount"
                            :detailId="detailId"
                        ></log>
                    </div>
                </div>
            </div>
        </el-drawer>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import configApi from "@/common/api/module/progress_setting";
import DetailHead from "./detail_head";
import DetailForm from "./detail_form";
import FileUpload from "./components/file_upload";
import Evolve from "@/components/evolve/index";
import Log from "@/components/log/index";
import subTaskContent from "./components/sub_task_content.vue";
export default {
    components: {
        DetailHead,
        DetailForm,
        Log,
        Evolve,
        FileUpload,
        subTaskContent
    },
    props: {},
    data() {
        return {
            formLabel: [],
            formData: {},
            drawer: false,
            leftActiveName: "",
            leftTabsList: [],
            rightTabsList: [
                {
                    label: "进展",
                    name: "progress"
                },
                {
                    label: "日志",
                    name: "log"
                }
            ],
            rightActiveName: "", // 右侧 进展/日志
            formLabelShow: [],
            formDataShow: {},
            detailId: "",
            logs: [],
            logsCount: 0,
            fileCount: 0,
            subTaskCount: 0,
            evolveCount: 0,
            permissionEdit: false, //表单编辑权限，下穿附件上传
            statusConfigList: [],
            subTaskConfig: [],
            cardEditInfo: {
                defaultStatus: null,
                status: {
                    start: null,
                    end: null
                },
                other: false
            }
        };
    },
    watch: {
        drawer: {
            handler(boolean) {
                if (!boolean) {
                    this.logsCount = 0;
                    this.fileCount = 0;
                    this.subTaskCount = 0;
                    this.evolveCount = 0;
                    this.leftActiveName = "";
                    this.rightActiveName = "";
                    this.$emit("refresh", this.cardEditInfo);
                    this.$refs.Evolve.clearEvolveShow();
                    // 重置 config form
                    this.formLabel = [];
                    this.formLabelShow = [];
                    this.formData = {};
                    this.formDataShow = {};
                    this.permissionEdit = false;
                    this.cardEditInfo = {
                        defaultStatus: null,
                        status: {
                            start: null,
                            end: null
                        },
                        other: false
                    };
                }
            }
        },
        leftActiveName: {
            handler(type) {
                // 请求刷新对应数据
                this.subTaskConfig = [];
                if (type === "tmpl_detail") {
                    this.getFormConfig();
                }
                if (type === "file") {
                    this.$refs.FileUpload.getFileAuth(); //获取文件权限
                    this.$refs.FileUpload.getFileList(); // 获取文件列表
                    this.$refs.FileUpload.downloadTypeCheck();
                }
                if (type === "sub_tmpl") {
                    this.$refs.subTask.getSubTmplAuth(); //获取子任务的增删权限
                    // 子任务
                    this.getSubTaskConfig();
                }
            }
        },
        rightActiveName: {
            handler(type) {
                if (type === "progress") {
                    this.$nextTick(() => {
                        this.$refs.Evolve.getEvolveDetail();
                    });
                }
                if (type === "log") {
                    this.$nextTick(() => {
                        // 重置
                        this.$refs.Logs.logsPage = {
                            size: 10,
                            page: 1
                        };
                        this.$refs.Logs.logs = [];
                        this.$refs.Logs.activeName = [];
                        this.$refs.Logs.isShowNoData = true;
                        this.$refs.Logs.loading = true;
                        this.$refs.Logs.imgShow = false;
                        this.$refs.Logs.getLog();
                    });
                }
            }
        }
    },
    computed: {
        wsId() {
            return this.$store.state.curSpace.id || "";
        },
        tempId() {
            return this.$route.params.id || "";
        }
    },
    methods: {
        // 子任务配置
        getSubTaskConfig() {
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId
            };
            configApi.getSubWorkList(params).then((res) => {
                if (res.resultCode === 200 && res.data) {
                    this.subTaskConfig = res.data || [];
                } else {
                    this.subTaskConfig = [];
                }
            });
        },
        // 删除成功后
        deleteRow() {
            this.cardEditInfo.other = true;
            this.drawer = false;
        },
        openDrawer(type, row) {
            // type  分为详情form 日志
            this.drawer = true;
            this.detailId = row._id;
            this.leftActiveName = type;
            this.cardEditInfo.defaultStatus = row.status.id;
            this.rightActiveName = "progress";
            this.getLeftTabMenu();
            this.getStatusMenusList();
            this.getSubTaskCount();
            this.$nextTick(() => {
                this.$refs.Evolve.fetAuthProgress();
                // 获取进展附件数量
                this.$refs.FileUpload.getFileList();
            });
        },
        // 左侧tab配置
        getLeftTabMenu() {
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId
            };
            configApi.getTabList(params).then((res) => {
                if (res.resultCode === 200 && res.data) {
                    this.leftTabsList = res.data || [];
                } else {
                    this.leftTabsList = [];
                }
            });
        },
        // 获取状态下拉列表
        getStatusMenusList() {
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
                    this.statusConfigList = res.data;
                } else {
                    this.statusConfigList = [];
                }
            });
        },
        // 获取项配置
        getFormConfig() {
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId,
                module: "detail"
            };
            api.getProgressColConfig(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    this.formLabel = res.data;
                    this.formLabelShow = _.cloneDeep(this.formLabel);
                    this.getFormData();
                } else {
                    this.formLabel = [];
                    this.formLabelShow = [];
                }
            });
        },
        // 获取数据
        getFormData() {
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId,
                id: this.detailId
            };
            api.getFormData(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    this.formData = res.data;
                    this.formDataShow = _.cloneDeep(this.formData);
                    this.permissionEdit =
                        res.data.permission &&
                        res.data.permission.edit === "yes";
                } else {
                    this.formData = {};
                    this.formDataShow = {};
                }
            });
        },
        // 提交编辑
        editFormItem(res, key, mode) {
            // 统计变更内容，按变更内容刷新卡片列表
            if (key === "status") {
                // 状态
                this.getFormData();
            } else {
                this.cardEditInfo.other = true;
                let file_key = key;
                let params = {
                    ws_id: this.wsId,
                    tmpl_id: this.tempId,
                    id: this.detailId,
                    field_key: key,
                    [file_key]: res
                };
                api.updateProgress(params).then((res) => {
                    if (res && res.resultCode === 200) {
                        // 刷新详情 刷新列表
                        this.getFormData();
                    } else {
                        this.getFormData();
                    }
                });
            }
        },
        statusChangeInfo(info) {
            if (!this.cardEditInfo.status.start) {
                this.cardEditInfo.status.start = info.start_status_id;
            }
            this.cardEditInfo.status.end = info.end_status_id;
        },
        refreshForm() {
            this.getFormData();
        },
        refresh() {
            this.$emit("refresh");
        },
        transmitFileCount(count) {
            this.fileCount = count;
        },
        // 查询子任务数量
        getSubTaskCount() {
            let params = {
                obj_id: this.detailId,
                ws_id: this.wsId,
                tmpl_id: this.tempId
            };
            api.getSubTaskCount(params).then((res) => {
                if (res.resultCode === 200 && res.data) {
                    this.subTaskCount = res.data || 0;
                } else {
                    this.subTaskCount = 0;
                }
            });
        },
        transmitEvolveCount(count) {
            this.evolveCount = count;
        }
    }
};
</script>

<style lang="scss" scoped>
@import "@/components/table/p_table/style.scss";
@import "@/pages/progress/detail/draw_and_single_route.scss";
@import "@/pages/progress/detail/single_route/evolve.scss";
</style>
<style lang="scss">
@import "@/pages/progress/detail/draw_and_single_route_drill.scss"; // @import "@/components/table/p_table/drill_style.scss";
</style>
