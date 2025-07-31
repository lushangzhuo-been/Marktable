<template>
    <!-- show-overflow-tooltip -->
    <el-table-column
        :min-width="getWidth(item.mode, item.field_key)"
        :width="getWidth(item.mode, item.field_key)"
        class-name="text"
        :prop="item.columnName"
    >
        <!-- :min-width="item.minWidth" -->
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
.el-progress.ptable-progress {
    // height: 12px;
    .el-progress-bar {
        .el-progress-bar__outer {
            height: 10px !important;
            border-radius: 0px;
            background-color: rgb(0, 0, 0, 0) !important;
            border: 1px solid var(--PRIMARY_COLOR);
            .el-progress-bar__inner {
                height: 10px !important;
                border-radius: 0px;
                background-color: var(--PRIMARY_COLOR) !important;
            }
        }
    }
    .el-progress__text {
        color: --color-1890ff !important;
        font-size: 12px !important;
    }
}
</style>
