<template>
    <div>
        <el-drawer
            :visible.sync="drawer"
            direction="rtl"
            size="60%"
            custom-class="node-drawer"
            :append-to-body="true"
        >
            <div slot="title" class="title-name">
                <!-- <tip-one :text="curBtnName"> -->
                <div class="name">
                    步骤流转:&nbsp;&nbsp;&nbsp;
                    <tip-one :text="curNodeRow.name">
                        <span
                            class="label-name"
                            :style="{
                                // color: `${curNodeRow.color || '#000'}`,
                                color: '#171e31',
                                backgroundColor: rgbToRgba(
                                    `${curNodeRow.color || '#fff'}`,
                                    0.3
                                )
                            }"
                            >{{ curNodeRow.name || "--" }}
                        </span>
                    </tip-one>
                    <img
                        :src="require('@/assets/image/common/right_arrow.png')"
                        alt=""
                        width="16px"
                        height="16px"
                        class="arrow"
                    />
                    <tip-one :text="curItem.label">
                        <span
                            class="label-name"
                            :style="{
                                // color: `${curItem.color || '#000'}`,
                                color: '#171e31',
                                backgroundColor: rgbToRgba(
                                    `${curItem.color || '#fff'}`,
                                    0.3
                                )
                            }"
                            >{{ curItem.label || "--" }}
                        </span>
                    </tip-one>
                </div>
                <!-- </tip-one> -->
                <el-tabs class="basic-ui" v-model="tabName">
                    <el-tab-pane
                        v-for="(item, index) in btnsList"
                        :key="index"
                        :label="item.label"
                        :name="item.key"
                        @click="handerClickTab(item)"
                    ></el-tab-pane>
                </el-tabs>
            </div>
            <!-- 内容 -->
            <div v-if="tabName === '转换界面'">
                <!-- <el-alert
                    v-if="isAlertShow"
                    :title="remindMessage"
                    type="warning"
                    show-icon
                    @close="isAlertShow = false"
                > 
                </el-alert>-->
                <div class="field-suggest">
                    <div class="field">
                        选择字段：
                        <el-select
                            class="basic-ui"
                            size="small"
                            v-model="field"
                            placeholder="请选择字段"
                            @change="fieldChange"
                        >
                            <el-option
                                v-for="item in fieldOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                            >
                            </el-option>
                        </el-select>
                    </div>
                    <div class="suggest">
                        <!-- <span class="text">转换意见</span>
                        <tip-question
                            tips="是否为当前转换开启转换意见，默认不开启，开启后则在转换时增加转换意见文本框。"
                        ></tip-question>
                        <el-switch
                            v-model="suggest"
                            :active-color="PRIMARY_COLOR"
                            inactive-color="#E6E9F0"
                        >
                        </el-switch> -->
                    </div>
                </div>
                <n-table
                    ref="screenTable"
                    class="table-content"
                    :data="converData"
                    :col="converCol"
                >
                    <template slot="field_name" slot-scope="{ column }">
                        <el-table-column
                            :show-overflow-tooltip="false"
                            :min-width="180"
                            :label="column.name"
                        >
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
                                    <span class="name">{{
                                        scope.row.field_name || "--"
                                    }}</span>
                                </span>
                            </template>
                        </el-table-column>
                    </template>
                    <template slot="require" slot-scope="{ column }">
                        <el-table-column
                            show-overflow-tooltip
                            :min-width="150"
                            :label="column.name"
                        >
                            <template slot-scope="scope">
                                <el-switch
                                    v-model="scope.row.required_flag"
                                    :active-color="PRIMARY_COLOR"
                                    inactive-color="#E6E9F0"
                                    @change="
                                        (val) => switchChange(val, scope.row)
                                    "
                                >
                                </el-switch>
                            </template>
                        </el-table-column>
                    </template>
                    <template slot="operation" slot-scope="{ column }">
                        <el-table-column
                            show-overflow-tooltip
                            :width="80"
                            :label="column.name"
                            :align="column.align"
                        >
                            <template slot-scope="scope">
                                <!-- <el-button
                                    @click="nodeChangeConfig(scope.row)"
                                    type="text"
                                    size="small"
                                >
                                    设置
                                </el-button> -->
                                <el-button
                                    @click="deleteField(scope.row)"
                                    type="text"
                                    size="small"
                                >
                                    删除
                                </el-button>
                            </template>
                        </el-table-column>
                    </template>
                </n-table>
            </div>
            <div v-if="tabName === '触发器'">
                <!-- 添加触发器 -->
                <div class="trigger-condition">
                    <el-button
                        size="small"
                        class="basic-ui"
                        type="primary"
                        @click="addTrigger"
                        >增加触发器</el-button
                    >
                </div>
                <n-table
                    class="table-content"
                    :data="converData"
                    :col="converCol"
                >
                    <template slot="operation" slot-scope="{ column }">
                        <el-table-column
                            show-overflow-tooltip
                            :width="200"
                            :label="column.name"
                        >
                            <template slot-scope="scope">
                                <div>
                                    <el-button
                                        @click="deleteField(column, scope.row)"
                                        type="text"
                                        size="small"
                                        class="basic-ui"
                                    >
                                        编辑
                                    </el-button>
                                </div>
                            </template>
                        </el-table-column>
                    </template>
                </n-table>
            </div>
            <div v-if="tabName === '条件'">
                <div class="trigger-condition">
                    <el-button
                        type="primary"
                        size="small"
                        class="basic-ui"
                        @click="addFilter"
                        >增加条件</el-button
                    >
                </div>
                <n-table
                    class="table-content"
                    :data="converData"
                    :col="converCol"
                >
                    <template slot="operation" slot-scope="{ column }">
                        <el-table-column
                            show-overflow-tooltip
                            :width="200"
                            :label="column.name"
                        >
                            <template slot-scope="scope">
                                <div>
                                    <el-button
                                        @click="deleteField(column, scope.row)"
                                        type="text"
                                        size="small"
                                    >
                                        删除
                                    </el-button>
                                </div>
                            </template>
                        </el-table-column>
                    </template>
                </n-table>
            </div>
            <!-- 限制器 -->
            <div class="limiter-block" v-if="tabName === '限制器'">
                <!-- <limiter :btnInfo="btnInfo" :rowInfo="rowInfo"></limiter> -->
                <limiter2
                    :btnInfo="btnInfo"
                    :rowInfo="rowInfo"
                    :curId="curId"
                ></limiter2>
            </div>
        </el-drawer>
        <add-trigger-dialog
            ref="AddTriggerDialog"
            :curId="curId"
        ></add-trigger-dialog>
        <add-filter-dialog
            ref="AddFilterDialog"
            :curId="curId"
        ></add-filter-dialog>
        <node-change-dialog
            ref="NodeChangeDialog"
            :curId="curId"
        ></node-change-dialog>
    </div>
</template>

<script>
import _ from "lodash";
import Sortable from "sortablejs";
import { rgbToRgba } from "@/assets/tool/func";
import TEXT from "@/assets/tool/text";
import api from "@/common/api/module/progress_setting";
import AddTriggerDialog from "./add_trigger_dialog";
import AddFilterDialog from "./add_filter_dialog";
import NTable from "@/pages/common/tables/common_table/table";
import { PRIMARY_COLOR_LIST } from "@/assets/tool/color_list";
import TipQuestion from "@/pages/common/tip_question.vue";
import TipOne from "@/pages/common/tooltip_one_line.vue";
import NodeChangeDialog from "./node_change_dialog.vue";
import Limiter from "./limiter.vue";
import Limiter2 from "./limiter_v2.vue";
export default {
    components: {
        AddTriggerDialog,
        AddFilterDialog,
        NTable,
        TipQuestion,
        TipOne,
        NodeChangeDialog,
        Limiter,
        Limiter2
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        }
    },
    data() {
        return {
            isAlertShow: true,
            remindMessage: TEXT.PROGRESS_SETTING_NODE_DREAWER,
            convertScreenDialogVisible: false,
            curItem: {},
            btnInfo: {}, //当前按钮信息
            rowInfo: {}, //当前行信息
            curNodeRow: {},
            convertRow: {},
            PRIMARY_COLOR: PRIMARY_COLOR_LIST.PRIMARY_COLOR,
            curBtnName: "",
            drawer: false,
            tabName: "",
            curBtn: "create",
            btnsList: [
                {
                    label: "转换界面",
                    key: "转换界面",
                    active: true
                },
                // {
                //     label: '触发器',
                //     key: '触发器',
                //     active: false,
                // },
                // {
                //     label: '条件',
                //     key: '条件',
                //     active: false,
                // },
                {
                    label: "限制器",
                    key: "限制器",
                    active: false
                }
            ],
            fieldOptions: [],
            field: "",
            suggest: false,
            converCol: [
                {
                    prop: "field_name",
                    name: "字段名称",
                    type: "slot"
                },
                {
                    prop: "require",
                    name: "是否必填",
                    type: "slot"
                },
                {
                    prop: "operation",
                    align: "left",
                    name: "操作",
                    type: "slot"
                }
            ],
            converData: [],
            curId: ""
        };
    },
    watch: {
        // drawer: {
        //     handler(Boolean) {
        //         if (!Boolean) {
        //         }
        //     }
        // }
    },

    methods: {
        handerClickTab(item) {},
        rgbToRgba(color, opacity) {
            return rgbToRgba(color, opacity);
        },
        rowDrop() {
            let _that = this;
            let ref = _that.$refs.screenTable;
            let tbody = ref.$el.querySelector(".el-table__body-wrapper tbody");
            Sortable.create(tbody, {
                group: {
                    name: "words",
                    pull: true,
                    put: true
                },
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
                    let curNodeRow = _that.converData.splice(oldIndex, 1)[0];
                    _that.converData.splice(newIndex, 0, curNodeRow);
                    // 调接口发送更改后的数据给后端并获取新数据
                    if (_that.converData.length < 2) return;
                    let tmpKeyList = _that.converData.map(
                        (item) => item.field_key
                    );
                    let params = {
                        ws_id: _that.curSpace.id,
                        tmpl_id: _that.$route.params.id,
                        start_status_id: _that.curNodeRow.id,
                        end_status_id: _that.curItem.id,
                        step_id: _that.curId
                    };
                    if (tmpKeyList.length) {
                        params.coordinate = tmpKeyList.join(",");
                    }
                    api.moveTSField(params).then((res) => {
                        if (res && res.resultCode === 200) {
                            //
                        }
                    });
                }
            });
        },
        deleteField(row) {
            this.convertRow = _.cloneDeep(row);
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                step_id: this.curId,
                id: row.id
            };
            api.delTSField(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.fetchFieldOptions();
                    this.fetchTableData();
                }
            });
        },
        nodeChangeConfig(row) {
            this.$refs.NodeChangeDialog.openDialog(row);
        },

        fieldChange(val) {
            this.field = "";
            this.createConverPage(val);
        },

        createConverPage(val) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                step_id: this.curId,
                field_key: val
            };
            api.createTSField(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.fetchFieldOptions();
                    this.fetchTableData();
                }
            });
        },
        switchChange(val, subRow) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                id: subRow.id,
                required: val ? "yes" : "no",
                step_id: this.curId
            };
            api.setTSRequired(params).then((res) => {
                if (res && res.resultCode === 200) {
                    // this.$message.success('设置成功')
                }
            });
        },
        clickBtn(curItem) {
            this.curBtn = curItem.key;
            this.curObj = curItem;
            this.btnsList.forEach((item) => {
                this.$set(item, "active", item.label === curItem.label);
            });
        },
        addTrigger() {
            this.$refs.AddTriggerDialog.openDialog();
        },
        addFilter() {
            this.$refs.AddFilterDialog.openDialog();
        },
        openDrawer(item, row) {
            this.isAlertShow = true;
            this.curItem = _.cloneDeep(item);
            this.curNodeRow = _.cloneDeep(row);
            if (row.steps && row.steps.length) {
                row.steps.forEach((step) => {
                    if (
                        step.start_status_id === row.id &&
                        step.end_status_id === item.id
                    ) {
                        this.curId = step.id;
                    }
                });
            }
            this.btnInfo = item;
            this.rowInfo = row;
            this.curBtnName = item ? item.name || "" : "";
            this.fetchFieldOptions();
            this.fetchTableData();
            this.drawer = true;
            this.tabName = "转换界面";
            this.$nextTick(() => {
                if (this.tabName === "转换界面") {
                    this.rowDrop();
                }
            });
        },
        fetchFieldOptions() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                step_id: this.curId
            };
            api.getTSOptions(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    res.data.length
                ) {
                    this.fieldOptions = this.formatField(res.data);
                } else {
                    this.fieldOptions = [];
                }
            });
        },
        formatField(arr) {
            return arr.map((item) => {
                return {
                    label: item.name,
                    value: item.field_key,
                    key: item.id
                };
            });
        },
        fetchTableData() {
            // 表格数据
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                step_id: this.curId
            };
            api.getTSList(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    res.data.length
                ) {
                    res.data.forEach((item) => {
                        this.$set(
                            item,
                            "required_flag",
                            item.required === "no" ? false : true
                        );
                    });
                    this.converData = res.data;
                } else {
                    this.converData = [];
                }
            });
        },
        handleClose(done) {
            this.$confirm("确认关闭？")
                .then((_) => {
                    done();
                })
                .catch((_) => {});
        }
    }
};
</script>

<style lang="scss" scoped>
.limiter-block {
    // 减去title和tabs高度
    height: calc(100% - 50px);
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
}
.el-button-group {
    margin-bottom: 20px;
    .el-button {
        border: 0;
        color: var(--GRAY_FONT_COLOR);
        border-radius: 4px !important;
    }
    .el-button.active-btn {
        color: var(--PRIMARY_COLOR);
        background-color: #eff8ff;
    }
}
.el-alert {
    margin-bottom: 16px;
}
.table-content {
    margin-top: 16px;
}
.field-suggest {
    display: flex;
    justify-content: space-between;
    font-size: 14px;
}
.trigger-condition {
    display: flex;
    justify-content: right;
}
.suggest {
    display: flex;
    align-items: center;
    .text {
        position: relative;
        top: -1px;
        color: var(--FONT_COLOR);
    }
    ::v-deep .el-switch {
        margin-left: 16px;
    }
}
</style>
<style lang="scss">
.el-drawer.node-drawer {
    overflow: visible;

    .el-drawer__header {
        margin-bottom: 2px;
        position: relative;
        .el-drawer__close-btn {
            position: absolute;
            top: 20px;
            left: -40px;
            .el-icon-close:before {
                color: #fff;
            }
        }
        .title-name {
            font-weight: 500;
            font-size: 20px;
            color: #171e31;
            width: calc(100% - 20px);
            .name {
                display: flex;
                align-items: center;
                margin-bottom: 12px;
            }
            img.arrow {
                margin: 0 12px;
            }
            .label-name {
                box-sizing: border-box;
                display: inline-block;
                height: 24px;
                line-height: 23px;
                border-radius: 4px;
                padding: 0 10px;
                font-size: 14px;
                font-weight: 400;
                color: #171e31;
                max-width: calc((100% - 60px - 40px) / 2);
                text-overflow: ellipsis;
                overflow: hidden;
                white-space: nowrap;
            }
        }
    }
    .el-drawer__body {
        padding: 0 24px 24px !important;
    }
}
</style>
