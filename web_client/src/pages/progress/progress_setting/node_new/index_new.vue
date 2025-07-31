<template>
    <div class="progress-setting-filed">
        <!-- <el-alert :title="remindMessage" type="warning" show-icon> </el-alert> -->
        <div class="btns">
            <el-button
                size="small"
                class="basic-ui"
                type="primary"
                @click="addNode"
            >
                <b class="add"></b>工作状态
            </el-button>
        </div>
        <div class="table-box" :class="curCol && curCol.length ? '' : 'empty'">
            <el-table
                class="normal-table"
                :data="tableData"
                row-key="id"
                id="nodeTable"
                :key="rendertable"
                ref="nodeTable"
                border
                :header-cell-style="handleHeaderCell"
                @cell-mouse-enter="mouseEnter"
                @cell-mouse-leave="mouseLeave"
            >
                <!-- 第一列(开始和目标状态) -->
                <el-table-column
                    :show-overflow-tooltip="false"
                    :min-width="240"
                    fixed="left"
                    :resizable="false"
                    v-if="curCol && curCol.length"
                >
                    <template slot="header" slot-scope="scope">
                        <div
                            style="text-align: right; margin: -4px 10px -4px 0"
                        >
                            目标状态
                        </div>
                        <div style="margin: -4px 0 -4px 10px">开始状态</div>
                    </template>
                    <template slot-scope="scope">
                        <div
                            style="
                                box-sizing: border-box;
                                display: flex;
                                justify-content: space-between;
                                align-items: center;
                                height: 50px;
                                background-color: #fafafa;
                                padding: 0 6px;
                            "
                        >
                            <span
                                style="display: flex; align-items: center"
                                :style="{
                                    width: nameWidth(scope.row)
                                }"
                            >
                                <img
                                    :src="
                                        require('@/assets/image/common/icon_dragtable_move.png')
                                    "
                                    alt=""
                                    width="20px"
                                    height="20px"
                                />
                                <tip-one
                                    :text="scope.row.name"
                                    position="top"
                                    :width="nameWidth(scope.row)"
                                >
                                    <span
                                        class="row-col-name row-name"
                                        :style="{
                                            color: '#171e31',
                                            backgroundColor: rgbToRgba(
                                                `${scope.row.color || '#fff'}`,
                                                0.3
                                            )
                                        }"
                                    >
                                        {{ scope.row.name || "--" }}</span
                                    >
                                </tip-one>
                            </span>

                            <span>
                                <span
                                    style="
                                        display: inline-block;
                                        height: 22px;
                                        line-height: 22px;
                                        padding: 0 8px;
                                        color: #a6abbc;
                                        font-size: 12px;
                                        border: 1px solid #ced2de;
                                        background-color: #fff;
                                        border-radius: 14px;
                                        margin-right: 2px;
                                    "
                                    v-if="
                                        scope.row.isShowStatus &&
                                        scope.row.mode === 'first'
                                    "
                                    >初始状态</span
                                >
                                <el-popover
                                    :ref="`popper_node_${scope.row.id}`"
                                    placement="bottom-start"
                                    trigger="click"
                                    popper-class="node-setting-nam-col"
                                    :visible-arrow="false"
                                    @show="
                                        (val) => poppoverShow(scope.row, val)
                                    "
                                    @hide="
                                        (val) => poppoverHide(scope.row, val)
                                    "
                                >
                                    <div
                                        class="pop-item"
                                        @click="setInitStatus(scope.row)"
                                    >
                                        <img
                                            :src="
                                                require(`@/assets/image/common/setting_status2.png`)
                                            "
                                            alt=""
                                            width="18px"
                                            height="18px"
                                        />
                                        设置初始化状态
                                    </div>
                                    <div
                                        class="pop-item"
                                        @click="editNode(scope.row)"
                                    >
                                        <img
                                            :src="
                                                require(`@/assets/image/common/edit.svg`)
                                            "
                                            alt=""
                                            width="18px"
                                            height="18px"
                                        />
                                        编辑工作状态
                                    </div>
                                    <div
                                        class="pop-item"
                                        @click="deleteNode(scope.row)"
                                    >
                                        <img
                                            :src="
                                                require(`@/assets/image/common/delete.svg`)
                                            "
                                            alt=""
                                            width="18px"
                                            height="18px"
                                        />
                                        删除状态
                                    </div>

                                    <img
                                        v-show="scope.row.isSettingShow"
                                        slot="reference"
                                        :src="
                                            require(`@/assets/image/common/setting_status.png`)
                                        "
                                        width="18px"
                                        height="18px"
                                        style="
                                            position: relative;
                                            top: 4px;
                                            cursor: pointer;
                                        "
                                    />
                                </el-popover>
                            </span>
                        </div>
                    </template>
                </el-table-column>
                <el-table-column
                    v-for="col in curCol"
                    :show-overflow-tooltip="false"
                    :min-width="160"
                    :key="col.id"
                    :align="col.align"
                >
                    <template slot="header" slot-scope="scope">
                        <tip-one :text="col.label" position="top">
                            <div
                                class="row-col-name"
                                :style="{
                                    color: '#171e31',
                                    backgroundColor: rgbToRgba(
                                        `${col.color || '#fff'}`,
                                        0.3
                                    )
                                }"
                            >
                                {{ col.label || "--" }}
                            </div>
                        </tip-one>
                    </template>
                    <template slot-scope="scope">
                        <div
                            v-if="scope.row.id === col.id"
                            style="
                                height: 50px;
                                line-height: 56px;
                                cursor: not-allowed;
                                text-align: center;
                                background-color: #fafafa;
                            "
                        >
                            <img
                                :src="
                                    require(`@/assets/image/common/not_allow_select.png`)
                                "
                                width="18px"
                                height="18px"
                            />
                        </div>
                        <el-checkbox
                            v-else
                            v-model="
                                scope.row[`row_${scope.row.id}_${col.id}`]
                                    .checked
                            "
                            @change="
                                (val) => checkboxChange(val, scope.row, col)
                            "
                            @click.native.prevent="
                                (e, val) =>
                                    handleCheck(
                                        e,
                                        scope.row[
                                            `row_${scope.row.id}_${col.id}`
                                        ].checked,
                                        scope.row,
                                        col
                                    )
                            "
                        ></el-checkbox>
                        <img
                            v-show="
                                scope.row[`row_${scope.row.id}_${col.id}`]
                                    .checked
                            "
                            :src="require(`@/assets/image/common/edit.svg`)"
                            width="18px"
                            height="18px"
                            style="
                                margin-left: 4px;
                                position: relative;
                                top: 4px;
                                cursor: pointer;
                            "
                            @click="openConvertDetail(col, scope.row)"
                        />
                    </template>
                </el-table-column>
                <template slot="empty">
                    <div style="height: 300px">
                        <no-data
                            :isTextShow="false"
                            :loading="loading"
                            :imgShow="imgShow"
                        ></no-data>
                    </div>
                </template>
            </el-table>
            <!-- 对角线 -->
            <div
                class="line"
                :style="{
                    width: lineWhide + 'px',
                    transform: 'rotate(' + lineDeg + 'deg)'
                }"
            ></div>
        </div>
        <node-drawer ref="NodeDrawer"></node-drawer>
        <add-node-dialog
            ref="AddNodeDialog"
            @on-confirm="refreshNodeList('update-node')"
        ></add-node-dialog>
        <delete-node-dialog
            ref="DeleteNodeDialog"
            @on-confirm="refreshNodeList('delete-node')"
        ></delete-node-dialog>
        <add-convert-btn-dialog
            ref="AddConvertBtnDialog"
            @on-confirm="refreshNodeList('update-btn')"
        ></add-convert-btn-dialog>
        <delete-convert-btn-dialog
            ref="DeleteConvertBtnDialog"
            @on-confirm="refreshNodeList('delete-btn')"
        ></delete-convert-btn-dialog>
        <!-- 移除勾选，即删除按钮到按钮的steps -->
        <delete-steps-dialog
            ref="DeleteStepsDialog"
            @on-confirm="refreshNodeList('delete-steps')"
        ></delete-steps-dialog>
    </div>
</template>

<script>
import _ from "lodash";
import Sortable from "sortablejs";
import TEXT from "@/assets/tool/text";
import api from "@/common/api/module/progress_setting";
// import NTable from "@/pages/common/tables/common_table/table";
import NodeDrawer from "./components/node_drawer";
import AddNodeDialog from "./components/node_add_dialog";
import DeleteNodeDialog from "./components/node_delete_dialog";
import AddConvertBtnDialog from "./components/convert_add_btn_dialog";
import DeleteConvertBtnDialog from "./components/convert_delete_btn_dialog";
import DeleteStepsDialog from "./components/steps_delete_dialog";
import TipOne from "@/pages/common/tooltip_one_line.vue";
import NoData from "@/pages/common/no_data.vue";
import { rgbToRgba } from "@/assets/tool/func";

export default {
    components: {
        NodeDrawer,
        AddNodeDialog,
        DeleteNodeDialog,
        AddConvertBtnDialog,
        DeleteConvertBtnDialog,
        DeleteStepsDialog,
        TipOne,
        NoData
    },
    computed: {
        nameWidth() {
            return (row) => {
                return row.isSettingShow
                    ? "calc(100% - 28px)"
                    : row.isShowStatus && row.mode === "first"
                    ? "calc(100% - 70px)"
                    : "100%";
            };
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        getStatus() {
            return (row, col) => {
                if (row.steps && row.steps.length) {
                    let arr = row.steps.filter((item) => {
                        return (
                            item.start_status_id === row.id &&
                            item.end_status_id === col.id
                        );
                    });
                    this.$set(
                        row,
                        `checked_${row.id}_${col.id}`,
                        arr.length ? true : false
                    );
                } else {
                    this.$set(row, `checked_${row.id}_${col.id}`, false);
                }
                return `checked_${row.id}_${col.id}`;
            };
        }
    },
    data() {
        return {
            remindMessage: TEXT.PROGRESS_SETTING_NODE,
            modeRadioList: [],
            curIndex: "",
            tableCol: [],
            curCol: [],
            tableData: [],
            isSettingShow: true,
            isShowStatus: false,
            rendertable: "default",
            loading: true,
            imgShow: false,
            //第一个单元格的类名
            firstCellClassName: "",
            //第一个单元格的宽度
            firstCellWidth: 0,
            //第一个单元格的高度
            firstCellHeight: 0,
            //对角线的宽度
            lineWhide: 0,
            //对角线的角度
            lineDeg: 0,
            //窗口宽度
            screenWidth: document.body.clientWidth,
            timer: null
        };
    },
    watch: {
        $route() {
            this.fetchNodeList();
        },
        screenWidth(val) {
            let _this = this;
            if (!_this.timer) {
                // 一旦监听到的screenWidth值改变，就将其重新赋给data里的screenWidth
                _this.screenWidth = val;
                _this.timer = true;
                setTimeout(function () {
                    _this.getsElement();
                    _this.timer = false;
                }, 400);
            }
        }
    },
    mounted() {
        this.fetchNodeConfig();
        let _this = this;
        //mounted里挂载window.onresize事件
        window.onresize = () => {
            return (() => {
                window.screenWidth = document.body.clientWidth;
                _this.screenWidth = window.screenWidth;
            })();
        };
    },
    methods: {
        handleHeaderCell({ row, column, rowIndex, columIndex }) {
            //如果使用的多级表头  就需要判断是第几个表头
            if (rowIndex == 0) {
                //获取第一个单元格的class类名
                this.firstCellClassName = row[0].id;
                //获取第一个单元格的宽度，这个方法直接返回了
                this.firstCellWidth = row[0].realWidth;
                //这个时候他是找不到元素的  所以在调用一次
                this.$nextTick(() => {
                    this.getsElement();
                });
            } else if (rowIndex == 1) {
                return { display: "none" };
            }
        },
        getsElement() {
            //用原生获取第一个单元格的高度
            var a = document.getElementsByClassName(this.firstCellClassName)[0];
            if (!a) return;
            this.firstCellHeight = a.offsetHeight;
            //获取对角线的长度 a²+b²=c²
            this.lineWhide = Math.sqrt(
                this.firstCellWidth * this.firstCellWidth +
                    this.firstCellHeight * this.firstCellHeight
            );

            //获取对角线旋转的角度
            this.lineDeg =
                (Math.atan(this.firstCellHeight / this.firstCellWidth) * 180) /
                Math.PI;
        },
        rgbToRgba(color, opacity) {
            return rgbToRgba(color, opacity);
        },
        // poppover显示
        poppoverShow(row, val) {
            this.$set(row, "isPopperShow", true);
        },
        // poppover隐藏
        poppoverHide(row, val) {
            this.$set(row, "isPopperShow", false);
            this.$set(row, "isSettingShow", false);
        },
        // 设置初始化状态
        setInitStatus(row) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                id: row.id
            };
            api.setFirst(params).then((res) => {
                if (res && res.resultCode === 200) {
                    document.body.click();
                    this.fetchNodeList();
                } else {
                    this.$set(row, "mode", "");
                }
            });
        },
        mouseLeave(row, col) {
            if (row.isPopperShow) return;
            this.$set(row, "isSettingShow", false);
            this.$set(row, "isShowStatus", true);
        },
        mouseEnter(row, col) {
            this.$set(row, "isShowStatus", false);
            this.$set(row, "isSettingShow", true);
        },
        downLoadCurrent() {},
        rowDrop() {
            let tbody = document
                .getElementById("nodeTable")
                .querySelector(
                    ".el-table__fixed .el-table__fixed-body-wrapper tbody"
                );
            let _this = this;
            Sortable.create(tbody, {
                // or { name: "...", pull: [true, false, 'clone', array], put: [true, false, array] }
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
                    let currRow = _this.tableData.splice(oldIndex, 1)[0];
                    _this.tableData.splice(newIndex, 0, currRow);
                    // 调接口发送更改后的数据给后端并获取新数据
                    let tmpKeyList = _this.tableData.map((item) => item.id);
                    let params = {
                        ws_id: _this.curSpace.id,
                        tmpl_id: _this.$route.params.id
                    };
                    if (tmpKeyList.length) {
                        params.coordinate = tmpKeyList.join(",");
                    }
                    api.moveOverviewTable(params).then((res) => {
                        if (res && res.resultCode === 200) {
                            _this.fetchNodeList();
                        }
                    });
                }
            });
        },
        rowDrop2() {
            let tbody = document
                .getElementById("nodeTable")
                .querySelector(".el-table__body-wrapper tbody");
            let _this = this;
            Sortable.create(tbody, {
                // or { name: "...", pull: [true, false, 'clone', array], put: [true, false, array] }
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
                    let currRow = _this.tableData.splice(oldIndex, 1)[0];
                    _this.tableData.splice(newIndex, 0, currRow);
                    // 调接口发送更改后的数据给后端并获取新数据
                    let tmpKeyList = _this.tableData.map((item) => item.id);
                    let params = {
                        ws_id: _this.curSpace.id,
                        tmpl_id: _this.$route.params.id
                    };
                    if (tmpKeyList.length) {
                        params.coordinate = tmpKeyList.join(",");
                    }
                    api.moveOverviewTable(params).then((res) => {
                        if (res && res.resultCode === 200) {
                            _this.fetchNodeList();
                        }
                    });
                }
            });
        },
        //  @click.native.prevent 为了阻止点击取消的时候还没判断就已经取消的问题，此时change方法不生效
        handleCheck(e, val, row, col) {
            if (val) {
                this.$refs.DeleteStepsDialog.openDialog(val, row, col);
            } else {
                //增加的时候
                let params = {
                    ws_id: this.curSpace.id,
                    tmpl_id: this.$route.params.id,
                    start_status_id: row.id,
                    end_status_id: col.id
                };

                api.addSteps(params).then((res) => {
                    if (res && res.resultCode === 200) {
                        this.fetchNodeList();
                    } else {
                        this.$set(
                            row[`row_${row.id}_${col.id}`],
                            "checked",
                            false
                        );
                    }
                });
            }
        },
        checkboxChange(val, row, col) {
            if (!val) {
                this.$refs.DeleteStepsDialog.openDialog(val, row, col);
            } else {
                //增加的时候
                let params = {
                    ws_id: this.curSpace.id,
                    tmpl_id: this.$route.params.id,
                    start_status_id: row.id,
                    end_status_id: col.id
                };

                api.addSteps(params).then((res) => {
                    if (res && res.resultCode === 200) {
                        this.fetchNodeList();
                    } else {
                        this.$set(
                            row[`row_${row.id}_${col.id}`],
                            "checked",
                            false
                        );
                    }
                });
            }
        },
        fetchNodeConfig() {
            api.getNodeConfig().then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.modeRadioList = _.cloneDeep(
                        res.data.mode_config || []
                    );
                    this.fetchNodeList();
                } else {
                    this.modeRadioList = [];
                }
            });
        },
        fetchNodeList() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id
            };
            api.getOverviewTable(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    res.data.length
                ) {
                    this.formatCol(res.data);
                    this.formatData(res.data);
                    this.$refs.nodeTable && this.$refs.nodeTable.doLayout();
                    this.rendertable = Math.random();
                    this.$nextTick(() => {
                        this.rowDrop();
                        // this.rowDrop2();
                    });
                } else {
                    this.tableData = [];
                    this.curCol = [];
                }
                this.loading = false;
                setTimeout(() => {
                    this.imgShow = true;
                }, 0);
            });
        },
        formatCol(data) {
            if (data && data.length) {
                let curTmp = data.map((item) => {
                    return {
                        label: item.name,
                        prop: `col_${item.id}`,
                        color: item.color,
                        type: "slot",
                        id: item.id,
                        steps: item.steps,
                        align: "center"
                    };
                });
                this.curCol = _.cloneDeep(curTmp);
            } else {
                this.curCol = [];
            }
        },
        formatData(data) {
            let tmpTableData = _.cloneDeep(data);
            tmpTableData.forEach((row) => {
                if (row.mode === "first") {
                    this.$set(row, "isShowStatus", true);
                } else {
                    this.$set(row, "isShowStatus", false);
                }
                this.curCol.forEach((col) => {
                    this.$set(row, `row_${row.id}_${col.id}`, {});
                    this.$set(row[`row_${row.id}_${col.id}`], "checked", false);
                    this.$set(row[`row_${row.id}_${col.id}`], "row_id", row.id);
                    this.$set(row[`row_${row.id}_${col.id}`], "col_id", col.id);
                    if (row.steps && row.steps.length) {
                        row.steps.forEach((step) => {
                            if (
                                step.start_status_id === row.id &&
                                step.end_status_id === col.id
                            ) {
                                this.$set(
                                    row[`row_${row.id}_${col.id}`],
                                    "checked",
                                    true
                                );
                            }
                        });
                    }
                });
            });
            this.tableData = tmpTableData;
        },
        refreshNodeList(type) {
            let mapping = {
                "update-node": "AddNodeDialog",
                "delete-node": "DeleteNodeDialog",
                "update-btn": "AddConvertBtnDialog",
                "delete-btn": "DeleteConvertBtnDialog",
                "delete-steps": "DeleteStepsDialog"
            };
            this.$refs[mapping[type]].cancel(); // 关闭对应的dialog
            this.fetchNodeList();
            document.body.click();
        },
        addNode() {
            // 新增节点
            this.$refs.AddNodeDialog.openDialog("add", this.modeRadioList);
        },
        editNode(row) {
            // 编辑节点
            this.$refs.AddNodeDialog.openDialog(
                "edit",
                this.modeRadioList,
                row
            );
        },
        deleteNode(row) {
            // 删除节点
            this.$refs.DeleteNodeDialog.openDialog(row);
        },
        addConvertBtn(item, row) {
            // 添加转换按钮
            this.$refs.AddConvertBtnDialog.openDialog(
                item,
                row,
                "add",
                this.tableData
            );
        },
        editConvertBtn(item, row) {
            // 编辑转换按钮
            this.$refs.AddConvertBtnDialog.openDialog(
                item,
                row,
                "edit",
                this.tableData
            );
        },
        deleteConvertBtn(item, row) {
            // 删除转换按钮
            this.$refs.DeleteConvertBtnDialog.openDialog(item, row);
        },
        openConvertDetail(item, row) {
            // 打开状态转换按钮设置详情状态
            this.$refs.NodeDrawer.openDrawer(item, row);
        }
    }
};
</script>

<style lang="scss" scoped>
.progress-setting-filed {
    .el-alert {
        margin-bottom: 16px;
    }
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
    .circle {
        display: inline-block;
        width: 8px;
        height: 8px;
        margin-right: 8px;
        border-radius: 50%;
    }
    .row-col-name {
        box-sizing: border-box;
        display: inline-block;
        height: 24px;
        line-height: 24px;
        border-radius: 4px;
        padding: 0 10px;
        color: #171e31;
        max-width: 100%;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        &.row-name {
            margin: 0 4px 0 4px;
            max-width: calc(100% - 26px);
        }
    }
    .table-box {
        position: relative;
        &.empty {
            ::v-deep {
                .el-table--border,
                .el-table--group {
                    border: none;
                }
            }
        }
        .line {
            position: absolute;
            /* position:fixed; */
            z-index: 9;
            top: 0;
            left: 0;
            height: 1px;
            width: 350px;
            background-color: #ebeef5;
            transform-origin: left;
            transform: rotate(24.2deg);
        }
    }
    ::v-deep
        .el-table--enable-row-hover
        .el-table__body
        tr:hover
        > td.el-table__cell {
        background-color: transparent;
    }
    ::v-deep .el-table__body .el-table__row.hover-row > td.el-table__cell {
        background-color: transparent;
    }
    ::v-deep .el-table--border {
        border-bottom: inherit;
    }
    ::v-deep .el-table--border .el-table__cell:first-child .cell {
        padding: 0 !important;
    }
    ::v-deep .el-table {
        &::before {
            height: 0;
        }
        th.el-table__cell > .cell {
            font-weight: 400;
            padding-left: 10px;
            padding-right: 10px;
        }
        th.el-table__cell {
            background-color: #fafafa;
        }
        .cell {
            padding: 0;
        }
        .el-table__cell {
            box-sizing: border-box;
            padding: 0;
            height: 50px;
            line-height: 49px;
        }
    }
    ::v-deep .el-button--mini {
        height: 24px;
        line-height: 24px;
        padding: 0 8px;
    }
    ::v-deep .el-button + .basic-ui.el-button,
    ::v-deep .el-checkbox.is-bordered + .basic-ui.el-checkbox.is-bordered {
        margin-left: 8px;
    }
    .btns-item {
        height: 32px;
        line-height: 32px;
        display: flex;
        align-items: center;
        .name {
            color: var(--PRIMARY_COLOR);
            text-decoration: underline;
            cursor: pointer;
        }
    }
    .status-flex {
        display: flex;
        align-items: center;
        .status-text {
            display: inline-block;
            width: calc(100% - 16px);
            overflow: hidden;
            white-space: nowrap;
            text-overflow: ellipsis;
        }
    }
}
</style>
<style lang="scss">
.node-setting-nam-col.el-popper[x-placement^="bottom"] {
    margin-top: 2px;
}
.node-setting-nam-col.el-popper[x-placement^="top"] {
    margin-bottom: 2px;
}
.node-setting-nam-col.el-popover {
    min-width: 20px;
    padding: 8px;
    border-radius: 4px;
    background-color: #fff;
    box-shadow: 2px 2px 8px 1px rgba(47, 56, 76, 0.3);
    .pop-item {
        display: flex;
        align-items: center;
        height: 32px;
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
