<template>
    <div class="group-by-list" :class="{ expand: expand, close: !expand }">
        <div class="title-block" :class="{ expand: expand, close: !expand }">
            <b
                v-if="!expand"
                class="expand-btn"
                :class="{ expand: expand, close: !expand }"
                @click="changeExpand"
            ></b>
            <div class="title" :class="{ expand: expand, close: !expand }">
                <b
                    class="type-box"
                    :class="{ expand: expand, close: !expand }"
                    :style="
                        getType(groupByFieldInfo.mode, groupByFieldInfo.exper)
                    "
                ></b>
                <div
                    class="title-name"
                    :class="{ expand: expand, close: !expand }"
                >
                    {{ groupByFieldInfo.name }}
                </div>
            </div>
            <!-- 收起按钮 -->
            <b
                v-if="expand"
                class="expand-btn"
                :class="{ expand: expand }"
                @click="changeExpand"
            ></b>
        </div>
        <!-- 搜索框 -->
        <div v-if="expand" class="search-block">
            <el-input
                class="basic-ui"
                size="small"
                placeholder="搜索"
                v-model="keyWords"
            ></el-input>
        </div>
        <div
            class="group-by-list-content"
            :class="{ expand: expand, close: !expand }"
        >
            <div
                class="group-by-item"
                :class="{
                    'cur-effect':
                        (curEffectEnum?.user_id &&
                            curEffectEnum?.user_id === item.user_id) ||
                        (curEffectEnum?.name &&
                            curEffectEnum?.name === item.name)
                }"
                v-for="(item, index) in groupByInfo"
                :key="index"
                @click="checkField(item)"
            >
                <div class="label" :class="{ expand: expand, close: !expand }">
                    {{ item.name }}
                </div>
                <div class="num" v-if="expand && item.filed_mode !== 'all'">
                    {{ "（" + item.value + "）" }}
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import { FieldType } from "@/assets/tool/const";
export default {
    components: {},
    props: {
        value: {
            type: String,
            default: ""
        },
        currentTab: {
            type: String,
            default: ""
        },
        groupByFieldInfo: {
            type: Object,
            default: () => {}
        },
        groupByList: {
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
            expand: true,
            keyWords: "",
            groupByInfo: [],
            curEffectEnum: {}
        };
    },
    computed: {},
    watch: {
        currentTab: {
            handler() {
                this.keyWords = "";
            }
        },
        keyWords: _.debounce(function () {
            this.filterKeyWords();
        }, 500),
        // 维度字段切换。枚举值列表  默认提交提第一个出去
        groupByList: {
            handler(arr) {
                this.groupByInfo = _.cloneDeep(arr);
                if (arr && arr.length) {
                    // 优先处理this.curEffectEnum，再this.cardFilterDown
                    if (Object.keys(this.curEffectEnum).length) {
                        // 当前选中的枚举值  需要确认列表中枚举值是否仍然存在
                        if (
                            this.groupByFieldInfo.mode === "person_com" &&
                            _.find(arr, {
                                user_id: this.curEffectEnum.user_id
                            })
                        ) {
                            // 当前人类型值仍有效
                        } else if (
                            this.groupByFieldInfo.mode !== "person_com" &&
                            _.find(arr, {
                                name: this.curEffectEnum.name
                            })
                        ) {
                            // 当前非人类型值仍有效
                        } else {
                            // 默认第一个
                            this.curEffectEnum = arr[0];
                        }
                        this.$emit("check-field-search", this.curEffectEnum);
                    } else if (this.cardFilterDown) {
                        // 视图存储的值
                        if (
                            this.groupByFieldInfo.mode === "person_com" &&
                            _.find(arr, {
                                user_id: this.cardFilterDown.group_value
                            })
                        ) {
                            // 优先取视图保存的group_value信息
                            this.curEffectEnum = _.find(arr, {
                                user_id: this.cardFilterDown.group_value
                            });
                        } else if (
                            this.groupByFieldInfo.mode !== "person_com" &&
                            _.find(arr, {
                                user_id: this.cardFilterDown.group_value
                            })
                        ) {
                            this.curEffectEnum = _.find(arr, {
                                name: this.cardFilterDown.group_value
                            });
                        } else {
                            // 默认第一个
                            this.curEffectEnum = arr[0];
                        }
                        this.$emit("check-field-search", this.curEffectEnum);
                    } else {
                        // 空值
                        this.curEffectEnum = arr[0];
                        this.$emit("check-field-search", this.curEffectEnum);
                    }
                }
            },
            immediate: true,
            deep: true
        }
    },
    mounted() {},
    methods: {
        filterKeyWords() {
            let filterArr = this.groupByList.filter((item) => {
                return (
                    item.name
                        .toLocaleUpperCase()
                        .indexOf(this.keyWords.trim().toLocaleUpperCase()) > -1
                );
            });
            this.groupByInfo = filterArr;
        },
        // 只刷新左侧枚举信息
        justRefreshEnumInfoNum(arr) {
            this.groupByInfo = _.cloneDeep(arr);
        },
        checkField(info) {
            this.curEffectEnum = info;
            this.$emit("check-field-search", info);
        },
        changeExpand() {
            this.expand = !this.expand;
        },
        getType(type, expr) {
            if (type && expr) {
                return {
                    "background-image": `url(${FieldType[type + "_" + expr]})`
                };
            } else {
                return {
                    "background-image": `url(${FieldType[type]})`
                };
            }
        }
    }
};
</script>

<style lang="scss" scoped>
.group-by-list {
    box-sizing: border-box;
    width: 40px;
    background: #fafafb;
    border-radius: 4px 0px 0px 4px;
    border: 1px solid #e6e9f0;
    &.expand {
        min-width: 200px;
    }
    .title-block {
        height: 40px;
        line-height: 40px;
        padding: 0 8px;
        display: flex;
        justify-content: space-between;
        &.close {
            height: fit-content;
            text-align: center;
            flex-direction: column;
            line-height: 24px;
        }
        .title {
            &.expand {
                display: flex;
            }
            &.close {
                margin-bottom: 8px;
            }
            .type-box {
                display: inline-block;
                width: 18px;
                height: 18px;
                background-size: 100% 100%;
                &.expand {
                    margin: 11px 0;
                }
                &.close {
                    margin: 0 2px;
                }
            }
            .title-name {
                &.expand {
                    margin-left: 4px;
                    display: inline-block;
                    overflow: hidden;
                    white-space: nowrap;
                    text-overflow: ellipsis;
                }
            }
        }
        .expand-btn {
            display: inline-block;
            width: 18px;
            height: 18px;
            background-size: 100% 100%;
            margin: 10px 2px;
            cursor: pointer;
            background-image: url("~@/assets/image/common/to-expand.svg");
            &.expand {
                background-image: url("~@/assets/image/common/to-close.svg");
            }
            &.close {
                margin-bottom: 4px;
            }
        }
    }
    .search-block {
        padding: 6px 8px;
    }
    .group-by-list-content {
        // max-height: ;
        .group-by-item {
            height: 40px;
            line-height: 40px;
            padding: 0 12px;
            display: flex;
            cursor: pointer;
            &:hover {
                background-color: #ecf4fb;
            }
            &.cur-effect {
                background-color: rgba($color: #1890ff, $alpha: 0.06);
            }
            .label {
                display: inline-block;
                white-space: nowrap;
                overflow: hidden;
                text-overflow: ellipsis;
                &.expand {
                    max-width: calc(100% - 32px);
                }
            }
            .num {
                max-width: 32px;
                display: inline-block;
                color: #82889e;
            }
        }
        &.close {
            .group-by-item {
                overflow: hidden;
            }
        }
    }
}
</style>
