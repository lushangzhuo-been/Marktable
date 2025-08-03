<template>
    <el-table-column type="expand" :fixed="true">
        <template slot-scope="props">
            <slot-content></slot-content>
        </template>
    </el-table-column>
</template>

<script>
import SlotContent from "./slot_content.vue";
export default {
    components: {
        SlotContent
    },
    data() {
        return {};
    },
    watch: {},
    mounted() {},
    methods: {}
};
</script>

<style lang="scss" scoped>
// @import "../../style.scss";
</style>
<style lang="scss">
// .text.el-table__cell {
//     .cell {
//         padding: 0;
//     }
// }
</style>
