<template>
    <div class="my-homepage">
        <div v-if="recentVisitedList.length">
            <div class="sub-title bottom36">最近访问</div>
            <div class="flex">
                <div
                    class="progress-box"
                    v-for="(item, index) in recentVisitedList"
                    :key="index"
                    @click="enterToProgress(item)"
                >
                    <div class="content">
                        <b class="icon-type"></b>
                        <tip-one :text="item.tmpl_name" position="top">
                            <div class="name">{{ item.tmpl_name || "--" }}</div>
                        </tip-one>

                        <div class="space-name">
                            所属空间：
                            <tip-one :text="item.ws_name" position="top">
                                <span>{{ item.ws_name || "--" }}</span>
                            </tip-one>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="sub-title title-flex">
            <span>我的空间</span>
            <el-button
                class="basic-ui"
                type="primary"
                size="small"
                @click="newSpace"
                >+新增空间</el-button
            >
        </div>
        <div class="flex" v-if="spaceList.length">
            <div
                class="space-box"
                v-for="(item, index) in spaceList"
                :key="index"
            >
                <div class="space-item" @click="jumpToCurSpace(item)">
                    <div class="item-top" :style="bgImgs(index)"></div>
                    <div class="item-bottom">
                        <div class="bottom-box">
                            <span :style="titleStyle(index)">{{
                                firstLetter(item.name)
                            }}</span>
                        </div>
                        <div class="bottom-right">
                            <tip-one :text="item.name">
                                <div class="name">{{ item.name || "--" }}</div>
                            </tip-one>
                            <tip-more :text="item.desc">
                                <div class="desc">{{ item.desc || "--" }}</div>
                            </tip-more>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div
            v-else
            class="no-space"
            v-loading="loading"
            element-loading-text="加载中"
            element-loading-background="rgba(255, 255, 255)"
        >
            <div v-show="imgShow">
                <img
                    :src="require(`@/assets/image/space/no_space_icon.png`)"
                    alt=""
                />
                <div class="text">
                    暂无工作空间<br />请创建您的第一个工作空间
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
import api_common from "@/common/api";
import AddSpaceDialog from "./new_space_dialog.vue";
import TipMore from "@/pages/common/tooltip_more_line.vue";
import TipOne from "@/pages/common/tooltip_one_line.vue";
import { TITLE_COLOT_ARR, SPACE_BG_ARR } from "@/assets/tool/color_list";
export default {
    components: {
        AddSpaceDialog,
        TipMore,
        TipOne
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
                return item.tmpl_mode === "private" ? "私" : "公";
            };
        },
        firstLetter() {
            return (name) => {
                return name ? name.substring(0, 1) : "";
            };
        },
        titleStyle() {
            return (index) => {
                let colorItem = this.colorList[index % this.colorList.length];
                return {
                    color: colorItem.color,
                    backgroundColor: colorItem.backgroundColor
                    // boxShadow: `inset 0px 3px 6px 1px ${this.hexToRgba(
                    //     colorItem.color,
                    //     0.1
                    // )}`,
                };
            };
        },
        bgImgs() {
            return (index) => {
                let colorItem =
                    this.spaceBgList[index % this.spaceBgList.length];
                return {
                    background: `url(${colorItem.bgImg})`,
                    backgroundRepeat: "no-repeat",
                    backgroundSize: "100%"
                };
            };
        }
    },
    data() {
        return {
            isJump: false,
            colorList: TITLE_COLOT_ARR,
            spaceBgList: SPACE_BG_ARR,
            recentVisitedList: [],
            spaceList: [],
            spaceList2: [{}],
            loading: true,
            imgShow: false
        };
    },
    watch: {
        workSpaceList: {
            handler() {
                if (this.workSpaceList && this.workSpaceList.length) {
                    this.spaceList = [...this.workSpaceList];
                } else {
                    this.spaceList = [];
                }
                this.loading = false;
                setTimeout(() => {
                    this.imgShow = true;
                }, 50);
            },
            deep: true
        }
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
        // this.isFirstEnter();
        // this.$nextTick(() => {
        //     if (this.curSpace && Object.keys(this.curSpace).length) {
        //         this.$router.push({
        //             path: `/workspace/${this.curSpace.id}`,
        //         })
        //     }
        // })
        this.getRecentItem();
        this.fetchSpaceList();
    },
    beforeDestroy() {
        window.name = "";
    },
    methods: {
        fetchSpaceList() {
            this.$store.dispatch("fetchSpaceList").then((res) => {
                let localSpace = localStorage.getItem("space");
                if (
                    !localSpace ||
                    !localSpace.length ||
                    localSpace === "undefined"
                ) {
                    this.$store.commit("setCurSpace", res[0]);
                } else {
                    let localId = JSON.parse(localStorage.getItem("space")).id;
                    if (res.some((item) => localId == item.id)) {
                        res.forEach((item) => {
                            if (localId == item.id) {
                                this.$store.commit("setCurSpace", item);
                            }
                        });
                    } else {
                        this.$store.commit("setCurSpace", res[0]);
                    }
                }
            });
        },
        enterToProgress(item) {
            this.$store.commit("setCurSpace", {
                id: item.ws_id,
                name: item.ws_name
            });
            if (item.tmpl_id) {
                this.$router.push({
                    path: `/progress/${item.ws_id}/${item.tmpl_id}`
                });
            }
        },
        getRecentItem() {
            api_common.getRecentItem().then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    res.data.length
                ) {
                    this.recentVisitedList = res.data.slice(0, 20);
                } else {
                    this.recentVisitedList = [];
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
        },
        // isFirstEnter() {
        //     if (window.name === '') {
        //         // 页面首次被加载
        //         if (this.workSpaceList && this.workSpaceList.length) {
        //             this.spaceList = [...this.workSpaceList];
        //         } else {
        //             this.spaceList = [];
        //         }
        //         window.name = 'home_page';
        //     }
        // },
        newSpace() {
            this.$refs.addSpaceDialog.openDialog();
        },
        jumpToCurSpace(item) {
            this.$store.commit("setCurSpace", item);
            // this.isJump = true
            this.$router.push({
                path: `/workspace/${item.id}`
            });
        },
        createSpaceConfirm(form) {
            this.curForm = _.cloneDeep(form);
            let params = {
                ...form
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
                                        path: `/workspace/${item.id}`
                                    });
                                }
                            });
                        }
                    });
                }
            });
        },
        enterSpace() {}
    }
};
</script>

<style lang="scss" scoped>
.my-homepage {
    font-family: MiSans, MiSans;

    .bottom36 {
        margin-bottom: 36px;
    }
    .flex {
        display: flex;
        flex-wrap: wrap;
    }
    .title-flex {
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
    .progress-box {
        position: relative;
        box-sizing: border-box;
        padding: 36px 24px 24px;
        margin: 0 20px 24px 0;
        background-color: #f8f8f8;
        border-radius: 8px;
        cursor: pointer;
        &:hover {
            background-color: #e9f0f8;
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
            background-image: url("~@/assets/image/common/progress-icon.png");
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

    .space-box {
        position: relative;
        border-radius: 8px;
        margin: 0 20px 24px 0;
        padding: 0 0 16px 0;
        box-sizing: border-box;
        background: #ffffff;
        box-shadow: 0px 4px 10px 1px rgba(0, 0, 0, 0.1);
        cursor: pointer;
        &:hover {
            box-shadow: 0px 4px 10px 1px rgba(0, 0, 0, 0.3);
        }
        .space-item {
            .item-top {
                height: 120px;
                border-radius: 8px 8px 0 0;
                background-repeat: no-repeat;
                background-size: 100%;
            }
            .item-bottom {
                display: flex;
                padding: 20px 16px 0;
                height: 74px; // 描述3行，80px；
                .bottom-box span {
                    display: inline-block;
                    width: 40px;
                    height: 40px;
                    line-height: 40px;
                    text-align: center;
                    margin-right: 10px;
                    font-size: 18px;
                    font-weight: 500;
                    border-radius: 8px;
                }
                .bottom-right {
                    width: calc(100% - 40px - 12px);
                    .name {
                        width: 100%;
                        display: block;
                        margin-bottom: 10px;
                        font-size: 16px;
                        font-weight: 600;
                        overflow: hidden;
                        text-overflow: ellipsis;
                        white-space: nowrap;
                    }

                    .desc {
                        font-size: 12px;
                        font-weight: 400;
                        color: #a6abbc;
                        line-height: 18px;
                        display: -webkit-box; //将对象作为弹性伸缩盒子模型显示。
                        -webkit-box-orient: vertical; //从上到下垂直排列子元素（设置伸缩盒子的子元素排列方式）
                        -webkit-line-clamp: 2; //限制行数
                        overflow: hidden; //超出部分隐藏
                        text-overflow: ellipsis; //用一个省略号代替超出的内容
                    }
                }
            }
        }
    }
    .no-space {
        box-sizing: border-box;
        text-align: center;
        height: 570px;
        padding: 150px 0 100px;
        img {
            width: 180px;
            height: 180px;
        }
        .text {
            color: #a6abbc;
            font-size: 14px;
            line-height: 24px;
            margin-top: -20px;
        }
    }
    @media screen and (max-width: 700px) {
        .progress-box,
        .space-box {
            width: calc(100% - 20px);
        }
    }
    @media screen and (min-width: 700px) and (max-width: 999px) {
        .progress-box {
            width: calc((100% - 20px) / 2);
            &:nth-child(2n) {
                margin-right: 0;
            }
        }
        .space-box {
            width: calc((100% - 20px) / 2);
            &:nth-child(2n) {
                margin-right: 0;
            }
        }
    }
    @media screen and (min-width: 1000px) and (max-width: 1399px) {
        .progress-box {
            width: calc((100% - 40px) / 3);
            &:nth-child(3n) {
                margin-right: 0;
            }
        }
        .space-box {
            width: calc((100% - 40px) / 3);
            &:nth-child(3n) {
                margin-right: 0;
            }
        }
    }
    @media screen and (min-width: 1400px) and (max-width: 1899px) {
        .progress-box {
            width: calc((100% - 60px) / 4);
            &:nth-child(4n) {
                margin-right: 0;
            }
        }
        .space-box {
            width: calc((100% - 60px) / 4);
            &:nth-child(4n) {
                margin-right: 0;
            }
        }
    }
    @media screen and (min-width: 1900px) {
        .progress-box {
            width: calc((100% - 80px) / 5);
        }
        .progress-box:nth-child(5n) {
            margin-right: 0;
        }
        .space-box {
            width: calc((100% - 80px) / 5);
        }
        .space-box:nth-child(5n) {
            margin-right: 0;
        }
    }
}
</style>
