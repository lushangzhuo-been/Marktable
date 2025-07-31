<template>
    <div
        class="column-block"
        ref="ColumnBlock"
        :class="{ isEditing: isEditing }"
    >
        <div>
            <el-popover
                ref="DropPopover"
                popper-class="progress-popover col-multi-select-popover"
                placement="bottom"
                :width="popoverWidth"
                trigger="click"
                @show="popoverShow"
                @after-leave="afterLeave"
            >
                <!-- 遍历item.list  做下拉内容 -->

                <div class="popover-content multi-select-col">
                    <el-input
                        ref="SearchInput"
                        class="basic-ui progress-input"
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
                            <b class="el-icon-error" @click="removeItem()"></b>
                        </span>
                    </div>
                    <div v-if="optionsList.length" class="options-box">
                        <div
                            class="drapdown-item"
                            v-for="(item, index) in optionsList"
                            :key="index"
                            :style="{
                                color: '#fff',
                                backgroundColor: item.color,
                            }"
                            @click="checkboxChange(item)"
                        >
                            {{ item.value }}
                        </div>
                    </div>
                    <div v-else class="no-data-text">暂无数据</div>
                </div>

                <!-- 标签化 -->
                <div slot="reference">
                    <!-- 名字列表 -->
                    <div
                        v-if="selectArr.length"
                        class="detail"
                        :class="{
                            'validate-failed': validateFailed,
                        }"
                    >
                        <span
                            class="first-tag"
                            :style="{
                                color: '#fff',
                                backgroundColor: selectArr[0].color,
                            }"
                        >
                            <span class="value">{{ selectArr[0].value }}</span>
                        </span>
                    </div>
                    <div
                        v-else
                        class="detail default-text"
                        :class="{
                            'validate-failed': validateFailed,
                        }"
                    >
                        请选择
                    </div>
                    <b class="triangle"></b>
                </div>
            </el-popover>
        </div>
        <div v-show="validateFailed" class="validate-desc">此项为必填</div>
    </div>
</template>

<script>
import _ from "lodash";
export default {
    props: {
        formItem: {
            type: Object,
            default: () => {},
        },
        formData: {
            type: Object,
            default: () => {},
        },
        validateOrder: {
            type: Number,
            default: 0,
        },
    },
    data() {
        return {
            validateFailed: false,
            isEditing: false,
            searchValue: "",
            vmodelArr: [],
            selectArr: [],
            frontArr: [],
            behandArr: [],
            popoverWidth: 220,
            optionsList: [],
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
            handler(data) {
                let values = data[this.formItem.field_key];
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
                                color: "#ccc",
                            });
                        }
                    });
                    this.selectArr = selectArr;
                } else {
                    this.selectArr = [];
                }
                this.optionsList = this.getOptionsList();
            },
            immediate: true,
            deep: true,
        },
        validateOrder: {
            handler() {
                this.doValidate();
            },
        },
    },
    mounted() {},
    methods: {
        doValidate() {
            this.curFormItem = this.formItem;
            if (this.curFormItem.required === "yes" && !this.selectArr.length) {
                this.validateFailed = true;
                this.curFormItem.validateFailed = true;
            } else {
                this.validateFailed = false;
                this.curFormItem.validateFailed = false;
            }
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
        removeItem() {
            this.selectArr = [];
            if (this.formItem.required === "yes" && !this.selectArr.length) {
                this.validateFailed = true;
            } else {
                this.validateFailed = false;
            }
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
        popoverShow() {
            this.searchValue = "";
            this.isEditing = true;
            this.popoverWidth = this.$refs.ColumnBlock.clientWidth;
            this.$nextTick(() => {
                if (this.$refs.SearchInput) {
                    this.$refs.SearchInput.focus();
                }
            });
        },
        afterLeave() {
            this.isEditing = false;
            let selectArr = [];
            if (this.selectArr) {
                this.selectArr.forEach((item) => {
                    selectArr.push(item.value);
                });
            }
            this.$emit(
                "edit-form-item",
                JSON.stringify(selectArr),
                this.formItem.field_key,
                this.formItem.mode
            );
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
        checkboxChange(item) {
            this.selectArr = [item];
            if (this.formItem.required === "yes" && !this.selectArr.length) {
                this.validateFailed = true;
            } else {
                this.validateFailed = false;
            }
            this.optionsList = this.getOptionsList();
            this.$refs.DropPopover.doClose();
        },
    },
};
</script>

<style lang="scss" scoped>
@import "../style.scss";
.detail {
    display: flex;
    align-items: center;
    border: 1px solid #dcdfe6;
    padding: 0 10px;
    white-space: nowrap;
    height: 32px;
    border-radius: 4px;
    &:focus {
        border: 1px solid var(--PRIMARY_COLOR);
    }
    .first-tag {
        display: inline-block;
        max-width: calc(100% - 24px);
        height: 24px;
        line-height: 24px;
        text-align: center;
        margin-right: 8px;
        padding: 0 4px;
        border-radius: 4px;
        text-overflow: ellipsis;
        overflow: hidden;
        white-space: nowrap;
    }
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
