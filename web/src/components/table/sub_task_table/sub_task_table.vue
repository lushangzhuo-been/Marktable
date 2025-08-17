<template>
    <div>
        <el-table class="progress-table" :data="data" style="width: 100%">
            <!-- @selection-change="handleSelectionChange" -->
            <!-- @header-dragend="headerDragend" -->
            <!-- #region -->
            <template v-for="(item, index) in col">
                <!-- 标题 创建时间 最后更新时间 -->
                <main-text-column
                    v-if="item.field_key === 'title'"
                    :key="index"
                    :item="item"
                    @open-detail="openDetail"
                ></main-text-column>
                <default-column
                    v-if="
                        item.field_key === 'created_at' ||
                        item.field_key === 'updated_at'
                    "
                    :key="index"
                    :item="item"
                ></default-column>
                <!-- 状态 -->
                <status-column
                    v-if="item.field_key === 'status'"
                    :key="index"
                    :item="item"
                ></status-column>
                <!-- 处理人/多人 创建人/单人 -->
                <multiple-people-column
                    v-if="
                        item.field_key === 'handler' ||
                        item.field_key === 'creator'
                    "
                    :key="index"
                    :item="item"
                ></multiple-people-column>
                <!-- 操作列 -->
                <operation-column
                    v-if="item.field_key === 'operation'"
                    :key="index"
                    :item="item"
                    :subTmplAuth="subTmplAuth"
                    @delete-row="deleteSubTask"
                ></operation-column>
            </template>
        </el-table>
    </div>
</template>

<script>
import DadaHandle from "./data_handle";
import MainTextColumn from "./columns/main_text_column/index";
import DefaultColumn from "./columns/default_column/index";
import MultiplePeopleColumn from "./columns/multiple_people_column/index";
import OperationColumn from "./columns/operation_column/index";
import StatusColumn from "./columns/status_column/index";
export default {
    components: {
        MainTextColumn,
        DefaultColumn,
        MultiplePeopleColumn,
        OperationColumn,
        StatusColumn
    },
    props: {
        data: {
            type: Array,
            default: () => []
        },
        tableName: {
            type: String,
            default: ""
        },
        subTmplAuth: {
            type: Boolean,
            default: false
        }
    },

    data() {
        return {
            col: DadaHandle.subTaskCol
        };
    },
    computed: {},
    methods: {
        openDetail(scope) {
            this.$emit("open-detail", scope);
        },
        deleteSubTask(row) {
            this.$emit("delete-sub-task", row);
        }
    }
};
</script>

<style lang="scss" scoped>
@import "@/components/table/p_table/style.scss";
</style>
<style lang="scss">
@import "@/components/table/p_table/drill_style.scss";
</style>
