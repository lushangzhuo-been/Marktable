<template>
    <el-table-column
        show-overflow-tooltip
        :min-width="column.minWidth"
        :width="column.width"
        :label="column.name"
    >
        <template slot-scope="scope">
            <enable-switch
                :column="column"
                :row="scope.row"
                @enable-change="enableChange"
            ></enable-switch>
        </template>
    </el-table-column>
</template>

<script>
import EnableSwitch from "./enable_switch.vue";
export default {
    props: {
        column: {
            type: Object,
            default: () => {}
        }
    },
    components: {
        EnableSwitch
    },
    data() {
        return {};
    },
    computed: {},
    watch: {},
    methods: {
        enableChange(val, col, row) {
            this.$emit("enable-change", val, col, row);
        }
    }
};
</script>

<style lang="scss" scoped></style>
