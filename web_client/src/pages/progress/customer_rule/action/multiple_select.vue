<template>
    <div
        class="column-block"
        ref="ColumnBlock"
        :class="{ isEditing: isEditing }"
    >
        <div>
            <el-popover
                ref="DropPopover"
                popper-class="progress-popover col-multi-select-popover"
                placement="bottom"
                :width="popoverWidth"
                trigger="click"
                @show="popoverShow"
                @after-leave="afterLeave"
            >
                <!-- 下拉内容 -->
                <div class="popover-content multi-select-col">
                    <el-input
                        class="basic-ui progress-input"
                        ref="SearchInput"
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
                            <span class="value"> {{ item.value }}</span>
                            <b
                                class="el-icon-error"
                                @click="removeItem(item, index)"
                            ></b>
                        </span>
                    </div>
                    <div v-show="optionsList.length" class="options-box">
                        <div
                            class="drapdown-item"
                            v-for="(item, index) in optionsList"
                            :key="index"
                            :style="{
                                color: '#fff',
                                backgroundColor: item.color
                            }"
                            @click="checkboxChange(item)"
                        >
                            {{ item.value }}
                        </div>
                    </div>
                    <div v-show="!optionsList.length" class="no-data-text">
                        暂无数据
                    </div>
                </div>
                <!-- 点击此展示下拉内容 -->
                <div slot="reference">
                    <!-- 名字列表 -->
                    <!-- 名字列表 -->
                    <div v-show="frontArr.length" class="detail">
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
                                        color: '#fff',
                                        backgroundColor: tagItem.color
                                    }"
                                >
                                    {{ tagItem.value }}
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
                                                color: '#fff',
                                                backgroundColor: item.color
                                            }"
                                        >
                                            {{ item.value }}
                                        </span>
                                    </div>
                                </div>
                                <b ref="num-box" class="num-box"
                                    >+{{ behandArr.length }}</b
                                >
                            </el-tooltip>
                        </template>
                        <span v-show="!selectArr.length"></span>
                    </div>
                    <div v-show="!frontArr.length" class="detail default-text">
                        请选择
                    </div>
                    <b class="triangle"></b>
                </div>
            </el-popover>
        </div>
    </div>
</template>

<script>
import { type } from "jquery";
import _ from "lodash";
export default {
    props: {
        fieldList: {
            type: Array,
            default: () => []
        },
        actionItem: {
            type: Object,
            default: () => {}
        },
        value: {
            type: String,
            default: ""
        }
    },
    components: {},
    data() {
        return {
            curFormItem: {},
            isEditing: false,
            searchValue: "",
            selectArr: [],
            frontArr: [],
            behandArr: [],
            popoverWidth: 220,
            optionsListOrigin: [],
            optionsList: [],
            labelIndex: 0,
            showNum: false
        };
    },
    watch: {
        actionItem: {
            handler(obj) {
                if (_.find(this.fieldList, { field_key: obj.field_key })) {
                    this.optionsListOrigin = _.find(this.fieldList, {
                        field_key: obj.field_key
                    }).enum_values
                        ? JSON.parse(
                              _.find(this.fieldList, {
                                  field_key: obj.field_key
                              }).enum_values
                          )
                        : [];
                    if (this.value) {
                        this.reshowSelectArr();
                    }
                }
                this.optionsList = this.getOptionsList();
            },
            deep: true,
            immediate: true
        },
        searchValue(str) {
            if (str.trim()) {
                this.optionsList = _.cloneDeep(
                    this.getOptionsList().filter((item) => {
                        return (
                            item.value
                                .toLocaleUpperCase()
                                .indexOf(str.toLocaleUpperCase()) > -1
                        );
                    })
                );
            } else {
                this.optionsList = this.getOptionsList();
            }
        },
        selectArr: {
            handler(arr) {
                let paramArr = [];
                arr.forEach((item) => {
                    paramArr.push(item.value);
                });
                let param = paramArr.join(",");
                this.$emit("input", param);
            },
            deep: true
        }
    },
    mounted() {},
    methods: {
        reshowSelectArr() {
            let arr = this.value.split(",");
            let selectArr = [];
            arr.forEach((value) => {
                if (
                    _.find(this.optionsListOrigin, {
                        value: value
                    })
                ) {
                    selectArr.push(
                        _.find(this.optionsListOrigin, {
                            value: value
                        })
                    );
                } else {
                    let obj = {
                        value: value,
                        color: "rgb(24, 144, 255)"
                    };
                    selectArr.push(obj);
                }
            });
            this.selectArr = selectArr;
            this.frontArr = this.selectArr;
            this.$nextTick(() => {
                this.getTagInit();
            });
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
        getShowLabel(labelIndex) {
            this.labelIndex = labelIndex;
            this.frontArr = this.getArrFront(this.selectArr);
            this.behandArr = this.getArrBehand(this.selectArr);
        },
        getAllLabel() {
            this.frontArr = this.selectArr;
            this.behandArr = [];
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
        removeItem(item, index) {
            this.selectArr.splice(index, 1);
            this.optionsList = this.getOptionsList();
            this.frontArr = this.selectArr;
            this.$nextTick(() => {
                this.getTagInit();
            });
        },
        getOptionsList() {
            let tmpList = [];
            if (this.optionsListOrigin) {
                tmpList = this.optionsListOrigin.filter((v) =>
                    this.selectArr.every((val) => val.value !== v.value)
                );
            } else {
                tmpList = this.optionsListOrigin || [];
            }
            if (this.searchValue.trim()) {
                return _.cloneDeep(
                    tmpList.filter((item) => {
                        return (
                            item.value
                                .toLocaleUpperCase()
                                .indexOf(
                                    this.searchValue.trim().toLocaleUpperCase()
                                ) > -1
                        );
                    })
                );
            } else {
                return tmpList;
            }
        },
        popoverShow() {
            this.searchValue = "";
            this.isEditing = true;
            this.popoverWidth = this.$refs.ColumnBlock.clientWidth;
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
            this.selectArr.push(item);
            this.optionsList = this.getOptionsList();
            this.frontArr = this.selectArr;
            this.$nextTick(() => {
                this.getTagInit();
            });
        }
    }
};
</script>

<style lang="scss" scoped>
@import "@/pages/progress/detail/components/node_operation_form/style.scss";
.default-text {
    color: #c0c4cc;
}
.detail {
    box-sizing: border-box;
    display: flex;
    align-items: center;
    border: 1px solid #dcdfe6;
    padding: 0 10px;
    height: 32px;
    border-radius: 4px;
    &:focus {
        border: 1px solid var(--PRIMARY_COLOR);
    }
    .tag-list {
        display: inline-block;
        height: 32px;
        white-space: wrap;
        max-width: calc(100% - 30px);
        &.show-num {
            white-space: nowrap;
        }
        .tag-item {
            display: inline-block;
            height: 24px;
            line-height: 24px;
            text-align: center;
            margin: 4px 8px 6px 0px;
            padding: 0 4px;
            border-radius: 4px;
        }
    }
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
    vertical-align: middle;
    font-weight: 400;
    font-size: 12px;
    color: #2f384c;
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
.el-popover.col-multi-select-popover {
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
            text-align: center;
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
