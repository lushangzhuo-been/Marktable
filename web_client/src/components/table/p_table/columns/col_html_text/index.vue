<template>
    <div>
        <el-table-column
            class-name="text"
            :min-width="getWidth(item.mode, item.field_key)"
            :width="getWidth(item.mode, item.field_key)"
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
                <div
                    class="html-icon column-block"
                    :class="{
                        'editing-effect': item.can_modify === 'yes',
                        'col-required': item.required === 'yes'
                    }"
                >
                    <img
                        :src="getHtmlImg(scope.row)"
                        :width="20"
                        :height="20"
                        @click="openHtmlDialog(scope.row)"
                    />
                </div>
            </template>
        </el-table-column>
    </div>
</template>

<script>
import { FieldType, FieldColumnWidth } from "@/assets/tool/const";
import { baseMixin } from "@/mixin.js";
export default {
    mixins: [baseMixin],
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
    computed: {
        getHtmlImg() {
            return (row) => {
                if (
                    row &&
                    this.item &&
                    Object.keys(this.item).length &&
                    row[this.item.field_key]
                ) {
                    return require(`@/assets/image/progress_column/col_html_content.png`);
                }
                return require(`@/assets/image/progress_column/col_html_empty.png`);
            };
        }
    },
    watch: {},
    mounted() {},
    methods: {
        openHtmlDialog(row) {
            this.$emit("open-html-col", row, this.item);
        },
        getType(type) {
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
.html-icon {
    text-align: center;
    img {
        vertical-align: middle;
        cursor: pointer;
    }
}
.column-block {
    box-sizing: border-box;
    height: 40px;
    line-height: 40px;
    border: 1px solid rgba(0, 0, 0, 0);
}
</style>
