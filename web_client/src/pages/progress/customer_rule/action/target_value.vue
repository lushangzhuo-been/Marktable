<template>
    <div class="action-target-border">
        <el-input
            v-if="getActionFieldType(action) === 'text'"
            class="basic-ui value"
            size="small"
            v-model="action.field_value"
            placeholder="请输入目标值"
        ></el-input>
        <el-input
            v-else-if="getActionFieldType(action) === 'number'"
            class="basic-ui value progress-input number hidden-arrow"
            type="number"
            size="small"
            v-model="action.field_value"
            placeholder="请输入目标值"
        ></el-input>
        <div v-else-if="getActionFieldType(action) === 'person'">
            <action-single-people
                v-model="action.field_value"
                :userInfoList="action.field_value_user"
            ></action-single-people>
        </div>
        <div v-else-if="getActionFieldType(action) === 'people'">
            <action-multiple-people
                v-model="action.field_value"
                :userInfoList="action.field_value_user"
            ></action-multiple-people>
        </div>
        <div v-else-if="getActionFieldType(action) === 'status'">
            <action-status-select
                :fieldList="fieldList"
                :actionItem="actionItem"
                v-model="action.field_value"
            ></action-status-select>
        </div>
        <div v-else-if="getActionFieldType(action) === 'select'">
            <action-single-select
                :fieldList="fieldList"
                :actionItem="actionItem"
                v-model="action.field_value"
            ></action-single-select>
        </div>
        <div v-else-if="getActionFieldType(action) === 'multiple-select'">
            <action-multiple-select
                :fieldList="fieldList"
                :actionItem="actionItem"
                v-model="action.field_value"
            ></action-multiple-select>
        </div>
        <div v-else-if="getActionFieldType(action) === 'time'">
            <!-- 时间 -->
            <el-date-picker
                ref="DatePicker"
                popper-class="progress-date-picker-popover"
                v-model="action.field_value"
                type="datetime"
                :editable="false"
                placeholder="选择时间"
                value-format="yyyy-MM-dd HH:mm:ss"
                size="small"
            >
            </el-date-picker>
        </div>
        <div v-else-if="getActionFieldType(action) === 'date'">
            <!-- 日期 -->
            <el-date-picker
                ref="DatePicker"
                popper-class="progress-date-picker-popover"
                v-model="action.field_value"
                type="date"
                :editable="false"
                placeholder="选择日期"
                value-format="yyyy-MM-dd"
                size="small"
            >
            </el-date-picker>
        </div>
        <div v-else-if="getActionFieldType(action) === 'related'">
            <!-- 关系 -->
            <action-related
                v-model="action.field_value"
                :relationInfo="getFieldKeyInfo(action.field_key)"
                :relationReshowInfo="action.field_value_documents"
            ></action-related>
        </div>
        <el-input
            v-else
            class="basic-ui"
            size="small"
            placeholder="请输入目标值"
            v-model="action.field_value"
        ></el-input>
    </div>
</template>

<script>
import _ from "lodash";
import { baseMixin } from "@/mixin.js";
import ActionSinglePeople from "./single_people.vue";
import ActionMultiplePeople from "./multiple_people.vue";
import ActionStatusSelect from "./status_select.vue";
import ActionSingleSelect from "./single_select.vue";
import ActionMultipleSelect from "./multiple_select.vue";
import ActionRelated from "./related.vue";
export default {
    mixins: [baseMixin],
    components: {
        ActionSinglePeople,
        ActionMultiplePeople,
        ActionStatusSelect,
        ActionSingleSelect,
        ActionMultipleSelect,
        ActionRelated
    },
    props: {
        fieldList: {
            type: Array,
            default: () => []
        },
        actionItem: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            action: {}
        };
    },
    watch: {
        actionItem: {
            handler(obj) {
                this.action = obj;
            },
            deep: true,
            immediate: true
        },
        action: {
            handler(obj) {
                this.$emit("input", obj.field_value);
            },
            deep: true
        }
    },
    methods: {
        getFieldKeyInfo(key) {
            let info = _.find(this.fieldList, { field_key: key });
            return info;
        },
        // 动作字段的类型
        getActionFieldType(item) {
            if (_.find(this.fieldList, { field_key: item.field_key })) {
                let detailObj = _.find(this.fieldList, {
                    field_key: item.field_key
                });
                let type = detailObj.mode;
                let expr = detailObj.expr;
                if (
                    type === "text_com" ||
                    type === "textarea_com" ||
                    type === "html_text_com"
                ) {
                    return "text";
                } else if (type === "select_com") {
                    return "select";
                } else if (type === "status_com") {
                    return "status";
                } else if (type === "multi_select_com") {
                    return "multiple-select";
                } else if (type === "person_com" && expr === "multiple") {
                    return "people";
                } else if (type === "person_com" && expr === "single") {
                    return "person";
                } else if (type === "date_com") {
                    return "date";
                } else if (type === "time_com") {
                    return "time";
                } else if (type === "number_com") {
                    return "number";
                } else if (type === "related_com") {
                    return "related";
                } else {
                    return "text";
                }
            }
        }
    }
};
</script>

<style lang="scss" scoped>
.action-target-border {
    .el-input {
        width: 100%;
    }
    .el-select {
        width: 100%;
    }
}
</style>
