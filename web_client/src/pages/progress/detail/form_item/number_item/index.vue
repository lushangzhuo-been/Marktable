<template>
    <div
        class="column-block form-width"
        :class="{
            isEditing: isEditing,
            'editing-effect': formItem.can_modify === 'yes',
            'col-required': formItem.required === 'yes'
        }"
    >
        <div v-if="!isEditing" class="detail" @click="checkScope()">
            {{ kSeq(text) }}
        </div>
        <div v-if="isEditing">
            <el-input
                class="basic-ui progress-input number hidden-arrow"
                ref="main-text-input"
                v-model.trim="text"
                type="number"
                placeholder="请输入内容"
                @blur="blur"
            ></el-input>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import { kSeq } from "@/assets/tool/func";
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
    data() {
        return {
            isEditing: false,
            text: ""
        };
    },
    watch: {
        formData: {
            handler(formData) {
                this.text = _.cloneDeep(formData[this.formItem.field_key]);
            }
        }
    },
    mounted() {},
    methods: {
        kSeq(num) {
            return kSeq(num);
        },
        checkScope(scope) {
            if (this.formItem.can_modify === "no") {
                return;
            }
            this.isEditing = !this.isEditing;
            this.$set(this.formItem, "isEditing", this.isEditing);
            this.$nextTick(() => {
                this.$refs["main-text-input"].focus();
            });
        },
        blur() {
            this.isEditing = false;
            this.$set(this.formItem, "isEditing", this.isEditing);
            if (
                this.formItem.required === "yes" &&
                this.formData[this.formItem.field_key] &&
                !this.text
            ) {
                // 必填以前有值被清空
                this.$message({
                    showClose: true,
                    message: "此为必填项",
                    type: "warning"
                });
                this.text = this.formData[this.formItem.field_key];
            } else if (
                this.formData[this.formItem.field_key] == this.text ||
                (!this.formData[this.formItem.field_key] && !this.text)
            ) {
                // 必填前后均不为空 或者
            } else {
                this.$emit(
                    "edit-form-item",
                    this.text,
                    this.formItem.field_key,
                    this.formItem.mode
                );
            }
        }
    }
};
</script>

<style lang="scss" scoped>
// @import "../../style.scss";
.column-block {
    box-sizing: border-box;
    height: 32px;
    line-height: 32px;
    border: 1px solid rgba(0, 0, 0, 0);
    padding: 0 10px;
    border-radius: 4px;
    position: relative;
    &:hover {
        background: #f5f6f8;
        border-radius: 4px 4px 4px 4px;
    }
    &.isEditing {
        border: 1px solid var(--PRIMARY_COLOR);
        background: #ffffff;
        box-shadow: 0px 4px 16px 1px rgba(0, 0, 0, 0.1);
        border-radius: 4px;
        border: 1px solid #e6e9f0;
    }
    &.form-width {
        width: 60%;
        min-width: 220px;
    }
}

.detail {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    // height: 40px;
    height: 32px;
}

::v-deep .progress-input.el-input {
    .el-input__inner {
        padding: 0;
        border: none;
        // height: 38px;
        // line-height: 38px;
        height: 30px;
        line-height: 30px;
        background-color: rgba(0, 0, 0, 0);
        font-weight: 400;
        font-size: 14px;
        color: #171e31;
        font-family: MiSans, MiSans;
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
}
</style>
