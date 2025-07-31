<template>
    <div
        class="column-block"
        ref="ColumnBlock"
        :class="{ isEditing: isEditing }"
    >
        <!-- 任何人/指定人 -->
        <el-popover
            ref="DropPopover"
            popper-class="progress-popover"
            placement="bottom-start"
            :width="144"
            trigger="click"
            @show="popoverShow"
            @after-leave="afterLeave"
        >
            <div>
                <div
                    class="list-item multiple-people-list-item"
                    @click="checkAny"
                >
                    <b class="value-type anyone"></b>
                    任意值
                </div>
                <el-popover
                    ref="DropPopover"
                    popper-class="progress-popover"
                    placement="right-start"
                    :width="196"
                    trigger="hover"
                >
                    <div>
                        <!-- 搜索框 -->
                        <div class="search-input">
                            <el-input
                                ref="SearchInput"
                                class="basic-ui progress-search-input"
                                v-model="searchInput"
                                size="small"
                                placeholder="搜索"
                                prefix-icon="el-icon-search"
                            ></el-input>
                        </div>
                        <!-- 所选tag -->
                        <div class="had-select-tag">
                            <!-- 遍历personList -->
                            <div
                                class="popover-tag-item"
                                v-for="(tagItem, tagIndex) in personList"
                                :key="tagIndex"
                            >
                                <el-avatar
                                    class="progress-avatar"
                                    icon="el-icon-user-solid"
                                    size="small"
                                    :src="getAvatar(tagItem.avatar)"
                                    :style="getPopoverTagListAvatar()"
                                ></el-avatar>
                                <div class="name-content">
                                    {{ tagItem.full_name }}
                                </div>
                                <i
                                    @click="removeTagItem(tagItem)"
                                    class="el-tag__close el-icon-close tag-delete"
                                ></i>
                            </div>
                        </div>
                        <!-- 列表 -->
                        <div class="list-content">
                            <div
                                class="list-item multiple-people-list-item"
                                v-for="(listItem, listIndex) in personOption"
                                :key="listIndex"
                                @click="confirmPeople(listItem)"
                            >
                                <el-avatar
                                    class="progress-avatar"
                                    icon="el-icon-user-solid"
                                    size="small"
                                    :src="getAvatar(listItem.avatar)"
                                ></el-avatar>
                                {{ listItem.full_name }}
                            </div>
                            <div
                                v-show="!personOption.length"
                                class="list-no-data"
                            >
                                暂无数据
                            </div>
                        </div>
                    </div>
                    <div
                        slot="reference"
                        class="list-item multiple-people-list-item"
                    >
                        <b class="value-type appoint"></b>
                        指定值
                    </div>
                </el-popover>
            </div>
            <!-- 遍历item.list  做下拉内容 -->
            <!-- 标签化 -->
            <div class="detail" slot="reference">
                <!-- 名字列表 -->
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
                            </div>
                        </div>
                        <b ref="num-box" class="num-box"
                            >+{{ behandArr.length }}</b
                        >
                    </el-tooltip>
                </div>
                <div v-if="anyCheckbox" style="color: #606266">任意值</div>
                <div
                    v-if="!anyCheckbox && !selectArr.length"
                    class="default-text"
                >
                    {{ placeholder }}
                </div>
                <b class="triangle"></b>
            </div>
        </el-popover>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import { imgHost } from "@/assets/tool/const";
import { type } from "jquery";
export default {
    props: {
        placeholder: {
            type: String,
            default: "请选择"
        },
        userInfoList: {
            type: Array,
            default: () => []
        },
        value: {
            type: String,
            default: ""
        }
    },
    data() {
        return {
            isEditing: false,
            selectArr: [], // 双向绑定id数组
            frontArr: [],
            behandArr: [],
            personList: [], // 向外传参，全部个人信息的数组
            searchInput: "",
            peopleList: [],
            personOption: [],
            labelIndex: 0,
            showNum: true,
            submitValue: null,
            anyCheckbox: false
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
    watch: {
        searchInput: _.debounce(function () {
            this.getPeopleList();
        }, 500),
        userInfoList: {
            handler(arr) {
                if (arr && arr.length) {
                    this.personList = arr;
                    this.frontArr = this.personList;
                    this.$nextTick(() => {
                        this.getTagInit();
                    });
                }
            },
            immediate: true,
            deep: true
        },
        value: {
            handler(str) {
                if (str === "任意值") {
                    this.anyCheckbox = true;
                } else if (str) {
                    this.selectArr = str.split(",");
                }
            },
            immediate: true
        }
    },
    mounted() {},
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
        getPeopleList() {
            let params = {
                ws_id: this.curSpace.id,
                ex: this.selectArr.join(","),
                key: this.searchInput,
                page_size: 50,
                page: 1
            };
            api.getPeopleList(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.peopleList = res.data;
                    this.personOption = res.data;
                } else {
                    this.peopleList = [];
                    this.personOption = [];
                }
            });
        },
        removeTagItem(people) {
            // 移除提交参数
            let arr = _.cloneDeep(this.selectArr);
            _.remove(arr, function (id) {
                return id === people.id;
            });
            this.selectArr = arr;
            this.getSubmitValue();
            // 移除页面绑定
            _.remove(this.personList, function (persion) {
                return persion.id === people.id;
            });
            this.frontArr = this.personList;
            this.$nextTick(() => {
                this.getTagInit();
            });
        },
        // 选择任意值
        checkAny() {
            this.anyCheckbox = true;
            this.selectArr = [];
            this.personList = [];
            this.frontArr = [];
            this.behandArr = [];
            // this.submitValue = ["任意值"];
            // this.$emit("user-change", this.submitValue);
            this.$emit("input", "任意值");
            this.$refs.DropPopover.doClose();
        },
        confirmPeople(people) {
            // 列表选择
            this.anyCheckbox = false;
            this.selectArr.push(people.id);
            this.getSubmitValue();
            this.removeHadSelect();
            this.getPeopleList();
            if (_.find(this.peopleList, { id: people.id })) {
                this.personList.push(
                    _.find(this.peopleList, { id: people.id })
                );
            }
            this.frontArr = this.personList;
            this.$nextTick(() => {
                this.getTagInit();
            });
        },
        removeHadSelect() {
            this.selectArr.forEach((id) => {
                if (_.find(this.showList, { id: id })) {
                    _.remove(this.showList, { id: id });
                }
            });
        },
        popoverShow() {
            this.isEditing = true;
            this.$nextTick(() => {
                if (this.$refs.SearchInput) {
                    this.$refs.SearchInput.focus();
                }
            });
            this.getPeopleList();
        },
        afterLeave() {
            this.isEditing = false;
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
        getPopoverTagListAvatar() {
            return {
                position: "relative",
                top: "-6px"
            };
        },
        getAvatarStack() {
            return {
                position: "relative",
                top: "-2px"
            };
        },
        // 即将提交的参数
        getSubmitValue() {
            this.submitValue = this.selectArr;
            // this.$emit("user-change", this.submitValue);
            let str = this.selectArr.join(",");
            this.$emit("input", str);
        }
    }
};
</script>

<style lang="scss" scoped>
// src\pages\progress\add_progress\style.scss
@import "@/pages/progress/add_progress/style.scss";
.column-block {
    .detail {
        display: flex;
        height: 32px;
        padding: 0 10px;
        white-space: wrap;
        .mem-list-content {
            display: flex;
            // white-space: nowrap;
            .tag-list {
                display: inline-block;
                height: 32px;
                &.show-num {
                    // width: calc(100% - 32px);
                    // white-space: nowrap;
                }
            }
        }
    }
}
.default-text {
    color: #c0c4cc;
}

.tag-item {
    display: inline-block;
    margin-right: 8px;
    overflow: hidden;
    text-wrap: nowrap;
    text-overflow: ellipsis;
    max-width: 148px;
}

.avatar-box {
    display: inline-block;
    min-width: 24px;
    height: 24px;
    border-radius: 50%;
    background-size: 100% 100%;
    vertical-align: middle;
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
    position: relative;
    top: 4px;
    vertical-align: middle;
    font-weight: 400;
    font-size: 12px;
    color: #2f384c;
}
.list-item {
    padding: 0 12px;
    height: 40px;
    line-height: 40px;
    &:hover {
        background-color: #f1f9ff;
    }
}
.popover-tag-item {
    display: inline-block;
    padding: 0 6px 0 0;
    height: 20px;
    line-height: 20px;
    font-size: 12px;
    color: #82889e;
    background-color: #e2e7f0;
    margin: 4px;
    border-radius: 10px;
    max-width: calc(100% - 8px);
    .name-content {
        display: inline-block;
        max-width: calc(100% - 40px);
        padding: 0 4px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }
}
.had-select-tag {
    padding: 4px;
}
.el-tag__close.el-icon-close.tag-delete {
    width: 12px;
    height: 12px;
    text-align: center;
    line-height: 12px;
    color: #ffffff;
    background-color: #a6abbc;
    border-radius: 50%;
    vertical-align: super;
    &:hover {
        background-color: #82889e;
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
.value-type {
    display: inline-block;
    width: 18px;
    height: 18px;
    background-size: 100% 100%;
    vertical-align: middle;
    &.anyone {
        background-image: url(@/assets/image/progress/add_rule/anyone.svg);
    }
    &.appoint {
        background-image: url(@/assets/image/progress/add_rule/appoint.svg);
    }
}
</style>
