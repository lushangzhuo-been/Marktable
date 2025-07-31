<template>
    <div>
        <div>
            编辑条件：
            <el-button
                @click="addFilter"
                class="basic-ui add-condition"
                size="small"
            >
                <b class="add-condition-icon"></b>
                添加条件
            </el-button>
        </div>
        <div class="filter-list">
            <div
                class="filter-list-item"
                v-for="(filterItem, filterIndex) in filterGroup"
                :key="filterIndex"
            >
                <!-- 条件间的关系 -->
                <div class="and-or-block" v-show="filterGroup.length > 1">
                    <span class="and-or-content" v-show="filterIndex !== 1">{{
                        filterIndex === 0
                            ? "当"
                            : andOr === "filter_and"
                            ? "且"
                            : "或"
                    }}</span>
                    <el-select
                        v-show="filterIndex === 1"
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
                <!-- 字段 -->
                <el-select
                    class="basic-ui field"
                    size="small"
                    v-model="filterItem.field_key"
                    placeholder="请选择字段"
                    @change="(val) => filterFieldChange(val, filterIndex)"
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
                        v-for="item in getFilterRelation(filterIndex)"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    >
                    </el-option>
                </el-select>
                <!-- 关键词 -->
                <!-- 根据类型有 字符串，单人，单选，； 暂无多人，多选 -->
                <div>
                    <el-input
                        v-if="getFilterType(filterIndex) === 'text'"
                        class="basic-ui value"
                        size="small"
                        v-model="filterItem.value"
                        placeholder="请输入条件"
                    ></el-input>
                    <el-input
                        v-else-if="getFilterType(filterIndex) === 'number'"
                        class="basic-ui value progress-input number hidden-arrow"
                        type="number"
                        size="small"
                        v-model="filterItem.value"
                        placeholder="请输入条件"
                    ></el-input>
                    <div v-else-if="getFilterType(filterIndex) === 'person'">
                        <!-- 单人 -->
                        <person-item
                            :param="filterItem.value"
                            v-model="filterItem.value"
                        ></person-item>
                    </div>
                    <div v-else-if="getFilterType(filterIndex) === 'select'">
                        <new-select-item
                            :option="getFilterOption(filterIndex)"
                            :param="filterItem.value"
                            v-model="filterItem.value"
                        >
                        </new-select-item>
                    </div>
                    <div v-else-if="getFilterType(filterIndex) === 'time'">
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
                    <div v-else-if="getFilterType(filterIndex) === 'date'">
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
                <b class="delete-box" @click="deleteItem(filterIndex)"></b>
            </div>
        </div>
    </div>
</template>

<script>
import api from "@/common/api/module/progress";
import { baseMixin } from "@/mixin.js";
// src\pages\progress\filter\filter_item\new_select.vue
import NewSelectItem from "@/pages/progress/filter/filter_item/new_select";
import PersonItem from "@/pages/progress/filter/filter_item/person";
export default {
    mixins: [baseMixin],
    components: {
        NewSelectItem,
        PersonItem
    },
    props: {
        curId: {
            type: [String, Number],
            default: ""
        }
    },
    data() {
        return {
            filterGroup: [], // 列表
            filterConfig: [], // 配置
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
            filterParams: [] //入参
        };
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        }
    },
    watch: {
        filterGroup: {
            handler() {
                // this.checkFilter();
                // 转化为入参的格式
                this.changeFilterParams();
            },
            deep: true
        }
    },

    mounted() {
        this.getFilterConfig();
    },
    methods: {
        changeFilterParams() {
            let filterParams = _.cloneDeep(this.filterGroup);
            if (filterParams.length) {
                let effectiveParamsNum = 0;
                filterParams.forEach((filterItem) => {
                    if (filterItem.field_key && filterItem.value) {
                        // 查询类型，人取id  选择取value
                        let key = filterItem.field_key;
                        let type = _.find(this.filterConfig, {
                            field_key: key
                        }).mode;
                        if (type === "person_com") {
                            // 单人
                            filterItem.value = filterItem.value[0].id;
                        } else if (
                            type === "select_com" ||
                            type === "multi_select_com"
                        ) {
                            filterItem.value;
                        } else if (type === "number_com") {
                            filterItem.value = Number(filterItem.value);
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
                this.filterParams = filterParams;
                this.effectiveParamsNum = effectiveParamsNum;
                if (effectiveParamsNum) {
                    // 防抖
                    this.emitFilterParams();
                }
            } else {
                this.filterParams = [];
                this.$emit("filter-params", [], [], this.andOr);
            }
        },
        emitFilterParams: _.debounce(function () {
            this.$emit(
                "filter-params",
                this.filterParams,
                this.filterGroup,
                this.andOr
            );
        }, 400),
        getFilterConfig() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                step_id: this.curId
            };
            api.getFilterConfig(params).then((res) => {
                if (res.resultCode === 200) {
                    this.filterConfig = res.data;
                    // 回显并查询
                }
            });
        },
        addFilter() {
            let obj = {
                field_key: "",
                op: "",
                value: ""
            };
            this.filterGroup.push(obj);
        },
        filterFieldChange(val, index) {
            // 重置对应的关系，值
            this.filterGroup[index].op = "";
            this.filterGroup[index].value = "";
        },
        getFilterOption(index) {
            let field_key = this.filterGroup[index].field_key;
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
        deleteItem(index) {
            this.filterGroup.splice(index, 1);
        },
        andOrChange() {
            // 关系手动切换需调接口入参
            this.$emit(
                "filter-params",
                this.filterParams,
                this.filterGroup,
                this.andOr
            );
        },
        getFilterRelation(index) {
            let field_key = this.filterGroup[index].field_key;
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
        getFilterType(index) {
            let field_key = this.filterGroup[index].field_key;
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
                } else {
                    return "text";
                }
            }
            return "text";
        }
    }
};
</script>

<style lang="scss" scoped>
.basic-ui.el-button.add-condition {
    color: #1890ff;
    border: 1px solid #1890ff;
    &:hover {
        border: 1px solid #1890ff;
    }
}
.add-condition-icon {
    display: inline-block;
    width: 10px;
    height: 10px;
    background-size: 100% 100%;
    background-image: url(@/assets/image/progress/add-card-active.png);
}
.filter-list {
    padding-left: 72px;
    margin-top: 16px;
    .filter-list-item {
        height: 32px;
        line-height: 32px;
        display: flex;
        margin-bottom: 12px;
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
</style>
