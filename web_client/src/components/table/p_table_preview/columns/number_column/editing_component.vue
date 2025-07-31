<template>
    <div class="column-block">
        <div class="detail">
            {{ kSeq(text) }}
        </div>
    </div>
</template>

<script>
// E:\mark3\mark3_client\src\assets\tool\func.js
import { kSeq } from "@/assets/tool/func";
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
            deep: true,
            immediate: true
        }
    },
    mounted() {},
    methods: {
        kSeq(num) {
            return kSeq(num);
        },
        checkScope(scope) {
            if (this.item.can_modify === "no") {
                return;
            }
            this.isEditing = true;
            this.$nextTick(() => {
                this.$refs["main-text-input"].focus();
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
.column-block {
    box-sizing: border-box;
    height: 40px;
    line-height: 40px;
    border: 1px solid rgba(0, 0, 0, 0);
    padding: 0 10px;
    &.isEditing {
        border: 1px solid var(--PRIMARY_COLOR);
    }
}

.detail {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    height: 40px;
}

::v-deep .progress-input.el-input {
    .el-input__inner {
        padding: 0;
        border: none;
        height: 38px;
        line-height: 38px;
        background-color: rgba(0, 0, 0, 0);
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
