<template>
    <div
        class="column-block"
        ref="column-block"
        :class="{ isEditing: isEditing }"
    >
        <el-popover
            ref="rich-text-popover"
            popper-class="progress-popover"
            placement="bottom"
            :width="popoverWidth"
            trigger="click"
            @show="popoverShow"
            @after-leave="afterLeave"
        >
            <div class="popover-content">
                <el-input
                    class="basic-ui progress-rich-text"
                    :class="{
                        'validate-failed': validateFailed
                    }"
                    ref="RichInput"
                    type="textarea"
                    :rows="2"
                    placeholder="请输入内容"
                    v-model.trim="text"
                    @input="inputChange"
                >
                </el-input>
            </div>
            <div slot="reference">
                <div
                    v-if="text"
                    class="detail"
                    :class="{
                        'validate-failed': validateFailed
                    }"
                >
                    {{ text }}
                </div>
                <div
                    class="detail default-text"
                    :class="{
                        'validate-failed': validateFailed
                    }"
                    v-else
                >
                    请输入
                </div>
                <b class="triangle"></b>
            </div>
        </el-popover>
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
            isEditing: false,
            text: "",
            validateFailed: false,
            popoverWidth: 220
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
        inputChange() {
            this.doValidate();
        },
        popoverShow() {
            this.isEditing = true;
            this.popoverWidth = this.$refs["column-block"].clientWidth;
            this.$nextTick(() => {
                this.$refs.RichInput.focus();
                this.doValidate();
            });
        },
        afterLeave() {
            this.isEditing = false;
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
.default-text {
    color: #c0c4cc;
}

::v-deep .progress-input.el-input {
    .el-input__inner {
        padding: 0;
        border: none;
        height: 39px;
        line-height: 39px;
        background-color: rgba(0, 0, 0, 0);
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
</style>
<style lang="scss">
.basic-ui.progress-rich-text.el-textarea {
    box-sizing: border-box;
    .el-textarea__inner {
        border-color: rgba(0, 0, 0, 0);
        padding: 16px;
        &:focus {
            border-color: rgba(0, 0, 0, 0);
        }
    }
}
</style>
