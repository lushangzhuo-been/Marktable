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
            type="date"
            placeholder="选择日期"
            @input="inputChange"
            @blur="blur"
            value-format="yyyy-MM-dd"
        >
        </el-date-picker>
        <div v-show="validateFailed" class="validate-desc">此项为必填</div>
        <b class="time-box"></b>
    </div>
</template>

<script>
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
                this.text = formData[this.formItem.prop];
            },
            immediate: true
        },
        validateOrder: {
            handler(order) {
                this.doValidate();
            }
        }
    },
    mounted() {},
    methods: {
        checkScope() {
            this.isEditing = !this.isEditing;
            this.$nextTick(() => {
                this.$refs["DatePicker"].focus();
            });
        },
        inputChange() {
            this.doValidate();
        },
        blur() {
            this.doValidate();
            this.isEditing = false;
            this.formData[this.formItem.field_key] = this.text;
        },
        doValidate() {
            if (this.formItem.required === "yes" && !this.text) {
                this.validateFailed = true;
                this.formItem.validateFailed = true;
            } else {
                this.validateFailed = false;
                this.formItem.validateFailed = false;
            }
        }
    }
};
</script>

<style lang="scss" scoped>
@import "../../style.scss";
::v-deep .el-date-editor.add-progress-date-picker {
    width: 100%;
    .el-input__inner {
        background-color: rgba(0, 0, 0, 0);
        height: 32px;
        line-height: 32px;
        padding: 0 10px;
    }
    &.validate-failed {
        .el-input__inner {
            border: 1px solid #fd5050;
        }
    }
    .el-input__prefix {
        display: none;
    }

    .el-input__suffix {
        display: none;
    }
}
</style>
