<template>
    <el-pagination
        class="basic-ui"
        :popper-class="`basic-ui-pagination-popper ${popperClass}`"
        @size-change="sizeChange"
        @current-change="currentChange"
        :current-page="page"
        :page-sizes="pageSizes"
        :page-size="size"
        :layout="layout"
        :total="total"
    >
    </el-pagination>
</template>

<script>
export default {
    props: {
        popperClass: {
            type: String,
            default: ""
        },
        pageSizes: {
            type: Array,
            default: () => {
                return [10, 15, 20, 50, 100];
            }
        },
        page: {
            type: Number,
            default: 1
        },
        size: {
            type: Number,
            default: 10
        },
        total: {
            type: Number,
            default: 0
        },
        layout: {
            type: String,
            default: "total, sizes, prev, pager, next"
        }
    },
    data() {
        return {};
    },
    methods: {
        // 切换每页条数
        sizeChange(size) {
            this.$emit("pageSizeChange", size);
        },
        // 切换页码
        currentChange(page) {
            this.$emit("curPageChange", page);
        }
    }
};
</script>

<style lang="scss">
.basic-ui-pagination-popper.el-select-dropdown.el-popper
    .el-scrollbar
    .el-select-dropdown__item.selected {
    color: var(--PRIMARY_COLOR);
}
</style>
