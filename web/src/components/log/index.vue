<template>
    <div class="no-data" v-if="isShowNoData">
        <no-data :isTextShow="false" :loading="loading" :imgShow="imgShow">
        </no-data>
    </div>
    <div v-else>
        <el-collapse class="aui-collapse" v-model="activeName">
            <el-collapse-item
                class="aui-collapse-item"
                v-for="(item, index) in logs"
                :key="index"
                :name="index"
            >
                <div slot="title" class="title-block">
                    <!-- 日志用户信息 -->
                    <div class="title">
                        <div class="left">
                            <el-avatar
                                icon="el-icon-user-solid"
                                size="small"
                                class="logs-avatar"
                                :src="getAvatar(item.creator.avatar)"
                            ></el-avatar>
                            <div class="full-name" v-overflow>
                                {{
                                    item.creator.full_name +
                                    " " +
                                    item.action_desc
                                }}
                            </div>
                            <b class="state-icon" :class="getState(index)"></b>
                        </div>
                        <div class="right">
                            {{ item.created_at }}
                        </div>
                    </div>
                    <div v-if="item.suggest" class="suggest-block" @click.stop>
                        <div class="suggest-block-decorate">
                            意见：{{ item.suggest }}
                        </div>
                    </div>
                </div>

                <div class="log-content">
                    <div
                        v-for="(logItem, logIndex) in item.content"
                        :key="logIndex"
                        class="log-item"
                        v-html="logItem"
                    ></div>
                </div>
            </el-collapse-item>
        </el-collapse>
        <pagination
            v-show="logsPage.count > 10"
            layout="total, prev, pager, next"
            :total="logsPage.count"
            @curPageChange="currentChange"
        ></pagination>
    </div>
</template>

<script>
import Pagination from "@/pages/common/pagination.vue";
import NoData from "@/pages/common/no_data.vue";
import { imgHost } from "@/assets/tool/const";
import api from "@/common/api/module/progress";
export default {
    components: { Pagination, NoData },
    props: {
        detailId: {
            type: String,
            default: ""
        }
    },
    data() {
        return {
            activeName: [],
            logs: [],
            logsPage: {
                size: 10,
                page: 1,
                count: 0
            },
            isShowNoData: true,
            loading: true,
            imgShow: false
        };
    },
    mounted() {},
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        }
    },
    watch: {},
    methods: {
        getState(index) {
            if (this.activeName.indexOf(index) > -1) {
                return "expand";
            } else {
                return "close";
            }
        },
        getAvatar(avatar) {
            if (avatar) {
                return `${imgHost}${avatar}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        },
        currentChange(page) {
            this.logsPage.page = page;
            this.activeName = [];
            this.getLog();
        },
        getLog() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                issue_id: this.detailId,
                page_num: this.logsPage.page,
                page_size: this.logsPage.size
            };
            api.getProgressLog(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length &&
                    res.data.list &&
                    res.data.list.length
                ) {
                    this.logs = res.data.list;
                    this.logsPage.count = res.data.cnt;
                    this.isShowNoData = false;
                } else {
                    this.logs = [];
                    this.logsPage.count = 0;
                    this.isShowNoData = true;
                }
                this.loading = false;
                setTimeout(() => {
                    this.imgShow = true;
                }, 50);
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.logs-avatar {
    vertical-align: middle;
    min-width: 28px;
}
.avatar-box {
    display: inline-block;
    width: 28px;
    height: 28px;
    border-radius: 50%;
    background-size: 100% 100%;
    vertical-align: middle;
    margin-right: 6px;
}
.log-content {
    padding-left: 38px;
}
.state-icon {
    display: inline-block;
    width: 20px;
    height: 20px;
    margin-left: 6px;
    vertical-align: middle;
    &.expand {
        background-image: url("~@/assets/image/progress_detail/expand.png");
        background-size: 100% 100%;
    }
    &.close {
        background-image: url("~@/assets/image/progress_detail/close.png");
        background-size: 100% 100%;
    }
}
.title-block {
    .title {
        display: flex;
        justify-content: space-between;
        align-items: center;
        height: 30px;
        line-height: 30px;
        .left {
            flex: 1;
            max-width: calc(100% - 146px);
            white-space: nowrap;
            display: flex;
            align-items: center;
            font-weight: 500;
            font-size: 14px;
            color: #2f384c;
            .full-name {
                display: inline-block;
                white-space: nowrap;
                overflow: hidden;
                text-overflow: ellipsis;
                font-size: 14px;
                margin: 0 4px 0 12px;
            }
        }
        .right {
            white-space: nowrap;
            font-weight: 400;
            font-size: 14px;
            color: #95979d;
        }
    }
}
.log-item {
    font-weight: 400;
    font-size: 14px;
    color: #95979d;
    line-height: 22px;
    white-space: pre-line;
    ::v-deep img {
        max-width: 100%;
        height: auto;
    }
}
.suggest-block {
    box-sizing: border-box;
    padding-left: 38px;
    .suggest-block-decorate {
        padding: 10px;
        background: #f8f8f8;
        border-radius: 4px;
        font-weight: 400;
        font-size: 14px;
        color: #171e31;
        line-height: 18px;
        margin-top: 6px;
    }
}
.no-data {
    height: 220px;
}
</style>
<style lang="scss">
.el-collapse.aui-collapse {
    border-top: none;
    border-bottom: none;
    .el-collapse-item.aui-collapse-item {
        margin-bottom: 24px;
        &:last-child {
            margin-bottom: 0px;
        }
        .el-collapse-item__header {
            display: block;
            border-bottom: none;
            height: auto;
            width: 100%;
            .el-collapse-item__arrow {
                display: none;
            }
        }
        .el-collapse-item__wrap {
            border-bottom: none;
            .el-collapse-item__content {
                margin-top: 12px;
                padding-bottom: 0;
            }
        }
    }
}
</style>
