<template>
    <div
        class="column-block"
        ref="ColumnBlock"
        :class="{
            isEditing: isEditing,
            'editing-effect': formItem.can_modify === 'yes',
            'col-required': formItem.required === 'yes'
        }"
    >
        <div v-if="!isEditing" slot="reference" @click="checkScope()">
            <div class="detail" slot="reference">
                <template
                    v-if="
                        formData.status && Object.keys(formData.status).length
                    "
                >
                    <tip-one :text="formData.status.name">
                        <div
                            class="status"
                            :style="{
                                backgroundColor: rgbToRgba(
                                    `${formData.status.color || '#fff'}`,
                                    0.3
                                ),
                                color: '#171e31'
                            }"
                        >
                            {{ formData.status.name || emptySpace() }}
                        </div>
                    </tip-one>
                </template>
                <span v-else style="padding: 0 10px">
                    {{ emptySpace() }}
                </span>
            </div>
        </div>
        <div v-if="isEditing">
            <el-popover
                ref="DropPopover"
                popper-class="col-multi-select-popover progress-popover"
                placement="bottom"
                :width="popoverWidth"
                trigger="click"
                @after-leave="afterLeave"
                @show="popperShow"
            >
                <div class="popover-content multi-select-col">
                    <el-input
                        ref="SearchInput"
                        class="basic-ui progress-input"
                        v-model="searchValue"
                        placeholder="搜索"
                        prefix-icon="el-icon-search"
                    ></el-input>
                    <div v-if="optionsList.length" class="options-box">
                        <div
                            class="drapdown-item"
                            v-for="(item, index) in optionsList"
                            :class="item.step_auth ? '' : 'disabled'"
                            :key="index"
                            :style="{
                                color: '#171e17',
                                backgroundColor: rgbToRgba(
                                    `${item.end_status_color || '#fff'}`,
                                    0.3
                                )
                            }"
                            @click="switchStatus(item)"
                        >
                            <tip-one :text="item.end_status_name">
                                <div class="font">
                                    {{ item.end_status_name }}
                                </div>
                            </tip-one>
                        </div>
                    </div>
                    <div v-else class="no-data-text">暂无数据</div>
                </div>

                <!-- 标签化 -->
                <div class="detail" slot="reference">
                    <template
                        v-if="
                            formData.status &&
                            Object.keys(formData.status).length
                        "
                    >
                        <tip-one :text="formData.status.name">
                            <div
                                class="status"
                                :style="{
                                    backgroundColor: rgbToRgba(
                                        `${formData.status.color || '#fff'}`,
                                        0.3
                                    ),
                                    color: '#171e31'
                                }"
                            >
                                {{ formData.status.name || emptySpace() }}
                            </div>
                        </tip-one>
                    </template>
                    <span v-else style="padding: 0 10px">
                        {{ emptySpace() }}
                    </span>
                </div>
            </el-popover>
        </div>
        <!-- 切换状态时有必填字段，显示弹窗 -->
        <node-operation-dialog
            ref="NodeOperationDialog"
            @confirm-next-state="confirmNextState"
            :nodeOperationConfig="nodeOperationConfig"
            :nodeOperationData="nodeOperationData"
        ></node-operation-dialog>
    </div>
</template>

<script>
import _ from "lodash";
import { rgbToRgba } from "@/assets/tool/func";
import TipOne from "@/pages/common/tooltip_one_line.vue";
import { emptySpace } from "@/assets/tool/func";
import api from "@/common/api/module/progress";
import NodeOperationDialog from "@/pages/progress/detail/components/node_operation_dialog";
export default {
    props: {
        formItem: {
            type: Object,
            default: () => {}
        },
        formData: {
            type: Object,
            default: () => {}
        }
    },
    components: {
        TipOne,
        NodeOperationDialog
    },
    data() {
        return {
            isEditing: false,
            searchValue: "",
            vmodelArr: [],
            selectArr: [],
            frontArr: [],
            behandArr: [],
            popoverWidth: 220,
            optionsList: [],
            tmpOptionsList: [],
            nodeOperationConfig: [],
            nodeOperationData: {}
        };
    },
    computed: {
        wsId() {
            return this.$store.state.curSpace.id || "";
        },
        tempId() {
            return this.$route.params.id || "";
        }
    },
    watch: {
        searchValue(str) {
            if (str.trim()) {
                this.optionsList = this.tmpOptionsList.filter((item) => {
                    return (
                        item.end_status_name
                            .toLocaleUpperCase()
                            .indexOf(str.toLocaleUpperCase()) > -1
                    );
                });
            } else {
                this.optionsList = this.tmpOptionsList;
            }
        },
        formData: {
            handler() {
                let values = this.formData[this.formItem.field_key];
                if (values && values.length) {
                    this.vmodelArr = _.cloneDeep(values);
                    let selectArr = [];
                    values.forEach((value) => {
                        let a = 0;
                        this.formItem.enum_values.forEach((item) => {
                            if (value === item.value) {
                                selectArr.push(item);
                                a++;
                            }
                        });
                        if (a === 0) {
                            selectArr.push({
                                value: value,
                                color: "#ccc"
                            });
                        }
                    });
                    this.selectArr = selectArr;
                } else {
                    this.selectArr = [];
                }
            },
            immediate: true,
            deep: true
        }
    },
    mounted() {},
    methods: {
        rgbToRgba(color, opacity) {
            return rgbToRgba(color, opacity);
        },
        // 打开popover时调用状态下拉选项
        popperShow() {
            this.getStatusList();
        },
        getStatusList() {
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId,
                id: this.formData._id
            };
            api.getStatusMenusList(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    res.data.length
                ) {
                    this.optionsList = res.data;
                    this.tmpOptionsList = _.cloneDeep(this.optionsList);
                } else {
                    this.optionsList = [];
                    this.tmpOptionsList = [];
                }
            });
        },
        emptySpace(param) {
            return emptySpace(param);
        },
        hexToRgba(hex, opacity) {
            return (
                "rgba(" +
                parseInt("0x" + hex.slice(1, 3)) +
                "," +
                parseInt("0x" + hex.slice(3, 5)) +
                "," +
                parseInt("0x" + hex.slice(5, 7)) +
                "," +
                opacity +
                ")"
            );
        },
        removeItem() {
            this.selectArr = [];
            this.optionsList = this.getOptionsList();
        },
        getOptionsList() {
            if (
                this.selectArr &&
                this.selectArr.length &&
                this.formItem.enum_values
            ) {
                return this.formItem.enum_values.filter((v) =>
                    this.selectArr.every((val) => val.value !== v.value)
                );
            } else {
                return this.formItem.enum_values || [];
            }
        },
        checkScope() {
            this.searchValue = "";
            this.fetAuthEdit();
        },
        fetAuthEdit() {
            // 获取进展权限
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId,
                id: this.formData._id,
                auth_mode: "edit",
                field_key: this.formItem.field_key
            };
            api.getUserAuth(params).then((res) => {
                if (res && res.resultCode === 200) {
                    if (res.data) {
                        this.isEditing = !this.isEditing;
                        if (this.isEditing) {
                            this.popoverWidth =
                                this.$refs.ColumnBlock.clientWidth;
                            setTimeout(() => {
                                this.$refs.DropPopover.doShow();
                                this.$nextTick(() => {
                                    if (this.$refs.SearchInput) {
                                        this.$refs.SearchInput.focus();
                                    }
                                });
                            }, 20);
                        } else {
                            this.afterLeave();
                        }
                    } else {
                        this.isEditing = false;
                    }
                } else {
                    this.isEditing = false;
                }
            });
        },
        getArrFront(arr) {
            let deepClone = _.cloneDeep(arr);
            let front = deepClone.splice(0, 2);
            return front;
        },
        getArrBehand(arr) {
            let deepClone = _.cloneDeep(arr);
            let behand = deepClone.splice(2);
            return behand;
        },
        switchStatus(item) {
            // 点击状态选项后，调如下接口，判断是否有限制字段需要填写
            if (!item.step_auth) return;
            this.hasLimitFiled(item);
        },
        hasLimitFiled(item) {
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId,
                step_id: item.id,
                id: this.formData._id
            };
            api.getBtnScreen(params).then((res) => {
                if (res && res.resultCode === 200) {
                    if (!res.data || (res.data && !res.data.length)) {
                        // 如果没有限制字段，直接放行， 设置状态为选中的，并且刷新表格内容
                        this.switchStepAction(item);
                    } else {
                        // res.data为获取的字段
                        this.nodeOperationConfig = res.data;
                        this.$refs.NodeOperationDialog.openDialog(item);
                        document.body.click(); //关闭poppover
                        // 获取节点配置数据
                        this.getNodeOperationData(item);
                    }
                }
            });
        },
        // 切换状态赋值
        switchStepAction(item) {
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId,
                step_id: item.id,
                id: this.formData._id
            };
            api.switchStepAction(params).then((res) => {
                if (res && res.resultCode === 200) {
                    document.body.click();
                    this.$emit(
                        "edit-form-item",
                        JSON.stringify(this.selectArr),
                        this.formItem.field_key,
                        this.formItem.mode
                    );
                    this.$emit("status-change-info", item);
                }
            });
        },
        // 获取弹窗节点配置数据
        getNodeOperationData(item) {
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId,
                step_id: item.id,
                id: this.formData._id
            };
            api.getFormData(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    this.nodeOperationData = res.data;
                } else {
                    this.nodeOperationData = {};
                }
            });
        },
        afterLeave() {
            this.isEditing = false;
        },
        // 弹窗跳出后填写完毕，提交
        confirmNextState(btn, form) {
            let params = {
                ws_id: this.wsId,
                tmpl_id: this.tempId,
                step_id: btn.id,
                id: this.formData._id,
                // 解构被编辑的参数
                ...form
            };
            api.switchStepAction(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.$emit(
                        "edit-form-item",
                        JSON.stringify(this.selectArr),
                        this.formItem.field_key,
                        this.formItem.mode
                    );
                    this.$emit("status-change-info", btn);
                    this.$refs.NodeOperationDialog.closeDialog();
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
@import "../../style.scss";
.detail {
    box-sizing: border-box;
    display: flex;
    align-items: center;
    height: 32px;
    padding: 0 16px 0 10px;
    border-radius: 4px;
    white-space: nowrap;
    min-width: 160px;
    max-width: 280px;
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
    .tag-list {
        display: inline-block;
        height: 32px;
        .tag-item {
            box-sizing: border-box;
            display: inline-block;
            height: 24px;
            line-height: 24px;
            text-align: center;
            margin: 4px 8px 4px 0;
            padding: 0 4px;
            border-radius: 4px;
            text-overflow: ellipsis;
            overflow: hidden;
            white-space: nowrap;
            max-width: 100%;
        }
    }
}

.avatar-box {
    display: inline-block;
    min-width: 24px;
    height: 24px;
    border-radius: 50%;
    background-size: 100% 100%;
    vertical-align: middle;
}

.num-box {
    box-sizing: border-box;
    display: inline-block;
    min-width: 22px;
    height: 22px;
    line-height: 22px;
    border-radius: 11px;
    padding: 0 2px;
    background-color: #f8f8f8;
    text-align: center;
    position: relative;
    vertical-align: middle;
    font-weight: 400;
    font-size: 12px;
    color: #2f384c;
}
</style>
<style lang="scss">
// tooltip数字提示内容
.el-tooltip__popper.col-mutil-select-tooltip {
    .tip-item {
        display: inline-block;
        font-size: 12px;
        padding: 0 4px;
        border-radius: 3px;
        height: 20px;
        box-sizing: border-box;
        line-height: 21px;
        margin: 0px 4px 3px 0px;
    }
}
// popover下拉框
.el-popover.col-multi-select-popover {
    padding: 4px 0;
    .progress-input.el-input {
        .el-input__inner {
            padding: 0 16px 0 32px;
            border-radius: 0;
            border: none;
            border-bottom: 1px solid #f0f1f5;
            height: 38px;
            line-height: 38px;
            background-color: rgba(0, 0, 0, 0);
        }
    }
    .selected-popover {
        padding: 8px 8px 2px;
        .selected-item {
            font-size: 12px;
            display: inline-block;
            padding: 0 4px;
            border-radius: 3px;
            box-sizing: border-box;
            max-width: 100%;
            background-color: rgb(245, 246, 248);
            border: 1px solid rgb(205, 213, 230);
            line-height: 1;
            margin: 0px 4px 3px 0px;
        }
        .value {
            display: inline-block;
            text-overflow: ellipsis;
            white-space: nowrap;
            overflow: hidden;
            width: calc(100% - 18px);
            height: 18px;
            line-height: 21px;
            position: relative;
            top: -1px;
        }
        .el-icon-error {
            position: relative;
            top: -4px;
            margin-left: 6px;
            cursor: pointer;
            color: #cdd5e6;
        }
    }

    .options-box {
        margin: 8px 0 4px 8px;
        padding-right: 8px;
        max-height: 210px;
        overflow-y: auto;
        .drapdown-item {
            box-sizing: border-box;
            padding: 0 10px;
            line-height: 32px;
            height: 32px;
            text-align: center;
            border-radius: 4px;
            cursor: pointer;
            margin-bottom: 4px;
            &.disabled {
                opacity: 0.5;
                cursor: not-allowed;
            }
            .font {
                width: 100%;
                text-overflow: ellipsis;
                overflow: hidden;
                white-space: nowrap;
            }
        }
    }
    .no-data-text {
        padding: 24px 0;
        text-align: center;
        color: #909399;
    }
}
</style>
