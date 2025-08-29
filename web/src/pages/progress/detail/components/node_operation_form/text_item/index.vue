<template>
    <div class="column-block">
        <el-input
            class="progress-input"
            :class="{
                'validate-failed': validateFailed
            }"
            ref="main-text-input"
            v-model.trim="text"
            placeholder="请输入内容"
            @blur="blur"
            @input="inputChange"
        ></el-input>
        <div v-show="validateFailed" class="validate-desc">此项为必填</div>
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
            isEditing: true,
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
        checkScope(scope) {
            this.isEditing = !this.isEditing;
            this.$nextTick(() => {
                this.$refs["main-text-input"].focus();
            });
        },
        blur() {
            this.isEditing = false;
            this.$emit(
                "edit-form-item",
                this.text,
                this.formItem.field_key,
                this.formItem.mode
            );
            this.doValidate();
        },
        inputChange() {
            this.doValidate();
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
    line-height: 32px;
    border: 1px solid rgba(0, 0, 0, 0);
    padding: 0;
    position: relative;
    .validate-desc {
        color: #f56c6c;
        font-size: 12px;
        height: 12px;
        line-height: 12px;
    }
}

.detail {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    height: 32px;
}
::v-deep .progress-input.el-input {
    .el-input__inner {
        padding: 0 10px;
    }
    &.validate-failed {
        .el-input__inner {
            border: 1px solid #fd5050;
        }
    }
}
</style>
