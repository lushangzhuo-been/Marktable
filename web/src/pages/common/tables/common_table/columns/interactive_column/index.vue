<template>
    <el-table-column
        show-overflow-tooltip
        :min-width="column.minWidth"
        :width="column.width"
        :align="column.align || 'left'"
    >
        <template slot="header">
            <div>{{ column.name }}</div>
        </template>
        <template slot-scope="scope">
            <span
                v-if="scope.row[column.prop]"
                class="interactive"
                @click="interactive(scope.row, column)"
                >{{ scope.row[column.prop] }}</span
            >
            <span v-else>--</span>
        </template>
    </el-table-column>
</template>

<script>
export default {
    components: {},
    props: {
        column: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {};
    },
    watch: {},
    mounted() {},
    methods: {
        interactive(row, col) {
            this.$emit("interactive", row, col);
        }
    }
};
</script>

<style lang="scss" scoped>
.interactive {
    color: #1890ff;
    cursor: pointer;
    &:hover {
        text-decoration: underline;
    }
}
</style>
