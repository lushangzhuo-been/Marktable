<template>
    <div>
        <template v-if="frontArr.length">
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
                v-show="behandArr.length"
                class="item"
                effect="dark"
                placement="top"
                popper-class="col-mutil-select-tooltip"
            >
                <div slot="content">
                    <div v-for="(item, index) in behandArr" :key="index">
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
                <div class="num-block">
                    <b ref="num-box" class="num-box">+{{ behandArr.length }}</b>
                </div>
            </el-tooltip>
        </template>

        <div v-show="showAny" class="show-any">任意值</div>
    </div>
</template>

<script>
import _ from "lodash";
export default {
    props: {
        option: {
            type: Array,
            default: () => []
        },
        selectValue: {
            type: String,
            default: ""
        }
    },
    data() {
        return {
            showNum: false,
            showAny: false,
            selectOption: [],
            frontArr: [],
            behandArr: []
        };
    },
    watch: {
        option: {
            handler(arr) {
                this.selectOption = arr;
            },
            deep: true,
            immediate: true
        },
        selectValue: {
            handler(str) {
                if (str === "任意值") {
                    this.showAny = true;
                } else if (str) {
                    this.showAny = false;
                    // 按值生成对应的全量信息
                    let valueArr = [];
                    valueArr = str.split(",");
                    let modelValue = [];
                    valueArr.forEach((item) => {
                        if (_.find(this.selectOption, { value: item })) {
                            let obj = _.find(this.selectOption, {
                                value: item
                            });
                            modelValue.push(obj);
                        } else {
                            let obj = {
                                color: "rgb(24, 144, 255)",
                                value: item
                            };
                            modelValue.push(obj);
                        }
                    });
                    this.selectArr = modelValue;
                    this.frontArr = this.selectArr;
                    this.$nextTick(() => {
                        this.getTagInit();
                    });
                } else {
                    return "--";
                }
            },
            deep: true,
            immediate: true
        }
    },
    methods: {
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
        getShowLabel(labelIndex) {
            this.labelIndex = labelIndex;
            this.frontArr = this.getArrFront(this.selectArr);
            this.behandArr = this.getArrBehand(this.selectArr);
        },
        getAllLabel() {
            this.frontArr = this.selectArr;
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
        }
    }
};
</script>

<style lang="scss" scoped>
.tag-list {
    display: inline-block;
    height: 40px;
    .tag-item {
        box-sizing: border-box;
        display: inline-block;
        height: 24px;
        line-height: 24px;
        text-align: center;
        margin: 8px 8px 8px 0;
        padding: 0 4px;
        border-radius: 4px;
        text-overflow: ellipsis;
        overflow: hidden;
        white-space: nowrap;
        max-width: 100%;
    }
}
.num-block {
    display: inline-block;
    vertical-align: top;
    height: 100%;
    padding: 9px 0;
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
        font-weight: 400;
        font-size: 12px;
        color: #2f384c;
    }
}
.show-any {
    height: 40px;
    line-height: 40px;
    color: #171e31;
}
</style>
