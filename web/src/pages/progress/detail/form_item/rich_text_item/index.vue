<template>
    <div
        class="column-block input-area"
        ref="column-block"
        :class="{
            isEditing: isEditing,
            'editing-effect': formItem.can_modify === 'yes',
            'col-required': formItem.required === 'yes'
        }"
    >
        <div v-if="!isEditing" class="detail" @click="checkScope">
            <tip-more :text="text">
                <div class="detail-content">
                    {{ emptySpace(text) }}
                </div>
            </tip-more>
        </div>
        <div v-if="isEditing">
            <el-input
                class="basic-ui progress-input"
                ref="main-text-input"
                v-model.trim="text"
                type="textarea"
                :rows="3"
                @blur="blur"
                placeholder="请输入内容"
            ></el-input>
            <!-- <el-popover
                ref="rich-text-popover"
                popper-class="progress-popover"
                placement="bottom"
                :width="popoverWidth"
                trigger="click"
                @after-leave="afterLeave"
            >
                <div class="popover-content">
                    <el-input
                        class="basic-ui progress-rich-text"
                        ref="RichInput"
                        type="textarea"
                        :rows="2"
                        placeholder="请输入内容"
                        v-model="text"
                    >
                    </el-input>
                </div>
                <div slot="reference" class="detail">
                    {{ text }}
                </div>
            </el-popover> -->
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import { emptySpace } from "@/assets/tool/func";
import TipMore from "@/pages/common/tooltip_more_line.vue";
import api from "@/common/api/module/progress";
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
            popoverWidth: 220,
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
                this.text =
                    _.cloneDeep(formData[this.formItem.field_key]) || "";
            }
        }
    },
    mounted() {},
    methods: {
        emptySpace(param) {
            return emptySpace(param);
        },
        checkScope() {
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
        },
        afterLeave() {
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
            } else if (this.formData[this.formItem.field_key] == this.text) {
                // 为空或者无变化
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
.avatar-box {
    display: inline-block;
    min-width: 24px;
    height: 24px;
    border-radius: 50%;
    background-size: 100% 100%;
    vertical-align: middle;
}
.detail {
    padding: 0 10px;
    .detail-content {
        box-sizing: border-box;
        min-height: 32px;
        line-height: 1.5;
        padding: 4px 0;
        display: -webkit-box;
        -webkit-box-orient: vertical;
        overflow: hidden;
        text-overflow: ellipsis;
        -webkit-line-clamp: 3;
    }
}
::v-deep .basic-ui.el-textarea.progress-input {
    .el-textarea__inner {
        font-size: 14px;
        padding: 4px 9px;
        border: none;
        &:focus {
            border: none;
            // border: 1px solid #dcdfe6;
            &:hover {
                // background-color: #e9f0f8;
                background-color: rgba(0, 0, 0, 0);
            }
        }
    }
}
</style>
<style lang="scss">
// .basic-ui.progress-rich-text.el-textarea {
//     box-sizing: border-box;
//     .el-textarea__inner {
//         border-color: rgba(0, 0, 0, 0);
//         padding: 16px;
//         &:focus {
//             border-color: rgba(0, 0, 0, 0);
//         }
//     }
// }
</style>
