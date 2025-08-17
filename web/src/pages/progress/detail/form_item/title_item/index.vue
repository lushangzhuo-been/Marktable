<template>
    <div
        class="column-block detail-title"
        :class="{
            isEditing: isEditing,
            'editing-effect': formItem.can_modify === 'yes',
            'col-required': formItem.required === 'yes'
        }"
    >
        <div v-if="!isEditing" class="detail" @click="checkScope">
            <div v-overflow class="detail-content">
                {{ emptySpace(text) }}
            </div>
        </div>
        <div class="detail" v-if="isEditing">
            <el-input
                class="basic-ui progress-input detail-title"
                ref="main-text-input"
                v-model.trim="text"
                @blur="blur"
                placeholder="请输入内容"
            ></el-input>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import { emptySpace } from "@/assets/tool/func";
import api from "@/common/api/module/progress";
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
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        }
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
            this.fetAuthEdit();
        },
        fetAuthEdit() {
            // 获取进展权限
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: this.formData._id,
                auth_mode: "edit",
                field_key: this.formItem.field_key
            };
            api.getUserAuth(params).then((res) => {
                if (res && res.resultCode === 200) {
                    if (res.data) {
                        this.isEditing = !this.isEditing;
                        this.$set(this.formItem, "isEditing", this.isEditing);
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
    padding: 0 16px 0 10px;
    height: 32px;
    .detail-content {
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }
}

::v-deep .progress-input.el-input {
    .el-input__inner {
        border: none;
        padding: 0;
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
