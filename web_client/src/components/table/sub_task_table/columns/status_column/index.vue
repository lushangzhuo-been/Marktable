<template>
    <el-table-column
        show-overflow-tooltip
        :min-width="120"
        :fixed="item.field_key === 'title'"
    >
        <template slot="header">
            <div>
                <b class="type-box" :style="getType(item)"></b>
                {{ item.name }}
            </div>
        </template>
        <template slot-scope="scope">
            <div class="detail">
                <div
                    class="status"
                    v-overflow
                    :style="{
                        backgroundColor: rgbToRgba(
                            `${scope.row[item.field_key].color || '#fff'}`,
                            0.3
                        ),
                        color: '#171e31'
                    }"
                >
                    {{ scope.row[item.field_key].name || emptySpace() }}
                </div>
            </div>
        </template>
    </el-table-column>
</template>

<script>
import { FieldType } from "@/assets/tool/const";
import { baseMixin } from "@/mixin.js";
import { rgbToRgba } from "@/assets/tool/func";
export default {
    mixins: [baseMixin],
    components: {},
    props: {
        item: {
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
        rgbToRgba(color, opacity) {
            return rgbToRgba(color, opacity);
        },
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
@import "@/components/table/p_table_preview/style.scss";
.detail {
    height: 24px;
    box-sizing: border-box;
    .status {
        display: inline-block;
        box-sizing: border-box;
        height: 24px;
        line-height: 22px;
        padding: 0 12px;
        font-size: 14px;
        background: #ffffff;
        border-radius: 4px;
        color: #fff;
        text-overflow: ellipsis;
        overflow: hidden;
        white-space: nowrap;
        max-width: 100%;
    }
}
</style>
<style lang="scss">
.text.el-table__cell {
    .cell {
        padding: 0;
    }
}
</style>
