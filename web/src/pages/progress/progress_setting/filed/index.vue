<template>
    <div class="progress-setting-filed">
        <div class="btns">
            <el-button
                size="small"
                class="basic-ui"
                type="primary"
                @click="addFiled"
            >
                <b class="add"></b>新增字段
            </el-button>
        </div>
        <!-- <n-table :data="filedTableData" :col="tableCol"></n-table> -->
        <el-table
            row-key="id"
            class="setting-filed-table"
            :data="filedTableData"
            style="width: 100%"
            :row-class-name="rowClass"
        >
            <el-table-column prop="name" label="名称">
                <template slot-scope="scope">
                    <span class="filed-name-col">
                        <img
                            :src="
                                require('@/assets/image/common/icon_dragtable_move.png')
                            "
                            alt=""
                            width="20px"
                            height="20px"
                        />
                        <span
                            class="name"
                            :class="
                                fieldDisabled(scope.row) ? 'disabled-style' : ''
                            "
                            v-overflow
                            >{{ scope.row.name || "--" }}</span
                        >
                    </span>
                </template>
            </el-table-column>
            <el-table-column prop="mode" show-overflow-tooltip label="类型">
                <template slot-scope="scope">
                    <span v-if="scope.row.mode">
                        <img
                            style="
                                position: relative;
                                top: 5px;
                                margin-right: 4px;
                            "
                            :src="typeIconMapping[scope.row.mode]"
                            alt=""
                            width="20px"
                            height="20px"
                        />
                        <span>{{ getFormatLabel(scope.row.mode) }}</span>
                    </span>
                    <span v-else>--</span>
                </template>
            </el-table-column>
            <el-table-column prop="desc" show-overflow-tooltip label="描述">
            </el-table-column>
            <el-table-column
                prop="action"
                show-overflow-tooltip
                label="操作"
                width="160"
            >
                <template slot-scope="scope">
                    <el-button
                        @click="editFiled(scope.row)"
                        type="text"
                        size="small"
                        :disabled="EditBtnDisabled(scope.row)"
                    >
                        编辑
                    </el-button>
                    <el-button
                        @click="configOnlyReady(scope.row)"
                        type="text"
                        size="small"
                        class="read-rule"
                    >
                        只读规则
                        <b
                            v-show="isShowOnlyRuleIcon(scope.row)"
                            class="content-mark"
                        ></b>
                    </el-button>
                    <el-button
                        @click.native.prevent="deleteFiled(scope.row)"
                        type="text"
                        size="small"
                        :disabled="DelBtnDisabled(scope.row)"
                    >
                        删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
        <pagination
            v-show="total > 10"
            :total="total"
            @pageSizeChange="pageSizeChange"
            @curPageChange="curPageChange"
        ></pagination>
        <!-- 添加字段弹窗 -->
        <template>
            <filed-add-dialog
                ref="addFiledDialog"
                @on-confirm="onConfirmAddFiled"
            ></filed-add-dialog>
        </template>
        <!-- 编辑字段弹窗 -->
        <template>
            <filed-edit-dialog
                ref="editFiledDialog"
                @on-confirm="onConfirmEditFiled"
            ></filed-edit-dialog>
        </template>
        <!-- 删除字段弹窗 -->
        <template>
            <filed-delete-dialog
                ref="deleteFiledDialog"
                @on-confirm="onConfirmDeleteFiled"
            ></filed-delete-dialog>
        </template>
        <only-read-dialog
            ref="OnlyReadDialog"
            @refresh-table="onRefreshTable"
        ></only-read-dialog>
    </div>
</template>

<script>
import Pagination from "@/pages/common/pagination.vue";
import FiledAddDialog from "./common/filed_add_dialog";
import FiledDeleteDialog from "./common/filed_delete_dialog";
import FiledEditDialog from "./common/filed_edit_dialog";
import dataHandle from "./data_handle";
// import NTable from "@/pages/common/tables/com_table/table";
import api from "@/common/api/module/progress_setting";
import OnlyReadDialog from "./common/only_read_dialog.vue";
import _ from "lodash";
import Sortable from "sortablejs";
export default {
    components: {
        Pagination,
        FiledAddDialog,
        FiledDeleteDialog,
        FiledEditDialog,
        OnlyReadDialog
        // NTable
    },
    computed: {
        isShowOnlyRuleIcon() {
            return (row) => {
                if (!row.read_only_rule) return;
                if (row.read_only_rule) {
                    let obj = JSON.parse(row.read_only_rule);
                    if (
                        obj &&
                        Object.keys(obj).length &&
                        obj.status_list &&
                        obj.status_list.length
                    ) {
                        return true;
                    } else {
                        return false;
                    }
                }
            };
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        fieldDisabled() {
            return (row) => {
                // 标题不可删除，不可编辑开关
                if (row.field_key === "title" || row.field_key === "handler") {
                    return true;
                } else {
                    return false;
                }
            };
        },
        getFormatLabel() {
            return (val) => {
                let label = "--";
                if (!val) return label;
                if (this.typeOptions.length) {
                    this.typeOptions.forEach((item) => {
                        if (item.value === val) {
                            label = item.label;
                        }
                    });
                } else {
                    label = "--";
                }
                return label;
            };
        },
        // 字段是否可编辑
        EditBtnDisabled() {
            return (row) => {
                if (!row.level || (row.level && row.level === "level1")) {
                    return true;
                } else {
                    return false;
                }
            };
        },
        // 字段是否可删除（level1，level2不能删除，level3可以删除）
        DelBtnDisabled() {
            return (row) => {
                if (!row.level || (row.level && row.level !== "level3")) {
                    return true;
                } else {
                    return false;
                }
            };
        }
    },

    data() {
        return {
            typeOptions: [],
            numberUnitList: [],
            timeRadioOptions: [],
            personRadioOptions: [],
            curIndex: "",
            typeIconMapping: _.cloneDeep(dataHandle.typeIconMapping),
            filedTableData: [],
            tableCol: [],
            page: 1,
            size: 10,
            total: 0
        };
    },
    watch: {
        $route(to, from) {
            this.fetchFiledConfig();
            this.fetchFiledList();
        }
    },
    mounted() {
        this.fetchFiledConfig();
        this.fetchFiledList();
        this.rowDrop();
    },

    methods: {
        rowClass({ row }) {
            if (row.field_key !== "title" && row.field_key !== "handler") {
                return "is-sort";
            } else {
                return "not-sort";
            }
        },
        rowDrop() {
            let tbody = document.querySelector(".el-table__body-wrapper tbody");
            let _this = this;
            Sortable.create(tbody, {
                // or { name: "...", pull: [true, false, 'clone', array], put: [true, false, array] }
                group: {
                    name: "words",
                    pull: true,
                    put: true
                },
                draggable: ".is-sort",
                animation: 150, // ms, number 单位：ms，定义排序动画的时间
                onAdd: function (evt) {
                    // 拖拽时候添加有新的节点的时候发生该事件
                },
                onUpdate: function (evt) {
                    // 拖拽更新节点位置发生该事件
                },
                onRemove: function (evt) {
                    // 删除拖拽节点的时候促发该事件
                },
                onStart: function (evt) {
                    // 开始拖拽出发该函数
                },
                onSort: function (evt) {
                    // 发生排序发生该事件
                },
                onEnd({ newIndex, oldIndex }) {
                    // 结束拖拽
                    let currRow = _this.filedTableData.splice(oldIndex, 1)[0];
                    _this.filedTableData.splice(newIndex, 0, currRow);
                    // 调接口发送更改后的数据给后端并获取新数据
                    if (_this.filedTableData.length < 2) return;
                    let tmpKeyList = _this.filedTableData.map(
                        (item) => item.field_key
                    );
                    let params = {
                        ws_id: _this.curSpace.id,
                        tmpl_id: _this.$route.params.id,
                        module: "filed"
                    };
                    if (tmpKeyList.length) {
                        params.coordinate = tmpKeyList.join(",");
                    }
                    api.moveScreenField(params).then((res) => {
                        if (res && res.resultCode === 200) {
                            //
                        }
                    });
                }
            });
        },
        onRefreshTable() {
            this.page = 1;
            this.fetchFiledList();
        },
        pageSizeChange(size) {
            this.size = size;
            this.page = 1;
            this.fetchFiledList();
        },
        curPageChange(page) {
            this.page = page;
            this.fetchFiledList();
        },
        fetchFiledConfig() {
            this.$store.dispatch("fetchFiledConfig").then((res) => {
                if (res && Object.keys(res).length) {
                    const {
                        mode_config,
                        number_com_config,
                        time_com_config,
                        person_com_config
                    } = res;
                    this.setTypeConfig(mode_config);
                    this.setNumberConfig(number_com_config);
                    this.setTimeConfig(time_com_config);
                    this.setPersonConfig(person_com_config);
                } else {
                    this.typeOptions = [];
                    this.numberUnitList = [];
                    this.timeRadioOptions = [];
                    this.personRadioOptions = [];
                }
            });
        },
        setTypeConfig(arr) {
            if (arr && arr.length) {
                let defaultConfig = _.cloneDeep(dataHandle.typeIconMapping);
                arr.forEach((item) => {
                    for (let key in defaultConfig) {
                        if (item.value === key) {
                            item.icon = defaultConfig[key];
                        }
                    }
                });
                this.typeOptions = _.cloneDeep(arr);
            } else {
                this.typeOptions = [];
            }
        },
        setNumberConfig(arr) {
            if (arr && arr.length) {
                arr.unshift({
                    label: "自定义",
                    value: "self"
                });
                this.numberUnitList = arr;
            } else {
                this.numberUnitList = [];
            }
        },
        setTimeConfig(arr) {
            this.timeRadioOptions = arr && arr.length ? arr : [];
        },
        setPersonConfig(arr) {
            this.personRadioOptions = arr && arr.length ? arr : [];
        },
        fetchFiledList() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id
            };
            api.getFieldList(params).then((res) => {
                if (res && res.resultCode === 200 && res.data.length) {
                    this.filedTableData = res.data;
                    this.total = res.cnt || 0;
                } else {
                    this.filedTableData = [];
                    this.total = 0;
                }
            });
        },
        // 添加字段
        addFiled() {
            this.$refs.addFiledDialog.openDialog({
                typeOptions: this.typeOptions,
                numberUnitList: this.numberUnitList,
                timeRadioOptions: this.timeRadioOptions,
                personRadioOptions: this.personRadioOptions
            });
        },
        onConfirmAddFiled() {
            this.$refs.addFiledDialog.cancel();
            this.fetchFiledList();
        },
        // 编辑字段
        editFiled(row) {
            this.$refs.editFiledDialog.openDialog(row, {
                typeOptions: this.typeOptions,
                numberUnitList: this.numberUnitList,
                timeRadioOptions: this.timeRadioOptions,
                personRadioOptions: this.personRadioOptions
            });
        },
        onConfirmEditFiled() {
            this.$refs.editFiledDialog.cancel();
            this.fetchFiledList();
        },
        // 删除字段
        deleteFiled(row) {
            this.$refs.deleteFiledDialog.openDialog(row);
        },
        onConfirmDeleteFiled(row) {
            if (!row || !row.id) return;
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                id: row.id
            };
            api.deleteField(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.$message.success("删除字段成功");
                    this.$refs.deleteFiledDialog.cancel();
                    this.fetchFiledList();
                }
            });
        },
        // 只读规则
        configOnlyReady(row) {
            this.$refs.OnlyReadDialog.openDialog(row);
        }
    }
};
</script>

<style lang="scss" scoped>
.progress-setting-filed {
    .btns {
        text-align: right;
        margin-bottom: 16px;
        ::v-deep .el-button--small {
            font-size: 14px;
        }
        .add {
            display: inline-block;
            width: 16px;
            height: 16px;
            margin-right: 2px;
            vertical-align: text-bottom;
            background-image: url(@/assets/image/common/add_white.png);
            background-size: 100% 100%;
        }
    }
    ::v-deep .el-table {
        .el-table__row.not-sort {
            cursor: not-allowed;
            // opacity: 0.5;
        }
        .filed-name-col {
            display: flex;
            align-items: center;
            img {
                position: relative;
                left: -4px;
            }
            .name {
                display: inline-block;
                width: calc(100% - 20px);
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
            }
            .disabled-style {
                // color: #c0c4cc;
            }
        }
        th.el-table__cell > .cell {
            font-weight: 400;
        }
        th.el-table__cell {
            background-color: #f5f6f8;
        }
        .el-table__cell {
            box-sizing: border-box;
            padding: 0;
            height: 40px;
            line-height: 39px;
        }
    }
    .read-rule {
        position: relative;
        .content-mark {
            position: absolute;
            bottom: 5px;
            right: -3px;
            display: inline-block;
            height: 12px;
            width: 12px;
            background-image: url("~@/assets/image/progress_setting/field/rule_mark.png");
            background-size: 100% 100%;
        }
    }
}
</style>
