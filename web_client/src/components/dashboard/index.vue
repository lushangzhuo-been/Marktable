<template>
    <div class="content">
        <grid-layout
            :layout.sync="layout"
            :col-num="colNum"
            :row-height="90"
            :is-draggable="true"
            :is-resizable="true"
            :is-mirrored="false"
            :vertical-compact="true"
            :margin="[10, 10]"
            :use-css-transforms="true"
        >
            <grid-item
                v-for="item in layout"
                :x="item.x"
                :y="item.y"
                :w="item.w"
                :h="item.h"
                :maxW="8"
                :maxH="4"
                :i="item.i"
                :key="item.i"
                @resize="resizeEvent"
                @move="moveEvent"
                @resized="resizedEvent"
                @container-resized="containerResizedEvent"
                @moved="movedEvent"
                drag-allow-from=".drag-icon"
            >
                <div class="grid-item-block">
                    <!-- 标题栏 -->
                    <div class="head-block">
                        <!-- 拖拽-标题 -->
                        <div class="left">
                            <b class="drag-icon"></b>
                            <el-input
                                v-if="item.rename"
                                :ref="`rename${item.i}`"
                                class="basic-ui"
                                size="small"
                                v-model="item.title"
                                @blur="renameInputBlur(item)"
                            ></el-input>
                            <div v-else class="title">
                                {{ item.title }}
                            </div>
                        </div>
                        <!-- 操作 -->
                        <div class="right">
                            <el-dropdown
                                trigger="click"
                                @command="
                                    (command) => gridOperation(command, item)
                                "
                            >
                                <b class="more-operation"></b>
                                <el-dropdown-menu slot="dropdown">
                                    <el-dropdown-item
                                        class="basic-ui"
                                        command="rename"
                                    >
                                        <b class="opertion-item-box rename"></b>
                                        编辑
                                    </el-dropdown-item>
                                    <el-dropdown-item
                                        class="basic-ui"
                                        command="delete"
                                    >
                                        <b class="opertion-item-box delete"></b>
                                        删除
                                    </el-dropdown-item>
                                    <el-dropdown-item
                                        class="basic-ui"
                                        command="delete"
                                    >
                                        <b class="opertion-item-box delete"></b>
                                        复制
                                    </el-dropdown-item>
                                    <el-dropdown-item
                                        class="basic-ui"
                                        command="delete"
                                    >
                                        <b class="opertion-item-box delete"></b>
                                        下载
                                    </el-dropdown-item>
                                </el-dropdown-menu>
                            </el-dropdown>
                        </div>
                    </div>
                    <div class="body-block">
                        <pie-chart
                            v-if="item.type === 'pie'"
                            :item="item"
                            :ref="'chart' + item.i"
                        ></pie-chart>
                        <bar-chart
                            v-if="item.type === 'bar'"
                            :item="item"
                            :ref="'chart' + item.i"
                        ></bar-chart>
                        <line-chart
                            v-if="item.type === 'line'"
                            :item="item"
                            :ref="'chart' + item.i"
                        ></line-chart>
                        <number
                            v-if="item.type === 'number'"
                            :item="item"
                            :ref="'chart' + item.i"
                        ></number>
                    </div>
                </div>
            </grid-item>
        </grid-layout>
        <delete-board-dialog
            ref="DeleteBoardDialog"
            @delete="deleteBoard"
        ></delete-board-dialog>
    </div>
</template>

<script>
import _ from "lodash";
import VueGridLayout from "vue-grid-layout";
import PieChart from "./echarts/pie_chart";
import BarChart from "./echarts/bar_chart";
import LineChart from "./echarts/line_chart";
import Number from "./echarts/number";
import DeleteBoardDialog from "./components/delete_board_dialog";
export default {
    components: {
        GridLayout: VueGridLayout.GridLayout,
        GridItem: VueGridLayout.GridItem,
        PieChart,
        BarChart,
        LineChart,
        Number,
        DeleteBoardDialog
    },
    data() {
        return {
            layout: [
                {
                    x: 0,
                    y: 4,
                    w: 3,
                    h: 4,
                    i: "0",
                    moved: false,
                    rename: false,
                    type: "pie",
                    title: "饼图"
                },
                {
                    x: 5,
                    y: 4,
                    w: 3,
                    h: 4,
                    i: "1",
                    moved: false,
                    rename: false,
                    type: "bar",
                    title: "柱形图"
                },
                {
                    x: 3,
                    y: 4,
                    w: 2,
                    h: 4,
                    i: "2",
                    moved: false,
                    rename: false,
                    type: "line",
                    title: "折线图"
                },
                {
                    x: 0,
                    y: 0,
                    w: 8,
                    h: 4,
                    i: "3",
                    moved: false,
                    rename: false,
                    type: "number",
                    title: "数字"
                }
            ],
            colNum: 8,
            index: 100
        };
    },
    watch: {},
    methods: {
        addChart(type, name) {
            let obj = {
                // 4是添加时默认的obj.w的宽度
                x: (this.layout.length * 4) % (this.colNum || 12),
                y: this.layout.length + (this.colNum || 12), // puts it at the bottom
                w: 4,
                h: 4,
                i: this.index,
                moved: false,
                rename: false,
                type: type,
                title: name
            };
            this.layout.push(obj);
            // Increment the counter to ensure key is always unique.
            this.index++;
        },
        resizeEvent() {},
        resizedEvent(i, newH, newW, newHPx, newWPx) {
            // 调用echarts重绘
            if (this.$refs["chart" + i][0].redraw) {
                this.$refs["chart" + i][0].redraw();
            }
        },
        moveEvent() {
            // 拿起时，获取dom最大高度/每行高度，可以获得当前最大行数， 用于配置阴影
        },
        movedEvent() {
            // 删除阴影
        },
        containerResizedEvent() {},
        gridOperation(command, item) {
            if (command === "rename") {
                this.girdItemRename(item);
            }
            if (command === "delete") {
                this.girdItemDelete(item);
            }
        },
        girdItemRename(item) {
            item.rename = true;
            this.$nextTick(() => {
                this.$refs[`rename${item.i}`][0].focus();
            });
        },
        renameInputBlur(item) {
            item.rename = false;
        },
        girdItemDelete(item) {
            this.$refs.DeleteBoardDialog.openDialog(item);
        },
        deleteBoard(msg) {
            let i = msg.i;
            let arr = _.cloneDeep(this.layout);
            let newArr = [];
            arr.forEach((item) => {
                if (item.i !== i) {
                    newArr.push(item);
                }
            });
            this.layout = newArr;
        }
    }
};
</script>

<style lang="scss" scoped>
.content {
    // background-color: #cccccc;
    background-color: #ffffff;
}

.vue-grid-item {
    box-sizing: border-box;
    background-color: #ffffff;
    touch-action: none;
    border: 1px solid #cccccc;
    border-radius: 8px;
}

// 内容结构
.grid-item-block {
    display: flex;
    flex-direction: column;
    height: 100%;

    .head-block {
        height: 40px;
        line-height: 40px;
        background: rgba($color: #f5f6f8, $alpha: 0.5);
        border-radius: 8px 8px 0 0;
        border-bottom: 1px solid #e6e9f0;
        display: flex;
        justify-content: space-between;
    }

    .body-block {
        flex: 1;
    }
}

.left {
    display: flex;
}

.drag-icon {
    display: inline-block;
    width: 24px;
    height: 24px;
    background-image: url(@/assets/image/common/icon_dragtable_move.png);
    background-size: 100% 100%;
    margin: 8px;
}
.right {
    .more-operation {
        display: inline-block;
        width: 24px;
        height: 24px;
        margin: 8px;
        background-image: url(@/assets/image/common/more-operation.png);
        background-size: 100% 100%;
        &:hover {
            background-image: url(@/assets/image/common/more-operation-active.png);
        }
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
}
.rename-input.el-input {
    width: 160px;
}
</style>
<style lang="scss">
.vue-grid-item.vue-grid-placeholder {
    background: #cccccc;
}
</style>
