<template>
    <el-table-column
        show-overflow-tooltip
        :min-width="widthMap[item.field_key] || 220"
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
                <div class="detail-content">
                    {{ scope.row[item.field_key] }}
                </div>
            </div>
        </template>
    </el-table-column>
</template>

<script>
import { FieldType } from "@/assets/tool/const";
import { baseMixin } from "@/mixin.js";
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
        return {
            widthMap: {
                title: 220,
                created_at: 160,
                updated_at: 160
            }
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
            } else if (item.field_key === "title") {
                type = "text_com";
            } else if (item.field_key === "status") {
                type = "select_com";
            } else if (
                item.field_key === "creator" ||
                item.field_key === "handler"
            ) {
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
.detail-content {
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
}
</style>
<style lang="scss">
.text.el-table__cell {
    .cell {
        padding: 0;
    }
}
</style>
