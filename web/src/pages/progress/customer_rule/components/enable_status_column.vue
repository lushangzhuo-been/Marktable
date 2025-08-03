<template>
    <el-table-column
        show-overflow-tooltip
        :min-width="column.minWidth"
        :width="column.width"
        :label="column.name"
    >
        <template slot-scope="scope">
            <div v-if="scope.row[column.prop] === 'yes'">
                <b class="status yes"></b>
                启用中
            </div>
            <div v-if="scope.row[column.prop] === 'no'">
                <b class="status no"></b>
                未启用
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
    components: {},
    data() {
        return {};
    },
    computed: {},
    watch: {},
    methods: {}
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
