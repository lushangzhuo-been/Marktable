<template>
    <div class="left-menu">
        <div class="content">
            <div class="logo-title" @click="enterHomePage">
                <!-- <img
                    :src="require('@/assets/image/common/a_logo_2.png')"
                    alt=""
                    :width="32"
                    :height="32"
                />
                项目管理 
                -->
                <!-- <img
                    :src="require('@/assets/image/common/a_logo.png')"
                    alt=""
                    :width="32"
                    :height="32"
                />
                Marktable -->
                <b class="logo-name"></b>
            </div>
            <!-- 我的首页、代办、通知 -->
            <div class="top" ref="top">
                <div
                    v-for="(item, index) in topList"
                    :key="index"
                    class="item"
                    :class="item.active ? 'active' : ''"
                    @click="jumpTo(item.name, 'top')"
                >
                    <img :src="item.img" alt="" :width="18" :height="18" />
                    <span class="name"
                        >{{ item.name }}
                        <span
                            class="message-unread-num"
                            v-if="item.name === '我的消息' && unReadNum"
                            >{{ unReadNum > 99 ? "99+" : unReadNum }}</span
                        >
                    </span>
                </div>
            </div>
            <!-- 空间 -->
            <space ref="space"></space>
            <!-- 流程树 -->
            <progress-tree
                :style="{ height: progressTreeHeight }"
            ></progress-tree>
        </div>
        <!-- 个人中心 -->
        <div class="bottom">
            <el-avatar
                icon="el-icon-user-solid"
                size="small"
                :src="getAvatar(userInfo.avatar)"
            ></el-avatar>
            <el-dropdown
                placement="right"
                trigger="click"
                width="200"
                @command="(val) => jumpTo(val, 'personal')"
            >
                <el-dropdown-menu slot="dropdown" class="personal-center">
                    <el-dropdown-item
                        :command="`${item.name}`"
                        v-for="(item, index) in personalCenterList"
                        :key="index"
                        class="personal-center-item"
                        :class="item.active ? 'active' : ''"
                    >
                        <img :src="item.img" alt="" class="img-more" />
                        <span class="name">{{ item.name }}</span>
                    </el-dropdown-item>
                </el-dropdown-menu>
                <span class="font">
                    <tip-one :text="userInfo.full_name">
                        <span
                            :style="{
                                color: isPersonalHighLight
                                    ? PRIMARY_COLOR
                                    : '#40485F'
                            }"
                            >{{ userInfo.full_name }}
                        </span>
                    </tip-one>
                    <img
                        :src="
                            require(`@/assets/image/common/${
                                isPersonalHighLight
                                    ? 'right_arraw_active'
                                    : 'right_arraw_defalut'
                            }.png`)
                        "
                        alt=""
                    />
                </span>
            </el-dropdown>
        </div>
    </div>
</template>

<script>
import api from "@/common/api/module/space";
import apiProgress from "@/common/api/module/progress";
import { imgHost } from "@/assets/tool/const";
import Space from "./components/space.vue";
import ProgressTree from "./components/progress_tree.vue";
import TipOne from "@/pages/common/tooltip_one_line.vue";
import { PRIMARY_COLOR_LIST } from "@/assets/tool/color_list";
export default {
    components: {
        Space,
        TipOne,
        ProgressTree
    },
    computed: {
        userInfo() {
            return this.$store.getters.userInfo;
        },
        // 个人中心高亮
        isPersonalHighLight() {
            if (this.$route.path.indexOf("/personal/center") > -1) {
                return true;
            } else {
                return false;
            }
        },
        getAvatar() {
            return (avatar) => {
                if (avatar) {
                    return `${imgHost}${avatar}`;
                }
                return require(`@/assets/image/common/default_avatar_big.png`);
            };
        }
    },
    data() {
        return {
            progressTreeHeight: "100px",
            PRIMARY_COLOR: PRIMARY_COLOR_LIST.PRIMARY_COLOR,
            topList: [
                {
                    name: "我的首页",
                    components: "myHomePage",
                    path: "/home",
                    status: true,
                    active: false,
                    img: require(`@/assets/image/common/index_homepage.png`)
                },
                // {
                //     name: '我的代办',
                //     status: true,
                //     components: 'myTodo',
                //     path: '/todo',
                //     active: false,
                //     img: require(`@/assets/image/common/index_my_todo.png`),
                // },
                {
                    name: "我的消息",
                    status: true,
                    components: "messageNotice",
                    path: "/notice",
                    active: false,
                    img: require(`@/assets/image/common/index_message_notice.png`)
                }
            ],
            personalCenterList: [
                {
                    name: "基本信息",
                    status: true,
                    components: "basicInfo",
                    path: "/personal/center/basic/info",
                    img: require(`@/assets/image/field_type_icon/people_single.svg`),
                    active: false
                },
                {
                    name: "修改密码",
                    status: true,
                    components: "modifyPwd",
                    path: "/personal/center/modify/pwd",
                    img: require(`@/assets/image/common/lock.png`),
                    active: false
                },
                {
                    name: "重置邮箱",
                    status: true,
                    components: "updateEmail",
                    path: "personal/center/update/email",
                    img: require(`@/assets/image/common/email.png`),
                    active: false
                },
                {
                    name: "退出登录",
                    status: true,
                    components: "updateEmail",
                    path: "/personal/center/update/email",
                    img: require(`@/assets/image/common/logout.png`),
                    active: false
                }
            ],
            unReadNum: 0
        };
    },
    watch: {
        $route(to, from) {
            // 排除顶部菜单和个人
            if (
                to.name.indexOf("my") === -1 &&
                to.name.indexOf("messageNotice") === -1 &&
                to.path.indexOf("personal/center") === -1
            ) {
                this.initData();
            } else {
                this.getCurrentMenu();
            }
        }
    },
    mounted() {
        this.getCurrentMenu();
        this.$nextTick(() => {
            let topHeight = this.$refs.top.offsetHeight;
            this.progressTreeHeight = `calc(100% - ${topHeight}px - 66px - 96px)`;
        });
        this.getUnReadNum();
    },

    methods: {
        // apiProgress
        getUnReadNum() {
            apiProgress.getUnReadNum().then((res) => {
                if (res && res.resultCode === 200) {
                    this.unReadNum = res.data;
                } else {
                    this.unReadNum = 0;
                }
            });
        },
        initData() {
            this.topList.forEach((item) => {
                this.$set(item, "active", false);
            });
            this.personalCenterList.forEach((item) => {
                this.$set(item, "active", false);
            });
        },
        // 点击logo或者汉字进入首页
        enterHomePage() {
            if (this.$route.path === "/home") return;
            this.$router.push({
                name: "myHomePage"
            });
        },
        // setStyle(arr, tmpObj) {
        //     arr.forEach((item) => {
        //         if (this.$route.path === item.path) {
        //             tmpObj = item
        //         }
        //         this.$set(item, 'active', this.$route.path === item.path)
        //     })
        // },
        getCurrentMenu() {
            this.$nextTick(() => {
                this.topList.forEach((item) => {
                    this.$set(item, "active", this.$route.path === item.path);
                });
                this.personalCenterList.forEach((item) => {
                    this.$set(item, "active", this.$route.path === item.path);
                });
            });
        },
        // 路由跳转
        jumpTo(curName, type) {
            if (curName === "退出登录") {
                this.$store.commit("delToken");
                this.$router.push({ path: "/login" });
                return;
            }
            if (curName === "删除") {
                return;
            }

            let mapping = {
                top: this.topList,
                personal: this.personalCenterList
            };
            if (type === "top") {
                this.personalCenterList.forEach((sub) => {
                    this.$set(sub, "active", false);
                });
            } else {
                this.topList.forEach((sub) => {
                    this.$set(sub, "active", false);
                });
            }
            let tmpObj = {};
            mapping[type] &&
                mapping[type].forEach((sub) => {
                    if (sub.name === curName) {
                        tmpObj = sub;
                    }
                    this.$set(sub, "active", sub.name === curName);
                });
            if (this.$route.path === tmpObj.path) return;
            this.$router.push({
                name: tmpObj.components
            });
        }
    }
};
</script>
<style lang="scss">
ul.personal-center.el-dropdown-menu.el-popper {
    padding: 8px;
    .el-dropdown-menu__item,
    .el-menu-item {
        // padding: 0 8px;
    }
}
ul.el-dropdown-menu.el-popper {
    box-sizing: border-box;
    padding: 8px;
    margin-top: 2px;
    .popper__arrow {
        display: none;
    }
    .el-dropdown-menu__item,
    .el-menu-item {
        padding: 0 8px;
    }
}
</style>
<style lang="scss" scoped>
.img-more {
    display: inline-block;
    width: 20px;
    height: 20px;
    margin-right: 8px;
    vertical-align: middle;
}

img {
    cursor: pointer;
    vertical-align: middle;
    margin-left: -4px;
}
.logo-title {
    display: flex;
    align-items: center;
    padding-left: 16px;
    height: 66px;
    font-family: DingTalk JinBuTi-Regular;
    font-size: 24px;
    cursor: pointer;
    img {
        margin-right: 8px;
    }
    .logo-name {
        display: inline-block;
        // width: 1066px;
        // height: 224px;
        width: 143px;
        height: 30px;
        background-image: url("@/assets/image/logo_name.png");
        background-size: 100% 100%;
    }
}
.content {
    height: calc(100% - 56px);
}
.bottom {
    height: 50px;
    display: flex;
    align-items: center;
    border-top: 1px solid #ced2de;
    margin: 0 16px;
    .el-avatar.el-avatar--icon.el-avatar--circle {
        margin-right: 8px;
    }
    .font {
        width: 100%;
        display: flex;
        align-items: center;
        cursor: pointer;
        span {
            display: inline-block;
            max-width: calc(100% - 16px - 8px);
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
        }
        img {
            position: relative;
            top: 2px;
            height: 16px;
            width: 16px;
            margin-left: 8px;
        }
    }
    .el-dropdown {
        width: calc(100% - 36px - 16px);
    }
}
.top {
    border-bottom: 1px solid #ebebea;
    padding-bottom: 10px;
    .item {
        height: 40px;
        display: flex;
        align-items: center;
        padding: 0 16px;
        border-left: solid 3px #f5f6f8;
        cursor: pointer;
        img {
            margin-right: 8px;
        }
        .name {
            height: 21px;
            line-height: 21px;
            .message-unread-num {
                box-sizing: border-box;
                display: inline-block;
                min-width: 18px;
                height: 18px;
                line-height: 18px;
                border-radius: 9px;
                padding: 0 2px;
                font-size: 12px;
                text-align: center;
                color: #ffffff;
                background-color: #ff3c3d;
            }
        }
        // .name:hover,
        // .name:focus {
        //     color: var(--PRIMARY_COLOR);
        // }
    }
    .item.disabled {
        visibility: 0.5;
        cursor: unset;
    }
    .item.active,
    .item.active:hover {
        border-left: solid 3px var(--PRIMARY_COLOR);
        background-color: var(--WHITE_COLOR);
    }
    .item:hover {
        background-color: #efefef;
    }
}

.personal-center-item {
    height: 40px;
    line-height: 40px;
    padding-left: 8px;
    cursor: pointer;
}
.personal-center-item.active {
    // color: var(--PRIMARY_COLOR);
    background-color: #f1f9ff;
}
.more {
    position: relative;
    right: -8px;
    display: inline-block;
    height: 20px;
    width: 20px;
    background-image: url("~@/assets/image/common/space_more_default.png");
    background-size: 100% 100%;
    cursor: pointer;
}

.is-select,
.more:hover {
    background-image: url("~@/assets/image/common/space_more_active.png");
    background-size: 100% 100%;
}
.absolute {
    // display: none;
    position: absolute;
    right: 8px;
}
</style>
