<template>
    <div>
        <!-- 传入一个userlist -->
        <div class="mem-list-content" v-if="frontArr.length">
            <div
                ref="listCon"
                class="tag-list"
                :class="{ 'show-num': showNum }"
            >
                <div
                    class="tag-item"
                    v-for="(tagItem, tagIndex) in frontArr"
                    :key="tagIndex"
                >
                    <user-message :userMessage="tagItem">
                        <span>
                            <el-avatar
                                class="progress-avatar"
                                icon="el-icon-user-solid"
                                size="small"
                                :src="getAvatar(tagItem.avatar)"
                                :style="
                                    getAvatarStack(tagItem.avatar, tagIndex)
                                "
                            ></el-avatar>
                            {{ tagItem.full_name }}
                        </span>
                    </user-message>
                </div>
            </div>
            <!-- 数字气泡 -->
            <el-tooltip
                v-show="showNum"
                class="item"
                effect="dark"
                placement="top"
            >
                <div slot="content">
                    <div
                        v-for="(tagItem, tagIndex) in behandArr"
                        :key="tagIndex"
                        class="member-list"
                    >
                        <user-message :userMessage="tagItem">
                            <span>
                                <el-avatar
                                    class="progress-avatar"
                                    icon="el-icon-user-solid"
                                    size="small"
                                    :src="getAvatar(tagItem.avatar)"
                                    :style="
                                        getAvatarStack(tagItem.avatar, tagIndex)
                                    "
                                ></el-avatar>
                                {{ tagItem.full_name }}
                            </span>
                        </user-message>
                    </div>
                </div>
                <b ref="num-box" class="num-box">+{{ behandArr.length }}</b>
            </el-tooltip>
        </div>
        <div v-show="showAny" class="show-any">任意值</div>
    </div>
</template>

<script>
import _ from "lodash";
import { imgHost } from "@/assets/tool/const";
import UserMessage from "@/components/user_message_tip";

export default {
    props: {
        userIdList: {
            type: String,
            default: ""
        },
        userInfoList: {
            type: Array,
            default: () => []
        }
    },
    components: {
        UserMessage
    },
    data() {
        return {
            showNum: false,
            personList: [],
            frontArr: [],
            behandArr: [],
            showAny: false
        };
    },
    watch: {
        userIdList: {
            handler(str) {
                if (str === "任意值") {
                    this.showAny = true;
                } else {
                    this.showAny = false;
                }
            },
            deep: true,
            immediate: true
        },
        userInfoList: {
            handler(arr) {
                if (arr) {
                    this.personList = arr;
                    this.frontArr = arr;
                    this.$nextTick(() => {
                        this.getTagInit();
                    });
                } else {
                    this.personList = [];
                    this.frontArr = [];
                }
            },
            deep: true,
            immediate: true
        }
    },
    methods: {
        getTagInit() {
            const listCon = this.$refs.listCon;
            if (listCon) {
                const labels = listCon.querySelectorAll(".tag-item");
                let labelIndex = 0; // 渲染到第几个
                const listConBottom = listCon.getBoundingClientRect().bottom; // 容器底部距视口顶部距离
                this.showNum = true;
                this.$nextTick(() => {
                    for (let i = 0; i < labels.length; i++) {
                        const _top = labels[i].getBoundingClientRect().top;
                        if (_top >= listConBottom) {
                            // 如果有标签顶部距离超过容器底部则表示超出容器隐藏
                            this.showNum = true;
                            labelIndex = i;
                            this.getShowLabel(labelIndex);
                            break;
                        } else {
                            this.getAllLabel();
                            this.showNum = false;
                        }
                    }
                });
            }
        },
        getShowLabel(labelIndex) {
            this.labelIndex = labelIndex;
            this.frontArr = this.getArrFront(this.personList);
            this.behandArr = this.getArrBehand(this.personList);
        },
        getAllLabel() {
            this.frontArr = this.personList;
            this.behandArr = [];
        },
        getArrFront(arr) {
            let deepClone = _.cloneDeep(arr);
            let front = deepClone.splice(0, this.labelIndex);
            return front;
        },
        getArrBehand(arr) {
            let deepClone = _.cloneDeep(arr);
            let behand = deepClone.splice(this.labelIndex);
            return behand;
        },
        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        },
        getAvatarStack() {
            return {
                position: "relative",
                top: "-3px"
            };
        }
    }
};
</script>

<style lang="scss" scoped>
.mem-list-content {
    display: flex;
    margin: 8px 0;
    .tag-list {
        display: inline-block;
        height: 20px;
        .tag-item {
            display: inline-block;
            margin-right: 8px;
            overflow: hidden;
            text-wrap: nowrap;
            text-overflow: ellipsis;
            max-width: 148px;
            height: 20px;
        }
    }
}
.member-list {
    margin: 6px 0;
    &:first-child {
        margin-top: 0;
    }
    &:last-child {
        margin-bottom: 0;
    }
}
.num-box {
    box-sizing: border-box;
    display: inline-block;
    min-width: 22px;
    height: 22px;
    line-height: 22px;
    border-radius: 11px;
    padding: 0 2px;
    background-color: #f8f8f8;
    text-align: center;
    vertical-align: top;
    font-weight: 400;
    font-size: 12px;
    color: #2f384c;
    margin: auto 0;
}
.show-any {
    height: 40px;
    line-height: 40px;
    color: #171e31;
}
</style>
