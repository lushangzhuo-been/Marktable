<template>
    <div>
        <el-popover
            ref="GroupBy"
            placement="bottom-start"
            popper-class="group-by-popover"
            width="200"
            trigger="click"
            @show="popoverShow"
            @hide="popoverHide"
        >
            <!-- 遍历item.list  做下拉内容 -->
            <div class="popover-content">
                <div class="group-by-list">
                    <div
                        class="group-by-list-item"
                        :class="{
                            active: colItem.id === currentGroupByInfo.id
                        }"
                        v-for="(colItem, colIndex) in GroupByListShow"
                        :key="colIndex"
                        @click="checkGroupBy(colItem)"
                    >
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
                        {{ colItem.name }}
                    </div>
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
                :class="{ editing: editing, active: currentGroupByInfo.id }"
            >
                <b class="group-by-box"></b>
                分组:{{ currentGroupByInfo.name }}
            </span>
        </el-popover>
    </div>
</template>

<script>
import { FieldType } from "@/assets/tool/const";
import DataHandle from "@/pages/progress/card_board/data_handle.js";
import _ from "lodash";
export default {
    components: {},
    props: {
        currentTab: {
            type: String,
            default: ""
        },
        groupByOption: {
            type: Array,
            default: () => []
        },
        cardFilterDown: {
            // 回显信息
            type: [String, Object],
            default: ""
        }
    },
    data() {
        return {
            currentGroupByInfo: {
                id: 0,
                ws_id: 0,
                tmpl_id: 0,
                field_key: "",
                name: "无分组",
                desc: "",
                mode: "",
                enum_values: "",
                expr: "",
                level: "",
                read_only_rule: "",
                related_tmpl_id: 0
            },
            GroupByListShow: [],
            editing: false
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
        currentTab: {
            handler(id) {
                // 重置
                this.groupByFieldInfo = {};
            }
        },
        groupByOption: {
            handler(arr) {
                let option = _.cloneDeep(arr);
                option.unshift(DataHandle.noGroupInfo);
                this.GroupByListShow = _.cloneDeep(option);
                if (this.cardFilterDown) {
                    this.currentGroupByInfo = _.find(this.GroupByListShow, {
                        field_key: this.cardFilterDown.group_axis
                    });
                    this.$emit(
                        "confim-group-by",
                        this.currentGroupByInfo,
                        "fromViewFilterDown"
                    );
                } else {
                    this.currentGroupByInfo = _.cloneDeep(
                        DataHandle.noGroupInfo
                    );
                }
                // 回显filterDown 内容
            },
            immediate: true,
            deep: true
        },
        curProgress: {
            // 流程切换
            handler() {
                if (this.$refs.GroupBy) {
                    this.$refs.GroupBy.doClose();
                }
            },
            immediate: true
        }
    },
    mounted() {},
    methods: {
        checkGroupBy(info) {
            this.currentGroupByInfo = info;
            this.$emit("confim-group-by", info);
            // this.$refs["GroupBy"].doClose();
        },
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
            let order = "";
            if (this.currentTab === "-1") {
                order = "add";
            } else {
                order = "update";
            }
            this.$emit("group-by-save-view", this.currentGroupByInfo, order);
            this.$refs["GroupBy"].doClose();
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
                this.$refs.GroupBy.doClose();
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
        .group-by-box {
            display: inline-block;
            width: 20px;
            height: 20px;
            background-size: 100% 100%;
            background-image: url(@/assets/image/progress/group-by-active.svg);
            vertical-align: middle;
        }
    }
    &.editing {
        color: #171e17;
        background-color: rgb(202, 228, 253);
        padding: 4px 8px;
        border-radius: 4px;
        .group-by-box {
            display: inline-block;
            width: 20px;
            height: 20px;
            background-size: 100% 100%;
            background-image: url(@/assets/image/progress/group-by-active.svg);
            vertical-align: middle;
        }
    }
    &.active {
        color: #171e17;
        background-color: rgb(202, 228, 253);
        padding: 4px 8px;
        border-radius: 4px;
        .group-by-box {
            display: inline-block;
            width: 20px;
            height: 20px;
            background-size: 100% 100%;
            background-image: url(@/assets/image/progress/group-by-active.svg);
            vertical-align: middle;
        }
    }
    .group-by-box {
        display: inline-block;
        width: 20px;
        height: 20px;
        background-size: 100% 100%;
        background-image: url(@/assets/image/progress/group-by.svg);
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
// @import "@/pages/progress/operation.scss";

// .operation-box {
//     display: inline-block;
//     width: 18px;
//     height: 18px;
//     background-size: 100% 100%;
//     vertical-align: middle;
//     &.save {
//         background-image: url(@/assets/image/progress/save.svg);
//         position: relative;
//         top: -1px;
//     }
// }
.el-popper.group-by-popover {
    box-sizing: border-box;
    margin-top: 0px;
    padding: 0;
    .popper__arrow {
        display: none;
    }
}
.popover-content {
    .group-by-list {
        padding: 4px 16px;
        max-height: 360px;
        overflow-y: auto;
        .group-by-list-item {
            height: 40px;
            line-height: 40px;
            padding: 0 4px;
            border-radius: 4px;
            .type-box {
                display: inline-block;
                width: 20px;
                height: 20px;
                background-size: 100% 100%;
                position: relative;
                top: 4px;
            }
            &:hover {
                background-color: #f1f9ff;
            }
            &.active {
                background-color: #f1f9ff;
            }
        }
    }
}
</style>
