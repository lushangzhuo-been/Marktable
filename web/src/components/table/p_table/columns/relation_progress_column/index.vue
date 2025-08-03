<template>
    <!-- show-overflow-tooltip -->
    <el-table-column
        class-name="main-text"
        :min-width="getWidth(item.mode, item.field_key)"
        :width="getWidth(item.mode, item.field_key)"
        :prop="item.columnName"
    >
        <!-- :min-width="item.minWidth" -->
        <template slot="header">
            <div class="main-text-head">
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
                @open-detail="openDetail"
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
        openDetail(scope, type) {
            this.$emit("open-detail", scope, type);
        },
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
.main-text-head {
    padding-left: 10px;
}
</style>
<style lang="scss">
.main-text.el-table__cell {
    .cell {
        padding-left: 0 !important;
        padding-right: 0;
    }
}
</style>
