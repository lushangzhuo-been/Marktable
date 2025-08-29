<template>
    <div class="no-data">
        <div class="top">
            <div class="top-content">
                <span class="tmpl-name">{{
                    progressInfo.tmpl.name || "--"
                }}</span>
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
import api from "@/common/api/module/progress";
import { TITLE_COLOT_ARR } from "@/assets/tool/color_list";
import UserMessage from "@/components/user_message_tip";
export default {
    components: {
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
    .top {
        .top-content {
            display: flex;
            align-items: center;
            flex-wrap: wrap;
            .tmpl-name {
                display: inline-block;
                font-size: 20px;
                font-weight: 500;
                margin-right: 24px;
                height: 26px;
                line-height: 23px;
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
