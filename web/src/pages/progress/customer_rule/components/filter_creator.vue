<template>
    <div>
        <el-popover
            ref="Sort"
            placement="bottom-start"
            popper-class="operation-popover"
            trigger="click"
            width="300"
            @show="popoverShow"
            @hide="popoverHide"
        >
            <!-- 遍历item.list  做下拉内容 -->
            <div class="popover-content">
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
                <div v-if="personList.length" class="had-select-tag">
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
                        v-for="(listItem, listIndex) in showList"
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
                    <div v-show="!showList.length" class="list-no-data">
                        暂无数据
                    </div>
                </div>
            </div>
            <span
                slot="reference"
                class="btn-icon"
                :class="{ editing: editing, effect: effectiveParamsNum }"
            >
                创建人
                <b class="icon-box"></b>
                <span v-if="effectiveParamsNum">
                    {{ `(${effectiveParamsNum})` || "" }}
                    <i
                        @click="removeAll"
                        class="el-tag__close el-icon-close filter-delete"
                    ></i>
                </span>
            </span>
        </el-popover>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import { imgHost } from "@/assets/tool/const";
import { cloneDeep } from "lodash";
export default {
    components: {},
    props: {},
    data() {
        return {
            editing: false,
            showList: [], //总列表备份
            peopleList: [], //总列表
            personList: [], //已选
            selectArr: [], //已选id
            searchInput: ""
        };
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        },
        effectiveParamsNum() {
            return this.selectArr.length;
        }
    },
    watch: {
        searchInput: _.debounce(function () {
            this.getPeopleList();
        }, 500),
        selectArr: {
            handler(arr) {
                let param = arr.join(",");
                this.$emit("input", param);
            }
        }
    },
    mounted() {},
    methods: {
        removeAll() {
            this.isIndeterminate = false;
            this.selectArr = [];
            this.personList = [];
            this.getPeopleList();
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
        removeTagItem(people) {
            // 移除提交参数
            let arr = cloneDeep(this.selectArr);
            _.remove(arr, function (id) {
                return id === people.id;
            });
            this.selectArr = arr;
            // 移除页面绑定
            _.remove(this.personList, function (persion) {
                return persion.id === people.id;
            });
            this.getPeopleList();
        },
        confirmPeople(people) {
            // 列表选择
            this.selectArr.push(people.id);
            this.removeHadSelect();
            this.getPeopleList();
            if (_.find(this.peopleList, { id: people.id })) {
                this.personList.push(
                    _.find(this.peopleList, { id: people.id })
                );
            }
        },
        // 手动移除所选
        removeHadSelect() {
            this.selectArr.forEach((id) => {
                if (_.find(this.showList, { id: id })) {
                    _.remove(this.showList, { id: id });
                }
            });
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
                    this.showList = _.cloneDeep(this.peopleList);
                    this.removeHadSelect();
                } else {
                    this.peopleList = [];
                    this.showList = [];
                }
            });
        },
        closePopver() {
            if (this.$refs.Sort) {
                // 重置参数
                this.$refs.Sort.doClose();
            }
        },
        popoverShow() {
            this.editing = true;
            this.getPeopleList();
        },
        popoverHide() {
            this.editing = false;
        }
    }
};
</script>

<style lang="scss" scoped>
@import "@/pages/progress/detail/components/node_operation_form/style.scss";
.icon-box {
    display: inline-block;
    width: 18px;
    height: 18px;
    background-image: url(@/assets/image/progress/add_rule/creator.svg);
    background-size: 100% 100%;
    vertical-align: middle;
    position: relative;
    top: -2px;
}
.btn-icon {
    cursor: pointer;
    color: #171e17;
    padding: 4px 8px;
    border-radius: 4px;
    &:hover {
        // color: var(--PRIMARY_COLOR);
        color: #171e17;
        background-color: rgb(202, 228, 253);
        padding: 4px 8px;
        border-radius: 4px;
    }
    &.editing,
    &.effect {
        color: #171e17;
        background-color: rgb(202, 228, 253);
        padding: 4px 8px;
        border-radius: 4px;
    }
}
.type-icon {
    display: inline-block;
    width: 18px;
    height: 18px;
    background-size: 100% 100%;
    vertical-align: text-top;
    &.yes {
        background-image: url(@/assets/image/progress/add_rule/update.svg);
    }
    &.no {
        background-image: url(@/assets/image/progress/add_rule/email.svg);
    }
}
// filter-delete
.el-tag__close.el-icon-close {
    &.tag-delete {
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
        &::before {
            content: "\e6db";
            font-size: 8px;
            vertical-align: middle;
        }
    }
    &.filter-delete {
        width: 12px;
        height: 12px;
        text-align: center;
        line-height: 12px;
        color: #ffffff;
        background-color: #a6abbc;
        border-radius: 50%;
        vertical-align: middle;
        margin-left: 2px;
        &:hover {
            background-color: #82889e;
        }
        &::before {
            content: "\e6db";
            font-size: 8px;
            position: relative;
            top: -2px;
        }
    }
}

.had-select-tag {
    padding: 4px;
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
</style>
<style lang="scss">
.el-popper.operation-popover {
    box-sizing: border-box;
    margin-top: 0px;
    padding: 0;
    .popper__arrow {
        display: none;
    }
}
</style>
