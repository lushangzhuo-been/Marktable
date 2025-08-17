<template>
    <div class="column-block" ref="ColumnBlock">
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
                    <span>
                        <el-avatar
                            class="progress-avatar"
                            icon="el-icon-user-solid"
                            size="small"
                            :src="getAvatar(tagItem.avatar)"
                            :style="getAvatarStack(tagItem.avatar, tagIndex)"
                        ></el-avatar>
                        {{ tagItem.full_name }}
                    </span>
                </div>
            </div>
        </div>

        <!-- 三角动画 -->
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
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
            isEditing: false,
            personList: [],
            selectArr: [],
            frontArr: [],
            behandArr: [],
            popoverWidth: 220,
            peopleList: [],
            showList: [],
            // 过滤关键词
            searchInput: "",
            labelIndex: 0,
            showNum: true
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
        scope: {
            handler(scope) {
                let selectArr = []; // 双向绑定id数组
                if (scope.row[this.item.field_key]) {
                    this.personList = _.cloneDeep(
                        scope.row[this.item.field_key]
                    );
                    this.personList.forEach((objItem) => {
                        selectArr.push(objItem.id);
                    });
                } else {
                    this.personList = [];
                }
                this.frontArr = this.personList;
                this.selectArr = selectArr;
                this.getTagInit();
            },
            immediate: true,
            deep: true
        },
        searchInput: _.debounce(function () {
            this.getPeopleList();
        }, 500)
    },
    mounted() {},
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
        removeHadSelect() {
            this.selectArr.forEach((id) => {
                if (_.find(this.showList, { id: id })) {
                    _.remove(this.showList, { id: id });
                }
            });
        },
        checkScope() {
            if (this.item.can_modify === "no") {
                return;
            }
            this.isEditing = !this.isEditing;
            if (this.isEditing) {
                this.popoverWidth = this.$refs.ColumnBlock.clientWidth;
                setTimeout(() => {
                    this.$refs.DropPopover.doShow();
                    this.$nextTick(() => {
                        if (this.$refs.SearchInput) {
                            this.$refs.SearchInput.focus();
                        }
                    });
                }, 20);
                this.getPeopleList();
            } else {
                this.afterLeave();
            }
        },
        afterLeave() {
            this.isEditing = false;
            // 校验
            if (
                this.item.required === "yes" &&
                this.scope.row[this.item.field_key] &&
                this.scope.row[this.item.field_key].length &&
                !this.selectArr.length
            ) {
                this.$message({
                    showClose: true,
                    message: "此为必填项",
                    type: "warning"
                });
                let selectArr = []; // 双向绑定id数组
                if (this.scope.row[this.item.field_key]) {
                    this.personList = _.cloneDeep(
                        this.scope.row[this.item.field_key]
                    );
                }
                if (this.personList.length) {
                    this.personList.forEach((objItem) => {
                        selectArr.push(objItem.id);
                    });
                }
                this.frontArr = this.personList;
                this.selectArr = selectArr;
                this.getTagInit();
            } else if (
                (!this.scope.row[this.item.field_key] &&
                    !this.personList.length) ||
                JSON.stringify(this.scope.row[this.item.field_key]) ===
                    JSON.stringify(this.personList)
            ) {
                // 前后相同
            } else {
                this.$emit(
                    "edit-form-item",
                    JSON.stringify(this.selectArr),
                    this.item.field_key,
                    this.scope.row._id,
                    this.item.mode
                );
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
        confirmPeople(people) {
            // 列表选择
            this.personList = [people];
            this.selectArr = [people.id];
            this.removeHadSelect();
            this.getPeopleList();
            this.frontArr = this.personList;
            this.getTagInit();
        },
        removeTagItem(people) {
            if (this.item.required === "yes" && this.selectArr.length === 1) {
                this.$message({
                    showClose: true,
                    message: "此为必填项",
                    type: "warning"
                });
                return;
            }
            // 标签移除
            // 移除提交参数
            _.remove(this.selectArr, function (id) {
                return id === people.id;
            });
            // 移除页面绑定
            _.remove(this.personList, function (persion) {
                return persion.id === people.id;
            });
            this.frontArr = this.personList;
            this.getTagInit();
            this.getPeopleList();
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
        },
        getPopoverTagListAvatar() {
            return {
                position: "relative",
                top: "-6px"
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
        &.show-num {
            white-space: nowrap;
            // overflow: hidden;
        }
    }
}
.tag-item {
    display: inline-block;
    padding-right: 8px;
    overflow: hidden;
    text-wrap: nowrap;
    text-overflow: ellipsis;
    max-width: 148px;
}
.column-block {
    box-sizing: border-box;
    height: 40px;
    line-height: 40px;
    border: 1px solid rgba(0, 0, 0, 0);
    position: relative;
    &:hover {
        .triangle {
            transform: rotate(0deg);
            transition-duration: 0.3s;
            position: absolute;
            right: 5px;
            top: 13px;
            display: inline-block;
            width: 16px;
            height: 16px;
            background-size: 100% 100%;
            background-image: url("~@/assets/image/common/triangle-down.png");
        }
    }
    &.isEditing {
        border: 1px solid var(--PRIMARY_COLOR);
        .triangle {
            transform: rotate(180deg);
            transition-duration: 0.3s;
            position: absolute;
            right: 5px;
            top: 13px;
            display: inline-block;
            width: 16px;
            height: 16px;
            background-size: 100% 100%;
            background-image: url("~@/assets/image/common/triangle-down-active.png");
        }
    }
}

::v-deep .progress-input.el-input {
    .el-input__inner {
        padding: 0;
        border: none;
        height: 39px;
        line-height: 39px;
        background-color: rgba(0, 0, 0, 0);
    }
}

.tag-item {
    display: inline-block;
    margin-right: 8px;
    overflow: hidden;
    text-wrap: nowrap;
    text-overflow: ellipsis;
    max-width: 100%;
    &.had-num {
        max-width: calc(100% - 40px);
    }
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
.list-item {
    padding: 0 12px;
    height: 40px;
    line-height: 40px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
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
.list-content {
    max-height: 200px;
    overflow-y: auto;
}
</style>
