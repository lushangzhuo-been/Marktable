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
                >
                    {{ tagItem.title }}
                </div>
            </div>
        </template>

        <div v-show="showAny" class="show-any">任意值</div>
    </div>
</template>

<script>
import _ from "lodash";
export default {
    props: {
        selectValue: {
            type: String,
            default: ""
        },
        relationReshowInfo: {
            type: Array,
            default: () => []
        }
    },
    data() {
        return {
            showNum: false,
            showAny: false,
            frontArr: [],
            behandArr: []
        };
    },
    watch: {
        selectValue: {
            handler(str) {
                if (str === "任意值") {
                    this.showAny = true;
                } else if (str) {
                    this.showAny = false;
                    this.frontArr = this.relationReshowInfo;
                } else {
                    return "--";
                }
            },
            deep: true,
            immediate: true
        }
    },
    methods: {}
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
