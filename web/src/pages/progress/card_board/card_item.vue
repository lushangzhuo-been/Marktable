<template>
    <div @click="openDetail">
        <div
            v-for="(fieldItem, index) in fieldKeyConfig"
            :key="index"
            class="field-item-block"
            :class="{ title: fieldItem.field_key === 'title' }"
        >
            <!-- 标题 -->
            <text-type
                v-if="
                    fieldItem.mode === 'text_com' &&
                    fieldItem.field_key === 'title' &&
                    cardDetail[fieldItem.field_key]
                "
                class="item-temp"
                :fieldItem="fieldItem"
                :detailInfo="cardDetail[fieldItem.field_key]"
            ></text-type>
            <!-- 人 -->
            <people-type
                v-if="
                    fieldItem.mode === 'person_com' &&
                    cardDetail[fieldItem.field_key] &&
                    cardDetail[fieldItem.field_key].length
                "
                class="item-temp multiple"
                :fieldItem="fieldItem"
                :detailInfo="cardDetail[fieldItem.field_key]"
            >
            </people-type>
            <!-- 选 -->
            <select-type
                v-if="
                    (fieldItem.mode === 'multi_select_com' ||
                        fieldItem.mode === 'select_com') &&
                    cardDetail[fieldItem.field_key] &&
                    cardDetail[fieldItem.field_key].length
                "
                class="item-temp multiple"
                :fieldItem="fieldItem"
                :detailInfo="cardDetail[fieldItem.field_key]"
            ></select-type>
            <!-- 进度 -->
            <progress-type
                v-if="
                    fieldItem.mode === 'progress_com' &&
                    cardDetail[fieldItem.field_key]
                "
                class="item-temp progress"
                :fieldItem="fieldItem"
                :detailInfo="cardDetail[fieldItem.field_key]"
            ></progress-type>
            <!-- 编辑器 -->
            <b
                v-if="
                    fieldItem.mode === 'html_text_com' &&
                    cardDetail[fieldItem.field_key]
                "
                class="detail item-temp html-box"
            ></b>
            <!-- 链接 -->
            <link-type
                v-if="
                    fieldItem.mode === 'link_com' &&
                    cardDetail[fieldItem.field_key] &&
                    (cardDetail[fieldItem.field_key].name ||
                        cardDetail[fieldItem.field_key].url)
                "
                :fieldItem="fieldItem"
                :detailInfo="cardDetail[fieldItem.field_key]"
            ></link-type>
            <!-- 关联 -->
            <relation-type
                v-if="
                    fieldItem.mode === 'related_com' &&
                    cardDetail[fieldItem.field_key] &&
                    cardDetail[fieldItem.field_key].length
                "
                :fieldItem="fieldItem"
                :detailInfo="cardDetail[fieldItem.field_key]"
            >
            </relation-type>
            <!-- 其他 -->
            <default-type
                v-if="
                    fieldItem.mode === 'time_com' &&
                    fieldItem.mode === 'date_com' &&
                    fieldItem.mode === 'textarea_com' &&
                    fieldItem.mode === 'number_com' &&
                    cardDetail[fieldItem.field_key]
                "
                class="item-temp"
                :fieldItem="fieldItem"
                :detailInfo="cardDetail[fieldItem.field_key]"
            ></default-type>
        </div>
    </div>
</template>

<script>
import TextType from "@/pages/progress/card_board/text_type.vue";
import LinkType from "@/pages/progress/card_board/link_type.vue";
import PeopleType from "@/pages/progress/card_board/people_type.vue";
import SelectType from "@/pages/progress/card_board/select_type.vue";
import ProgressType from "@/pages/progress/card_board/progress_type.vue";
import DefaultType from "@/pages/progress/card_board/default_type.vue";
import RelationType from "@/pages/progress/card_board/relation_type.vue";
export default {
    components: {
        TextType,
        LinkType,
        PeopleType,
        SelectType,
        ProgressType,
        DefaultType,
        RelationType
    },
    props: {
        fieldKeyConfig: {
            type: Array,
            default: () => []
        },
        cardDetail: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {};
    },
    mounted() {},
    methods: {
        openDetail() {
            this.$emit("open-detail", this.cardDetail);
        }
    }
};
</script>

<style lang="scss" scoped>
.field-item-block {
    display: inline-block;
    max-width: 100%;
    margin-bottom: 8px;
    &.title {
        display: block;
        width: 100%;
        .item-temp {
            max-width: 100%;
        }
    }
    .item-temp {
        max-width: 80px;
        margin-right: 6px;
        &.progress {
            min-width: 120px;
            width: 120px;
            max-width: 120px;
        }
        &.multiple {
            max-width: 228px;
        }
    }
    &:last-child {
        .item-temp {
            margin-bottom: 0;
        }
    }
}
.html-box {
    margin-top: 2px;
    display: inline-block;
    width: 20px;
    height: 20px;
    background-image: url("@/assets/image/progress_column/col_html_content.png");
    background-size: 100% 100%;
}
</style>
