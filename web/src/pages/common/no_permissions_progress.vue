<template>
    <div class="no-data">
        <div class="top">
            <div class="top-content">
                <!-- <span :class="['icon-type', this.progressInfo.tmpl.mode]">{{
                    tmplMode
                }}</span> -->
                <!-- <b :class="['icon-type', this.progressInfo.tmpl.mode]"></b> -->
                <span class="tmpl-name">{{
                    progressInfo.tmpl.name || "--"
                }}</span>
                <!-- <div class="space-admin">
                    <span class="label">
                        <span class="flex">
                            <img
                                :src="
                                    require('@/assets/image/field_type_icon/people_multiple.svg')
                                "
                                :width="16"
                                :height="16"
                            />
                            <span>超级管理员:</span>
                        </span>
                    </span>
                    <span class="value">
                        <template
                            v-if="
                                progressInfo.admin_list &&
                                progressInfo.admin_list.length
                            "
                        >
                            <span
                                v-for="(item, index) in progressInfo.admin_list"
                                :key="index"
                                class="value-item"
                            >
                                <user-message :userMessage="item">
                                    <span class="flex">
                                        <el-avatar
                                            icon="el-icon-user-solid"
                                            size="small"
                                            :src="getAvatar(item)"
                                        ></el-avatar>
                                        <span>{{
                                            item.full_name || "--"
                                        }}</span>
                                    </span>
                                </user-message>
                            </span>
                        </template>
                        <template v-else>--</template>
                    </span>
                </div> -->
            </div>
            <div class="desc">{{ this.progressInfo.tmpl.desc || "--" }}</div>
        </div>
        <div class="content-img" :style="height">
            <div class="sub-box">
                <img
                    :src="require(`@/assets/image/space/no_space_icon.png`)"
                    alt=""
                />
                <div class="no-data-text">暂无权限, 请联系空间管理员</div>
            </div>
        </div>
    </div>
</template>

<script>
import { imgHost } from "@/assets/tool/const";
// import api from '@/common/api/module/space';
import api from "@/common/api/module/progress";
import TipMore from "@/pages/common/tooltip_more_line.vue";
import TipOne from "@/pages/common/tooltip_one_line.vue";
import { TITLE_COLOT_ARR } from "@/assets/tool/color_list";
import UserMessage from "@/components/user_message_tip";
export default {
    components: {
        TipMore,
        TipOne,
        UserMessage
    },
    props: {
        height: {
            type: Number,
            default: 300
        }
    },
    computed: {
        tmplMode() {
            if (this.progressInfo.tmpl && this.progressInfo.tmpl.mode) {
                let mode = this.progressInfo.tmpl.mode;
                return mode === "public" ? "公" : "私";
            } else {
                return "--";
            }
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        },
        getAvatar() {
            return (item) => {
                if (item.avatar) {
                    return `${imgHost}${item.avatar}`;
                }
                return require(`@/assets/image/common/default_avatar_big.png`);
            };
        },
        firstLetter() {
            return (name) => {
                return name ? name.substring(0, 1) : "";
            };
        },
        titleStyle() {
            let obj = {};
            if (this.workSpaceList && this.workSpaceList.length) {
                this.workSpaceList.forEach((item, index) => {
                    if (item.name === this.curSpace.name) {
                        let colorItem =
                            this.colorList[index % this.colorList.length];
                        obj = {
                            color: colorItem.color,
                            backgroundColor: colorItem.backgroundColor
                        };
                    }
                });
            } else {
                obj = {
                    color: "",
                    borderColor: ""
                };
            }
            return obj;
        }
    },

    data() {
        return {
            colorList: TITLE_COLOT_ARR,
            progressInfo: {
                admin_list: {},
                tmpl: {}
            }
        };
    },
    watch: {
        curProgress() {
            console.log(666);
            this.getProgressInfo();
        }
    },
    mounted() {
        this.getProgressInfo();
    },
    methods: {
        getProgressInfo() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress
            };
            api.getProgressInfo(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    this.progressInfo.admin_list = res.data.admin_list || [];
                    this.progressInfo.tmpl = res.data.tmpl || {};
                } else {
                    this.progressInfo.admin_list = [];
                    this.progressInfo.tmpl = {};
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.no-data {
    // position: relative;
    // height: 100%;
    // width: 100%;

    // ::v-deep .el-avatar {
    //     margin-right: 8px;
    //     height: 16px;
    //     width: 16px;
    //     line-height: 16px;
    // }
    // .top-img {
    //     height: 220px;
    //     background: url('@/assets/image/common/space_homepage_top_bg.png');
    //     background-size: 100% 100%;
    // }
    // .flex {
    //     display: flex;
    //     align-items: center;
    // }
    // .desc {
    //     color: #a6abbc;
    //     font-size: 12px;
    //     margin-top: 4px;
    // }
    // .box {
    //     margin: 16px 88px 0 264px;
    //     position: relative;
    //     .sub-box {
    //         position: absolute;
    //         top: -48px;
    //         left: -150px;
    //         height: 118px;
    //         line-height: 118px;
    //         width: 118px;
    //         text-align: center;
    //         font-size: 56px;
    //         border: 3px solid #fff;
    //         border-radius: 8px;
    //         background-color: #e5ebfd;
    //         color: #1890ff;
    //         box-shadow: 0px 3px 10px 1px rgba(0, 0, 0, 0.1);
    //     }
    //     .space-name {
    //         width: 100%;
    //         font-size: 24px;
    //         font-weight: 600;
    //         margin-bottom: 10px;
    //         overflow: hidden;
    //         text-overflow: ellipsis;
    //         white-space: nowrap;
    //     }
    //     .space-desc {
    //         width: 100%;
    //         line-height: 20px;
    //         font-size: 12px;
    //         color: #a6abbc;
    //         margin-bottom: 10px;
    //         display: -webkit-box; //将对象作为弹性伸缩盒子模型显示。
    //         -webkit-box-orient: vertical; //从上到下垂直排列子元素（设置伸缩盒子的子元素排列方式）
    //         -webkit-line-clamp: 2; //限制行数
    //         overflow: hidden; //超出部分隐藏
    //         text-overflow: ellipsis; //用一个省略号代替超出的内容
    //     }
    //     .space-admin {
    //         display: flex;
    //         align-items: start;
    //         padding: 17px 24px 6px;
    //         background-color: #f8f8f8;
    //         border-radius: 8px;
    //         margin-bottom: 16px;
    //         .label {
    //             display: inline-block;
    //             margin-bottom: 10px;
    //             width: 116px;
    //         }
    //         .value {
    //             display: inline-block;
    //             width: calc(100% - 116px);
    //         }
    //         .value-item {
    //             display: inline-block;
    //             margin: 0 12px 10px 0;
    //         }
    //     }
    //     .space-num {
    //         display: inline-block;
    //         height: 52px;
    //         line-height: 52px;
    //         padding: 0 24px;
    //         margin: 0 16px 16px 0;
    //         background-color: #f8f8f8;
    //         border-radius: 8px;

    //         .value {
    //             display: inline-block;
    //             min-width: 80px;
    //             text-align: right;
    //             font-weight: 500;
    //             font-size: 24px;
    //             color: #40485f;
    //         }
    //     }
    //     img {
    //         position: relative;
    //         top: 1px;
    //         margin-right: 8px;
    //     }
    // }
    .top {
        .top-content {
            display: flex;
            align-items: center;
            flex-wrap: wrap;
            .flex {
                display: flex;
                align-items: center;
            }
            ::v-deep .el-avatar {
                margin-right: 8px;
                height: 16px;
                width: 16px;
                line-height: 16px;
            }
            .icon-type {
                display: inline-block;
                width: 16px;
                height: 16px;
                line-height: 16px;
                text-align: center;
                margin-right: 8px;
                font-size: 12px;
                background-size: 100% 100%;
            }
            .icon-type.private {
                background-image: url("~@/assets/image/common/icon_private.svg");
                // color: #7b5b12;
                // background: linear-gradient(91deg, #ffd980 0%, #fdf5d8 100%);
            }
            .icon-type.public {
                background-image: url("~@/assets/image/common/icon_public.svg");
                // color: #394c84;
                // background: linear-gradient(270deg, #f0f4f8 0%, #ccdbf7 100%);
            }
            .tmpl-name {
                display: inline-block;
                font-size: 20px;
                font-weight: 500;
                margin-right: 24px;
                height: 26px;
                line-height: 23px;
            }
            .space-admin {
                display: flex;
                align-items: start;
                margin-top: 2px;
                .label {
                    display: inline-block;
                    width: 100px;
                    margin-right: 16px;
                    color: #40485f;
                    img {
                        margin-right: 4px;
                    }
                }
                .value {
                    display: inline-block;
                    flex: 1;
                    color: #171e31;
                }
                .value-item {
                    display: inline-block;
                    margin-right: 8px;
                }
            }
        }
        .desc {
            color: #a6abbc;
            font-size: 12px;
            margin-top: 4px;
        }
    }
    .content-img {
        margin: auto 0;
        text-align: center;
        .sub-box {
            margin: 180px 0;
        }
        img {
            height: 180px;
            width: 180px;
        }
        .no-data-text {
            color: #a6abbc;
            font-size: 14px;
        }
    }
}
</style>
