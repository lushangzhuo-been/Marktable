<template>
    <div v-if="isShow">
        <div class="space">
            <div
                @mouseleave="spaceNameMouseleave"
                @mouseenter="spaceNameMouseenter"
                class="space-line"
            >
                <!-- 切换空间 -->
                <el-dropdown
                    style="width: calc(100% - 20px)"
                    placement="bottom-start"
                    trigger="click"
                    width="200"
                    @command="switchSpace"
                >
                    <el-dropdown-menu
                        slot="dropdown"
                        class="absolute space-more-dropdown"
                    >
                        <div class="box">
                            <el-dropdown-item
                                v-for="(item, index) in spaceList"
                                :key="index"
                                :command="`${item.id}`"
                                class="space-more-item-space-list"
                                :class="item.active ? 'active' : ''"
                            >
                                <!-- <b class="icon"></b> -->
                                <span
                                    class="first-letter"
                                    :style="titleStyle(index)"
                                    >{{ firstLetter(item.name) }}</span
                                >
                                <tip-one :text="item.name" position="bottom">
                                    <span class="name">{{
                                        item.name || "--"
                                    }}</span>
                                </tip-one>
                            </el-dropdown-item>
                        </div>
                    </el-dropdown-menu>
                    <div class="switch-space-name">
                        <span class="space-name">
                            <img
                                v-show="!isMouseOverSwitch"
                                :src="
                                    require('@/assets/image/common/switch_space.png')
                                "
                                alt=""
                                :width="20"
                                :height="20"
                            />
                            <img
                                v-show="isMouseOverSwitch"
                                :src="
                                    require('@/assets/image/common/switch_space_active.png')
                                "
                                alt=""
                                :width="20"
                                :height="20"
                            />
                            <span
                                class="first-letter"
                                :style="{ ...titleStyleObj }"
                                >{{ firstLetter(spaceName) }}</span
                            >
                            <tip-one :text="spaceName" position="top">
                                <span class="name">{{
                                    spaceName || "--"
                                }}</span>
                            </tip-one>
                        </span>
                    </div>
                </el-dropdown>
                <!-- 切换空间右侧的三个点及下拉 -->
                <el-dropdown
                    placement="bottom-start"
                    trigger="click"
                    width="200"
                    popper-class="dropdown"
                    @command="clickSpaceMoreItem"
                    @visible-change="spaceMoreChange"
                >
                    <el-dropdown-menu
                        slot="dropdown"
                        class="absolute space-more-dropdown-right"
                    >
                        <el-dropdown-item
                            :command="`${item.name}`"
                            v-for="(item, index) in spaceMoreList"
                            :key="index"
                            class="space-more-item"
                            :class="
                                item.active
                                    ? 'active'
                                    : item.is_admin === 'no'
                                    ? 'disabled'
                                    : ''
                            "
                        >
                            <img :src="item.img" alt="" class="img-more" />
                            <span class="name">{{ item.name }}</span>
                        </el-dropdown-item>
                    </el-dropdown-menu>

                    <b
                        v-show="isMouseOver"
                        class="more"
                        :class="
                            isPopoverShow || isSpaceRoute ? 'is-select' : ''
                        "
                        @click.stop
                    ></b>
                </el-dropdown>
            </div>
            <el-dropdown
                placement="bottom"
                trigger="click"
                class="width100-2"
                style="padding: 0 16px"
                popper-class="dropdown"
                :visible-arrow="false"
                @command="addBtn"
            >
                <el-dropdown-menu
                    slot="dropdown"
                    class="absolute"
                    style="width: 220px"
                >
                    <el-dropdown-item
                        :command="`${item.key}`"
                        v-for="(item, index) in addBtnList"
                        :key="index"
                        class="space-more-item"
                        :class="item.active ? 'active' : ''"
                    >
                        <img :src="item.img" alt="" class="img-more" />
                        <span class="name">{{ item.name }}</span>
                    </el-dropdown-item>
                </el-dropdown-menu>
                <el-button
                    class="basic-ui width100 add-btn"
                    type="primary"
                    size="small"
                    :disabled="isAdd !== 'yes'"
                >
                    <b class="add"></b>添加
                </el-button>
            </el-dropdown>
        </div>
        <!-- 添加文件夹 -->
        <template>
            <add-folder-dialog
                ref="addFolderDialog"
                @on-confirm="onConfirmFolder"
            ></add-folder-dialog>
        </template>
        <!-- 外层添加流程 -->
        <template>
            <add-progress-dialog
                ref="addProgressDialog"
                @add-progress-success-refresh="addProgressSuccessRefresh"
            ></add-progress-dialog>
        </template>
        <template>
            <remove-current-space
                ref="removeSpaceDialog"
                @on-confirm="onConfirmRemoveSpace"
            ></remove-current-space>
        </template>
    </div>
</template>

<script>
// 三个点显示逻辑
// 1.鼠标悬浮整体： 显示默认； 鼠标离开： 不显示
// 2.鼠标悬浮三个点： 显示蓝色，popver打开： 显示蓝色， 路由包含space， 显示蓝色
// 3. popover关闭，不显示
import _ from "lodash";
import api from "@/common/api/module/space";
import AddFolderDialog from "../dialog/add_folder.vue";
import AddProgressDialog from "../dialog/add_progress.vue";
import TipOne from "@/pages/common/tooltip_one_line.vue";
import RemoveCurrentSpace from "../dialog/remove_current_space.vue";
import { TITLE_COLOT_ARR } from "@/assets/tool/color_list";
// 全局弹窗
import { JumpDialogBox } from "@/assets/tool/login";
export default {
    components: {
        AddFolderDialog,
        AddProgressDialog,
        TipOne,
        RemoveCurrentSpace
    },
    computed: {
        // 空间名更多三个点高亮
        isSpaceRoute() {
            if (this.$route.path.indexOf("/space/") > -1) {
                return true;
            } else {
                return false;
            }
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        workSpaceList() {
            return this.$store.state.workSpaceList;
        },
        firstLetter() {
            return (name) => {
                return name ? name.substring(0, 1) : "--";
            };
        },
        titleStyle() {
            return (index) => {
                let colorItem = this.colorList[index % this.colorList.length];
                return {
                    color: colorItem.color,
                    backgroundColor: colorItem.backgroundColor
                };
            };
        },
        switchProgress() {
            return this.$store.state.switchProgress;
        }
    },
    data() {
        return {
            colorList: TITLE_COLOT_ARR,
            isShow: false,
            // 鼠标悬浮整体
            isMouseOver: false,
            // popover显示
            isPopoverShow: false,
            isMouseOverSwitch: false,
            spaceName: "",
            spaceList: [],
            titleStyleObj: {
                color: "",
                backgroundColor: ""
            },
            addBtnList: [
                {
                    name: "添加流程",
                    key: "progress",
                    img: require(`@/assets/image/common/icon_add_progress.png`)
                },
                {
                    name: "添加文件夹",
                    key: "folder",
                    img: require(`@/assets/image/common/icon_add_folder.png`)
                }
            ],
            curSpaceMoreItem: {},
            spaceMoreList: [
                {
                    name: "空间详情",
                    path: "/space/detail",
                    img: require(`@/assets/image/common/progress_update.png`),
                    components: "spaceDetail"
                },
                {
                    name: "空间成员",
                    path: "/space/member",
                    img: require(`@/assets/image/field_type_icon/people_multiple.svg`),
                    components: "spaceMember"
                },
                {
                    name: "删除",
                    img: require(`@/assets/image/common/delete.svg`)
                }
            ],
            isAdd: "no"
        };
    },

    watch: {
        $route(to, from) {
            if (to.path.indexOf("/space/") === -1) {
                this.initData();
                if (to.path.indexOf("/progress/") > -1 && to.params.wsId) {
                    if (this.spaceList && this.spaceList.length) {
                        let tmp = this.spaceList.filter(
                            (space) => space.id == to.params.wsId
                        );
                        if (tmp.length) {
                            if (!this.switchProgress) {
                                this.switchSpace(to.params.wsId, "fresh");
                            }
                        } else {
                            JumpDialogBox.install({
                                code: 403,
                                msg: "空间不存在"
                            });
                        }
                    }
                }
                this.$store.commit("switchProgress", false);
            } else {
                this.getCurrentItem();
            }
        },
        workSpaceList: {
            // 深度监听空间列表，
            handler(arr) {
                if (arr && arr.length) {
                    this.spaceList = _.cloneDeep(arr);
                    this.isShow = true;
                } else {
                    this.spaceList = [];
                    this.isShow = false;
                }
            },
            deep: true
        },
        curSpace: {
            handler(obj) {
                if (!obj || !Object.keys(obj).length) return;
                if (this.workSpaceList && this.workSpaceList.length) {
                    let localSpace = localStorage.getItem("space");
                    if (
                        !localSpace ||
                        !localSpace.length ||
                        localSpace === "undefined"
                    ) {
                        this.$store.commit(
                            "setCurSpace",
                            this.workSpaceList[0]
                        );
                        this.spaceName = this.workSpaceList[0].name;
                        this.spaceId = this.workSpaceList[0].id;
                    } else {
                        let localId = JSON.parse(localSpace).id;
                        this.workSpaceList.forEach((item) => {
                            if (localId == item.id) {
                                this.spaceId = item.id;
                                this.spaceName = item.name;
                            }
                        });
                    }
                } else {
                    this.spaceId = "";
                    this.spaceName = "";
                }
                this.spaceList.length &&
                    this.spaceList.forEach((item, index) => {
                        if (item.id == this.spaceId) {
                            let colorItem = _.cloneDeep(
                                this.colorList[index % this.colorList.length]
                            );
                            this.titleStyleObj = {
                                color: colorItem.color,
                                backgroundColor: colorItem.backgroundColor
                            };
                            // this.titleStyleObj = _.cloneDeep(
                            //     this.colorList[index % this.colorList.length]
                            // )
                            // this.titleStyle.backgroundColor = '#E5EBFD'
                            this.$set(item, "active", true);
                        } else {
                            this.$set(item, "active", false);
                        }
                    });
                this.getCurWorkspaceInfo();
            },
            deep: true
        }
    },
    mounted() {
        if (this.$route.name !== "myHomePage") {
            this.fetchSpaceList();
        }
        this.getCurrentItem();
    },
    methods: {
        getCurWorkspaceInfo() {
            if (this.spaceName && this.spaceId) {
                api.getCurWorkspaceInfo({ ws_id: this.spaceId }).then((res) => {
                    if (
                        res &&
                        res.resultCode === 200 &&
                        res.data &&
                        Object.keys(res.data)
                    ) {
                        this.$store.commit(
                            "setWorkspaceHomePageInfo",
                            res.data
                        );
                        this.isAdd = res.data.is_admin;
                        this.spaceMoreList.forEach((item) => {
                            this.$set(item, "is_admin", res.data.is_admin);
                        });
                    } else {
                        this.$store.commit("setWorkspaceHomePageInfo", {});
                        this.isAdd = "no";
                        this.spaceMoreList.forEach((item) => {
                            this.$set(item, "is_admin", false);
                        });
                    }
                });
            }
        },
        // 删除空间确认
        onConfirmRemoveSpace(obj) {
            let params = {
                id: obj.spaceId
            };
            api.deleteWorkspace(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.$message({
                        showClose: true,
                        message: "删除空间成功",
                        type: "success"
                    });
                    this.$refs.removeSpaceDialog.cancel();
                    this.fetchSpaceList(true);
                }
            });
        },
        fetchSpaceList(isDelete = false) {
            // 调接口获取空间列表并设置当前空间
            this.$store.dispatch("fetchSpaceList").then((res) => {
                this.checkCurSpaceAndStyle(res, isDelete);
            });
        },
        checkCurSpaceAndStyle(res, isDelete) {
            if (res && res.length) {
                let curSpace = localStorage.getItem("space");
                if (!curSpace) {
                    this.$store.commit("setCurSpace", res[0]);
                    this.spaceName = res[0].name;
                    this.spaceId = res[0].id;
                } else {
                    let localId = JSON.parse(localStorage.getItem("space")).id;
                    if (res.some((item) => localId == item.id)) {
                        res.forEach((item) => {
                            if (localId == item.id) {
                                this.spaceId = item.id;
                                this.spaceName = item.name;
                            }
                        });
                    } else {
                        this.$store.commit("setCurSpace", res[0]);
                        this.spaceName = res[0].name;
                        this.spaceId = res[0].id;
                    }
                }
                this.switchSpace(this.spaceId, "fresh");
            } else {
                this.spaceName = "";
                this.spaceId = "";
                if (isDelete) {
                    this.$router.push({
                        name: "home"
                    });
                }
            }
        },
        initData() {
            this.spaceMoreList.forEach((item) => {
                this.$set(item, "active", false);
            });
            // 鼠标悬浮整体
            this.isMouseOver = false;
            // popover显示
            this.isPopoverShow = false;
            this.isMouseOverSwitch = false;
        },
        // 鼠标滑过空间名
        spaceNameMouseenter() {
            this.isMouseOver = true;
            this.isMouseOverSwitch = true;
        },
        spaceNameMouseleave() {
            this.isMouseOverSwitch = false;
            // 鼠标离开的时候判断，是否路由存在或者popover存在， 如果存在就为true
            if (this.isPopoverShow || this.isSpaceRoute) {
                this.isMouseOver = true;
            } else {
                this.isMouseOver = false;
            }
        },
        // 空间名更多下拉出现或隐藏
        spaceMoreChange(val) {
            this.isPopoverShow = val;
            this.spaceNameMouseleave();
        },
        // 刚进入check路由添加或移除样式
        getCurrentItem() {
            this.$nextTick(() => {
                this.spaceMoreList.length &&
                    this.spaceMoreList.forEach((item) => {
                        if (this.$route.path === item.path) {
                            this.curSpaceMoreItem = item;
                            this.$set(item, "active", true);
                        } else {
                            this.$set(item, "active", false);
                        }
                    });
                this.spaceNameMouseleave();
            });
        },
        // 切换更多下拉item
        clickSpaceMoreItem(name) {
            let isNoAuth = this.spaceMoreList.some(
                (sub) => name === sub.name && sub.is_admin === "no"
            );
            if (isNoAuth) return;
            if (name === "删除") {
                this.$refs.removeSpaceDialog.openDialog({
                    spaceName: this.spaceName,
                    spaceId: this.spaceId
                });
                return;
            }
            this.spaceMoreList.forEach((item) => {
                if (item.name === name) {
                    this.curSpaceMoreItem = item;
                    this.$set(item, "active", true);
                } else {
                    this.$set(item, "active", false);
                }
            });
            if (this.$route.path === this.curSpaceMoreItem.path) return;
            this.$router.push({ name: this.curSpaceMoreItem.components });
        },
        addBtn(key) {
            if (key === "folder") {
                this.$refs.addFolderDialog.openDialog(this.curSpace);
            } else {
                this.$refs.addProgressDialog.openDialog(this.curSpace);
            }
        },
        // 切换空间, 存入vuex中
        switchSpace(command, type) {
            if (type !== "fresh" && this.spaceId == command) return;
            this.spaceList.forEach((item, index) => {
                if (item.id == command) {
                    this.spaceName = item.name;
                    let colorItem = _.cloneDeep(
                        this.colorList[index % this.colorList.length]
                    );
                    this.titleStyleObj = {
                        color: colorItem.color,
                        backgroundColor: colorItem.backgroundColor
                    };
                    // this.$set(item, 'active', true)
                    this.$store.commit("setCurSpace", item);
                    if (type !== "fresh") {
                        this.$router.push({
                            path: `/workspace/${command}`
                        });
                    }
                } else {
                    // this.$set(item, 'active', false)
                }
            });
            // this.getCurrentItem()
        },
        // 添加文件夹成功后，刷新流程树
        onConfirmFolder() {
            this.$store.dispatch("fetchProgressTree", {
                ws_id: this.curSpace.id
            });
            // 将新增的流程存储到vuex中，在progress_tree.vue中监听流程树做样式路由选中处理
            this.$store.commit("deleteOrAddTree", {
                action: "add", // 删除: 'delete', 增加: 'add'
                id: "",
                type: "", // 文件夹还是文件 'tmpl' 或者 'ws_file'
                file_id: ""
            });
            this.$refs.addFolderDialog.cancel();
        },
        // 添加外层流程成功后，刷新流程树
        addProgressSuccessRefresh() {
            this.$store.dispatch("fetchProgressTree", {
                ws_id: this.curSpace.id
            });
            // 将新增的流程存储到vuex中，在progress_tree.vue中监听流程树做样式路由选中处理
            this.$store.commit("deleteOrAddTree", {
                action: "add", // 删除: 'delete', 增加: 'add'
                id: "",
                type: "", // 文件夹还是文件 'tmpl' 或者 'ws_file'
                file_id: ""
            });
            this.$refs.addProgressDialog.cancel();
        }
    }
};
</script>
<style lang="scss">
// ul.personal-center.el-dropdown-menu.el-popper {
//     // width: 112px !important;
//     padding: 8px;
//     .el-dropdown-menu__item,
//     .el-menu-item {
//         padding: 0 8px;
//     }
// }

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
ul.space-more-dropdown-right.el-dropdown-menu.el-popper {
    margin-top: 2px;
    padding: 8px;
}
ul.space-more-dropdown.el-dropdown-menu.el-popper {
    margin-top: 2px;
    padding: 0px;
    .box {
        margin: 8px 0 8px 8px;
        padding: 0 8px 0 0;
        max-height: 400px;
        overflow-y: auto;
    }
    .first-letter {
        display: inline-block;
        width: 20px;
        height: 20px;
        line-height: 20px;
        text-align: center;
        margin-right: 8px;
        font-size: 12px;
        font-weight: 500;
        border-radius: 4px;
    }
}
</style>
<style lang="scss" scoped>
.space {
    margin: 0 0 16px;

    .space-line {
        display: flex;
        align-items: center;
        padding: 0 16px 0 13px;
        margin: 6px 0;
        border-left: 3px solid #f5f6f8;
        // &:hover {
        //     background-color: #efefef;
        //     border-left: 3px solid #f5f6f8;
        //     .switch-space-name .name {
        //         color: var(--PRIMARY_COLOR);
        //     }
        // }
    }
    .switch-space-name {
        // height: 48px;
        cursor: pointer;
        .space-name {
            display: flex;
            align-items: center;
            width: calc(100% - 10px);
            max-width: 190px;
            border-radius: 6px;
            padding: 6px 6px 6px 4px;
            &:hover {
                background-color: #efefef;
                // .name {
                //     color: var(--PRIMARY_COLOR);
                // }
            }
        }
        img {
            vertical-align: middle;
            margin-right: 4px;
        }
        .first-letter {
            display: inline-block;
            width: 20px;
            height: 20px;
            line-height: 20px;
            text-align: center;
            margin-right: 8px;
            font-size: 12px;
            font-weight: 500;
            border-radius: 4px;
        }
        .name {
            display: inline-block;
            width: calc(100% - 26px - 30px);
            line-height: 20px;
            font-weight: 500;
            font-size: 16px;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
        }
    }
    .width100 {
        width: 100%;
    }
    .add-btn {
        ::v-deep &.el-button--small {
            font-size: 14px;
        }
        .add {
            display: inline-block;
            width: 16px;
            height: 16px;
            margin-right: 2px;
            vertical-align: text-bottom;
            background-image: url(@/assets/image/common/add_white.png);
            background-size: 100% 100%;
        }
    }
    .width100-2 {
        width: calc(100% - 32px);
    }
}
.img-more {
    display: inline-block;
    width: 20px;
    height: 20px;
    margin-right: 8px;
    vertical-align: middle;
}
.more {
    position: relative;
    top: 2px;
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
.space-more-item {
    height: 40px;
    line-height: 40px;
    padding-left: 8px;
    cursor: pointer;
}
.space-more-item-space-list {
    height: 40px;
    line-height: 40px;
    padding-left: 8px;
    display: flex;
    align-items: center;
    cursor: pointer;
    .name {
        display: inline-block;
        width: 166px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }
}
.space-more-item.active,
.space-more-item-space-list.active {
    // color: var(--PRIMARY_COLOR);
    background-color: #f1f9ff;
}
.space-more-item.disabled,
.space-more-item.disabled:hover,
.space-more-item.disabled:focus {
    background-color: #fff;
    color: #ccc;
    cursor: not-allowed;
    .img-more {
        opacity: 0.3;
    }
}
</style>
