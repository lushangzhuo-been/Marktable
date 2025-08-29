<template>
    <div>
        <div class="title-block">
            <div class="title" @click="changeExpand">
                <span class="title-content">
                    {{ subTaskInfo.sub_tmpl_name }}
                </span>
                <b
                    class="down-box"
                    :class="{
                        expand: expand,
                        close: !expand
                    }"
                ></b>
            </div>
            <div
                v-if="subTmplAuth"
                class="add-sub-task"
                :class="subTmplAuth ? '' : 'disabled'"
                @click="addSubTask"
            >
                <b class="plus-box"></b>
                添加子任务
            </div>
        </div>
        <div
            class="table-content"
            :class="{
                expand: expand,
                close: !expand
            }"
        >
            <sub-task-table
                :data="tableData"
                :subTmplAuth="subTmplAuth"
                @delete-sub-task="deleteSubTask"
                @open-detail="openSubTaskDetail"
            ></sub-task-table>
            <el-pagination
                v-show="page.count > 10"
                class="basic-ui"
                @current-change="currentChange"
                layout="total, prev, pager, next"
                :current-page="page.page_num"
                :page-size="page.page_size"
                :total="page.count"
            >
            </el-pagination>
        </div>
        <delete-table-row-dialog
            ref="deleteTableRowDialog"
            @on-confirm="comfirmDeleteTableRow"
        ></delete-table-row-dialog>
        <add-progress
            ref="AddProgress"
            from="subTask"
            :detailId="detailId"
            :formLabel="createCol"
            :subTaskInfo="subTaskInfo"
            @refresh="refreshTable"
        ></add-progress>
    </div>
</template>

<script>
import SubTaskTable from "@/components/table/sub_task_table/sub_task_table";
import { baseMixin } from "@/mixin.js";
import api from "@/common/api/module/progress";
// src\
import AddProgress from "@/pages/progress/add_progress/index.vue";
import DeleteTableRowDialog from "@/pages/progress/components/delete_table_row_dialog.vue";
export default {
    mixins: [baseMixin],
    components: {
        SubTaskTable,
        AddProgress,
        DeleteTableRowDialog
    },
    props: {
        subTaskInfo: {
            type: Object,
            default: () => {}
        },
        detailId: {
            type: String,
            default: ""
        },
        subTmplAuth: {
            type: Boolean,
            default: false
        }
    },
    watch: {
        subTaskInfo: {
            handler(obj) {
                this.getListData();
            },
            immediate: true
        }
    },
    data() {
        return {
            tableData: [],
            page: {
                page_num: 1,
                page_size: 10,
                count: 0
            },
            tableLoading: true,
            createCol: [],
            expand: true
        };
    },
    methods: {
        // 跳子任务详情
        openSubTaskDetail(scope) {
            let rowInfo = scope.row;
            let origin = window.location.origin;
            let newUrl =
                origin +
                "/#/" +
                "task/detail/" +
                rowInfo.ws_id +
                "/" +
                rowInfo.tmpl_id +
                "/" +
                rowInfo._id;
            window.open(newUrl, "_blank");
        },
        changeExpand() {
            this.expand = !this.expand;
        },
        addSubTask() {
            this.getAddProgressConfig();
            this.$refs.AddProgress.openDialog();
        },
        getAddProgressConfig() {
            let params = {
                ws_id: this.subTaskInfo.ws_id,
                tmpl_id: this.subTaskInfo.sub_tmpl_id,
                module: "create"
            };
            api.getProgressColConfig(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    Object.keys(res.data).length
                ) {
                    this.createCol = res.data;
                }
            });
        },
        deleteSubTask(row) {
            this.$refs.deleteTableRowDialog.openDialog(row);
        },
        comfirmDeleteTableRow(row) {
            let params = {
                parent_tmpl_id: this.subTaskInfo.tmpl_id,
                obj_id: this.detailId,
                ws_id: this.subTaskInfo.ws_id,
                tmpl_id: this.subTaskInfo.sub_tmpl_id,
                ids: row._id
            };
            api.deleteSubAction(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.$refs.deleteTableRowDialog.cancel();
                    this.$message({
                        showClose: true,
                        message: "删除成功",
                        type: "success"
                    });
                    document.body.click();
                    this.getListData();
                    this.$emit("refresh-sub-task-count");
                }
            });
        },
        refreshTable() {
            this.page.page_num = 1;
            this.getListData();
            // 更新子任务总数
            this.$emit("refresh-sub-task-count");
        },
        getListData() {
            let params = {
                parent_tmpl_id: this.subTaskInfo.tmpl_id,
                obj_id: this.detailId,
                ws_id: this.subTaskInfo.ws_id,
                tmpl_id: this.subTaskInfo.sub_tmpl_id,
                page_num: this.page.page_num,
                page_size: this.page.page_size
            };
            api.getProgressListData(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.tableData = res.data.list || [];
                    this.page.count = res.data.cnt || 0;
                } else {
                    this.tableData = [];
                    this.page.count = 0;
                }
                this.tableLoading = false;
            });
        },
        currentChange(page) {
            this.tableLoading = true;
            this.page.page_num = page;
            this.getListData();
        }
    }
};
</script>

<style lang="scss" scoped>
.title-block {
    height: 52px;
    line-height: 52px;
    display: flex;
    justify-content: space-between;
    .title {
        font-weight: 500;
        font-size: 14px;
        color: #171e31;
        cursor: pointer;
        .title-content {
            margin-right: 4px;
        }
        .down-box {
            display: inline-block;
            position: relative;
            top: -1px;
            width: 12.21px;
            height: 6.86px;
            background-size: 100% 100%;
            background-image: url(@/assets/image/common/to-down.svg);
            transition-duration: 0.3s;
            &.close {
                transform: rotate(180deg);
                transition-duration: 0.3s;
            }
        }
    }
    .add-sub-task {
        cursor: pointer;
        font-weight: 400;
        font-size: 14px;
        color: #1890ff;
    }
}
.table-content {
    overflow: hidden;
    &.expand {
        height: fit-content;
        transition-duration: 0.3s;
    }
    &.close {
        height: 0;
        transition-duration: 0.3s;
    }
}
.plus-box {
    display: inline-block;
    vertical-align: middle;
    position: relative;
    top: -2px;
    width: 20px;
    height: 20px;
    background-size: 100% 100%;
    background-image: url(@/assets/image/common/icon_plus.png);
}
</style>
