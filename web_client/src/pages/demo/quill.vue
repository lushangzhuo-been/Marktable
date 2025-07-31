<template>
    <!-- https://blog.csdn.net/Kylincsg/article/details/135751130 -->
    <div id="comment-out-box">
        <div class="footer-comment-send-box">
            <div
                id="comment-input-data"
                contenteditable="true"
                @input="handleInput"
                @keydown="handleKeydown"
                @keyup="handleKeyup"
                @click.stop="handleKeyup"
            ></div>
            <div class="footer-power-box">
                <div
                    class="comment-emote-btn"
                    @click.stop="clickShowDrawer(1)"
                ></div>
                <div
                    class="comment-file-btn"
                    v-if="!showSendBtn"
                    @click.stop="clickShowDrawer(2)"
                    :class="{
                        'comment-file-btn_active':
                            showDrawerDown && showEmoteOrFile === 2,
                    }"
                ></div>
                <div
                    class="comment-send-enter"
                    v-else
                    @click.self="enterSendCommentOrReply(1)"
                >
                    发 送
                </div>
            </div>
            <!-- :style="{
                    bottom: mentionBottom + 'px',
                    left: mentionLeft + 'px',
                }" -->
            <div
                class="mention-box"
                v-if="isShowBox"
            >
                <div>所有人（{{ filterUserList.length }}）</div>
                <div class="user-box">
                    <div
                        :data-id="item.userId"
                        v-for="item in filterUserList"
                        :key="item.userId"
                        @click.stop="selectUser(item)"
                    >
                        <img class="avatarImg" :src="item.avatar" alt="" />
                        <span
                            v-html="searchHeightLight(item, selectionName)"
                        ></span>
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
export default {
    name: "tagDetail",
    // components: { VuePdf },
    data() {
        return {
            showDrawerDown: false,
            showSendBtn: false,
            ignoreUserDataList: [
                {
                    userName: "张三",
                    userId: "1",
                },
                {
                    userName: "张三2",
                    userId: "2",
                },
                {
                    userName: "张三3",
                    userId: "3",
                },
            ], // @功能获取存储用户的数据
            ignoreUserIdList: [], // @功能存储用户的选中id集合
            filterUserList: [
                {
                    userName: "张三",
                    userId: "1",
                },
                {
                    userName: "张三2",
                    userId: "2",
                },
                {
                    userName: "张三3",
                    userId: "3",
                },
            ], // @列表
            isShowBox: false, // 是否显示@列表
            mentionBottom: 0, // @列表位置 bottom
            mentionLeft: 20, // @列表位置 left
            savedSelectionMention: null, // 保存选中位置
            savedSelectionRange: null, // 保存选中位置
            selectionName: false, // @列表筛选用户
            isRecording: false, // 是否正在记录
            // ... 其他功能的变量 ...
        };
    },

    computed: {
        rangeLength() {
            return (
                this.savedSelectionRange.length > 0 &&
                this.savedSelectionRange[0].startOffset -
                    this.savedSelectionMention[0].startOffset
            );
        },
    },
    methods: {
        /**
         * @函数描述: 输入框输入变化时的回调函数
         **/
        handleInput(event) {
            const commentBox = document.querySelector("#comment-input-data");
            // 显示隐藏评论按钮
            if (commentBox.innerHTML.length > 0) {
                this.showSendBtn = true;
                if (this.showDrawerDown && this.showEmoteOrFile === 2) {
                    this.showDrawerDown = false;
                }
            } else {
                this.showSendBtn = false;
            }
            if (this.isRecording) {
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
                    if (charBeforeCursor === "@") {
                        this.ignoreUserIdList = this.mentionCreateUserIdList(
                            commentBox.innerHTML
                        );
                        this.savedSelectionMention = this.saveSelection();

                        const rect = range.getClientRects()[0];
                        this.mentionBottom =
                            rect.top -
                            event.target.getBoundingClientRect().top +
                            30;
                        this.mentionLeft =
                            rect.left -
                            event.target.getBoundingClientRect().left +
                            20;
                        if (event.target.clientWidth - this.mentionLeft < 85) {
                            this.mentionLeft = this.mentionLeft - 200;
                        }
                        this.isRecording = true;
                        let data = {
                            ignoreUserIdList: this.ignoreUserIdList,
                            tagId: this.tagId,
                        };
                        // get_interaction_mentioning(data).then((res) => {
                        //     this.ignoreUserDataList = res.data.data;
                        //     this.filterUserList = res.data.data;
                        //     this.isShowBox = true;
                        // });
                        // this.filterUserList = [
                        //     {
                        //         userName: "张三",
                        //         userId: "1",
                        //     },
                        //     {
                        //         userName: "张三2",
                        //         userId: "2",
                        //     },
                        //     {
                        //         userName: "张三3",
                        //         userId: "3",
                        //     },
                        // ];
                    }
                }
            }
        },

        /**
         * @函数描述: 键盘输入时的回调函数 - 【弹起】
         **/
        handleKeyup() {
            const commentBox = document.querySelector("#comment-input-data");
            this.savedSelectionRange = this.saveSelection();
            if (commentBox.innerHTML === "") {
                this.isRecording = false;
            }
            if (this.isRecording) {
                if (this.rangeLength < 0) {
                    this.isShowBox = false;
                } else {
                    this.isShowBox = true;
                    this.filterSelectionName();
                }
            }
        },

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
            if (event.keyCode === 13) {
                //回车执行查询
                event.preventDefault();
                this.enterSendCommentOrReply(1);
            }
        },

        /**
         * @函数描述: 过滤输入框内的姓名
         **/
        filterSelectionName() {
            const commentBox = document.querySelector("#comment-input-data");
            this.selectionName = commentBox.innerHTML
                .replace(/<[^>]*>[^<]*<\/[^>]*>/g, function (match) {
                    return match.replace(/[^]*/g, ""); // 使用正则表达式替换标签的内容
                })
                .split("@")[1]
                .substr(0, this.rangeLength);

            this.filterUserList = this.ignoreUserDataList.filter(
                (item) => item.userName.indexOf(this.selectionName) > -1
            );
        },

        /**
         * 选择用户
         * @param row
         */
        selectUser(row) {
            const span = document.createElement("span");
            span.className = "mention-at";
            span.id = row.userId;
            span.textContent = `@${row.userName}`;
            span.setAttribute("contenteditable", "false");

            const commentBox = document.querySelector("#comment-input-data");

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
            let roleName = `(<span>${content.roleName}</span>)`;
            if (content.userName.indexOf(title) !== -1) {
                reg = content.userName.replace(title, dataToReplace) + roleName;
            } else {
                reg = content.userName + roleName;
            }
            return reg;
        },

        /**
         * @函数描述: 点击表情添加到输入框
         * @param: {Object} event 点击的对象
         **/
        insertEmoji(emoji) {
            const commentBox = document.querySelector("#comment-input-data");
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
        },
        enterSendCommentOrReply() {
            const commentBox = document.querySelector("#comment-input-data");
        },
    },
};
</script>

<style lang="scss" scoped>
.mention-box {
    position: absolute;
    min-width: 120px;
    max-width: 200px;
    padding: 10px;
    background: #fff;
    border-radius: 5px;
    border: 1px solid gainsboro;
    box-sizing: border-box;
    box-shadow: 0 0 5px rgba(0, 0, 0, 0.1);
    max-height: 245px;
    z-index: 9999999;
    font-size: 12px;

    & > div:first-child {
        font-weight: bold;
        padding: 0 0 5px 0;
    }

    .user-box {
        padding: 0;
        margin: 0;
        overflow: hidden;
        overflow-y: auto;
        max-height: 200px;

        &::-webkit-scrollbar {
            display: none;
        }

        div {
            padding: 5px 10px;
            cursor: pointer;
            display: flex;
            align-items: center;

            &:hover {
                background: #46a8ff;
                color: #fff;
            }

            .avatarImg {
                width: 20px;
                height: 20px;
                border-radius: 50%;
                margin-right: 5px;
                vertical-align: sub;
            }

            span {
                width: 86%;
                overflow: hidden;
                white-space: nowrap;
                text-overflow: ellipsis;
                display: inline-block;
            }
        }
    }
}

</style>
