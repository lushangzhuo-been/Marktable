<template>
    <div
        class="column-block"
        :class="{
            isEditing: isEditing,
            'editing-effect': formItem.can_modify === 'yes',
            'col-required': formItem.required === 'yes'
        }"
    >
        <div v-if="!isEditing" class="detail input-area" @click="checkScope">
            <div v-overflow class="detail-content">
                {{ emptySpace(text) }}
            </div>
        </div>
        <div class="detail" v-if="isEditing">
            <el-input
                class="basic-ui progress-input"
                ref="main-text-input"
                v-model.trim="text"
                placeholder="请输入内容"
                @blur="blur"
            ></el-input>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import { emptySpace } from "@/assets/tool/func";
import TipMore from "@/pages/common/tooltip_more_line.vue";
export default {
    components: {
        TipMore
    },
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
        emptySpace(param) {
            return emptySpace(param);
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
            // 校验
            if (
                this.formItem.required === "yes" &&
                this.formData[this.formItem.field_key] &&
                !this.text
            ) {
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
                // 不提交
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
@import "../../style.scss";
.detail {
    height: 32px;
    padding: 0 10px;
    .detail-content {
        box-sizing: border-box;
        height: 32px;
        line-height: 32px;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
    }
}
::v-deep .basic-ui.el-textarea.progress-input {
    .el-input__inner {
        padding: 0;
        border: none;
        height: 30px;
        line-height: 30px;
        background-color: rgba(0, 0, 0, 0);
        font-weight: 400;
        font-size: 14px;
        color: #171e31;
        font-family: MiSans, MiSans;
    }
    .el-textarea__inner {
        font-size: 14px;
        padding: 4px 9px;
        border: rgba(0, 0, 0, 0);
        &:focus {
            border: 1px solid #dcdfe6;
        }
    }
}
::v-deep .basic-ui.progress-input {
    height: 32px;
    .el-input__inner {
        padding: 0;
        border: rgba(0, 0, 0, 0);
        height: 32px;
        line-height: 32px;
        background-color: rgba(0, 0, 0, 0);
        font-weight: 400;
        font-size: 14px;
        color: #171e31;
        font-family: MiSans, MiSans;
    }
}
</style>
