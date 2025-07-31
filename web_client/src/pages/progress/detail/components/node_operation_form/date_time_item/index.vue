<template>
    <div class="column-block" :class="{ isEditing: isEditing }">
        <el-date-picker
            ref="DatePicker"
            class="add-progress-date-picker"
            :class="{
                'validate-failed': validateFailed
            }"
            popper-class="progress-date-picker-popover"
            v-model.trim="text"
            type="datetime"
            placeholder="选择日期"
            @blur="blur"
            value-format="yyyy-MM-dd HH:mm:ss"
        >
        </el-date-picker>
        <div v-show="validateFailed" class="validate-desc">此项为必填</div>
        <b class="time-box"></b>
    </div>
</template>

<script>
import _ from "lodash";
export default {
    props: {
        formItem: {
            type: Object,
            default: () => {}
        },
        formData: {
            type: Object,
            default: () => {}
        },
        validateOrder: {
            type: Number,
            default: 0
        }
    },
    data() {
        return {
            isEditing: false,
            text: "",
            validateFailed: false
        };
    },
    watch: {
        formData: {
            handler(formData) {
                this.text =
                    _.cloneDeep(formData[this.formItem.field_key]) || "";
                this.$emit(
                    "edit-form-item",
                    this.text,
                    this.formItem.field_key,
                    this.formItem.mode
                );
            },
            immediate: true
        },
        validateOrder: {
            handler() {
                this.doValidate();
            }
        }
    },
    mounted() {},
    methods: {
        doValidate() {
            if (this.formItem.required === "yes" && !this.text) {
                this.validateFailed = true;
                this.formItem.validateFailed = true;
            } else {
                this.validateFailed = false;
                this.formItem.validateFailed = false;
            }
        },
        checkScope() {
            this.isEditing = true;
            this.$nextTick(() => {
                this.$refs["DatePicker"].focus();
            });
        },
        blur() {
            this.isEditing = false;
            this.doValidate();
            this.$emit(
                "edit-form-item",
                this.text,
                this.formItem.field_key,
                this.formItem.mode
            );
        }
    }
};
</script>

<style lang="scss" scoped>
@import "../style.scss";
::v-deep .el-date-editor.add-progress-date-picker {
    width: 100%;
    .el-input__inner {
        background-color: rgba(0, 0, 0, 0);
        height: 32px;
        line-height: 32px;
        padding: 0 10px;
    }

    .el-input__prefix {
        display: none;
    }

    .el-input__suffix {
        display: none;
    }
    &.validate-failed {
        .el-input__inner {
            border: 1px solid #fd5050;
        }
    }
}
</style>
