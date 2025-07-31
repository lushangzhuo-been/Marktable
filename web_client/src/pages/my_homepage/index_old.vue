<template>
    <div>
        <div v-if="lastVisitedList.length">
            <div class="sub-title bottom36">最近访问</div>
            <div class="flex">
                <div
                    class="progress-box"
                    v-for="(item, index) in lastVisitedList"
                    :key="index"
                >
                    <div class="content">
                        <!-- <div :class="['icon-type', item.mode]">
                            {{ progressMode(item) }}
                        </div> -->
                        <b
                            :class="['icon-type', item.mode]"
                        ></b>
                        <tip-one :text="item.name" position="top">
                            <div class="name">{{ item.name }}</div>
                        </tip-one>

                        <div class="space-name">
                            所属空间：
                            <tip-one :text="item.spaceName" position="top">
                                <span>{{ item.spaceName }}</span>
                            </tip-one>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="sub-title">
            <img
                :src="require(`@/assets/image/common/index_homepage.png`)"
                alt=""
                width="20px"
                height="20px"
                class="icon"
            />
            我的空间
        </div>
        <div class="flex">
            <div
                :class="['box', index === 0 ? '' : 'bg-img']"
                v-for="(item, index) in spaceList"
                :key="index"
            >
                <div v-if="index === 0" class="add-space" @click="newSpace">
                    <img
                        :src="require('@/assets/image/space/add_space.png')"
                        alt=""
                        :width="32"
                        :height="32"
                    />
                    <div class="name">{{ item.name }}</div>
                </div>
                <div v-else class="space-item">
                    <!-- <div class="top-bar"></div> -->
                    <div class="item-top">
                        <div class="top-left">
                            <span :style="titleStyle(index)">{{
                                firstLetter(item.name)
                            }}</span>
                        </div>
                        <div class="top-right">
                            <tip-one :text="item.name">
                                <div class="name">{{ item.name || "--" }}</div>
                            </tip-one>
                            <tip-more :text="item.desc">
                                <div class="desc">{{ item.desc || "--" }}</div>
                            </tip-more>
                        </div>
                    </div>
                    <div class="item-bottom">
                        <div class="bot-left">
                            <div class="user">
                                创建人:
                                <tip-one :text="item.created_at">
                                    <span class="created-user">{{
                                        item.creator_name || "--"
                                    }}</span>
                                </tip-one>
                            </div>
                            <div class="date">
                                创建时间:
                                <tip-one :text="item.created_at">
                                    <span class="created-date">{{
                                        item.created_at || "--"
                                    }}</span>
                                </tip-one>
                            </div>
                        </div>
                        <div class="bot-right">
                            <div
                                class="btn-enter"
                                @click="jumpToCurSpace(item)"
                            >
                                进入空间
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <template>
            <add-space-dialog
                ref="addSpaceDialog"
                @on-confirm="createSpaceConfirm"
            ></add-space-dialog>
        </template>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/space";
import AddSpaceDialog from "./new_space_dialog.vue";
import TipMore from "@/pages/common/tooltip_more_line.vue";
import TipOne from "@/pages/common/tooltip_one_line.vue";
import { TITLE_COLOT_ARR } from "@/assets/tool/color_list";
export default {
    components: {
        AddSpaceDialog,
        TipMore,
        TipOne,
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        workSpaceList() {
            return this.$store.state.workSpaceList;
        },
        progressTreeList() {
            return this.$store.state.progressTree;
        },
        progressMode() {
            return (item) => {
                return item.mode === "private" ? "私" : "公";
            };
        },
        firstLetter() {
            return (name) => {
                return name ? name.substring(0, 1) : "";
            };
        },
        titleStyle() {
            return (index) => {
                let colorItem =
                    this.colorList[(index - 1) % this.colorList.length];
                return {
                    color: colorItem.color,
                    backgroundColor: colorItem.backgroundColor,
                    // boxShadow: `inset 0px 3px 6px 1px ${this.hexToRgba(
                    //     colorItem.color,
                    //     0.1
                    // )}`,
                };
            };
        },
    },
    data() {
        return {
            isJump: false,
            colorList: TITLE_COLOT_ARR,
            lastVisitedList: [
                {
                    name: "测试测试测试测试测试",
                    spaceName: "数字化平台",
                    date: "2024-02-15",
                    mode: "private",
                    user: "梁杏",
                },
                {
                    name: "testtesttesttesttesttest",
                    spaceName: "混合云持续运营",
                    date: "2024-02-15",
                    mode: "public",
                    user: "梁杏",
                },
                {
                    name: "四月十六日空间测试",
                    spaceName:
                        "设计组得空间设计组得空间设计组得空间设计组得空间",
                    date: "2024-02-15",
                    mode: "private",
                    user: "梁杏",
                },
                {
                    name: "四月十六日空间测试四月十六日空间测试",
                    spaceName: "设计组得空间",
                    date: "2024-02-15",
                    mode: "private",
                    user: "梁杏",
                },
                {
                    name: "测试",
                    spaceName: "设计组得空间",
                    date: "2024-02-15",
                    mode: "public",
                    user: "梁杏",
                },
            ],
            spaceList: [
                {
                    name: "新建空间",
                },
            ],
            spaceList2: [{}],
        };
    },
    watch: {
        workSpaceList: {
            handler() {
                if (this.workSpaceList && this.workSpaceList.length) {
                    this.spaceList = [
                        {
                            name: "新建空间",
                        },
                        ...this.workSpaceList,
                    ];
                } else {
                    this.spaceList = [
                        {
                            name: "新建空间",
                        },
                    ];
                }
            },
            deep: true,
        },
        // progressTreeList(arr) {
        //     // if (this.isJump) {
        //     if (arr && arr.length) {
        //         // this.$router.push({
        //         //     path: `/progress/${arr[0].id}`,
        //         // })
        //         this.$router.push({
        //             path: `/workspace/${arr[0].id}`,
        //         })
        //     }
        //     this.isJump = false
        //     // }
        // },
        // curSpace(obj) {
        //     this.$router.push({
        //         path: `/workspace/${obj.id}`,
        //     })
        // },
    },
    mounted() {
        this.isFirstEnter();
        // this.$nextTick(() => {
        //     if (this.curSpace && Object.keys(this.curSpace).length) {
        //         this.$router.push({
        //             path: `/workspace/${this.curSpace.id}`,
        //         })
        //     }
        // })
    },
    beforeDestroy() {
        window.name = "";
    },
    methods: {
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
        },
        isFirstEnter() {
            if (window.name === "") {
                // 页面首次被加载
                if (this.workSpaceList && this.workSpaceList.length) {
                    this.spaceList = [
                        {
                            name: "新建空间",
                        },
                        ...this.workSpaceList,
                    ];
                } else {
                    this.spaceList = [
                        {
                            name: "新建空间",
                        },
                    ];
                }
                window.name = "home_page";
            }
        },
        newSpace() {
            this.$refs.addSpaceDialog.openDialog();
        },
        jumpToCurSpace(item) {
            this.$store.commit("setCurSpace", item);
            // this.isJump = true
            this.$router.push({
                path: `/workspace/${item.id}`,
            });
        },
        createSpaceConfirm(form) {
            this.curForm = _.cloneDeep(form);
            let params = {
                ...form,
            };
            api.creatWorkspace(params).then((response) => {
                if (response.resultCode === 200) {
                    this.$refs.addSpaceDialog.cancel();
                    this.$store.dispatch("fetchSpaceList").then((res) => {
                        if (res && res.length) {
                            res.forEach((item) => {
                                if (item.name === this.curForm.name) {
                                    this.$store.commit("setCurSpace", item);
                                    this.$router.push({
                                        path: `/workspace/${item.id}`,
                                    });
                                }
                            });
                        }
                    });
                }
            });
        },
        enterSpace() {},
    },
};
</script>

<style lang="scss" scoped>
.bottom36 {
    margin-bottom: 36px;
}
.flex {
    display: flex;
    flex-wrap: wrap;
}
.progress-box {
    width: 270px;
    position: relative;
    box-sizing: border-box;
    padding: 36px 24px 24px;
    margin: 0 20px 24px 0;
    background-color: #f8f8f8;
    border-radius: 8px;
    cursor: pointer;
    &:hover {
        box-shadow: 2px 2px 8px 1px rgba(47, 56, 76, 0.1);
    }
    .icon-type {
        position: absolute;
        top: -16px;
        width: 40px;
        height: 40px;
        line-height: 40px;
        text-align: center;
        border-radius: 4px;
        background-size: 100% 100%;
    }
    .icon-type.private {
        background-image: url("~@/assets/image/common/icon_private.svg");
        //         color: #7b5b12;
        // background: linear-gradient(91deg, #ffd980 0%, #fdf5d8 100%);
    }
    .icon-type.public {
        background-image: url("~@/assets/image/common/icon_public.svg");
        // color: #394c84;
        // background: linear-gradient(270deg, #f0f4f8 0%, #ccdbf7 100%);
    }
    .name {
        font-weight: 700;
        margin-bottom: 8px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }
    .space-name {
        display: flex;
        align-items: center;
        font-size: 12px;
        color: #82889e;
        span {
            display: inline-block;
            width: calc(100% - 64px);
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
        }
    }
}
.bg-img {
    background: url("~@/assets/image/common/home_page_space_bg.png") no-repeat;
    background-size: 100% 100%;
}
.box {
    position: relative;
    width: 270px;
    border: 1px solid #e6e9f0;
    border-radius: 8px;
    margin: 0 20px 24px 0;
    padding: 16px;
    box-sizing: border-box;
    cursor: pointer;
    &:hover {
        box-shadow: 2px 2px 8px 1px rgba(47, 56, 76, 0.1);
    }
    .add-space {
        position: relative;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        text-align: center;
        .name {
            margin-top: 16px;
            color: #2f384c;
        }
    }
    .space-item {
        margin-top: 8px;
        .item-top {
            display: flex;
            height: 76px;
            padding-bottom: 12px;
            border-bottom: 1px solid #cdd5e6;
        }
        .top-left span {
            display: inline-block;
            width: 40px;
            height: 40px;
            line-height: 40px;
            text-align: center;
            margin-right: 10px;
            font-size: 18px;
            font-weight: 500;
            border: 2px solid #fff;
            // box-shadow: inset 0px 3px 6px 1px rgba(25, 124, 240, 0.1);
            box-shadow: 0px 3px 10px 1px rgba(0, 0, 0, 0.1);
            border-radius: 8px;
        }
        .top-right {
            width: calc(100% - 40px - 12px);
            .name {
                width: 100%;
                display: block;
                margin-bottom: 6px;
                font-size: 16px;
                font-weight: 600;
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
            }

            .desc {
                margin-bottom: 10px;
                font-size: 12px;
                font-weight: 400;
                color: #a6abbc;
                display: -webkit-box; //将对象作为弹性伸缩盒子模型显示。
                -webkit-box-orient: vertical; //从上到下垂直排列子元素（设置伸缩盒子的子元素排列方式）
                -webkit-line-clamp: 3; //限制行数
                overflow: hidden; //超出部分隐藏
                text-overflow: ellipsis; //用一个省略号代替超出的内容
            }
        }
        .item-bottom {
            display: flex;
            margin-top: 12px;
            .bot-left {
                width: 160px;
                color: #82889e;
                font-size: 12px;
                .user {
                    display: flex;
                    align-items: center;
                    margin-bottom: 10px;
                    .created-user {
                        margin-left: 8px;
                        display: inline-block;
                        width: calc(100% - 50px);
                        overflow: hidden;
                        text-overflow: ellipsis;
                        white-space: nowrap;
                    }
                }
                .date {
                    display: flex;
                    align-items: center;
                    .created-date {
                        margin-left: 8px;
                        display: inline-block;
                        width: calc(100% - 62px);
                        overflow: hidden;
                        text-overflow: ellipsis;
                        white-space: nowrap;
                    }
                }
            }
            .bot-right {
                position: relative;
                width: calc(100% - 160px);
                .btn-enter {
                    position: absolute;
                    top: 10px;
                    right: 0;
                    width: 70px;
                    height: 24px;
                    line-height: 24px;
                    text-align: center;
                    font-size: 12px;
                    color: #1890ff;
                    border-radius: 17px;
                    border: 1px solid #1890ff;
                }
            }
        }
    }
}
</style>
