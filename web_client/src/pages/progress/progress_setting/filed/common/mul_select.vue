<template>
    <div
        class="column-block form-width"
        ref="ColumnBlock"
        :class="{
            isEditing: isEditing
        }"
    >
        <!-- 编辑 -->
        <div>
            <el-popover
                ref="DropPopover"
                popper-class="role-multi-select-popover progress-popover"
                placement="bottom-start"
                :width="popoverWidth"
                trigger="click"
                @after-leave="afterLeave"
            >
                <div v-if="optionsList.length" class="options-box">
                    <div class="status-item check-all">
                        <el-checkbox
                            class="role-status"
                            :indeterminate="isIndeterminate"
                            v-model="checkAll"
                            @change="clickAll"
                            >全选</el-checkbox
                        >
                    </div>
                    <!-- 选项列表 -->
                    <div class="check-group-block">
                        <el-checkbox-group
                            v-model="checkedArr"
                            @change="checkboxGroupChange"
                        >
                            <div
                                v-for="(item, index) in optionsList"
                                :key="index"
                                class="status-item"
                            >
                                <el-checkbox
                                    :label="item.name"
                                    class="role-status"
                                    @click="checkboxChange(item)"
                                >
                                    <div
                                        v-overflow
                                        class="status-tag"
                                        :style="{
                                            color: '#171e31',
                                            backgroundColor: rgbToRgba(
                                                `${item.color || '#fff'}`,
                                                0.3
                                            )
                                        }"
                                    >
                                        {{ item.name }}
                                    </div>
                                </el-checkbox>
                            </div>
                        </el-checkbox-group>
                    </div>
                </div>
                <div v-else class="no-data-text">暂无数据</div>
                <!-- 标签化 -->
                <div class="detail" slot="reference">
                    <!-- 名字列表 -->
                    <template>
                        <div
                            ref="listCon"
                            class="tag-list"
                            :class="{ 'show-num': showNum }"
                        >
                            <div
                                v-overflow
                                class="tag-item"
                                v-for="(tagItem, tagIndex) in frontArr"
                                :key="tagIndex"
                                :style="{
                                    color: '#171e31',
                                    backgroundColor: rgbToRgba(
                                        `${tagItem.color || '#fff'}`,
                                        0.3
                                    )
                                }"
                            >
                                {{ tagItem.name }}
                            </div>
                        </div>
                        <!-- 数字气泡 -->
                        <el-tooltip
                            v-show="showNum"
                            class="item"
                            effect="dark"
                            placement="top"
                            popper-class="col-mutil-select-tooltip"
                        >
                            <div slot="content">
                                <div
                                    v-for="(item, index) in behandArr"
                                    :key="index"
                                >
                                    <span
                                        class="tip-item"
                                        :style="{
                                            color: '#171e31',
                                            backgroundColor: rgbToRgba(
                                                `${item.color || '#fff'}`,
                                                0.3
                                            )
                                        }"
                                    >
                                        {{ item.name }}
                                    </span>
                                </div>
                            </div>
                            <b ref="num-box" class="num-box"
                                >+{{ behandArr.length }}</b
                            >
                        </el-tooltip>
                    </template>
                    <span v-if="!behandArr.length"></span>
                </div>
            </el-popover>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import { rgbToRgba } from "@/assets/tool/func";
import { emptySpace } from "@/assets/tool/func";
import { type } from "jquery";
export default {
    props: {
        option: {
            type: Array,
            default: () => []
        },
        value: {
            type: Array,
            default: () => []
        },
        statusList: {
            type: Array,
            default: () => []
        }
    },
    data() {
        return {
            isEditing: false,
            frontArr: [],
            behandArr: [],
            popoverWidth: 220,
            optionsList: [],
            labelIndex: 0,
            showNum: false,
            checkAll: false,
            isIndeterminate: false,
            checkedArr: [],
            selectOption: []
        };
    },
    watch: {
        option: {
            handler(arr) {
                this.optionsList = _.cloneDeep(arr);
            },
            deep: true
        },
        statusList: {
            handler(arr) {
                if (arr && arr.length) {
                    this.checkedArr = arr.map((item) => item.name);
                } else {
                    this.checkedArr;
                }
                this.$nextTick(() => {
                    this.checkboxGroupChange(arr, true);
                });
            },
            deep: true,
            immediate: true
        },
        checkedArr: {
            handler(arr) {
                let ids = [];
                //    提交必填校验
                if (this.option && this.option.length) {
                    ids = this.option
                        .filter((item) => {
                            return arr.includes(item.name);
                        })
                        .map((sub) => sub.id);
                }
                this.$emit("checked-arr", ids);
            },
            deep: true,
            immediate: true
        }
    },
    mounted() {},
    methods: {
        rgbToRgba(color, opacity) {
            return rgbToRgba(color, opacity);
        },
        getTagInit() {
            // 会因为突然显示的数字球  导致换行
            const listCon = this.$refs.listCon;
            if (listCon) {
                const labels = listCon.querySelectorAll(".tag-item");
                let labelIndex = 0; // 渲染到第几个
                const listConBottom = listCon.getBoundingClientRect().bottom; // 容器底部距视口顶部距离
                this.showNum = false;
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
        getShowLabel(labelIndex) {
            this.labelIndex = labelIndex;
            this.frontArr = this.getArrFront(this.selectOption);
            this.behandArr = this.getArrBehand(this.selectOption);
        },
        getAllLabel() {
            this.frontArr = this.selectOption;
            this.behandArr = [];
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
        checkScope() {
            this.isEditing = !this.isEditing;
        },
        afterLeave() {
            this.isEditing = false;
        },
        checkboxChange() {
            let arr = [];
            this.checkAll.forEach((item) => {
                if (_.find(this.optionsList, { name: item })) {
                    arr.push(_.find(this.optionsList, { name: item }));
                }
            });
            this.selectOption = arr;
        },
        clickAll(val) {
            if (val) {
                let arr = [];
                this.optionsList.forEach((item) => {
                    arr.push(item.name);
                });
                this.checkedArr = arr;
                this.selectOption = this.optionsList;
                this.frontArr = this.optionsList;
            } else {
                this.checkedArr = [];
                this.selectOption = [];
                this.frontArr = [];
            }
            this.isIndeterminate = false;
            this.$nextTick(() => {
                this.getTagInit();
            });
        },
        checkboxGroupChange(selectArr, flag) {
            let arr = [];
            selectArr.forEach((item) => {
                if (flag) {
                    if (_.find(this.optionsList, { name: item.name })) {
                        arr.push(_.find(this.optionsList, { name: item.name }));
                    }
                } else {
                    if (_.find(this.optionsList, { name: item })) {
                        arr.push(_.find(this.optionsList, { name: item }));
                    }
                }
            });
            this.selectOption = arr;
            this.frontArr = arr;
            let checkedCount = selectArr.length;
            // 有值且相等
            this.checkAll =
                checkedCount && checkedCount === this.optionsList.length;
            this.isIndeterminate =
                checkedCount > 0 && checkedCount < this.optionsList.length;
            this.$nextTick(() => {
                this.getTagInit();
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.empty-notic {
    font-size: 14px;
    color: #c0c4cc;
    vertical-align: middle;
    display: inline-block;
    line-height: 24px;
    height: 24px;
}
.column-block {
    height: 32px;
    background: #ffffff;
    border-radius: 4px;
    border: 1px solid #e6e9f0;
}
.detail {
    display: flex;
    align-items: center;
    height: 32px;
    padding: 0 10px;
    border-radius: 4px;
    .tag-list {
        display: inline-block;
        height: 32px;
        max-width: calc(100% - 40px);
        &.show-num {
            white-space: nowrap;
        }
        .tag-item {
            box-sizing: border-box;
            display: inline-block;
            height: 24px;
            line-height: 24px;
            text-align: center;
            margin: 4px 8px 4px 0;
            padding: 0 4px;
            border-radius: 4px;
            text-overflow: ellipsis;
            overflow: hidden;
            white-space: nowrap;
            max-width: 100%;
        }
    }
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
    vertical-align: middle;
    font-weight: 400;
    font-size: 12px;
    color: #2f384c;
}
::v-deep.status-item {
    height: 40px;
    line-height: 40px;
    padding: 0 16px;
    .role-status {
        display: flex;
        flex-direction: row-reverse;
        justify-content: space-between;
        height: 100%;
        .status-tag {
            box-sizing: border-box;
            padding: 0 10px;
            line-height: 24px;
            height: 24px;
            text-align: center;
            border-radius: 4px;
            text-overflow: ellipsis;
            overflow: hidden;
            white-space: nowrap;
            cursor: pointer;
            margin: 8px 0;
        }
        .el-checkbox__input {
            line-height: inherit;
        }
        .el-checkbox__label {
            max-width: calc(100% - 36px);
            line-height: 40px;
            padding: 0;
        }
    }
    &:hover {
        background-color: #f5faff;
    }
    &.check-all {
        border-bottom: 1px solid #f0f1f5;
    }
}
</style>
<style lang="scss">
// tooltip数字提示内容
.el-tooltip__popper.col-mutil-select-tooltip {
    .tip-item {
        display: inline-block;
        font-size: 12px;
        padding: 0 4px;
        border-radius: 3px;
        height: 20px;
        box-sizing: border-box;
        line-height: 21px;
        margin: 0px 4px 3px 0px;
    }
}
// popover下拉框
.el-popover.role-multi-select-popover {
    max-height: 360px;
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
            border: 1px solid rgb(205, 28, 230);
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
        // margin: 8px 0 4px 8px;
        // padding-right: 8px;
        // max-height: 210px;
        overflow-y: auto;
        position: relative;
        .check-all {
            position: absolute;
            top: 0;
            width: calc(100% - 32px);
        }
        .check-group-block {
            margin-top: 40px;
            max-height: 320px;
            overflow-y: auto;
        }
    }
    .no-data-text {
        padding: 24px 0;
        text-align: center;
        color: #909399;
    }
}
</style>
