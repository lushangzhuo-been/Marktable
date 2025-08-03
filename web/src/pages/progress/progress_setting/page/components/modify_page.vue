<template>
    <div>
        <!-- <el-alert
            v-show="modules"
            :title="remindMessage"
            type="warning"
            show-icon
        >
        </el-alert> -->
        <div class="margin-top-16">
            <span class="label">选择字段:</span>
            <el-select
                class="basic-ui"
                v-model="filed"
                size="small"
                placeholder="请选择字段"
                @change="filedChange"
            >
                <el-option
                    v-for="item in filedOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                >
                </el-option>
            </el-select>
        </div>
        <el-table
            row-key="id"
            :row-class-name="rowClass"
            class="setting-filed-table"
            :data="screenTableData"
            style="width: 100%"
        >
            <!-- <el-table-column prop="icon" width="46px">
                <template slot-scope="scope">
                    <img
                        :src="
                            require('@/assets/image/common/icon_dragtable_move.png')
                        "
                        alt=""
                        width="24px"
                        height="24px"
                    />
                </template>
            </el-table-column> -->
            <el-table-column
                prop="field_name"
                show-overflow-tooltip
                label="字段名称"
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
                        <span
                            class="name"
                            :class="
                                fieldDisabled(scope.row) ? 'disabled-style' : ''
                            "
                            >{{ scope.row.field_name || "--" }}</span
                        >
                    </span>
                </template>
            </el-table-column>
            <!-- <el-table-column
                prop="can_modify"
                show-overflow-tooltip
                label="是否可编辑"
            >
                <template slot-scope="scope">
                    <el-switch
                        v-model="scope.row.can_modify"
                        :active-color="PRIMARY_COLOR"
                        inactive-color="#E6E9F0"
                        :disabled="fieldDisabled(scope.row)"
                        @change="
                            (val) => switchChange(val, scope.row, 'can_modify')
                        "
                    >
                    </el-switch>
                </template>
            </el-table-column> -->
            <el-table-column
                prop="required_flag"
                show-overflow-tooltip
                label="是否必填"
            >
                <template slot-scope="scope">
                    <el-switch
                        v-model="scope.row.required_flag"
                        :active-color="PRIMARY_COLOR"
                        inactive-color="#E6E9F0"
                        :disabled="fieldDisabled(scope.row)"
                        @change="
                            (val) =>
                                switchChange(val, scope.row, 'required_flag')
                        "
                    >
                    </el-switch>
                </template>
            </el-table-column>

            <el-table-column
                prop="action"
                show-overflow-tooltip
                label="操作"
                width="120"
            >
                <template slot-scope="scope">
                    <el-button
                        @click="removeItem(scope.row)"
                        type="text"
                        size="small"
                        :disabled="fieldDisabled(scope.row)"
                    >
                        删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
import Sortable from "sortablejs";
import api from "@/common/api/module/progress_setting";
import _ from "lodash";
import { PRIMARY_COLOR_LIST } from "@/assets/tool/color_list";
import TEXT from "@/assets/tool/text";
export default {
    props: {
        modules: {
            type: String,
            default: ""
        }
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        remindMessage() {
            const message = TEXT.PROGRESS_SETTING_SCREEN;
            return message.UPDATED_REMIND || "";
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
        }
    },
    data() {
        return {
            title: "创建界面：根据创建任务时所需的字段，在下拉列表中进行选择，加入下方列表后，即可设置是否必填或者顺序调整。",
            filed: "",
            filedOptions: [],
            screenTableData: [],
            PRIMARY_COLOR: PRIMARY_COLOR_LIST.PRIMARY_COLOR
        };
    },
    watch: {
        modules(val) {
            if (val) {
                this.fetchFiledOptions();
                this.fetchTableData();
            }
        },
        $route() {
            if (this.modules) {
                this.fetchFiledOptions();
                this.fetchTableData();
            }
        }
    },
    mounted() {
        // 行拖拽排序
        if (this.modules) {
            this.fetchFiledOptions();
            this.fetchTableData();
        }
        this.rowDrop();
    },
    methods: {
        fetchFiledOptions() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                module: this.modules
            };
            api.getRestScreenList(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    res.data.length
                ) {
                    this.filedOptions = this.formatFiled(res.data);
                } else {
                    this.filedOptions = [];
                }
            });
        },
        formatFiled(arr) {
            return arr.map((item) => {
                return {
                    label: item.name,
                    value: item.field_key,
                    key: item.id
                };
            });
        },
        switchChange(val, row, prop) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                id: row.id,
                module: this.modules
                // required: val ? 'yes' : 'no',
                // can_modify: val ? 'yes' : 'no',
            };
            if (prop === "required_flag") {
                params.required = val ? "yes" : "no";

                // if (val) {
                //     // 必填，一起必填
                //     params.can_modify = val ? "yes" : "no";
                //     row.can_modify = val;
                // } else {
                //     // 不必填，只控制自己
                //     params.can_modify = row.can_modify ? "yes" : "no";
                // }
            }
            if (prop === "can_modify") {
                params.can_modify = val ? "yes" : "no";

                if (!val) {
                    // 可以修改
                    params.required = "no";
                    row.required_flag = val;
                } else {
                    // 不可以修改
                    params.required = row.required_flag ? "yes" : "no";
                }
            }
            api.setIsRequired(params).then((res) => {
                if (res && res.resultCode !== 200) {
                    // row.required_flag = !row.required_flag
                }
            });
            // 调接口发送更改后的数据给后端并获取新数据
        },
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
                    let currRow = _this.screenTableData.splice(oldIndex, 1)[0];
                    _this.screenTableData.splice(newIndex, 0, currRow);
                    // 调接口发送更改后的数据给后端并获取新数据
                    if (_this.screenTableData.length < 2) return;
                    let tmpKeyList = _this.screenTableData.map(
                        (item) => item.field_key
                    );
                    let params = {
                        ws_id: _this.curSpace.id,
                        tmpl_id: _this.$route.params.id,
                        module: _this.modules
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
        removeItem(row) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                module: this.modules,
                id: row.id
            };
            api.delScreen(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.fetchFiledOptions();
                    this.fetchTableData();
                }
            });
        },
        filedChange(val) {
            this.filed = "";
            this.createPage(val);
        },
        createPage(val) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                module: this.modules,
                field_key: val
            };
            api.addScreen(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.fetchFiledOptions();
                    this.fetchTableData();
                }
            });
        },
        fetchTableData() {
            // 表格数据
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                module: this.modules
            };
            api.getScreenList(params).then((res) => {
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
                        this.$set(
                            item,
                            "can_modify",
                            item.can_modify === "no" ? false : true
                        );
                    });
                    this.screenTableData = res.data;
                } else {
                    this.screenTableData = [];
                }
            });
            // this.screenTableData = [
            //     {
            //         name: '创建人',
            //         required: true,
            //         id: 1,
            //     },
            //     {
            //         name: 'bug名称',
            //         required: true,
            //         id: 2,
            //     },
            //     {
            //         name: '期望修复日期',
            //         required: true,
            //         id: 3,
            //     },
            //     {
            //         name: 'bug描述',
            //         required: true,
            //         id: 4,
            //     },
            //     {
            //         name: '负责人',
            //         required: true,
            //         id: 5,
            //     },
            // ]
            // 下拉选项
            // this.filedOptions = [
            //     {
            //         label: '创建人',
            //         value: '创建人',
            //     },
            //     {
            //         label: 'bug名称',
            //         value: 'bug名称',
            //     },
            //     {
            //         label: '期望修复日期',
            //         value: '期望修复日期',
            //     },
            //     {
            //         label: 'bug描述',
            //         value: 'bug描述',
            //     },
            //     {
            //         label: '负责人',
            //         value: '负责人',
            //     },
            // ]
        }
    }
};
</script>

<style lang="scss" scoped>
.margin-top-16 {
    margin: 16px 0;
}
.label {
    margin-right: 16px;
}
::v-deep .el-table {
    .el-table__row.not-sort {
        cursor: not-allowed;
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
            color: #c0c4cc;
        }
    }
}
</style>
