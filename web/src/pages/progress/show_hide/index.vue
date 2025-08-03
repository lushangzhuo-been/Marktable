<template>
    <div>
        <el-popover
            ref="HideShow"
            placement="bottom-start"
            popper-class="hide-show-popover"
            width="200"
            trigger="click"
            @show="popoverShow"
            @hide="popoverHide"
        >
            <!-- 遍历item.list  做下拉内容 -->
            <div class="popover-content">
                <div class="hide-show-list">
                    <el-checkbox-group
                        v-model="checkArr"
                        class="hide-show-checkbox-group"
                        @change="checkChange"
                    >
                        <draggable
                            v-model="hiddenShowColumnsListShow"
                            handle=".drag"
                            :move="onMove"
                            @change="draggableChange"
                        >
                            <transition-group>
                                <div
                                    class="hide-show-list-item"
                                    v-for="(
                                        colItem, colIndex
                                    ) in hiddenShowColumnsListShow"
                                    :key="colIndex"
                                >
                                    <el-checkbox
                                        :label="colItem.field_key"
                                        :disabled="
                                            colItem.field_key === 'title'
                                        "
                                        class="hide-show-checkbox"
                                    >
                                        <b class="show-hide-drag-icon drag"></b>
                                        <b
                                            class="type-box"
                                            :style="
                                                getType(
                                                    colItem.mode,
                                                    colItem.expr,
                                                    colItem.field_key
                                                )
                                            "
                                        ></b>
                                        <!-- {{ getName(colItem.field_key) }} -->
                                        {{ colItem.name }}
                                    </el-checkbox>
                                </div>
                            </transition-group>
                        </draggable>
                    </el-checkbox-group>
                </div>
                <div class="btn-group">
                    <div class="save-block" @click="confirm">
                        <b class="operation-box save"></b>
                        {{ currentTab === "-1" ? "保存为新视图" : "保存" }}
                    </div>
                </div>
            </div>
            <span
                slot="reference"
                class="btn-icon"
                :class="{ editing: editing, active: hiddenNum }"
            >
                <b class="hide-show-box"></b>
                显示/隐藏列{{ hiddenNum || "" }}
            </span>
        </el-popover>
    </div>
</template>

<script>
import draggable from "vuedraggable";
import { FieldType } from "@/assets/tool/const";
import api from "@/common/api/module/progress";
import _ from "lodash";
export default {
    components: {
        draggable
    },
    props: {
        hiddenShowColumnsList: {
            type: Array,
            default: () => []
        },
        currentTab: {
            type: String,
            default: ""
        }
    },
    data() {
        return {
            hiddenShowColumnsListShow: [],
            checkArr: [], // hiddenShowColumnsListShow中show为yes的
            allColumnConfig: [], // 全量列， 匹配寻找对应类型，name等。
            editing: false,
            hiddenNum: 0
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
        hiddenShowColumnsList: {
            handler(arr) {
                this.checkArr = [];
                this.hiddenShowColumnsListShow = _.cloneDeep(arr);
                let hiddenNum = 0;
                arr.forEach((item) => {
                    if (item.show === "yes") {
                        this.checkArr.push(item.field_key);
                    } else {
                        hiddenNum += 1;
                    }
                });
                this.hiddenNum = hiddenNum;
            },
            immediate: true,
            deep: true
        },
        curProgress: {
            // 流程切换
            handler() {
                if (this.$refs.HideShow) {
                    this.$refs.HideShow.doClose();
                }
            },
            immediate: true
        }
    },
    mounted() {},
    methods: {
        getType(type, expr, field_key) {
            if (type === "person_com") {
                return {
                    "background-image": `url(${FieldType[type + "_" + expr]})`
                };
            } else if (type === "blank_com") {
                if (field_key === "status") {
                    return {
                        "background-image": `url(${FieldType.select_com})`
                    };
                }
                if (field_key === "creator") {
                    return {
                        "background-image": `url(${FieldType.person_com_single})`
                    };
                }
                if (field_key === "created_at" || field_key === "updated_at") {
                    return {
                        "background-image": `url(${FieldType.time_com})`
                    };
                }
            } else if (type) {
                return {
                    "background-image": `url(${FieldType[type]})`
                };
            } else {
                return {
                    "background-image": `url(${FieldType.text_com})`
                };
            }
        },
        confirm() {
            // 将chekc_arr中的列保存为true，向外提交
            let hiddenShowResult = _.cloneDeep(this.hiddenShowColumnsListShow);
            hiddenShowResult.forEach((item) => {
                if (this.checkArr.indexOf(item.field_key) > -1) {
                    item.show = "yes";
                } else {
                    item.show = "no";
                }
            });
            let order = "";
            if (this.currentTab === "-1") {
                order = "add";
            } else {
                order = "update";
            }
            this.$emit("save-hidden-show", hiddenShowResult, order);
            this.$refs["HideShow"].doClose();
        },
        checkChange(arr) {
            // 按照hiddenShowColumnsList的顺序  将arr中存在的show设为yes 否则 no
            let newHiddenShowColumnList = _.cloneDeep(
                this.hiddenShowColumnsList
            );
            let hiddenNum = 0;
            newHiddenShowColumnList.forEach((item) => {
                if (arr.indexOf(item.field_key) > -1) {
                    item.show = "yes";
                } else {
                    item.show = "no";
                    hiddenNum += 1;
                }
            });
            this.hiddenNum = hiddenNum;
            this.$emit("check-change", newHiddenShowColumnList);
        },
        draggableChange() {
            let newHiddenShowColumnList = _.cloneDeep(
                this.hiddenShowColumnsListShow
            );
            let hiddenNum = 0;
            newHiddenShowColumnList.forEach((item) => {
                if (this.checkArr.indexOf(item.field_key) > -1) {
                    item.show = "yes";
                } else {
                    item.show = "no";
                    hiddenNum += 1;
                }
            });
            this.hiddenNum = hiddenNum;
            this.$emit("check-change", newHiddenShowColumnList);
        },
        onMove(e) {
            //不允许停靠
            if (e.relatedContext.element.field_key === "title") {
                return false;
            }
            //不允许拖拽
            if (e.draggedContext.element.field_key === "title") {
                return false;
            }
            return true;
        },
        popoverShow() {
            this.editing = true;
        },
        popoverHide() {
            this.editing = false;
        },
        closePopver() {
            if (this.$refs.Sort) {
                // 重置参数
                this.$refs.HideShow.doClose();
            }
        }
    }
};
</script>

<style lang="scss" scoped>
@import "@/pages/progress/operation.scss";
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
        .hide-show-box {
            display: inline-block;
            width: 20px;
            height: 20px;
            background-size: 100% 100%;
            background-image: url(@/assets/image/progress/hide-show-active.png);
            vertical-align: middle;
        }
    }
    &.editing {
        // color: var(--PRIMARY_COLOR);
        color: #171e17;
        background-color: rgb(202, 228, 253);
        padding: 4px 8px;
        border-radius: 4px;
        .hide-show-box {
            display: inline-block;
            width: 20px;
            height: 20px;
            background-size: 100% 100%;
            background-image: url(@/assets/image/progress/hide-show-active.png);
            vertical-align: middle;
        }
    }
    &.active {
        // color: var(--PRIMARY_COLOR);
        color: #171e17;
        background-color: rgb(202, 228, 253);
        padding: 4px 8px;
        border-radius: 4px;
        .hide-show-box {
            display: inline-block;
            width: 20px;
            height: 20px;
            background-size: 100% 100%;
            background-image: url(@/assets/image/progress/hide-show-active.png);
            vertical-align: middle;
        }
    }
    .hide-show-box {
        display: inline-block;
        width: 20px;
        height: 20px;
        background-size: 100% 100%;
        background-image: url(@/assets/image/progress/hide-show.png);
        vertical-align: middle;
    }
}
.btn-group {
    padding: 8px;
    display: flex;
    justify-content: right;
    box-shadow: 0px -3px 8px 1px rgba(0, 0, 0, 0.1);
}
.save-block {
    cursor: pointer;
    &:hover {
        color: var(--PRIMARY_COLOR);
        .save {
            background-image: url(@/assets/image/progress/save-active.svg);
        }
    }
    &.disabled {
        color: #a6abbc;
        cursor: not-allowed;
        .save {
            background-image: url(@/assets/image/progress/save-disabled.svg);
        }
    }
}
</style>
<style lang="scss">
.el-popper.hide-show-popover {
    box-sizing: border-box;
    margin-top: 0px;
    padding: 0;
    .popper__arrow {
        display: none;
    }
}
.popover-content {
    .hide-show-list {
        padding: 4px 0;
        max-height: 360px;
        overflow-y: auto;
        .hide-show-list-item {
            height: 40px;
            line-height: 40px;
            padding-right: 4px;
            padding: 0 12px 0 0;
            .show-hide-drag-icon {
                display: inline-block;
                width: 20px;
                height: 20px;
                background-size: 100% 100%;
                position: relative;
                top: 4px;
            }
            .type-box {
                display: inline-block;
                width: 20px;
                height: 20px;
                background-size: 100% 100%;
                position: relative;
                top: 4px;
            }
            &:hover {
                .show-hide-drag-icon {
                    display: inline-block;
                    width: 20px;
                    height: 20px;
                    background-image: url(@/assets/image/common/icon_dragtable_move.png);
                    background-size: 100% 100%;
                    position: relative;
                    top: 4px;
                }
            }
        }
    }
}

// 多选
.el-checkbox.hide-show-checkbox {
    width: 100%;
    display: flex;
    flex-direction: row-reverse;
    justify-content: space-between;
    height: 40px;
    line-height: 40px;
    &.is-checked {
        .el-checkbox__label {
            color: #606266;
        }
    }
    .el-checkbox__label {
        padding-left: 0px;
        line-height: 40px;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
    }
    .el-checkbox__input {
        line-height: 44px;
    }
}
</style>
