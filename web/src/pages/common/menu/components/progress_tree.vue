<template>
    <!-- 流程树 -->
    <div class="progress-tree">
        <el-tree
            v-if="processData.length"
            ref="tree"
            node-key="id"
            highlight-current
            :default-expanded-keys="
                defaultExpandedKeys[curSpace.id] &&
                defaultExpandedKeys[curSpace.id].length
                    ? defaultExpandedKeys[curSpace.id]
                    : []
            "
            draggable
            :allow-drag="allowDrag"
            :allow-drop="allowDrop"
            :data="processData"
            :props="defaultProps"
            @node-click="handleNodeClick"
            @node-drop="nodeDrop"
            @node-expand="onNodeExpand"
            @node-collapse="onNodeCollapse"
            @node-drag-over="nodeDragOver"
        >
            <!-- 
              :class="!node.isLeaf ? 'not-leaf' : 'folder-leaf'" 
              叶子节点添加folder
          -->
            <span
                class="custom-tree-node"
                :class="!node.isLeaf ? 'not-leaf' : 'folder-leaf'"
                slot-scope="{ node, data }"
                @mouseenter="progressMouseEnter(data)"
                @mouseleave="progressMouseLeave(data)"
            >
                <!-- <span>{{ getIsPublicClass(data) }}</span> -->
                <b :class="getIsPublicClass(data)"></b>
                <tip-one :text="data.name" position="bottom-end">
                    <span>{{ data.name }}</span>
                </tip-one>
                <el-dropdown
                    placement="bottom-start"
                    trigger="click"
                    width="200"
                    style="position: absolute; right: 16px"
                    @visible-change="(val) => processMoreChange(val, data)"
                    @command="(item) => progressMorejumpTo(item, data)"
                >
                    <el-dropdown-menu
                        slot="dropdown"
                        class="absolute progress-more-dropdown"
                    >
                        <el-dropdown-item
                            v-for="(item, index) in getProgressDrapdownList(
                                data
                            )"
                            :key="index"
                            :command="`${item.name}`"
                            :style="{
                                display:
                                    data.mode === 'public' &&
                                    item.name === '角色&成员'
                                        ? 'none'
                                        : 'block'
                            }"
                            class="progress-more-item"
                            :class="
                                item.active && $route.path.indexOf(data.id) > -1
                                    ? 'active'
                                    : data.is_admin === 'no' ||
                                      (item.name === '删除' &&
                                          data.impl_table_name === 'ws_file' &&
                                          data.children &&
                                          data.children.length)
                                    ? 'disabled'
                                    : ''
                            "
                        >
                            <img :src="item.img" alt="" class="img-more" />
                            <span class="name">{{ item.name }}</span>
                        </el-dropdown-item>
                    </el-dropdown-menu>
                    <b
                        v-show="
                            data.isShow ||
                            (data.id === currentActivedItem.id &&
                                currentActivedItem.isShow)
                        "
                        class="more"
                        :class="
                            data.isSelect ||
                            (data.id === currentActivedItem.id &&
                                currentActivedItem.isSelect)
                                ? 'is-select'
                                : ''
                        "
                        @visible-change="morePopoverShow()"
                        @click.stop
                    ></b>
                </el-dropdown>
            </span>
        </el-tree>
        <!-- 移除流程弹窗 -->
        <template>
            <remove-folder-progress
                ref="removeProgressDialog"
                @on-confirm="onConfirmRemove"
            ></remove-folder-progress>
        </template>
        <!-- 文件夹更多中的添加流程 -->
        <template>
            <add-progress-dialog
                ref="addProgressDialog"
                @add-progress-success-refresh="addProgressSuccessRefresh"
            ></add-progress-dialog>
        </template>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress_setting";
import TipOne from "@/pages/common/tooltip_one_line.vue";
import RemoveFolderProgress from "../dialog/remove_folder_progress.vue";
import AddProgressDialog from "../dialog/add_progress.vue";
// import RemoveFolderProgress from "../dialog/remove_folder_progress.vue";
export default {
    //  1. 鼠标悬浮显示， 鼠标离开不显示
    //  2. 当前data.isShow
    components: {
        RemoveFolderProgress,
        TipOne,
        AddProgressDialog
    },
    computed: {
        // 删除或添加流程标识 默认为""， 删除流程时为"yes", 新增流程时"true"
        // 1. 新增流程时， 刷新流程树，不管是外层还是文件夹内层流程，都选中当前新增的流程
        // 2. 新增文件夹时，刷新流程树， 路由不发生变化
        // 3. 删除流程时， 如果删除的是当前选中的流程或者当前选中的流程三个点时， 跳转到首页， 如果删除的是非选中的流程，刷新流程树， 路由不发生变化
        // 4. 删除文件夹时， 如果删除的是当前选中的文件夹三个点时，刷新流程树， 跳转到首页， 如果删除的是其他，刷新流程树， 路由不发生变化
        deleteOrAddTree() {
            return this.$store.state.deleteOrAddTree;
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        progressTree() {
            return this.$store.state.progressTree;
        },
        getIsPublicClass() {
            return (data) => {
                if (data.impl_table_name === "ws_file") return "icon folder";
                return data.mode === "public" ? "icon public" : "icon private";
            };
        },
        isMoreProgressHighLight() {
            return (item) => {
                if (this.$route.path.indexOf(`/progress/${item.id}`) > -1) {
                    return true;
                } else {
                    return false;
                }
            };
        },
        workspaceHomePageInfo() {
            return this.$store.state.workspaceHomePageInfo;
        }
    },

    data() {
        return {
            defaultExpandedKeys: {},
            tmpProcessData: [],
            curActiveNode: {},
            processData: [],
            progressGroupMoreList: [
                {
                    name: "基本信息修改",
                    path: "/group/basicinfo/update",
                    img: require(`@/assets/image/common/progress_update.png`),
                    module: "progressGroupInfo"
                },
                {
                    name: "添加流程",
                    img: require(`@/assets/image/common/icon_add_progress.png`)
                },
                {
                    name: "删除",
                    img: require(`@/assets/image/common/delete.svg`)
                }
            ],

            progressMoreList: [
                {
                    name: "基本信息修改",
                    path: "/basicinfo/update",
                    img: require(`@/assets/image/common/progress_update.png`),
                    module: "progressInfo"
                },
                {
                    name: "流程设置",
                    path: "/setting",
                    children: [{}],
                    img: require(`@/assets/image/common/progress_setting.png`),
                    module: "progressSetting"
                },
                {
                    name: "自定义规则",
                    path: "/customer_rule",
                    children: [{}],
                    img: require(`@/assets/image/common/customer_rule.svg`),
                    module: "customerRole"
                },
                {
                    name: "角色&成员",
                    path: "/member",
                    children: [{}],
                    img: require(`@/assets/image/field_type_icon/people_multiple.svg`),
                    module: "progressMember"
                },
                {
                    name: "删除",
                    img: require(`@/assets/image/common/delete.svg`)
                }
            ],
            // folderMoreList: [
            //     {
            //         name: '基本信息修改',
            //         path: '/basicinfo/update',
            //         img: require(`@/assets/image/common/progress_update.png`),
            //     },
            //     {
            //         name: '删除',
            //         img: require(`@/assets/image/common/delete.svg`),
            //     },
            // ],
            currentActivedItem: {},
            defaultProps: {
                children: "children",
                label: "label",
                disabled: function (data, node) {
                    return !node.isLeaf;
                }
            }
        };
    },

    watch: {
        $route(to, from) {
            if (to.path.indexOf("/progress/") === -1) {
                // 如果是非流程树相关的页面，清除所有流程树的选中样式
                this.clearAllSelectStyle();
            } else {
                // 如果是流程树相关的路由页面，去设置选中样式
                this.$nextTick(() => {
                    this.getCurrentItem();
                });
            }
        },
        curSpace: {
            handler(obj) {
                if (!obj || !Object.keys(obj).length) return;
                this.fetchProgressTree();
            },
            deep: true
        },
        progressTree: {
            // 深度监听空间列表，
            handler(arr) {
                if (arr && arr.length) {
                    this.processData = _.cloneDeep(arr);
                    this.getCurrentItem();
                    this.isShow = true;
                    // arr.forEach((item) => {
                    //     if (
                    //         item.impl_table_name === "ws_file" &&
                    //         (!item.children || !item.children.length)
                    //     ) {
                    //         this.removeDefaultKeys(item.id);
                    //     }
                    // });
                } else {
                    this.processData = [];
                    this.isShow = false;
                }
            },
            deep: true
        }
    },
    mounted() {
        this.$nextTick(() => {
            this.getCurrentItem();
            this.defaultExpandedKeys = localStorage.getItem(
                "defaultExpandedKeys"
            )
                ? JSON.parse(localStorage.getItem("defaultExpandedKeys"))
                : {};
        });
    },
    methods: {
        morePopoverShow() {},
        fetchProgressTree() {
            this.$store
                .dispatch("fetchProgressTree", {
                    ws_id: this.curSpace.id
                })
                .then((res) => {
                    this.getCurrentItem();
                });
        },
        // 确定移除文件夹或流程调接口
        onConfirmRemove(data, form) {
            if (data.data.impl_table_name === "ws_file") {
                // 删除文件夹
                let params = {
                    ws_id: this.curSpace.id,
                    id: data.data.id
                };
                api.deleteFolder(params).then((res) => {
                    if (res && res.resultCode === 200) {
                        this.$message({
                            showClose: true,
                            message: "文件夹删除成功",
                            type: "success"
                        });
                        this.$refs.removeProgressDialog.cancel();
                        this.$store.commit("deleteOrAddTree", {
                            action: "delete", // 删除: 'delete', 增加: 'add'
                            id: data.data.id,
                            type: "ws", // 文件夹还是文件 'tmpl' 或者 'ws_file'
                            file_id: ""
                        });
                        this.fetchProgressTree(true);
                        this.removeDefaultKeys(data.data.id);
                    }
                });
            } else {
                // 删除流程
                let params = {
                    ws_id: this.curSpace.id,
                    id: data.data.id
                };
                api.deleteProgressTreeDetail(params).then((res) => {
                    if (res && res.resultCode === 200) {
                        this.$message({
                            showClose: true,
                            message: "流程删除成功",
                            type: "success"
                        });
                        this.$refs.removeProgressDialog.cancel();

                        // 删除的流程是当前选中的流程或者当前流程更多中的，跳到首页
                        // this.$store.commit("deleteOrAddTree", "yes");
                        this.$store.commit("deleteOrAddTree", {
                            action: "delete", // 删除: 'delete', 增加: 'add'
                            id: data.data.id,
                            type: "tmpl", // 文件夹还是文件 'tmpl' 或者 'ws_file'
                            file_id: ""
                        });
                        this.$store
                            .dispatch("fetchProgressTree", {
                                ws_id: this.curSpace.id
                            })
                            .then((res) => {
                                this.getCurrentItem();
                                this.processData.length &&
                                    this.processData.forEach((item) => {
                                        if (
                                            item.impl_table_name ===
                                                "ws_file" &&
                                            (!item.children ||
                                                !item.children.length)
                                        ) {
                                            this.removeDefaultKeys(item.id);
                                        }
                                    });
                            });
                    }
                });
            }
        },
        clearAllSelectStyle() {
            this.$nextTick(() => {
                // 清除三个点选中
                if (this.processData && this.processData.length) {
                    this.processData.forEach((item) => {
                        this.$set(item, "isShow", false);
                        this.$set(item, "isSelect", false);
                        if (item.children && item.children.length) {
                            item.children.forEach((sub) => {
                                this.$set(sub, "isShow", false);
                                this.$set(sub, "isSelect", false);
                            });
                        }
                    });
                }
                this.currentActivedItem = {};
                // 清除node选中
                this.$refs.tree && this.$refs.tree.setCurrentKey(null);
            });
        },
        // 鼠标移入显示default 三个点
        progressMouseEnter(data) {
            this.$set(data, "isShow", true);
        },
        // 鼠标移除时，如果data的isSelect为true， 显示active三个点， false不显示
        progressMouseLeave(data) {
            if (data.isSelect) {
                this.$set(data, "isShow", true);
            } else {
                this.$set(data, "isShow", false);
            }
        },

        // 刚进入页面check路由添加或移除样式
        getCurrentItem() {
            if (!this.processData.length) return; // 流程树不存在不做任何处理
            let route = this.$route;
            if (this.curSpace.id !== Number(route.params.wsId)) return; // 重复选中当前路由不做任何处理
            if (
                route.path.indexOf("/progress/") > -1 &&
                route.params &&
                route.params.wsId &&
                route.params.id
            ) {
                this.getCurrentActivedItem(this.processData); // 该方法获取到当前激活的文件夹或者流程节点this.currentActivedItem
                if (
                    !this.currentActivedItem ||
                    _.isEmpty(this.currentActivedItem)
                )
                    return; // 如果当前激活的文件夹或者流程节点，return不做任何样式处理
                if (route.meta.module !== "progress") {
                    // 如果module不是progress，说明选中的是文件夹或者流程的三个点中的item,
                    this.$set(this.currentActivedItem, "isShow", true);
                    this.$set(this.currentActivedItem, "isSelect", true);
                    this.getProgressDrapdownList(
                        this.currentActivedItem
                    ).forEach((item) => {
                        if (item.module === route.meta.module) {
                            this.$set(item, "active", true);
                        } else {
                            this.$set(item, "active", false);
                        }
                    });
                    // }
                } else {
                    // 选中的是流程，设置当前node为高亮, 文件夹不做处理
                    if (this.currentActivedItem.impl_table_name === "tmpl") {
                        this.$nextTick(() => {
                            this.$refs.tree &&
                                this.$refs.tree.setCurrentKey(
                                    Number(route.params.id)
                                );
                        });
                    }
                }
                // 删除的流程或者文件夹 是当前激活的流程或者三个点，则跳转到首页, 否则不做处理
                if (this.deleteOrAddTree.action === "delete") {
                    if (this.deleteOrAddTree.id == this.$route.params.id) {
                        this.$router.push({
                            name: "home" //注意使用 params 时一定不能使用 path
                        });
                    }
                    this.$store.commit("deleteOrAddTree", {});
                }
            }
        },

        getCurrentActivedItem(arr) {
            // 第一次进来，如果路由是文件夹或者流程的更多时候，循环流程树
            // 如果第一层就有流程id等于路由的id，则需要给当前流程右边的三个点高亮选中
            if (arr && arr) {
                arr.forEach((item) => {
                    if (item.id === Number(this.$route.params.id)) {
                        this.currentActivedItem = item;
                    }
                    this.$set(item, "active", false);
                    if (item.children && item.children.length) {
                        this.getCurrentActivedItem(item.children);
                    }
                });
            }
        },
        // 获取三个点选项，文件夹和流程的三个点选项不相同
        getProgressDrapdownList(data) {
            if (data.impl_table_name === "ws_file") {
                return this.progressGroupMoreList;
            } else {
                return this.progressMoreList;
            }
        },
        // 文件夹或者流程的三个点收起或者展开
        processMoreChange(flag, data) {
            if (!data) return;
            this.currentActivedItem = _.cloneDeep(data);
            if (flag) {
                this.$set(this.currentActivedItem, "isShow", true);
                this.$set(this.currentActivedItem, "isSelect", true);
            } else {
                this.$set(this.currentActivedItem, "isShow", false);
                this.$set(this.currentActivedItem, "isSelect", false);
            }
        },
        // 点击文件夹或者流程的三个点调转
        progressMorejumpTo(curName, data) {
            if (data.is_admin === "no") return; // 不管是文件夹还有流程的更多，如果没有权限点击都不操作
            // 点击删除文件或者删除流程
            if (curName === "删除") {
                if (
                    data.impl_table_name === "ws_file" &&
                    data.children &&
                    data.children.length
                ) {
                    return;
                }
                this.$refs.removeProgressDialog.openDialog({
                    data
                });
                return;
            }
            if (curName === "添加流程") {
                // 文件夹更多中的添加流程
                this.$refs.addProgressDialog.openDialog(this.curSpace, data);
                return;
            }
            let tmpObj = {};
            if (data.impl_table_name === "ws_file") {
                this.progressGroupMoreList.forEach((sub) => {
                    this.$set(sub, "active", curName === sub.name);
                    if (curName === sub.name) {
                        tmpObj = sub;
                    }
                });
            } else {
                this.progressMoreList.forEach((sub) => {
                    this.$set(sub, "active", curName === sub.name);
                    if (curName === sub.name) {
                        tmpObj = sub;
                    }
                });
            }
            // 如果当前的路由和选中的路由一样，就不跳转
            if (
                this.$route.meta.module === tmpObj.module &&
                Number(this.$route.params.id) == data.id
            )
                return;
            this.$store.commit("switchProgress", true);
            this.$router.push({
                path: `/progress/${this.curSpace.id}/${data.id}${tmpObj.path}`
            });
            this.processData.forEach((item) => {
                // 当前是点击的更多跳转，所以要循环整棵树，将所有的外层里层流程的选中清除掉
                this.$set(item, "isShow", false);
                this.$set(item, "isSelect", false);
                if (item.children && item.children.length) {
                    item.children.forEach((sub) => {
                        this.$set(sub, "isShow", false);
                        this.$set(sub, "isSelect", false);
                    });
                }
            });
            this.$refs.tree.setCurrentKey(null);
            this.$set(data, "isShow", true);
            this.$set(data, "isSelect", true);
        },
        // 确定新增流程
        addProgressSuccessRefresh(params) {
            this.addDefaultKeys(params.ws_file_id); // 根据文件夹中的流程数量打开收起
            this.$store.dispatch("fetchProgressTree", {
                ws_id: this.curSpace.id
            });
            this.$store.commit("deleteOrAddTree", {
                action: "add", // 删除: 'delete', 增加: 'add'
                id: "",
                type: "", // 文件夹还是文件 'tmpl' 或者 'ws_file'
                file_id: ""
            });
            this.$refs.addProgressDialog.cancel();
        },
        nodeDragStart(node, event) {},
        nodeDragEnter(node1, node2, event) {},
        removeActivedNode(val) {
            if (this.$refs.tree) {
                this.$refs.tree.setCurrentKey(val);
            }
        },
        // 不允许拿起
        allowDrag(node) {
            // 拿起之前先备份树
            this.tmpProcessData = _.cloneDeep(this.processData);
            // 仅当前空间的管理员可拖拽
            if (this.workspaceHomePageInfo.is_admin === "yes") {
                return true;
            }
            return false;
        },
        // 拖拽成功后, 排序
        nodeDrop(fromNode, toNode, dropType, event) {
            let params = {
                ws_id: this.curSpace.id,
                from_id: fromNode.data.id,
                from_type: fromNode.data.impl_table_name,
                to_id: toNode.data.id,
                to_type: toNode.data.impl_table_name,
                order: dropType
            };
            api.moveProgressTree(params).then((res) => {
                if (res && res.resultCode === 200) {
                    // this.fetchProgressTree();
                    // 移动的节点是文件夹下的
                    this.processData.length &&
                        this.processData.forEach((item) => {
                            if (
                                (!item.children || !item.children.length) &&
                                item.impl_table_name === "ws_file"
                            ) {
                                this.removeDefaultKeys(item.id);
                            }
                        });
                    if (
                        fromNode.data.impl_table_name === "tmpl" &&
                        toNode.data.impl_table_name === "tmpl" &&
                        toNode.data.ws_file_id
                    ) {
                        this.addDefaultKeys(toNode.data.ws_file_id);
                    }
                    if (
                        fromNode.data.impl_table_name === "tmpl" &&
                        toNode.data.impl_table_name === "ws_file" &&
                        toNode.data.id &&
                        dropType === "inner"
                    ) {
                        // 往空文件夹中拖
                        // targetFolderId = toNode.data.id;
                        this.addDefaultKeys(toNode.data.id);
                    }
                    this.processData = _.cloneDeep(this.processData);
                    this.getCurrentItem();
                } else {
                    this.processData = _.cloneDeep(this.tmpProcessData);
                }
            });
        },
        // 不允许放下,draggingNode拿起的node， dropNode放下的node， type操作类型
        allowDrop(draggingNode, dropNode, type) {
            if (
                draggingNode.data.impl_table_name === "ws_file" &&
                type === "inner"
            ) {
                return false;
            }
            // 如果目标节点是流程， 流程不能放任何东西
            if (dropNode.data.impl_table_name === "tmpl" && type === "inner") {
                return false;
            }
            if (
                draggingNode.data.impl_table_name === "ws_file" &&
                dropNode.level === 2
            ) {
                return false;
            }
            return true;
        },

        handleNodeClick(data, node, event) {
            //路由会切换
            if (node.data.impl_table_name === "ws_file") {
                this.$refs.tree.setCurrentKey(this.curActiveNode.id || null);
                return;
            }
            if (
                this.$route.path ===
                `/progress/${this.curSpace.id}${node.data.id}`
            )
                return;

            this.curActiveNode = _.cloneDeep(node.data);
            this.processData.forEach((item) => {
                this.$set(item, "isSelect", false);
                this.$set(item, "isShow", false);
                if (item.children && item.children.length) {
                    item.children.forEach((sub) => {
                        this.$set(sub, "isSelect", false);
                        this.$set(item, "isShow", false);
                    });
                }
            });
            this.$set(data, "isShow", true);
            // this.$set(data, 'isSelect', true)
            this.progressMoreList.forEach((sub) => {
                this.$set(sub, "active", false);
            });
            this.$store.commit("setCurProgressNodeData", node.data);
            this.$store.commit("switchProgress", true);
            this.$router.push({
                name: "progress", //注意使用 params 时一定不能使用 path
                params: { wsId: this.curSpace.id, id: node.data.id }
            });
            // this.switchProgress = true;
        },
        // 节点被展开时
        onNodeExpand(data, node, event) {
            this.addDefaultKeys(data.id);
        },
        // 节点被收起时
        onNodeCollapse(data, node, event) {
            this.$set(node, "expanded", false);
            this.removeDefaultKeys(data.id);
        },
        // 拖拽时候样式优化
        nodeDragOver() {
            //在此事件内修正offset解决el-tree在有滚动条时拖拽提示线错位的问题
            // let offsetTop = this.$refs["tree"].$el.scrollTop || 0;
            // let originalTop = document.querySelector(
            //     ".el-tree__drop-indicator"
            // ).offsetTop;
            // document.querySelector(".el-tree__drop-indicator").style.top =
            //     offsetTop + originalTop + "px";
        },
        addDefaultKeys(id) {
            if (
                this.defaultExpandedKeys[this.curSpace.id] &&
                this.defaultExpandedKeys[this.curSpace.id].length
            ) {
                let flag = false;
                this.defaultExpandedKeys[this.curSpace.id].some((item) => {
                    if (item == id) {
                        // 判断当前节点是否存在， 存在不做处理
                        flag = true;

                    }
                });
                if (!flag) {
                    // 不存在则存到数组里
                    this.defaultExpandedKeys[this.curSpace.id].push(id);
                }
            } else {
                this.defaultExpandedKeys[this.curSpace.id] = [];
                this.defaultExpandedKeys[this.curSpace.id].push(id);
            }
            localStorage.setItem(
                "defaultExpandedKeys",
                JSON.stringify(this.defaultExpandedKeys)
            );
        },
        removeDefaultKeys(id) {
            if (
                this.defaultExpandedKeys[this.curSpace.id] &&
                this.defaultExpandedKeys[this.curSpace.id].length &&
                this.defaultExpandedKeys[this.curSpace.id].indexOf(id) !== -1
            ) {
                this.defaultExpandedKeys[this.curSpace.id].splice(
                    this.defaultExpandedKeys[this.curSpace.id].indexOf(id),
                    1
                );
            } else {
                this.defaultExpandedKeys[this.curSpace.id] = [];
            }
            localStorage.setItem(
                "defaultExpandedKeys",
                JSON.stringify(this.defaultExpandedKeys)
            );
        }
    }
};
</script>
<style lang="scss">
.progress-tree {
    height: 100%;
    overflow-y: auto;
}
.progress-tree::-webkit-scrollbar {
    width: 6px;
    height: 10px;
}
.progress-tree::-webkit-scrollbar-track {
    background-color: #f5f6f8;
    border-radius: 3px;
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
ul.progress-more-dropdown.el-dropdown-menu.el-popper {
    // width: 144px !important;
    margin-top: 0px;
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
.progress-more-item {
    height: 40px;
    line-height: 40px;
    padding-left: 8px;
    cursor: pointer;
}
.progress-more-item.active {
    // color: var(--PRIMARY_COLOR);
    background-color: #f1f9ff;
}
.progress-more-item.disabled,
.progress-more-item.disabled:hover,
.progress-more-item.disabled:focus {
    background-color: #fff;
    color: #ccc;
    cursor: not-allowed;
    .img-more {
        opacity: 0.3;
    }
}
::v-deep .el-tree {
    height: 100%;
    overflow-y: auto;
    background-color: #f5f6f8;
    .el-tree__drop-indicator {
        display: none !important;
        z-index: 9999 !important;
        left: 0 !important;
    }
    .el-tree-node__expand-icon.is-leaf {
        display: none;
    }
    .el-tree-node__content > .el-tree-node__expand-icon {
        padding: 0;
    }
    .el-tree-node__expand-icon.el-icon-caret-right {
        position: absolute;
        left: 20px;
        z-index: 996;
        // opacity: 0;
        color: #a6abbc;
        border-radius: 2px;
        &::before {
            font-size: 16px;
        }
        &:hover {
            background-color: #e6e6e6;
            z-index: 1000;
        }
    }
    .el-tree-node__expand-icon.expanded {
        transform: unset;
        &::before {
            content: "\e790";
        }
    }
    .el-tree-node__content {
        padding: 0 !important;
        height: 40px;
        border-left: solid 3px #f5f6f8;
        position: relative;
        &:hover {
            background-color: #efefef;
            .progress-more {
                display: inline-block;
            }
        }
    }
    .custom-tree-node {
        display: flex;
        align-items: center;
        width: calc(100% - 32px);
        height: 40px;
        padding: 0 16px;

        span {
            display: inline-block;
            width: calc(100% - 40px);
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
        }
    }
    .custom-tree-node.not-leaf {
        // padding: 0 16px 0 0;
        // width: calc(100% - 16px);
        // span {
        //     width: calc(100% - 24px);
        // }
    }
    .custom-tree-node.not-leaf:hover {
        .icon.folder {
            opacity: 0;
            z-index: 988;
        }
    }
    //非叶子节点移除选中状态
    //选中.el-tree-node__content下包含.custom-tree-node.not-leaf类的.el-tree-node__content元素
    .el-tree-node.is-current
        > .el-tree-node__content:has(> .custom-tree-node.not-leaf) {
        border-left: solid 3px #f5f6f8;
        background-color: #f5f6f8;
    }
    .el-tree-node .el-tree-node__children .custom-tree-node {
        padding-left: 36px;
        width: calc(100% - 52px);
        span {
            display: inline-block;
            width: calc(100% - 40px);
        }
    }
    .el-tree-node.is-current > .el-tree-node__content {
        border-left: solid 3px var(--PRIMARY_COLOR);
        background-color: var(--WHITE_COLOR);
    }

    .icon {
        display: inline-block;
        width: 16px;
        height: 16px;
        margin: 4px;
    }
    .folder {
        background-image: url("~@/assets/image/common/icon_folder.png");
        background-size: 100% 100%;
        background-color: #f5f6f8;
        opacity: 1;
        z-index: 999;
    }
    .public {
        background-image: url("~@/assets/image/common/icon_public.svg");
        background-size: 100% 100%;
    }
    .private {
        background-image: url("~@/assets/image/common/icon_private.svg");
        background-size: 100% 100%;
    }
    .more {
        position: relative;
        right: -8px;
        top: 2px;
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
}
</style>
