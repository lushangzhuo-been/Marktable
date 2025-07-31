<template>
    <div class="page-content">
        <!-- 标题&tab 吸顶 -->
        <div class="page-head">
            <div class="page-title">我的消息</div>
            <div>
                <el-tabs class="basic-ui" v-model="activeName">
                    <el-tab-pane
                        v-for="(item, index) in tabsList"
                        :label="item.label"
                        :name="item.name"
                        :key="index"
                    >
                        <span slot="label" class="table-pane-item">
                            {{ item.label }}
                            <!-- <span
                                class="num-tip"
                                >{{item.name === 'newest' ? 1:2 }}</span
                            > -->
                            <span
                                class="num-tip"
                                v-if="
                                    item.name === 'newest' && lastestUnread > 0
                                "
                                >{{ lastestUnread }}</span
                            >
                            <span
                                class="num-tip"
                                v-if="item.name === '@i' && atListUnread > 0"
                                >{{ atListUnread }}</span
                            >
                        </span>
                    </el-tab-pane>
                </el-tabs>
            </div>
        </div>
        <!-- 列表 -->
        <div class="page-body">
            <div class="list-content">
                <div v-show="isShowMarkRead" class="not-read-msg">
                    <span class="mark-read" @click="markAllRead"
                        >全部标记已读</span
                    >
                </div>
                <!-- 最新 -->
                <div v-if="activeName === 'newest'">
                    <div
                        class="list-item"
                        v-for="(item, index) in lastestList"
                        :key="index"
                    >
                        <!-- 信息 & 描述 & 事件 -->
                        <div class="item-msg-head">
                            <div class="msg-left">
                                <el-avatar
                                    class="progress-avatar"
                                    icon="el-icon-user-solid"
                                    size="small"
                                    :src="getAvatar(item.origin_user.avatar)"
                                ></el-avatar>
                                <div class="name">
                                    {{ item.origin_user.full_name }}
                                </div>
                                <div
                                    class="title"
                                    :class="{ undo: item.read_sign === 'no' }"
                                >
                                    给你分配了任务
                                </div>
                            </div>
                            <div class="msg-right">{{ item.created_at }}</div>
                        </div>
                        <div class="item-msg-body">
                            <div
                                class="task-title"
                                @click="jumpToDetail(item)"
                                v-overflow
                            >
                                <!-- 解json item.content  取issue_title -->
                                {{ JSON.parse(item.content).issue_title }}
                            </div>
                            <div class="task-detail">
                                <div class="detail-item">
                                    <!-- 解json item.content  拼ws_name > tmpl_name -->
                                    {{
                                        `${
                                            JSON.parse(item.content).ws_name
                                        } > ${
                                            JSON.parse(item.content).tmpl_name
                                        }`
                                    }}
                                </div>
                                <!-- <div class="detail-item">创建人:xxx</div>
                                <div class="detail-item">
                                    创建时间:2024-08-06 15:23:32
                                </div> -->
                            </div>
                        </div>
                    </div>
                </div>
                <!-- @我 -->
                <div v-if="activeName === '@i'">
                    <div
                        class="list-item"
                        v-for="(item, index) in atList"
                        :key="index"
                    >
                        <!-- 信息 & 描述 & 事件 -->
                        <div class="item-msg-head">
                            <div class="msg-left">
                                <el-avatar
                                    class="progress-avatar"
                                    icon="el-icon-user-solid"
                                    size="small"
                                    :src="getAvatar(item.origin_user.avatar)"
                                ></el-avatar>
                                <div class="name">
                                    {{ item.origin_user.full_name }}
                                </div>
                                <div
                                    class="title"
                                    :class="{ undo: item.read_sign === 'no' }"
                                >
                                    @我
                                </div>
                            </div>
                            <div class="msg-right">{{ item.created_at }}</div>
                        </div>
                        <!-- v-html 解json item.content 取 progress_content -->
                        <!-- <div class="evolve" v-overflow>这个问题你看一下</div> -->
                        <div
                            class="evolve"
                            v-html="JSON.parse(item.content).progress_content"
                        ></div>
                        <div class="item-msg-body">
                            <div
                                class="task-title"
                                @click="jumpToDetail(item)"
                                v-overflow
                            >
                                {{ JSON.parse(item.content).issue_title }}
                            </div>
                            <div class="task-detail">
                                <div class="detail-item">
                                    {{
                                        `${
                                            JSON.parse(item.content).ws_name
                                        } > ${
                                            JSON.parse(item.content).tmpl_name
                                        }`
                                    }}
                                </div>
                                <!-- <div class="detail-item">创建人:xxx</div>
                                <div class="detail-item">
                                    创建时间:2024-08-06 15:23:32
                                </div> -->
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <!-- 无数据 -->
            <no-data
                v-if="
                    (activeName === 'newest' && !lastestList.length) ||
                    (activeName === '@i' && !atList.length)
                "
                :loading="loading"
                :imgShow="imgShow"
            ></no-data>
        </div>
        <!-- 分页吸底 -->
        <div class="page-foot">
            <el-pagination
                v-show="page.count > 20"
                class="basic-ui"
                @current-change="currentChange"
                @size-change="sizeChange"
                layout="total,sizes, prev, pager, next"
                :current-page="page.pageNum"
                :page-size="page.size"
                :total="page.count"
            >
            </el-pagination>
        </div>
    </div>
</template>

<script>
import { imgHost } from "@/assets/tool/const";
import api from "@/common/api/module/progress";
import NoData from "@/pages/common/no_data.vue";
export default {
    components: { NoData },
    data() {
        return {
            title: "登录页",
            activeName: "newest",
            tabsList: [
                {
                    label: "最新",
                    name: "newest"
                },
                {
                    label: "@我",
                    name: "@i"
                }
            ],
            page: {
                count: 40,
                pageNum: 1,
                size: 20,
                pageSize: [20, 40, 100]
            },
            lastestList: [],
            atList: [],
            loading: true,
            imgShow: false,
            lastestUnread: 0,
            atListUnread: 0
        };
    },
    computed: {
        isShowMarkRead() {
            return this.activeName === "newest"
                ? this.lastestUnread
                    ? true
                    : false
                : this.atListUnread
                ? true
                : false;
        }
    },
    watch: {
        activeName: {
            handler(tab) {
                this.page.pageNum = 1;
                this.getMessageList(tab);
            }
        }
    },
    mounted() {
        // 初始两类型都请求，获取未读数量，之后每次切换tab刷新
        this.getMessageList("newest");
        this.getMessageList("@i");
        this.$emit("refresh-unread");
    },
    methods: {
        // 标记为全部已读
        markAllRead() {
            let mode =
                this.activeName === "newest"
                    ? "reassign_mode"
                    : this.activeName === "@i"
                    ? "at_person_mode"
                    : "";
            api.markAllReadRead({ mode }).then((res) => {
                if (res && res.resultCode === 200) {
                    this.$emit("refresh-unread");
                    this.getMessageList(this.activeName);
                }
            });
        },
        jumpToDetail(item) {
            // json解item.content 取 issue_url
            if (
                JSON.parse(item.content) &&
                JSON.parse(item.content).issue_url
            ) {
                let url = JSON.parse(item.content).issue_url;
                let split = url.split("/#/");
                let origin = window.location.origin;
                let newUrl = origin + "/#/" + split[1];
                window.open(newUrl, "_blank");
                // 调已读接口
                this.hadReadMessage(item);
            }
        },
        getMessageList(tab) {
            this.loading = true;
            let params = {
                mode: "",
                page_num: this.page.pageNum,
                page_size: this.page.size
            };
            if (tab === "newest") {
                params.mode = "reassign_mode";
            }
            if (tab === "@i") {
                params.mode = "at_person_mode";
            }
            api.getMessageList(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    Object.keys(res.data).length
                ) {
                    if (tab === "newest") {
                        this.lastestList = res.data.list || [];
                        this.lastestUnread = res.data.un_read_cnt || 0;
                    }
                    if (tab === "@i") {
                        this.atList = res.data.list || [];
                        this.atListUnread = res.data.un_read_cnt || 0;
                    }
                    this.page.count = res.data.cnt;
                } else {
                    this.lastestList = [];
                    this.atList = [];
                    this.page.count = 0;
                }
                this.loading = false;
                if (!this.page.count) {
                    this.imgShow = true;
                }
            });
        },
        hadReadMessage(item) {
            let params = {
                id: item._id
            };
            api.hadReadMessage(params).then((res) => {
                if (res && res.resultCode === 200) {
                    // 调列表刷新，向上请求刷新
                    this.$emit("refresh-unread");
                    // 手动减少未读消息数量
                    item.read_sign = "yes";
                }
                // if (this.activeName === "newest") {
                //     this.lastestUnread -= 1;
                // }
                // if (this.activeName === "@i") {
                //     this.atListUnread -= 1;
                // }
                this.getMessageList(this.activeName);
            });
        },
        currentChange(page) {
            this.page.pageNum = page;
            this.getMessageList(this.activeName);
        },
        sizeChange(size) {
            this.page.pageNum = 1;
            this.page.size = size;
            this.getMessageList(this.activeName);
        },
        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        }
    }
};
</script>

<style lang="scss" scoped>
.table-pane-item {
    position: relative;
    .num-tip {
        display: inline-block;
        box-sizing: border-box;
        background-color: #ff3c3d;
        min-width: 18px;
        height: 18px;
        font-size: 12px;
        color: #fff;
        line-height: 18px;
        padding: 0 2px;
        border-radius: 9px;
        text-align: center;
        position: absolute;
        top: -6px;
        right: -14px;
    }
}
.progress-avatar.el-avatar {
    width: 28px;
    height: 28px;
    line-height: 1;
    vertical-align: middle;
}
.page-content {
    height: calc(100% - 44px);
}
.page-head {
    .page-title {
        font-family: MiSans;
        font-weight: 500;
        font-size: 20px;
        color: #171e31;
        margin-bottom: 4px;
    }
}
.page-body {
    height: calc(100% - 100px);
    overflow-y: auto;
    .list-content {
        width: 720px;
        margin: 0 auto;
        .not-read-msg {
            display: flex;
            align-items: center;
            justify-content: flex-end;
            padding: 4px 0 10px 0;
            margin-bottom: 8px;
            border-bottom: 1px solid #e4e7ed;
            font-weight: 700;
            .mark-read {
                display: inline-block;
                height: 24px;
                line-height: 24px;
                margin-left: 16px;
                padding: 0 8px;
                border-radius: 4px;
                color: #1890ff;
                font-weight: 400;
                cursor: pointer;
                &:hover {
                    background-color: #f0f8ff;
                }
            }
        }
        .list-item {
            padding: 12px 0;
            border-bottom: 1px solid #f7f7f7;
            .item-msg-head {
                height: 28px;
                line-height: 28px;
                display: flex;
                justify-content: space-between;
                margin-bottom: 8px;
                .msg-left {
                    display: flex;
                    .name {
                        margin-left: 12px;
                        color: #66677c;
                    }
                    .title {
                        margin-left: 12px;
                        color: #171e31;
                        max-width: 340px;
                        position: relative;
                        &.undo {
                            &:after {
                                content: "";
                                display: inline-block;
                                width: 8px;
                                height: 8px;
                                background: #ff3c3d;
                                border-radius: 50%;
                                position: absolute;
                                right: -10px;
                                top: 6px;
                            }
                        }
                    }
                }
                .msg-right {
                    color: #66677c;
                }
            }
            .evolve {
                color: #171e31;
                margin-left: 40px;
                margin-bottom: 12px;
                overflow: hidden;
                white-space: nowrap;
                text-overflow: ellipsis;
            }
            .item-msg-body {
                margin-left: 40px;
                padding: 12px;
                background: #fafafb;
                border-radius: 4px 4px 4px 4px;
                .task-title {
                    height: 24px;
                    line-height: 24px;
                    color: #1890ff;
                    width: 100%;
                    overflow: hidden;
                    white-space: nowrap;
                    text-overflow: ellipsis;
                    cursor: pointer;
                    &:hover {
                        text-decoration: underline;
                    }
                }
                .task-detail {
                    display: flex;
                    .detail-item {
                        height: 18px;
                        line-height: 18px;
                        font-size: 12px;
                        color: #82889e;
                        padding: 0 12px;
                        position: relative;
                        &:after {
                            content: "";
                            margin: 2px 0;
                            height: 10px;
                            border-right: 1px solid #82889e;
                            display: inline-block;
                            position: absolute;
                            right: 0;
                        }
                        &:first-child {
                            padding-left: 0;
                        }
                        &:last-child {
                            &:after {
                                border-right: none;
                            }
                        }
                    }
                }
            }
        }
    }
}
</style>
