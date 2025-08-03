<template>
    <!-- 详情及单路由 共用form部分抽取 -->
    <div>
        <el-form
            class="basic-ui progress-detail-form"
            ref="form"
            :model="formDataShow"
            :label-position="'left'"
            label-width="140px"
            size="small"
        >
            <div
                v-for="(formItem, formIndex) in filterDisplayField(
                    formLabelShow
                )"
                :key="formIndex"
                class="circulate-content"
            >
                <!-- 按照类型 分配组件 -->
                <!-- 流程关联列 -->
                <template v-if="formItem.mode === 'related_com'">
                    <el-form-item>
                        <div
                            class="form-item-label"
                            :class="{
                                required: formItem.required === 'yes'
                            }"
                            slot="label"
                        >
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode, formItem.expr)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </div>
                        <relation-progress-item
                            :formItem="formItem"
                            :formData="formDataShow"
                            @edit-form-item="editFormItem"
                        ></relation-progress-item>
                    </el-form-item>
                </template>
                <!-- 文本列 -->
                <template v-if="formItem.mode === 'text_com'">
                    <el-form-item>
                        <div
                            class="form-item-label"
                            :class="{
                                required: formItem.required === 'yes'
                            }"
                            slot="label"
                        >
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </div>
                        <text-item
                            :formItem="formItem"
                            :formData="formDataShow"
                            @edit-form-item="editFormItem"
                        ></text-item>
                    </el-form-item>
                </template>
                <!-- 数字 -->
                <template v-if="formItem.mode === 'number_com'">
                    <el-form-item>
                        <div
                            class="form-item-label"
                            :class="{
                                required: formItem.required === 'yes'
                            }"
                            slot="label"
                        >
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name }}
                            <span v-show="formItem.expr"
                                >({{ formItem.expr }})</span
                            >:
                        </div>
                        <number-item
                            :formItem="formItem"
                            :formData="formDataShow"
                            @edit-form-item="editFormItem"
                        ></number-item>
                    </el-form-item>
                </template>
                <!-- 单选item -->
                <template v-if="formItem.mode === 'select_com'">
                    <el-form-item>
                        <div
                            class="form-item-label"
                            :class="{
                                required: formItem.required === 'yes'
                            }"
                            slot="label"
                        >
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </div>
                        <radio-select-item
                            :item="formItem"
                            :formData="formDataShow"
                            @edit-form-item="editFormItem"
                        ></radio-select-item>
                    </el-form-item>
                </template>
                <!-- 多选item -->
                <template v-if="formItem.mode === 'multi_select_com'">
                    <el-form-item>
                        <div
                            class="form-item-label"
                            :class="{
                                required: formItem.required === 'yes'
                            }"
                            slot="label"
                        >
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </div>
                        <multi-select-item
                            :item="formItem"
                            :formData="formDataShow"
                            @edit-form-item="editFormItem"
                        ></multi-select-item>
                    </el-form-item>
                </template>
                <!-- 单人列 -->
                <template
                    v-if="
                        formItem.mode === 'person_com' &&
                        formItem.expr === 'single'
                    "
                >
                    <el-form-item>
                        <div
                            class="form-item-label"
                            :class="{
                                required: formItem.required === 'yes'
                            }"
                            slot="label"
                        >
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode, formItem.expr)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </div>
                        <single-people-item
                            class="detail-content-person"
                            :formItem="formItem"
                            :formData="formDataShow"
                            @edit-form-item="editFormItem"
                        ></single-people-item>
                    </el-form-item>
                </template>
                <!-- 多人项 -->
                <template
                    v-if="
                        formItem.mode === 'person_com' &&
                        formItem.expr === 'multiple'
                    "
                >
                    <el-form-item>
                        <div
                            class="form-item-label"
                            :class="{
                                required: formItem.required === 'yes'
                            }"
                            slot="label"
                        >
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode, formItem.expr)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </div>
                        <multiple-people-item
                            class="detail-content-person"
                            :formItem="formItem"
                            :formData="formDataShow"
                            @edit-form-item="editFormItem"
                        ></multiple-people-item>
                    </el-form-item>
                </template>
                <!-- 日期 年月日 -->
                <template v-if="formItem.mode === 'date_com'">
                    <el-form-item>
                        <div
                            class="form-item-label"
                            :class="{
                                required: formItem.required === 'yes'
                            }"
                            slot="label"
                        >
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </div>
                        <date-item
                            :formItem="formItem"
                            :formData="formDataShow"
                            @edit-form-item="editFormItem"
                        ></date-item>
                    </el-form-item>
                </template>
                <!-- 时间 年月日 时分秒-->
                <template v-if="formItem.mode === 'time_com'">
                    <el-form-item>
                        <div
                            class="form-item-label"
                            :class="{
                                required: formItem.required === 'yes'
                            }"
                            slot="label"
                        >
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </div>
                        <date-time-item
                            :formItem="formItem"
                            :formData="formDataShow"
                            @edit-form-item="editFormItem"
                        ></date-time-item>
                    </el-form-item>
                </template>
                <!-- 富文本 -->
                <template v-if="formItem.mode === 'textarea_com'">
                    <el-form-item>
                        <div
                            class="form-item-label"
                            :class="{
                                required: formItem.required === 'yes'
                            }"
                            slot="label"
                        >
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </div>
                        <rich-text-item
                            :formItem="formItem"
                            :formData="formDataShow"
                            @edit-form-item="editFormItem"
                        ></rich-text-item>
                    </el-form-item>
                </template>
                <!-- HTML编辑框 -->
                <template v-if="formItem.mode === 'html_text_com'">
                    <el-form-item>
                        <div
                            class="form-item-label"
                            :class="{
                                required: formItem.required === 'yes'
                            }"
                            slot="label"
                        >
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </div>
                        <html-text-item
                            :formItem="formItem"
                            :formData="formDataShow"
                            @edit-form-item="editFormItem"
                        ></html-text-item>
                    </el-form-item>
                </template>
                <!-- 链接 -->
                <template v-if="formItem.mode === 'link_com'">
                    <el-form-item>
                        <div
                            class="form-item-label"
                            :class="{
                                required: formItem.required === 'yes'
                            }"
                            slot="label"
                        >
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </div>
                        <link-item
                            :formItem="formItem"
                            :formData="formDataShow"
                            @edit-form-item="editFormItem"
                        ></link-item>
                    </el-form-item>
                </template>
                <!-- 进展 -->
                <template v-if="formItem.mode === 'progress_com'">
                    <el-form-item>
                        <div
                            class="form-item-label"
                            :class="{
                                required: formItem.required === 'yes'
                            }"
                            slot="label"
                        >
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </div>
                        <progress-item
                            :formItem="formItem"
                            :formData="formDataShow"
                            @edit-form-item="editFormItem"
                        ></progress-item>
                    </el-form-item>
                </template>
            </div>
        </el-form>
    </div>
</template>

<script>
import _ from "lodash";
import { FieldType } from "@/assets/tool/const";
import TextItem from "./form_item/text_item/index.vue";
import NumberItem from "./form_item/number_item/index.vue";
import MultiSelectItem from "./form_item/multi_select_item/index2.vue";
import RadioSelectItem from "./form_item/radio_select_item/index2.vue";
import HtmlTextItem from "./form_item/html_text_item/index.vue";
import SinglePeopleItem from "./form_item/single_people_item/index2.vue";
import RelationProgressItem from "./form_item/relation_progress_item/index2.vue"; // 流程关联列
import MultiplePeopleItem from "./form_item/multiple_people_item/index2.vue";
import DateItem from "./form_item/date_item/index.vue";
import DateTimeItem from "./form_item/date_time_item/index.vue";
import RichTextItem from "./form_item/rich_text_item/index.vue";
import LinkItem from "./form_item/link_item/index.vue";
import ProgressItem from "./form_item/progress_item/index.vue";
export default {
    components: {
        TextItem,
        NumberItem,
        MultiSelectItem,
        RadioSelectItem,
        HtmlTextItem,
        RelationProgressItem,
        SinglePeopleItem,
        MultiplePeopleItem,
        DateItem,
        DateTimeItem,
        RichTextItem,
        LinkItem,
        ProgressItem
    },
    props: {
        formLabelShow: {
            type: Array,
            default: () => []
        },
        formDataShow: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {};
    },
    computed: {
        filterDisplayField() {
            // 过滤不展示标题、状态、处理人、创建人、创建时间、最后更新时间
            return (arr) => {
                if (arr && arr.length) {
                    let filterArr = _.cloneDeep(arr).filter((item) => {
                        return (
                            item.field_key !== "title" &&
                            item.field_key !== "status" &&
                            item.field_key !== "handler" &&
                            item.field_key !== "creator" &&
                            item.field_key !== "created_at" &&
                            item.field_key !== "updated_at"
                        );
                    });
                    return filterArr;
                } else {
                    return [];
                }
            };
        }
    },
    methods: {
        getType(type, expr) {
            if (type && !expr) {
                return {
                    "background-image": `url(${FieldType[type]})`
                };
            } else {
                return {
                    "background-image": `url(${FieldType[type + "_" + expr]})`
                };
            }
        },
        // 提交编辑
        editFormItem(res, key, id, mode) {
            this.$emit("edit-form-item", res, key, id, mode);
        }
    }
};
</script>

<style lang="scss" scoped>
@import "@/components/table/p_table/style.scss";
@import "@/pages/progress/detail/draw_and_single_route.scss";
</style>
<style lang="scss">
@import "@/pages/progress/detail/draw_and_single_route_drill.scss"; // @import "@/components/table/p_table/drill_style.scss";
</style>
