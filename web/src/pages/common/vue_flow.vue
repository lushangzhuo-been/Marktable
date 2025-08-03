<template>
    <el-container class="workflow-container">
        <div class="super-flow-demo">
            <div class="node-container">
                <div
                    class="node-item"
                    v-for="(item, index) in nodeItemList"
                    :key="index"
                    @mousedown="(evt) => nodeItemMouseDown(evt, item.value)"
                >
                    {{ item.label }}
                </div>
            </div>
            <div class="flow-container" ref="flowContainer">
                <super-flow
                    ref="superFlow"
                    :graph-menu="graphMenu"
                    :node-menu="nodeMenu"
                    :link-menu="linkMenu"
                    :link-desc="linkDesc"
                    :node-list="nodeList"
                    :link-list="linkList"
                >
                    <template v-slot:node="{ meta }">
                        <div
                            :class="meta.type ? `flow-node-${meta.type}` : ''"
                            class="flow-node ellipsis"
                        >
                            <div class="node-content" :title="meta.name">
                                {{ meta.name }}
                            </div>
                        </div>
                    </template>
                </super-flow>
                <el-button
                    type="primary"
                    class="saveIcon basic-ui"
                    size="small"
                    @click="saveFlow"
                >
                    保存
                </el-button>
            </div>
        </div>

        <el-dialog
            :title="drawerConf.title"
            :visible.sync="drawerConf.visible"
            :close-on-click-modal="false"
            append-to-body
            width="500px"
            class="basic-ui"
        >
            <el-form
                @keyup.native.enter="settingSubmit"
                @submit.native.prevent
                v-show="drawerConf.type === drawerType.node"
                ref="nodeSetting"
                :model="nodeSetting"
            >
                <el-form-item label="节点名称" prop="name">
                    <el-input
                        class="basic-ui"
                        v-model="nodeSetting.name"
                        placeholder="请输入节点名称"
                        maxlength="30"
                    >
                    </el-input>
                </el-form-item>
                <el-form-item label="节点描述" prop="desc">
                    <el-input
                        class="basic-ui"
                        v-model="nodeSetting.desc"
                        placeholder="请输入节点描述"
                        maxlength="30"
                    >
                    </el-input>
                </el-form-item>
            </el-form>
            <el-form
                @keyup.native.enter="settingSubmit"
                @submit.native.prevent
                v-show="drawerConf.type === drawerType.link"
                ref="linkSetting"
                :model="linkSetting"
            >
                <el-form-item label="连线描述" prop="desc">
                    <el-input
                        v-model="linkSetting.desc"
                        placeholder="请输入连线描述"
                        class="basic-ui"
                    >
                    </el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button
                    @click="drawerConf.cancel"
                    class="basic-ui"
                    size="small"
                >
                    取消
                </el-button>
                <el-button
                    type="primary"
                    class="basic-ui"
                    size="small"
                    @click="settingSubmit"
                >
                    确定
                </el-button>
            </span>
        </el-dialog>
    </el-container>
</template>

<script>
import SuperFlow from "vue-super-flow";
import "vue-super-flow/lib/index.css";

const drawerType = {
    node: 0,
    link: 1
};

export default {
    components: {
        SuperFlow
    },
    data() {
        return {
            drawerType,
            nodeList: [
                {
                    id: "nodeeOC74zVI2ahlm7dz",
                    width: 100,
                    height: 32,
                    coordinate: [241, 30],
                    meta: {
                        label: "start",
                        name: "开始",
                        type: "start"
                    }
                },
                {
                    id: "nodeS9enKZNbCUytCiHe",
                    width: 168,
                    height: 168,
                    coordinate: [207, 109],
                    meta: {
                        label: "if",
                        name: "PM审批",
                        type: "if",
                        desc: "PM审批"
                    }
                },
                {
                    id: "nodeB4V7WNN6tcFJ2Bnk",
                    width: 150,
                    height: 32,
                    coordinate: [448, 177],
                    meta: {
                        label: "process",
                        name: "不同意",
                        type: "process"
                    }
                },

                {
                    id: "node52nfG098vaGn6M0D",
                    width: 168,
                    height: 168,
                    coordinate: [207, 320.0000122189522],
                    meta: {
                        label: "if",
                        name: "考勤专员审批",
                        type: "if",
                        desc: "考勤专员审批"
                    }
                },
                {
                    id: "nodeMzqPYGYFzLPEOWVO",
                    width: 100,
                    height: 32,
                    coordinate: [241, 535.8000000119209],
                    meta: {
                        label: "end",
                        name: "结束",
                        type: "end"
                    }
                }
            ],
            linkList: [
                {
                    id: "linkeUj01PShc10CiAqX",
                    startId: "nodeeOC74zVI2ahlm7dz",
                    endId: "nodeS9enKZNbCUytCiHe",
                    startAt: [50, 32],
                    endAt: [84, 0],
                    meta: {
                        desc: "提交"
                    }
                },
                {
                    id: "linkUn187rCtyJsLyjGS",
                    startId: "nodeB4V7WNN6tcFJ2Bnk",
                    endId: "nodeeOC74zVI2ahlm7dz",
                    startAt: [75, 0],
                    endAt: [100, 16],
                    meta: {
                        desc: "重新提交"
                    }
                },
                {
                    id: "link91yZLsmOwUcnsSyR",
                    startId: "node52nfG098vaGn6M0D",
                    endId: "nodeMzqPYGYFzLPEOWVO",
                    startAt: [84, 168],
                    endAt: [50, 0],
                    meta: {
                        desc: "同意并核销"
                    }
                },
                {
                    id: "linkEoBdEm2gCjbjQXNf",
                    startId: "nodeS9enKZNbCUytCiHe",
                    endId: "node52nfG098vaGn6M0D",
                    startAt: [84, 168],
                    endAt: [84, 0],
                    meta: {
                        desc: "同意"
                    }
                },
                {
                    id: "linkrxvLLA4aJvOa9QGo",
                    startId: "node52nfG098vaGn6M0D",
                    endId: "nodeB4V7WNN6tcFJ2Bnk",
                    startAt: [168, 84],
                    endAt: [75, 32],
                    meta: {
                        desc: "驳回"
                    }
                },
                {
                    id: "link27eL0iPUXCrFFVC1",
                    startId: "nodeS9enKZNbCUytCiHe",
                    endId: "nodeB4V7WNN6tcFJ2Bnk",
                    startAt: [168, 84],
                    endAt: [0, 16],
                    meta: {
                        desc: "驳回"
                    }
                }
            ],
            drawerConf: {
                title: "",
                visible: false,
                type: null,
                info: null,
                open: (type, info) => {
                    const conf = this.drawerConf;
                    conf.visible = true;
                    conf.type = type;
                    conf.info = info;
                    if (conf.type === drawerType.node) {
                        conf.title = "节点";
                        if (this.$refs.nodeSetting)
                            this.$refs.nodeSetting.resetFields();
                        this.$set(this.nodeSetting, "name", info.meta.name);
                        this.$set(this.nodeSetting, "desc", info.meta.desc);
                    } else {
                        conf.title = "连线";
                        if (this.$refs.linkSetting)
                            this.$refs.linkSetting.resetFields();
                        this.$set(
                            this.linkSetting,
                            "desc",
                            info.meta ? info.meta.desc : ""
                        );
                    }
                },
                cancel: () => {
                    this.drawerConf.visible = false;
                    if (this.drawerConf.type === drawerType.node) {
                        this.$refs.nodeSetting.clearValidate();
                    } else {
                        this.$refs.linkSetting.clearValidate();
                    }
                }
            },
            linkSetting: {
                desc: ""
            },
            nodeSetting: {
                name: "",
                desc: ""
            },
            nodeItemList: [
                {
                    label: "开始",
                    value: () => ({
                        width: 100,
                        height: 32,
                        meta: {
                            label: "start",
                            name: "开始",
                            type: "start"
                        }
                    })
                },
                {
                    label: "处理",
                    value: () => ({
                        width: 150,
                        height: 32,
                        meta: {
                            label: "process",
                            name: "处理",
                            type: "process"
                        }
                    })
                },
                {
                    label: "条件",
                    value: () => ({
                        width: 168,
                        height: 168,
                        meta: {
                            label: "if",
                            name: "条件",
                            type: "if"
                        }
                    })
                },
                {
                    label: "结束",
                    value: () => ({
                        width: 100,
                        height: 32,
                        meta: {
                            label: "end",
                            name: "结束",
                            type: "end"
                        }
                    })
                }
            ],
            graphMenu: [
                [
                    {
                        // 选项 label
                        label: "开始",
                        // 选项是否禁用
                        disable(graph) {
                            return !!graph.nodeList.find(
                                (node) => node.meta.label === "start"
                            );
                        },
                        // 选项选中后回调函数
                        selected(graph, coordinate) {
                            graph.addNode({
                                width: 100,
                                height: 32,
                                coordinate,
                                meta: {
                                    label: "start",
                                    name: "开始",
                                    type: "start"
                                }
                            });
                        }
                    },
                    {
                        label: "处理",
                        selected(graph, coordinate) {
                            graph.addNode({
                                width: 150,
                                height: 32,
                                coordinate,
                                meta: {
                                    label: "process",
                                    name: "处理",
                                    type: "process"
                                }
                            });
                        }
                    },
                    {
                        label: "条件",
                        selected(graph, coordinate) {
                            graph.addNode({
                                width: 168,
                                height: 168,
                                coordinate,
                                meta: {
                                    label: "if",
                                    name: "条件",
                                    type: "if"
                                }
                            });
                        }
                    }
                ],
                [
                    {
                        label: "结束",
                        selected(graph, coordinate) {
                            graph.addNode({
                                width: 100,
                                height: 32,
                                coordinate,
                                meta: {
                                    label: "end",
                                    name: "结束",
                                    type: "end"
                                }
                            });
                        }
                    }
                ],
                [
                    {
                        label: "全选",
                        selected: (graph) => {
                            graph.selectAll();
                        }
                    }
                ]
            ],
            nodeMenu: [
                [
                    {
                        label: "删除",
                        selected: (node) => {
                            node.remove();
                        }
                    },
                    {
                        label: "编辑",
                        selected: (node) => {
                            this.drawerConf.open(drawerType.node, node);
                        }
                    }
                ]
            ],
            linkMenu: [
                [
                    {
                        label: "删除",
                        selected: (link) => {
                            link.remove();
                        }
                    },
                    {
                        label: "编辑",
                        selected: (link) => {
                            this.drawerConf.open(drawerType.link, link);
                        }
                    }
                ]
            ],
            dragConf: {
                isDown: false,
                isMove: false,
                offsetTop: 0,
                offsetLeft: 0,
                clientX: 0,
                clientY: 0,
                ele: null,
                info: null
            }
        };
    },
    mounted() {
        document.addEventListener("mousemove", this.docMousemove);
        document.addEventListener("mouseup", this.docMouseup);
        this.$once("hook:beforeDestroy", () => {
            document.removeEventListener("mousemove", this.docMousemove);
            document.removeEventListener("mouseup", this.docMouseup);
        });
    },
    methods: {
        saveFlow() {
            this.nodeList = this.$refs.superFlow.toJSON().nodeList;
            this.linkList = this.$refs.superFlow.toJSON().linkList;
        },
        linkDesc(link) {
            return link.meta ? link.meta.desc : "";
        },
        settingSubmit() {
            const conf = this.drawerConf;
            if (this.drawerConf.type === drawerType.node) {
                if (!conf.info.meta) conf.info.meta = {};
                Object.keys(this.nodeSetting).forEach((key) => {
                    this.$set(conf.info.meta, key, this.nodeSetting[key]);
                });
                this.$refs.nodeSetting.resetFields();
            } else {
                if (!conf.info.meta) conf.info.meta = {};
                Object.keys(this.linkSetting).forEach((key) => {
                    this.$set(conf.info.meta, key, this.linkSetting[key]);
                });
                this.$refs.linkSetting.resetFields();
            }
            conf.visible = false;
        },
        docMousemove({ clientX, clientY }) {
            const conf = this.dragConf;
            if (conf.isMove) {
                conf.ele.style.top = clientY - conf.offsetTop + "px";
                conf.ele.style.left = clientX - conf.offsetLeft + "px";
            } else if (conf.isDown) {
                // 鼠标移动量大于 5 时 移动状态生效
                conf.isMove =
                    Math.abs(clientX - conf.clientX) > 5 ||
                    Math.abs(clientY - conf.clientY) > 5;
            }
        },
        docMouseup({ clientX, clientY }) {
            const conf = this.dragConf;
            conf.isDown = false;

            if (conf.isMove) {
                const { top, right, bottom, left } =
                    this.$refs.flowContainer.getBoundingClientRect();

                // 判断鼠标是否进入 flow container
                if (
                    clientX > left &&
                    clientX < right &&
                    clientY > top &&
                    clientY < bottom
                ) {
                    // 获取拖动元素左上角相对 super flow 区域原点坐标
                    const coordinate = this.$refs.superFlow.getMouseCoordinate(
                        clientX - conf.offsetLeft,
                        clientY - conf.offsetTop
                    );
                    // 添加节点
                    this.$refs.superFlow.addNode({
                        coordinate,
                        ...conf.info
                    });
                }
                conf.isMove = false;
            }
            if (conf.ele) {
                conf.ele.remove();
                conf.ele = null;
            }
        },
        nodeItemMouseDown(evt, infoFun) {
            const { clientX, clientY, currentTarget } = evt;

            const { top, left } = evt.currentTarget.getBoundingClientRect();

            const conf = this.dragConf;
            const ele = currentTarget.cloneNode(true);

            Object.assign(this.dragConf, {
                offsetLeft: clientX - left,
                offsetTop: clientY - top,
                clientX: clientX,
                clientY: clientY,
                info: infoFun(),
                ele,
                isDown: true
            });

            ele.style.position = "fixed";
            ele.style.margin = "0";
            ele.style.top = clientY - conf.offsetTop + "px";
            ele.style.left = clientX - conf.offsetLeft + "px";

            this.$el.appendChild(this.dragConf.ele);
        }
    }
};
</script>

<style lang="scss" scoped>
.workflow-container {
    width: calc(100vw - 80px);
    height: 700px;
    box-shadow: 0px 3px 1px -2px rgb(0 0 0 / 20%),
        0px 2px 2px 0px rgb(0 0 0 / 14%), 0px 1px 5px 0px rgb(0 0 0 / 12%);
    margin: 32px;
    padding: 0;
    background: #fff;
    // overflow: hidden;
}
.ellipsis {
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
    word-wrap: break-word;
}
.super-flow-demo {
    position: relative;
    margin: 20px;
    background: #f5f5f5;
    height: 100%;
    width: 100%;
    .node-container {
        width: 100%;
        height: 50px;
        background-color: #ffffff;

        .node-item {
            display: inline-block;
            font-size: 14px;
            height: 30px;
            width: 120px;
            margin: 0 20px 0 0;
            background: #ffffff;
            line-height: 30px;
            box-shadow: 1px 1px 4px rgba(0, 0, 0, 0.3);
            cursor: pointer;
            user-select: none; // 防止鼠标左键拖动选中页面的文字
            text-align: center;
            &:hover {
                box-shadow: 1px 1px 8px rgba(0, 0, 0, 0.4);
            }
        }
    }
    .flow-container {
        width: 100%;
        height: calc(100% - 50px);

        .super-flow {
            overflow: auto;
        }
    }
    .saveIcon {
        position: absolute;
        right: 0px;
        top: 0px;
    }
    .super-flow__node {
        .flow-node {
            box-sizing: border-box;
            width: 100%;
            height: 100%;
            line-height: 32px;
            padding: 0 6px;
            font-size: 16px;
            color: #fff;
            font-weight: bold;
            .node-content {
                text-align: center;
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
            }
        }
    }
    /*开始节点样式*/
    .ellipsis.flow-node-start {
        background: #55abfc;
        border-radius: 10px;
        border: 1px solid #b4b4b4;
    }
    /*流程节点样式*/
    .ellipsis.flow-node-process {
        position: relative;
        background: #30b95c;
        border: 1px solid #b4b4b4;
    }
    /*条件节点样式*/
    .ellipsis.flow-node-if {
        width: 120px;
        height: 120px;
        position: relative;
        top: 24px;
        left: 24px;
        background: #bc1d16;
        border: 1px solid #b4b4b4;
        transform: rotateZ(45deg); //倾斜
        .node-content {
            position: absolute;
            top: 50%;
            left: 20px;
            width: 100%;
            transform: rotateZ(-45deg) translateY(-75%);
        }
    }
    /*结束节点样式*/
    .ellipsis.flow-node-end {
        background: #000;
        border-radius: 10px;
        border: 1px solid #b4b4b4;
    }
}
</style>
<style>
.super-flow-demo .super-flow__node {
    border: none;
    background: none;
    box-shadow: none;
}
</style>
