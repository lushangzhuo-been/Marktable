<template>
    <div>
        <el-popover
            ref="Sort"
            placement="bottom-start"
            popper-class="sort-popove"
            trigger="click"
            @show="popoverShow"
            @hide="popoverHide"
        >
            <!-- 遍历item.list  做下拉内容 -->
            <div class="popover-content">
                <div class="sort-list">
                    <draggable
                        v-model="sortGroup"
                        handle=".drag"
                        :move="onMove"
                    >
                        <transition-group>
                            <div
                                class="sort-list-item"
                                v-for="(sortItem, sortIndex) in sortGroup"
                                :key="sortIndex"
                            >
                                <!-- 顺序按钮 -->
                                <b class="show-hide-drag-icon drag"></b>
                                <!-- 字段 -->
                                <el-select
                                    class="basic-ui field"
                                    size="small"
                                    v-model="sortItem.field_key"
                                    placeholder="请选择字段"
                                    @change="
                                        (val) => sortFieldChange(val, sortIndex)
                                    "
                                >
                                    <el-option
                                        v-for="item in fileConfig"
                                        :key="item.field_key"
                                        :label="item.name"
                                        :value="item.field_key"
                                    >
                                    </el-option>
                                </el-select>
                                <!-- 正倒序 -->
                                <el-select
                                    class="basic-ui relationship"
                                    size="small"
                                    v-model="sortItem.sort_order"
                                    placeholder="请选择"
                                    @change="sortOrderChange"
                                >
                                    <el-option
                                        v-for="item in sortRelation"
                                        :key="item.value"
                                        :label="item.label"
                                        :value="item.value"
                                    >
                                    </el-option>
                                </el-select>
                                <b
                                    class="delete-box"
                                    @click="deleteItem(sortIndex)"
                                ></b>
                            </div>
                        </transition-group>
                    </draggable>
                </div>
                <div class="operation-block">
                    <div class="add-sort" @click="addSort">
                        <b class="operation-box add"></b>
                        添加条件
                    </div>
                    <!-- v-show="sortGroup.length"  -->
                    <div class="others-operation">
                        <div
                            v-show="sortGroup.length"
                            class="clear-sort"
                            @click="clearSort"
                        >
                            <b class="operation-box clear"></b>
                            清空
                        </div>
                        <div
                            class="save-sort"
                            :class="{ disabled: sortGroupCheck }"
                            @click="saveSort"
                        >
                            <b class="operation-box save"></b>
                            {{ currentTab === "-1" ? "保存为新视图" : "保存" }}
                        </div>
                    </div>
                </div>
            </div>
            <span
                slot="reference"
                class="btn-icon"
                :class="{ editing: editing }"
            >
                <b class="sort-box"></b>
                排序{{ effectiveParamsNum || "" }}
            </span>
        </el-popover>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import draggable from "vuedraggable";
export default {
    components: {
        draggable
    },
    props: {
        currentTab: {
            type: String,
            default: ""
        },
        sortOrder: {
            type: Array,
            default: () => []
        }
    },
    data() {
        return {
            sortGroup: [], // 列表
            fileConfig: [], // 字段选项配置
            editing: false,
            sortGroupCheck: false, // true为禁用 ， false为可用
            sortRelation: [
                {
                    label: "升序",
                    value: 1
                },
                {
                    label: "降序",
                    value: -1
                }
            ],
            effectiveParamsNum: 0
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
        sortOrder: {
            handler(arr) {
                this.sortGroup = arr;
            },
            deep: true
        },
        sortGroup: {
            handler() {
                this.checkSort();
            },
            deep: true
        }
    },
    mounted() {},
    methods: {
        closePopver() {
            if (this.$refs.Sort) {
                // 重置参数
                this.$refs.Sort.doClose();
            }
        },
        sortFieldChange() {
            // 重置对应的关系，值
            this.$emit("sort-search", this.sortGroup);
        },
        sortOrderChange() {
            this.$emit("sort-search", this.sortGroup);
        },
        checkSort() {
            // 检查是否每条filter都有值
            let res = false;
            let effectiveParamsNum = 0;
            this.sortGroup.forEach((filterItem) => {
                // if (!filterItem.field_key || !filterItem.sort_order) {
                //     res = true;
                //     return res;
                // }
                if (filterItem.field_key && filterItem.sort_order) {
                    effectiveParamsNum += 1;
                } else {
                    res = true;
                    // return res;
                }
            });
            this.effectiveParamsNum = effectiveParamsNum;
            this.sortGroupCheck = res;
        },
        addSort() {
            let obj = {
                field_key: "",
                sort_order: 1
            };
            this.sortGroup.push(obj);
        },
        saveSort() {
            if (this.sortGroupCheck) {
                return;
            }
            let order = "";
            if (this.currentTab === "-1") {
                order = "add";
            } else {
                order = "update";
            }
            this.$emit("save-sort", this.sortGroup, order);
        },
        clearSort() {
            this.sortGroup = [];
            this.$emit("sort-search", this.sortGroup);
        },
        mountedSort() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress
            };
            api.getFilterConfig(params).then((res) => {
                this.fileConfig = res.data;
            });
        },
        onMove() {
            return true;
        },
        deleteItem(index) {
            this.sortGroup.splice(index, 1);
        },
        popoverShow() {
            this.editing = true;
            this.mountedSort();
        },
        popoverHide() {
            this.editing = false;
        }
    }
};
</script>

<style lang="scss" scoped>
.sort-list-item {
    .show-hide-drag-icon {
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
        .sort-box {
            display: inline-block;
            width: 20px;
            height: 20px;
            background-size: 100% 100%;
            background-image: url(@/assets/image/progress/sort-active.svg);
            vertical-align: middle;
        }
    }
    &.editing {
        // color: var(--PRIMARY_COLOR);
        color: #171e17;
        background-color: rgb(202, 228, 253);
        padding: 4px 8px;
        border-radius: 4px;
        .sort-box {
            display: inline-block;
            width: 20px;
            height: 20px;
            background-size: 100% 100%;
            background-image: url(@/assets/image/progress/sort-active.svg);
            vertical-align: middle;
        }
    }
    .sort-box {
        display: inline-block;
        width: 20px;
        height: 20px;
        background-size: 100% 100%;
        background-image: url(@/assets/image/progress/sort.svg);
        vertical-align: middle;
    }
}
.view-had-sort {
    .btn-icon {
        cursor: pointer;
        // color: var(--PRIMARY_COLOR);
        color: #171e17;
        background-color: rgb(202, 228, 253);
        padding: 4px 8px;
        border-radius: 4px;
        .sort-box {
            display: inline-block;
            width: 20px;
            height: 20px;
            background-size: 100% 100%;
            background-image: url(@/assets/image/progress/sort-active.svg);
            vertical-align: middle;
        }
    }
}
</style>
<style lang="scss">
@import "@/pages/progress/operation.scss";
// .operation-box {
//     display: inline-block;
//     width: 18px;
//     height: 18px;
//     background-size: 100% 100%;
//     vertical-align: middle;
//     &.add {
//         background-image: url(@/assets/image/progress/add.png);
//     }
//     &.clear {
//         background-image: url(@/assets/image/common/delete.svg);
//     }
//     &.save {
//         background-image: url(@/assets/image/progress/save.svg);
//     }
// }
.operation-block {
    display: flex;
    justify-content: space-between;
    .add-sort {
        color: var(--PRIMARY_COLOR);
        cursor: pointer;
    }
    .others-operation {
        display: flex;
        justify-content: space-between;
        margin-left: 24px;
        .clear-sort {
            margin-right: 21px;
            cursor: pointer;
            &:hover {
                color: var(--PRIMARY_COLOR);
                .clear {
                    background-image: url(@/assets/image/common/delete-active.svg);
                }
            }
        }
        .save-sort {
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
    }
}

.popover-content {
    .sort-list {
        max-height: 540px;
        overflow-y: auto;
        .sort-list-item {
            height: 32px;
            line-height: 32px;
            display: flex;
            margin-bottom: 12px;
            .delete-box {
                display: inline-block;
                width: 18px;
                height: 18px;
                background-size: 100% 100%;
                background-image: url(@/assets/image/progress/delete.png);
                vertical-align: middle;
                cursor: pointer;
                margin: 7px;
                &:hover {
                    background-image: url(@/assets/image/progress/delete-active.png);
                }
            }
        }
    }
}

.el-popper.sort-popove {
    box-sizing: border-box;
    margin-top: 0px;
    padding: 12px 16px;
    .popper__arrow {
        display: none;
    }
}
.and-or-block {
    width: 58px;
    margin-right: 8px;
    .and-or-content {
        padding-left: 15px;
    }
}
.el-select.and-or {
    width: 58px;
}
.el-select.field {
    width: 140px;
    margin-right: 24px;
}
.el-select.relationship {
    width: 120px;
    margin-right: 8px;
}
.el-input.value {
    width: 160px;
}
.el-date-editor.el-input.progress-date-picker.time {
    width: 200px;
}
.el-date-editor.el-input.progress-date-picker.date {
    width: 160px;
}
::v-deep .progress-input.el-input {
    .el-input__inner {
        padding: 0;
        border: none;
        // height: 38px;
        // line-height: 38px;
        height: 30px;
        line-height: 30px;
        background-color: rgba(0, 0, 0, 0);
        font-weight: 400;
        font-size: 14px;
        color: #171e31;
        font-family: MiSans, MiSans;
    }
    &.number.hidden-arrow {
        input::-webkit-outer-spin-button,
        input::-webkit-inner-spin-button {
            -webkit-appearance: none;
        }
        input[type="number"] {
            -moz-appearance: textfield;
        }
    }
}
</style>
