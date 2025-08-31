<template>
    <div>
        <el-table
            class="normal-table"
            :data="data"
            style="width: 100%"
            :row-key="rowKey"
        >
            <template v-for="(item, index) in col">
                <interactive-column
                    v-if="item.type === 'interactive'"
                    :name="item.prop"
                    :column="item"
                    :key="index"
                    @interactive="interactive"
                ></interactive-column>
                <slot
                    v-if="item.type === 'slot'"
                    :name="item.prop"
                    :column="item"
                >
                </slot>
                <default-column
                    v-if="!item.type"
                    :key="index"
                    :column="item"
                ></default-column>
            </template>
        </el-table>
    </div>
</template>

<script>
import DefaultColumn from "./columns/default_column/index";
import InteractiveColumn from "./columns/interactive_column/index";
export default {
    components: {
        DefaultColumn,
        InteractiveColumn
    },
    props: {
        col: {
            type: Array,
            default: () => []
        },
        data: {
            type: Array,
            default: () => []
        },
        rowKey: {
            type: String,
            default: "id"
        }
    },
    data() {
        return {
            tableKey: new Date().getTime()
        };
    },
    watch: {
        col: {
            handler() {
                this.tableKey = new Date().getTime();
            },
            immediate: true
        }
    },
    methods: {
        interactive(row, col) {
            this.$emit("interactive", row, col);
        }
    }
};
</script>

<style lang="scss" scoped>
@import "./style.scss";
::v-deep.normal-table.el-table {
    th.el-table__cell > .cell {
        font-weight: 400;
    }
    th.el-table__cell {
        background-color: #f5f6f8;
    }
    .el-table__cell {
        box-sizing: border-box;
        padding: 0;
        height: 40px;
        line-height: 39px;
    }
}
</style>
<style lang="scss">
@import "./drill_style.scss";
</style>
