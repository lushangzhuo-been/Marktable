<template>
    <el-dialog
        title="新增"
        :visible.sync="dialogVisible"
        :close-on-click-modal="false"
        class="basic-ui add-progress-dialog"
        append-to-body
        @close="cancel"
    >
        <el-form
            id="form"
            ref="form"
            :model="formData"
            class="basic-ui add-progress-form"
            size="small"
            label-width="140px"
            :label-position="'left'"
            :validate-on-rule-change="false"
        >
            <div
                v-for="(formItem, formIndex) in formLabelShow"
                :key="formIndex"
                class="circulate-content"
            >
                <!-- 按照类型 分配组件 -->
                <!-- 关联流程列 -->
                <template v-if="formItem.mode === 'related_com'">
                    <el-form-item
                        :label="formItem.name + ':'"
                        :prop="formItem.field_key"
                    >
                        <span slot="label">
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode, formItem.expr)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </span>
                        <relation-progress-item
                            :formItem="formItem"
                            :formData="formData"
                            :validateOrder="validateOrder"
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
                    <el-form-item
                        :label="formItem.name + ':'"
                        :prop="formItem.field_key"
                    >
                        <span slot="label">
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </span>
                        <text-item
                            :formItem="formItem"
                            :formData="formData"
                            :validateOrder="validateOrder"
                        ></text-item>
                    </el-form-item>
                </template>
                <!-- 数字 -->
                <template v-if="formItem.mode === 'number_com'">
                    <el-form-item
                        :label="formItem.name + ':'"
                        :prop="formItem.field_key"
                    >
                        <span slot="label">
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
                        </span>
                        <number-item
                            :formItem="formItem"
                            :formData="formData"
                            :validateOrder="validateOrder"
                        ></number-item>
                    </el-form-item>
                </template>
                <!-- 进展 -->
                <template v-if="formItem.mode === 'progress_com'">
                    <el-form-item
                        :label="formItem.name + ':'"
                        :prop="formItem.field_key"
                    >
                        <span slot="label">
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </span>
                        <progress-item
                            :formItem="formItem"
                            :formData="formData"
                            :validateOrder="validateOrder"
                        ></progress-item>
                    </el-form-item>
                </template>
                <!-- 下拉列单选 -->
                <template v-if="formItem.mode === 'select_com'">
                    <el-form-item
                        :label="formItem.name + ':'"
                        :prop="formItem.field_key"
                    >
                        <span slot="label">
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </span>
                        <radio-select-item
                            :formItem="formItem"
                            :formData="formData"
                            :validateOrder="validateOrder"
                            @refresh="refresh"
                            @on-change="formChange"
                        ></radio-select-item>
                    </el-form-item>
                </template>
                <!-- 单人列 -->
                <template
                    v-if="
                        formItem.mode === 'person_com' &&
                        formItem.expr === 'single'
                    "
                >
                    <el-form-item
                        :label="formItem.name + ':'"
                        :prop="formItem.field_key"
                    >
                        <span slot="label">
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode, formItem.expr)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </span>
                        <single-people-item
                            :formItem="formItem"
                            :formData="formData"
                            :validateOrder="validateOrder"
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
                    <el-form-item
                        :label="formItem.name + ':'"
                        :prop="formItem.field_key"
                    >
                        <span slot="label">
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode, formItem.expr)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </span>
                        <multiple-people-item
                            :formItem="formItem"
                            :formData="formData"
                            :validateOrder="validateOrder"
                        ></multiple-people-item>
                    </el-form-item>
                </template>
                <!-- 日期 -->
                <template v-if="formItem.mode === 'date_com'">
                    <el-form-item
                        :label="formItem.name + ':'"
                        :prop="formItem.field_key"
                    >
                        <span slot="label">
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </span>
                        <date-item
                            :formItem="formItem"
                            :formData="formData"
                            :validateOrder="validateOrder"
                        ></date-item>
                    </el-form-item>
                </template>
                <!-- 时间 -->
                <template v-if="formItem.mode === 'time_com'">
                    <el-form-item
                        :label="formItem.name + ':'"
                        :prop="formItem.field_key"
                    >
                        <span slot="label">
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </span>
                        <date-time-item
                            :formItem="formItem"
                            :formData="formData"
                            :validateOrder="validateOrder"
                        ></date-time-item>
                    </el-form-item>
                </template>
                <!-- 多选 -->
                <template v-if="formItem.mode === 'multi_select_com'">
                    <el-form-item
                        :label="formItem.name + ':'"
                        :prop="formItem.field_key"
                    >
                        <span slot="label">
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </span>
                        <multi-select-item
                            :formItem="formItem"
                            :formData="formData"
                            :validateOrder="validateOrder"
                            @on-change="formChange"
                        ></multi-select-item>
                    </el-form-item>
                </template>
                <!-- 富文本 -->
                <template v-if="formItem.mode === 'textarea_com'">
                    <el-form-item
                        :label="formItem.name + ':'"
                        :prop="formItem.field_key"
                    >
                        <span slot="label">
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </span>
                        <rich-text-item
                            :formItem="formItem"
                            :formData="formData"
                            :validateOrder="validateOrder"
                        ></rich-text-item>
                    </el-form-item>
                </template>
                <!-- html富文本 -->
                <template v-if="formItem.mode === 'html_text_com'">
                    <el-form-item
                        :label="formItem.name + ':'"
                        :prop="formItem.field_key"
                    >
                        <span slot="label">
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </span>
                        <html-text-item
                            :formItem="formItem"
                            :formData="formData"
                            :validateOrder="validateOrder"
                            @on-change="formChange"
                        ></html-text-item>
                    </el-form-item>
                </template>
                <!-- 链接 -->
                <template v-if="formItem.mode === 'link_com'">
                    <el-form-item
                        :label="formItem.name + ':'"
                        :prop="formItem.field_key"
                    >
                        <span slot="label">
                            <b
                                class="type-box"
                                :class="{
                                    required: formItem.required === 'yes'
                                }"
                                :style="getType(formItem.mode)"
                            ></b>
                            {{ formItem.name + ":" }}
                        </span>
                        <link-item
                            :formItem="formItem"
                            :formData="formData"
                            :validateOrder="validateOrder"
                        ></link-item>
                    </el-form-item>
                </template>
            </div>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <div :class="{ 'shadow-block': dialogFooterShadow }"></div>
            <el-button size="small" @click="cancel">取 消</el-button>
            <el-button size="small" type="primary" @click="onConfirm"
                >确 定</el-button
            >
        </div>
    </el-dialog>
</template>

<script>
import _ from "lodash";
import { DefauleAvatar } from "@/assets/tool/const";
import TextItem from "./form_item/text_item/index.vue";
import NumberItem from "./form_item/number_item/index.vue";
import ProgressItem from "./form_item/progress_item/index.vue";

import RelationProgressItem from "./form_item/relation_progress_item/index.vue";
import SinglePeopleItem from "./form_item/single_people_item/index.vue";
import MultiplePeopleItem from "./form_item/multiple_people_item/index2.vue";
import DateItem from "./form_item/date_item/index.vue";
import DateTimeItem from "./form_item/date_time_item/index.vue";
import RichTextItem from "./form_item/rich_text_item/index.vue";
import HtmlTextItem from "./form_item/html_text_item/index.vue";
import MultiSelectItem from "./form_item/multi_select_item/index2.vue";
import RadioSelectItem from "./form_item/radio_select_item/index.vue";
import LinkItem from "./form_item/link_item/index.vue";
import { FieldType } from "@/assets/tool/const";
import api from "@/common/api/module/progress";
export default {
    components: {
        TextItem,
        NumberItem,
        ProgressItem,
        RelationProgressItem,
        SinglePeopleItem,
        MultiplePeopleItem,
        DateItem,
        DateTimeItem,
        RichTextItem,
        HtmlTextItem,
        RadioSelectItem,
        MultiSelectItem,
        LinkItem
    },
    props: {
        formLabel: {
            type: Array,
            default: () => []
        },
        from: {
            type: String,
            default: ""
        },
        subTaskInfo: {
            type: Object,
            default: () => {}
        },
        detailId: {
            type: String,
            default: ""
        }
        // formData: {
        //     type: Object,
        //     default: () => {},
        // },
    },
    data() {
        return {
            dialogVisible: false,
            formData: {},
            dialogFooterShadow: false,
            rules: {},
            formLabelShow: {},
            validateOrder: 0 //校验指令
        };
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        }
    },
    watch: {
        formLabel: {
            handler(arr) {
                // 初始化数据格式
                this.getFormData(arr);
                // 遍历生成表单规则
                // this.setFormRule();
            }
            // immediate: true,
        },
        dialogVisible: {
            handler(boolean) {
                if (!boolean) {
                    this.formLabelShow = [];
                }
            }
        }
    },
    methods: {
        formChange(a, b, c, d) {
            this.$set(this.formData, `${b}`, a);
        },
        setFormRule() {
            let rules = {};
            this.formLabelShow.forEach((formItem) => {
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
        getFormData(arr) {
            // 添加流程的col,按照类型生成formData,遍历数组，取prop做key,取type判断类型给数据类型
            this.formData = {};
            arr.forEach((item) => {
                if (
                    item.mode === "mainText" ||
                    item.mode === "text_com" ||
                    item.mode === "textarea_com" ||
                    item.mode === "number_com" ||
                    item.mode === "html_text_com"
                ) {
                    // 主文本、文本、富文本
                    this.$set(this.formData, item.field_key, "");
                } else if (item.mode === "select_com") {
                    // 单选
                    let arr = [];
                    this.$set(this.formData, item.field_key, arr);
                } else if (
                    (item.mode === "person_com" && item.expr === "single") ||
                    item.mode === "related_com"
                ) {
                    // 单人
                    // let obj = {
                    //     id: "",
                    //     avatar: "",
                    // };
                    // this.$set(this.formData, item.field_key, obj);
                    let arr = [];
                    this.$set(this.formData, item.field_key, arr);
                } else if (
                    item.mode === "person_com" &&
                    item.expr === "multiple"
                ) {
                    // 多人
                    let arr = [];
                    this.$set(this.formData, item.field_key, arr);
                } else if (item.mode === "time_com") {
                    // 日期与时间
                    this.$set(this.formData, item.field_key, "");
                } else if (item.mode === "multi_select_com") {
                    // 多选
                    let arr = [];
                    this.$set(this.formData, item.field_key, arr);
                } else if (item.mode === "file") {
                    // 文件
                    let arr = [];
                    this.$set(this.formData, item.field_key, arr);
                } else if (item.mode === "link_com") {
                    // 链接
                    let obj = {
                        name: "",
                        url: ""
                    };
                    this.$set(this.formData, item.field_key, obj);
                }
            });
            this.formLabelShow = arr;
            this.$nextTick(() => {
                // 获取form实际高度 如果大于480，设计阴影样式
                if (document.getElementById("form")) {
                    let formHeight =
                        document.getElementById("form").clientHeight;
                    if (formHeight > 480) {
                        this.dialogFooterShadow = true;
                    } else {
                        this.dialogFooterShadow = false;
                    }
                }
            });
        },
        openDialog() {
            this.dialogVisible = true;
        },
        cancel() {
            this.dialogVisible = false;
        },
        onConfirm() {
            // 提交处理参数，将数组等提取唯一值后转为json
            let keysArr = Object.keys(this.formData);
            // 需处理 单选， 多选， 单人，多人
            let apiForm = {};
            keysArr.forEach((formItenKey) => {
                // 查询类型
                let type = _.find(this.formLabelShow, {
                    field_key: formItenKey
                }).mode;
                if (type === "select_com") {
                    // 单选
                    apiForm[formItenKey] = this.getItemKeyArr(
                        this.formData[formItenKey]
                    );
                } else if (type === "multi_select_com") {
                    apiForm[formItenKey] = this.getItemKeyArr(
                        this.formData[formItenKey]
                    );
                } else if (type === "link_com") {
                    apiForm[formItenKey] = JSON.stringify(
                        this.formData[formItenKey]
                    );
                } else if (type === "person_com") {
                    let expr = _.find(this.formLabelShow, {
                        field_key: formItenKey
                    }).expr;
                    if (expr === "single") {
                        // 单人
                        apiForm[formItenKey] = this.getPersonIdArr(
                            this.formData[formItenKey]
                        );
                    } else {
                        // 多人
                        apiForm[formItenKey] = this.getPersonIdArr(
                            this.formData[formItenKey]
                        );
                    }
                } else if (type === "related_com") {
                    apiForm[formItenKey] = this.getPersonIdArr(
                        this.formData[formItenKey],
                        "related_com"
                    );
                } else {
                    // 其他皆为字符串类型
                    apiForm[formItenKey] = this.formData[formItenKey];
                }
            });
            // 向下索取必填校验结果
            this.validateOrder += 1;
            let validateFailedNum = 0;
            this.$nextTick(() => {
                this.formLabelShow.forEach((item) => {
                    if (item.validateFailed) {
                        validateFailedNum += 1;
                    }
                });
                if (!validateFailedNum) {
                    if (this.from === "subTask") {
                        // 子任务
                        let params = {
                            parent_tmpl_id: this.subTaskInfo.tmpl_id,
                            obj_id: this.detailId,
                            ws_id: this.subTaskInfo.ws_id,
                            tmpl_id: this.subTaskInfo.sub_tmpl_id,
                            ...apiForm
                        };
                        api.createSubAction(params).then((res) => {
                            // 成功后关闭dialog，刷新列表
                            if (res && res.resultCode === 200) {
                                this.dialogVisible = false;
                                this.$emit("refresh");
                            }
                        });
                    } else {
                        // 流程
                        let params = {
                            ws_id: this.curSpace.id,
                            tmpl_id: this.curProgress,
                            ...apiForm
                        };
                        api.createProgressItem(params).then((res) => {
                            // 成功后关闭dialog，刷新列表
                            if (res && res.resultCode === 200) {
                                this.dialogVisible = false;
                                this.$emit("refresh");
                            }
                        });
                    }
                }
            });
        },
        getItemKeyObj(arr) {
            // 返回唯一值为对象的方法
            let apiArr = [];
            // if (Object.keys(obj).length) {
            //     apiArr.push(obj.value);
            //     return JSON.stringify(apiArr);
            // } else {
            //     return null;
            // }
            if (arr.length) {
                apiArr.push(arr[0]);
            }
            return JSON.stringify(apiArr);
        },
        getItemKeyArr(arr) {
            // 返回唯一值为数组的方法
            let apiArr = [];
            if (arr.length) {
                arr.forEach((arrItem) => {
                    apiArr.push(arrItem);
                });
            }
            return JSON.stringify(apiArr);
        },
        getPersonIdObj(obj) {
            let apiArr = [];
            if (Object.keys(obj).length) {
                apiArr.push(obj.id);
            }
            return JSON.stringify(apiArr);
        },
        getPersonIdArr(arr, type) {
            let apiArr = [];
            if (arr.length) {
                arr.forEach((arrItem) => {
                    apiArr.push(type === "related_com" ? arrItem : arrItem.id);
                });
            }
            return JSON.stringify(apiArr);
        },
        getType(type, expr) {
            if (type && !expr) {
                return {
                    "background-image": `url(${FieldType[type]})`
                };
            }
            if (type && expr) {
                return {
                    "background-image": `url(${FieldType[type + "_" + expr]})`
                };
            }
        },
        refresh(key) {
            let arr = [];
            arr.push(key);
            this.$refs.form.validateField(arr);
        }
    }
};
</script>

<style lang="scss" scoped>
@import "@/components/table/p_table/drill_style.scss";
.type-box {
    display: inline-block;
    width: 20px;
    height: 20px;
    background-size: 100% 100%;
    vertical-align: middle;
    position: relative;
    top: -2px;
    &.required {
        &::before {
            content: "*";
            color: #f56c6c;
            position: relative;
            top: -8px;
            left: -8px;
        }
    }
}
.basic-ui.add-progress-form {
    margin: 6px 0;
}
::v-deep.basic-ui.add-progress-dialog {
    .el-dialog {
        .el-dialog__footer {
            padding: 0;
            .dialog-footer {
                box-sizing: border-box;
                position: relative;
                padding: 12px 32px 16px;
                .shadow-block {
                    height: 1px;
                    box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.1);
                    position: absolute;
                    top: 0;
                    left: 0;
                    width: 100%;
                }
            }
        }
    }
}
</style>
