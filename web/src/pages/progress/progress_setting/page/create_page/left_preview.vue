<template>
    <div class="preview-box-left">
        <div class="title">
            <span>新增</span>
            <i class="el-icon-close"></i>
        </div>
        <div class="body">
            <!-- {{ createFiled }} -->
            <div
                class="form-item"
                v-for="(item, index) in createFiled"
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
                <!-- 进度条 -->
                <span
                    class="value progress"
                    v-if="item.mode === 'progress_com'"
                >
                    <div class="column-block flex-content">
                        <div class="el-slider silder-block">
                            <!---->
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
                                <!----><!---->
                            </div>
                        </div>
                        <div class="num-block">0%</div>
                    </div>
                </span>
                <!-- Html编辑器 -->
                <div
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
                </div>
                <span class="value" v-else>
                    {{ getPlaceholder(item) }}
                    <i class="form-item-icon" :class="iconClass(item)"></i>
                </span>
            </div>
        </div>
        <div class="footer">
            <span>
                <span class="footer-btn">取消</span>
                <span class="footer-btn primay">确定</span>
            </span>
        </div>
    </div>
</template>
<script>
import { FieldType } from "@/assets/tool/const";
import { mapState } from "vuex";
import api from "@/common/api/module/progress";
export default {
    data() {
        return {
            createFiled: []
        };
    },

    computed: {
        // 映射 user 模块的 state 和 getters
        ...mapState("progress_setting", ["createPageFiled"]),
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
        createPageFiled: {
            handler() {
                this.fetchCreatePage();
            },
            deep: true,
            immediate: true
        }
    },
    // mounted() {},
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
        fetchCreatePage() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                module: "create"
            };
            api.getProgressColConfig(params).then((res) => {
                if (this.$route.name !== "progress") return;
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    this.createFiled = res.data;
                } else {
                    this.createFiled = [];
                }
            });
        }
    }
};
</script>
<style lang="scss" scoped>
.preview-box-left {
    box-sizing: border-box;
    height: calc(100% - 100px);
    width: calc(100% - 60px);
    max-width: 540px;
    padding: 24px;
    background: #ffffff;
    box-shadow: 0px 3px 10px 1px rgba(0, 0, 0, 0.1);
    border-radius: 4px 4px 4px 4px;
    .title {
        height: 30px;
        margin-bottom: 20px;
        font-size: 16px;
        font-weight: 500;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
    .body {
        height: calc(100% - 42px - 60px);
        // overflow-y: auto;
        // padding: 10px 0;
        padding-right: 20px;
        overflow-y: auto;
        .label {
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
            width: 140px;
            // height: 32px;
            line-height: 32px;
            margin-right: 16px;
            word-break: break-all;
            color: #5c6479;
        }
        .value {
            position: relative;
            display: inline-block;
            box-sizing: border-box;
            width: calc(100% - 156px);
            // max-width: 300px;
            padding: 0 16px;
            height: 32px;
            line-height: 32px;
            border-radius: 4px 4px 4px 4px;
            border: 1px solid #e6e9f0;
            color: #a6abbc;
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
                background-image: url("~@/assets/image/common/time.svg");
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
