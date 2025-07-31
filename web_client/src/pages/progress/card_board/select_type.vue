<template>
    <el-tooltip popper-class="basic-ui" effect="dark" placement="top">
        <div slot="content" class="card-tooltip-content">
            {{ fieldItem.name + "：" }}
            <span v-for="(item, index) in selectArr" :key="index"
                >{{ item.value }}
                <span v-show="index !== selectArr.length - 1">、</span>
            </span>
        </div>
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
                    :style="{
                        color: '#fff',
                        backgroundColor: tagItem.color
                    }"
                >
                    {{ tagItem.value }}
                </div>
            </div>
            <b v-show="showNum" class="num-box" ref="num-box"
                >+{{ behandArr.length }}</b
            >
        </div>
    </el-tooltip>
</template>

<script>
import _ from "lodash";
export default {
    props: {
        fieldItem: {
            type: Object,
            default: () => {}
        },
        detailInfo: {
            type: Array,
            default: () => []
        }
    },
    data() {
        return {
            selectArr: [],
            frontArr: [],
            behandArr: [],
            labelIndex: 1,
            showNum: true
        };
    },
    watch: {
        detailInfo: {
            handler(info) {
                let values = info;
                if (values && values.length) {
                    let selectArr = [];
                    // 双重循环拿到选项的颜色
                    values.forEach((value, index) => {
                        selectArr.push({
                            value: value,
                            color: this.getColor(value) // 在配置中找到对应的颜色
                            // color: "#F0F1F5" // 在配置中找到对应的颜色
                        });
                    });
                    this.selectArr = selectArr;
                    if (this.selectArr && this.selectArr.length > 1) {
                        this.showNum = true;
                    } else {
                        this.showNum = false;
                    }
                    this.$nextTick(() => {
                        this.frontArr = this.getArrFront(this.selectArr);
                        this.behandArr = this.getArrBehand(this.selectArr);
                    });
                } else {
                    this.selectArr = [];
                }
            },
            immediate: true,
            deep: true
        }
    },
    methods: {
        getColor(value) {
            let color;
            if (_.find(this.fieldItem.enum_values, { value: value })) {
                color = _.find(this.fieldItem.enum_values, {
                    value: value
                }).color;
            } else {
                color = "#F0F1F5";
            }
            return color;
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
.detail {
    display: flex;
    align-items: center;
    height: 24px;
    .tag-list {
        display: inline-block;
        height: 24px;
        &.show-num {
            white-space: nowrap;
        }
        .tag-item {
            display: inline-block;
            min-width: 24px;
            max-width: 48px;
            box-sizing: border-box;
            height: 24px;
            line-height: 24px;
            text-align: center;
            margin: 0px 8px 0px 0;
            padding: 0 4px;
            border-radius: 4px;
            text-overflow: ellipsis;
            overflow: hidden;
            white-space: nowrap;
            // max-width: 100%;
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
    vertical-align: middle;
    font-weight: 400;
    font-size: 12px;
    color: #2f384c;
}
</style>
