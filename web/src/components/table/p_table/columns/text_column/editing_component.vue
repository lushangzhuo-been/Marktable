<template>
    <div
        class="column-block"
        :class="{
            isEditing: isEditing,
            'editing-effect': item.can_modify === 'yes',
            'col-required': item.required === 'yes'
        }"
    >
        <!-- :class="{
            isEditing: isEditing,
            'editing-effect':
                item.can_modify === 'yes',
            'col-required': item.required === 'yes'
        }" -->
        <div
            v-if="!isEditing"
            v-overflow
            class="detail"
            @click="checkScope(scope)"
        >
            {{ text || emptySpace() }}
        </div>
        <div v-if="isEditing">
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
import { emptySpace } from "@/assets/tool/func";
import api from "@/common/api/module/progress";
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
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        }
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
            this.fetAuthEdit();
        },
        fetAuthEdit() {
            // 获取进展权限
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: this.scope.row._id,
                auth_mode: "edit",
                field_key: this.item.field_key
            };
            api.getUserAuth(params).then((res) => {
                if (res && res.resultCode === 200) {
                    if (res.data) {
                        this.isEditing = true;
                        this.$nextTick(() => {
                            this.$refs["main-text-input"].focus();
                        });
                    } else {
                        this.isEditing = false;
                    }
                } else {
                    this.isEditing = false;
                }
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
    &.editing-effect {
        &:before {
            transition-duration: 0.3s;
            content: "";
            display: block;
            position: absolute;
            right: 0;
            top: 0;
            width: 0px;
            height: 0px;
            background: linear-gradient(
                42deg,
                #cdd5e6 30%,
                rgba(0, 0, 0, 0) 40%
            );
        }
        &:hover {
            &:before {
                transition-duration: 0.3s;
                content: "";
                display: block;
                position: absolute;
                right: 0;
                top: 0;
                width: 12px;
                height: 12px;
                background: linear-gradient(
                    42deg,
                    #cdd5e6 30%,
                    rgba(0, 0, 0, 0) 40%
                );
            }
        }
        &.isEditing {
            border: 1px solid var(--PRIMARY_COLOR);
            &:before {
                transition-duration: 0.3s;
                content: "";
                display: block;
                position: absolute;
                right: 0;
                top: 0;
                width: 12px;
                height: 12px;
                background: linear-gradient(
                    42deg,
                    #cdd5e6 30%,
                    rgba(0, 0, 0, 0) 40%
                );
            }
        }
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
}
</style>
