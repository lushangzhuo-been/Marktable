<template>
    <el-dialog
        title="规则执行日志"
        :visible.sync="dialogVisible"
        :width="'860px'"
        append-to-body
        :destroy-on-close="true"
        class="log-detail-dialog basic-ui"
    >
        <div>
            <div class="trigger-block">
                <div class="title-text">触发条件</div>
                <div class="desc-block trigger">
                    <!-- 动作描述 （创建/删除） -->
                    <div class="desc-text title">
                        当&nbsp;&nbsp;任务&nbsp;
                        <span
                            v-if="
                                triggerDetail.triggerType === 'create' ||
                                triggerDetail.triggerType === 'delete'
                            "
                        >
                            {{
                                triggerDetail.triggerType === "create"
                                    ? "创建时"
                                    : "删除时"
                            }}
                        </span>
                    </div>
                    <!-- 任务名 -->
                    <div class="normal-text task-name">
                        <div class="desc-text-block">
                            {{ triggerDetail.taskName }}
                        </div>
                    </div>
                    <!-- 前后变化 -->
                    <div
                        class="field-change"
                        v-if="triggerDetail.triggerType === 'update'"
                    >
                        <div class="desc-text title">字段</div>
                        <div class="normal-text field-change-text">
                            <div class="desc-text-block">
                                {{ getFieldKeyName(triggerDetail.fieldKey) }}
                            </div>
                            <!-- 人类型 -->
                            <template
                                v-if="
                                    getFieldKeyType(triggerDetail.fieldKey) ===
                                    'person_com'
                                "
                            >
                                <div class="desc-text">
                                    &nbsp;&nbsp;从&nbsp;&nbsp;
                                </div>
                                <!-- 人 -->
                                <user-show
                                    class="user-show"
                                    :userIdList="triggerDetail.oldValue"
                                    :userInfoList="triggerDetail.oldValueUser"
                                ></user-show>
                                <div class="desc-text-block desc-text">
                                    &nbsp;&nbsp;变更为&nbsp;&nbsp;
                                </div>
                                <user-show
                                    class="user-show"
                                    :userIdList="triggerDetail.newValue"
                                    :userInfoList="triggerDetail.newValueUser"
                                ></user-show>
                            </template>
                            <!-- 选择类型 -->
                            <template
                                v-else-if="
                                    getFieldKeyType(triggerDetail.fieldKey) ===
                                        'select_com' ||
                                    getFieldKeyType(triggerDetail.fieldKey) ===
                                        'multi_select_com'
                                "
                            >
                                <div class="desc-text">
                                    &nbsp;&nbsp;从&nbsp;&nbsp;
                                </div>

                                <select-show
                                    class="select-show"
                                    :option="
                                        getFieldKeyOption(
                                            triggerDetail.fieldKey
                                        )
                                    "
                                    :selectValue="triggerDetail.oldValue"
                                ></select-show>
                                <div class="desc-text-block desc-text">
                                    &nbsp;&nbsp;变更为&nbsp;&nbsp;
                                </div>
                                <select-show
                                    class="select-show"
                                    :option="
                                        getFieldKeyOption(
                                            triggerDetail.fieldKey
                                        )
                                    "
                                    :selectValue="triggerDetail.newValue"
                                ></select-show>
                            </template>
                            <!-- 状态类型 -->
                            <template
                                v-else-if="
                                    getFieldKeyType(triggerDetail.fieldKey) ===
                                    'status_com'
                                "
                            >
                                <div class="desc-text">
                                    &nbsp;&nbsp;从&nbsp;&nbsp;
                                </div>

                                <status-show
                                    class="select-show"
                                    :option="
                                        getFieldKeyOption(
                                            triggerDetail.fieldKey
                                        )
                                    "
                                    :selectValue="triggerDetail.oldValue"
                                ></status-show>
                                <div class="desc-text-block desc-text">
                                    &nbsp;&nbsp;变更为&nbsp;&nbsp;
                                </div>
                                <status-show
                                    class="select-show"
                                    :option="
                                        getFieldKeyOption(
                                            triggerDetail.fieldKey
                                        )
                                    "
                                    :selectValue="triggerDetail.newValue"
                                ></status-show>
                            </template>
                            <!-- 日期 时间 文本 数值 进展 -》 “变更时”替代 -->
                            <div
                                class="text-change-show"
                                v-else-if="
                                    getFieldKeyType(triggerDetail.fieldKey) ===
                                        'number_com' ||
                                    getFieldKeyType(triggerDetail.fieldKey) ===
                                        'text_com' ||
                                    getFieldKeyType(triggerDetail.fieldKey) ===
                                        'textarea_com' ||
                                    getFieldKeyType(triggerDetail.fieldKey) ===
                                        'time_com' ||
                                    getFieldKeyType(triggerDetail.fieldKey) ===
                                        'date_com' ||
                                    getFieldKeyType(triggerDetail.fieldKey) ===
                                        'progress_com'
                                "
                            >
                                <div class="desc-text">&nbsp;&nbsp;变更时</div>
                            </div>
                            <!-- 关系类型 -->
                            <template
                                v-else-if="
                                    getFieldKeyType(triggerDetail.fieldKey) ===
                                    'related_com'
                                "
                            >
                                <div class="desc-text">
                                    &nbsp;&nbsp;从&nbsp;&nbsp;
                                </div>
                                <!-- <RelatedShow -->
                                <related-show
                                    class="select-show"
                                    :selectValue="triggerDetail.oldValue"
                                    :relationReshowInfo="
                                        triggerDetail.oldValueDocuments
                                    "
                                ></related-show>
                                <div class="desc-text-block desc-text">
                                    &nbsp;&nbsp;变更为&nbsp;&nbsp;
                                </div>
                                <related-show
                                    class="select-show"
                                    :selectValue="triggerDetail.newValue"
                                    :relationReshowInfo="
                                        triggerDetail.newValueDocuments
                                    "
                                ></related-show>
                                <!-- 关系字段 -->
                            </template>
                            <!-- 其他 -->
                            <div class="text-change-show" v-else>
                                <div class="desc-text">
                                    &nbsp;&nbsp;从&nbsp;&nbsp;
                                </div>
                                {{ triggerDetail.oldValue }}
                                <div class="desc-text">变更为</div>
                                {{ triggerDetail.newValue }}
                            </div>
                        </div>
                    </div>
                    <!-- 定时任务 -->
                    <div v-if="triggerDetail.triggerType === 'schedule'">
                        <!-- 定时任务 -->
                        <div class="desc-text desc-text-field">字段</div>
                        <div class="normal-text schedule-block">
                            {{
                                getFieldKeyName(triggerDetail.fieldKey)
                            }}&nbsp;&nbsp; 到达&nbsp;&nbsp;<span
                                class="desc-text"
                                >{{
                                    triggerDetail.triggerDay === 0
                                        ? "当天"
                                        : triggerDetail.triggerDay < 0
                                        ? `在${-triggerDetail.triggerDay}天前的`
                                        : `在${triggerDetail.triggerDay}天后的`
                                }}{{ triggerDetail.triggerTime }}</span
                            >
                        </div>
                    </div>
                    <!-- filter 满足条件 -->
                    <div class="filter-list" v-if="formatFilterList.length">
                        <!-- {{ filterList }} -->
                        <div class="desc-text title">满足条件</div>
                        <div
                            v-for="(
                                filterItem, filterIndex
                            ) in formatFilterList"
                            :key="filterIndex"
                            class="filter-item"
                            :class="{ first: !filterIndex }"
                        >
                            <div v-if="filterIndex" class="filter-head-and">
                                且
                            </div>
                            <div class="filter-body">
                                <!-- 字段 -->
                                <div class="filter-field-key" v-overflow>
                                    {{ getFieldKeyName(filterItem.field_key) }}
                                </div>
                                <!-- 关系 -->
                                <div class="filter-relation" v-overflow>
                                    {{
                                        getFieldKeyOp(
                                            filterItem.field_key,
                                            filterItem.op
                                        )
                                    }}
                                </div>
                                <!-- 人 回显头像 -->
                                <div
                                    v-if="
                                        getFieldKeyType(
                                            filterItem.field_key
                                        ) === 'person_com'
                                    "
                                    class="filter-content"
                                >
                                    <!-- 人类型--{{ filterItem.value }} -->
                                    <user-message
                                        :userMessage="filterItem.value[0]"
                                    >
                                        <span class="member-item">
                                            <el-avatar
                                                class="progress-avatar"
                                                icon="el-icon-user-solid"
                                                size="small"
                                                :src="
                                                    getAvatar(
                                                        filterItem.value[0]
                                                            .avatar
                                                    )
                                                "
                                                :style="getAvatarStack()"
                                            ></el-avatar>
                                            {{ filterItem.value[0].full_name }}
                                        </span>
                                    </user-message>
                                </div>
                                <!-- 文本 富文本 日期 时间  -->
                                <div
                                    v-if="
                                        getFieldKeyType(
                                            filterItem.field_key
                                        ) === 'text_com' ||
                                        getFieldKeyType(
                                            filterItem.field_key
                                        ) === 'textarea_com' ||
                                        getFieldKeyType(
                                            filterItem.field_key
                                        ) === 'date_com' ||
                                        getFieldKeyType(
                                            filterItem.field_key
                                        ) === 'time_com' ||
                                        getFieldKeyType(
                                            filterItem.field_key
                                        ) === 'number_com' ||
                                        getFieldKeyType(
                                            filterItem.field_key
                                        ) === 'progress_com'
                                    "
                                    class="filter-content"
                                    v-overflow
                                >
                                    {{ filterItem.value }}
                                </div>
                                <!-- 单多选 取配置中的enum_values 配置标签样式 -->
                                <div
                                    v-if="
                                        getFieldKeyType(
                                            filterItem.field_key
                                        ) === 'select_com' ||
                                        getFieldKeyType(
                                            filterItem.field_key
                                        ) === 'multi_select_com'
                                    "
                                    class="filter-content"
                                >
                                    <!-- 单多选类型{{ filterItem.value }} -->
                                    <div
                                        v-overflow
                                        class="tag-item"
                                        :style="{
                                            color: '#fff',
                                            backgroundColor:
                                                getSelectFilterDetail(
                                                    filterItem.field_key,
                                                    filterItem.value
                                                ).color
                                        }"
                                    >
                                        {{ filterItem.value }}
                                    </div>
                                </div>
                                <!-- 状态  value中取标签颜色 -->
                                <div
                                    v-if="
                                        getFieldKeyType(
                                            filterItem.field_key
                                        ) === 'status_com'
                                    "
                                    class="filter-content"
                                >
                                    <!-- {{ filterItem.value }} -->
                                    <div
                                        class="tag-item"
                                        :style="{
                                            color: '#fff',
                                            backgroundColor:
                                                filterItem.value.color
                                        }"
                                    >
                                        {{ filterItem.value.name }}
                                    </div>
                                </div>
                                <!-- 关系字段 -->
                                <div
                                    v-if="
                                        getFieldKeyType(
                                            filterItem.field_key
                                        ) === 'related_com'
                                    "
                                    class="filter-content"
                                >
                                    {{ filterItem.value[0].title }}
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="action-block" v-show="actionList.length">
                <div class="title-text">执行动作</div>
                <div class="desc-block">
                    <div
                        v-for="(item, index) in actionList"
                        :key="index"
                        class="action-item"
                    >
                        <div v-if="item.action_type === 'email'">
                            <div class="action-title">
                                发送&nbsp;&nbsp;邮件&nbsp;&nbsp;通知
                            </div>
                            <div class="action-content">
                                <email-to-user
                                    :userInfoList="item.email_user"
                                ></email-to-user>
                                <div class="action-status">
                                    <b
                                        class="action-status-icon"
                                        :class="item.status"
                                    ></b>
                                    {{ getStatusName(item.status) }}
                                </div>
                            </div>
                        </div>
                        <div v-if="item.action_type === 'update'">
                            <div class="action-title">修改&nbsp;&nbsp;字段</div>
                            <!-- 分单/多选 单/多人 / 状态/ 其他 -->
                            <div
                                v-if="
                                    getFieldKeyType(item.field_key) ===
                                        'person_com' &&
                                    getFieldKeyExpr(item.field_key) === 'single'
                                "
                                class="action-content"
                            >
                                <div class="action-detail">
                                    <div class="field-action-desc">
                                        <div class="action-desc-text">
                                            将&nbsp;&nbsp;
                                        </div>
                                        <div class="action-content-text">
                                            {{
                                                getFieldKeyName(item.field_key)
                                            }}
                                        </div>
                                        <div class="action-desc-text">
                                            &nbsp;&nbsp;变更为&nbsp;&nbsp;
                                        </div>
                                    </div>
                                    <action-multiple-user
                                        class="action-tem"
                                        :userInfo="item.field_value_user"
                                    ></action-multiple-user>
                                </div>
                                <div class="action-status">
                                    <b
                                        class="action-status-icon"
                                        :class="item.status"
                                    ></b>
                                    <!-- item.status === "success"
                                            ? "执行成功"
                                            : "执行失败" -->
                                    {{ getStatusName(item.status) }}
                                </div>
                            </div>
                            <div
                                v-else-if="
                                    getFieldKeyType(item.field_key) ===
                                        'person_com' &&
                                    getFieldKeyExpr(item.field_key) ===
                                        'multiple'
                                "
                                class="action-content"
                            >
                                <div class="action-detail">
                                    <div class="field-action-desc">
                                        <div class="action-desc-text">
                                            将&nbsp;&nbsp;
                                        </div>
                                        <div class="action-content-text">
                                            {{
                                                getFieldKeyName(item.field_key)
                                            }}
                                        </div>
                                        <div class="action-desc-text">
                                            &nbsp;&nbsp;变更为&nbsp;&nbsp;
                                        </div>
                                    </div>
                                    <action-multiple-user
                                        class="action-tem"
                                        :userInfo="item.field_value_user"
                                    ></action-multiple-user>
                                </div>
                                <div class="action-status">
                                    <b
                                        class="action-status-icon"
                                        :class="item.status"
                                    ></b>
                                    {{ getStatusName(item.status) }}
                                </div>
                            </div>
                            <div
                                v-else-if="
                                    getFieldKeyType(item.field_key) ===
                                    'select_com'
                                "
                                class="action-content"
                            >
                                <div class="action-detail">
                                    <div class="field-action-desc">
                                        <div class="action-desc-text">
                                            将&nbsp;&nbsp;
                                        </div>
                                        <div class="action-content-text">
                                            {{
                                                getFieldKeyName(item.field_key)
                                            }}
                                        </div>
                                        <div class="action-desc-text">
                                            &nbsp;&nbsp;变更为&nbsp;&nbsp;
                                        </div>
                                    </div>
                                    <action-single-select
                                        class="action-tem"
                                        :option="
                                            getFieldKeyOption(item.field_key)
                                        "
                                        :selectValue="item.field_value"
                                    ></action-single-select>
                                </div>
                                <div class="action-status">
                                    <b
                                        class="action-status-icon"
                                        :class="item.status"
                                    ></b>
                                    {{ getStatusName(item.status) }}
                                </div>
                            </div>
                            <div
                                v-else-if="
                                    getFieldKeyType(item.field_key) ===
                                    'multi_select_com'
                                "
                                class="action-content"
                            >
                                <div class="action-detail">
                                    <div class="field-action-desc">
                                        <div class="action-desc-text">
                                            将&nbsp;&nbsp;
                                        </div>
                                        <div class="action-content-text">
                                            {{
                                                getFieldKeyName(item.field_key)
                                            }}
                                        </div>
                                        <div class="action-desc-text">
                                            &nbsp;&nbsp;变更为&nbsp;&nbsp;
                                        </div>
                                    </div>
                                    <action-multiple-select
                                        class="action-tem"
                                        :option="
                                            getFieldKeyOption(item.field_key)
                                        "
                                        :selectValue="item.field_value"
                                    ></action-multiple-select>
                                </div>
                                <div class="action-status">
                                    <b
                                        class="action-status-icon"
                                        :class="item.status"
                                    ></b>
                                    {{ getStatusName(item.status) }}
                                </div>
                            </div>
                            <div
                                v-else-if="
                                    getFieldKeyType(item.field_key) ===
                                    'status_com'
                                "
                                class="action-content"
                            >
                                <div class="action-detail">
                                    <div class="field-action-desc">
                                        <div class="action-desc-text">
                                            将&nbsp;&nbsp;
                                        </div>
                                        <div class="action-content-text">
                                            {{
                                                getFieldKeyName(item.field_key)
                                            }}
                                        </div>
                                        <div class="action-desc-text">
                                            &nbsp;&nbsp;变更为&nbsp;&nbsp;
                                        </div>
                                    </div>
                                    <action-status-select
                                        class="action-tem"
                                        :option="
                                            getFieldKeyOption(item.field_key)
                                        "
                                        :selectValue="item.field_value"
                                    ></action-status-select>
                                </div>
                                <div class="action-status">
                                    <b
                                        class="action-status-icon"
                                        :class="item.status"
                                    ></b>
                                    {{ getStatusName(item.status) }}
                                </div>
                            </div>
                            <!-- 关联字段 -->
                            <div
                                v-else-if="
                                    getFieldKeyType(item.field_key) ===
                                    'related_com'
                                "
                                class="action-content"
                            >
                                <div class="action-detail">
                                    <div class="field-action-desc">
                                        <div class="action-desc-text">
                                            将&nbsp;&nbsp;
                                        </div>
                                        <div class="action-content-text">
                                            {{
                                                getFieldKeyName(item.field_key)
                                            }}
                                        </div>
                                        <div class="action-desc-text">
                                            &nbsp;&nbsp;变更为&nbsp;&nbsp;
                                        </div>
                                        {{
                                            item.field_value_documents[0].title
                                        }}
                                    </div>
                                </div>
                                <div class="action-status">
                                    <b
                                        class="action-status-icon"
                                        :class="item.status"
                                    ></b>
                                    {{ getStatusName(item.status) }}
                                </div>
                            </div>
                            <div v-else class="action-content">
                                <div class="action-detail other">
                                    <div class="action-desc-text">
                                        将&nbsp;&nbsp;
                                    </div>
                                    <div class="action-content-text">
                                        {{ getFieldKeyName(item.field_key) }}
                                    </div>
                                    <div class="action-desc-text">
                                        &nbsp;&nbsp;变更为&nbsp;&nbsp;
                                    </div>
                                    <div class="text-block-content" v-overflow>
                                        {{ item.field_value }}
                                    </div>
                                </div>
                                <div class="action-status">
                                    <b
                                        class="action-status-icon"
                                        :class="item.status"
                                    ></b>
                                    {{ getStatusName(item.status) }}
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </el-dialog>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress_rule";
import ProgressApi from "@/common/api/module/progress";
import UserShow from "./user_show";
import SelectShow from "./select_show";
import RelatedShow from "./related_show";
import StatusShow from "./status_show";
import EmailToUser from "./email_to_user";
import ActionSingleUser from "./action_single_user";
import ActionMultipleUser from "./action_multiple_user";
import ActionSingleSelect from "./action_single_select";
import ActionMultipleSelect from "./action_multiple_select";
import ActionStatusSelect from "./action_status_select";
import { baseMixin } from "@/mixin.js";
import { imgHost } from "@/assets/tool/const";
import UserMessage from "@/components/user_message_tip";
import { color } from "echarts";
export default {
    mixins: [baseMixin],
    components: {
        UserShow,
        SelectShow,
        StatusShow,
        EmailToUser,
        RelatedShow,
        ActionSingleUser,
        ActionStatusSelect,
        ActionMultipleUser,
        ActionSingleSelect,
        ActionMultipleSelect,
        UserMessage
    },
    data() {
        return {
            dialogVisible: false,
            fieldList: [],
            triggerDetail: {
                taskName: "",
                triggerType: "",
                fieldKey: "",
                oldValue: "",
                oldValueUser: "",
                newValue: "",
                newValueUser: "",
                triggerDay: "",
                triggerTime: ""
            },
            actionList: [],
            filterList: [],
            formatFilterList: []
        };
    },
    methods: {
        getStatusName(status) {
            if (status === "success") {
                return "执行成功";
            }
            if (status === "fail") {
                return "执行失败";
            }
            if (status === "no") {
                return "待执行";
            }
            return "状态异常";
        },
        getFieldKeyOption(key) {
            if (
                _.find(this.fieldList, { field_key: key }) &&
                _.find(this.fieldList, { field_key: key }).enum_values
            ) {
                let optionJson = _.find(this.fieldList, {
                    field_key: key
                }).enum_values;
                let option = JSON.parse(optionJson);
                return option;
            } else {
                return [];
            }
        },
        getFieldKeyType(key) {
            if (_.find(this.fieldList, { field_key: key })) {
                return _.find(this.fieldList, { field_key: key }).mode;
            } else {
                return key;
            }
        },
        getFieldKeyExpr(key) {
            if (_.find(this.fieldList, { field_key: key })) {
                return _.find(this.fieldList, { field_key: key }).expr;
            } else {
                return "";
            }
        },
        getFieldKeyName(key) {
            if (_.find(this.fieldList, { field_key: key })) {
                return _.find(this.fieldList, { field_key: key }).name;
            } else {
                return key;
            }
        },
        getFieldKeyOp(key, op) {
            if (_.find(this.fieldList, { field_key: key })) {
                let opList = _.find(this.fieldList, { field_key: key }).op;
                let opObj = _.find(opList, { value: op });
                return opObj.label;
            } else {
                return key;
            }
        },
        getSelectFilterDetail(key, value) {
            if (_.find(this.fieldList, { field_key: key })) {
                if (_.find(this.fieldList, { field_key: key }).enum_values) {
                    let enumValuesList = JSON.parse(
                        _.find(this.fieldList, { field_key: key }).enum_values
                    );
                    if (_.find(enumValuesList, { value: value })) {
                        let selectDetail = _.find(enumValuesList, {
                            value: value
                        });
                        return selectDetail;
                    }
                }
                let obj = {
                    color: "#dddddd"
                };
                return obj;
            } else {
                let obj = {
                    color: "#dddddd"
                };
                return obj;
            }
        },
        openDialog(rowDetail) {
            this.rowDetail = rowDetail;
            this.getFilterConfig();
            this.triggerDetail.taskName = rowDetail.data_name;
            this.dialogVisible = true;
        },
        // 获取字段列表&配置
        getFilterConfig() {
            // 触发条件的字段配置
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress
            };
            ProgressApi.getFilterConfig(params).then((res) => {
                if (res.resultCode === 200) {
                    this.fieldList = res.data || [];
                    this.getDetail();
                } else {
                    this.fieldList = [];
                }
            });
        },
        getDetail() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                rule_execute_uuid: this.rowDetail.rule_execute_uuid
            };
            api.getLogDetail(params).then((res) => {
                if (res.resultCode === 200) {
                    let logDetail = res.data;
                    let rule_detail = logDetail.rule_detail;
                    let actionList = logDetail.list || [];
                    let filterList = rule_detail.filter
                        ? JSON.parse(rule_detail.filter)
                        : [];
                    this.filterList = filterList;
                    let formatFilterList = rule_detail.format_filter
                        ? JSON.parse(rule_detail.format_filter)
                        : [];
                    this.formatFilterList = formatFilterList;
                    let obj = {
                        triggerType: rule_detail.rule_type,
                        fieldKey: rule_detail.field_key,
                        oldValue: rule_detail.old_value,
                        oldValueUser: rule_detail.old_value_user,
                        oldValueDocuments: rule_detail.old_value_documents,
                        newValue: rule_detail.new_value,
                        newValueUser: rule_detail.new_value_user,
                        newValueDocuments: rule_detail.new_value_documents,
                        triggerDay: rule_detail.trigger_day,
                        triggerTime: rule_detail.trigger_time
                    };
                    Object.assign(this.triggerDetail, obj);
                    this.actionList = actionList;
                }
            });
        },
        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            } else {
                return require(`@/assets/image/common/default_avatar_big.png`);
            }
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
::v-deep.log-detail-dialog.basic-ui {
    .el-dialog {
        .el-dialog__header {
            padding: 17px 30px;
            border-bottom: 2px solid #e6e9f0;
        }
        .el-dialog__body {
            max-height: 70vh;
            height: 70vh;
            padding: 32px 30px 0;
            margin-bottom: 32px;
        }
    }
}
.user-show,
.select-show,
.text-change-show {
    display: inline-block;
    box-sizing: border-box;
    height: 40px;
}
.text-change-show {
    line-height: 40px;
}

.title-text {
    font-weight: 500;
    font-size: 20px;
    color: #171e31;
}
.desc-text {
    font-size: 18px;
    color: #82889e;
    line-height: 40px;
    &.title {
        line-height: 18px;
        margin-bottom: 12px;
    }
}
.normal-text {
    font-size: 14px;
    color: #171e31;
    .desc-text {
        font-size: 14px;
    }
}
.desc-block {
    padding: 20px 24px;
    background-color: #fafafb;
    border-radius: 8px;
    margin-top: 20px;
    border: 1px solid #e6e9f0;
    &.trigger {
        margin-bottom: 20px;
    }
}
.field-change {
    margin-top: 16px;
}
.task-name {
    height: 40px;
    line-height: 40px;
}
.task-name,
.schedule-block,
.field-change-text {
    // margin-top: 8px;
    display: flex;
    padding: 0 12px;
    height: 40px;
    background-color: #fff;
    border-radius: 4px;
    border: 1px solid #e6e9f0;
    box-sizing: border-box;
    .desc-text-block {
        display: inline-block;
        height: 40px;
        line-height: 40px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }
}
.filter-list {
    margin-top: 16px;
    .filter-item {
        margin-bottom: 8px;
        display: flex;
        &.first {
            .filter-body {
                width: 100%;
            }
        }
        &:last-child {
            margin-bottom: 0px;
        }
        .filter-head-and {
            font-size: 14px;
            color: #82889e;
            line-height: 40px;
            margin-right: 4px;
            display: flex;
            padding: 0 12px;
            width: 40px;
            min-width: 40px;
            height: 40px;
            background-color: #fff;
            border-radius: 4px;
            border: 1px solid #e6e9f0;
            box-sizing: border-box;
        }
        .filter-body {
            width: calc(100% - 44px);
            display: flex;
            padding: 0 12px;
            height: 40px;
            background-color: #fff;
            border-radius: 4px;
            border: 1px solid #e6e9f0;
            box-sizing: border-box;
            .filter-field-key {
                margin-right: 16px;
                display: inline-block;
                max-width: 220px;
                height: 40px;
                line-height: 40px;
                white-space: nowrap;
                overflow: hidden;
                text-overflow: ellipsis;
            }
            .filter-relation {
                max-width: 220px;
                margin-right: 16px;
                font-size: 14px;
                color: #82889e;
                line-height: 40px;
                white-space: nowrap;
                overflow: hidden;
                text-overflow: ellipsis;
            }
            .filter-content {
                flex: 1;
                min-width: calc(100% - 440px);
                display: inline-block;
                height: 40px;
                line-height: 40px;
                white-space: nowrap;
                overflow: hidden;
                text-overflow: ellipsis;
                .tag-item {
                    margin: 8px 8px 8px 0px;
                    display: inline-block;
                    height: 24px;
                    line-height: 24px;
                    text-align: center;
                    padding: 0 8px;
                    border-radius: 4px;
                }
                .member-item {
                    line-height: 40px;
                }
            }
        }
    }
}

.schedule-block {
    line-height: 40px;
}

.action-item {
    margin-bottom: 16px;
    &:last-child {
        margin-bottom: 0;
    }
    .action-title {
        margin-bottom: 12px;
        font-size: 18px;
        line-height: 18px;
        color: #66677c;
    }
    .action-content {
        box-sizing: border-box;
        background: #ffffff;
        border-radius: 4px 4px 4px 4px;
        border: 1px solid #e6e9f0;
        height: 40px;
        line-height: 40px;
        padding: 0 12px;
        display: flex;
        justify-content: space-between;
        .action-detail {
            display: flex;

            .field-action-desc {
                display: inline-block;
                white-space: nowrap;
            }
            .action-desc-text {
                display: inline-block;
                color: #82889e;
            }
            .action-content-text {
                display: inline-block;
            }
            .text-block-content {
                display: inline-block;
                white-space: nowrap;
                overflow: hidden;
                text-overflow: ellipsis;
            }
            .action-tem {
                display: inline-block;
                height: 40px;
            }
            &.other {
                display: flex;
                max-width: calc(100% - 140px);
                white-space: nowrap;
                overflow: hidden;
                text-overflow: ellipsis;
            }
        }
        .action-status {
            display: inline-block;
            white-space: nowrap;
            .action-status-icon {
                display: inline-block;
                width: 12px;
                height: 12px;
                background-size: 100% 100%;
                &.success {
                    background-image: url(@/assets/image/progress/add_rule/success.png);
                }
                &.fail {
                    background-image: url(@/assets/image/progress/add_rule/fail.png);
                }
                &.no {
                    background-image: url(@/assets/image/progress/add_rule/wait.svg);
                }
            }
        }
    }
}
</style>
