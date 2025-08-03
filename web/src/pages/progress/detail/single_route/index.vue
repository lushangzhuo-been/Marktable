<template>
    <div class="window-page">
        <!-- 详情单路由 -->
        <div
            v-if="linkEffective"
            class="page-content"
            v-loading="pageLoading"
            element-loading-text="加载中"
            element-loading-background="rgba(255, 255, 255)"
        >
            <!-- 头 -->
            <div class="page-head single-detail">
                <DetailHead
                    jumpBack
                    :wsId="wsId"
                    :tempId="tempId"
                    :detailId="detailId"
                    :formLabelShow="formLabelShow"
                    :formDataShow="formDataShow"
                    :permissionEdit="permissionEdit"
                    @edit-form-item="editFormItem"
                    @refresh-form="getFormData"
                ></DetailHead>
            </div>
            <!-- 身子 -->
            <div class="page-body single-detail">
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
                        <div v-if="leftActiveName === 'sub_tmpl'">
                            <sub-task-content
                                :detailId="detailId"
                                :subTaskConfig="subTaskConfig"
                            ></sub-task-content>
                        </div>
                    </div>
                    <!-- <el-tabs class="tabs detail-tabs" v-model="activeName">
                        <el-tab-pane
                            v-for="(item, index) in tabsList"
                            :label="item.label"
                            :name="item.name"
                            :key="index"
                        >
                            <span slot="label">
                                {{ item.label
                                }}<span v-show="item.name === 'file'">{{
                                    "(" + fileCount + ")"
                                }}</span>
                            </span>
                        </el-tab-pane>
                    </el-tabs>
                    <div class="block-content">
                        <div v-show="activeName === 'detail'">
                            <DetailForm
                                :wsId="wsId"
                                :tempId="tempId"
                                :detailId="detailId"
                                :formLabelShow="formLabelShow"
                                :formDataShow="formDataShow"
                                @edit-form-item="editFormItem"
                            ></DetailForm>
                        </div>
                        <div v-show="activeName === 'file'">
                            <file-upload
                                ref="FileUpload"
                                :detailId="detailId"
                                :permissionEdit="permissionEdit"
                                @transmit-file-count="transmitFileCount"
                            ></file-upload>
                        </div>
                        <div v-show="activeName === 'log'" class="logs-content">
                            <log
                                ref="Logs"
                                :logs="logs"
                                :count="logsCount"
                                :detailId="detailId"
                            ></log>
                        </div>
                    </div> -->
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
                <!-- <div class="body-right">

                    <div class="evolve-title">进展({{ evolveCount }})</div>
                    <div class="evolve-content">
                        <evolve
                            ref="Evolve"
                            :detailId="detailId"
                            @transmit-evolve-count="transmitEvolveCount"
                        ></evolve>
                    </div>
                </div> -->
            </div>
        </div>
        <!-- 失效 -->
        <div v-else class="invalid-msg">
            <div class="invalid-border">
                <b class="invalid-box"></b>
            </div>
            <div class="invalid-text">此任务链接已被删除或者已失效</div>
            <div class="invalid-btn">
                <el-button
                    class="basic-ui"
                    type="primary"
                    size="small"
                    @click="jumpToHome"
                >
                    返回首页
                </el-button>
            </div>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import DetailHead from "../detail_head";
import DetailForm from "../detail_form";
import Log from "@/components/log/index";
import Evolve from "@/components/evolve/index";
import FileUpload from "../components/file_upload";
import configApi from "@/common/api/module/progress_setting";
// src\pages\progress\detail\components\sub_task_content.vue
import subTaskContent from "@/pages/progress/detail/components/sub_task_content.vue";
export default {
    components: {
        DetailHead,
        DetailForm,
        subTaskContent,
        Log,
        Evolve,
        FileUpload
    },
    props: {},
    data() {
        return {
            formLabel: [],
            formData: {},
            leftActiveName: "tmpl_detail",
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
            nodeOperationConfig: [],
            nodeOperationData: {},
            logs: [],
            logsCount: 0,
            fileCount: 0,
            evolveCount: 0,
            permissionEdit: false, //表单编辑权限，下穿附件上传
            wsId: null,
            tempId: null,
            detailId: null,
            pageLoading: true,
            linkEffective: true,
            subTaskConfig: []
        };
    },
    watch: {
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
        },
        "$route.params.wsid"() {
            this.loadView();
        },
        "$route.params.id"() {
            this.loadView();
        },
        "$route.params.taskid"() {
            this.loadView();
        }
    },
    computed: {},
    mounted() {
        this.loadView();
    },
    methods: {
        // 左侧tab配置
        getLeftTabMenu() {
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId
            };
            configApi.getTabList(params).then((res) => {
                if (res.resultCode === 200 && res.data) {
                    this.leftTabsList = res.data || [];
                    this.leftActiveName = "tmpl_detail";
                    this.rightActiveName = "progress";
                } else {
                    this.leftTabsList = [];
                }
            });
        },
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
        loadView() {
            this.wsId = this.$route.params.wsid;
            this.tempId = this.$route.params.id;
            this.detailId = this.$route.params.taskid;
            this.$nextTick(() => {
                if (this.wsId && this.tempId && this.detailId) {
                    let obj = {
                        id: this.wsId
                    };
                    this.$store.commit("setCurSpace", obj);
                    // 调接口
                    this.getLeftTabMenu();
                    this.getFormConfig();
                    // this.$refs.Evolve.getEvolveDetail();
                    // this.$refs.FileUpload.getFileList();
                } else {
                    this.linkEffective = false;
                }
            });
        },

        jumpToHome() {
            this.$router.push({
                name: "home"
            });
        },
        jumpBack() {
            this.$router.push({
                path: `/progress/${this.tempId}`
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
                    this.linkEffective = false;
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
                    this.linkEffective = false;
                }
                this.pageLoading = false;
            });
        },
        // 提交编辑
        editFormItem(res, key) {
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
                }
            });
        },
        transmitFileCount(count) {
            this.fileCount = count;
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
.page-content {
    box-sizing: border-box;
    // padding: 16px 0 0;
    height: 100vh;
}
.window-page {
    width: 100%;
    height: 100%;
    position: relative;
    background-color: #fafafb;
    .invalid-msg {
        box-sizing: border-box;
        width: 480px;
        height: 304px;
        padding: 26px 26px 32px;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        text-align: center;

        box-shadow: 0px 0px 16px 1px rgba(0, 0, 0, 0.1);
        border-radius: 4px 4px 4px 4px;
        .invalid-border {
            .invalid-box {
                display: inline-block;
                width: 182px;
                height: 136px;
                background-image: url("@/assets/image/state_picture/invalid.svg");
                background-size: 100% 100%;
            }
        }
        .invalid-text {
            font-size: 16px;
            color: #5c6479;
            line-height: 60px;
        }
    }
}
</style>
<style lang="scss">
@import "@/pages/progress/detail/draw_and_single_route_drill.scss";
</style>
