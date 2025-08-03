<template>
    <div>
        <div
            v-for="(item, index) in modelArr"
            :key="index"
            class="first-tag"
            :style="{
                color: '#fff',
                backgroundColor: item.color
            }"
        >
            <span class="value">{{ item.value }}</span>
        </div>
    </div>
</template>

<script>
import { color } from "echarts";
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
            modelArr: []
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
            },
            immediate: true
        }
    }
};
</script>

<style lang="scss" scoped>
.first-tag {
    display: inline-block;
    max-width: 148px;
    height: 24px;
    line-height: 24px;
    text-align: center;
    // margin-right: 8px;
    padding: 0 4px;
    border-radius: 4px;
    text-overflow: ellipsis;
    overflow: hidden;
    white-space: nowrap;
    margin: 8px;
}
</style>
