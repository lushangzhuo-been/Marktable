<template>
    <div
        class="column-block"
        ref="column-block"
        :class="{ isEditing: isEditing }"
    >
        <el-popover
            ref="popover"
            popper-class="progress-popover padding12"
            placement="bottom"
            :width="popoverWidth"
            trigger="click"
            @show="popoverShow"
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
                            v-model="linkObj.url"
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
                            v-model="linkObj.name"
                            placeholder="请输入名称"
                        ></el-input>
                    </div>
                </div>
            </div>
            <div slot="reference">
                <div
                    v-if="linkObj.url || linkObj.name"
                    class="detail"
                    :class="{
                        'validate-failed': validateFailed
                    }"
                >
                    <span class="could-jump" @click.stop="jump(formData)">
                        {{ linkObj.name || linkObj.url }}
                    </span>
                </div>
                <div
                    class="detail default-text"
                    :class="{
                        'validate-failed': validateFailed
                    }"
                    v-else
                >
                    请编辑
                </div>
                <b class="triangle"></b>
            </div>
        </el-popover>
        <div v-show="validateFailed" class="validate-desc">此项为必填</div>
    </div>
</template>

<script>
import _ from "lodash";
export default {
    props: {
        formItem: {
            type: Object,
            default: () => {}
        },
        formData: {
            type: Object,
            default: () => {}
        },
        validateOrder: {
            handler(order) {
                this.doValidate();
            }
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
            },
            validateFailed: false
        };
    },
    watch: {
        formData: {
            handler(formData) {
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
                this.$emit(
                    "edit-form-item",
                    JSON.stringify(this.linkObj),
                    this.formItem.field_key,
                    this.formItem.mode
                );
            },
            immediate: true
        },
        validateOrder: {
            handler(order) {
                this.doValidate();
            }
        }
    },
    mounted() {},
    methods: {
        popoverShow() {
            this.isEditing = true;
            this.popoverWidth = this.$refs["column-block"].clientWidth;
        },
        doValidate() {
            if (
                this.formItem.required === "yes" &&
                !this.linkObj.url &&
                !this.linkObj.name
            ) {
                this.validateFailed = true;
                this.formItem.validateFailed = true;
            } else {
                this.validateFailed = false;
                this.formItem.validateFailed = false;
            }
        },
        afterLeave() {
            this.edit = false;
            this.isEditing = false;
            this.$emit(
                "edit-form-item",
                JSON.stringify(this.linkObj),
                this.formItem.field_key,
                this.formItem.mode
            );
            this.doValidate();
        },
        copyLink() {
            let tempTextarea = document.createElement("textarea");
            if (this.linkObj.url) {
                // 设置textarea的value为当前网址
                tempTextarea.value = this.linkObj.url;
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
                this.$refs["popover"].doClose();
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
            this.$refs["popover"].doClose();
        },
        jump() {
            if (this.linkObj.url) {
                let url = this.linkObj.url;
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
        }
    }
};
</script>

<style lang="scss" scoped>
@import "../style.scss";
.default-text {
    color: #c0c4cc;
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
