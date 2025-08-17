<template>
    <div
        class="column-block"
        ref="ColumnBlock"
        :class="{
            isEditing: isEditing,
            'editing-effect': item.can_modify === 'yes',
            'col-required': item.required === 'yes'
        }"
    >
        <!-- 非编辑状态 -->
        <div
            class="detail"
            v-overflow
            v-if="!isEditing"
            @click="checkScope(scope)"
        >
            <span
                class="not-edit-progress-name"
                v-if="selectArr && selectArr.length"
                @click.stop="openDetail()"
                >{{ selectArr[0].title }}</span
            >

            <span v-else>--</span>
        </div>
        <!-- 编辑状态 -->
        <div v-if="isEditing">
            <el-popover
                ref="DropPopover"
                popper-class="progress-popover"
                placement="bottom"
                :width="popoverWidth"
                trigger="click"
                @after-leave="afterLeave"
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
                    <!-- 已选中流程 -->
                    <div class="had-select-tag" v-show="selectedArrEdit.length">
                        <div
                            class="popover-tag-item"
                            v-for="(tagItem, tagIndex) in selectedArrEdit"
                            :key="tagIndex"
                        >
                            <div class="name-content">
                                {{ tagItem.title }}
                            </div>
                            <!-- 移除icon -->
                            <i
                                @click="removeTagItem(tagItem)"
                                class="el-tag__close el-icon-close tag-delete"
                            ></i>
                        </div>
                    </div>
                    <!-- 下拉列表 -->
                    <div class="list-content">
                        <div
                            v-overflow
                            class="list-item multiple-people-list-item"
                            v-for="(listItem, listIndex) in progressDataList"
                            :key="listIndex"
                            @click="confirmPeople(listItem)"
                        >
                            {{ listItem.title }}
                        </div>
                        <div
                            v-show="!progressDataList.length"
                            class="list-no-data"
                        >
                            暂无数据
                        </div>
                    </div>
                </div>
                <!-- 标签化 -->
                <div class="detail" v-overflow slot="reference">
                    <span v-if="selectedArrEdit && selectedArrEdit.length">{{
                        selectedArrEdit[0].title
                    }}</span>

                    <span v-else></span>
                </div>
            </el-popover>
        </div>
        <!-- 三角动画 -->
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import api_common from "@/common/api";
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
            selectedArrEdit: [],
            behandArr: [],
            popoverWidth: 220,
            progressDataList: [],
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
        // 监听行数据
        scope: {
            handler(scope) {
                let curValue = scope.row[this.item.field_key];
                this.selectArr = curValue && curValue.length ? curValue : [];
                this.selectedArrEdit = _.cloneDeep(this.selectArr);
            },
            immediate: true,
            deep: true
        },
        searchInput: _.debounce(function () {
            this.getProgressDataList();
        }, 500)
    },
    methods: {
        getProgressDataList() {
            let params = {
                ws_id: this.curSpace.id,
                ex: this.selectedArrEdit.length
                    ? this.selectedArrEdit.map((item) => item._id).join(",")
                    : "",
                key: this.searchInput,
                page_size: 20,
                page: 1,
                tmpl_id: this.item.related_tmpl_id // 关联流程列的流程
            };
            api_common.getOtherTmplDataList(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.progressDataList = res.data.list || [];
                } else {
                    this.progressDataList = [];
                }
            });
        },
        // 流程数据跳转
        openDetail() {
            let curItem = this.selectArr[0];
            if (curItem.is_member === "yes") {
                if (curItem.issue_url) {
                    window.open(curItem.issue_url, "_blank");
                } else {
                    this.$message({
                        showClose: true,
                        message: "当前issue_url不存在",
                        type: "warning"
                    });
                }
            } else {
                this.$message({
                    showClose: true,
                    message:
                        "您暂无访问权限,请联系该私有流程超级管理员添加权限",
                    type: "warning"
                });
            }
        },
        // 点击时
        checkScope() {
            this.fetAuthEdit();
        },
        fetAuthEdit() {
            // 获取进展权限
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: this.scope.row._id,
                auth_mode: "edit",
                field_key: this.item.field_key
            };
            api.getUserAuth(params).then((res) => {
                if (res && res.resultCode === 200) {
                    if (res.data) {
                        // 可编辑时执行以下
                        this.isEditing = !this.isEditing;
                        if (this.isEditing) {
                            this.popoverWidth =
                                this.$refs.ColumnBlock.clientWidth;
                            setTimeout(() => {
                                this.$refs.DropPopover.doShow();
                                this.$nextTick(() => {
                                    if (this.$refs.SearchInput) {
                                        this.$refs.SearchInput.focus();
                                    }
                                });
                            }, 20);
                            this.getProgressDataList();
                        } else {
                            //
                            this.afterLeave();
                        }
                    } else {
                        this.isEditing = false;
                    }
                } else {
                    this.isEditing = false;
                }
            });
        },
        // 关闭popover
        afterLeave() {
            let curValue = this.scope.row[this.item.field_key];
            // 校验
            if (
                this.item.required === "yes" &&
                curValue &&
                curValue.length &&
                !this.selectedArrEdit.length
            ) {
                // 如果是必填&原来有数据&此时把数据移除了就提示
                this.$message({
                    showClose: true,
                    message: "此为必填项",
                    type: "warning"
                });
                this.selectedArrEdit = _.cloneDeep(this.selectArr);
                this.isEditing = false;
            } else {
                if (
                    (!this.selectArr.length && !this.selectedArrEdit.length) ||
                    (this.selectArr.length &&
                        this.selectedArrEdit.length &&
                        this.selectArr[0]._id === this.selectedArrEdit[0]._id)
                ) {
                    this.isEditing = false;
                    return;
                }
                this.selectArr = _.cloneDeep(this.selectedArrEdit);
                this.isEditing = false;
                this.$emit(
                    "edit-form-item",
                    JSON.stringify(
                        this.selectedArrEdit.map((item) => item._id)
                    ),
                    this.item.field_key,
                    this.scope.row._id,
                    this.item.mode
                );
            }
        },
        // 选择下拉选项时
        confirmPeople(people) {
            this.selectedArrEdit = [people];
            this.getProgressDataList();
        },
        // 移除选中的
        removeTagItem() {
            this.selectedArrEdit = [];
            this.getProgressDataList();
        }
    }
};
</script>

<style lang="scss" scoped>
.detail {
    height: 40px;
    padding: 0 10px;
    overflow: hidden;
    text-wrap: nowrap;
    text-overflow: ellipsis;
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
.not-edit-progress-name {
    color: var(--PRIMARY_COLOR);
    cursor: pointer;
    &:hover {
        text-decoration: underline;
    }
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
.had-select-tag {
    padding: 4px;
    margin-top: 4px;
}
.popover-tag-item {
    // box-sizing: border-box;
    display: inline-block;
    padding: 0 8px;
    height: 20px;
    line-height: 19px;
    font-size: 12px;
    color: #82889e;
    background-color: #e2e7f0;
    /* margin: 4px; */
    border-radius: 10px;
    max-width: calc(100% - 16px);
    .name-content {
        display: inline-block;
        max-width: calc(100% - 16px);
        /* padding: 0 4px; */
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }
}

.el-tag__close.el-icon-close.tag-delete {
    width: 12px;
    height: 12px;
    margin-left: 4px;
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
    max-height: 260px;
    overflow-y: auto;
    .list-no-data {
        height: 40px;
        line-height: 40px;
        text-align: center;
        font-weight: 400;
        font-size: 14px;
        color: #a6abbc;
        margin-bottom: 8px;
    }
}
</style>
