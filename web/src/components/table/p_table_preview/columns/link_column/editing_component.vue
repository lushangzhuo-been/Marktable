<template>
    <div class="column-block" ref="ColumnBlock">
        <div class="detail">
            <span
                v-if="linkObj.name || linkObj.url"
                class="could-jump"
                @click.stop="jump(scope)"
            >
                {{ linkObj.name || linkObj.url }}
            </span>
            <span v-else>
                {{ emptySpace() }}
            </span>
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
            edit: false,
            popoverWidth: 220,
            linkObj: {
                name: "",
                url: ""
            }
        };
    },
    watch: {
        scope: {
            handler(scope) {
                // 分情况 接口有返回此字段/无返回此字段
                if (
                    scope.row[this.item.field_key] &&
                    Object.values(scope.row[this.item.field_key]).length
                ) {
                    this.linkObj = _.cloneDeep(scope.row[this.item.field_key]);
                } else {
                    this.linkObj = {
                        name: "",
                        url: ""
                    };
                }
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
                this.popoverWidth = this.$refs.ColumnBlock.clientWidth;
                this.$nextTick(() => {
                    setTimeout(() => {
                        this.$refs.DropPopover.doShow();
                    }, 20);
                });
            } else {
                this.afterLeave();
            }
        },
        afterLeave() {
            this.edit = false;
            this.isEditing = false;
            // 校验
            if (
                this.item.required === "yes" &&
                this.scope.row[this.item.field_key] &&
                (this.scope.row[this.item.field_key].name ||
                    this.scope.row[this.item.field_key].url) &&
                !this.linkObj.name &&
                !this.linkObj.url
            ) {
                this.$message({
                    showClose: true,
                    message: "此为必填项",
                    type: "warning"
                });
                this.linkObj = this.scope.row[this.item.field_key];
            } else if (
                this.item.required === "yes" &&
                (!this.scope.row[this.item.field_key] ||
                    (!this.scope.row[this.item.field_key].name &&
                        !this.scope.row[this.item.field_key].url)) &&
                !this.linkObj.name &&
                !this.linkObj.url
            ) {
                // 不提交
            } else if (
                JSON.stringify(this.scope.row[this.item.field_key]) ==
                JSON.stringify(this.linkObj)
            ) {
                // 不提交
            } else {
                this.$emit(
                    "edit-form-item",
                    JSON.stringify(this.linkObj),
                    this.item.field_key,
                    this.scope.row._id,
                    this.item.mode
                );
            }
        },
        copyLink(scope) {
            let tempTextarea = document.createElement("textarea");
            // 设置textarea的value为当前网址
            tempTextarea.value = scope.row[this.item.field_key].url;
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
        jump(scope) {
            if (scope.row[this.item.field_key].url) {
                let url = scope.row[this.item.field_key].url;
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
.column-block {
    box-sizing: border-box;
    height: 40px;
    line-height: 40px;
    border: 1px solid rgba(0, 0, 0, 0);
    padding: 0 10px;
    position: relative;
    // &.modify {
    //     &:hover {
    //         .triangle {
    //             transform: rotate(0deg);
    //             transition-duration: 0.3s;
    //             position: absolute;
    //             right: 5px;
    //             top: 13px;
    //             display: inline-block;
    //             width: 16px;
    //             height: 16px;
    //             background-size: 100% 100%;
    //             background-image: url("~@/assets/image/common/triangle-down.png");
    //         }
    //     }
    //     &.isEditing {
    //         border: 1px solid var(--PRIMARY_COLOR);
    //         .triangle {
    //             transform: rotate(180deg);
    //             transition-duration: 0.3s;
    //             position: absolute;
    //             right: 5px;
    //             top: 13px;
    //             display: inline-block;
    //             width: 16px;
    //             height: 16px;
    //             background-size: 100% 100%;
    //             background-image: url("~@/assets/image/common/triangle-down-active.png");
    //         }
    //     }
    // }
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

.could-jump {
    color: var(--PRIMARY_COLOR);
    cursor: pointer;

    &:hover {
        text-decoration: underline;
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
