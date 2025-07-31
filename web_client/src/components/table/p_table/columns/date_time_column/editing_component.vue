<template>
    <div
        class="column-block"
        :class="{
            isEditing: isEditing,
            'editing-effect': item.can_modify === 'yes',
            'col-required': item.required === 'yes'
        }"
    >
        <div
            v-overflow
            class="detail"
            v-if="!isEditing"
            @click="checkScope(scope)"
        >
            {{ text || emptySpace() }}
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
import { emptySpace } from "@/assets/tool/func";
export default {
    props: {
        item: {
            type: Object,
            default: () => {}
        },
        scope: {
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
        scope: {
            handler(scope) {
                this.text = scope.row[this.item.field_key] || "";
            },
            immediate: true
        }
    },
    mounted() {},
    methods: {
        emptySpace(param) {
            return emptySpace(param);
        },
        checkScope(scope) {
            if (this.item.can_modify === "no") {
                return;
            }
            this.isEditing = !this.isEditing;
            this.$nextTick(() => {
                this.$refs["DatePicker"].focus();
            });
        },
        blur() {
            this.isEditing = false;
            if (
                this.item.required === "yes" &&
                this.scope.row[this.item.field_key] &&
                !this.text
            ) {
                this.$message({
                    showClose: true,
                    message: "此为必填项",
                    type: "warning"
                });
                this.text = this.scope.row[this.item.field_key];
            } else if (
                this.scope.row[this.item.field_key] == this.text ||
                (!this.scope.row[this.item.field_key] && !this.text)
            ) {
                //
            } else {
                this.$emit(
                    "edit-form-item",
                    this.text,
                    this.item.field_key,
                    this.scope.row._id,
                    this.item.mode
                );
            }
        }
    }
};
</script>

<style lang="scss" scoped>
@import "../../style.scss";
.column-block {
    box-sizing: border-box;
    height: 40px;
    line-height: 40px;
    border: 1px solid rgba(0, 0, 0, 0);
    &.isEditing {
        border: 1px solid var(--PRIMARY_COLOR);
    }
}

.detail {
    height: 40px;
    padding: 0 10px;
    font-weight: 400;
    font-size: 14px;
    color: #171e31;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}

::v-deep .el-date-editor.progress-date-picker {
    width: 100%;
    border: none;

    .el-input__inner {
        background-color: rgba(0, 0, 0, 0);
        height: 38px;
        line-height: 38px;
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
