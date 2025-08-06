<template>
    <div
        class="column-block"
        ref="column-block"
        :class="{ isEditing: isEditing }"
    >
        <el-popover
            ref="popover"
            popper-class="progress-popover email-rule"
            placement="bottom"
            :width="popoverWidth"
            trigger="click"
            @show="popoverShow"
            @hide="popoverHide"
        >
            <div class="popover-content">
                <!-- 邮件主题 -->
                <div>
                    <el-input
                        v-model="emailTitle"
                        placeholder="邮件主题"
                        size="small"
                    ></el-input>
                </div>
                <!-- 邮件内容 -->
                <div class="email-border-box" :class="{ focus: focus }">
                    <div
                        :id="`email-input` + id"
                        ref="EmailInput"
                        class="input-border"
                        contenteditable="paintext-only"
                        :data-placeholder="placeholder"
                        @input="handleInput"
                        @keyup="handleKeyup"
                        @keydown="handleKeydown"
                        @click.stop="handleKeyup"
                        @paste="optimizePasteEvent"
                        @focus="inputFocus"
                        @blur="inputBlur"
                    ></div>
                </div>
                <div class="field-list">
                    <span
                        class="field-tag-item"
                        v-for="(fieldItem, fieldIndex) in fieldList"
                        :key="fieldIndex"
                        @click.stop="insertField(fieldItem)"
                    >
                        <b
                            class="type-box"
                            :style="getType(fieldItem.mode)"
                        ></b>
                        {{ fieldItem.name }}</span
                    >
                </div>
            </div>
            <div slot="reference">
                <div class="output">
                    <span :class="{ empty: !emailTitle }">
                        <!-- 有名字优先展示名字，不然url -->
                        {{ emailTitle || "邮件主题" }}
                    </span>
                </div>
            </div>
        </el-popover>
        <b class="triangle"></b>
    </div>
</template>

<script>
import _ from "lodash";
import { FieldType } from "@/assets/tool/const";
export default {
    props: {
        title: {
            type: String,
            default: ""
        },
        value: {
            type: String,
            default: ""
        },
        id: {
            type: [Number, String],
            default: "necessaryAction"
        },
        fieldList: {
            type: Array,
            default: () => []
        }
    },
    data() {
        return {
            innerHTML: "",
            popoverWidth: 358,
            isEditing: false,
            focus: false,
            emailTitle: "",
            emailContents: "",
            placeholder: "邮件内容",
            savedSelectionMention: null, // 保存选中位置
            savedSelectionRange: null // 保存选中位置
        };
    },
    computed: {
        rangeLength() {
            if (this.savedSelectionRange) {
                return (
                    this.savedSelectionRange.length > 0 &&
                    this.savedSelectionRange[0].startOffset -
                        this.savedSelectionMention[0].startOffset
                );
            } else {
                return 0;
            }
        }
    },
    watch: {
        title: {
            handler(title) {
                this.emailTitle = title;
            },
            immediate: true
        },
        value: {
            handler(contents) {
                this.emailContents = contents;
            },
            immediate: true
        },
        emailTitle: {
            handler(title) {
                this.$emit("email-update", title, this.emailContents);
            }
        },
        emailContents: {
            handler(contents) {
                this.$emit("input", contents);
            }
        }
    },
    mounted() {},
    methods: {
        // 插入字段到标签
        insertField(field) {
            const span = document.createElement("span");
            span.className = "field-tag";
            span.id = field.field;
            span.textContent = `${field.name}`;
            span.setAttribute("contenteditable", "false");
            // 重新聚焦到评论框并恢复之前的选择
            const commentBox = document.querySelector(`#email-input` + this.id);
            commentBox.focus();
            const selection = window.getSelection();
            if (
                this.savedSelectionMention &&
                this.savedSelectionMention.length > 0
            ) {
                selection.removeAllRanges();
                selection.addRange(this.savedSelectionMention[0]);
            }
            if (selection.rangeCount > 0) {
                const range = selection.getRangeAt(0);
                // // 如果@有内容 又左右键改了位置 ，使用回车会出问题
                // range是input-border
                // range.startOffset 期望是当前光标位置
                range.setStart(range.startContainer, range.startOffset);
                range.setEnd(
                    range.startContainer,
                    range.endOffset + this.rangeLength
                );
                // newRange是记录了开始位置的input-border
                const newRange = document.createRange();
                range.insertNode(span);
                newRange.setStartAfter(span);
                newRange.collapse(true);
                // APPLY THE NEW RANGE
                selection.removeAllRanges();
                selection.addRange(newRange);
            }
            // 需要刷新savedSelectionMention位置
            this.savedSelectionRange = this.saveSelection();
            this.savedSelectionMention = this.saveSelection();
            this.emailContents = commentBox.innerHTML;
        },
        inputFocus() {

        },
        inputBlur() {

        },
        // 粘贴事件
        optimizePasteEvent(e) {
            // 监听粘贴内容到输入框事件，对内容进行处理 处理掉复制的样式标签，只拿取文本部分
            e.stopPropagation();
            e.preventDefault();
            let text = "",
                event = e.originalEvent || e;
            if (event.clipboardData && event.clipboardData.getData) {
                text = event.clipboardData.getData("text/plain");
            } else if (window.clipboardData && window.clipboardData.getData) {
                text = window.clipboardData.getData("text");
            }
            if (document.queryCommandSupported("insertText")) {
                document.execCommand("insertText", false, text);
            } else {
                document.execCommand("paste", false, text);
            }
        },
        /**
         * @函数描述: 输入框输入变化时的回调函数
         **/
        handleInput(event) {
            // 输入框内容有实际变动时屏蔽回车
            this.savedSelectionRange = this.saveSelection();
            // 获取当前选区的范围对象
            const selection = window.getSelection();
            if (selection.rangeCount > 0) {
                const range = selection.getRangeAt(0);
                // 折叠range以确保它的start和end在相同位置
                range.collapse(true);
                // 如果不在编辑div的起始位置，创建一个新的range来查找前一个字符
                if (range.startOffset > 0) {
                    this.savedSelectionMention = this.saveSelection();
                    // 复制当前选区的范围
                    const rangeStart = range.cloneRange();
                    // 将这个新range的起始点向前移动一个字符
                    rangeStart.setStart(
                        rangeStart.startContainer,
                        range.startOffset
                    );
                    // 选择这个字符
                    rangeStart.setEnd(
                        rangeStart.startContainer,
                        range.startOffset
                    );
                    // 检查当前元素
                    let parentLocalName =
                        rangeStart.startContainer.parentNode.localName;
                    let parentClassName =
                        rangeStart.startContainer.parentNode.className;
                    if (
                        parentLocalName !== "span" &&
                        parentClassName !== "tag-item"
                    ) {
                        this.savedSelectionMention = this.saveSelection();
                    }
                }
            }
            // this.submitInner();
            const commentBox = document.querySelector(`#email-input` + this.id);
            this.emailContents = commentBox.innerHTML;
        },
        /**
         * 这里是之前提到的saveSelection函数的示例实现
         * @returns {*[]}
         */
        saveSelection() {
            const ranges = [];
            const selection = window.getSelection();
            for (let i = 0; i < selection.rangeCount; i++) {
                ranges.push(selection.getRangeAt(i));
            }
            return ranges;
        },
        handleKeyup(e) {
            // 光标切换位置要刷新
            this.savedSelectionMention = this.saveSelection();
            this.savedSelectionRange = this.saveSelection();
        },

        /**
         * @函数描述: 键盘输入时的回调函数 - 【按下】
         * @param: {object} event
         **/
        handleKeydown(event) {},
        popoverShow() {
            const commentBox = document.querySelector(`#email-input` + this.id);
            if (this.emailContents !== commentBox.innerHTML) {
                // 不一样就需要回显
                if (commentBox) {
                    commentBox.innerHTML = this.emailContents;
                }
            }
        },
        popoverHide() {},
        getType(type) {
            if (type) {
                return {
                    "background-image": `url(${FieldType[type]})`
                };
            }
        }
    }
};
</script>

<style lang="scss" scoped>
.output {
    height: 32px;
    line-height: 32px;
    border-radius: 4px;
    padding: 0 10px;
    border: 1px solid #e6e9f0;
    cursor: pointer;
    .empty {
        color: #a6abbc;
    }
}
.email-border-box {
    margin-top: 12px;
    padding: 14px 6px;
    border-top: 1px dotted #dcdfe6;
}
.field-list {
    // margin-top: 12px;
    border-top: 1px dotted #dcdfe6;
    padding-top: 8px;
    .field-tag-item {
        box-sizing: border-box;
        display: inline-block;
        padding: 4px 8px;
        font-size: 12px;
        height: 24px;
        // line-height: 24px;
        margin-right: 4px;
        margin-bottom: 4px;
        background-color: #f5f5f5;
        border-radius: 4px;
        cursor: pointer;
    }
}
.input-border {
    box-sizing: border-box;
    height: 108px;
    font-size: 14px;
    // padding: 14px 6px;
    overflow-y: auto;
    &:empty {
        &:before {
            font-size: 14px;
            color: #c0c4cc;
            content: attr(data-placeholder);
            vertical-align: middle;
            display: inline-block;
            line-height: 24px;
            height: 24px;
        }
    }
    &:not(:empty) {
        // height: 72px;
    }
    &:focus {
        // height: 72px;
    }
    &:focus-visible {
        outline: none;
    }
}
.type-box {
    display: inline-block;
    width: 14px;
    height: 14px;
    background-size: 100% 100%;
    vertical-align: middle;
    position: relative;
    top: -1px;
}
</style>
<style lang="scss">
.field-tag {
    color: #1890ff;
    text-decoration: underline;
    margin: 0 4px;
}
.el-popper.progress-popover.email-rule {
    padding: 12px 12px 8px 12px;
}
</style>
