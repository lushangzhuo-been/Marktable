<template>
    <div class="no-data">
        <div class="top-img" :style="bgImgs"></div>
        <div class="box">
            <div class="left-box" :style="titleStyle">
                {{ firstLetter(curWorkSpaceInfo.name) }}
            </div>
            <div class="right-box">
                <tip-one :text="curWorkSpaceInfo.name">
                    <div class="space-name">
                        {{ curWorkSpaceInfo.name || "--" }}
                    </div>
                </tip-one>
                <tip-more :text="curWorkSpaceInfo.desc">
                    <div class="space-desc">
                        {{ curWorkSpaceInfo.desc || "--" }}
                    </div>
                </tip-more>
                <div class="space-admin">
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
                        <template v-if="admins && admins.length">
                            <span
                                v-for="(item, index) in admins"
                                :key="index"
                                class="value-item"
                            >
                                <span class="flex">
                                    <user-message :userMessage="item">
                                        <span>
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
                                    <span v-if="index < admins.length - 1"
                                        >,</span
                                    >
                                </span>
                            </span>
                        </template>
                        <template v-else>--</template>
                    </span>
                </div>
                <div>
                    <div class="space-num">
                        <img
                            :src="
                                require('@/assets/image/common/progress_setting.png')
                            "
                            :width="16"
                            :height="16"
                        />
                        <span class="label">流程:</span>
                        <span class="value">{{ progressNum || 0 }}个</span>
                    </div>
                    <div class="space-num">
                        <img
                            :src="
                                require('@/assets/image/field_type_icon/people_multiple.svg')
                            "
                            :width="16"
                            :height="16"
                        />
                        <span class="label">成员:</span>
                        <span class="value">{{ usersNum || 0 }}个</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { imgHost } from "@/assets/tool/const";
import api from "@/common/api/module/space";
import TipMore from "@/pages/common/tooltip_more_line.vue";
import TipOne from "@/pages/common/tooltip_one_line.vue";
import { TITLE_COLOT_ARR, SPACE_BG_ARR } from "@/assets/tool/color_list";
import UserMessage from "@/components/user_message_tip";
export default {
    components: {
        TipMore,
        TipOne,
        UserMessage
    },
    computed: {
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
        workSpaceList() {
            return this.$store.state.workSpaceList;
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        workspaceHomePageInfo() {
            return this.$store.state.workspaceHomePageInfo || {};
        },
        bgImgs() {
            let obj = {};
            if (this.workSpaceList && this.workSpaceList.length) {
                this.workSpaceList.forEach((item, index) => {
                    if (item.name === this.curSpace.name) {
                        let colorItem =
                            this.spaceBgList[index % this.spaceBgList.length];
                        obj = {
                            background: `url(${colorItem.bgImg})`,
                            backgroundRepeat: "no-repeat",
                            backgroundSize: "100%"
                        };
                    }
                });
            } else {
                obj = {
                    background: "",
                    backgroundRepeat: "no-repeat",
                    backgroundSize: "100%"
                };
            }
            return obj;
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
        },
        progressTree() {
            return this.$store.state.progressTree;
        }
    },
    data() {
        return {
            colorList: TITLE_COLOT_ARR,
            spaceBgList: SPACE_BG_ARR,
            admin: [
                {
                    name: "发发发",
                    img: ""
                },
                {
                    name: "祥哥",
                    img: ""
                },
                {
                    name: "成辉成",
                    img: ""
                },
                {
                    name: "鹏飞鹏飞",
                    img: ""
                },
                {
                    name: "发发",
                    img: ""
                },
                {
                    name: "梁杏杏",
                    img: ""
                },
                {
                    name: "发发",
                    img: ""
                },
                {
                    name: "梁杏杏",
                    img: ""
                },
                {
                    name: "发发",
                    img: ""
                }
            ],
            curWorkSpaceInfo: {
                name: "",
                desc: ""
            },
            admins: [],
            progressNum: 0,
            usersNum: 0
        };
    },
    watch: {
        workspaceHomePageInfo: {
            handler(obj) {
                this.setInfo(obj);
            },
            deep: true,
            immediate: true
        },
        progressTree: {
            handler() {
                this.fetchCurWorkspaceInfo();
            },
            deep: true
        }
    },
    methods: {
        setInfo(data) {
            if (data && Object.keys(data).length) {
                this.curWorkSpaceInfo = data.ws || {
                    name: "",
                    desc: ""
                };
                this.admins = data.admin_list || [];
                this.progressNum = data.tmpl_cnt;
                this.usersNum = data.member_cnt;
            } else {
                this.curWorkSpaceInfo = {
                    name: "",
                    desc: ""
                };
                this.admins = [];
                this.progressNum = 0;
                this.usersNum = 0;
            }
        },
        fetchCurWorkspaceInfo() {
            api.getCurWorkspaceInfo({ ws_id: this.curSpace.id }).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    this.curWorkSpaceInfo = res.data.ws || {
                        name: "",
                        desc: ""
                    };
                    this.admins = res.data.admin_list || [];
                    this.progressNum = res.data.tmpl_cnt;
                    this.usersNum = res.data.member_cnt;
                } else {
                    this.curWorkSpaceInfo = {
                        name: "",
                        desc: ""
                    };
                    this.admins = [];
                    this.progressNum = 0;
                    this.usersNum = 0;
                }
            });
        },
        hexToRgba(hex, opacity) {
            return (
                "rgba(" +
                parseInt("0x" + hex.slice(1, 3)) +
                "," +
                parseInt("0x" + hex.slice(3, 5)) +
                "," +
                parseInt("0x" + hex.slice(5, 7)) +
                "," +
                opacity +
                ")"
            );
        }
    }
};
</script>

<style lang="scss" scoped>
.no-data {
    position: relative;
    height: 100%;
    width: 100%;

    ::v-deep .el-avatar {
        margin-right: 8px;
        height: 16px;
        width: 16px;
        line-height: 16px;
    }

    .top-img {
        height: 220px;
        background: url("@/assets/image/space/homepage.png");
        background-repeat: no-repeat;
        background-size: 100%;
    }
    ::v-deep .el-avatar {
        position: relative;
        top: 4px;
    }
    .box {
        display: flex;
        margin: 16px 10% 0 10%;
        .left-box {
            height: 118px;
            line-height: 118px;
            width: 118px;
            margin-right: 24px;
            text-align: center;
            font-size: 56px;
            // border: 3px solid #fff;
            border-radius: 8px;
            background-color: #e5ebfd;
            color: #1890ff;
            // box-shadow: 0px 3px 10px 1px rgba(0, 0, 0, 0.1);
        }
        .right-box {
            width: calc(100% - 115px);
        }
        .space-name {
            font-size: 24px;
            font-weight: 600;
            margin-bottom: 16px;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
        }
        .space-desc {
            line-height: 20px;
            font-size: 12px;
            color: #a6abbc;
            margin-bottom: 24px;
            display: -webkit-box; //将对象作为弹性伸缩盒子模型显示。
            -webkit-box-orient: vertical; //从上到下垂直排列子元素（设置伸缩盒子的子元素排列方式）
            -webkit-line-clamp: 2; //限制行数
            overflow: hidden; //超出部分隐藏
            text-overflow: ellipsis; //用一个省略号代替超出的内容
        }
        .space-admin {
            display: flex;
            align-items: start;
            padding-top: 24px;
            margin-bottom: 6px;
            border-top: 1px solid #ccc;
            .label {
                display: inline-block;
                width: 100px;
                margin: 0 16px 10px 0;
                color: #40485f;
            }
            .value {
                display: inline-block;
                width: calc(100% - 116px);
                color: #171e31;
            }
            .value-item {
                display: inline-block;
                margin: 0 12px 10px 0;
            }
        }
        .space-num {
            display: flex;
            align-items: center;
            margin-bottom: 16px;
            .label {
                margin-right: 16px;
                color: #40485f;
            }
            .value {
                color: #171e31;
            }
        }
        img {
            position: relative;
            top: 1px;
            margin-right: 8px;
        }
    }
    .flex {
        display: flex;
        align-items: center;
    }
}
</style>
