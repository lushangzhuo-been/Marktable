<template>
    <div>
        <el-table
            class="progress-table"
            :data="data"
            :key="tableKey"
            border
            style="width: 100%"
            @selection-change="handleSelectionChange"
            @header-dragend="headerDragend"
        >
            <!-- #region -->
            <template v-for="(item, index) in col">
                <el-table-column
                    v-if="item.mode === 'selection'"
                    :key="index"
                    :selectable="selectable"
                    type="selection"
                    width="40"
                >
                </el-table-column>
                <!-- 操作列 -->
                <operation-column
                    :key="index"
                    v-else-if="item.mode === 'operation'"
                    :item="item"
                    @copy-link="copyLink"
                    @delete-row="deleteRow"
                ></operation-column>
                <!-- 标题title列 -->
                <main-text-column
                    :key="index"
                    v-if="
                        item.mode === 'text_com' && item.field_key === 'title'
                    "
                    :item="item"
                    @open-detail="openDetail"
                    @edit-form-item="editFormItem"
                ></main-text-column>
                <!-- 关联流程列 -->
                <relation-progress-column
                    :key="index"
                    v-if="item.mode === 'related_com'"
                    :item="item"
                    @open-detail="openDetail"
                    @edit-form-item="editFormItem"
                ></relation-progress-column>
                <!-- 纯文本列 -->
                <text-column
                    :key="index"
                    v-if="
                        item.mode === 'text_com' && item.field_key !== 'title'
                    "
                    :item="item"
                    @edit-form-item="editFormItem"
                ></text-column>
                <number-column
                    :key="index"
                    v-if="item.mode === 'number_com'"
                    :item="item"
                    @edit-form-item="editFormItem"
                ></number-column>
                <!-- 下拉-单选 -->
                <col-radio-select
                    :key="index"
                    v-if="item.mode === 'select_com'"
                    :item="item"
                    @edit-form-item="editFormItem"
                ></col-radio-select>
                <!-- 下拉-多选 -->
                <col-multi-select
                    :key="index"
                    v-if="item.mode === 'multi_select_com'"
                    :item="item"
                    @edit-form-item="editFormItem"
                ></col-multi-select>
                <single-people-column
                    :key="index"
                    v-if="item.mode === 'person_com' && item.expr === 'single'"
                    :item="item"
                    @edit-form-item="editFormItem"
                ></single-people-column>
                <multiple-people-column
                    :key="index"
                    v-if="
                        item.mode === 'person_com' && item.expr === 'multiple'
                    "
                    :item="item"
                    @edit-form-item="editFormItem"
                ></multiple-people-column>
                <!-- 日期列 -->
                <date-column
                    :key="index"
                    v-if="item.mode === 'date_com'"
                    :item="item"
                    @edit-form-item="editFormItem"
                ></date-column>
                <!-- 可编辑的时间列 -->
                <date-time-column
                    :key="index"
                    v-if="
                        item.mode === 'time_com' &&
                        item.field_key !== 'created_at' &&
                        item.field_key !== 'updated_at'
                    "
                    :item="item"
                    @edit-form-item="editFormItem"
                ></date-time-column>
                <rich-text-column
                    :key="index"
                    v-if="item.mode === 'textarea_com'"
                    :item="item"
                    @edit-form-item="editFormItem"
                ></rich-text-column>
                <!-- html编辑器 -->
                <col-html-text
                    :key="index"
                    v-if="item.mode === 'html_text_com'"
                    :item="item"
                    @open-html-col="openHtmlCol"
                ></col-html-text>
                <progress-column
                    :key="index"
                    v-if="item.mode === 'progress_com'"
                    :item="item"
                    @edit-form-item="editFormItem"
                ></progress-column>
                <link-column
                    :key="index"
                    v-if="item.mode === 'link_com'"
                    :item="item"
                    @edit-form-item="editFormItem"
                ></link-column>
                <!-- blank列-不可编辑列 -->
                <blank-column
                    :key="index"
                    v-if="
                        item.field_key === 'created_at' ||
                        item.field_key === 'updated_at'
                    "
                    :item="item"
                ></blank-column>
                <!-- 状态列，单独一列可编辑，以前为blank列 -->
                <status-column
                    :key="index"
                    v-if="item.mode === 'status_com'"
                    :item="item"
                    @refresh-table-data="refreshTableData"
                ></status-column>
            </template>
            <!-- #endregion -->
        </el-table>
    </div>
</template>

<script>
import MainTextColumn from "./columns/main_text_column/index";
import RelationProgressColumn from "./columns/relation_progress_column/index"; // 关联关系
import TextColumn from "./columns/text_column/index";
import NumberColumn from "./columns/number_column/index";
import ColRadioSelect from "./columns/col_radio_select/index";
import ColMultiSelect from "./columns/col_multi_select/index";
import ColHtmlText from "./columns/col_html_text/index";
import MultiplePeopleColumn from "./columns/multiple_people_column/index";
import SinglePeopleColumn from "./columns/single_people_column/index";
import DateColumn from "./columns/date_column/index";
import DateTimeColumn from "./columns/date_time_column/index";
import RichTextColumn from "./columns/rich_text_column/index";
import ProgressColumn from "./columns/progress_column/index";
import LinkColumn from "./columns/link_column/index";
import BlankColumn from "./columns/blank_column/index";
import StatusColumn from "./columns/status_column/index"; //状态列
import TimeColumn from "./columns/time_column/index"; //时间日期列
import api from "@/common/api/module/progress";
import OperationColumn from "./columns/operation_column/index"; //时间列
export default {
    components: {
        MainTextColumn,
        RelationProgressColumn,
        TextColumn,
        NumberColumn,
        // StatusColumn,
        ColRadioSelect,
        ColMultiSelect,
        ColHtmlText,
        MultiplePeopleColumn,
        SinglePeopleColumn,
        DateColumn,
        DateTimeColumn,
        RichTextColumn,
        ProgressColumn,
        LinkColumn,
        BlankColumn,
        StatusColumn,
        TimeColumn,
        OperationColumn
    },
    props: {
        col: {
            type: Array,
            default: () => []
        },
        data: {
            type: Array,
            default: () => []
        }
    },
    data() {
        return {
            multipleSelection: [],
            tableKey: ""
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
    methods: {
        // 操作列复制连接
        copyLink(row) {},
        // 操作列删除行
        deleteRow(row) {
            this.$emit("delete-row", row);
        },
        selectable(row, index) {
            if (row.permission && row.permission.delete === "no") {
                return false; //禁用
            } else {
                return true; //可选
            }
        },
        openHtmlCol(row, col) {
            this.$emit("open-html-col", row, col);
        },
        refreshTableKey() {
            this.tableKey = JSON.stringify(new Date().getTime());
        },
        handleSelectionChange(arr) {
            this.multipleSelection = arr;
            this.$emit("multiple-selection", this.multipleSelection);
        },
        headerDragend(newWidth, oldWidth, column, event) {
            this.$emit("change-column-width", column.property, newWidth);
        },
        openDetail(scope, type) {
            this.$emit("open-detail", scope, type);
        },
        editFormItem(res, key, id, mode) {
            let file_key = key;
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: id,
                field_key: key,
                [file_key]: res
            };
            api.updateProgress(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.$emit("refresh");
                } else {
                    this.$emit("refresh");
                }
            });
        },
        // 状态列更新后，直接刷新表格数据
        refreshTableData() {
            this.$emit("refresh");
        }
    }
};
</script>

<style lang="scss" scoped>
@import "./style.scss";
</style>
<style lang="scss">
@import "./drill_style.scss";
</style>
