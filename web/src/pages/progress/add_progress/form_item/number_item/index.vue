<template>
    <div class="column-block">
        <el-input
            class="basic-ui progress-input number hidden-arrow"
            :class="{
                'validate-failed': validateFailed
            }"
            ref="main-text-input"
            type="number"
            v-model.trim="text"
            placeholder="请输入内容"
            @input="inputChange"
            @blur="blur"
        ></el-input>
        <div v-show="validateFailed" class="validate-desc">此项为必填</div>
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
            isEditing: true,
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
        checkScope(scope) {
            this.isEditing = !this.isEditing;
            this.$nextTick(() => {
                this.$refs["main-text-input"].focus();
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
.column-block {
    box-sizing: border-box;
    height: 32px;
    line-height: 32px;
    border: 1px solid rgba(0, 0, 0, 0);
    padding: 0;
    .validate-desc {
        color: #f56c6c;
        font-size: 12px;
        height: 12px;
        line-height: 12px;
        position: absolute;
        bottom: -16px;
    }
}
::v-deep .progress-input.el-input {
    .el-input__inner {
        padding: 0 10px;
    }
    &.number.hidden-arrow {
        input::-webkit-outer-spin-button,
        input::-webkit-inner-spin-button {
            -webkit-appearance: none;
        }
        input[type="number"] {
            -moz-appearance: textfield;
        }
    }
    &.validate-failed {
        .el-input__inner {
            border: 1px solid #fd5050;
        }
    }
}
</style>
