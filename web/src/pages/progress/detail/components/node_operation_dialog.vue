<template>
    <div>
        <el-dialog
            :title="btnDetail.name || ''"
            :visible.sync="dialogVisible"
            :close-on-click-modal="false"
            :close-on-press-escape="false"
            append-to-body
            class="basic-ui"
        >
            <!-- 节点进度form表单 -->
            <el-form
                id="form"
                ref="form"
                :model="formData"
                class="basic-ui add-progress-form"
                size="small"
                label-width="140px"
                :label-position="'left'"
            >
                <!-- :rules="rules" -->
                <div
                    v-for="(formItem, formIndex) in formLabel"
                    :key="formIndex"
                    class="circulate-content"
                >
                    <!-- 按照类型 分配组件 -->
                    <!-- 流程关联列 -->
                    <template v-if="formItem.mode === 'related_com'">
                        <el-form-item :label="formItem.name + ':'">
                            <div slot="label">
                                <b
                                    class="type-box"
                                    :class="{
                                        required: formItem.required === 'yes'
                                    }"
                                    :style="
                                        getType(formItem.mode, formItem.expr)
                                    "
                                ></b>
                                {{ formItem.name + ":" }}
                            </div>
                            <relation-progress-item
                                :formItem="formItem"
                                :formData="formData"
                                :validateOrder="validateOrder"
                                @edit-form-item="editFormItem"
                            ></relation-progress-item>
                        </el-form-item>
                    </template>
                    <!-- 文本列 -->
                    <template
                        v-if="
                            formItem.mode === 'mainText' ||
                            formItem.mode === 'text_com'
                        "
                    >
                        <el-form-item :label="formItem.name + ':'">
                            <div slot="label">
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
                                :formData="formData"
                                :validateOrder="validateOrder"
                                @edit-form-item="editFormItem"
                            ></text-item>
                        </el-form-item>
                    </template>
                    <!-- 数字 -->
                    <template v-if="formItem.mode === 'number_com'">
                        <el-form-item :label="formItem.name + ':'">
                            <div slot="label">
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
                                :formData="formData"
                                :validateOrder="validateOrder"
                                @edit-form-item="editFormItem"
                            ></number-item>
                        </el-form-item>
                    </template>
                    <!-- 进展 -->
                    <template v-if="formItem.mode === 'progress_com'">
                        <el-form-item :label="formItem.name + ':'">
                            <div slot="label">
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
                                :formData="formData"
                                @edit-form-item="editFormItem"
                            ></progress-item>
                        </el-form-item>
                    </template>
                    <!-- 下拉单选 -->
                    <template v-if="formItem.mode === 'select_com'">
                        <el-form-item :label="formItem.name + ':'">
                            <div slot="label">
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
                                :formItem="formItem"
                                :formData="formData"
                                :validateOrder="validateOrder"
                                @edit-form-item="editFormItem"
                            ></radio-select-item>
                        </el-form-item>
                    </template>
                    <!-- 下拉多选 -->
                    <template v-if="formItem.mode === 'multi_select_com'">
                        <el-form-item :label="formItem.name + ':'">
                            <div slot="label">
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
                                :formItem="formItem"
                                :formData="formData"
                                :validateOrder="validateOrder"
                                @edit-form-item="editFormItem"
                            ></multi-select-item>
                        </el-form-item>
                    </template>
                    <!-- html编辑 -->
                    <template v-if="formItem.mode === 'html_text_com'">
                        <el-form-item :label="formItem.name + ':'">
                            <div slot="label">
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
                                :formData="formData"
                                :validateOrder="validateOrder"
                                @edit-form-item="editFormItem"
                            ></html-text-item>
                        </el-form-item>
                    </template>
                    <!-- 单人列 -->
                    <template
                        v-if="
                            formItem.mode === 'person_com' &&
                            formItem.expr === 'single'
                        "
                    >
                        <el-form-item :label="formItem.name + ':'">
                            <div slot="label">
                                <b
                                    class="type-box"
                                    :class="{
                                        required: formItem.required === 'yes'
                                    }"
                                    :style="
                                        getType(formItem.mode, formItem.expr)
                                    "
                                ></b>
                                {{ formItem.name + ":" }}
                            </div>
                            <single-people-item
                                :formItem="formItem"
                                :formData="formData"
                                :validateOrder="validateOrder"
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
                        <el-form-item :label="formItem.name + ':'">
                            <div slot="label">
                                <b
                                    class="type-box"
                                    :class="{
                                        required: formItem.required === 'yes'
                                    }"
                                    :style="
                                        getType(formItem.mode, formItem.expr)
                                    "
                                ></b>
                                {{ formItem.name + ":" }}
                            </div>
                            <MultiplePeopleItem
                                :formItem="formItem"
                                :formData="formData"
                                :validateOrder="validateOrder"
                                @edit-form-item="editFormItem"
                            >
                            </MultiplePeopleItem>
                        </el-form-item>
                    </template>
                    <!-- 日期 -->
                    <template v-if="formItem.mode === 'date_com'">
                        <el-form-item :label="formItem.name + ':'">
                            <div slot="label">
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
                                :formData="formData"
                                :validateOrder="validateOrder"
                                @edit-form-item="editFormItem"
                            ></date-item>
                        </el-form-item>
                    </template>
                    <!-- 时间 -->
                    <template v-if="formItem.mode === 'time_com'">
                        <el-form-item :label="formItem.name + ':'">
                            <div slot="label">
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
                                :formData="formData"
                                :validateOrder="validateOrder"
                                @edit-form-item="editFormItem"
                            ></date-time-item>
                        </el-form-item>
                    </template>
                    <!-- 富文本 -->
                    <template v-if="formItem.mode === 'textarea_com'">
                        <el-form-item :label="formItem.name + ':'">
                            <div slot="label">
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
                                :formData="formData"
                                :validateOrder="validateOrder"
                                @edit-form-item="editFormItem"
                            ></rich-text-item>
                        </el-form-item>
                    </template>
                    <!-- 链接 -->
                    <template v-if="formItem.mode === 'link_com'">
                        <el-form-item :label="formItem.name + ':'">
                            <div slot="label">
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
                                :formData="formData"
                                :validateOrder="validateOrder"
                                @edit-form-item="editFormItem"
                            ></link-item>
                        </el-form-item>
                    </template>
                </div>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button class="basic-ui" size="small" @click="cancel"
                    >取消</el-button
                >
                <el-button
                    class="basic-ui"
                    size="small"
                    type="primary"
                    @click="onConfirm"
                >
                    确定
                </el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
import TextItem from "@/pages/progress/detail/components/node_operation_form/text_item/index.vue";
import NumberItem from "@/pages/progress/detail/components/node_operation_form/number_item/index.vue";
import ProgressItem from "@/pages/progress/detail/components/node_operation_form/progress_item/index.vue";
import MultiplePeopleItem from "@/pages/progress/detail/components/node_operation_form/multiple_people_item/index2.vue";
import SinglePeopleItem from "@/pages/progress/detail/components/node_operation_form/single_people_item/index2.vue";
import RelationProgressItem from "@/pages/progress/detail/components/node_operation_form/relation_progress_item/index2.vue";
import DateItem from "@/pages/progress/detail/components/node_operation_form/date_item/index.vue";
import DateTimeItem from "@/pages/progress/detail/components/node_operation_form/date_time_item/index.vue";
import RichTextItem from "@/pages/progress/detail/components/node_operation_form/rich_text_item/index.vue";
import LinkItem from "@/pages/progress/detail/components/node_operation_form/link_item/index.vue";
import { FieldType } from "@/assets/tool/const";
import MultiSelectItem from "@/pages/progress/detail/components/node_operation_form/multi_select_item/index2.vue";
import RadioSelectItem from "@/pages/progress/detail/components/node_operation_form/radio_select_item/index.vue";
import HtmlTextItem from "@/pages/progress/detail/components/node_operation_form/html_text_item/index.vue";

export default {
    components: {
        TextItem,
        NumberItem,
        ProgressItem,
        MultiplePeopleItem,
        SinglePeopleItem,
        RelationProgressItem,
        DateItem,
        DateTimeItem,
        RichTextItem,
        LinkItem,
        MultiSelectItem,
        RadioSelectItem,
        HtmlTextItem
    },
    props: {
        nodeOperationConfig: {
            type: Array,
            default: () => []
        },
        nodeOperationData: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            title: "确认bug(当前流程所处状态)",
            dialogVisible: false,
            form: {},
            formLabel: [],
            formData: {},
            btnDetail: {},
            changeForm: {},
            rules: {},
            validateOrder: 0 //校验指令
        };
    },
    watch: {
        nodeOperationConfig: {
            handler(arr) {
                this.formLabel = arr;
            }
        },
        nodeOperationData: {
            handler(obj) {
                // let modes = ['', '', '']
                // sub.mode.indexOf(modes) > -1  ?
                this.formData = obj;
                if (this.formData && this.formLabel) {
                    this.formLabel.forEach((sub) => {
                        for (let key in this.formData) {
                            if (key === sub.field_key) {
                                this.changeForm[sub.field_key] = JSON.stringify(
                                    this.formData[key]
                                );
                            }
                        }
                    });
                }
            }
        },
        dialogVisible: {
            handler(boolean) {
                if (!boolean) {
                    this.formLabel = [];
                }
            }
        }
    },
    methods: {
        setFormRule() {
            let rules = {};
            this.formLabel.forEach((formItem) => {
                if (formItem.required === "yes") {
                    let key = formItem.field_key;
                    rules[key] = [
                        {
                            required: true,
                            message: "此项为必填"
                        }
                    ];
                }
            });
            this.rules = rules;
        },
        closeDialog() {
            this.dialogVisible = false;
        },
        openDialog(btnDetail) {
            this.btnDetail = btnDetail;
            this.dialogVisible = true;
            this.$nextTick(() => {
                if (this.$refs.form) {
                    this.$refs.form.clearValidate();
                }
            });
        },
        cancel() {
            this.$refs.form.clearValidate();
            this.dialogVisible = false;
            this.$emit("cancel-submit", this.btnDetail);
        },
        onConfirm() {
            // 向下索取必填校验结果
            this.validateOrder += 1;
            let validateFailedNum = 0;
            this.$nextTick(() => {
                this.formLabel.forEach((item) => {
                    if (item.validateFailed) {
                        validateFailedNum += 1;
                    }
                });
                if (!validateFailedNum) {
                    this.$emit(
                        "confirm-next-state",
                        this.btnDetail,
                        this.changeForm
                    );
                }
            });
        },
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
        editFormItem(res, key, mode) {
            this.$set(this.changeForm, key, res);
            this.$emit("change-form", this.changeForm);
        }
    }
};
</script>

<style lang="scss" scoped>
.add-trigger-dialog {
    .el-dialog {
        width: 520px;
        border-radius: 4px;
        .el-dialog__header {
            padding: 24px 32px;
            font-size: 16px;
            color: #2f384c;
        }
        .el-dialog__body {
            padding: 0 32px;
        }
        .el-dialog__footer {
            padding: 0 32px 32px;
        }
        .el-dialog__headerbtn {
            top: 24px;
            right: 32px;
        }
        .el-form-item {
            margin-bottom: 24px;
        }
        .el-dialog__footer {
            text-align: right;
        }
    }
}
.basic-ui.add-progress-form {
    margin: 6px 0;
}
.type-box {
    display: inline-block;
    width: 20px;
    height: 20px;
    background-size: 100% 100%;
    vertical-align: middle;
    position: relative;
    top: -2px;
    &.required {
        position: relative;
        &::before {
            content: "*";
            color: #f56c6c;
            position: absolute;
            top: -8px;
            left: -8px;
        }
    }
}
</style>
