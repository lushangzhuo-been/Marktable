<template>
    <!-- 进展 -->
    <div class="evolve-area">
        <div
            class="content-block"
            v-infinite-scroll="loadMoreEvolve"
            infinite-scroll-delay="400"
            infinite-scroll-distance="30"
            :class="{ 'prop-up': propUp }"
        >
            <div
                v-for="(evolveItem, evolveIndex) in evolve"
                :key="evolveIndex"
                class="evolve-item"
            >
                <div class="info-block">
                    <div class="left">
                        <el-avatar
                            class="progress-avatar"
                            icon="el-icon-user-solid"
                            size="small"
                            :src="getAvatar(evolveItem.creator.avatar)"
                        ></el-avatar>
                        {{ evolveItem.creator.full_name }}
                    </div>
                    <div class="right">
                        <div>
                            {{ evolveItem.created_at }}
                        </div>
                        <div class="stick"></div>
                        <!-- 更多操作 -->
                        <div class="more-operation">
                            <el-dropdown
                                placement="bottom-end"
                                trigger="click"
                                class="operation-block-dropdown basic-ui"
                                width="200"
                                @command="
                                    (command) =>
                                        evolveOperation(command, evolveItem)
                                "
                                @visible-change="
                                    (boolean) =>
                                        evolveDropShow(boolean, evolveItem)
                                "
                            >
                                <el-dropdown-menu slot="dropdown">
                                    <el-dropdown-item
                                        class="basic-ui height32"
                                        command="edit"
                                        :disabled="
                                            evolveItem.permission !== 'yes'
                                        "
                                    >
                                        <div
                                            :class="{
                                                disabled:
                                                    evolveItem.permission !==
                                                    'yes'
                                            }"
                                        >
                                            <b
                                                class="operation-item-box edit"
                                            ></b>
                                            编辑
                                        </div>
                                    </el-dropdown-item>
                                    <el-dropdown-item
                                        class="basic-ui height32"
                                        command="delete"
                                        :disabled="
                                            evolveItem.permission !== 'yes'
                                        "
                                    >
                                        <div
                                            :class="{
                                                disabled:
                                                    evolveItem.permission !==
                                                    'yes'
                                            }"
                                        >
                                            <b
                                                class="operation-item-box delete"
                                            ></b>
                                            删除
                                        </div>
                                    </el-dropdown-item>
                                </el-dropdown-menu>
                                <b
                                    class="tabs-partition"
                                    :class="{ active: evolveItem.operating }"
                                ></b>
                            </el-dropdown>
                        </div>
                    </div>
                </div>
                <div class="msg-block">
                    <!-- 编辑模式 -->
                    <at-mention
                        v-if="evolveItem.editing"
                        :id="`${evolveIndex}`"
                        cancleWithSave
                        placeholder="请输入进展或@责任人"
                        :msgObj="evolveItem"
                        v-model="evolveItem.content"
                        @cancle-content="cancleUpdateEvolve"
                        @confirm-content="updateEvolve"
                    ></at-mention>
                    <!-- 详情模式 -->
                    <div v-else>
                        <span v-html="evolveItem.content"></span>
                    </div>
                </div>
                <div v-if="false" class="main-reply-block">
                    <div @click="replyToFloor(evolveItem, evolveIndex)">
                        回复
                    </div>
                </div>
                <div v-if="false" class="reply-content">
                    <div
                        v-for="(replyItem, replyIndex) in evolveItem.reply"
                        :key="replyIndex"
                        class="floor-item"
                    >
                        <div class="info-block">
                            <div class="left">
                                <el-avatar
                                    class="progress-avatar"
                                    icon="el-icon-user-solid"
                                    size="small"
                                    :src="getAvatar(evolveItem.creator.avatar)"
                                ></el-avatar>
                                {{ replyItem.name }}
                            </div>
                            <div class="right">
                                {{ replyItem.time }}
                            </div>
                        </div>
                        <div class="msg-block-in-floor">
                            {{ replyItem.content }}
                        </div>
                    </div>
                    <!-- 楼中楼回复按钮 -->
                    <div
                        v-show="evolveItem.reply.length"
                        class="floor-reply-clock"
                    >
                        <el-button
                            type="text"
                            class="basic-ui"
                            size="small"
                            @click="replyToFloor(evolveItem, evolveIndex)"
                        >
                            回复
                        </el-button>
                    </div>
                    <div
                        v-show="evolveItem.replying"
                        class="floor-reply-block"
                        :class="{ 'empty-content': !evolveItem.reply.length }"
                    >
                        <el-input
                            :ref="'FloorReply' + evolveIndex"
                            class="basic-ui"
                            size="small"
                            v-model="evolveItem.replyingContent"
                            placeholder="请输入进展或@责任人"
                            @blur="floorInputBlur(evolveItem)"
                        ></el-input>
                        <div class="floor-btn-block">
                            <el-button
                                v-show="evolveItem.replyingContent"
                                type="primary"
                                class="basic-ui"
                                size="small"
                                @click="
                                    replyFloor(
                                        evolveItem.reply,
                                        evolveItem.replyingContent,
                                        evolveItem
                                    )
                                "
                            >
                                发送</el-button
                            >
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <!-- 吸底分页器 和输入框 -->
        <div class="bottom-block">
            <div class="input-block">
                <at-mention
                    ref="atMention"
                    confirmBtn
                    mentionListPosition="top"
                    placeholder="请输入进展或@责任人"
                    @confirm-content="addEvolve"
                    @get-propUp="getPropUp"
                    v-model="mainReplyText"
                ></at-mention>
            </div>
        </div>
    </div>
</template>

<script>
import atSomeone from "@/components/at_Someone/index";
import atMention from "@/components/at_mention/index";
import { DefauleAvatar } from "@/assets/tool/const";
import api from "@/common/api/module/progress";
import { imgHost } from "@/assets/tool/const";
export default {
    components: {
        atSomeone,
        atMention
    },
    props: {
        detailId: {
            type: String,
            default: ""
        }
    },
    data() {
        return {
            mainReplyText: "",
            evolve: [],
            list: {
                page: 1,
                size: 10,
                count: 0
            },
            propUp: false
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
    mounted() {},
    methods: {
        loadMoreEvolve() {
            // 计算是否为最后一页
            let lastPage = Math.ceil(this.list.count / this.list.size);
            if (this.list.page >= lastPage) {
                return;
            }
            this.list.page += 1;
            this.getEvolveDetail();
        },
        getPropUp(propUp) {
            this.propUp = propUp;
        },
        currentChange(page) {
            this.list.page = page;
            this.getEvolveDetail();
        },
        // 获取数据
        getEvolveDetail(order) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                issue_id: this.detailId,
                page_num: this.list.page,
                page_size: this.list.size
            };
            // 新增或者删除时
            if (order === "refresh") {
                params.page_num = 1;
                params.page_size = this.list.page * this.list.size;
            }
            api.getEvolveDetail(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    if (order === "refresh") {
                        // 删除后 pagesize*pagenum
                        this.evolve = [];
                    }
                    // 触底加载逻辑
                    if (res.data.list && res.data.list.length) {
                        this.evolve = [...this.evolve, ...res.data.list];
                    }
                    this.list.count = res.data.cnt || 0;
                } else {
                    this.evolve = [];
                    this.list.count = 0;
                }
                this.$emit("transmit-evolve-count", this.list.count);
            });
        },
        // 新增进展
        addEvolve(inner) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                issue_id: this.detailId,
                content: this.mainReplyText
            };
            api.addEvolve(params).then((res) => {
                this.mainReplyText = "";
                this.$refs.atMention.reset();
                this.getEvolveDetail("refresh");
            });
        },
        // 编辑进展
        updateEvolve(inner, evolve) {
            if (evolve.permission !== "yes") {
                return;
            }
            let params = {
                id: evolve._id,
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                issue_id: this.detailId,
                content: evolve.content
            };
            api.updateEvolve(params).then((res) => {
                this.getEvolveDetail("refresh");
                this.$set(evolve, "editing", false);
            });
        },
        // 编辑进展
        // editEvolve(evolve) {
        //     if (evolve.permission !== "yes") {
        //         return;
        //     }
        //     let params = {
        //         id: evolve._id,
        //         ws_id: this.curSpace.id,
        //         tmpl_id: this.curProgress,
        //         issue_id: this.detailId,
        //         content: evolve.content
        //     };
        //     api.updateEvolve(params).then((res) => {
        //         this.getEvolveDetail();
        //     });
        // },
        // 取消编辑
        cancleUpdateEvolve(evolve) {
            this.$set(evolve, "editing", false);
            this.$set(evolve, "content", evolve.content_backup);
        },
        cancleEditEvolve(evolve) {
            this.$set(evolve, "editing", false);
            this.$set(evolve, "content", evolve.content_backup);
        },
        // 删除进展
        deleteEvolve(evolve) {
            if (evolve.permission !== "yes") {
                return;
            }
            let params = {
                id: evolve._id,
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                issue_id: this.detailId
            };
            api.deleteEvolve(params).then((res) => {
                this.getEvolveDetail("refresh");
            });
        },
        replyToFloor(evolveItem, evolveIndex) {
            this.$set(evolveItem, "replying", true);
            this.$set(evolveItem, "replyingContent", "");
            this.$nextTick(() => {
                this.$refs[`FloorReply` + evolveIndex][0].focus();
            });
        },
        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        },
        floorInputBlur(evolveItem) {
            if (!evolveItem.replyingContent) {
                evolveItem.replying = false;
            }
        },
        // 清空页面进展显示
        clearEvolveShow() {
            this.mainReplyText = "";
            this.$refs.atMention.reset();
            this.evolve = [];
        },
        evolveOperation(command, evolve) {
            if (command === "edit") {
                this.$set(evolve, "editing", true);
                this.$set(evolve, "content_backup", evolve.content);
            }
            if (command === "delete") {
                this.deleteEvolve(evolve);
            }
        },
        evolveDropShow(boolean, evolve) {
            if (boolean) {
                this.$set(evolve, "operating", true);
            } else {
                this.$set(evolve, "operating", false);
            }
        }
    }
};
</script>

<style lang="scss" scoped>
.input-block {
    .btns-group {
        display: flex;
        justify-content: right;
        margin-top: 12px;
    }
}
.evolve-area {
    width: 100%;
    height: 100%;
    position: relative;
    .content-block {
        padding: 0 16px;
        height: calc(100% - 58px);
        overflow-y: auto;
        &.prop-up {
            height: calc(100% - 148px);
        }
    }
    .bottom-block {
        width: calc(100% - 32px);
        padding: 16px 16px 0 16px;
        border-top: 1px solid #e6e9f0;
        position: absolute;
        bottom: 0;
        .basic-ui.el-pagination {
            margin: 4px 0;
        }
    }
}
.avatar-box {
    display: inline-block;
    width: 32px;
    height: 32px;
    border-radius: 50%;
    background-size: 100% 100%;
    vertical-align: middle;
    margin-right: 6px;
}

.evolve-item {
    padding: 12px;
    margin-bottom: 12px;
    border-radius: 4px;
    background-color: #fafafb;
}
.info-block {
    // padding-left: 20px;
    height: 32px;
    line-height: 32px;
    display: flex;
    justify-content: space-between;
    .left {
        font-weight: 500;
        font-size: 14px;
        color: #2f384c;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        .progress-avatar.el-avatar {
            width: 28px;
            height: 28px;
            line-height: 1;
            vertical-align: middle;
            margin-right: 8px;
        }
    }
    .right {
        font-weight: 400;
        font-size: 14px;
        color: #95979d;
        display: flex;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }
}
.msg-block {
    // padding: 0 6px 0 44px;
    padding: 0 6px 0 0px;
    font-weight: 400;
    font-size: 14px;
    color: #2f384c;
    margin-top: 12px;
    line-height: 1.5;
    word-break: break-word;
}
.floor-item {
    padding-top: 12px;
}
.msg-block-in-floor {
    padding: 0 20px 0 62px;
    font-weight: 400;
    font-size: 14px;
    color: #2f384c;
    margin-top: 12px;
}
.main-reply-block {
    height: 30px;
    line-height: 30px;
    border-top: 1px solid #e9e9e9;
    border-bottom: 1px solid #e9e9e9;
    display: flex;
    justify-content: center;
}
.floor-reply-clock {
    padding-left: 62px;
    margin-bottom: 6px;
}
.floor-btn-block {
    display: flex;
    justify-content: right;
    margin-top: 12px;
    margin-bottom: 12px;
}
.floor-reply-block {
    padding: 0 30px;
    &.empty-content {
        margin-top: 12px;
    }
}
.stick {
    display: inline-block;
    width: 0px;
    height: 14px;
    margin: auto 7px;
}
.more-operation {
    .tabs-partition {
        display: inline-block;
        width: 20px;
        height: 20px;
        vertical-align: sub;
        background-image: url(@/assets/image/progress/operation.png);
        background-size: 100% 100%;
        &:hover {
            background-image: url(@/assets/image/progress/operation-active.png);
        }
        &.active {
            background-image: url(@/assets/image/progress/operation-active.png);
        }
    }
}

.operation-item-box {
    display: inline-block;
    width: 20px;
    height: 20px;
    background-size: 100% 100%;
    vertical-align: sub;
    &.edit {
        background-image: url(@/assets/image/common/edit.svg);
    }
    &.delete {
        background-image: url(@/assets/image/common/delete.svg);
    }
}
.evolve-btn-group {
    display: flex;
    justify-content: flex-end;
    margin-top: 12px;
}
</style>
