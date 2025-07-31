<template>
    <el-table-column
        show-overflow-tooltip
        :min-width="column.minWidth"
        :width="column.width"
        :label="column.name"
    >
        <template slot-scope="scope">
            <div>
                <el-button
                    @click="operation(column, scope.row, 'edit')"
                    type="text"
                    size="small"
                >
                    编辑
                </el-button>
                <el-button
                    @click="operation(column, scope.row, 'log')"
                    type="text"
                    size="small"
                >
                    日志
                </el-button>
                <el-button
                    @click="operation(column, scope.row, 'copy')"
                    type="text"
                    size="small"
                >
                    复制
                </el-button>
                <el-button
                    @click="operation(column, scope.row, 'delete')"
                    type="text"
                    size="small"
                >
                    删除
                </el-button>
            </div>
        </template>
    </el-table-column>
</template>

<script>
export default {
    props: {
        column: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {};
    },
    methods: {
        operation(col, row, type) {
            this.$emit("list-item-operation", col, row, type);
        }
    }
};
</script>

<style lang="scss" scoped></style>
