<template>
    <div
        class="column-block form-width"
        ref="ColumnBlock"
        :class="{
            isEditing: isEditing,
            'editing-effect': formItem.can_modify === 'yes',
            'col-required': formItem.required === 'yes'
        }"
    >
        <div v-if="!isEditing" class="detail" @click="checkScope">
            <span
                v-if="linkObj.name || linkObj.url"
                class="could-jump"
                @click.stop="jump(formData)"
            >
                {{ linkObj.name || linkObj.url }}
            </span>
            <span v-else>
                {{ emptySpace() }}
            </span>
        </div>
        <div v-if="isEditing">
            <el-popover
                ref="DropPopover"
                popper-class="progress-popover padding12"
                placement="bottom-start"
                :width="popoverWidth"
                trigger="click"
                @after-leave="afterLeave"
            >
                <div class="popover-content">
                    <!-- 三个操作 与 编辑面板 -->
                    <div v-show="!edit" class="operation-list">
                        <div class="operation-item" @click="copyLink(formData)">
                            <b class="operation-item-box copy"></b>
                            复制网址
                        </div>
                        <div class="operation-item" @click="editLink">
                            <b class="operation-item-box edit"></b>
                            编辑
                        </div>
                        <div class="operation-item" @click="deleteLink">
                            <b class="operation-item-box delete"></b>
                            删除
                        </div>
                    </div>
                    <div v-show="edit">
                        <div class="edit-item">
                            <div class="title">
                                <b class="operation-item-box web"></b>
                                网址
                            </div>
                            <el-input
                                ref="UrlInput"
                                class="basic-ui"
                                v-model.trim="linkObj.url"
                                placeholder="请输入网址"
                            ></el-input>
                        </div>
                        <div class="edit-item">
                            <div class="title">
                                <b class="operation-item-box name"></b>
                                名称
                            </div>
                            <el-input
                                class="basic-ui"
                                v-model.trim="linkObj.name"
                                placeholder="请输入名称"
                            ></el-input>
                        </div>
                    </div>
                </div>
                <div slot="reference" class="detail">
                    <span class="could-jump" @click.stop="jumpLink(linkObj)">
                        {{ linkObj.name || linkObj.url }}
                    </span>
                </div>
            </el-popover>
        </div>
        <b class="triangle" @click="checkScope"></b>
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
            edit: false,
            popoverWidth: 220,
            linkObj: {
                name: "",
                url: ""
            }
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
                // 分情况 接口有返回此字段/无返回此字段
                if (
                    formData[this.formItem.field_key] &&
                    Object.values(formData[this.formItem.field_key]).length
                ) {
                    this.linkObj = _.cloneDeep(
                        formData[this.formItem.field_key]
                    );
                } else {
                    this.linkObj = {
                        name: "",
                        url: ""
                    };
                }
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
                        if (this.isEditing) {
                            // this.popoverWidth = this.$refs.ColumnBlock.clientWidth;
                            this.popoverWidth = 220;
                            this.$nextTick(() => {
                                setTimeout(() => {
                                    this.$refs.DropPopover.doShow();
                                }, 20);
                            });
                        } else {
                            this.afterLeave();
                        }
                    } else {
                        this.isEditing = false;
                    }
                } else {
                    this.isEditing = false;
                }
            });
        },
        afterLeave() {
            this.edit = false;
            this.isEditing = false;
            this.$set(this.formItem, "isEditing", this.isEditing);
            // 校验
            if (
                this.formItem.required === "yes" &&
                this.formData[this.formItem.field_key] &&
                this.formData[this.formItem.field_key].name &&
                this.formData[this.formItem.field_key].url &&
                (!this.linkObj.name || !this.linkObj.url)
            ) {
                // 必填 新值两个字段都为空
                this.$message({
                    showClose: true,
                    message: "此为必填项",
                    type: "warning"
                });
                this.linkObj = _.cloneDeep(
                    this.formData[this.formItem.field_key]
                );
            } else if (
                (!this.formData[this.formItem.field_key] &&
                    !this.linkObj.name &&
                    !this.linkObj.url) ||
                (this.formData[this.formItem.field_key] &&
                    !this.formData[this.formItem.field_key].name &&
                    !this.formData[this.formItem.field_key].url &&
                    !this.linkObj.name &&
                    !this.linkObj.url) ||
                JSON.stringify(this.formData[this.formItem.field_key]) ==
                    JSON.stringify(this.linkObj)
            ) {
                // 新旧都为空或相同不提交
            } else {
                this.$emit(
                    "edit-form-item",
                    JSON.stringify(this.linkObj),
                    this.formItem.field_key,
                    this.formItem.mode
                );
            }
        },
        copyLink() {
            let tempTextarea = document.createElement("textarea");
            if (this.formData[this.formItem.field_key].url) {
                // 设置textarea的value为当前网址
                tempTextarea.value = this.formData[this.formItem.field_key].url;
                // 将textarea添加到DOM中
                document.body.appendChild(tempTextarea);
                // 选中textarea中的文本
                tempTextarea.select();
                // 复制选中的文本
                document.execCommand("copy");
                // 移除临时的textarea元素
                document.body.removeChild(tempTextarea);
                this.$message({
                    showClose: true,
                    message: "复制成功",
                    type: "success"
                });
                this.$refs.DropPopover.doClose();
            } else {
                this.$message({
                    showClose: true,
                    message: "无链接可以复制",
                    type: "warning"
                });
            }
        },
        editLink() {
            this.edit = true;
            this.$nextTick(() => {
                if (this.$refs.UrlInput) {
                    this.$refs.UrlInput.focus();
                }
            });
        },
        deleteLink() {
            this.$refs.DropPopover.doClose();
        },
        jump(formData) {
            if (formData[this.formItem.field_key].url) {
                let url = formData[this.formItem.field_key].url;
                const encodedUrl = encodeURI(url);
                // 打开新页面
                let reg = /http(s)?:\/\/([\w-]+\.)+[\w-]+(\/[\w- .\/?%&=]*)?/;
                if (reg.test(encodedUrl)) {
                    window.open(encodedUrl, "_blank");
                } else {
                    window.open("http://" + encodedUrl, "_blank");
                }
            } else {
                this.$message({
                    showClose: true,
                    message: "复制成功",
                    type: "warning"
                });
            }
        },
        jumpLink(obj) {
            if (obj.url) {
                let url = obj.url;
                const encodedUrl = encodeURI(url);
                // 打开新页面
                let reg = /http(s)?:\/\/([\w-]+\.)+[\w-]+(\/[\w- .\/?%&=]*)?/;
                if (reg.test(encodedUrl)) {
                    window.open(encodedUrl, "_blank");
                } else {
                    window.open("http://" + encodedUrl, "_blank");
                }
            } else {
                this.$message({
                    showClose: true,
                    message: "链接地址未填写",
                    type: "warning"
                });
            }
        }
    }
};
</script>

<style lang="scss" scoped>
@import "../../style.scss";
.could-jump {
    color: var(--PRIMARY_COLOR);
    cursor: pointer;

    &:hover {
        text-decoration: underline;
    }
}
.detail {
    height: 32px;
    padding: 0 10px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}

.rich-text-popover {
    padding: 0;
}
.operation-item-box {
    display: inline-block;
    width: 20px;
    height: 20px;
    background-size: 100% 100%;
    vertical-align: middle;
    &.copy {
        background-image: url(@/assets/image/progress_column/copy.svg);
    }
    &.edit {
        background-image: url(@/assets/image/common/edit.svg);
    }
    &.delete {
        background-image: url(@/assets/image/common/delete.svg);
    }
    &.web {
        background-image: url(@/assets/image/progress_column/web.svg);
    }
    &.name {
        background-image: url(@/assets/image/progress_column/name.svg);
    }
}
.edit-item {
    margin-bottom: 12px;
    &:last-child {
        margin-bottom: 0px;
    }
    .title {
        height: 20px;
        line-height: 20px;
        margin-bottom: 8px;
    }
}
</style>
<style lang="scss">
.operation-list {
    .operation-item {
        height: 40px;
        line-height: 40px;
        padding: 0 8px;

        &:hover {
            background-color: #f1f9ff;
        }
    }
}
</style>
