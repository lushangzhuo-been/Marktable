<template>
    <el-table-column
        show-overflow-tooltip
        :min-width="column.minWidth"
        :width="column.width"
        :label="column.name"
    >
        <template slot-scope="scope">
            <div>
                {{ getTypeName(scope.row[column.prop]) }}
            </div>
        </template>
    </el-table-column>
</template>

<script>
import _ from "lodash";
export default {
    props: {
        column: {
            type: Object,
            default: () => {}
        }
    },
    components: {},
    data() {
        return {
            ruleType: [
                {
                    value: "update",
                    label: "字段变更"
                },
                {
                    value: "schedule",
                    label: "定时触发"
                },
                {
                    value: "create",
                    label: "创建任务"
                },
                {
                    value: "delete",
                    label: "删除任务"
                }
            ]
        };
    },
    computed: {},
    watch: {},
    methods: {
        getTypeName(value) {
            if (_.find(this.ruleType, { value })) {
                return _.find(this.ruleType, { value }).label;
            } else {
                return "--";
            }
        }
    }
};
</script>

<style lang="scss" scoped>
.status {
    display: inline-block;
    width: 8px;
    height: 8px;
    border-radius: 50%;
    margin-right: 2px;
    &.yes {
        background-color: #1890ff;
    }
    &.no {
        background-color: #ffa414;
    }
}
</style>
