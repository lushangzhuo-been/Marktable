<template>
    <el-table-column
        show-overflow-tooltip
        class-name="text"
        :min-width="getWidth(item.mode, item.field_key)"
        :width="getWidth(item.mode, item.field_key)"
        :prop="item.columnName"
    >
        <template slot="header">
            <div>
                <b class="type-box" :style="getType(item)"></b>
                {{ item.name }}
            </div>
        </template>
        <template slot-scope="scope">
            <editing-component :item="item" :scope="scope"></editing-component>
        </template>
    </el-table-column>
</template>

<script>
import EditingComponent from "./editing_component";
import { FieldType, FieldColumnWidth } from "@/assets/tool/const";
import { baseMixin } from "@/mixin.js";
export default {
    mixins: [baseMixin],
    components: {
        EditingComponent
    },
    props: {
        item: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            isEditing: false
        };
    },
    watch: {},
    mounted() {},
    methods: {
        getType(item) {
            // 按照类型引对应的icon
            let type;
            if (
                item.field_key === "created_at" ||
                item.field_key === "updated_at"
            ) {
                type = "time_com";
            } else if (item.field_key === "status") {
                type = "select_com";
            } else if (item.field_key === "creator") {
                type = "person_com_single";
            }
            if (type) {
                return {
                    "background-image": `url(${FieldType[type]})`
                };
            }
        }
    }
};
</script>

<style lang="scss" scoped>
@import "../../style.scss";
</style>
<style lang="scss">
.text.el-table__cell {
    .cell {
        padding: 0;
    }
}
</style>
