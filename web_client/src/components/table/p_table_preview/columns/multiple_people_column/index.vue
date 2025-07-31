<template>
    <el-table-column
        class-name="text"
        :prop="item.columnName"
        :min-width="getWidth(item.mode, item.field_key, item.expr)"
    >
        <!-- :min-width="getWidth(item.mode, item.expr)" -->
        <template slot="header">
            <div>
                <b class="type-box" :style="getType(item.mode, item.expr)"></b>
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
import EditingComponent from "./editing_component2";
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
        getType(type, expr) {
            if (type) {
                return {
                    "background-image": `url(${FieldType[type + "_" + expr]})`
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
