<template>
    <div>
        <div ref="listCon" class="tag-list" :class="{ 'show-num': showNum }">
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
            <b ref="num-box" class="num-box">+{{ behandArr.length }}</b>
        </el-tooltip>
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
            modelArr: [],
            frontArr: [],
            behandArr: [],
            labelIndex: 0,
            showNum: false
        };
    },
    watch: {
        option: {
            handler(arr) {},
            immediate: true,
            deep: true
        },
        selectValue: {
            handler(str) {
                let valueArr = [];
                let modelArr = [];
                valueArr = str.split(",");
                valueArr.forEach((item) => {
                    if (_.find(this.option, { value: item })) {
                        let obj = _.find(this.option, { value: item });
                        modelArr.push(obj);
                    } else {
                        let obj = {
                            color: "rgb(240, 241, 245)",
                            value: item
                        };
                        modelArr.push(obj);
                    }
                });
                this.modelArr = modelArr;
                this.frontArr = this.modelArr;
                this.$nextTick(() => {
                    this.getTagInit();
                });
            },
            immediate: true
        }
    },
    methods: {
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
            this.frontArr = this.getArrFront(this.modelArr);
            this.behandArr = this.getArrBehand(this.modelArr);
        },
        getAllLabel() {
            this.frontArr = this.modelArr;
            this.behandArr = [];
        }
    }
};
</script>

<style lang="scss" scoped>
.tag-list {
    display: inline-block;
    height: 40px;
    white-space: wrap;
    // max-width: calc(100% - 40px);
    margin-left: 8px;
    .tag-item {
        margin: 8px 8px 8px 0px;
        display: inline-block;
        height: 24px;
        line-height: 24px;
        text-align: center;
        padding: 0 4px;
        border-radius: 4px;
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
    // vertical-align: top;
    font-weight: 400;
    font-size: 12px;
    color: #2f384c;
    margin: auto 0;
}
</style>
