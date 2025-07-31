<template>
    <div class="preview-box-left">
        <div class="title">
            <span>{{ title }}</span>
            <i class="el-icon-close"></i>
        </div>
        <div class="header-info">
            <div class="info-line">
                <div style="margin-right: 24px">
                    <span class="lable">
                        <span></span>
                        <span>状态：</span>
                    </span>
                    <span class="value">
                        <span class="status">进行中</span>
                    </span>
                </div>
                <div>
                    <span class="lable">
                        <span></span>
                        <span>处理人：</span>
                    </span>
                    <span class="value">
                        <el-avatar
                            class="progress-avatar"
                            icon="el-icon-user-solid"
                            size="small"
                            :src="getAvatar('')"
                        ></el-avatar>
                        <span>张三</span>
                    </span>
                </div>
            </div>
            <div class="info-line bg">
                <div class="msg-item" style="max-width: calc(26% - 32px)">
                    <span class="msg-item-title">创建人：</span>
                    <span class="msg-item-content" v-overflow>
                        <el-avatar
                            class="progress-avatar"
                            icon="el-icon-user-solid"
                            size="small"
                            :src="getAvatar('')"
                        ></el-avatar>
                        张三</span
                    >
                </div>
                <div class="msg-item" style="max-width: calc(37% - 32px)">
                    <span class="msg-item-title">创建时间：</span>
                    <span class="msg-item-content" v-overflow
                        >2023-06-08 12:38:46</span
                    >
                </div>
                <div class="msg-item" style="max-width: 37%">
                    <span class="msg-item-title">最后修改时间：</span>
                    <span class="msg-item-content" v-overflow
                        >2023-06-08 12:38:46</span
                    >
                </div>
            </div>
        </div>

        <div class="main-info">
            <div class="left">
                <div
                    class="tabs-box"
                    :class="currentSelect === 'tabMenu' ? 'active' : ''"
                    @click="onClickTabs"
                >
                    <el-tabs
                        class="basic-ui font"
                        :class="
                            tabsList && tabsList.length === 1 ? 'only-one' : ''
                        "
                        size="small"
                        v-model="activeName"
                        @click.native.stop="handleClick(activeName)"
                    >
                        <el-tab-pane
                            v-for="(item, index) in tabsList"
                            :label="item.label"
                            :name="item.value"
                            :key="index"
                            :disabled="item.value === 'file' ? true : false"
                        ></el-tab-pane>
                    </el-tabs>
                </div>
                <div class="body" v-if="activeName === 'tmpl_detail'">
                    <div
                        class="form-item"
                        v-for="(item, index) in detailFiled"
                        :key="index"
                    >
                        <span
                            class="label"
                            :class="{ required: item.required === 'yes' }"
                            ><b
                                class="type-box"
                                :style="getType(item.mode, item.expr)"
                            ></b
                            ><span>{{ item.name }}:</span></span
                        >
                        <span class="value"> -- </span>
                        <!-- 进度条 -->
                        <!-- <span
                            class="value progress"
                            v-if="item.mode === 'progress_com'"
                        >
                            <div class="column-block flex-content">
                                <div class="el-slider silder-block">
                                    <div class="el-slider__runway">
                                        <div
                                            class="el-slider__bar"
                                            style="width: 0%; left: 0%"
                                        ></div>
                                        <div
                                            tabindex="0"
                                            class="el-slider__button-wrapper"
                                            style="left: 0%"
                                        >
                                            <div
                                                class="el-tooltip el-slider__button"
                                                aria-describedby="el-tooltip-2058"
                                                tabindex="0"
                                            ></div>
                                        </div>
                                    </div>
                                </div>
                                <div class="num-block">0%</div>
                            </div>
                        </span> -->
                        <!-- Html编辑器 -->
                        <!-- <div
                            class="value html-img"
                            v-else-if="item.mode === 'html_text_com'"
                        >
                            <img
                                :src="
                                    require(`@/assets/image/progress_column/col_html_empty.png`)
                                "
                                :width="20"
                                :height="20"
                            />
                        </div> -->
                        <!-- <span class="value" v-else>
                            {{ getPlaceholder(item) }}
                            <i
                                class="form-item-icon"
                                :class="iconClass(item)"
                            ></i>
                        </span> -->
                    </div>
                </div>
                <div class="body sub-work" v-if="activeName === 'sub_tmpl'">
                    <template v-if="subWorkArr && subWorkArr.length">
                        <div
                            class="table-item"
                            v-for="(item, index) in subWorkArr"
                            :key="index"
                        >
                            <div class="title-line">
                                <span
                                    style="
                                        display: flex;
                                        width: calc(100% - 96px);
                                    "
                                >
                                    <span class="name" v-overflow>{{
                                        item.sub_tmpl_name
                                    }}</span>
                                    <i class="el-icon-arrow-down"></i>
                                </span>
                                <span style="color: #1890ff">
                                    <i class="plus"></i>
                                    添加子任务
                                </span>
                            </div>
                            <NTable
                                ref="ProgressTable"
                                :data="tableData"
                                :col="tableCol"
                            >
                                <template slot="status" slot-scope="{ column }">
                                    <el-table-column
                                        :min-width="column.minWidth"
                                        :width="column.width"
                                        :align="column.align || 'left'"
                                    >
                                        <template slot="header">
                                            <div>
                                                <b
                                                    class="type-box"
                                                    :style="
                                                        getType(
                                                            column.mode,
                                                            column.expr
                                                        )
                                                    "
                                                ></b>
                                                {{ column.name }}
                                            </div>
                                        </template>
                                        <template slot-scope="scope">
                                            <span class="status-col"
                                                >{{
                                                    scope.row[column.prop] ||
                                                    "--"
                                                }}
                                            </span>
                                        </template>
                                    </el-table-column>
                                </template>
                            </NTable>
                        </div>
                    </template>
                    <!-- <span v-else class="empty">暂无内容</span> -->
                    <span v-else class="sub-tmpl-empty">
                        <div class="top-line">
                            <span class="name" v-overflow>子任务</span>
                            <span style="color: #1890ff">
                                <i class="plus"></i>
                                添加子任务
                            </span>
                        </div>
                        <div class="gray">暂无子任务</div>
                        <div class="bottom-line"></div>
                    </span>
                </div>
            </div>
            <div class="right">
                <div style="height: calc(100% - 60px); overflow-y: auto">
                    <el-tabs
                        class="basic-ui font"
                        size="small"
                        v-model="activeNameRight"
                    >
                        <el-tab-pane
                            v-for="(item, index) in tabsListRight"
                            :label="item.label"
                            :name="item.name"
                            :key="index"
                            :disabled="true"
                        ></el-tab-pane>
                    </el-tabs>
                    <div>
                        <div
                            v-for="(item, index) in progressArr"
                            :key="index"
                            class="evolve-item"
                        >
                            <div class="info-block">
                                <div class="block-left" v-overflow>
                                    <el-avatar
                                        class="progress-avatar"
                                        icon="el-icon-user-solid"
                                        size="small"
                                        :src="getAvatar(item.creator.avatar)"
                                    ></el-avatar>
                                    {{ item.creator.full_name }}
                                </div>
                                <div class="block-right">
                                    <span class="time" v-overflow>
                                        {{ item.created_at }}
                                    </span>
                                    <b class="tabs-partition"></b>
                                </div>
                            </div>
                            <div class="msg-block">
                                <span
                                    ><span v-if="item.name" class="mention-at"
                                        >@{{ item.name }}</span
                                    >{{ item.info }}</span
                                >
                            </div>
                        </div>
                    </div>
                </div>
                <div class="bottom">
                    <div class="input-box">请输入进展</div>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import { imgHost } from "@/assets/tool/const";
import NTable from "@/pages/common/tables/common_table/table";
import { FieldType } from "@/assets/tool/const";
import { mapState } from "vuex";
import api from "@/common/api/module/progress";
export default {
    components: {
        NTable
    },
    props: {
        modules: {
            type: String,
            default: ""
        }
    },
    data() {
        return {
            currentSelect: "tabMenu",
            subWorkArr: [],
            progressArr: [
                {
                    creator: { full_name: "张三", avatar: "" },
                    created_at: "2023-06-08 12:38:46",
                    info: "目前进展一切顺利"
                },
                {
                    creator: { full_name: "张三", avatar: "" },
                    created_at: "2023-06-08 12:38:46",
                    name: "张三",
                    info: "目前进展一切顺利"
                }
            ],
            tableData: [
                {
                    title: "名称",
                    status: "开发中",
                    handler: "张三",
                    creator: "李四",
                    creat_time: "2025-6-20"
                }
            ],
            tableCol: [
                {
                    prop: "title",
                    name: "标题",
                    minWidth: 80,
                    iconCol: true,
                    mode: "text_com"
                },
                {
                    prop: "status",
                    name: "状态",
                    mode: "status_com",
                    minWidth: 80,
                    type: "slot"
                },
                {
                    prop: "handler",
                    name: "处理人",
                    minWidth: 92,
                    iconCol: true,
                    mode: "person_com",
                    expr: "single"
                },
                {
                    prop: "creator",
                    name: "创建人",
                    minWidth: 92,
                    iconCol: true,
                    mode: "person_com",
                    expr: "single"
                },
                {
                    prop: "creat_time",
                    name: "创建时间",
                    minWidth: 102,
                    iconCol: true,
                    mode: "date_com"
                }
            ],
            title: "文档详情编辑后，保存不了",
            activeName: "",
            tabsList: [],
            activeNameRight: "progress",
            tabsListRight: [
                {
                    label: "进展(2)",
                    name: "progress"
                },
                {
                    label: "日志",
                    name: "log"
                }
            ],
            detailFiled: []
        };
    },
    computed: {
        ...mapState("progress_setting", [
            "detailPageFiled",
            "detailTabFiled",
            "detailSubFiled"
        ]),
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        },
        getPlaceholder() {
            return (item) => {
                let placeholder = "";
                if (item.field_key === "title" || item.mode === "number_com") {
                    placeholder = "请输入内容";
                } else if (item.mode === "person_com") {
                    placeholder = `请选择${
                        item.expr === "single" ? "(单人)" : "(多人)"
                    }`;
                } else if (
                    item.mode === "textarea_com" ||
                    item.mode === "text_com"
                ) {
                    placeholder = "请输入";
                } else if (item.mode === "date_com") {
                    placeholder = "请选择日期";
                } else if (item.mode === "time_com") {
                    placeholder = "请选择时间";
                } else if (item.mode === "link_com") {
                    placeholder = "请编辑";
                } else if (
                    item.mode === "multi_select_com" ||
                    item.mode === "select_com"
                ) {
                    placeholder = "请选择";
                } else if (item.mode === "related_com") {
                    placeholder = "请选择关联流程";
                }
                return placeholder;
            };
        },
        iconClass() {
            return (item) => {
                let icon = "";
                if (
                    item.mode === "person_com" ||
                    item.mode === "multi_select_com" ||
                    item.mode === "select_com"
                ) {
                    icon = "el-icon-caret-bottom";
                } else if (
                    item.mode === "time_com" ||
                    item.mode === "date_com"
                ) {
                    icon = "time";
                }
                return icon;
            };
        }
    },
    watch: {
        // 映射 user 模块的 state 和 getters
        detailPageFiled: {
            handler() {
                this.fetchDetailPage();
            },
            deep: true,
            immediate: true
        },
        // 映射 user 模块的 state 和 getters
        detailSubFiled: {
            handler(arr) {
                this.subWorkArr = arr;
            },
            deep: true,
            immediate: true
        },
        detailTabFiled: {
            handler(arr) {
                if (arr && arr.length) {
                    this.tabsList = arr.map((item) => {
                        return {
                            label: item.tab_cn,
                            value: item.tab,
                            id: item.id
                        };
                    });
                } else {
                    this.tabsList = [
                        {
                            label: "流程详情",
                            value: "tmpl_detail"
                        }
                    ];
                }
                if (
                    this.tabsList
                        .map((item) => item.value)
                        .includes(this.activeName)
                )
                    return;
                this.activeName = this.tabsList[0].value;
            },
            deep: true,
            immediate: true
        },
        currentSelect: {
            handler(val) {
                if (this.currentSelect === "file" || this.activeName === "file")
                    return;
                this.$emit("select-change", val);
            },
            immediate: true
        },
        module: {
            handler() {}
        },
        $route() {
            this.currentSelect = "tabMenu";
            if (this.currentSelect === "file" || this.activeName === "file")
                return;
            this.$emit("select-change", this.currentSelect);
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                module: this.modules
            };
            setTimeout(() => {
                if (this.activeName === "tmpl_detail") {
                    this.fetchDetailPage();
                } else if (this.activeName === "sub_tmpl") {
                    this.$store.dispatch(
                        "progress_setting/fetchDetailSubFiled",
                        params
                    );
                }
            }, 0);
        },
        activeName(val) {
            if (this.currentSelect === "file" || val === "file") return;
            this.$emit("select-change", this.currentSelect);
        }
    },
    methods: {
        getType(type, expr = "") {
            if (!type) return;
            if ((type && !expr) || type === "number_com") {
                return {
                    "background-image": `url(${FieldType[type]})`
                };
            }
            if (type && expr) {
                return {
                    "background-image": `url(${FieldType[type + "_" + expr]})`
                };
            }
        },
        fetchDetailPage() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                module: "detail"
            };
            api.getProgressColConfig(params).then((res) => {
                if (this.$route.name !== "progress") return;
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    this.detailFiled = res.data.filter((item) => {
                        return (
                            item.field_key !== "title" &&
                            item.field_key !== "handler" &&
                            item.field_key !== "creator" &&
                            item.field_key !== "status" &&
                            item.field_key !== "updated_at" &&
                            item.field_key !== "created_at"
                        );
                    });
                } else {
                    this.detailFiled = [];
                }
            });
        },

        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        },
        // 点击顶部的整个tabsbar
        onClickTabs() {
            this.currentSelect = "tabMenu";
        },
        // 点击tab item
        handleClick(tab) {
            if (tab === "file") return;
            this.currentSelect = tab;
        }
    }
};
</script>
<style lang="scss" scoped>
.preview-box-left {
    box-sizing: border-box;
    height: calc(100% - 100px);
    width: calc(100% - 60px);
    padding: 24px;
    background: #ffffff;
    box-shadow: 0px 3px 10px 1px rgba(0, 0, 0, 0.1);
    border-radius: 4px 4px 4px 4px;
    .type-box {
        display: inline-block;
        width: 20px;
        height: 20px;
        background-size: 100% 100%;
        vertical-align: middle;
        position: relative;
        top: -2px;
    }
    .status-col {
        display: inline-block;
        height: 21px;
        background: #b4f298;
        border-radius: 4px 4px 4px 4px;
        padding: 0 4px;
        line-height: 20px;
    }
    ::v-deep .el-tabs__item.is-disabled {
        color: #303133;
        cursor: default;
    }
    ::v-deep .el-tabs__item.is-disabled.is-active {
        color: #409eff;
    }
    .title {
        height: 30px;
        margin-bottom: 20px;
        font-size: 16px;
        font-weight: 500;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
    .header-info {
        .progress-avatar.el-avatar {
            position: relative;
            top: -2px;
            width: 20px;
            height: 20px;
            margin-right: 2px;
        }
        .info-line {
            display: flex;
            margin-bottom: 14px;
            .label {
                color: #82889e;
            }
        }

        .status {
            display: inline-block;
            height: 21px;
            background: #b4f298;
            border-radius: 4px 4px 4px 4px;
            padding: 0 4px;
            line-height: 20px;
            margin-right: 32px;
        }
        .msg-item {
            display: flex;
            align-items: center;
            white-space: nowrap;
            margin-right: 32px;
            height: 32px;
            &:last-child {
                margin-right: 0;
            }
            .msg-item-title {
                color: #82889e;
                display: inline-block;
                height: 32px;
                line-height: 32px;
                vertical-align: top;
            }
            .msg-item-content {
                color: #2a3048;
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
                &.create-time {
                    width: calc(100% - 72px);
                }
                &.last-time {
                    width: calc(100% - 100px);
                }
            }
        }
        .bg {
            box-sizing: border-box;
            width: 100%;
            height: 30px;
            line-height: 30px;
            padding: 0 8px;
            background: #f5f6f8;
            border-radius: 4px 4px 4px 4px;
        }
    }
    .main-info {
        display: flex;
        height: calc(100% - 50px - 64px);
        margin-top: 20px;
        border-top: 1px solid #e6e9f0;
        .left {
            box-sizing: border-box;
            border-right: 1px solid #e6e9f0;
            width: 60%;
            padding: 1px 0;
        }
        .tabs-box {
            box-sizing: border-box;
            height: 50px;
            line-height: 11px;
            padding: 11px 12px 6px;
            border: 1px dashed #cdd5e6;
            margin-bottom: 12px;
            margin-right: 12px;
            cursor: pointer;
            cursor: pointer;
            &:hover {
                border: 1px dashed #1890ff;
            }
            &.active {
                border: 1px dashed #1890ff;
            }
            ::v-deep .el-tabs {
                display: inline-block;
                max-width: 100%;
            }
            ::v-deep .el-tabs__nav-wrap::after {
                background-color: transparent;
            }
            ::v-deep .el-tabs__header {
                margin: 0 !important;
            }
            ::v-deep .el-tabs__nav-next,
            ::v-deep .el-tabs__nav-prev {
                line-height: 26px !important;
            }
            ::v-deep
                .basic-ui.only-one.el-tabs.el-tabs--top
                .el-tabs__header
                .el-tabs__item.is-bottom:nth-child(2),
            ::v-deep
                .basic-ui.only-one.el-tabs.el-tabs--top
                .el-tabs__header
                .el-tabs__item.is-top:nth-child(2) {
                padding: 0;
            }
        }
        .body {
            height: calc(100% - 42px - 48px);
            padding: 12px;
            // padding-right: 20px;
            // margin-right: 12px;
            overflow-y: auto;
            &.sub-work {
                height: calc(100% - 72px);
                padding: 0 12px;
                .empty {
                    position: relative;
                    display: inline-block;
                    top: 50%;
                    left: 50%;
                    transform: translate(-50%, -50%);
                    color: #a6abbc;
                }
                .sub-tmpl-empty {
                    .top-line {
                        display: flex;
                        justify-content: space-between;
                        margin-bottom: 18px;
                        .name {
                            display: inline-block;
                            max-width: calc(100% - 100px);
                        }
                        .plus {
                            position: relative;
                            top: 4px;
                            display: inline-block;
                            width: 20px;
                            height: 20px;
                            background-size: 100% 100%;
                            background-image: url(@/assets/image/common/icon_plus.png);
                        }
                    }
                    .gray {
                        color: #a6abbc;
                    }
                    .bottom-line {
                        height: 1px;
                        margin-top: 10px;
                        background-color: #e6e9f0;
                    }
                }
            }
            .table-item {
                margin-bottom: 16px;
                .title-line {
                    display: flex;
                    justify-content: space-between;
                    align-items: center;
                    margin-bottom: 16px;
                }
                .name {
                    display: inline-block;
                    max-width: calc(100% - 115px);
                    margin-right: 4px;
                    white-space: nowrap;
                    overflow: hidden;
                    text-overflow: ellipsis;
                }
                .el-icon-arrow-down:before {
                    position: relative;
                    top: 5px;
                    color: #66677c;
                }
                .plus {
                    position: relative;
                    top: 4px;
                    display: inline-block;
                    width: 20px;
                    height: 20px;
                    background-size: 100% 100%;
                    background-image: url(@/assets/image/common/icon_plus.png);
                }
            }
        }
        .form-item {
            display: flex;
            align-items: flex-start;
            margin-bottom: 16px;
            &:last-child {
                margin-bottom: 0;
            }
            .label {
                position: relative;
                display: inline-block;
                width: 160px;
                text-align: left;
                // height: 32px;
                line-height: 32px;
                margin-right: 16px;
                word-break: break-all;
                color: #5c6479;
                &.required::before {
                    position: absolute;
                    content: "*";
                    color: red;
                }
                .type-box {
                    display: inline-block;
                    width: 20px;
                    height: 20px;
                    margin-left: 6px;
                    margin-right: 3px;
                    background-size: 100% 100%;
                    vertical-align: middle;
                    position: relative;
                    top: -2px;
                }
            }
            .value {
                position: relative;
                display: inline-block;
                box-sizing: border-box;
                width: calc(100% - 176px);
                max-width: 300px;
                padding: 0 16px;
                height: 32px;
                line-height: 32px;
                border-radius: 4px 4px 4px 4px;
                // border: 1px solid #e6e9f0;
                color: #2a3048;
                &.html-img {
                    border: none;
                    padding: 0;
                    img {
                        vertical-align: middle;
                        cursor: pointer;
                    }
                }
                &.progress {
                    border: none;
                    padding: 0 0 0 12px;
                    .flex-content {
                        display: flex;
                        width: 100%;
                        .el-slider.silder-block {
                            width: calc(100% - 50px);
                        }
                        .el-slider__button-wrapper {
                            top: -14px;
                            &:hover {
                                cursor: default;
                            }
                            .el-tooltip.el-slider__button {
                                box-sizing: border-box;
                                width: 12px;
                                height: 12px;
                                border: 2px solid #10a862;
                                background-color: #10a862;
                                transition: 0;
                                &:hover {
                                    transform: none;
                                    cursor: default;
                                }
                            }
                        }
                        .num-block {
                            width: 50px;
                            text-align: center;
                            color: #10a862;
                            font-size: 12px;
                        }
                        .el-slider__runway {
                            box-sizing: border-box;
                            height: 10px;
                            margin: 11px 0;
                            background-color: rgba(0, 0, 0, 0);
                            border: 1px solid #10a862;
                            border-radius: 5px;
                            cursor: default;
                        }
                    }
                }
            }
            .form-item-icon {
                display: inline-block;
                position: absolute;
                height: 32px;
                line-height: 32px;
                right: 16px;
                &.time {
                    width: 16px;
                    background-size: 100% 100%;
                    margin-right: 2px;
                    background-image: url("~@/assets/image/common/time.svg");
                }
            }
        }
        .right {
            position: relative;
            box-sizing: border-box;
            width: 40%;
            padding: 13px 12px;
            .evolve-item {
                padding: 12px;
                margin-bottom: 12px;
                border-radius: 4px;
                background-color: #fafafb;
                .info-block {
                    height: 32px;
                    line-height: 32px;
                    display: flex;
                    justify-content: space-between;
                    .block-left {
                        width: 80px;
                        font-weight: 500;
                        font-size: 14px;
                        color: #2f384c;
                        white-space: nowrap;
                        overflow: hidden;
                        text-overflow: ellipsis;
                    }
                    .block-right {
                        width: calc(100% - 80px);
                        text-align: right;
                        font-weight: 400;
                        font-size: 14px;
                        color: #95979d;
                        display: flex;
                    }
                    .time {
                        display: inline-block;
                        width: calc(100% - 20px);
                        margin-right: 2px;
                        white-space: nowrap;
                        overflow: hidden;
                        text-overflow: ellipsis;
                    }
                    .tabs-partition {
                        position: relative;
                        top: 5px;
                        left: 3px;
                        display: inline-block;
                        width: 20px;
                        height: 20px;
                        vertical-align: sub;
                        background-image: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAYAAACM/rhtAAAAAXNSR0IArs4c6QAAANpJREFUWEftlDEKwkAURP9Hg3bGG3kD7bVIkQNop1262HkBAyuYPkeLXUDCF60EFwlMFhKYrWfg8XZ2VQZ+dOB8QkD0hmiQBlEDaJ8bpEHUANrnBmkQNYD2x7nBJMni56S5qOhGVGIzc1E7PziX1T4jIfNeg9v0WKno+hvGTG5lkSc+wJB5L+AuPdkPiEl9L/KlDzBkvjOgiT3K6znuCthXfpxX/Bn9tHEislLRxXt/UTvb/30kgfLj/GbQz7XPPg2iNmmQBlEDaJ8bpEHUANrnBmkQNYD2B7/BF48alimY6dGEAAAAAElFTkSuQmCC);
                        background-size: 100% 100%;
                    }
                }
                .msg-block {
                    // padding: 0 6px 0 44px;
                    padding: 0 6px 0 0px;
                    font-weight: 400;
                    font-size: 14px;
                    color: #2f384c;
                    margin-top: 12px;
                    line-height: 1.5;
                    word-break: break-word;
                    .mention-at {
                        display: inline-block;
                        box-sizing: border-box;
                        color: #1890ff;
                        margin: 2px;
                        background-color: #e3f0fc;
                        border-radius: 4px;
                        padding: 0 4px;
                        height: 20px;
                        line-height: 18px;
                        font-size: 14px;
                    }
                }
            }
            .bottom {
                width: calc(100% + 24px);
                position: relative;
                left: -12px;
                height: 58px;
                display: flex;
                align-items: center;
                justify-content: center;
                border-top: 1px solid #e6e9f0;
                .input-box {
                    box-sizing: border-box;
                    height: 32px;
                    width: calc(100% - 20px);
                    border: 1px solid #cdd5e6;
                    line-height: 32px;
                    border-radius: 4px;
                    color: #a6abbc;
                    padding: 0 10px;
                    box-sizing: border-box;
                }
            }
        }
    }
    .footer {
        box-sizing: border-box;
        padding-top: 20px;
        height: 60px;
        text-align: right;
        .footer-btn {
            display: inline-block;
            line-height: 1;
            white-space: nowrap;
            cursor: pointer;
            background: #fff;
            border: 1px solid #dcdfe6;
            color: #606266;
            text-align: center;
            box-sizing: border-box;
            outline: 0;
            margin: 0;
            transition: 0.1s;
            font-weight: 500;
            padding: 12px 20px;
            font-size: 14px;
            border-radius: 4px;
        }
        .primay {
            color: #fff;
            background-color: var(--PRIMARY_COLOR);
            border-color: var(--PRIMARY_COLOR);
            margin-left: 10px;
        }
    }
}
</style>
