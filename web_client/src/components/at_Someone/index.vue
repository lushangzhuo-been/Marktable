<template>
    <div class="at-someone-content" :class="{ focus: editorFocus }">
        <div
            ref="editor"
            class="editor"
            :id="id"
            @keydown="enterEv($event)"
        ></div>
        <!-- @列表 -->
        <div
            v-show="showFlag"
            class="at-someone"
            :style="{
                left: left + 'px',
                top: top + 'px',
                visibility: showFlag
            }"
        >
            <div class="at-someone-box" ref="atSomeoneBox">
                <!-- 搜索框 -->
                <div class="search-input">
                    <el-input
                        ref="searchInput"
                        class="basic-ui progress-search-input"
                        v-model="searchKeyword"
                        size="small"
                        placeholder="搜索"
                        prefix-icon="el-icon-search"
                        @input="handlerSearchInput"
                        @blur="searchBlue"
                    ></el-input>
                </div>
                <div :loading="userLoading">
                    <!-- 列表 -->
                    <div class="list-content">
                        <div
                            class="list-item multiple-people-list-item"
                            v-for="(listItem, listIndex) in userList"
                            :key="listIndex"
                            @click="selectPerson(listItem)"
                        >
                            <el-avatar
                                class="progress-avatar"
                                icon="el-icon-user-solid"
                                size="small"
                                :src="getAvatar(listItem.avatar)"
                            ></el-avatar>
                            {{ listItem.full_name }}
                        </div>
                        <div v-show="!userList.length" class="list-no-data">
                            暂无数据
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import Vue from "vue";
import E from "wangeditor";
import { position, offset } from "caret-pos";
import { imgHost } from "@/assets/tool/const";
import api from "@/common/api/module/progress";
export default {
    name: "AtSomeone",
    props: {
        from: {
            type: String,
            default: ""
        },
        placeholder: {
            type: String,
            default: ""
        },
        id: {
            type: String,
            default: "editor"
        },
        record: {
            type: String,
            default: ""
        },
        height: {
            type: Number,
            default: 72
        },
        value: {
            type: String,
            default: ""
        }
    },
    data() {
        return {
            showEmojiFlag: "hidden",
            showFlag: "hidden",
            userLoading: false,
            left: "",
            top: "",
            browserType: "",
            editor: null,
            position: {
                range: "",
                selection: ""
            },
            selectionNode: null,
            userList: [],
            searchKeyword: "",
            editorFocus: false
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
    watch: {},
    mounted() {
        this.createEditor();
    },
    beforeDestroy() {
        // 销毁编辑器
        this.editor.destroy();
        this.editor = null;
    },
    methods: {
        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        },
        reset() {
            this.editor.txt.clear();
        },
        searchBlue() {
            let time = setTimeout(() => {
                this.showFlag = "hidden";
                this.searchKeyword = "";
                clearTimeout(time);
                time = null;
            }, 300);
        },
        handlerSearchInput() {
            this.userLoading = true;
            // 获取用户列表
            this.getPeopleList();
        },
        getPeopleList() {
            let params = {
                ws_id: this.curSpace.id,
                key: this.searchKeyword,
                page_size: 50,
                page: 1
            };
            api.getPeopleList(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.userList = res.data;
                } else {
                    this.userList = [];
                }
            });
        },
        createEditor() {
            let editor = new E(`#${this.id}`);
            editor.config.onchange = (newHtml) => {
                // 获取纯文本
                const text = editor.txt.text().replace(/&nbsp;/gi, "");
                // 超出1000字处理
                if (text.length >= 1000) {
                    // 因为是富文本编辑器，无法做截取只能提示
                    // this.$message.warning("最多可输入1000字!");
                    this.$message({
                        showClose: true,
                        message: "最多可输入1000字!",
                        type: "warning"
                    });
                }
                this.$emit("updataRecordHtml", newHtml, text);
            };
            editor.config.onfocus = (newHtml) => {
                this.editorFocus = true;
            };
            editor.config.onblur = (newHtml) => {
                this.editorFocus = false;
            };
            // 配置触发 onchange 的时间频率，默认为 200ms
            editor.config.onchangeTimeout = 500; // 修改为 500ms
            editor.config.height = this.height;
            editor.config.placeholder = this.placeholder;
            //菜单置空
            editor.config.menus = [];
            //创建
            editor.create();
            editor.txt.html(this.record);
            this.editor = editor;
        },
        enterEv(e) {
            // 优先确保以聚焦
            let ele = this.editor.$textElem.elems[0];
            //getSelection是另一种获取用户选择的文本范围或光标的当前位置的方法
            let selection = getSelection();
            this.selectionNode = selection;
            //判断输入的是@符号
            if (
                ((e.keyCode === 229 && e.key === "@") ||
                    (e.keyCode === 229 && e.code === "Digit2") ||
                    e.keyCode === 50) &&
                e.shiftKey
            ) {
                // setTimeout(() => {
                e.preventDefault ? e.preventDefault() : (e.returnValue = false);
                this.position = {
                    range: selection.getRangeAt(0),
                    selection: selection
                };
                try {
                    if (
                        this.editor.$textElem.elems[0].firstChild.tagName ===
                        "BR"
                    ) {
                        this.editor.$textElem.elems[0].removeChild(
                            this.editor.$textElem.elems[0].firstChild
                        );
                    }
                    if (
                        this.editor.$textElem.elems[0].firstChild.firstChild
                            .tagName === "BR"
                    ) {
                        // 解决不输入内容直接@时候出现的<br>换行标签
                        this.editor.$textElem.elems[0].firstChild.removeChild(
                            this.editor.$textElem.elems[0].firstChild.firstChild
                        );
                    }
                } catch (e) {
                    console.log(e);
                }
                //根据输入@的位置定位下拉框的位置
                // this.getPosition();
                this.showFlag = "visible";
                //下拉框搜索框自动获取焦点
                this.$nextTick(() => {
                    this.searchKeyword = "";
                    this.$refs.searchInput.focus();
                });
                this.handlerSearchInput();
                // }, 200);
                // 兼容
            } else if (e.code === "Backspace" || e.key === "Backspace") {
                // e.preventDefault()
                //  let selection = getSelection()
                let range = selection.getRangeAt(0);
                let removeNode = null;
                if (this.browserType === "Firefox") {
                    if (range.startContainer.className !== "at-text") {
                        removeNode =
                            range.startContainer.previousElementSibling;
                    }
                    if (
                        range.startContainer.parentElement.className !==
                            "at-text" &&
                        range.startContainer.lastChild !== null
                    ) {
                        removeNode =
                            range.startContainer.lastElementChild
                                .previousElementSibling;
                    }
                }
                if (this.browserType === "Chrome") {
                    if (
                        range.startContainer.textContent.length === 1 &&
                        range.startContainer.textContent.trim() === ""
                    ) {
                        removeNode =
                            range.startContainer.previousElementSibling;
                    }
                    if (
                        range.startContainer.parentNode.className === "at-text"
                    ) {
                        e.preventDefault
                            ? e.preventDefault()
                            : (e.returnValue = false);
                        removeNode = range.startContainer.parentNode;
                    }
                }
                if (this.browserType === "IE") {
                    if (
                        range.startContainer.nodeName !== "P" &&
                        range.startContainer.nodeValue.trim() === "" &&
                        range.startContainer.previousSibling.className ===
                            "at-text"
                    ) {
                        removeNode = range.startContainer.previousSibling;
                    }
                    if (
                        range.startContainer.parentNode.className ===
                            "at-text" &&
                        range.startContainer.nodeName === "#text" &&
                        range.startContainer.previousSibling == null
                    ) {
                        removeNode = range.startContainer.parentNode;
                    }
                }
                if (removeNode) {
                    // ele.removeChild(removeNode)
                    ele.firstChild.removeChild(removeNode);
                }
                this.showFlag = "hidden";
            }
        },

        selectPerson(item) {
            const { full_name, id } = item;
            this.showFlag = "hidden";
            // 获取选区对象
            let selection = this.position.selection;
            let range = this.position.range;
            // 清空手动输入的@符号  !!! 光标位置会重置为0 range.startOffset
            // let startOffset = range.startOffset
            // range.startContainer.data = range.startContainer.data.replace(
            //     /@$/,
            //     ""
            // );
            // let index = range.startContainer.data.length;
            // 按照自有属性搜索node data-we-empty-p
            // range.setStart(this.selectionNode, 0);
            // range.setEnd(this.selectionNode, index);
            // 结束
            // 生成需要显示的内容，包括一个 span 和一个空格。
            let spanNode1 = document.createElement("span");
            let spanNode2 = document.createElement("span");
            spanNode1.className = "at-text";
            spanNode1.style.color = "#3E74CA";
            spanNode1.innerHTML = "@" + full_name;
            spanNode1.dataset.id = id;
            //  设置@人的节点不可编辑
            spanNode1.contentEditable = false;
            spanNode2.innerHTML = "&nbsp;";
            // 将生成内容打包放在 Fragment 中，并获取生成内容的最后一个节点，也就是空格。
            //创建一个新的空白的文档片段
            let frag = document.createDocumentFragment(),
                node,
                lastNode;
            frag.appendChild(spanNode1);
            while ((node = spanNode2.firstChild)) {
                lastNode = frag.appendChild(node);
            }
            // 将 Fragment 中的内容放入 range 中，并将光标放在空格之后。
            range.insertNode(frag);
            selection.collapse(lastNode, 1);
            //将当前的选区折叠到最末尾的一个点
            selection.collapseToEnd();
        },
        showInner(html) {
            // 优先确保以聚焦
            //getSelection是另一种获取用户选择的文本范围或光标的当前位置的方法
            let selection = getSelection();
            // 将html插入
            let node = document.createElement("span");
            node.innerHTML = html;
            let frag = document.createDocumentFragment();
            frag.appendChild(node);
            selection.collapseToEnd();
        },
        //获取@位置
        getPosition() {
            this.showFlag = "visible";
            const ele = this.editor.$textElem.elems[0];
            const childEle = document.getElementsByClassName("at-someone")[0];
            let parentW = ele.offsetWidth;
            let parentH = ele.offsetHeight;
            let childW = childEle.offsetWidth;
            let childH = childEle.offsetHeight;
            const pos = position(ele);
            const off = offset(ele);
            // 弹框偏移超出父元素的宽高
            if (parentW - pos.left < childW) {
                this.left = off.left - childW;
            } else {
                this.left = off.left;
            }
            if (childH + off.top + 20 > document.documentElement.clientHeight) {
                this.top =
                    childH +
                    parentH +
                    off.top +
                    off.height +
                    40 -
                    document.documentElement.clientHeight;
            } else {
                this.top = off.top + 20;
            }

            // 如果是消息归档页面则写死left
            if (["message"].includes(this.from)) {
                this.left = 40;
            }
        }
    }
};
</script>

<style lang="scss" scoped>
.top-team-search-empty {
    display: flex;
    height: 200px;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    small {
        font-size: 12px !important;
        color: #c1c4cb;
        margin-top: 6px !important;
        display: block;
    }
}
.at-someone-content {
    border-radius: 4px;
    width: 100%;
    :deep(.w-e-toolbar) {
        display: none;
    }
    :deep(.w-e-text-container) {
        border: 1px solid #dcdfe6 !important;
        border-radius: 4px;
        z-index: 1 !important;
        .placeholder {
            left: 0;
            top: 0;
            padding: 4px 8px;
        }
        .w-e-text {
            text-align: left;
            padding: 4px !important;
            p {
                margin: 0;
            }
        }
    }
    &.focus {
        :deep(.w-e-text-container) {
            border: 1px solid #1890ff !important;
        }
    }
    .search-input {
        border-bottom: 1px solid #f0f1f5;
        padding-bottom: 4px;
        .basic-ui.progress-search-input {
            .el-input__inner {
                border: 1px solid rgba(0, 0, 0, 0);
                &:focus {
                    border: 1px solid rgba(0, 0, 0, 0);
                }
            }
        }
    }
    .at-someone {
        max-width: 260px;
        position: fixed;
        z-index: 5;
        border-radius: 3px;
        background: #fff;
        padding: 8px 0;
        border: 1px solid #e4e5ec;
        box-shadow: 0 4px 10px rgb(0 0 0 / 10%);
        // &::after {
        //   content: "";
        //   position: absolute;
        //   left: 14px;
        //   top: 0;
        //   display: block;
        //   box-sizing: border-box;
        //   width: 8px;
        //   height: 8px;
        //   transform: translate(-50%, -50%) rotate(45deg);
        //   z-index: 1;
        //   border: 1px solid #E4E5EB;
        //   border-right: 0;
        //   border-bottom: 0;
        //   background: #fff;
        //   border-top-left-radius: 2px;
        // }
        :deep(.arco-input-wrapper) {
            border-left: 0;
            border-top: 0;
            border-right: 0;
            border-color: #e4e5ec;
            padding: 0 16px;
            box-sizing: border-box;
            border-radius: 0;
            &:hover {
                background: #fff;
            }
        }
    }
}
//
.list-content {
    max-height: 200px;
    overflow-y: auto;
}
.list-item {
    padding: 0 12px;
    height: 40px;
    line-height: 40px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    text-align: left;
    &:hover {
        background-color: #f1f9ff;
    }
}
</style>
<style lang="scss">
@import "@/components/table/p_table/drill_style.scss";
</style>
