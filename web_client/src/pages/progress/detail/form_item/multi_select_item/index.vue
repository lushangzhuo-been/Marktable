<template>
    <div
        class="column-block form-width"
        ref="ColumnBlock"
        :class="{
            isEditing: isEditing,
            'editing-effect': item.can_modify === 'yes',
            'col-required': item.required === 'yes'
        }"
    >
        <!-- 展示(非编辑) -->
        <div class="detail" v-if="!isEditing" @click="checkScope">
            <!-- 名字列表 -->
            <template v-if="selectArr.length">
                <tip-one :text="selectArr[0].value" v-show="!behandArr.length">
                    <span
                        class="first-tag"
                        :style="{
                            color: '#fff',
                            backgroundColor: selectArr[0].color,
                            maxWidth:
                                selectArr.length === 1
                                    ? 'calc(100% - 24px)'
                                    : 'calc(100% - 30px - 24px)',
                            marginRight: selectArr.length === 1 ? 0 : '8px'
                        }"
                    >
                        {{ selectArr[0].value }}
                    </span>
                </tip-one>
                <span
                    v-show="behandArr.length"
                    class="first-tag"
                    :style="{
                        color: '#fff',
                        backgroundColor: selectArr[0].color,
                        maxWidth:
                            selectArr.length === 1
                                ? 'calc(100% - 24px)'
                                : 'calc(100% - 30px - 24px)',
                        marginRight: selectArr.length === 1 ? 0 : '8px'
                    }"
                >
                    {{ selectArr[0].value }}
                </span>
                <!-- 数字气泡 -->
                <el-tooltip
                    class="item"
                    effect="dark"
                    placement="top"
                    popper-class="col-mutil-select-tooltip"
                >
                    <div slot="content">
                        <div v-for="(item, index) in behandArr" :key="index">
                            <span
                                class="tip-item"
                                :style="{
                                    color: '#fff',
                                    backgroundColor: item.color
                                }"
                            >
                                {{ item.value }}
                            </span>
                        </div>
                    </div>
                    <b v-if="behandArr.length" class="num-box"
                        >+{{ behandArr.length }}</b
                    >
                </el-tooltip>
            </template>
            <span v-else>
                {{ emptySpace() }}
            </span>
        </div>
        <!--编辑 -->
        <div v-if="isEditing">
            <el-popover
                ref="DropPopover"
                popper-class="col-multi-select-popover progress-popover"
                placement="bottom-start"
                :width="popoverWidth"
                trigger="click"
                @after-leave="afterLeave"
            >
                <!-- 遍历item.list  做下拉内容 -->
                <el-input
                    class="basic-ui progress-input"
                    ref="SearchInput"
                    v-model="searchValue"
                    placeholder="搜索"
                    prefix-icon="el-icon-search"
                ></el-input>
                <div class="selected-popover" v-if="selectArr.length">
                    <span
                        v-for="(item, index) in selectArr"
                        :key="index"
                        class="selected-item"
                    >
                        <span class="value"> {{ item.value }}</span>
                        <b
                            class="el-icon-error"
                            @click="removeItem(item, index)"
                        ></b>
                    </span>
                </div>
                <div v-if="optionsList.length" class="options-box">
                    <div
                        class="drapdown-item"
                        v-for="(item, index) in optionsList"
                        :key="index"
                        :style="{
                            color: '#fff',
                            backgroundColor: item.color
                        }"
                        @click="checkboxChange(item)"
                    >
                        {{ item.value }}
                    </div>
                </div>
                <div v-else class="no-data-text">暂无数据</div>
                <!-- 标签化 -->
                <div class="detail" slot="reference">
                    <!-- 名字列表 -->
                    <template v-if="selectArr.length">
                        <span
                            class="first-tag"
                            :style="{
                                color: '#fff',
                                backgroundColor: selectArr[0].color,
                                maxWidth:
                                    selectArr.length === 1
                                        ? 'calc(100% - 24px)'
                                        : 'calc(100% - 30px - 24px)',
                                marginRight: selectArr.length === 1 ? 0 : '8px'
                            }"
                        >
                            {{ selectArr[0].value }}
                        </span>
                        <!-- 数字气泡 -->
                        <el-tooltip
                            class="item"
                            effect="dark"
                            placement="top"
                            popper-class="col-mutil-select-tooltip"
                        >
                            <div slot="content">
                                <div
                                    v-for="(item, index) in behandArr"
                                    :key="index"
                                >
                                    <span
                                        class="tip-item"
                                        :style="{
                                            color: '#fff',
                                            backgroundColor: item.color
                                        }"
                                    >
                                        {{ item.value }}
                                    </span>
                                </div>
                            </div>
                            <b v-if="behandArr.length" class="num-box"
                                >+{{ behandArr.length }}</b
                            >
                        </el-tooltip>
                    </template>
                    <span v-else></span>
                </div>
            </el-popover>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import TipOne from "@/pages/common/tooltip_one_line.vue";
import { emptySpace } from "@/assets/tool/func";
export default {
    props: {
        item: {
            type: Object,
            default: () => {}
        },
        formData: {
            type: Object,
            default: () => {}
        }
    },
    components: {
        TipOne
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
            optionsList: []
        };
    },
    watch: {
        searchValue(str) {
            if (str.trim()) {
                this.optionsList = _.cloneDeep(
                    this.getOptionsList().filter((item) => {
                        return (
                            item.value
                                .toLocaleUpperCase()
                                .indexOf(str.toLocaleUpperCase()) > -1
                        );
                    })
                );
            } else {
                this.optionsList = this.getOptionsList();
            }
        },
        formData: {
            handler(formData) {
                let values = formData[this.item.field_key];
                if (values && values.length) {
                    this.vmodelArr = _.cloneDeep(values);
                    let selectArr = [];
                    // 双重循环拿到选项的颜色
                    if (this.item.enum_values && this.item.enum_values.length) {
                        values.forEach((value, index) => {
                            // selectArr.push(
                            //     _.find(this.item.enum_values, { value: value })
                            // );
                            let a = 0;
                            this.item.enum_values.forEach((item) => {
                                if (value === item.value) {
                                    selectArr.push(item);
                                    a++;
                                }
                            });
                            if (a === 0) {
                                selectArr.push({
                                    value: value,
                                    color: "#F0F1F5"
                                });
                            }
                        });
                    }
                    this.selectArr = selectArr;
                    // this.frontArr = this.getArrFront(this.selectArr);
                    if (this.selectArr.length > 1) {
                        this.behandArr = this.getArrBehand(this.selectArr);
                    } else {
                        this.behandArr = [];
                    }
                } else {
                    this.selectArr = [];
                }
                this.optionsList = this.getOptionsList();
            },
            immediate: true,
            deep: true
        }
    },
    mounted() {},
    methods: {
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

        rgbToRgba(color, alp) {
            let rgbaAttr = color.match(/[\d.]+/g);
            if (rgbaAttr.length >= 3) {
                let r, g, b;
                r = rgbaAttr[0];
                g = rgbaAttr[1];
                b = rgbaAttr[2];
                return "rgba(" + r + "," + g + "," + b + "," + alp + ")";
            }
        },
        removeItem(item, index) {
            this.selectArr.splice(index, 1);
            this.optionsList = this.getOptionsList();
            // this.frontArr = this.getArrFront(this.selectArr);
            if (this.selectArr.length > 1) {
                this.behandArr = this.getArrBehand(this.selectArr);
            } else {
                this.behandArr = [];
            }
        },
        getOptionsList() {
            let tmpList = [];
            if (
                this.selectArr &&
                this.selectArr.length &&
                this.item.enum_values
            ) {
                tmpList = this.item.enum_values.filter((v) =>
                    this.selectArr.every((val) => val.value !== v.value)
                );
            } else {
                tmpList = this.item.enum_values || [];
            }
            if (this.searchValue.trim()) {
                return _.cloneDeep(
                    tmpList.filter((item) => {
                        return (
                            item.value
                                .toLocaleUpperCase()
                                .indexOf(
                                    this.searchValue.trim().toLocaleUpperCase()
                                ) > -1
                        );
                    })
                );
            } else {
                return tmpList;
            }
        },
        checkScope() {
            this.searchValue = "";
            if (this.item.can_modify === "no") {
                return;
            }
            this.isEditing = !this.isEditing;
            this.$set(this.item, "isEditing", this.isEditing);
            if (this.isEditing) {
                // this.popoverWidth = this.$refs.ColumnBlock.clientWidth;
                this.popoverWidth = 220;
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
        },
        afterLeave() {
            this.isEditing = false;
            this.$set(this.item, "isEditing", this.isEditing);
            let validateSelectArr = [];
            if (this.selectArr && this.selectArr.length) {
                this.selectArr.forEach((item) => {
                    validateSelectArr.push(item.value);
                });
            }
            // 校验
            if (
                this.item.required === "yes" &&
                !this.selectArr.length &&
                this.vmodelArr.length
            ) {
                // 前端回显
                this.$message({
                    showClose: true,
                    message: "此为必填项",
                    type: "warning"
                });
                this.vmodelArr = _.cloneDeep(
                    this.formData[this.item.field_key]
                );
                let selectArr = [];
                this.formData[this.item.field_key].forEach((value) => {
                    let a = 0;
                    this.item.enum_values.forEach((item) => {
                        if (value === item.value) {
                            selectArr.push(item);
                            a++;
                        }
                    });
                    if (a === 0) {
                        selectArr.push({
                            value: value,
                            color: "#F0F1F5"
                        });
                    }
                });
                this.selectArr = selectArr;
                if (this.selectArr.length > 1) {
                    this.behandArr = this.getArrBehand(this.selectArr);
                } else {
                    this.behandArr = [];
                }
                this.optionsList = this.getOptionsList();
            } else if (
                JSON.stringify(validateSelectArr) ==
                JSON.stringify(this.vmodelArr)
            ) {
                // 前后皆为空不提交
            } else {
                let selectArr = [];
                if (this.selectArr) {
                    this.selectArr.forEach((item) => {
                        selectArr.push(item.value);
                    });
                }
                this.$emit(
                    "edit-form-item",
                    JSON.stringify(selectArr),
                    this.item.field_key,
                    this.formData._id,
                    this.item.mode
                );
            }
        },
        getArrFront(arr) {
            let deepClone = _.cloneDeep(arr);
            let front = deepClone.splice(0, 2);
            return front;
        },
        getArrBehand(arr) {
            let deepClone = _.cloneDeep(arr);
            let behand = deepClone.splice(1);
            return behand;
        },
        checkboxChange(item) {
            this.selectArr.push(item);
            this.optionsList = this.getOptionsList();
            if (this.selectArr.length > 1) {
                this.behandArr = this.getArrBehand(this.selectArr);
            } else {
                this.behandArr = [];
            }
        }
    }
};
</script>

<style lang="scss" scoped>
@import "../../style.scss";
.detail {
    display: flex;
    align-items: center;
    height: 32px;
    padding: 0 10px;
    border-radius: 4px;
    white-space: nowrap;
    .first-tag {
        display: inline-block;
        height: 24px;
        line-height: 24px;
        text-align: center;
        margin-right: 8px;
        padding: 0 4px;
        max-width: 62px;
        border-radius: 4px;
        text-overflow: ellipsis;
        overflow: hidden;
        white-space: nowrap;
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
            border: 1px solid rgb(205, 28, 230);
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
            width: 100%;
            box-sizing: border-box;
            padding: 0 10px;
            line-height: 32px;
            height: 32px;
            text-align: center;
            border-radius: 4px;
            text-overflow: ellipsis;
            overflow: hidden;
            white-space: nowrap;
            cursor: pointer;
            margin-bottom: 4px;
        }
    }
    .no-data-text {
        padding: 24px 0;
        text-align: center;
        color: #909399;
    }
}
</style>
