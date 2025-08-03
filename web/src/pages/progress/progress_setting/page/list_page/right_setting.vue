<template>
    <div class="create-setting-right">
        <div class="title">
            {{ `${tab === "card" ? "卡片" : "列表"}字段设置` }}
        </div>
        <!-- <el-tabs class="basic-ui font" size="small" v-model="activeName">
            <el-tab-pane
                v-for="(item, index) in tabsList"
                :label="item.label"
                :name="item.name"
                :key="index"
            ></el-tab-pane>
        </el-tabs> -->
        <div class="content">
            <div class="left-all-filed box">
                <div class="sub-title">全部字段</div>
                <div class="box-content left-border">
                    <template v-if="leftData && leftData.length">
                        <div
                            class="data-item"
                            v-for="(item, index) in leftData"
                            :key="index"
                        >
                            <span class="name">{{ item.label }}</span>
                            <i class="el-icon-plus" @click="addFiled(item)"></i>
                        </div>
                    </template>
                    <span v-else class="empty-data">暂无内容</span>
                </div>
            </div>
            <div class="right-all-filed box">
                <div class="sub-title">列表字段</div>
                <div class="box-content">
                    <div class="data-item label">
                        <span class="right-name">字段名称</span>
                        <span style="display: inline-block; width: 60px"
                            >是否必填</span
                        >
                        <span style="width: 40px; text-align: right">移除</span>
                    </div>
                    <div id="createFiledData">
                        <template v-if="rightData && rightData.length">
                            <div
                                class="data-item"
                                :class="itemClass(item)"
                                v-for="(item, index) in rightData"
                                :key="index"
                            >
                                <span class="right-name">
                                    <img
                                        style="position: relative; top: 5px"
                                        :src="
                                            require('@/assets/image/common/icon_dragtable_move.png')
                                        "
                                        alt=""
                                        width="20px"
                                        height="20px"
                                    />
                                    <span class="right-name">{{
                                        item.field_name
                                    }}</span>
                                </span>
                                <!-- <el-switch v-model="item.required"></el-switch> -->
                                <span
                                    style="display: inline-block; width: 60px"
                                >
                                    <el-switch
                                        v-model="item.required_flag"
                                        :active-color="PRIMARY_COLOR"
                                        inactive-color="#E6E9F0"
                                        :disabled="fieldDisabled(item)"
                                        @change="
                                            (val) => switchChange(val, item)
                                        "
                                    >
                                    </el-switch>
                                </span>
                                <!-- <i
                                class="el-icon-close"
                                @click="removeFiled(item)"
                            ></i> -->
                                <el-button
                                    style="width: 40px; text-align: right"
                                    @click="removeFiled(item)"
                                    type="text"
                                    size="small"
                                    :disabled="fieldDisabled(item)"
                                >
                                    <i
                                        v-if="!fieldDisabled(item)"
                                        class="el-icon-close"
                                    ></i>
                                </el-button>
                            </div>
                        </template>
                        <span v-else class="empty-data">暂无内容</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import Sortable from "sortablejs";
import api from "@/common/api/module/progress_setting";
import { PRIMARY_COLOR_LIST } from "@/assets/tool/color_list";
import _ from "lodash";
import { param } from "jquery";
import { active } from "sortablejs";
export default {
    props: {
        modules: {
            type: String,
            default: ""
        },
        tab: {
            type: String,
            default: "table"
        }
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        // remindMessage() {
        //     const message = TEXT.PROGRESS_SETTING_SCREEN;
        //     return message.CREATED_REMIND || "";
        // },
        fieldDisabled() {
            return (row) => {
                // 标题不可删除，不可编辑开关
                if (row.field_key === "title" || row.field_key === "handler") {
                    return true;
                } else {
                    return false;
                }
            };
        }
    },
    data() {
        return {
            // activeName: "table",
            // tabsList: [
            //     {
            //         label: "表格",
            //         name: "table"
            //     },
            //     {
            //         label: "卡片",
            //         name: "card"
            //     }
            // ],
            leftData: [],
            rightData: [],
            PRIMARY_COLOR: PRIMARY_COLOR_LIST.PRIMARY_COLOR
        };
    },
    watch: {
        tab: {
            handler(val) {
                if (val) {
                    this.getLeftData();
                    this.getRightData();
                }
            },
            immediate: true
        },
        modules: {
            handler(val) {
                if (val) {
                    this.getLeftData();
                    this.getRightData();
                }
            },
            immediate: true
        },
        $route() {
            if (this.modules) {
                this.getLeftData();
                this.getRightData();
            }
        }
    },
    mounted() {
        this.$nextTick(() => {
            this.dropFiled();
        });
    },
    methods: {
        switchChange(val, item) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                id: item.id,
                required: val ? "yes" : "no",
                can_modify: "yes",
                module: this.modules
            };
            api.setIsRequired(params).then((res) => {
                if (res && res.resultCode !== 200) {
                    item.required_flag = !item.required_flag;
                }
            });
            // 调接口发送更改后的数据给后端并获取新数据
        },
        itemClass(item) {
            if (item.field_key !== "title" && item.field_key !== "handler") {
                return "is-sort";
            } else {
                return "not-sort";
            }
        },
        dropFiled() {
            let tbody = document.querySelector("#createFiledData");
            let _this = this;
            Sortable.create(tbody, {
                // or { name: "...", pull: [true, false, 'clone', array], put: [true, false, array] }
                group: {
                    name: "words",
                    pull: true,
                    put: true
                },
                draggable: ".is-sort",
                animation: 150, // ms, number 单位：ms，定义排序动画的时间
                onAdd: function (evt) {
                    // 拖拽时候添加有新的节点的时候发生该事件
                },
                onUpdate: function (evt) {
                    // 拖拽更新节点位置发生该事件
                },
                onRemove: function (evt) {
                    // 删除拖拽节点的时候促发该事件
                },
                onStart: function (evt) {
                    // 开始拖拽出发该函数
                },
                onSort: function (evt) {
                    // 发生排序发生该事件
                },
                onEnd({ newIndex, oldIndex }) {
                    // 结束拖拽
                    let currRow = _this.rightData.splice(oldIndex, 1)[0];
                    _this.rightData.splice(newIndex, 0, currRow);
                    // 调接口发送更改后的数据给后端并获取新数据
                    if (_this.rightData.length < 2) return;
                    let tmpKeyList = _this.rightData.map(
                        (item) => item.field_key
                    );
                    let params = {
                        ws_id: _this.curSpace.id,
                        tmpl_id: _this.$route.params.id,
                        module: _this.modules
                    };
                    if (tmpKeyList.length) {
                        params.coordinate = tmpKeyList.join(",");
                    }
                    api.moveScreenField(params).then((res) => {
                        if (res && res.resultCode === 200) {
                            //
                        }
                    });
                }
            });
        },
        getLeftData() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                module: this.modules
            };
            api.getRestScreenList(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    res.data.length
                ) {
                    this.leftData = this.formatFiled(res.data);
                } else {
                    this.leftData = [];
                }
            });
        },
        formatFiled(arr) {
            return arr.map((item) => {
                return {
                    label: item.name,
                    value: item.field_key,
                    key: item.id
                };
            });
        },
        getRightData() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                module: this.modules
            };
            api.getScreenList(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    res.data.length
                ) {
                    res.data.forEach((item) => {
                        this.$set(
                            item,
                            "required_flag",
                            item.required === "no" ? false : true
                        );
                    });
                    this.rightData = res.data;
                } else {
                    this.rightData = [];
                }
            });
        },
        addFiled(item) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                module: this.modules,
                field_key: item.value
            };
            api.addScreen(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.getLeftData();
                    this.getRightData();
                    this.$emit("filed-change");
                }
            });
        },
        removeFiled(item) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                module: this.modules,
                id: item.id
            };
            api.delScreen(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.getLeftData();
                    this.getRightData();
                    this.$emit("filed-change");
                }
            });
        }
    }
};
</script>
<style lang="scss" scoped>
.create-setting-right {
    height: calc(100% - 40px);
    .title {
        font-weight: 700;
        font-size: 16px;
        margin-bottom: 20px;
    }
    .content {
        display: flex;
        /* align-items: center; */
        height: calc(100% - 42px);
        .box {
            position: relative;
            box-sizing: border-box;
            margin-right: 24px;
            padding: 12px;
            background-color: #fafafa;
            border-radius: 0px 0px 0px 0px;
            // border: 1px solid #a6abbc;
            border: 1px dashed #a6abbc;
            &:last-child {
                margin-right: 0;
                &::before {
                    content: url("~@/assets/image/progress_setting/page/arrow.svg");
                    position: absolute;
                    top: 50%;
                    left: -20px;
                }
            }
            &.left-all-filed {
                width: calc(40% - 24px);
            }
            &.right-all-filed {
                width: calc(60%);
            }
        }
        .sub-title {
            margin-bottom: 12px;
            font-size: 14px;
            font-weight: 700;
        }
        .box-content {
            height: calc(100% - 33px);
            overflow-y: auto;
            padding-right: 16px;
            margin-right: -16px;
            &.left-border {
            }
        }
        .data-item {
            height: 32px;
            display: flex;
            align-items: center;
            justify-content: space-between;
            .name {
                width: calc(100% - 30px);
                white-space: nowrap;
                overflow: hidden;
                text-overflow: ellipsis;
            }
            .right-name {
                width: calc((100% - 90px));
                white-space: nowrap;
                overflow: hidden;
                text-overflow: ellipsis;
            }
            &.label {
                color: #82889e;
            }
            &.not-sort {
                cursor: not-allowed;
                opacity: 0.5;
            }
        }
    }
    .el-icon-close:before {
        cursor: pointer;
        content: url("~@/assets/image/progress_setting/page/remove.svg");
    }
    .el-icon-plus:before {
        cursor: pointer;
        content: url("~@/assets/image/progress_setting/page/add.svg");
    }
    .el-icon-plus:before {
        content: url("~@/assets/image/progress_setting/page/add.svg");
    }
    .empty-data {
        position: relative;
        display: inline-block;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        color: #a6abbc;
    }
}
</style>
