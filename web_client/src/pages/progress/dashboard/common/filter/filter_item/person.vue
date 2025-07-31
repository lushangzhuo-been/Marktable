<template>
    <div
        class="column-block"
        ref="column-block"
        :class="{ isEditing: isEditing }"
    >
        <el-popover
            ref="status-popover"
            :width="popoverWidth"
            popper-class="progress-popover"
            placement="bottom"
            trigger="click"
            @show="popoverShow"
            @after-leave="afterLeave"
        >
            <!-- 遍历item.list  做下拉内容 -->
            <div>
                <!-- 搜索框 -->
                <div class="search-input">
                    <el-input
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
                        v-for="(tagItem, tagIndex) in frontArr"
                        :key="tagIndex"
                        @click="removeTagItem(tagItem)"
                    >
                        <el-avatar
                            class="progress-avatar"
                            icon="el-icon-user-solid"
                            size="small"
                            :src="getAvatar(tagItem.avatar)"
                            :style="getPopoverTagListAvatar(tagItem.avatar)"
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
                    <div v-show="!personOption.length" class="list-no-data">
                        暂无数据
                    </div>
                </div>
            </div>
            <div slot="reference" class="detail">
                <div v-if="frontArr.length">
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
                            :style="getAvatarStack(tagItem.avatar, tagIndex)"
                        ></el-avatar>
                        {{ tagItem.full_name }}
                    </div>
                </div>
                <div class="default-text" v-else>请选择(单人)</div>
                <b class="triangle"></b>
            </div>
        </el-popover>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import { imgHost } from "@/assets/tool/const";
export default {
    props: {
        param: {
            type: [Array, String],
            default: () => []
        }
    },
    data() {
        return {
            isEditing: false,
            frontArr: [],
            selectText: "",
            submitArr: [],
            popoverWidth: 220,
            personOption: [],
            searchInput: "",
            selectArr: []
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
    searchInput: _.debounce(function () {
        this.getPeopleList();
    }, 500),
    watch: {
        param: {
            handler(arr) {
                let selectArr = [];
                if (arr) {
                    arr.forEach((objItem) => {
                        selectArr.push(objItem.id);
                    });
                    this.submitArr = arr;
                }

                this.selectArr = selectArr;
                this.frontArr = this.getArrFront(this.submitArr);
            },
            immediate: true
        }
    },
    mounted() {},
    methods: {
        getArrFront(arr) {
            let deepClone = _.cloneDeep(arr);
            let front = deepClone.splice(0, 1);
            return front;
        },
        popoverShow() {
            this.isEditing = true;
            this.popoverWidth = this.$refs["column-block"].clientWidth;
            // 调查询人的接口
            this.getPeopleList();
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
                    this.personOption = res.data;
                }
            });
        },
        afterLeave() {
            this.isEditing = false;
        },
        getAvatarStack(src, index) {
            return {
                "background-image": `url(${src})`,
                position: "relative",
                top: "-2px"
            };
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
                top: "-7px"
            };
        },
        confirmPeople(people) {
            // 列表选择
            let singleObj;
            if (_.find(this.personOption, { id: people.id })) {
                singleObj = _.find(this.personOption, {
                    id: people.id
                });
            }
            this.submitArr = [singleObj];
            this.frontArr = this.getArrFront(this.submitArr);
            this.$emit("input", this.submitArr);
            this.$refs["status-popover"].doClose();
        },
        removeTagItem(people) {
            // 标签移除
            // 移除提交参数
            _.remove(this.selectArr, function (id) {
                return id === people.id;
            });
            // 移除页面绑定
            _.remove(this.submitArr, function (persion) {
                return persion.id === people.id;
            });
            this.frontArr = this.getArrFront(this.submitArr);
            this.getPeopleList();
            this.$emit("input", this.submitArr);
        }
    }
};
</script>

<style lang="scss" scoped>
@import "./style.scss";
.default-text {
    color: #c0c4cc;
}

.avatar-box {
    display: inline-block;
    min-width: 24px;
    height: 24px;
    border-radius: 50%;
    background-size: 100% 100%;
    vertical-align: middle;
}
.list-item {
    padding: 0 12px;
    height: 40px;
    line-height: 40px;
    &:hover {
        background-color: #f1f9ff;
    }
}
.tag-item {
    display: inline-block;
    overflow: hidden;
    text-wrap: nowrap;
    text-overflow: ellipsis;
    max-width: 100%;
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
</style>
