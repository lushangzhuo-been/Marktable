<template>
    <div
        class="column-block"
        ref="ColumnBlock"
        :class="{ isEditing: isEditing }"
    >
        <div>
            <!-- v-removeAriaHidden -->
            <el-popover
                ref="DropPopover"
                popper-class="progress-popover related-action-popover"
                placement="bottom"
                :width="popoverWidth"
                trigger="click"
                @show="popoverShow"
                @after-leave="afterLeave"
            >
                <!-- 遍历item.list  做下拉内容 -->
                <div class="popover-content multi-select-col">
                    <el-input
                        ref="SearchInput"
                        class="basic-ui progress-input"
                        v-model="searchValue"
                        placeholder="搜索"
                        prefix-icon="el-icon-search"
                    ></el-input>
                    <div class="selected-popover" v-if="selectArr.length">
                        <span
                            v-for="(item, index) in selectArr"
                            :key="index"
                            class="selected-item"
                        >
                            <span class="value"> {{ item.title }}</span>
                            <b class="el-icon-error" @click="removeItem()"></b>
                        </span>
                    </div>
                    <div v-if="progressDataList.length" class="options-box">
                        <div
                            class="drapdown-item"
                            v-for="(item, index) in progressDataList"
                            :key="index"
                            @click="checkboxChange(item)"
                        >
                            {{ item.title }}
                        </div>
                    </div>
                    <div v-else class="no-data-text">暂无数据</div>
                </div>

                <!-- 标签化 -->
                <div slot="reference">
                    <!-- 名字列表 -->
                    <div v-if="selectArr.length" class="detail">
                        <span class="first-tag" v-overflow>
                            {{ selectArr[0].title }}
                        </span>
                    </div>
                    <div v-else class="detail default-text">请选择</div>
                    <b class="triangle"></b>
                </div>
            </el-popover>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import api_common from "@/common/api";
import { baseMixin } from "@/mixin.js";
export default {
    mixins: [baseMixin],
    props: {
        relationInfo: {
            type: Object,
            default: () => {}
        },
        relationReshowInfo: {
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
            validateFailed: false,
            isEditing: false,
            searchValue: "",
            selectArr: [],
            popoverWidth: 220,
            progressDataList: [],
            relatedTmplId: ""
        };
    },
    watch: {
        relationInfo: {
            handler(obj) {
                this.relatedTmplId = obj.related_tmpl_id;
                this.$nextTick(() => {
                    this.$refs.DropPopover.doClose();
                });
                this.searchValue = "";
                this.selectArr = [];
                this.progressDataList = [];
            },
            immediate: true,
            deep: true
        },
        relationReshowInfo: {
            handler(arr) {
                if (arr && arr.length) {
                    this.selectArr = arr;
                }
            },
            immediate: true
        },
        searchValue: _.debounce(function () {
            this.getRelationTemList();
        }, 500)
    },
    mounted() {},
    methods: {
        getRelationTemList() {
            let params = {
                ws_id: this.curSpace.id,
                ex: this.selectArr.length
                    ? this.selectArr.map((item) => item._id).join(",")
                    : "",
                key: this.searchValue,
                page_size: 20,
                page: 1,
                tmpl_id: this.relatedTmplId // 关联流程列的流程
            };
            api_common.getOtherTmplDataList(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.progressDataList = res.data.list || [];
                } else {
                    this.progressDataList = [];
                }
            });
        },
        removeItem() {
            this.selectArr = [];
            let str = "";
            this.$emit("input", str);
        },
        popoverShow() {
            this.searchValue = "";
            this.isEditing = true;
            this.popoverWidth = this.$refs.ColumnBlock.clientWidth;
            this.getRelationTemList();
            this.$nextTick(() => {
                if (this.$refs.SearchInput) {
                    this.$refs.SearchInput.focus();
                }
            });
        },
        afterLeave() {
            this.isEditing = false;
        },
        checkboxChange(item) {
            this.selectArr = [item];
            this.$emit("this.selectArr", this.selectArr);
            let str = this.selectArr[0]._id;
            this.$emit("input", str);
            this.$nextTick(() => {
                this.$refs.DropPopover.doClose();
            });
        }
    }
};
</script>

<style lang="scss" scoped>
@import "@/pages/progress/detail/components/node_operation_form/style.scss";
.detail {
    display: flex;
    align-items: center;
    border: 1px solid #dcdfe6;
    padding: 0 10px;
    white-space: nowrap;
    height: 32px;
    border-radius: 4px;
    &:focus {
        border: 1px solid var(--PRIMARY_COLOR);
    }
    .first-tag {
        display: inline-block;
        max-width: calc(100% - 24px);
        height: 24px;
        line-height: 24px;
        text-align: center;
        margin-right: 8px;
        padding: 0 4px;
        border-radius: 4px;
        text-overflow: ellipsis;
        overflow: hidden;
        white-space: nowrap;
    }
}
</style>
<style lang="scss">
div[aria-hidden="true"] {
    display: none !important;
}
input[aria-hidden="true"] {
    display: none !important;
}
// popover下拉框
.el-popover.related-action-popover {
    padding: 4px 0;
    .progress-input.el-input {
        .el-input__inner {
            padding: 0 16px 0 32px;
            border-radius: 0;
            border: none;
            border-bottom: 1px solid #f0f1f5;
            height: 38px;
            line-height: 38px;
            background-color: rgba(0, 0, 0, 0);
        }
    }
    .selected-popover {
        padding: 8px 8px 2px;
        .selected-item {
            font-size: 12px;
            display: inline-block;
            padding: 0 4px;
            border-radius: 3px;
            box-sizing: border-box;
            max-width: 100%;
            background-color: rgb(245, 246, 248);
            border: 1px solid rgb(205, 213, 230);
            line-height: 1;
            margin: 0px 4px 3px 0px;
        }
        .value {
            display: inline-block;
            text-overflow: ellipsis;
            white-space: nowrap;
            overflow: hidden;
            width: calc(100% - 18px);
            height: 18px;
            line-height: 21px;
            position: relative;
            top: -1px;
        }
        .el-icon-error {
            position: relative;
            top: -4px;
            margin-left: 6px;
            cursor: pointer;
            color: #cdd5e6;
        }
    }

    .options-box {
        margin: 8px 0 4px 8px;
        padding-right: 8px;
        max-height: 210px;
        overflow-y: auto;
        .drapdown-item {
            width: 100%;
            box-sizing: border-box;
            padding: 0 10px;
            line-height: 32px;
            height: 32px;
            text-align: left;
            border-radius: 4px;
            text-overflow: ellipsis;
            overflow: hidden;
            white-space: nowrap;
            cursor: pointer;
            margin-bottom: 4px;
        }
    }
    .no-data-text {
        padding: 24px 0;
        text-align: center;
        color: #909399;
    }
}
</style>
