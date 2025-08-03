<template>
    <el-table-column
        class-name="text"
        :min-width="getWidth(item.mode, item.field_key)"
        :prop="item.columnName"
    >
        <template slot="header">
            <div>
                <b class="type-box" :style="getType(item.mode)"></b>
                {{ item.name }}
                <el-tooltip
                    popper-class="basic-ui"
                    effect="dark"
                    placement="top"
                >
                    <div slot="content">
                        {{ item.desc }}
                    </div>
                    <b v-show="item.desc" class="tip-box"></b>
                </el-tooltip>
            </div>
        </template>
        <template slot-scope="scope">
            <editing-component
                :item="item"
                :scope="scope"
                @edit-form-item="editFormItem"
            ></editing-component>
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
        getType(type) {
            if (type) {
                return {
                    "background-image": `url(${FieldType[type]})`
                };
            }
        },
        editFormItem(res, key, id, mode) {
            this.$emit("edit-form-item", res, key, id, mode);
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
