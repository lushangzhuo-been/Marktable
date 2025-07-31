<template>
    <div class="column-block sub-task" ref="ColumnBlock">
        <div class="detail">
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
                        <user-message :userMessage="tagItem" position="left">
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
                <b class="num-box" ref="num-box">+{{ behandArr.length }}</b>
            </el-tooltip>
            <div v-if="!personList.length">
                {{ emptySpace() }}
            </div>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import { imgHost } from "@/assets/tool/const";
import { emptySpace } from "@/assets/tool/func";
import UserMessage from "@/components/user_message_tip";
export default {
    components: {
        UserMessage
    },
    props: {
        item: {
            type: Object,
            default: () => {}
        },
        scope: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            personList: [],
            selectArr: [],
            frontArr: [],
            behandArr: [],
            popoverWidth: 220,
            labelIndex: 0,
            showNum: true
        };
    },
    computed: {},
    watch: {
        scope: {
            handler(scope) {
                if (scope.row[this.item.field_key]) {
                    this.personList = _.cloneDeep(
                        scope.row[this.item.field_key]
                    );
                } else {
                    this.personList = [];
                }
                this.frontArr = this.personList;
                this.getTagInit();
            },
            immediate: true,
            deep: true
        }
    },
    mounted() {},
    updated() {},
    methods: {
        getShowLabel(labelIndex) {
            this.$nextTick(() => {
                this.labelIndex = labelIndex;
                this.frontArr = this.getArrFront(this.personList);
                this.behandArr = this.getArrBehand(this.personList);
            });
        },
        getAllLabel() {
            this.frontArr = this.personList;
            this.behandArr = [];
        },
        getTagInit() {
            this.showNum = false;
            this.$nextTick(() => {
                const listCon = this.$refs.listCon;
                if (listCon) {
                    const labels = listCon.querySelectorAll(".tag-item");
                    let labelIndex = 0; // 渲染到第几个
                    const listConBottom =
                        listCon.getBoundingClientRect().bottom; // 容器底部距视口顶部距离
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
                }
            });
        },
        emptySpace(param) {
            return emptySpace(param);
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
                top: "-2px"
            };
        }
    }
};
</script>

<style lang="scss" scoped>
.detail {
    display: flex;
    height: 40px;
    padding: 0 10px;
    .tag-list {
        display: inline-block;
        height: 40px;
        max-width: calc(100% - 30px);
        white-space: normal;
        &.show-num {
            white-space: nowrap;
        }
        .tag-item {
            display: inline-block;
            padding-right: 8px;
            overflow: hidden;
            text-wrap: nowrap;
            text-overflow: ellipsis;
            max-width: 148px;
        }
    }
}

.column-block {
    box-sizing: border-box;
    height: 40px;
    line-height: 40px;
    border: 1px solid rgba(0, 0, 0, 0);
    position: relative;
}
.list-content {
    max-height: 200px;
    overflow-y: auto;
}
.num-box {
    display: inline-block;
    min-width: 22px;
    height: 22px;
    line-height: 22px;
    border-radius: 11px;
    padding: 0 2px;
    background-color: #f8f8f8;
    text-align: center;
    position: relative;
    top: 8px;
    vertical-align: middle;
    font-weight: 400;
    font-size: 12px;
    color: #2f384c;
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
</style>
