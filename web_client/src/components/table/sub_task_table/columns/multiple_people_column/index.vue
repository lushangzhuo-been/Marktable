<template>
    <el-table-column
        show-overflow-tooltip
        class-name="text"
        :min-width="widthMap[item.field_key] || 220"
        :prop="item.columnName"
    >
        <template slot="header">
            <div>
                <b class="type-box" :style="getType(item)"></b>
                {{ item.name }}
            </div>
        </template>
        <template slot-scope="scope">
            <people-content :item="item" :scope="scope"></people-content>
        </template>
    </el-table-column>
</template>

<script>
import { FieldType } from "@/assets/tool/const";
import { baseMixin } from "@/mixin.js";
import { imgHost } from "@/assets/tool/const";
import PeopleContent from "./content.vue";
export default {
    mixins: [baseMixin],
    components: {
        PeopleContent
    },
    props: {
        item: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            widthMap: {
                handler: 180,
                creator: 120
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
        },
        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        },
        getAvatarStack() {
            return {
                position: "relative",
                top: "-2px"
            };
        }
    }
};
</script>

<style lang="scss" scoped>
@import "@/components/table/p_table/style.scss";
</style>
<style lang="scss">
.text.el-table__cell {
    .cell {
        padding: 0;
    }
}
</style>
