<template>
    <!-- https://blog.csdn.net/Kylincsg/article/details/135751130 -->
    <div id="comment-out-box">
        <div class="footer-comment-send-box" :class="{ focus: focus }">
            <div
                :id="`comment-input-data${id}`"
                class="input-border"
                :class="{ 'pointer-focus': focus, disabled: !hasProgressAuth }"
                :contenteditable="hasProgressAuth ? 'plaintext-only' : 'false'"
                :data-placeholder="placeholder"
                @input="handleInput"
                @keyup="handleKeyup"
                @keydown="handleKeydown"
                @click.stop="handleKeyup"
                @paste="optimizePasteEvent"
                @focus="inputFocus"
                @blur="inputBlur"
            ></div>
            <div v-if="cancleWithSave" class="bottom-area">
                <el-button
                    class="basic-ui height24"
                    size="small"
                    @click="cancleContent"
                >
                    取消</el-button
                >
                <el-button
                    class="basic-ui height24"
                    size="small"
                    type="primary"
                    :disabled="!value"
                    @click="confirmContent"
                >
                    保存</el-button
                >
            </div>
            <div v-if="confirmBtn && (focus || value)" class="bottom-area">
                <el-button
                    class="basic-ui height24"
                    size="small"
                    type="primary"
                    :disabled="!value.trim()"
                    @click="confirmContent"
                >
                    发送</el-button
                >
            </div>
            <div
                :id="`mention-dropdown-list${id}`"
                class="mention-box"
                ref="MentionBox"
                v-if="isShowBox"
                :style="{
                    top: mentionTop + 'px',
                    left: mentionLeft + 'px'
                }"
            >
                <div class="user-box" ref="UserBox" v-loading="mentionLoading">
                    <div
                        class="mention-item"
                        :class="{
                            'current-selected': currentSelectId === item.id
                        }"
                        v-for="item in filterUserList"
                        :data-id="item.id"
                        :key="item.id"
                        @click.stop="selectUser(item)"
                        @mouseenter="selectEnter(item)"
                    >
                        <el-avatar
                            class="progress-avatar"
                            icon="el-icon-user-solid"
                            size="small"
                            :src="getAvatar(item.avatar)"
                        ></el-avatar>
                        <span
                            v-html="searchHeightLight(item, selectionName)"
                        ></span>
                    </div>
                    <div v-if="!filterUserList.length" class="no-msg">
                        暂无数据
                    </div>
                </div>
            </div>
        </div>
        <!--     下拉面板 表情+文件    -->
        <!-- <div class="footer-power-drawer" v-if="showDrawerDown">
            <div class="emote-box" v-if="showEmoteOrFile === 1">
                <span
                    class="emote-item"
                    v-for="emote in emoteData"
                    :key="emote"
                    @click="insertEmoji"
                    >{{ emote }}</span
                >
            </div>
        </div> -->
    </div>
</template>
<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import { imgHost } from "@/assets/tool/const";
import { range } from "lodash";
export default {
    name: "tagDetail",
    props: {
        msgObj: {
            type: Object,
            default: () => {}
        },
        cancleWithSave: {
            type: Boolean,
            default: false
        },
        confirmBtn: {
            type: Boolean,
            default: false
        },
        mentionListPosition: {
            type: String,
            default: "bottom"
        },
        record: {
            type: String,
            default: ""
        },
        value: {
            type: String,
            default: ""
        },
        id: {
            type: String,
            default: "editor"
        },
        placeholder: {
            type: String,
            default: "请输入"
        },
        hasProgressAuth: {
            type: Boolean,
            default: false
        }
    },
    data() {
        return {
            innerHTML: "",
            showDrawerDown: false,
            showSendBtn: false,
            ignoreUserDataList: [], // @功能获取存储用户的数据
            ignoreUserIdList: [], // @功能存储用户的选中id集合
            filterUserList: [], // @列表
            isShowBox: false, // 是否显示@列表
            mentionTop: 0,
            mentionBottom: 0, // @列表位置 bottom
            mentionLeft: 20, // @列表位置 left
            savedSelectionMention: null, // 保存选中位置
            savedSelectionRange: null, // 保存选中位置
            selectionName: false, // @列表筛选用户
            isRecording: false, // 是否正在记录
            // ... 其他功能的变量 ...
            mentionLoading: true, // @loading
            currentSelectId: null, // @预选，回车可以直接选择
            focus: false
        };
    },
    computed: {
        // focus || innerHTML
        // 输入框撑起来了，向外传递高度
        propUp() {
            if (this.focus || this.value) {
                return true;
            } else {
                return false;
            }
        },
        rangeLength() {
            return (
                this.savedSelectionRange.length > 0 &&
                this.savedSelectionRange[0].startOffset -
                    this.savedSelectionMention[0].startOffset
            );
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        }
    },
    watch: {
        propUp(boolean) {
            this.$emit("get-propUp", boolean);
        }
    },
    mounted() {
        // if (this.record) {
        //     const commentBox = document.querySelector(
        //         `#comment-input-data${this.id}`
        //     );
        //     commentBox.innerHTML = this.record;
        // }
        if (this.value) {
            const commentBox = document.querySelector(
                `#comment-input-data${this.id}`
            );
            commentBox.innerHTML = this.value === "<br>" ? "" : this.value;
        }
        // document.querySelector(
        //         `#comment-input-data${this.id}`
        //     ).addEventListener('keydown')
        // 注册@mention列表回车事件
        let that = this;
        document
            .querySelector(`#comment-input-data${this.id}`)
            .addEventListener("keydown", function (e) {
                // 回车
                if (that.isShowBox && e.key === "Enter") {
                    // 如果列表展开了 就选择第一个
                    // if (that.isShowBox && that.ignoreUserDataList.length) {
                    if (that.ignoreUserDataList.length) {
                        e.preventDefault();
                        // document.execCommand("insertHTML", false, "<br>");
                        let row = _.find(that.ignoreUserDataList, {
                            id: that.currentSelectId
                        });
                        if (row) {
                            that.selectUser(row);
                        }
                    }
                }
                // 上下键
                if (
                    that.isShowBox &&
                    (e.key === "ArrowUp" || e.key === "ArrowDown")
                ) {
                    // 如果列表展开了 就选择第一个
                    // if (that.isShowBox && that.ignoreUserDataList.length) {
                    if (that.ignoreUserDataList.length) {
                        e.preventDefault();
                        // 处理选择,处理列表随着上下键滚动
                        if (e.key === "ArrowDown") {
                            e.preventDefault(); // 阻止默认的向下滚动行为
                            // 切换到下一个选项
                            let currentIndex = _.findIndex(
                                that.ignoreUserDataList,
                                {
                                    id: that.currentSelectId
                                }
                            );
                            // 计算当前选中的是否为第一个或最后一个
                            if (
                                currentIndex <
                                that.ignoreUserDataList.length - 1
                            ) {
                                currentIndex += 1;
                                that.currentSelectId =
                                    that.ignoreUserDataList[currentIndex].id;
                            }
                            // 滚动逻辑
                            let list = that.$refs.UserBox;
                            if (currentIndex > 4) {
                                list.scrollTop += 40;
                            }
                        } else if (e.key === "ArrowUp") {
                            e.preventDefault(); // 阻止默认的向上滚动行为
                            // 切换到上一个选项
                            let currentIndex = _.findIndex(
                                that.ignoreUserDataList,
                                {
                                    id: that.currentSelectId
                                }
                            );
                            // 计算当前选中的是否为第一个或最后一个
                            if (currentIndex > 0) {
                                currentIndex -= 1;
                                that.currentSelectId =
                                    that.ignoreUserDataList[currentIndex].id;
                            }
                            // 滚动逻辑
                            let list = that.$refs.UserBox;
                            if (currentIndex < 5) {
                                list.scrollTop -= 40;
                            }
                        }
                    }
                }
                // 空格
                if (that.isShowBox && e.keyCode === 32) {
                    // 关闭下拉box
                    that.isRecording = false;
                    that.isShowBox = false;
                }
            });
    },
    methods: {
        // 取消事件
        cancleContent() {
            this.$emit("cancle-content", this.msgObj);
        },
        // 发送按钮事件
        confirmContent() {
            // const commentBox = document.querySelector(
            //     `#comment-input-data${this.id}`
            // );
            this.$emit("confirm-content", this.value, this.msgObj);
        },
        inputFocus() {
            if (!this.hasProgressAuth) return;
            // 聚焦
            this.focus = true;
        },
        inputBlur() {
            this.focus = false;
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

        selectEnter(row) {
            this.currentSelectId = row.id;
        },
        blur: _.debounce(function () {
            // 关闭mention列表 取消关键词搜索逻辑
            this.isRecording = false;
            this.isShowBox = false;
        }, 200),
        // 重置
        reset() {
            this.isRecording = false;
            this.isShowBox = false;
            // const commentBox = document.querySelector("#comment-input-data");
            const commentBox = document.querySelector(
                `#comment-input-data${this.id}`
            );
            commentBox.innerHTML = "";
        },
        submitInner() {
            // const commentBox = document.querySelector("#comment-input-data");
            const commentBox = document.querySelector(
                `#comment-input-data${this.id}`
            );
            if (commentBox.innerHTML === "<br>") {
                this.$emit("input", "");
            } else {
                // 判断是否为空
                let valiStr = commentBox.innerHTML.replace(/&nbsp;/gi, "");
                if (valiStr.trim()) {
                    this.$emit("input", commentBox.innerHTML);
                } else {
                    this.$emit("input", "");
                }
            }
        },
        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        },
        /**
         * @函数描述: 输入框输入变化时的回调函数
         **/
        handleInput(event) {
            // const commentBox = document.querySelector("#comment-input-data");
            const commentBox = document.querySelector(
                `#comment-input-data${this.id}`
            );
            // this.innerHTML = commentBox.innerHTML;
            // 显示隐藏评论按钮
            if (commentBox.innerHTML.length > 0) {
                this.showSendBtn = true;
                if (this.showDrawerDown && this.showEmoteOrFile === 2) {
                    this.showDrawerDown = false;
                }
            } else {
                this.showSendBtn = false;
            }
            // 输入框内容有实际变动时屏蔽回车
            if (this.isRecording && event.inputType) {
                this.savedSelectionRange = this.saveSelection();
            }
            // 获取当前选区的范围对象
            const selection = window.getSelection();
            if (selection.rangeCount > 0) {
                const range = selection.getRangeAt(0);
                // 折叠range以确保它的start和end在相同位置
                range.collapse(true);

                // 如果不在编辑div的起始位置，创建一个新的range来查找前一个字符
                if (range.startOffset > 0) {
                    // 复制当前选区的范围
                    const rangeStart = range.cloneRange();

                    // 将这个新range的起始点向前移动一个字符
                    rangeStart.setStart(
                        rangeStart.startContainer,
                        range.startOffset - 1
                    );
                    // 选择这个字符
                    rangeStart.setEnd(
                        rangeStart.startContainer,
                        range.startOffset
                    );
                    // 获取这个字符并检查它是否是 '@'
                    const charBeforeCursor = rangeStart.toString();
                    let parentLocalName =
                        rangeStart.startContainer.parentNode.localName;
                    let parentClassName =
                        rangeStart.startContainer.parentNode.className;
                    // 当前输入符号是@ 且父级标签不能是span类名不能是mention-at
                    if (
                        charBeforeCursor === "@" &&
                        parentLocalName !== "span" &&
                        parentClassName !== "mention-at"
                    ) {
                        this.ignoreUserIdList = this.mentionCreateUserIdList(
                            commentBox.innerHTML
                        );
                        this.savedSelectionMention = this.saveSelection();
                        // rect 是输入的高度位置
                        // event.target 是@符号的位置
                        // 222 是mention列表的高度
                        const rect = range.getClientRects()[0];
                        this.getUserListPosition(event);
                        // this.mentionBottom =
                        //     rect.top -
                        //     event.target.getBoundingClientRect().top +
                        //     30;
                        // if (this.mentionListPosition === "top") {
                        //     this.mentionTop =
                        //         rect.top -
                        //         event.target.getBoundingClientRect().top -
                        //         222;
                        // } else {
                        //     this.mentionTop =
                        //         rect.top -
                        //         event.target.getBoundingClientRect().top +
                        //         20;
                        // }
                        // this.mentionLeft =
                        //     rect.left -
                        //     event.target.getBoundingClientRect().left;
                        // if (event.target.clientWidth - this.mentionLeft < 85) {
                        //     this.mentionLeft = this.mentionLeft - 260;
                        // }
                        this.isRecording = true;
                    }
                    // 可补充或事件 手动移除中文输入法下的所选区域
                }
            }
            this.submitInner();
        },

        /**
         * @函数描述: 键盘输入时的回调函数 - 【弹起】
         **/
        handleKeyup: _.debounce(function (e) {
            // const commentBox = document.querySelector("#comment-input-data");
            const commentBox = document.querySelector(
                `#comment-input-data${this.id}`
            );
            // 屏蔽左右键
            if (e.key !== "ArrowLeft" && e.key !== "ArrowRight") {
                this.savedSelectionRange = this.saveSelection();
            }
            if (commentBox.innerHTML === "") {
                this.isRecording = false;
            }
            if (this.isRecording) {
                if (this.rangeLength < 0) {
                    this.isShowBox = false;
                } else {
                    this.isShowBox = true;
                    // 非上下左右回车再请求
                    if (
                        e.key !== "ArrowDown" &&
                        e.key !== "ArrowUp" &&
                        e.key !== "ArrowLeft" &&
                        e.key !== "ArrowRight" &&
                        e.key !== "keydown"
                    ) {
                        this.filterSelectionName(e);
                    }
                }
            }
        }, 100),

        /**
         * @函数描述: 键盘输入时的回调函数 - 【按下】
         * @param: {object} event
         **/
        handleKeydown(event) {
            const BACKSPACE_KEY = "Backspace";
            const selection = window.getSelection();
            if (selection.rangeCount > 0) {
                const range = selection.getRangeAt(0);
                // 如果按下的是返回键
                if (event.key === BACKSPACE_KEY) {
                    // 检查并且消除选区
                    if (!range.collapsed) {
                        // 用户有一个激活的选区，可能会删除多个字符
                        return;
                    }
                    // 如果不在内容的开头位置
                    if (range.startOffset > 0) {
                        // 创建一个新的range来查找要删除的字符
                        const rangeToDelete = range.cloneRange();
                        rangeToDelete.setStart(
                            rangeToDelete.startContainer,
                            range.startOffset - 1
                        );
                        rangeToDelete.setEnd(
                            rangeToDelete.startContainer,
                            range.startOffset
                        );
                        const charBeforeCursor = rangeToDelete.toString();
                        if (charBeforeCursor === "@") {
                            this.isShowBox = false;
                            this.savedSelectionMention = null;
                            this.savedSelectionRange = null;
                            this.isRecording = false;
                        }
                    } else {
                        this.isRecording = false;
                        this.savedSelectionMention = null;
                    }
                    this.showSendBtn = false;
                }
            }
        },

        /**
         * @函数描述: 过滤输入框内的姓名
         **/
        // filterSelectionName: _.debounce(function () {
        filterSelectionName(e) {
            // 关键词取当前selecttion内的
            // const commentBox = document.querySelector(
            //     `#comment-input-data${this.id}`
            // );
            // this.selectionName = commentBox.innerHTML
            //     .replace(/<[^>]*>[^<]*<\/[^>]*>/g, function (match) {
            //         return match.replace(/[^]*/g, ""); // 使用正则表达式替换标签的内容
            //     })
            //     .split("@")[1]
            //     .substr(0, this.rangeLength);
            const selection = window.getSelection();
            let range = selection.getRangeAt(0);
            //   在 range.startContainer 中取 range.startOffset位数往前取直到@
            let area = range.startContainer.data;
            let start = range.endOffset - this.rangeLength;
            let end = range.endOffset;
            let keyword = area.substring(start, end);
            this.selectionName = keyword;
            // this.selectionName = selection.focusNode.nodeValue
            //     .replace(/<[^>]*>[^<]*<\/[^>]*>/g, function (match) {
            //         return match.replace(/[^]*/g, ""); // 使用正则表达式替换标签的内容
            //     })
            //     .split("@")[1]
            //     .substr(0, this.rangeLength);
            // 换成接口查询
            // 非上下左右键采用请求
            this.getUserList(e);
        },
        // }, 100),
        getUserList(e) {
            this.mentionLoading = true;
            let params = {
                ws_id: this.curSpace.id,
                key: this.selectionName,
                page_size: 50,
                page: 1
            };
            api.getPeopleList(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.ignoreUserDataList = res.data || [];
                    this.filterUserList = res.data || [];
                    this.isShowBox = true;
                    // 将第一个存入当前选中
                    this.currentSelectId = this.ignoreUserDataList[0].id;
                } else {
                    this.ignoreUserDataList = [];
                    this.filterUserList = [];
                    this.currentSelectId = null;
                }
                this.mentionLoading = false;
                this.$nextTick(() => {
                    this.getUserListPosition(e);
                });
            });
        },
        getUserListPosition(event) {
            const selection = window.getSelection();
            let range = selection.getRangeAt(0);
            const rect = range.getClientRects()[0];
            let MentionBox = this.$refs.MentionBox;
            let height = 220;
            if (MentionBox) {
                height = MentionBox.clientHeight;
            }
            this.mentionBottom =
                rect.top - event.target.getBoundingClientRect().top + 30;
            if (this.mentionListPosition === "top") {
                this.mentionTop =
                    rect.top -
                    event.target.getBoundingClientRect().top -
                    height;
            } else {
                this.mentionTop =
                    rect.top - event.target.getBoundingClientRect().top + 20;
            }
            this.mentionLeft =
                rect.left - event.target.getBoundingClientRect().left;
            if (event.target.clientWidth - this.mentionLeft < 85) {
                this.mentionLeft = this.mentionLeft - 260;
            }
        },
        /**
         * 选择用户
         * @param row
         */
        selectUser(row) {
            const span = document.createElement("span");
            span.className = "mention-at";
            span.id = row.id;
            span.textContent = `@${row.full_name}`;
            span.setAttribute("contenteditable", "false");
            // const commentBox = document.querySelector("#comment-input-data");
            const commentBox = document.querySelector(
                `#comment-input-data${this.id}`
            );
            // 重新聚焦到评论框并恢复之前的选择
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
                // 如果@有内容 又左右键改了位置 ，使用回车会出问题
                range.setStart(range.startContainer, range.startOffset - 1);
                range.setEnd(
                    range.startContainer,
                    range.endOffset + this.rangeLength
                );
                range.deleteContents(); // 删除 "@"
                const newRange = document.createRange();
                range.insertNode(span);
                newRange.setStartAfter(span);
                newRange.collapse(true);
                // APPLY THE NEW RANGE
                selection.removeAllRanges();
                selection.addRange(newRange);
            }
            this.isRecording = false;
            this.isShowBox = false;
            // 清除savedSelection状态，因为已经不再需要了
            this.savedSelectionMention = null;
            // this.savedSelectionRange = null;
            this.savedSelectionRange = this.saveSelection();
            this.submitInner();
            this.ignoreUserDataList = [];
            this.filterUserList = [];
            this.mentionLoading = true;
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

        /**
         * @函数描述: 模糊搜索下拉框精确匹配文字高亮
         * @param {object} content 下拉项的内容
         */
        searchHeightLight(content, title) {
            let reg = "";
            let dataToReplace = `<span style="color: #ff4d51">${title}</span>`;
            // let roleName = `(<span>${content.roleName}</span>)`;
            if (content.full_name.indexOf(title) !== -1) {
                // reg =
                //     content.full_name.replace(title, dataToReplace) + roleName;
                reg = content.full_name.replace(title, dataToReplace);
            } else {
                // reg = content.full_name + roleName;
                reg = content.full_name;
            }
            return reg;
        },

        /**
         * @函数描述: 点击表情添加到输入框
         * @param: {Object} event 点击的对象
         **/
        insertEmoji(emoji) {
            return;
            // const commentBox = document.querySelector("#comment-input-data");
            const commentBox = document.querySelector(
                `#comment-input-data${this.id}`
            );
            // 重新聚焦到评论框并恢复之前的选择
            commentBox.focus();
            const selection = window.getSelection();
            if (
                this.savedSelectionRange &&
                this.savedSelectionRange.length > 0
            ) {
                selection.removeAllRanges();
                selection.addRange(this.savedSelectionRange[0]);
            }
            if (selection.rangeCount > 0) {
                const range = selection.getRangeAt(0);
                range.setStart(range.startContainer, range.startOffset);
                range.setEnd(range.startContainer, range.endOffset);
                range.deleteContents();
                const newRange = document.createRange();
                const emojiNode = document.createTextNode(emoji);
                range.insertNode(emojiNode);
                newRange.setStartAfter(emojiNode);
                newRange.collapse(true);

                // APPLY THE NEW RANGE
                selection.removeAllRanges();
                selection.addRange(newRange);
                this.savedSelectionRange = this.saveSelection();
            }

            this.showSendBtn = true;
        },

        /**
         * @函数描述: 获取@提及功能的用户id集合
         * @param: {string} text
         **/
        mentionCreateUserIdList(string) {
            const parser = new DOMParser();
            const doc = parser.parseFromString(string, "text/html");
            const spanElements = doc.querySelectorAll("span");
            const idValues = Array.from(spanElements).map((span) => span.id);
            const filterValues = idValues.filter((item) => item !== "");
            return filterValues;
        }
    }
};
</script>

<style lang="scss" scoped>
.footer-comment-send-box {
    position: relative;
    .bottom-area {
        padding-top: 10px;
        display: flex;
        justify-content: end;
    }
}
.no-msg {
    height: 40px;
    line-height: 40px;
    text-align: center;
}
.progress-avatar.el-avatar {
    width: 20px;
    height: 20px;
    line-height: 1;
    vertical-align: middle;
    margin-right: 4px;
}
// mention列表
.mention-box {
    position: absolute;
    min-width: 120px;
    max-width: 200px;
    padding: 10px 0;
    background: #fff;
    border-radius: 4px;
    border: 1px solid gainsboro;
    box-sizing: border-box;
    box-shadow: 0 0 5px rgba(0, 0, 0, 0.1);
    max-height: 245px;
    z-index: 9999999;
    font-size: 12px;
    // & > div:first-child {
    //     font-weight: bold;
    //     padding: 0 0 4px 0;
    // }
    .user-box {
        padding: 0;
        margin: 0;
        overflow: hidden;
        overflow-y: auto;
        max-height: 200px;
        // &::-webkit-scrollbar {
        //     display: none;
        // }
        .mention-item {
            font-size: 14px;
            padding: 0 12px;
            height: 40px;
            line-height: 40px;
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;
            cursor: pointer;
            display: flex;
            align-items: center;
            &.current-selected {
                background: #f1f9ff;
            }
            // &:hover {
            //     background: #f1f9ff;
            // }
        }
    }
}
.input-border {
    box-sizing: border-box;
    // height: 34px;
    font-size: 14px;
    padding: 4px 8px;
    overflow-y: auto;
    border: 1px solid #dcdfe6;
    border-radius: 4px;
    background-color: #fff;
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
    // &:not(:empty) {
    //     height: 98px;
    // }
    &.pointer-focus {
        height: 98px;
        border: 1px solid #1890ff;
    }
    &.disabled {
        opacity: 0.5;
        cursor: not-allowed;
        background-color: #f5f5f5;
        border: 1px solid #dcdfe6;
        // pointer-events: none;
        &:hover {
            border: 1px solid #dcdfe6;
        }
    }
    &:hover {
        border: 1px solid #1890ff;
    }
    &:focus-visible {
        outline: none;
    }
}
</style>
