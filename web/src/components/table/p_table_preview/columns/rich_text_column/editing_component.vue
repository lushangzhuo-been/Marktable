<template>
    <div class="column-block" ref="column-block">
        <div v-overflow class="detail">
            {{ text || emptySpace() }}
        </div>
    </div>
</template>

<script>
import _ from "lodash";
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
            text: "",
            popoverWidth: 380
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
        checkScope() {
            if (this.item.can_modify === "no") {
                return;
            }
            this.isEditing = !this.isEditing;
            if (this.isEditing) {
                // this.popoverWidth = this.$refs["column-block"].clientWidth;
                this.$nextTick(() => {
                    setTimeout(() => {
                        this.$refs["rich-text-popover"].doShow();
                        this.$nextTick(() => {
                            this.$refs.RichInput.focus();
                        });
                    }, 20);
                });
            } else {
                this.afterLeave();
            }
        },
        afterLeave() {
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
    position: relative;
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

.detail {
    height: 40px;
    padding: 0 10px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}
</style>
