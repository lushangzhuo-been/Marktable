<template>
    <el-dialog
        title="添加自动化规则"
        :visible.sync="dialogVisible"
        class="add-rule-dialog basic-ui"
        :width="'1180px'"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :destroy-on-close="true"
        append-to-body
        @close="cancel"
    >
        <div v-if="dialogVisible" class="base-msg-block">
            <el-form
                ref="formInfo"
                :model="formInfo"
                class="basic-ui base-info"
                :rules="formRules"
            >
                <div class="form-row">
                    <el-form-item
                        class="basic-ui form-name"
                        label="名称"
                        prop="name"
                    >
                        <el-input
                            v-model="formInfo.name"
                            class="basic-ui"
                            size="small"
                        ></el-input>
                    </el-form-item>
                    <el-form-item
                        class="basic-ui form-desc"
                        label="描述"
                        prop="description"
                    >
                        <el-input
                            v-model="formInfo.description"
                            class="basic-ui"
                            size="small"
                        ></el-input>
                    </el-form-item>
                </div>
            </el-form>
        </div>
        <div v-if="dialogVisible" class="body-block">
            <div class="trigger-block-outter">
                <div class="guide-block">
                    <div class="title">
                        <b class="type-icon trigger-icon"></b>
                        当
                    </div>
                    <div class="desc">某个条件发生时</div>
                </div>
                <div ref="ConditionBlock" class="condition-block">
                    <div ref="ConditionBlockContent" class="scroll-block">
                        <div class="trigger-condition">
                            <div
                                class="type-select-border"
                                :class="{
                                    'had-detail':
                                        rule_type === 'update' ||
                                        rule_type === 'schedule'
                                }"
                            >
                                <el-select
                                    class="basic-ui field-select"
                                    size="small"
                                    v-model="rule_type"
                                    placeholder="请选择"
                                >
                                    <el-option
                                        v-for="item in triggerOptions"
                                        :key="item.value"
                                        :label="item.label"
                                        :value="item.value"
                                    >
                                        <div>
                                            <b
                                                class="type-box"
                                                :class="item.value"
                                            ></b>
                                            {{ item.label }}
                                        </div>
                                    </el-option>
                                </el-select>
                            </div>
                            <div
                                v-if="
                                    rule_type === 'update' ||
                                    rule_type === 'schedule'
                                "
                                class="detail-block"
                            >
                                <!-- 字段变更 -->
                                <div v-if="rule_type === 'update'">
                                    <!-- 流程的字段从a变成b -->
                                    <div>
                                        <el-input
                                            class="basic-ui field-select"
                                            size="small"
                                            placeholder="请输入条件"
                                            v-model="templateName"
                                            disabled
                                        ></el-input>
                                    </div>
                                    <div class="desc">的</div>
                                    <div>
                                        <el-select
                                            class="basic-ui field-select"
                                            size="small"
                                            v-model="field_key"
                                            placeholder="请选择字段"
                                            @change="fieldKeyChange"
                                        >
                                            <el-option
                                                v-for="item in filterFieldList"
                                                :key="item.value"
                                                :label="item.name"
                                                :value="item.field_key"
                                            >
                                                <b
                                                    class="type-box"
                                                    :style="getType(item.mode)"
                                                ></b>
                                                {{ item.name }}
                                            </el-option>
                                        </el-select>
                                    </div>
                                    <!-- 日期 时间 文本 数值 进展 -》 “变更时”替代 -->
                                    <div
                                        class="desc end"
                                        v-if="
                                            (field_key &&
                                                getTriggerTyple(field_key)
                                                    .mode === 'number_com') ||
                                            getTriggerTyple(field_key).mode ===
                                                'text_com' ||
                                            getTriggerTyple(field_key).mode ===
                                                'textarea_com' ||
                                            getTriggerTyple(field_key).mode ===
                                                'time_com' ||
                                            getTriggerTyple(field_key).mode ===
                                                'date_com' ||
                                            getTriggerTyple(field_key).mode ===
                                                'progress_com'
                                        "
                                    >
                                        变更时
                                    </div>
                                    <div v-else>
                                        <div v-if="field_key" class="desc">
                                            从
                                        </div>
                                        <!-- 单选 -->
                                        <div
                                            v-if="
                                                getTriggerTyple(field_key)
                                                    .mode === 'select_com'
                                            "
                                        >
                                            <select-change
                                                v-model="old_value"
                                                :option="
                                                    getFieldOption(field_key)
                                                "
                                            ></select-change>
                                            <div class="desc">到</div>
                                            <select-change
                                                v-model="new_value"
                                                :option="
                                                    getFieldOption(field_key)
                                                "
                                            ></select-change>
                                        </div>
                                        <!-- 多选 -->
                                        <div
                                            v-if="
                                                getTriggerTyple(field_key)
                                                    .mode === 'multi_select_com'
                                            "
                                        >
                                            2
                                            <multiple-select-change
                                                v-model="old_value"
                                                :option="
                                                    getFieldOption(field_key)
                                                "
                                            ></multiple-select-change>
                                            <div class="desc">到</div>
                                            <select-change
                                                v-model="new_value"
                                                :option="
                                                    getFieldOption(field_key)
                                                "
                                            ></select-change>
                                        </div>
                                        <!-- 人 -->
                                        <div
                                            v-if="
                                                getTriggerTyple(field_key)
                                                    .mode === 'person_com'
                                            "
                                        >
                                            <people-change
                                                v-model="old_value"
                                                :userInfoList="oldUserListInfo"
                                            ></people-change>
                                            <div class="desc">到</div>
                                            <people-change
                                                v-model="new_value"
                                                :userInfoList="newUserListInfo"
                                            ></people-change>
                                        </div>
                                        <!-- 状态 -->
                                        <div
                                            v-if="
                                                getTriggerTyple(field_key)
                                                    .mode === 'status_com'
                                            "
                                        >
                                            <status-change
                                                v-model="old_value"
                                                :option="
                                                    getFieldOption(field_key)
                                                "
                                            ></status-change>
                                            <div class="desc">到</div>
                                            <status-change
                                                v-model="new_value"
                                                :option="
                                                    getFieldOption(field_key)
                                                "
                                            ></status-change>
                                        </div>
                                        <!-- 关系 -->
                                        <div
                                            v-if="
                                                getTriggerTyple(field_key)
                                                    .mode === 'related_com'
                                            "
                                        >
                                            <!-- 关系a -->
                                            <relation-change
                                                v-model="old_value"
                                                :relationInfo="
                                                    getTriggerTyple(field_key)
                                                "
                                                :relationReshowInfo="
                                                    oldRelationListInfo
                                                "
                                            ></relation-change>
                                            <div class="desc">到</div>
                                            <!-- 关系b -->
                                            <relation-change
                                                v-model="new_value"
                                                :relationInfo="
                                                    getTriggerTyple(field_key)
                                                "
                                                :relationReshowInfo="
                                                    newRelationListInfo
                                                "
                                            ></relation-change>
                                        </div>
                                    </div>
                                </div>
                                <!-- 定时触发 -->
                                <div v-if="rule_type === 'schedule'">
                                    <div>
                                        <el-select
                                            class="basic-ui field-select"
                                            size="small"
                                            v-model="triggerDateFieldKey"
                                            placeholder="请选择日期字段"
                                        >
                                            <el-option
                                                v-for="item in dateFieldList"
                                                :key="item.value"
                                                :label="item.name"
                                                :value="item.field_key"
                                            >
                                                <div>
                                                    <b
                                                        class="type-box"
                                                        :style="
                                                            getType(item.mode)
                                                        "
                                                    ></b>
                                                    {{ item.name }}
                                                </div>
                                            </el-option>
                                        </el-select>
                                    </div>
                                    <div class="msg-transform">
                                        <span class="desc"
                                            >到达<span class="time"
                                                >{{
                                                    triggerDay === 0
                                                        ? "当天"
                                                        : triggerDay < 0
                                                        ? `在${-triggerDay}天前的`
                                                        : `在${triggerDay}天后的`
                                                }}{{ triggerTime }}</span
                                            ></span
                                        >
                                    </div>
                                    <div class="day-and-time">
                                        <!-- 日期 -->
                                        <div class="day-block">
                                            <el-input-number
                                                v-model="triggerDay"
                                                controls-position="right"
                                                :min="-365"
                                                :max="365"
                                                width="180"
                                            >
                                            </el-input-number>
                                            <div class="define-append">天</div>
                                        </div>
                                        <!-- 时间 -->
                                        <div class="time-block">
                                            <el-time-picker
                                                v-model="triggerTime"
                                                format="HH:mm"
                                                value-format="HH:mm"
                                                placeholder="08:00"
                                            >
                                            </el-time-picker>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <!-- filter条件列表 -->
                        <div
                            class="trigger-item list-item"
                            v-for="(triggerItem, triggerIndex) in triggerList"
                            :key="triggerIndex"
                        >
                            <!-- 字段 -->
                            <div class="field-row">
                                <span class="desc">且</span>
                                <el-select
                                    class="basic-ui"
                                    size="small"
                                    v-model="triggerItem.field_key"
                                    placeholder="请选择字段"
                                    @change="
                                        (val) => filterFieldChange(triggerItem)
                                    "
                                >
                                    <el-option
                                        v-for="item in filterFieldList"
                                        :key="item.value"
                                        :label="item.name"
                                        :value="item.field_key"
                                    >
                                        <b
                                            class="type-box"
                                            :style="getType(item.mode)"
                                        ></b>
                                        {{ item.name }}
                                    </el-option>
                                </el-select>
                            </div>
                            <!-- 关系 -->
                            <div class="relation-value-row">
                                <el-select
                                    class="basic-ui op-select"
                                    size="small"
                                    v-model="triggerItem.op"
                                    placeholder="请选择关系"
                                >
                                    <el-option
                                        v-for="item in getFilterRelation(
                                            triggerItem.field_key
                                        )"
                                        :key="item.value"
                                        :label="item.label"
                                        :value="item.value"
                                    >
                                    </el-option>
                                </el-select>
                                <!-- 值 值需要按类型 -->
                                <div class="trigger-value-border">
                                    <!-- 文字 -->
                                    <el-input
                                        v-if="
                                            getTriggerFilterFieldType(
                                                triggerItem
                                            ) === 'text'
                                        "
                                        class="basic-ui value"
                                        size="small"
                                        v-model="triggerItem.value"
                                        placeholder="请输入条件"
                                    ></el-input>
                                    <!-- 数值 -->
                                    <el-input
                                        v-else-if="
                                            getTriggerFilterFieldType(
                                                triggerItem
                                            ) === 'number'
                                        "
                                        class="basic-ui value progress-input number hidden-arrow"
                                        type="number"
                                        size="small"
                                        v-model="triggerItem.value"
                                        placeholder="请输入条件"
                                    ></el-input>
                                    <!-- 单多人 -->
                                    <div
                                        v-else-if="
                                            getTriggerFilterFieldType(
                                                triggerItem
                                            ) === 'person'
                                        "
                                    >
                                        <person-filter
                                            :fieldKey="triggerItem.field_key"
                                            :triggerListReshow="
                                                triggerListReshow
                                            "
                                            v-model="triggerItem.value"
                                        ></person-filter>
                                    </div>
                                    <!-- 单多选 -->
                                    <div
                                        v-else-if="
                                            getTriggerFilterFieldType(
                                                triggerItem
                                            ) === 'select'
                                        "
                                    >
                                        <new-select-item
                                            :option="
                                                getFilterOption(triggerItem)
                                            "
                                            :param="triggerItem.value"
                                            v-model="triggerItem.value"
                                        >
                                        </new-select-item>
                                    </div>
                                    <!-- 状态 -->
                                    <div
                                        v-else-if="
                                            getTriggerFilterFieldType(
                                                triggerItem
                                            ) === 'status'
                                        "
                                    >
                                        <status-filter
                                            :fieldList="filterFieldList"
                                            :actionItem="triggerItem"
                                            v-model="triggerItem.value"
                                        ></status-filter>
                                    </div>
                                    <!-- 时间 -->
                                    <div
                                        v-else-if="
                                            getTriggerFilterFieldType(
                                                triggerItem
                                            ) === 'time'
                                        "
                                    >
                                        <!-- 时间 -->
                                        <el-date-picker
                                            ref="DatePicker"
                                            popper-class="progress-date-picker-popover"
                                            v-model="triggerItem.value"
                                            type="datetime"
                                            :editable="false"
                                            placeholder="选择时间"
                                            value-format="yyyy-MM-dd HH:mm:ss"
                                            size="small"
                                        >
                                        </el-date-picker>
                                    </div>
                                    <!-- 日期 -->
                                    <div
                                        v-else-if="
                                            getTriggerFilterFieldType(
                                                triggerItem
                                            ) === 'date'
                                        "
                                    >
                                        <!-- 日期 -->
                                        <el-date-picker
                                            ref="DatePicker"
                                            popper-class="progress-date-picker-popover"
                                            v-model="triggerItem.value"
                                            type="date"
                                            :editable="false"
                                            placeholder="选择日期"
                                            value-format="yyyy-MM-dd"
                                            size="small"
                                        >
                                        </el-date-picker>
                                    </div>
                                    <!-- 关联关系 -->
                                    <div
                                        v-else-if="
                                            getTriggerFilterFieldType(
                                                triggerItem
                                            ) === 'related'
                                        "
                                    >
                                        <!-- {{ triggerItem }} -->
                                        <related-filter
                                            :fieldKey="triggerItem.field_key"
                                            :triggerListReshow="
                                                triggerListReshow
                                            "
                                            :relationInfo="
                                                getTriggerTyple(
                                                    triggerItem.field_key
                                                )
                                            "
                                            v-model="triggerItem.value"
                                        ></related-filter>
                                    </div>
                                    <el-input
                                        v-else
                                        class="basic-ui"
                                        size="small"
                                        placeholder="请输入条件"
                                        v-model="triggerItem.value"
                                    ></el-input>
                                </div>

                                <b
                                    class="delete-icon"
                                    @click="
                                        deleteTrigger(triggerItem, triggerIndex)
                                    "
                                ></b>
                            </div>
                        </div>
                        <!-- 添加条件按钮 -->
                        <div class="add-trigger-block" @click="addTrigger">
                            <b class="add-trigger-icon"></b>
                        </div>
                    </div>
                </div>
                <!-- 中线 -->
                <div class="middle-line"></div>
            </div>
            <div class="middle-block">
                <div class="middle-btn"></div>
                <div class="middle-line"></div>
            </div>
            <div class="action-block-outter">
                <div class="guide-block">
                    <div class="title">
                        <b class="type-icon action-icon"></b>
                        则
                    </div>
                    <div class="desc">执行某个动作</div>
                </div>
                <div class="action-block" ref="ActionBlock">
                    <div ref="ActionBlockContent" class="scroll-block">
                        <!-- 必要的动作 -->
                        <div class="action-detail">
                            <div class="action-select">
                                <el-select
                                    class="basic-ui width100"
                                    size="small"
                                    v-model="necessaryAction.action_type"
                                    @change="necessaryActionTypeChange"
                                    placeholder="请选择"
                                >
                                    <el-option
                                        v-for="item in actionOptions"
                                        :key="item.value"
                                        :label="item.label"
                                        :value="item.value"
                                    >
                                        <b
                                            class="type-box"
                                            :class="item.value"
                                        ></b>
                                        {{ item.label }}
                                    </el-option>
                                </el-select>
                            </div>
                            <div class="detail-block action-item">
                                <!-- 修改字段 -->
                                <div
                                    v-if="
                                        necessaryAction.action_type === 'update'
                                    "
                                >
                                    <div class="desc if">将</div>
                                    <el-select
                                        class="basic-ui field-select"
                                        size="small"
                                        v-model="necessaryAction.field_key"
                                        @change="
                                            (value) =>
                                                fieldChange(
                                                    value,
                                                    necessaryAction
                                                )
                                        "
                                        placeholder="请选择字段"
                                    >
                                        <el-option
                                            v-for="item in actionFieldList"
                                            :key="item.value"
                                            :label="item.name"
                                            :value="item.field_key"
                                        >
                                            <b
                                                class="type-box"
                                                :style="getType(item.mode)"
                                            ></b>
                                            {{ item.name }}
                                        </el-option>
                                    </el-select>
                                    <div class="desc">变更为</div>
                                    <set-target-value
                                        :fieldList="actionFieldList"
                                        :actionItem="necessaryAction"
                                        v-model="necessaryAction.field_value"
                                    ></set-target-value>
                                </div>
                                <!-- 发送通知 -->
                                <div
                                    v-if="
                                        necessaryAction.action_type === 'email'
                                    "
                                >
                                    <div class="desc if">发送</div>
                                    <div>
                                        <email
                                            :title="necessaryAction.email_title"
                                            v-model="
                                                necessaryAction.email_contents
                                            "
                                            :fieldList="actionFieldList"
                                            :id="'necessaryAction'"
                                            @email-update="
                                                (title, content) =>
                                                    emailUpdate(
                                                        title,
                                                        content,
                                                        'necessaryAction'
                                                    )
                                            "
                                        ></email>
                                    </div>
                                    <div class="desc">给</div>
                                    <div>
                                        <email-to
                                            placeholder="收件人"
                                            :userList="
                                                necessaryAction.user_list
                                            "
                                            :userInfoList="
                                                necessaryAction.user_list_info
                                            "
                                            :roleCheckList="
                                                necessaryAction.issue_role_list
                                            "
                                            @userlist-change="
                                                (arr) =>
                                                    userChange(
                                                        arr,
                                                        'necessaryAction'
                                                    )
                                            "
                                            @rolelist-change="
                                                (arr) =>
                                                    roleChange(
                                                        arr,
                                                        'necessaryAction'
                                                    )
                                            "
                                        ></email-to>
                                    </div>
                                </div>
                            </div>
                            <b
                                v-if="actionList.length"
                                class="delete-icon"
                                @click="deleteNecessaryAction()"
                            ></b>
                        </div>
                        <!-- 补充的动作列表 -->
                        <div
                            class="action-item list-item"
                            v-for="(actionItem, actionIndex) in actionList"
                            :key="actionIndex"
                        >
                            <div v-if="actionItem.action_type === 'update'">
                                <div class="desc if">将</div>
                                <el-select
                                    class="basic-ui field-select"
                                    size="small"
                                    v-model="actionItem.field_key"
                                    @change="
                                        (value) =>
                                            fieldChange(value, actionItem)
                                    "
                                    placeholder="字段"
                                >
                                    <el-option
                                        v-for="item in actionFieldList"
                                        :key="item.value"
                                        :label="item.name"
                                        :value="item.field_key"
                                    >
                                        <b
                                            class="type-box"
                                            :style="getType(item.mode)"
                                        ></b>
                                        {{ item.name }}
                                    </el-option>
                                </el-select>
                                <div class="desc">变更为</div>
                                <set-target-value
                                    :fieldList="actionFieldList"
                                    :actionItem="actionItem"
                                    v-model="actionItem.field_value"
                                ></set-target-value>
                                <b
                                    class="delete-icon"
                                    @click="
                                        deleteAction(actionItem, actionIndex)
                                    "
                                ></b>
                            </div>
                            <div v-if="actionItem.action_type === 'email'">
                                <div class="desc if">发送</div>
                                <div>
                                    <email
                                        :title="actionItem.email_title"
                                        v-model="actionItem.email_contents"
                                        :id="actionIndex"
                                        :fieldList="actionFieldList"
                                        @email-update="
                                            (title, content) =>
                                                emailUpdate(
                                                    title,
                                                    content,
                                                    actionIndex
                                                )
                                        "
                                    ></email>
                                </div>
                                <div class="desc">给</div>
                                <div>
                                    <email-to
                                        placeholder="收件人"
                                        :userList="actionItem.user_list"
                                        :userInfoList="
                                            actionItem.user_list_info
                                        "
                                        :roleCheckList="
                                            actionItem.issue_role_list
                                        "
                                        @userlist-change="
                                            (arr) =>
                                                userChange(arr, actionIndex)
                                        "
                                        @rolelist-change="
                                            (arr) =>
                                                roleChange(arr, actionIndex)
                                        "
                                    ></email-to>
                                </div>
                                <b
                                    class="delete-icon"
                                    @click="
                                        deleteAction(actionItem, actionIndex)
                                    "
                                ></b>
                            </div>
                        </div>
                        <!-- 添加条件按钮-选类型 -->
                        <el-dropdown
                            trigger="click"
                            class="add-action-dropdown"
                            placement="bottom"
                            @command="addAction"
                        >
                            <div class="add-action-block">
                                <b class="add-action-icon"></b>
                            </div>
                            <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item
                                    class="basic-ui"
                                    command="editField"
                                >
                                    <b class="opertion-item-box rename"></b>
                                    修改字段
                                </el-dropdown-item>
                                <el-dropdown-item
                                    class="basic-ui"
                                    command="sendMsg"
                                >
                                    <b class="opertion-item-box delete"></b>
                                    发送通知
                                </el-dropdown-item>
                            </el-dropdown-menu>
                        </el-dropdown>
                    </div>
                </div>
                <!-- 中线 -->
                <div class="middle-line"></div>
            </div>
        </div>
        <span slot="footer" class="dialog-footer">
            <el-button size="small" @click="cancel">取消</el-button>
            <el-button
                class="basic-ui"
                size="small"
                type="primary"
                @click="onConfirm"
                >{{
                    operationType === "edit" ? "确定" : "添加至列表"
                }}</el-button
            >
        </span>
    </el-dialog>
</template>

<script>
import _ from "lodash";
import { baseMixin } from "@/mixin.js";
import api from "@/common/api/module/progress_rule";
import settingApi from "@/common/api/module/progress_setting";
import ProgressApi from "@/common/api/module/progress";
import SetTargetValue from "@/pages/progress/customer_rule/action/target_value.vue";
import Email from "@/pages/progress/customer_rule/action/email.vue";
import EmailTo from "@/pages/progress/customer_rule/action/email_to.vue";
import PeopleChange from "@/pages/progress/customer_rule/trigger/people_change.vue";
import SelectChange from "@/pages/progress/customer_rule/trigger/select_change.vue";
import MultipleSelectChange from "@/pages/progress/customer_rule/trigger/select_change.vue";
import RelationChange from "@/pages/progress/customer_rule/trigger/relation_change.vue";
import StatusChange from "@/pages/progress/customer_rule/trigger/status_change.vue";
import NewSelectItem from "@/pages/progress/filter/filter_item/new_select";
// import PersonItem from "@/pages/progress/filter/filter_item/person";
// src\pages\progress\customer_rule\filter\person.vue
// src\pages\progress\customer_rule\filter\person.vue
import PersonFilter from "@/pages/progress/customer_rule/filter/person.vue";
import RelatedFilter from "@/pages/progress/customer_rule/filter/related.vue";
// import StatusFilter from "@/pages/progress/filter/filter_item/status_select.vue";
import StatusFilter from "@/pages/progress/customer_rule/action/status_select.vue";
import { FieldType } from "@/assets/tool/const";

export default {
    mixins: [baseMixin],
    components: {
        PeopleChange,
        SelectChange,
        MultipleSelectChange,
        StatusChange,
        RelationChange,
        // EmailPeople,
        Email,
        EmailTo,
        NewSelectItem,
        PersonFilter,
        SetTargetValue,
        StatusFilter,
        RelatedFilter
    },
    props: {},
    data() {
        return {
            dialogVisible: false,
            operationType: "",
            detail: null,
            formInfo: {
                name: "",
                description: ""
            },
            formRules: {
                name: [
                    {
                        required: true,
                        message: "请输入规则名称",
                        trigger: "blur"
                    }
                ]
            },
            // 触发
            triggerOptions: [],
            rule_type: null, // 触发类型
            // 日期字段选择
            dateOptions: [
                {
                    value: 1,
                    label: "创建日期"
                },
                {
                    value: 2,
                    label: "更新日期"
                },
                {
                    value: 3,
                    label: "闭环日期"
                }
            ],
            // fieldList: [], // 字段列表
            filterFieldList: [], // 触发器 字段
            dateFieldList: [], // 定时触发-》字段 mode为date的
            actionFieldList: [], // 动作列表-变更
            triggerDateFieldKey: "", // 定时触发的字段
            triggerDay: 0,
            triggerTime: "08:00",
            triggerList: [],
            triggerListReshow: [], //回显用
            // 必要的动作 分为邮件 和 修改字段
            necessaryAction: {
                action_type: "email",
                email_title: "",
                email_contents: "",
                user_list: ""
            },
            // 动作类型列表
            actionOptions: [],
            // 动作列表
            actionList: [],
            field_key: "", // 触发的字段
            new_value: null, // 准备好入参的字符串
            old_value: null, // 准备好入参的字符串
            newUserListInfo: [], // 新旧人如果为人时的全量信息
            oldUserListInfo: [],
            oldRelationListInfo: [], // 关系回显
            newRelationListInfo: [],
            filterFieldType: [
                "text_com",
                "person_com",
                "status_com",
                "time_com",
                "date_com",
                "progress_com",
                "number_com",
                "select_com",
                // 'link_com',
                "textarea_com",
                "multi_select_com",
                "related_com"
            ],
            templateName: ""
        };
    },
    computed: {},
    watch: {
        dialogVisible: {
            handler(Boolean) {
                if (!Boolean) {
                    this.resetForm();
                }
            }
        }
    },
    methods: {
        getFieldKeyInfo(key) {
            let info = _.find(this.filterFieldList, { field_key: key });
            return info;
        },
        filterFieldChange(filterItem) {
            this.$set(filterItem, "op", "");
            this.$set(filterItem, "value", "");
            // 如果关联关系字段变更 要清空tem列表
        },
        // 触发条件 字段对应的关系
        getFilterRelation(field_key) {
            let option;
            if (field_key) {
                option = _.find(this.filterFieldList, {
                    field_key: field_key
                }).op;
            }
            if (option) {
                return option;
            } else {
                return [];
            }
        },
        // 动作中的字段变更动作
        fieldKeyChange(field_key) {
            if (
                this.getTriggerTyple(field_key).mode === "number_com" ||
                this.getTriggerTyple(field_key).mode === "text_com" ||
                this.getTriggerTyple(field_key).mode === "textarea_com" ||
                this.getTriggerTyple(field_key).mode === "time_com" ||
                this.getTriggerTyple(field_key).mode === "date_com" ||
                this.getTriggerTyple(field_key).mode === "progress_com"
            ) {
                this.old_value = "任意值";
                this.new_value = "任意值";
            } else {
                this.old_value = "";
                this.oldUserListInfo = [];
                this.new_value = "";
                this.newUserListInfo = [];
            }
        },
        // 动作字段变更
        fieldChange(value, item) {
            item.field_value = "";
        },
        // 触发器filter获取字段的类型
        getTriggerFilterFieldType(item) {
            if (_.find(this.filterFieldList, { field_key: item.field_key })) {
                let type = _.find(this.filterFieldList, {
                    field_key: item.field_key
                }).mode;
                if (
                    type === "text_com" ||
                    type === "textarea_com" ||
                    type === "html_text_com" ||
                    !type
                ) {
                    return "text";
                } else if (
                    type === "select_com" ||
                    type === "multi_select_com"
                ) {
                    return "select";
                } else if (type === "status_com") {
                    return "status";
                } else if (type === "person_com") {
                    return "person";
                } else if (type === "date_com") {
                    return "date";
                } else if (type === "time_com") {
                    return "time";
                } else if (type === "number_com") {
                    return "number";
                } else if (type === "related_com") {
                    return "related";
                } else {
                    return "text";
                }
            }
        },
        // 获取filter字段的关系列表
        getFilterOption(item) {
            let option;
            if (item.field_key) {
                option = _.find(this.filterFieldList, {
                    field_key: item.field_key
                }).enum_values;
                if (option) {
                    return JSON.parse(option);
                } else {
                    return [];
                }
            } else {
                return [];
            }
        },
        // 触发单/多标签
        triggerValueChange(arr, mark) {
            if (mark === "before") {
                this.old_value = arr.join(",");
            } else {
                this.new_value = arr.join(",");
            }
        },
        // 邮件内容更新
        emailUpdate(title, content, mark) {
            if (mark === "necessaryAction") {
                this.necessaryAction.email_title = title;
            } else {
                this.actionList[mark].email_title = title;
            }
        },
        // 动作-邮件-用户选择
        userChange(arr, mark) {
            if (mark === "necessaryAction") {
                this.necessaryAction.user_list = arr.join(",");
            } else {
                this.actionList[mark].user_list = arr.join(",");
            }
        },
        roleChange(arr, mark) {
            // necessaryAction
            if (mark === "necessaryAction") {
                this.necessaryAction.issue_role_list = arr.join(",");
            } else {
                this.actionList[mark].issue_role_list = arr.join(",");
            }
        },

        // 获取当前字段类型
        getTriggerTyple(field) {
            if (field) {
                let detail;
                if (_.find(this.filterFieldList, { field_key: field })) {
                    detail = _.find(this.filterFieldList, { field_key: field });
                    return detail;
                }
            } else {
                return "未选择";
            }
        },
        // 获取当前字段可配置项
        getFieldOption(field) {
            if (field && _.find(this.filterFieldList, { field_key: field })) {
                let type;
                type = _.find(this.filterFieldList, {
                    field_key: field
                }).enum_values;
                if (type) {
                    return JSON.parse(type);
                } else {
                    return [];
                }
            } else {
                return [];
            }
        },
        // 必要的动作
        necessaryActionTypeChange(type) {
            if (type === "update") {
                this.necessaryAction = {
                    action_type: "update",
                    field_key: "",
                    field_value: ""
                };
            }
            if (type === "email") {
                this.necessaryAction = {
                    action_type: "email",
                    email_title: "",
                    email_contents: "",
                    user_list: "",
                    issue_role_list: ""
                };
            }
        },
        // 添加触发条件
        addTrigger() {
            let option = {
                field_key: "",
                op: "",
                value: ""
            };
            this.triggerList.push(option);
            this.$nextTick(() => {
                let height = this.$refs.ConditionBlockContent.clientHeight;
                this.$refs.ConditionBlock.scrollTo({
                    top: height
                });
            });
        },
        deleteTrigger(item, index) {
            this.triggerList.splice(index, 1);
        },
        // 动作
        addAction(command) {
            let obj = {};
            if (command === "editField") {
                obj = {
                    action_type: "update",
                    field_key: "",
                    field_value: ""
                };
            }
            if (command === "sendMsg") {
                obj = {
                    action_type: "email",
                    email_title: "",
                    email_contents: "",
                    user_list: ""
                };
            }
            this.actionList.push(obj);
            this.$nextTick(() => {
                let height = this.$refs.ActionBlockContent.clientHeight;
                this.$refs.ActionBlock.scrollTo({
                    top: height
                });
            });
        },
        deleteAction(item, index) {
            this.actionList.splice(index, 1);
        },
        deleteNecessaryAction() {
            // 将actionList第一个赋给necessaryAction   并删除
            let newNecessaryAction = this.actionList.splice(0, 1);
            this.necessaryAction = {
                ...newNecessaryAction[0]
            };
        },
        cancel() {
            this.dialogVisible = false;
        },
        onConfirm() {
            // 校验
            this.$refs.formInfo.validate((vali) => {
                if (vali) {
                    // 触发器条件
                    let triggerFilter = _.cloneDeep(this.triggerList);
                    if (triggerFilter && triggerFilter.length) {
                        triggerFilter.forEach((item) => {
                            if (item.value instanceof Array) {
                                let arr = [];
                                item.value.forEach((iitem) => {
                                    arr.push(iitem.id);
                                });
                                item.value = arr.join(",");
                            }
                        });
                    }
                    // 动作列表
                    let action_list = [
                        this.necessaryAction,
                        ...this.actionList
                    ];
                    // 入参参数
                    let params = {
                        ...this.formInfo,
                        // 触发
                        ws_id: this.curSpace.id,
                        tmpl_id: this.curProgress,
                        rule_type: this.rule_type,
                        filter: triggerFilter.length
                            ? JSON.stringify(triggerFilter)
                            : "",
                        // 动作
                        action_list: JSON.stringify(action_list)
                    };
                    let triggerParams;
                    // 字段变更参数
                    if (this.rule_type === "update") {
                        triggerParams = {
                            field_key: this.field_key,
                            old_value: this.old_value,
                            new_value: this.new_value
                        };
                    }
                    // 定时触发参数
                    if (this.rule_type === "schedule") {
                        triggerParams = {
                            trigger_day: this.triggerDay,
                            trigger_time: this.triggerTime,
                            field_key: this.triggerDateFieldKey
                        };
                    }
                    params = {
                        ...params,
                        ...triggerParams
                    };
                    if (
                        this.operationType === "add" ||
                        this.operationType === "copy"
                    ) {
                        api.ruleCreate(params).then((res) => {
                            if (res.resultCode === 200) {
                                this.$emit("refresh");
                                this.dialogVisible = false;
                            }
                        });
                    }
                    if (this.operationType === "edit") {
                        params.id = this.detail.id;
                        api.ruleUpdate(params).then((res) => {
                            if (res.resultCode === 200) {
                                this.$emit("refresh");
                                this.dialogVisible = false;
                            }
                        });
                    }
                }
            });
        },
        openDialog(type, detail) {
            this.dialogVisible = true;
            this.operationType = type;
            this.detail = detail;
            // 请求规则配置
            this.getRuleConfig(type, detail);
            // 获取流程名
            let id = parseInt(this.curProgress);
            let curProgressInfo = this.curProgressInfo(id);
            this.templateName = curProgressInfo.name || "--";
        },
        // settingApi
        getRuleConfig(type, detail) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress
            };
            api.getRuleConfig(params).then((res) => {
                if (res.resultCode === 200) {
                    // 触发类型
                    this.triggerOptions = res.data.rule_type || [];
                    this.rule_type = this.triggerOptions.length
                        ? this.triggerOptions[0].value
                        : null;

                    // 动作类型
                    this.actionOptions = res.data.rule_action_type || [];
                    // 请求字段列表
                    this.getFilterConfig(type, detail);
                } else {
                    this.triggerOptions = [];
                    this.rule_type = null;
                    this.actionOptions = [];
                }
            });
        },
        getFilterConfig(type, detail) {
            // 触发条件的字段配置
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress
            };
            ProgressApi.getFilterConfig(params).then((res) => {
                if (res.resultCode === 200) {
                    this.fieldList = res.data || [];
                    // 监听字段变更过滤编辑器类型 链接类型
                    let filterFieldList = [];
                    filterFieldList = this.fieldList.filter((item) => {
                        return this.filterFieldType.indexOf(item.mode) > -1;
                    });
                    this.filterFieldList = filterFieldList;
                    // 定时触发的时间/日期字段
                    let dateFieldList = [];
                    dateFieldList = this.fieldList.filter((item) => {
                        return (
                            item.mode === "time_com" || item.mode === "date_com"
                        );
                    });
                    this.dateFieldList = dateFieldList;
                    this.getFieldList(type, detail);
                } else {
                    this.fieldList = [];
                }
            });
        },
        getFieldList(type, detail) {
            // 动作中变更字段配置
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                has_status: "yes"
            };
            settingApi.getFieldList(params).then((res) => {
                if (res.resultCode === 200) {
                    let fieldList = res.data;
                    let actionFieldList = [];
                    actionFieldList = fieldList.filter((item) => {
                        return this.filterFieldType.indexOf(item.mode) > -1;
                    });
                    this.actionFieldList = actionFieldList;
                    if (type === "edit" || type === "copy") {
                        this.reshowDetail(type, detail);
                    }
                } else {
                    this.actionFieldList = [];
                }
            });
        },

        resetForm() {
            this.formInfo = {
                name: "",
                description: ""
            };
            // 触发
            this.triggerOptions = [];
            this.rule_type = null; // 触发类型
            // 日期字段选择
            this.triggerDateFieldKey = ""; // 定时触发的字段
            this.triggerDay = 0;
            this.triggerTime = "08:00";
            this.triggerList = [];
            this.triggerListReshow = [];
            // 必要的动作 分为邮件 和 修改字段
            this.necessaryAction = {
                action_type: "email",
                email_title: "",
                email_contents: "",
                user_list: ""
            };
            // 动作列表
            this.actionList = [];
            this.field_key = ""; // 触发的字段
            this.old_value = null;
            this.new_value = null;
        },
        reshowDetail(type, detail) {
            // 编辑 复制回显
            if (type === "edit") {
                this.formInfo = {
                    name: detail.name,
                    description: detail.description
                };
            }
            // 出发类型
            this.rule_type = detail.rule_type;
            // 字段变更
            if (detail.rule_type === "update") {
                this.field_key = detail.field_key;
                this.old_value = detail.old_value;
                this.oldUserListInfo = detail.old_value_user;
                this.oldRelationListInfo = detail.old_value_documents;
                this.new_value = detail.new_value;
                this.newUserListInfo = detail.new_value_user;
                this.newRelationListInfo = detail.new_value_documents;
            }
            // 定时触发
            if (detail.rule_type === "schedule") {
                this.triggerDateFieldKey = detail.field_key; // 定时触发的字段
                this.triggerDay = detail.trigger_day;
                this.triggerTime = detail.trigger_time;
            }
            // 触发filter
            this.triggerList = detail.filter ? JSON.parse(detail.filter) : [];
            this.triggerListReshow = detail.format_filter
                ? JSON.parse(detail.format_filter)
                : [];
            // 动作 -  拆出来来要动作 和 其他列表
            // this.actionList = detail.actionList ? detail.actionList : [];
            if (detail.action_list) {
                let necessaryAction = detail.action_list.slice(0, 1);
                let actionList = detail.action_list.slice(1);
                this.necessaryAction = necessaryAction[0];
                this.actionList = actionList ? actionList : [];
            }
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
.opertion-item-box {
    display: inline-block;
    width: 18px;
    height: 18px;
    background-size: 100% 100%;
    vertical-align: middle;
    position: relative;
    top: -1px;
    &.rename {
        background-image: url(@/assets/image/progress/add_rule/email.svg);
    }
    &.delete {
        background-image: url(@/assets/image/progress/add_rule/delete.svg);
    }
}
::v-deep.add-rule-dialog.basic-ui {
    .el-dialog {
        .el-dialog__header {
            padding: 18px 32px;
            font-size: 16px;
            color: #2f384c;
        }
        .el-dialog__body {
            max-height: 70vh;
            height: 70vh;
            overflow-y: hidden;
            padding: 0;
            border-top: 2px solid #e6e9f0;
        }
        .el-dialog__footer {
            padding: 20px 32px;
        }
        .el-dialog__headerbtn {
            top: 24px;
            right: 32px;
        }
        .el-form-item {
            // margin-bottom: 24px;
        }
        .el-dialog__footer {
            text-align: right;
            box-shadow: 0px -3px 6px 2px rgba(0, 0, 0, 0.04);
            margin-top: 12px;
        }
    }
}
::v-deep.base-msg-block {
    box-sizing: border-box;
    padding: 14px 32px;
    border-bottom: 1px dotted #cdd5e6;
    background-color: #f5f5f5;
    .base-info {
        .form-row {
            display: flex;
            justify-content: space-between;
            .el-form-item {
                margin-bottom: 0;
                display: inline-block;
                .el-form-item__label {
                    height: 32px;
                    line-height: 32px;
                }
                .el-form-item__content {
                    display: inline-block;
                    height: 32px;
                    line-height: 32px;
                }
            }
            .form-name {
                .el-input {
                    width: 296px;
                }
            }
            .form-desc {
                .el-input {
                    width: 640px;
                }
            }
        }
    }
}
.body-block {
    box-sizing: border-box;
    padding: 16px 96px;
    display: flex;
    justify-content: space-between;
    height: calc(100% - 72px);
    .trigger-block-outter {
        width: calc(50% - 106px);
        position: relative;
        .middle-line {
            width: 0;
            height: 100%;
            border-right: 1px dashed #cdd5e6;
            margin: 0 auto;
            position: absolute;
            top: 0;
            left: 50%;
        }
    }
    .condition-block {
        position: relative;
        width: 100%;
        padding-right: 24px;
        overflow-y: auto;
        height: calc(100% - 100px);
        overflow-x: hidden;
        // .middle-line {
        //     width: 0;
        //     height: 100%;
        //     border-right: 1px dotted #cdd5e6;
        //     margin: 0 auto;
        //     position: absolute;
        //     top: 0;
        //     left: 50%;
        // }
        .trigger-condition {
            position: relative;
            border: 1px solid #e6e9f0;
            // box-shadow: 0px 3px 8px 1px rgba(0, 0, 0, 0.1);
            z-index: 1;
            border-radius: 8px;
            .type-select-border {
                border-radius: 8px;
                padding: 16px;
                background: #fafafb;
                &.had-detail {
                    border-radius: 8px 8px 0 0;
                }
            }
            .detail-block {
                padding: 16px;
                background-color: #ffffff;
                border-radius: 0 0 8px 8px;
                .desc {
                    margin: 16px 0 12px;
                    font-size: 14px;
                    color: #171e31;
                    line-height: 14px;
                    &.end {
                        margin-bottom: 0;
                    }
                }
            }
        }
    }
    .middle-block {
        width: 64px;
        position: relative;
        .middle-btn {
            position: relative;
            top: 18px;
            box-sizing: border-box;
            display: inline-block;
            width: 64px;
            height: 64px;
            background-image: url(@/assets/image/progress/add_rule/middle-icon.svg);
            background-size: 100% 100%;
            border-radius: 8px;
            border: 1px solid #cdd5e6;
            // box-shadow: 0px 3px 8px 1px rgba(0, 0, 0, 0.1);
            z-index: 1;
        }

        .middle-line {
            width: 0;
            height: 100%;
            border-right: 1px dashed #cdd5e6;
            margin: 0 auto;
            position: absolute;
            top: 0;
            left: 50%;
            margin-top: 18px;
            // height: calc(100% + 24px);
        }
    }

    .action-block-outter {
        width: calc(50% - 106px);
        position: relative;
        .middle-line {
            width: 0;
            height: 100%;
            border-right: 1px dashed #cdd5e6;
            margin: 0 auto;
            position: absolute;
            top: 0;
            left: 50%;
        }
    }
    .action-block {
        position: relative;
        width: 100%;
        padding-right: 24px;
        height: calc(100% - 100px);
        overflow-y: auto;
        overflow-x: hidden;
        // .middle-line {
        //     width: 0;
        //     height: 100%;
        //     border-right: 1px solid #cdd5e6;
        //     margin: 0 auto;
        //     position: absolute;
        //     top: 0;
        //     left: 50%;
        // }
        .action-detail {
            position: relative;
            // margin-top: 20px;
            border: 1px solid #e6e9f0;
            // box-shadow: 0px 3px 8px 1px rgba(0, 0, 0, 0.1);
            z-index: 1;
            border-radius: 8px;
            .action-select {
                border-radius: 8px 8px 0 0;
                padding: 16px;
                background: #fafafb;
            }
            .detail-block {
                padding: 16px;
                background-color: #ffffff;
                border-radius: 0 0 8px 8px;
            }
            .delete-icon {
                position: absolute;
                bottom: -9px;
                left: calc(50% - 9px);
                display: inline-block;
                width: 18px;
                height: 18px;
                background-image: url(@/assets/image/progress/add_rule/trigger_delete.svg);
                background-size: 100% 100%;
                cursor: pointer;
            }
        }
    }
}
.scroll-block {
    box-sizing: border-box;
    width: 400px;
    min-height: 100%;
    position: relative;
    left: -6px;
    // padding-bottom: 24px;
    padding: 0 6px 24px;
}
.guide-block {
    position: relative;
    // width: calc(100% - 24px);
    margin-bottom: 20px;
    z-index: 1;
    padding: 16px;
    border: 1px solid #e6e9f0;
    border-radius: 8px;
    background-color: #fafafb;
    // box-shadow: 0px 3px 8px 1px rgba(0, 0, 0, 0.1);
    .title {
        line-height: 32px;
        font-size: 20px;
        color: #171e31;
        margin-bottom: 12px;
        .type-icon {
            box-sizing: border-box;
            display: inline-block;
            width: 32px;
            height: 32px;
            background-size: 100% 100%;
            vertical-align: bottom;
            margin-right: 12px;
            border-radius: 4px;
            border: 1px solid #cdd5e6;
        }
        .trigger-icon {
            background-image: url(@/assets/image/progress/add_rule/trigger_icon.svg);
        }
        .action-icon {
            background-image: url(@/assets/image/progress/add_rule/action_icon.svg);
        }
    }
    .desc {
        padding-left: 48px;
        font-size: 14px;
        color: #a6abbc;
    }
}
.day-and-time {
    display: flex;
    .day-block {
        margin-right: 8px;
    }
    .time-block {
    }
}
::v-deep.day-block {
    height: 32px;
    border: 1px solid #dcdfe6;
    width: fit-content;
    display: flex;
    border-radius: 4px;
    .el-input-number {
        height: 32px;
        line-height: 32px;
        .el-input__inner {
            height: 32px;
            line-height: 32px;
            border: none !important;
        }
        .el-input-number__increase {
            width: 18px;
            line-height: 16px;
            border: 1px solid #dcdfe6;
            right: 0px;
            top: 0;
            border-radius: 0;
        }
        .el-input-number__decrease {
            width: 18px;
            line-height: 16px;
            border: 1px solid #dcdfe6;
            right: 0px;
            bottom: 0;
            border-radius: 0;
        }
    }
    .el-input-number--mini {
        flex: 1;
    }

    .define-append {
        width: 32px;
        height: 32px;
        line-height: 32px;
        display: inline-block;
        background: #f5f7fa;
        border-left: none;
        color: #909399;
        font-size: 12px;
        text-align: center;
    }
}
::v-deep.time-block {
    .el-date-editor {
        // width: 136px;
        width: inherit;
        .el-input__inner {
            height: 32px;
            line-height: 32px;
        }
        .el-input__prefix {
            display: none;
        }
        .el-input__suffix {
            display: none;
        }
    }
}
.list-item {
    margin-top: 16px;
    position: relative;
    z-index: 1;
    padding: 16px;
    border: 1px solid #e6e9f0;
    border-radius: 8px;
    background-color: #ffffff;
    // box-shadow: 0px 3px 8px 1px rgba(0, 0, 0, 0.1);
    .field-row {
        .desc {
            font-size: 14px;
        }
        .el-select {
            margin-left: 104px;
            width: calc(100% - 118px);
        }
        margin-bottom: 8px;
    }
    .relation-value-row {
        height: 32px;
        .el-select.op-select {
            width: 110px;
            vertical-align: top;
        }
        .trigger-value-border {
            display: inline-block;
            margin-left: 8px;
            width: calc(100% - 118px);
            .column-block {
                width: 100%;
            }
            .el-input {
                width: 100%;
            }
            .el-select {
                width: 100%;
            }
            .el-date-editor {
                width: 100%;
            }
        }
    }
    .delete-icon {
        position: absolute;
        bottom: -9px;
        left: calc(50% - 9px);
        display: inline-block;
        width: 18px;
        height: 18px;
        background-image: url(@/assets/image/progress/add_rule/trigger_delete.svg);
        background-size: 100% 100%;
        cursor: pointer;
    }
}
.field-select.el-select {
    box-sizing: border-box;
    width: 100%;
}
.action-item {
    .desc {
        height: 14px;
        line-height: 14px;
        margin: 16px 0 12px;
        &.if {
            margin-top: 0;
        }
    }
}
.add-trigger-block {
    display: inline-block;
    box-sizing: border-box;
    margin: 24px auto;
    width: 40px;
    height: 40px;
    padding: 12px;
    position: relative;
    left: calc(50% - 20px);
    z-index: 1;
    border-radius: 8px;
    background-color: #ffffff;
    // box-shadow: 0px 3px 8px 1px rgba(0, 0, 0, 0.1);
    border: 1px solid #e6e9f0;
    cursor: pointer;
    .add-trigger-icon {
        display: inline-block;
        width: 16px;
        height: 16px;
        background-image: url(@/assets/image/progress/add_rule/trigger_add.svg);
        background-size: 100% 100%;
    }
}
.add-action-dropdown {
    box-sizing: border-box;
    margin: 24px auto;
    left: calc(50% - 20px);
    .add-action-block {
        box-sizing: border-box;
        width: 40px;
        height: 40px;
        padding: 12px;
        position: relative;
        z-index: 1;
        border-radius: 8px;
        background-color: #ffffff;
        // box-shadow: 0px 3px 8px 1px rgba(0, 0, 0, 0.1);
        border: 1px solid #e6e9f0;
        cursor: pointer;
        .add-action-icon {
            display: inline-block;
            width: 16px;
            height: 16px;
            background-image: url(@/assets/image/progress/add_rule/trigger_add.svg);
            background-size: 100% 100%;
        }
    }
}
.msg-transform {
    margin: 16px 0 12px;
    .desc {
        height: 14px;
        font-size: 14px;
        color: #171e31;
        .time {
            margin-left: 8px;
            font-size: 14px;
            color: #959595;
            line-height: 14px;
        }
    }
}
::v-deep .number-input.el-input {
    &.number.hidden-arrow {
        input::-webkit-outer-spin-button,
        input::-webkit-inner-spin-button {
            -webkit-appearance: none;
        }
        input[type="number"] {
            -moz-appearance: textfield;
        }
    }
}
.type-box {
    display: inline-block;
    width: 18px;
    height: 18px;
    background-size: 100% 100%;
    vertical-align: middle;
    position: relative;
    top: -2px;
    &.update {
        background-image: url(@/assets/image/progress/add_rule/update.svg);
    }
    &.schedule {
        background-image: url(@/assets/image/progress/add_rule/schedule.svg);
    }
    &.create {
        background-image: url(@/assets/image/progress/add_rule/create.svg);
    }
    &.delete {
        background-image: url(@/assets/image/progress/add_rule/delete.svg);
    }
    &.email {
        background-image: url(@/assets/image/progress/add_rule/email.svg);
    }
}
</style>
