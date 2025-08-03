<template>
    <div
        class="column-block form-width"
        :class="{
            isEditing: isEditing,
            'editing-effect': formItem.can_modify === 'yes',
            'col-required': formItem.required === 'yes'
        }"
    >
        <div class="detail" v-if="!isEditing" @click="checkScope">
            {{ emptySpace(text) }}
        </div>
        <div v-if="isEditing">
            <el-date-picker
                ref="DatePicker"
                class="progress-date-picker"
                popper-class="progress-date-picker-popover"
                v-model="text"
                type="datetime"
                placeholder="选择日期"
                @blur="blur"
                value-format="yyyy-MM-dd HH:mm:ss"
            >
            </el-date-picker>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import { emptySpace } from "@/assets/tool/func";
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
        emptySpace(param) {
            return emptySpace(param);
        },
        checkScope() {
            if (this.formItem.can_modify === "no") {
                return;
            }
            this.isEditing = true;
            this.$set(this.formItem, "isEditing", this.isEditing);
            this.$nextTick(() => {
                this.$refs["DatePicker"].focus();
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
                this.$message({
                    showClose: true,
                    message: "此为必填项",
                    type: "warning"
                });
                this.text = this.formData[this.formItem.field_key];
            } else if (this.formData[this.formItem.field_key] == this.text) {
                //
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
}
::v-deep .el-date-editor.progress-date-picker {
    width: 100%;
    border: none;

    .el-input__inner {
        background-color: rgba(0, 0, 0, 0);
        // height: 38px;
        // line-height: 38px;
        height: 30px;
        line-height: 30px;
        border: none;
        padding: 0 10px;
        font-weight: 400;
        font-size: 14px;
        color: #171e31;
        font-family: MiSans, MiSans;
        &:hover {
            border: none;
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
