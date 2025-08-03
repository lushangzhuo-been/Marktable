<template>
    <div>
        <el-popover
            ref="Filter"
            placement="bottom-start"
            popper-class="filter-popove"
            trigger="click"
            @show="popoverShow"
            @hide="popoverHide"
        >
            <!-- 遍历item.list  做下拉内容 -->
            <div class="popover-content">
                <div class="filter-list">
                    <div
                        v-for="(groupItem, groupIndex) in filterGroup"
                        :key="groupIndex"
                        class="group-border"
                    >
                        <!-- 组间的关系 -->
                        <div
                            class="and-or-block"
                            v-show="filterGroup.length > 1"
                        >
                            <span
                                class="and-or-content"
                                v-show="groupIndex !== 1"
                                >{{
                                    groupIndex === 0
                                        ? "当"
                                        : andOr === "filter_and"
                                        ? "且"
                                        : "或"
                                }}</span
                            >
                            <el-select
                                v-show="groupIndex === 1"
                                class="basic-ui and-or"
                                size="small"
                                v-model="andOr"
                                @input="andOrChange"
                            >
                                <el-option
                                    v-for="item in andOrList"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.value"
                                >
                                </el-option>
                            </el-select>
                        </div>
                        <!-- 二级filter -->
                        <div class="filter-group-border">
                            <div
                                class="filter-list-item"
                                v-for="(
                                    filterItem, filterIndex
                                ) in groupItem.data"
                                :key="filterIndex"
                            >
                                <!-- filter-item之间的关系 -->
                                <div class="and-or-block">
                                    <span
                                        class="and-or-content"
                                        v-show="filterIndex !== 1"
                                        >{{
                                            filterIndex === 0
                                                ? "当"
                                                : groupItem.lor === "filter_and"
                                                ? "且"
                                                : "或"
                                        }}</span
                                    >
                                    <el-select
                                        v-show="filterIndex === 1"
                                        class="basic-ui and-or"
                                        size="small"
                                        v-model="groupItem.lor"
                                        @input="andOrChange"
                                    >
                                        <el-option
                                            v-for="item in andOrList"
                                            :key="item.value"
                                            :label="item.label"
                                            :value="item.value"
                                        >
                                        </el-option>
                                    </el-select>
                                </div>
                                <div class="filter-list-item-relation-filter">
                                    <!-- 字段 -->
                                    <el-select
                                        class="basic-ui field"
                                        size="small"
                                        v-model="filterItem.field_key"
                                        placeholder="请选择字段"
                                        @change="
                                            (val) =>
                                                filterFieldChange(
                                                    val,
                                                    filterItem,
                                                    filterIndex
                                                )
                                        "
                                    >
                                        <el-option
                                            v-for="item in filterConfig"
                                            :key="item.field_key"
                                            :label="item.name"
                                            :value="item.field_key"
                                        >
                                        </el-option>
                                    </el-select>
                                    <!-- 关系 -->
                                    <el-select
                                        class="basic-ui relationship"
                                        size="small"
                                        v-model="filterItem.op"
                                        placeholder="请选择关系"
                                    >
                                        <el-option
                                            v-for="item in getFilterRelation(
                                                filterItem
                                            )"
                                            :key="item.value"
                                            :label="item.label"
                                            :value="item.value"
                                        >
                                        </el-option>
                                    </el-select>
                                    <!-- 关键词 -->
                                    <!-- 根据类型有 字符串，单人，单选，； 暂无多人，多选 -->
                                    <div class="filter-itme-value">
                                        <el-input
                                            v-if="
                                                getFilterType(filterItem) ===
                                                'text'
                                            "
                                            class="basic-ui value"
                                            size="small"
                                            v-model="filterItem.value"
                                            placeholder="请输入条件"
                                        ></el-input>
                                        <el-input
                                            v-else-if="
                                                getFilterType(filterItem) ===
                                                'number'
                                            "
                                            class="basic-ui value progress-input number hidden-arrow"
                                            type="number"
                                            size="small"
                                            v-model="filterItem.value"
                                            placeholder="请输入条件"
                                        ></el-input>
                                        <!-- 关联关系 -->
                                        <relation-item
                                            v-else-if="
                                                getFilterType(filterItem) ===
                                                'related'
                                            "
                                            :configData="{
                                                filterConfig: filterConfig,
                                                filterItem: filterItem
                                            }"
                                            :param="filterItem.value"
                                            v-model="filterItem.value"
                                        ></relation-item>
                                        <div
                                            v-else-if="
                                                getFilterType(filterItem) ===
                                                'person'
                                            "
                                        >
                                            <!-- 单人 -->
                                            <person-item
                                                :param="filterItem.value"
                                                v-model="filterItem.value"
                                            ></person-item>
                                        </div>
                                        <div
                                            v-else-if="
                                                getFilterType(filterItem) ===
                                                'select'
                                            "
                                        >
                                            <new-select-item
                                                :option="
                                                    getFilterOption(filterItem)
                                                "
                                                :param="filterItem.value"
                                                v-model="filterItem.value"
                                            >
                                            </new-select-item>
                                        </div>
                                        <!-- 状态筛选 -->
                                        <div
                                            v-else-if="
                                                getFilterType(filterItem) ===
                                                'status'
                                            "
                                        >
                                            <status-select
                                                :param="filterItem.value"
                                                v-model="filterItem.value"
                                            >
                                            </status-select>
                                        </div>
                                        <div
                                            v-else-if="
                                                getFilterType(filterItem) ===
                                                'time'
                                            "
                                        >
                                            <!-- 时间 -->
                                            <el-date-picker
                                                ref="DatePicker"
                                                class="progress-date-picker time"
                                                popper-class="progress-date-picker-popover"
                                                v-model="filterItem.value"
                                                type="datetime"
                                                :editable="false"
                                                placeholder="选择时间"
                                                value-format="yyyy-MM-dd HH:mm:ss"
                                                size="small"
                                            >
                                            </el-date-picker>
                                        </div>
                                        <div
                                            v-else-if="
                                                getFilterType(filterItem) ===
                                                'date'
                                            "
                                        >
                                            <!-- 日期 -->
                                            <el-date-picker
                                                ref="DatePicker"
                                                class="progress-date-picker date"
                                                popper-class="progress-date-picker-popover"
                                                v-model="filterItem.value"
                                                type="date"
                                                :editable="false"
                                                placeholder="选择日期"
                                                value-format="yyyy-MM-dd"
                                                size="small"
                                            >
                                            </el-date-picker>
                                        </div>
                                    </div>
                                    <b
                                        class="delete-box"
                                        @click="
                                            deleteItem(
                                                groupItem,
                                                filterIndex,
                                                groupIndex
                                            )
                                        "
                                    ></b>
                                </div>
                            </div>
                            <div
                                class="add-filter"
                                @click="addFilter(groupItem)"
                            >
                                <b class="operation-box add"></b>
                                添加条件
                            </div>
                        </div>
                    </div>
                </div>
                <div class="operation-block">
                    <div class="add-filter" @click="addGroup">
                        <b class="operation-box add"></b>
                        添加组
                    </div>
                    <!-- v-show="filterGroup.length" -->
                    <div class="others-operation">
                        <div
                            v-show="filterGroup.length"
                            class="clear-filter"
                            @click="clearFilter"
                        >
                            <b class="operation-box clear"></b>
                            清空
                        </div>
                        <div
                            class="save-filter"
                            :class="{ disabled: filterGroupCheck }"
                            @click="saveFilter"
                        >
                            <b class="operation-box save"></b>
                            {{ currentTab === "-1" ? "保存为新视图" : "保存" }}
                        </div>
                    </div>
                </div>
            </div>
            <span
                slot="reference"
                class="btn-icon"
                :class="{ editing: editing }"
            >
                <b class="filter-box"></b>
                筛选{{ effectiveParamsNum || "" }}
            </span>
        </el-popover>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import NewSelectItem from "./filter_item/new_select";
import StatusSelect from "./filter_item/status_select";
import PersonItem from "./filter_item/person";
import RelationItem from "./filter_item/related";
export default {
    components: {
        NewSelectItem,
        PersonItem,
        StatusSelect,
        RelationItem
    },
    props: {
        viewFilter: {
            type: Array,
            default: () => []
        },
        viewFilterStandby: {
            type: Array,
            default: () => []
        },
        filterRelation: {
            type: String,
            default: "filter_and"
        },
        currentTab: {
            type: String,
            default: ""
        }
    },
    data() {
        return {
            filterGroup: [], // 列表
            filterConfig: [], // 配置
            editing: false,
            andOrList: [
                {
                    label: "且",
                    value: "filter_and"
                },
                {
                    label: "或",
                    value: "filter_or"
                }
            ],
            andOr: "filter_and",
            filterGroupCheck: false, // true为禁用 ， false为可用
            filterParams: [], //入参
            effectiveParamsNum: 0
        };
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        }
    },
    watch: {
        viewFilterStandby: {
            handler() {
                this.getEffectiveParamsNum();
            },
            immediate: true,
            deep: true
        },
        filterRelation: {
            handler(filterRelation) {
                this.andOr = filterRelation;
            }
        },
        // 每添加一条过滤或者过滤的数据变化都会执行
        filterGroup: {
            handler() {
                this.watchFilterGroup();
            },
            deep: true
        },
        curProgress: {
            // 流程切换
            handler() {
                if (this.$refs.Filter) {
                    // 重置参数
                    // this.filterGroup = [];
                    this.filterParams = [];
                    this.$refs.Filter.doClose();
                }
            },
            immediate: true
        }
    },
    mounted() {},
    methods: {
        watchFilterGroup: _.debounce(function () {
            // 两件事  check && 改为入参格式
            this.checkFilter();
            // 转化为入参的格式
            this.changeFilterParams();
        }, 600),
        getEffectiveParamsNum() {
            let effectiveParamsNum = 0;
            this.viewFilterStandby.forEach((groupItem) => {
                groupItem.data.forEach((filterItem) => {
                    if (
                        filterItem.field_key &&
                        filterItem.op &&
                        filterItem.value
                    ) {
                        effectiveParamsNum += 1;
                    }
                });
            });
            this.effectiveParamsNum = effectiveParamsNum;
        },
        andOrChange() {
            // 关系手动切换需调接口入参
            this.$emit(
                "filter-search",
                this.filterParams,
                this.filterGroup,
                this.andOr
            );
            this.$emit("save-filter", this.filterParams, this.andOr);
        },
        closePopver() {
            if (this.$refs.Filter) {
                // 重置参数
                this.$refs.Filter.doClose();
            }
        },
        mountedFilter() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress
            };
            api.getFilterConfig(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.filterConfig = res.data;
                } else {
                    this.filterConfig = [];
                }
                // 回显并查询
                // filterGroup为每一条过滤的数据， viewFilter默认是空数组，在子列表会有回显的值
                this.filterGroup = _.cloneDeep(this.viewFilter);
            });
        },
        changeFilterParams() {
            let filterParams = _.cloneDeep(this.filterGroup);
            if (filterParams.length) {
                let effectiveParamsNum = 0;
                filterParams.forEach((groupItem) => {
                    groupItem.data.forEach((filterItem) => {
                        if (groupItem.data.length > 1) {
                            groupItem.filter_type = "group";
                        } else {
                            groupItem.filter_type = "filter";
                        }
                        if (filterItem.field_key && filterItem.value) {
                            // 查询类型，人取id  选择取value
                            let key = filterItem.field_key;
                            let type = _.find(this.filterConfig, {
                                field_key: key
                            }).mode;
                            if (type === "person_com") {
                                // 单人
                                filterItem.value = filterItem.value[0].id;
                            } else if (type === "related_com") {
                                // 单人
                                filterItem.value =
                                    filterItem.value && filterItem.value.length
                                        ? filterItem.value[0]._id
                                        : "";
                            } else if (
                                type === "select_com" ||
                                type === "multi_select_com"
                            ) {
                                filterItem.value;
                            } else if (type === "number_com") {
                                filterItem.value = Number(filterItem.value);
                            } else if (type === "status_com") {
                                filterItem.value = filterItem.value.id;
                            } else {
                                filterItem.value;
                            }
                        }
                        if (
                            filterItem.field_key &&
                            filterItem.op &&
                            filterItem.value
                        ) {
                            effectiveParamsNum += 1;
                        }
                    });
                });
                this.filterParams = filterParams;
                this.effectiveParamsNum = effectiveParamsNum;
                if (effectiveParamsNum) {
                    // 防抖
                    this.emitFilterParams();
                }
            } else {
                this.filterParams = [];
                this.$emit("filter-search", [], [], this.andOr);
                this.$emit("save-filter", this.filterParams, this.andOr);
            }
        },
        emitFilterParams: _.debounce(function () {
            this.$emit(
                "filter-search",
                this.filterParams,
                this.filterGroup,
                this.andOr
            );
            this.$emit("save-filter", this.filterParams, this.andOr);
        }, 400),
        filterFieldChange(val, item, index) {
            // 重置对应的关系，值
            item.op = "";
            item.value = "";
        },
        getFilterOption(filterItem) {
            let field_key = filterItem.field_key;
            let option;
            if (field_key) {
                option = _.find(this.filterConfig, {
                    field_key: field_key
                }).enum_values;
                if (option) {
                    return JSON.parse(option);
                } else {
                    return [];
                }
            } else {
                return [];
            }
        },
        getFilterRelation(filterItem) {
            let field_key = filterItem.field_key;
            let option;
            if (field_key) {
                option = _.find(this.filterConfig, {
                    field_key: field_key
                }).op;
            }
            if (option) {
                return option;
            } else {
                return [];
            }
        },
        getFilterType(filterItem) {
            let field_key = filterItem.field_key;
            if (
                _.find(this.filterConfig, {
                    field_key: field_key
                })
            ) {
                let type = _.find(this.filterConfig, {
                    field_key: field_key
                }).mode;
                if (
                    type === "text_com" ||
                    type === "textarea_com" ||
                    type === "html_text_com" ||
                    !type
                ) {
                    return "text";
                } else if (
                    type === "select_com" ||
                    type === "multi_select_com"
                ) {
                    return "select";
                } else if (type === "person_com") {
                    return "person";
                } else if (type === "date_com") {
                    return "date";
                } else if (type === "time_com") {
                    return "time";
                } else if (type === "number_com") {
                    return "number";
                } else if (type === "status_com") {
                    return "status";
                } else if (type === "related_com") {
                    return "related";
                } else {
                    return "text";
                }
            }
            return "text";
        },
        getFilterConfig() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress
            };
            api.getFilterConfig(params).then((res) => {
                this.filterConfig = res.data;
            });
        },
        checkFilter() {
            // 检查是否每条filter都有值
            let res = false;
            this.filterGroup.forEach((groupItem) => {
                groupItem.data.forEach((filterItem) => {
                    if (
                        !filterItem.field_key ||
                        !filterItem.op ||
                        !filterItem.value
                    ) {
                        res = true;
                        return res;
                    }
                });
            });
            this.filterGroupCheck = res;
        },
        // 添加组
        addGroup() {
            let obj = {
                lor: "filter_and",
                data: [
                    {
                        field_key: "",
                        op: "",
                        value: ""
                    }
                ]
            };
            this.filterGroup.push(obj);
        },
        // 添加二级条件
        addFilter(groupItem) {
            let obj = {
                field_key: "",
                op: "",
                value: ""
            };
            groupItem.data.push(obj);
        },
        clearFilter() {
            this.filterGroup = [];
        },
        saveFilter() {
            let order = "";
            if (this.currentTab === "-1") {
                order = "add";
            } else {
                order = "update";
            }
            this.$emit("save-filter", this.filterParams, this.andOr, "save");
        },
        deleteItem(groupItem, filterIndex, groupIndex) {
            groupItem.data.splice(filterIndex, 1);
            if (!groupItem.data.length) {
                this.filterGroup.splice(groupIndex, 1);
            }
        },
        popoverShow() {
            this.editing = true;
            this.$emit("get-view-filter");
            // 过滤的poppover打开时， 调用filter的config接口
            this.mountedFilter();
        },
        popoverHide() {
            this.editing = false;
        }
    }
};
</script>

<style lang="scss" scoped>
@import "@/pages/progress/operation.scss";

.btn-icon {
    cursor: pointer;
    color: #171e17;
    padding: 4px 8px;
    border-radius: 4px;
    &:hover {
        // color: var(--PRIMARY_COLOR);
        color: #171e17;
        background-color: rgb(202, 228, 253);
        padding: 4px 8px;
        border-radius: 4px;
        .filter-box {
            display: inline-block;
            width: 20px;
            height: 20px;
            background-size: 100% 100%;
            background-image: url(@/assets/image/progress/filter-active.png);
            vertical-align: middle;
        }
    }
    &.editing {
        // color: var(--PRIMARY_COLOR);
        color: #171e17;
        background-color: rgb(202, 228, 253);
        padding: 4px 8px;
        border-radius: 4px;
        .filter-box {
            display: inline-block;
            width: 20px;
            height: 20px;
            background-size: 100% 100%;
            background-image: url(@/assets/image/progress/filter-active.png);
            vertical-align: middle;
        }
    }
    .filter-box {
        display: inline-block;
        width: 20px;
        height: 20px;
        background-size: 100% 100%;
        background-image: url(@/assets/image/progress/filter.png);
        vertical-align: middle;
    }
}
.view-had-filter {
    .btn-icon {
        cursor: pointer;
        // color: var(--PRIMARY_COLOR);
        color: #171e17;
        background-color: rgb(202, 228, 253);
        padding: 4px 8px;
        border-radius: 4px;
        .filter-box {
            display: inline-block;
            width: 20px;
            height: 20px;
            background-size: 100% 100%;
            background-image: url(@/assets/image/progress/filter-active.png);
            vertical-align: middle;
        }
    }
}
</style>
<style lang="scss">
// @import "@/pages/progress/operation.scss";

// .operation-box {
//     display: inline-block;
//     width: 18px;
//     height: 18px;
//     background-size: 100% 100%;
//     vertical-align: middle;
//     &.add {
//         vertical-align: top;
//         background-image: url(@/assets/image/progress/add.png);
//     }
//     &.clear {
//         background-image: url(@/assets/image/common/delete.svg);
//     }
//     &.save {
//         background-image: url(@/assets/image/progress/save.svg);
//     }
// }
.operation-block {
    display: flex;
    justify-content: space-between;

    .others-operation {
        display: flex;
        justify-content: space-between;
        margin-left: 24px;
        .clear-filter {
            margin-right: 21px;
            cursor: pointer;
            &:hover {
                color: var(--PRIMARY_COLOR);
                .clear {
                    background-image: url(@/assets/image/common/delete-active.svg);
                }
            }
        }
        .save-filter {
            cursor: pointer;
            &:hover {
                color: var(--PRIMARY_COLOR);
                .save {
                    background-image: url(@/assets/image/progress/save-active.svg);
                }
            }
            &.disabled {
                color: #a6abbc;
                cursor: not-allowed;
                .save {
                    background-image: url(@/assets/image/progress/save-disabled.svg);
                }
            }
        }
    }
}

.popover-content {
    .filter-list {
        max-height: 540px;
        overflow-y: auto;
        .filter-list-item {
            height: 32px;
            line-height: 32px;
            display: flex;
            margin-bottom: 12px;
            .filter-list-item-relation-filter {
                display: flex;
            }
            .delete-box {
                display: inline-block;
                width: 18px;
                height: 18px;
                background-size: 100% 100%;
                background-image: url(@/assets/image/progress/delete.png);
                vertical-align: middle;
                cursor: pointer;
                margin: 7px;
                &:hover {
                    background-image: url(@/assets/image/progress/delete-active.png);
                }
            }
        }
    }
}

.el-popper.filter-popove {
    box-sizing: border-box;
    margin-top: 0px;
    padding: 12px 16px;
    .popper__arrow {
        display: none;
    }
}
.and-or-block {
    width: 58px;
    margin-right: 8px;
    .and-or-content {
        padding-left: 15px;
    }
}
.el-select.and-or {
    width: 58px;
    .el-input--suffix {
        .el-input__inner {
            padding-right: 24px;
        }
    }
}
.el-select.field {
    width: 140px;
    margin-right: 24px;
}
.el-select.relationship {
    width: 120px;
    margin-right: 8px;
}
.el-input.value {
    width: 160px;
}
.el-date-editor.el-input.progress-date-picker.time {
    width: 200px;
}
.el-date-editor.el-input.progress-date-picker.date {
    width: 160px;
}
::v-deep .progress-input.el-input {
    .el-input__inner {
        padding: 0;
        border: none;
        // height: 38px;
        // line-height: 38px;
        height: 30px;
        line-height: 30px;
        background-color: rgba(0, 0, 0, 0);
        font-weight: 400;
        font-size: 14px;
        color: #171e31;
        font-family: MiSans, MiSans;
    }
    &.number.hidden-arrow {
        input::-webkit-outer-spin-button,
        input::-webkit-inner-spin-button {
            -webkit-appearance: none;
        }
        input[type="number"] {
            -moz-appearance: textfield;
        }
    }
}
.group-border {
    display: flex;
    margin-bottom: 16px;
    .filter-group-border {
        padding: 10px;
        border-radius: 8px;
        background-color: #fafafb;
    }
}

.filter-itme-value {
    display: inline-block;
}
.add-filter {
    display: inline-block;
    color: var(--PRIMARY_COLOR);
    cursor: pointer;
}
</style>
