<template>
    <el-table-column
        show-overflow-tooltip
        :min-width="column.minWidth"
        :width="column.width"
        :align="column.align || 'left'"
    >
        <template slot="header">
            <div v-if="column.iconCol">
                <b
                    class="type-box"
                    :style="getType(column.mode, column.expr)"
                ></b>
                {{ column.name }}
            </div>
            <div v-else>{{ column.name }}</div>
        </template>
        <template slot-scope="scope">
            {{ scope.row[column.prop] || "--" }}
        </template>
    </el-table-column>
</template>

<script>
import { FieldType } from "@/assets/tool/const";
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
        getType(type, expr = "") {
            if (!type) return;
            if (type && !expr) {
                return {
                    "background-image": `url(${FieldType[type]})`
                };
            }
            if (type && expr) {
                return {
                    "background-image": `url(${FieldType[type + "_" + expr]})`
                };
            }
        }
    }
};
</script>

<style lang="scss" scoped>
// @import "../../style.scss";
.type-box {
    display: inline-block;
    width: 20px;
    height: 20px;
    background-size: 100% 100%;
    vertical-align: middle;
    position: relative;
    top: -2px;
}
</style>
<style lang="scss">
// .text.el-table__cell {
//     .cell {
//         padding: 0;
//     }
// }
</style>
