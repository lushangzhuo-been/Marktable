<template>
    <div class="create-setting-right">
        <div class="title">{{ getTitle }}</div>
        <div class="content">
            <div
                class="left-all-filed box"
                :style="{ width: selected !== 'tmpl_detail' ? '50%' : '40%' }"
            >
                <div class="sub-title">{{ getLeftTitle }}</div>
                <div class="box-content left-border">
                    <el-input
                        v-if="this.selected === 'sub_tmpl'"
                        v-model="search"
                        class="basic-ui"
                        placeholder="搜索"
                        prefix-icon="el-icon-search"
                        size="small"
                    ></el-input>
                    <div
                        :style="{
                            height:
                                selected === 'sub_tmpl'
                                    ? 'calc(100% - 40px)'
                                    : 'calc(100% - 25px)'
                        }"
                        class="overflow-y"
                    >
                        <template v-if="isNotEmpty">
                            <div
                                class="data-item"
                                v-for="item in leftData"
                                :key="item.value"
                                v-show="item.checked"
                            >
                                <span class="left-name" v-overflow
                                    >{{ item.label }}
                                </span>
                                <i
                                    class="el-icon-plus"
                                    @click="addFiled(item)"
                                ></i>
                            </div>
                        </template>
                        <span v-else class="empty-data">暂无内容</span>
                    </div>
                </div>
            </div>
            <div
                class="right-all-filed box"
                :style="{ width: selected !== 'tmpl_detail' ? '50%' : '60%' }"
            >
                <div class="sub-title">{{ getRightTitle }}</div>
                <div class="box-content">
                    <div class="data-item label">
                        <span
                            class="right-name"
                            :style="{
                                width:
                                    selected === 'tmpl_detail'
                                        ? 'calc(100% - 90px)'
                                        : 'calc(100% - 50px)'
                            }"
                            >字段名称</span
                        >
                        <span
                            v-if="selected === 'tmpl_detail'"
                            style="display: inline-block; width: 60px"
                            >是否必填</span
                        >
                        <span style="width: 40px; text-align: right">移除</span>
                    </div>
                    <div id="detailFiledData">
                        <template v-if="rightData && rightData.length">
                            <div
                                class="data-item"
                                :class="itemClass(item)"
                                v-for="item in rightData"
                                :key="item.value"
                            >
                                <span
                                    v-overflow
                                    class="right-name"
                                    :style="{
                                        width:
                                            selected === 'tmpl_detail'
                                                ? 'calc(100% - 90px)'
                                                : 'calc(100% - 50px)'
                                    }"
                                >
                                    <img
                                        :src="
                                            require('@/assets/image/common/icon_dragtable_move.png')
                                        "
                                        alt=""
                                        width="20px"
                                        height="20px"
                                        style="position: relative; top: 5px"
                                    />
                                    <span class="name">{{ item.label }}</span>
                                </span>
                                <!-- <el-switch v-model="item.required"></el-switch> -->
                                <span
                                    style="display: inline-block; width: 60px"
                                    v-if="selected === 'tmpl_detail'"
                                >
                                    <el-switch
                                        v-model="item.required_flag"
                                        :active-color="PRIMARY_COLOR"
                                        inactive-color="#E6E9F0"
                                        :disabled="!fieldDisabled(item)"
                                        @change="
                                            (val) => switchChange(val, item)
                                        "
                                    >
                                    </el-switch>
                                </span>
                                <el-button
                                    style="width: 40px; text-align: right"
                                    @click="removeFiled(item)"
                                    type="text"
                                    size="small"
                                    :disabled="!fieldDisabled(item)"
                                >
                                    <i
                                        v-if="fieldDisabled(item)"
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
export default {
    props: {
        modules: {
            type: String,
            default: ""
        },
        selected: {
            type: String,
            default: ""
        }
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        fieldDisabled() {
            return (item) => {
                // 标题不可删除，不可编辑开关
                // if (row.field_key === "title" || row.field_key === "handler") {
                //     return true;
                // } else {
                //     return false;
                // }
                if (this.selected === "tmpl_detail") {
                    if (item.value !== "title" && item.value !== "handler") {
                        return true;
                    } else {
                        return false;
                    }
                } else if (this.selected === "tabMenu") {
                    if (item.value !== "tmpl_detail") {
                        return true;
                    } else {
                        return false;
                    }
                } else if (this.selected === "sub_tmpl") {
                    return true;
                }
            };
        },
        getTitle() {
            return this.selected === "tabMenu"
                ? "模块标签页设置"
                : this.selected === "tmpl_detail"
                ? "详情字段设置"
                : "子任务字段设置";
        },
        getLeftTitle() {
            return this.selected === "tabMenu"
                ? "全部标签"
                : this.selected === "tmpl_detail"
                ? "全部字段"
                : "可选流程";
        },
        getRightTitle() {
            return this.selected === "tabMenu"
                ? "模块标签"
                : this.selected === "tmpl_detail"
                ? "详情字段"
                : "已选流程";
        },
        isNotEmpty() {
            return (
                (this.selected !== "tabMenu" &&
                    this.leftData &&
                    this.leftData.length) ||
                (this.selected === "tabMenu" &&
                    this.leftData.some((item) => item.checked))
            );
        }
    },
    data() {
        return {
            search: "",
            leftData: [],
            tempLeftData: [],
            rightData: [],
            apiMapping: {
                tabMenu: "getTabConfig", // tab
                tmpl_detail: "getRestScreenList", // 详情
                sub_tmpl: "getSubWorkConfig" // 子任务
            },
            PRIMARY_COLOR: PRIMARY_COLOR_LIST.PRIMARY_COLOR
        };
    },
    watch: {
        search: {
            handler(val) {
                if (this.selected !== "sub_tmpl") {
                    this.leftData = _.cloneDeep(this.tempLeftData);
                } else {
                    if (val.trim()) {
                        let tmpVal = val.trim().toLowerCase();
                        this.leftData = this.tempLeftData.filter((item) => {
                            return (
                                item.label.toLowerCase().indexOf(tmpVal) > -1
                            );
                        });
                    } else {
                        this.leftData = _.cloneDeep(this.tempLeftData);
                    }
                }
            }
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
        selected: {
            handler(val) {
                if (val) {
                    this.search = "";
                    this.getLeftData();
                    this.getRightData();
                }
            },
            immediate: true
        },
        $route() {
            if (this.modules && this.selected) {
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
                module: "detail"
            };
            api.setIsRequired(params).then((res) => {
                if (res && res.resultCode !== 200) {
                    item.required_flag = !item.required_flag;
                } else {
                    this.getRightData();
                }
            });
            // 调接口发送更改后的数据给后端并获取新数据
        },
        itemClass(item) {
            if (this.selected === "tmpl_detail") {
                if (item.value !== "title" && item.value !== "handler") {
                    return "is-sort";
                } else {
                    return "not-sort";
                }
            } else if (this.selected === "tabMenu") {
                if (item.value !== "tmpl_detail") {
                    return "is-sort";
                } else {
                    return "not-sort";
                }
            } else {
                return "is-sort";
            }
        },
        dropFiled() {
            let tbody = document.querySelector("#detailFiledData");
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
                    let tmpKeyList = _this.rightData.map((item) => item.value);
                    let mapping = {
                        tmpl_detail: "detail",
                        tabMenu: "tab",
                        sub_tmpl: "sub_tmpl"
                    };
                    let params = {
                        ws_id: _this.curSpace.id,
                        tmpl_id: _this.$route.params.id,
                        module: mapping[_this.selected]
                    };
                    if (tmpKeyList.length) {
                        params.coordinate = tmpKeyList.join(",");
                    }
                    api.moveScreenField(params).then((res) => {
                        if (res && res.resultCode === 200) {
                            _this.getRightData();
                        }
                    });
                }
            });
        },
        getLeftData() {
            if (!this.selected) return;
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                module: "detail"
            };

            api[this.apiMapping[this.selected]](params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    res.data.length
                ) {
                    if (this.selected !== "sub_tmpl") {
                        this.leftData = this.formatLeftFiled(res.data);
                        this.tempLeftData = _.cloneDeep(this.leftData);
                    } else {
                        this.tempLeftData = this.formatLeftFiled(res.data);
                        if (this.search.trim()) {
                            let tmpVal = this.search.trim().toLowerCase();
                            this.leftData = this.tempLeftData.filter((item) => {
                                return (
                                    item.label.toLowerCase().indexOf(tmpVal) >
                                    -1
                                );
                            });
                        } else {
                            this.leftData = this.formatLeftFiled(res.data);
                        }
                    }
                } else {
                    this.leftData = [];
                    this.tempLeftData = _.cloneDeep(this.leftData);
                }
            });
        },
        // 格式化左侧全部字段
        formatLeftFiled(arr) {
            return arr.map((item) => {
                return {
                    label: item.name || item.label,
                    value: item.field_key || item.value || item.id, // 流程详情是field_key, tab取value, 子任务目取id
                    key: item.id || item.value,
                    checked: this.selected === "tabMenu" ? !item.checked : true
                };
            });
        },
        // 格式化右侧已选字段
        formatRightFiled(arr) {
            return arr.map((item) => {
                if (this.selected === "tmpl_detail") {
                    return {
                        label: item.field_name,
                        value: item.field_key,
                        id: item.id,
                        required_flag: item.required_flag
                    };
                } else if (this.selected === "tabMenu") {
                    return {
                        label: item.tab_cn,
                        value: item.tab,
                        id: item.id
                    };
                } else {
                    return {
                        label: item.sub_tmpl_name,
                        value: item.sub_tmpl_id,
                        id: item.sub_tmpl_id
                    };
                }
            });
        },
        getRightData() {
            if (!this.selected) return;
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                module: "detail"
            };
            if (this.selected === "tmpl_detail") {
                this.$store
                    .dispatch("progress_setting/fetchDetailPageFiled", params)
                    .then((res) => {
                        let resData = _.cloneDeep(res);
                        if (resData.length) {
                            resData.forEach((item) => {
                                this.$set(
                                    item,
                                    "required_flag",
                                    item.required === "no" ? false : true
                                );
                            });
                            this.rightData = this.formatRightFiled(resData);
                        } else {
                            this.rightData = [];
                        }
                    });
            } else if (this.selected === "tabMenu") {
                delete params.module;
                this.$store
                    .dispatch("progress_setting/fetchDetailTabFiled", params)
                    .then((res) => {
                        let resData = _.cloneDeep(res);
                        if (resData.length) {
                            this.rightData = this.formatRightFiled(resData);
                        } else {
                            this.rightData = [];
                        }
                    });
            } else if (this.selected === "sub_tmpl") {
                delete params.module;
                this.$store
                    .dispatch("progress_setting/fetchDetailSubFiled", params)
                    .then((res) => {
                        // return;
                        let resData = _.cloneDeep(res);
                        if (resData.length) {
                            this.rightData = this.formatRightFiled(resData);
                        } else {
                            this.rightData = [];
                        }
                    });
            }
        },
        addFiled(item) {
            if (this.selected === "tmpl_detail") {
                let params = {
                    ws_id: this.curSpace.id,
                    tmpl_id: this.$route.params.id,
                    module: "detail",
                    field_key: item.value
                };
                api.addScreen(params).then((res) => {
                    if (res && res.resultCode === 200) {
                        this.getLeftData();
                        this.getRightData();
                        this.$emit("filed-change");
                    }
                });
            } else {
                let mapping = {
                    tabMenu: "tabCreate",
                    sub_tmpl: "subWorkCreate"
                };
                let params = {
                    ws_id: this.curSpace.id,
                    tmpl_id: this.$route.params.id
                };
                if (this.selected === "sub_tmpl") {
                    params.sub_tmpl_id = item.value;
                } else {
                    params.tab = item.value;
                }
                api[mapping[this.selected]](params).then((res) => {
                    if (res && res.resultCode === 200) {
                        this.getLeftData();
                        this.getRightData();
                        this.$emit("filed-change");
                    }
                });
            }
        },
        removeFiled(item) {
            if (this.selected === "tmpl_detail") {
                let params = {
                    ws_id: this.curSpace.id,
                    tmpl_id: this.$route.params.id,
                    module: "detail",
                    id: item.id
                };
                api.delScreen(params).then((res) => {
                    if (res && res.resultCode === 200) {
                        this.getLeftData();
                        this.getRightData();
                        this.$emit("filed-change");
                    }
                });
            } else {
                let mapping = {
                    tabMenu: "tabDelete",
                    sub_tmpl: "subWorkDelete"
                };
                let params = {
                    ws_id: this.curSpace.id,
                    tmpl_id: this.$route.params.id
                };
                if (this.selected === "sub_tmpl") {
                    params.sub_tmpl_id = item.value;
                } else {
                    params.tab = item.value;
                }
                api[mapping[this.selected]](params).then((res) => {
                    if (res && res.resultCode === 200) {
                        this.getLeftData();
                        this.getRightData();
                        this.$emit("filed-change");
                    }
                });
            }
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
            box-sizing: border-box;
            height: 100%;
            padding: 12px;
            background-color: #fafafa;
            border-radius: 0px 0px 0px 0px;
            border: 1px dashed #a6abbc;
            .overflow-y {
                overflow-y: auto;
                padding-right: 12px;
                margin-right: -11px;
            }
            .el-input {
                margin-bottom: 8px;
            }
        }
        .data-item {
            height: 32px;
            display: flex;
            align-items: center;
            justify-content: space-between;
            .left-name {
                display: inline-block;
                width: calc(100% - 30px);
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
            }
            .right-name {
                display: inline-block;
                width: calc(100% - 90px);
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
        #detailFiledData {
            height: calc(100% - 32px);
            overflow-y: auto;
            padding-right: 16px;
            margin-right: -11px;
        }
    }
    .el-icon-close:hover:before {
        cursor: pointer;
        content: url("~@/assets/image/progress_setting/page/remove_hover.svg");
    }
    .el-icon-close:before {
        cursor: pointer;
        content: url("~@/assets/image/progress_setting/page/remove.svg");
    }

    .el-icon-plus:hover:before {
        cursor: pointer;
        content: url("~@/assets/image/progress_setting/page/add_hover.svg");
    }
    .el-icon-plus:before {
        cursor: pointer;
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
