<template>
    <div>
        <div class="mem-list-content">
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
                                :style="getAvatarStack()"
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
                                    :style="getAvatarStack()"
                                ></el-avatar>
                                {{ tagItem.full_name }}
                            </span>
                        </user-message>
                    </div>
                </div>
                <b ref="num-box" class="num-box">+{{ behandArr.length }}</b>
            </el-tooltip>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import { imgHost } from "@/assets/tool/const";
import UserMessage from "@/components/user_message_tip";

export default {
    props: {
        userInfo: {
            type: Array,
            default: () => []
        }
    },
    components: {
        UserMessage
    },
    data() {
        return {
            frontArr: [],
            behandArr: [],
            personList: [],
            labelIndex: 0,
            showNum: true
        };
    },
    watch: {
        userInfo: {
            handler(arr) {
                this.personList = _.cloneDeep(arr);
                this.frontArr = this.personList;
                this.$nextTick(() => {
                    this.getTagInit();
                });
            },
            immediate: true,
            deep: true
        }
    },
    methods: {
        getShowLabel(labelIndex) {
            this.labelIndex = labelIndex;
            this.frontArr = this.getArrFront(this.personList);
            this.behandArr = this.getArrBehand(this.personList);
        },
        getAllLabel() {
            this.frontArr = this.personList;
            this.behandArr = [];
        },
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
            } else {
                return require(`@/assets/image/common/default_avatar_big.png`);
            }
        },
        getAvatarStack() {
            return {
                position: "relative",
                top: "-2px"
            };
        }
    }
};
</script>

<style lang="scss" scoped>
.mem-list-content {
    display: flex;
    height: 40px;
    padding: 0 10px;
    white-space: wrap;
    max-width: 400px;
    .tag-list {
        display: inline-block;
        height: 40px;
        .tag-item {
            display: inline-block;
            margin-right: 8px;
            max-width: 148px;
            overflow: hidden;
            white-space: nowrap;
            text-overflow: ellipsis;
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
</style>
