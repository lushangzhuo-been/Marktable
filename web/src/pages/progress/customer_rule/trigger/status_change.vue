<template>
    <div
        class="column-block"
        ref="ColumnBlock"
        :class="{ isEditing: isEditing }"
    >
        <!-- 任何人/指定人 -->
        <el-popover
            ref="DropPopover"
            popper-class="progress-popover"
            placement="bottom-start"
            :width="144"
            trigger="click"
            @show="popoverShow"
            @after-leave="afterLeave"
        >
            <div>
                <div
                    @click="checkAny"
                    class="list-item multiple-people-list-item"
                >
                    <b class="value-type anyone"></b>
                    任意值
                </div>
                <el-popover
                    ref="DropPopover"
                    popper-class="col-multi-select-popover progress-popover"
                    placement="right-start"
                    :width="196"
                    trigger="hover"
                >
                    <!-- <div class="popover-content multi-select-col"> -->
                    <el-input
                        class="basic-ui progress-input"
                        ref="SearchInput"
                        v-model="searchInput"
                        placeholder="搜索"
                        prefix-icon="el-icon-search"
                    ></el-input>
                    <div class="selected-popover" v-if="selectArr.length">
                        <span
                            v-for="(item, index) in selectArr"
                            :key="index"
                            class="selected-item"
                        >
                            <span class="value"> {{ item.name }}</span>
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
                            {{ item.name }}
                        </div>
                    </div>
                    <div v-show="!optionsList.length" class="no-data-text">
                        暂无数据
                    </div>
                    <!-- </div> -->
                    <div
                        slot="reference"
                        class="list-item multiple-people-list-item"
                    >
                        <b class="value-type appoint"></b>
                        指定值
                    </div>
                </el-popover>
            </div>
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
                                color: '#fff',
                                backgroundColor: tagItem.color
                            }"
                        >
                            {{ tagItem.name }}
                        </div>
                    </div>
                    <!-- 数字气泡 -->
                    <el-tooltip
                        v-show="behandArr.length"
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
                                    {{ item.name }}
                                </span>
                            </div>
                        </div>
                        <b ref="num-box" class="num-box"
                            >+{{ behandArr.length }}</b
                        >
                    </el-tooltip>
                </template>
                <div v-if="anyCheckbox" style="color: #606266">任意值</div>
                <div
                    v-if="!anyCheckbox && !selectArr.length"
                    style="color: #c0c4cc"
                >
                    请选择
                </div>
                <b class="triangle" @click="checkScope"></b>
            </div>
        </el-popover>
    </div>
</template>

<script>
import _ from "lodash";
export default {
    props: {
        placeholder: {
            type: String,
            default: "请选择"
        },
        option: {
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
            isEditing: false,
            selectArr: [], // 回显用的name 与 颜色
            frontArr: [],
            behandArr: [],
            popoverWidth: 220,
            optionsList: [],
            searchInput: "",
            labelIndex: 0,
            showNum: true,
            anyCheckbox: null
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
        option: {
            handler(arr) {
                this.optionsList = arr;
            },
            immediate: true
        },
        searchInput(str) {
            if (str.trim()) {
                this.optionsList = _.cloneDeep(
                    this.getOptionsList().filter((item) => {
                        return (
                            item.name
                                .toLocaleUpperCase()
                                .indexOf(str.toLocaleUpperCase()) > -1
                        );
                    })
                );
            } else {
                this.optionsList = this.getOptionsList();
            }
        },
        value: {
            handler(str) {
                if (str === "任意值") {
                    this.anyCheckbox = true;
                } else if (str) {
                    // 组装
                    let arr = str.split(",");
                    let showArr = [];
                    arr.forEach((id) => {
                        let numId = parseInt(id);
                        if (_.find(this.optionsList, { id: numId })) {
                            showArr.push(
                                _.find(this.optionsList, { id: numId })
                            );
                        } else {
                            let obj = {
                                color: "rgb(24, 144, 255)",
                                id: numId,
                                name: "已被删除"
                            };
                            showArr.push(obj);
                        }
                    });
                    this.selectArr = showArr;
                    this.frontArr = this.selectArr;
                    this.optionsList = this.getOptionsList();
                    this.$nextTick(() => {
                        this.getTagInit();
                    });
                }
            },
            immediate: true
        }
    },
    mounted() {},
    methods: {
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
            const listCon = this.$refs.listCon;
            if (listCon) {
                const labels = listCon.querySelectorAll(".tag-item");
                let labelIndex = 0; // 渲染到第几个
                const listConBottom = listCon.getBoundingClientRect().bottom; // 容器底部距视口顶部距离
                this.showNum = true;
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
            this.getSubmitValue();
            this.optionsList = this.getOptionsList();
            this.frontArr = this.selectArr;
            this.$nextTick(() => {
                this.getTagInit();
            });
        },
        getOptionsList() {
            let tmpList = [];
            if (this.selectArr.length && this.optionsList) {
                tmpList = this.option.filter((v) =>
                    this.selectArr.every((val) => val.name !== v.name)
                );
            } else {
                tmpList = this.option || [];
            }
            if (this.searchInput.trim()) {
                return _.cloneDeep(
                    tmpList.filter((item) => {
                        return (
                            item.name
                                .toLocaleUpperCase()
                                .indexOf(
                                    this.searchInput.trim().toLocaleUpperCase()
                                ) > -1
                        );
                    })
                );
            } else {
                return tmpList;
            }
        },
        checkScope() {
            this.searchValue = "";
            if (this.formItem.can_modify === "no") {
                return;
            }
            this.isEditing = !this.isEditing;
            if (this.isEditing) {
                this.popoverWidth = this.$refs.ColumnBlock.clientWidth;
            } else {
                this.afterLeave();
            }
        },
        popoverShow() {
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
        // 选择任意值
        checkAny() {
            this.anyCheckbox = true;
            this.selectArr = [];
            this.frontArr = [];
            this.behandArr = [];
            this.$emit("input", "任意值");
            this.$refs.DropPopover.doClose();
        },
        checkboxChange(item) {
            this.anyCheckbox = false;
            this.selectArr.push(item);
            this.getSubmitValue();
            this.optionsList = this.getOptionsList();
            this.frontArr = this.selectArr;
            this.$nextTick(() => {
                this.getTagInit();
            });
        },
        // 即将提交的参数
        getSubmitValue() {
            let submitValue = [];
            this.selectArr.forEach((item) => {
                submitValue.push(item.id);
            });
            let str = submitValue.join(",");
            this.$emit("input", str);
        }
    }
};
</script>

<style lang="scss" scoped>
@import "@/pages/progress/add_progress/style.scss";
.column-block {
    .detail {
        display: flex;
        height: 32px;
        padding: 0 10px;
        white-space: wrap;
        .tag-list {
            display: inline-block;
            height: 32px;
            .tag-item {
                box-sizing: border-box;
                display: inline-block;
                height: 24px;
                line-height: 24px;
                text-align: center;
                margin: 3px 8px 3px 0;
                padding: 0 4px;
                border-radius: 4px;
                text-overflow: ellipsis;
                overflow: hidden;
                white-space: nowrap;
                max-width: 100%;
            }
        }
    }
}
.default-text {
    color: #c0c4cc;
}

.tag-item {
    display: inline-block;
    margin-right: 8px;
    overflow: hidden;
    text-wrap: nowrap;
    text-overflow: ellipsis;
    max-width: 148px;
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
    top: 4px;
    vertical-align: middle;
    font-weight: 400;
    font-size: 12px;
    color: #2f384c;
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
.member-list {
    margin: 6px 0;
    &:first-child {
        margin-top: 0;
    }
    &:last-child {
        margin-bottom: 0;
    }
}
.value-type {
    display: inline-block;
    width: 18px;
    height: 18px;
    background-size: 100% 100%;
    vertical-align: middle;
    &.anyone {
        background-image: url(@/assets/image/progress/add_rule/anyone.svg);
    }
    &.appoint {
        background-image: url(@/assets/image/progress/add_rule/appoint.svg);
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
            cursor: pointer;
            margin-bottom: 4px;
            .sub-item {
                display: inline-block;
                width: calc(100% - 20px);
                text-overflow: ellipsis;
                overflow: hidden;
                white-space: nowrap;
            }
        }
    }
    .no-data-text {
        padding: 24px 0;
        text-align: center;
        color: #909399;
    }
}
</style>
