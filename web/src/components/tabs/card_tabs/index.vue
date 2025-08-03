<template>
    <!-- 卡片切换 流程页面用  1.支持拖拽排序 2.支持编辑 -->
    <div class="card-tabs-block">
        <el-tabs
            class="progress-tabs"
            :id="customerClass"
            :ref="customerClass"
            v-model="currentTabShow"
            type="card"
            @keydown.native.capture.stop
        >
            <el-tab-pane
                v-for="(item, index) in tabsListShow"
                :key="item.id"
                :label="item.name"
                :name="item.jsonId"
                :class="{}"
            >
                <div
                    class="tab-name-block"
                    :class="{
                        editing: item.editing,
                        current: currentTabShow
                    }"
                    slot="label"
                >
                    <b v-show="item.pin === 'yes'" class="thumbtack"></b>
                    <b
                        class="tab-type-icon drag-sort-block"
                        :class="item.mode"
                    ></b>
                    <div class="name-content" v-overflow>
                        <!-- <el-input
                            v-if="item.editing"
                            :ref="`rename${item.id}`"
                            class="basic-ui height24 padding8"
                            size="small"
                            v-model="item.name"
                        ></el-input> -->
                        <!-- @blur="renameInputBlur(item)" -->
                        <input
                            v-if="item.editing"
                            :ref="`rename${item.id}`"
                            class="ori-input rename-input"
                            size="small"
                            v-model="item.name"
                            @input.stop
                            @blur="renameInputBlur(item)"
                        />
                        <span v-show="!item.editing">
                            {{ item.name }}
                        </span>
                    </div>
                    <!-- 操作 -->
                    <el-dropdown
                        v-show="index !== 0"
                        trigger="click"
                        @command="(command) => tabOperation(command, item)"
                        class="operation-block-dropdown basic-ui"
                        placement="bottom-start"
                        @visible-change="
                            (boolean) => tabDropShow(boolean, item)
                        "
                    >
                        <b
                            class="operation-block"
                            :class="{ operating: item.operating }"
                            @click.stop
                        ></b>
                        <el-dropdown-menu slot="dropdown">
                            <el-dropdown-item class="basic-ui" command="rename">
                                <b class="opertion-item-box rename"></b>
                                重命名
                            </el-dropdown-item>
                            <el-dropdown-item
                                v-show="item.pin === 'no'"
                                class="basic-ui"
                                command="lock"
                            >
                                <b class="opertion-item-box lock"></b>
                                锁定视图
                            </el-dropdown-item>
                            <el-dropdown-item
                                v-show="item.pin === 'yes'"
                                class="basic-ui"
                                command="unlock"
                            >
                                <b class="opertion-item-box unlock"></b>
                                取消锁定
                            </el-dropdown-item>
                            <el-dropdown-item class="basic-ui" command="delete">
                                <b class="opertion-item-box delete"></b>
                                删除视图
                            </el-dropdown-item>
                        </el-dropdown-menu>
                    </el-dropdown>
                    <b class="tabs-partition"></b>
                </div>
            </el-tab-pane>
        </el-tabs>
        <el-dropdown
            trigger="click"
            @command="addOperation"
            class="add-btn-dropdown basic-ui"
            :style="{ left: addBtnLeft + 'px' }"
            placement="bottom-start"
            @visible-change="addDropShow"
        >
            <b class="add-btn" :class="{ operating: addOperating }"></b>
            <el-dropdown-menu slot="dropdown">
                <el-dropdown-item class="basic-ui" command="list">
                    <b class="tab-type-icon list"></b>
                    列表
                </el-dropdown-item>
                <el-dropdown-item class="basic-ui" command="board">
                    <b class="tab-type-icon board"></b>
                    仪表盘
                </el-dropdown-item>
                <el-dropdown-item class="basic-ui" command="card">
                    <b class="tab-type-icon card"></b>
                    看板
                </el-dropdown-item>
            </el-dropdown-menu>
        </el-dropdown>
        <delete-tab-view-dialog
            ref="DeleteTabViewDialog"
            @delete-view="deleteView"
        ></delete-tab-view-dialog>
    </div>
</template>

<script>
import DeleteTabViewDialog from "@/pages/progress/components/delete_tab_view_dialog.vue";
import Sortable from "sortablejs";
import api from "@/common/api/module/progress";
import index from "@/components/dashboard/index.vue";
import draggable from "vuedraggable";
import _ from "lodash";
import TipOne from "@/pages/common/tooltip_one_line.vue";
export default {
    components: {
        DeleteTabViewDialog,
        draggable,
        TipOne
    },
    props: {
        customerClass: {
            type: String,
            default: Date.now().toString()
        },
        tabsList: {
            type: Array,
            default: () => []
        },
        currentTab: {
            type: String,
            default: ""
        }
    },
    data() {
        return {
            currentTabShow: "", // 定位当前卡片
            tabsListShow: [],
            tabIndex: 2, //新增用
            addBtnLeft: 0,
            addOperating: false,
            draggableIdArr: ""
        };
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        }
    },
    watch: {
        tabsList: {
            handler(arr) {
                this.tabsListShow = arr;
                setTimeout(() => {
                    // 配置是否可拖动
                    // this.setSortAble();
                    this.getDraggableIdArr();
                    this.getAddBtnLeft();
                }, 60);
            },
            immediate: true
        },
        currentTab: {
            handler(str) {
                this.currentTabShow = str;
                this.$nextTick(() => {
                    this.getAddBtnLeft();
                });
            },
            immediate: true,
            deep: true
        },
        currentTabShow: {
            handler(str, oldStr) {
                let tab = _.find(this.tabsListShow, { jsonId: str });
                this.$emit("tab-change", tab);
            },
            deep: true
        }
    },
    mounted() {
        setTimeout(() => {
            this.getAddBtnLeft();
        }, 60);
        window.addEventListener("resize", this.getAddBtnLeft, true);
    },
    beforeDestroy() {
        window.removeEventListener("resize", this.getAddBtnLeft, true);
    },
    // onBeforeUnmount() {
    //     window.removeEventListener('resize', this.getAddBtnLeft);
    // },
    //     onBeforeUnmount(() => {
    // 	window.removeEventListener('resize', instanceResize);
    // });
    methods: {
        getDraggableIdArr() {
            let arr = [];
            this.tabsList.forEach((tabItem, tabIndex) => {
                if (tabItem.id !== -1) {
                    let id = "#tab-" + tabItem.id;
                    arr.push(id);
                }
            });
            this.draggableIdArr = arr.join(", ");
            this.rowDrop();
        },
        // tabClick() {
        //     this.$nextTick(() => {
        //         this.setSortAble();
        //     });
        // },
        // createView(){
        //     let params={
        //         ws_id: this.curSpace.id,
        //         tmpl_id: this.curProgress,
        //     }
        // },
        // refreshView() {
        //     let params = {
        //         ws_id: this.curSpace.id,
        //         tmpl_id: this.curProgress,
        //     };
        //     api.getView(params).then((res) => {
        //         this.tabsList = res.data;
        //         if (this.tabsList && this.tabsList instanceof Array) {
        //             this.tabsList.forEach((item) => {
        //                 item.jsonId = JSON.stringify(item.id);
        //             });
        //             this.currentTabShow = this.tabsList;
        //         }
        //     });
        // },
        onMove(e) {
            //不允许停靠
            // if (e.relatedContext.element.type === "mainText") {
            //     return false;
            // }
            //不允许拖拽
            // if (e.draggedContext.element.type === "mainText") {
            //     return false;
            // }
            return true;
        },
        tabDropShow(boolean, item) {
            if (boolean) {
                this.$set(item, "operating", true);
            } else {
                this.$set(item, "operating", false);
            }
        },
        addDropShow(boolean) {
            if (boolean) {
                this.addOperating = true;
            } else {
                this.addOperating = false;
            }
        },
        tabOperation(command, item) {
            if (command === "rename") {
                this.tabItemRename(item);
            }
            if (command === "delete") {
                this.tabItemDelete(item);
            }
            if (command === "lock") {
                this.tabItemLock(item);
            }
            if (command === "unlock") {
                this.tabItemUnlock(item);
            }
        },
        tabItemRename(item) {
            this.$set(item, "editing", true);
            this.$nextTick(() => {
                setTimeout(() => {
                    this.getAddBtnLeft();
                    this.$refs[`rename${item.id}`][0].focus();
                }, 60);
            });
        },
        tabItemDelete(item) {
            this.$refs.DeleteTabViewDialog.openDialog(item);
        },
        deleteView(view) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: view.id
            };
            api.deleteView(params).then((res) => {
                if (this.$route.name !== "progress") return;
                // 如果当前选中的是被删除的视图,则先置空
                this.$emit("delete-view");
            });
        },
        tabItemLock(item) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: item.id
            };
            api.pinView(params).then((res) => {
                if (this.$route.name !== "progress") return;
                if (res && res.resultCode === 200) {
                    this.$emit("lock-tab", item);
                }
            });
        },
        coordinate() {
            let idSortArr = [];
            this.tabsListShow.forEach((item) => {
                // 移除主视图
                if (item.id !== -1) {
                    idSortArr.push(item.id);
                }
            });
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                coordinate: idSortArr.join(",")
            };
            api.coordinate(params).then((res) => {});
        },
        tabItemUnlock(item) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: item.id
            };
            api.unPinView(params).then((res) => {
                if (this.$route.name !== "progress") return;
                if (res && res.resultCode === 200) {
                    this.$emit("unlock-tab", item);
                }
            });
        },
        renameInputBlur(item) {
            this.renameApi(item);
            this.$set(item, "editing", false);
            this.getAddBtnLeft();
        },
        renameApi(item) {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: item.id,
                name: item.name
            };
            api.viewRename(params).then((res) => {});
        },
        addOperation(command) {
            this.$emit("add-tab", command);
            // if (command === "list") {
            //     // 调接口添加
            //     this.$emit("add-tab", "list");
            // }
            // if (command === "board") {
            //     // 调接口添加
            //     this.$emit("add-tab", "board");
            // }
        },
        // 获取el-tabs__nav的宽度， 用于设置add-btn的相对定位
        handleResize(rect) {
            // 在这里处理DOM元素宽高变化后的逻辑
        },
        getAddBtnLeft() {
            // ref获取 tabs中的scroll 跟 nav
            setTimeout(() => {
                if (
                    this.$refs[this.customerClass] &&
                    this.$refs[this.customerClass].$refs.nav
                ) {
                    let father = this.$refs[this.customerClass];
                    let son = father.$refs.nav;
                    let scroll = son.$refs.navScroll;
                    let scrollWidth = scroll.clientWidth;
                    let nav = son.$refs.nav;
                    let navWidth = nav.clientWidth;
                    if (navWidth < scrollWidth) {
                        this.addBtnLeft = navWidth + 14;
                    } else {
                        this.addBtnLeft = scrollWidth + 44;
                    }
                }
            }, 60);
        },

        rowDrop() {
            let tabsDom = document.querySelector(
                ".progress-tabs .el-tabs__nav"
            );
            if (!tabsDom) {
                return;
            }
            let _this = this;
            Sortable.create(tabsDom, {
                draggable: _this.draggableIdArr,
                handle: ".drag-sort-block",
                animation: 150, // ms, number 单位：ms，定义排序动画的时间
                onAdd: function (evt) {
                    // 拖拽时候添加有新的节点的时候发生该事件
                },
                onUpdate: function (evt) {
                    // 拖拽更新节点位置发生该事件
                },
                onRemove: function (evt) {
                    // 删除拖拽节点的时候促发该事件
                },
                onStart: function (evt) {
                    // 开始拖拽出发该函数
                },
                onSort: function (evt) {
                    // 发生排序发生该事件
                },

                onEnd({ newIndex, oldIndex }) {
                    let currRow = _this.tabsListShow.splice(oldIndex, 1)[0];
                    _this.tabsListShow.splice(newIndex, 0, currRow);
                    _this.coordinate();
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
::v-deep.ori-input.rename-input {
    box-sizing: border-box;
    width: 100%;
    height: 24px;
    padding: 0 8px;
    border: 1px solid #dcdfe6;
    border-radius: 4px;
    &:focus {
        border: 1px solid #1890ff;
    }
    &:focus-visible {
        border: 1px solid #1890ff;
        outline: none;
    }
}
::v-deep.progress-tabs.el-tabs {
    .el-tabs__header {
        .el-tabs__nav-wrap {
            width: calc(100% - 120px);
            .el-tabs__nav-scroll {
                .el-tabs__nav {
                    border: 1px solid rgba(0, 0, 0, 0);
                    height: 32px;
                    .el-tabs__item {
                        box-sizing: border-box;
                        height: 32px;
                        line-height: 32px;
                        padding: 0;
                        max-width: 160px;
                        border-left: none;
                        &.is-active {
                            border-radius: 4px 4px 0 0;
                            background: linear-gradient(
                                to bottom,
                                #e5effa,
                                #ffffff
                            );
                            .tab-name-block {
                                border: 1px solid #e6e9f0;
                                border-bottom: 1px solid #ffffff;
                                border-radius: 4px 4px 0 0;
                                .tabs-partition {
                                    display: none;
                                }
                            }
                        }
                        &:hover {
                            .tab-name-block {
                                .operation-block {
                                    background-image: url(@/assets/image/common/more-operation-active.png);
                                }
                                &.editing {
                                    .operation-block {
                                        display: none;
                                    }
                                }
                            }
                        }
                        .tab-name-block {
                            padding: 0 20px 0 12px;
                            position: relative;
                            box-sizing: border-box;
                            height: 31px;
                            .name-content {
                                display: inline-block;
                                margin-left: 4px;
                                min-width: 48px;
                                max-width: 104px;
                                white-space: nowrap;
                                overflow: hidden;
                                text-overflow: ellipsis;
                                vertical-align: bottom;
                                font-weight: 500;
                                font-size: 14px;
                                color: #2f384c;
                            }
                            .operation-block-dropdown {
                                position: absolute;
                                right: 4px;
                                .operation-block {
                                    display: inline-block;
                                    width: 16px;
                                    height: 16px;
                                    vertical-align: middle;
                                    background-size: 100% 100%;
                                    &.operating {
                                        background-image: url(@/assets/image/common/more-operation-active.png);
                                    }
                                }
                            }
                            .tabs-partition {
                                position: absolute;
                                right: -1px;
                                top: 4px;
                                display: inline-block;
                                width: 0px;
                                height: 20px;
                                border-right: 1px solid #e6e9f0;
                            }
                        }
                    }
                }
            }
            .el-tabs__nav-prev {
                line-height: 32px;
            }
            .el-tabs__nav-next {
                line-height: 32px;
            }
        }
    }
}
.card-tabs-block {
    position: relative;
    .add-btn-dropdown {
        position: absolute;
        top: 2px;
        cursor: pointer;
        padding: 4px;
    }
}
.add-btn {
    display: inline-block;
    width: 12px;
    height: 12px;
    vertical-align: middle;
    background-image: url("~@/assets/image/progress/add-card.png");
    background-size: 100% 100%;
    &:hover {
        background-image: url("~@/assets/image/progress/add-card-active.png");
    }
    &.operating {
        background-image: url("~@/assets/image/progress/add-card-active.png");
    }
}
.opertion-item-box {
    display: inline-block;
    width: 20px;
    height: 20px;
    vertical-align: middle;
    background-size: 100% 100%;
    &.rename {
        background-image: url(@/assets/image/common/edit.svg);
    }
    &.delete {
        background-image: url(@/assets/image/common/delete.svg);
    }
    &.lock {
        background-image: url(@/assets/image/common/lock.svg);
    }
    &.unlock {
        background-image: url(@/assets/image/common/unlock.svg);
    }
}
.thumbtack {
    position: absolute;
    top: 4px;
    left: 4px;
    display: inline-block;
    width: 6px;
    height: 8px;
    vertical-align: middle;
    background-size: 100% 100%;
    background-image: url(@/assets/image/progress/thumbtack.png);
}
.tab-type-icon {
    display: inline-block;
    width: 18px;
    height: 18px;
    vertical-align: middle;
    position: relative;
    top: -2px;
    background-size: 100% 100%;
    &.list {
        background-image: url(@/assets/image/progress/list.svg);
    }
    &.board {
        background-image: url(@/assets/image/progress/board.svg);
    }
    &.card {
        background-image: url(@/assets/image/progress/card.svg);
    }
}
</style>
