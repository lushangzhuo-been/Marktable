<template>
    <div>
        <div class="sub-title">
            <img
                :src="require(`@/assets/image/common/progress_setting.png`)"
                alt=""
                width="20px"
                height="20px"
                class="icon"
            />
            权限配置
        </div>
        <div class="auth-content-flex">
            <div class="auth-left">
                <div class="small-title">权限类型</div>
                <div>
                    <div
                        class="auth-item"
                        v-for="item in authList"
                        :key="item.value"
                        :style="{
                            backgroundColor: item.active ? '#f1f9ff' : ''
                        }"
                        @click="authChange(item)"
                    >
                        {{ item.label }}
                    </div>
                </div>
            </div>

            <div class="auth-right">
                <div class="name">添加权限成员</div>
                <div class="desc">{{ curItem.desc || "" }}</div>
                <SelectDrapDownSpecial
                    :curItem="curItem"
                ></SelectDrapDownSpecial>
                <!-- <div class="note">以下成员拥有此操作权限</div>
                <div v-if="curItem.auth && curItem.auth.length">
                    <div
                        v-for="auth in curItem.auth"
                        :key="auth.key"
                        class="auth-item"
                    >
                        <span class="label">{{ `${auth.label}:` }}</span>
                        <span class="value">{{ auth.value }}</span>
                    </div>
                </div>-->
            </div>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import SelectDrapDownSpecial from "./common/select_drop_special";
export default {
    components: {
        // Filed,
        // Page,
        // Node,
        // NodeNew,
        // NoData,
        // SpaceInfo
        SelectDrapDownSpecial
    },
    computed: {
        progressTreeList() {
            return this.$store.state.progressTree;
        },
        // isNoData() {
        //     return (
        //         (this.progressTreeList && this.progressTreeList.length) || false
        //     )
        // },
        curSpace() {
            return this.$store.state.curSpace || {};
        }
    },
    data() {
        return {
            authList: [
                {
                    id: "view",
                    name: "查看",
                    desc: "设置列表、详情、附件、子任务的编辑权限"
                },
                {
                    id: "view1",
                    name: "导出",
                    desc: "设置列表、详情、附件、子任务的编辑权限"
                },
                {
                    id: "view2",
                    name: "创建",
                    desc: "设置列表、详情、附件、子任务的编辑权限"
                },
                {
                    id: "view3",
                    name: "编辑"
                },
                {
                    id: "view4",
                    name: "删除"
                },
                {
                    id: "view5",
                    name: "进展"
                }
            ],
            curItem: {
                // title: "添加权限成员",
                // desc: "设置列表、详情、附件、子任务的编辑权限",
                // auth: [
                //     {
                //         label: "空间角色",
                //         value: ["所有人", "所有人2", "所有人3"]
                //     },
                //     {
                //         label: "任务属性",
                //         value: ["所有人", "所有人2", "所有人3"]
                //     },
                //     {
                //         label: "成员",
                //         value: ["所有人", "所有人2", "所有人3"]
                //     }
                // ]
            }
        };
    },
    watch: {
        // $route() {
        //     this.activeName = "成员";
        //     this.searchInput = "";
        //     if (this.curItem && Object.keys(this.curItem).length) {
        //         this.switchAuth();
        //         this.getSpecialPersonList(); // 获取任务属性下拉
        //     }
        // },
        progressTreeList(arr) {
            // if (arr && arr.length) {
            //     if (
            //         arr.some(
            //             (item) => item.id === Number(this.$route.params.id)
            //         )
            //     ) {
            //         if (this.activeName === "filed") {
            //             this.$nextTick(() => {
            //                 this.$refs.filed.fetchFiledConfig();
            //                 this.$refs.filed.fetchFiledList();
            //             });
            //         } else if (this.activeName === "page") {
            //             this.$nextTick(() => {
            //                 this.$refs.page.fetchPageConfig();
            //             });
            //         } else {
            //             this.$nextTick(() => {
            //                 this.$refs.node.fetchNodeConfig();
            //             });
            //         }
            //         this.isShowSpaceInfo = false;
            //     } else {
            //         this.isShowSpaceInfo = true;
            //     }
            // }
        },
        $route() {
            this.authChange(this.authList[0], true); // 路由刷新时为true
        }
    },
    mounted() {
        this.getConfig();
    },
    methods: {
        getConfig() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id
            };
            this.$store
                .dispatch("progress_setting/fetchAuthConfig", params)
                .then((res) => {
                    if (res && Object.keys(res).length) {
                        this.authList = _.cloneDeep(res.auth_config || []);
                        if (res.auth_config && res.auth_config.length) {
                            this.authList = _.cloneDeep(res.auth_config);
                            this.$set(this.authList[0], "active", true);
                            this.curItem = this.authList[0];
                        } else {
                            this.authList = [];
                            this.curItem = {};
                        }
                    } else {
                        this.authList = [];
                    }
                });
        },
        authChange(item, type = false) {
            if (item.active && !type) return;
            this.curItem = _.cloneDeep(item);
            this.authList.forEach((sub) => {
                this.$set(sub, "active", sub.value === item.value);
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.auth-content-flex {
    display: flex;
    height: 800px;
    .auth-left {
        width: 200px;
        height: 100%;
        margin-right: 24px;
        background-color: #ffffff;
        border-radius: 4px 0px 0px 4px;
        border: 1px solid #e6e9f0;
        .small-title {
            height: 40px;
            line-height: 40px;
            margin-bottom: 10px;
            font-weight: 500;
            padding: 0 12px;
            background-color: #fafafb;
        }
        .auth-item {
            height: 40px;
            line-height: 40px;
            padding: 0 12px;
            cursor: pointer;
            &:hover {
                background-color: #fafafb;
            }
        }
    }
    .auth-right {
        width: calc(100% - 228px);
        height: 100%;
        .name {
            font-size: 16px;
            font-weight: 500;
            margin-bottom: 12px;
        }
        .desc {
            font-size: 12px;
            color: #959595;
            margin-bottom: 20px;
        }
        .note {
            font-weight: 700;
            margin: 16px 0 20px;
        }
        .auth-item {
            display: flex;
            align-items: center;
            margin-bottom: 12px;
            .label {
                display: inline-block;
                width: 64px;
                margin-right: 10px;
                text-align: right;
                color: #82889e;
            }
            .value {
                display: inline-block;
                width: calc(100% - 74px);
            }
        }
    }
}
</style>
