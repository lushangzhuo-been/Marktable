<template>
    <div>
        <!-- <el-alert :title="title" type="warning" show-icon> </el-alert> -->
        <div class="margin-top-16">
            <span class="label">选择字段:</span>
            <el-select
                class="basic-ui"
                v-model="filed"
                size="small"
                placeholder="请选择字段"
                @change="filedChange"
            >
                <el-option
                    v-for="item in filedOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                >
                </el-option>
            </el-select>
        </div>
        <el-table
            class="setting-filed-table"
            :data="createPageTableData"
            style="width: 100%"
        >
            <el-table-column prop="name" show-overflow-tooltip label="字段名称">
            </el-table-column>
            <!-- <el-table-column
                prop="required"
                show-overflow-tooltip
                label="是否必填"
            >
                <template slot-scope="scope">
                    <el-switch
                        v-model="scope.row.required"
                        :active-color="PRIMARY_COLOR"
                        inactive-color="#E6E9F0"
                    >
                    </el-switch>
                </template>
            </el-table-column> -->

            <el-table-column
                prop="action"
                show-overflow-tooltip
                label="操作"
                width="120"
            >
                <template slot-scope="scope">
                    <el-button
                        @click="removeItem(scope.$index)"
                        type="text"
                        size="small"
                    >
                        删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
import Sortable from "sortablejs";
import { PRIMARY_COLOR_LIST } from "@/assets/tool/color_list";
export default {
    data() {
        return {
            title: "列表界面：根据任务列表所需的字段，在下拉列表中进行选择，加入下方列表后，即可设置是否必填或者顺序调整。",
            filed: "",
            filedOptions: [],
            createPageTableData: [],
            PRIMARY_COLOR: PRIMARY_COLOR_LIST.PRIMARY_COLOR
        };
    },
    mounted() {
        this.fetchTableData();
        // 行拖拽排序
        this.rowDrop();
    },
    methods: {
        switchChange() {
            // 调接口发送更改后的数据给后端并获取新数据
        },
        rowDrop() {
            let tbody = document.querySelector(".el-table__body-wrapper tbody");
            let _this = this;
            Sortable.create(tbody, {
                // or { name: "...", pull: [true, false, 'clone', array], put: [true, false, array] }
                group: {
                    name: "words",
                    pull: true,
                    put: true
                },
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
                    // 结束拖拽
                    let currRow = _this.createPageTableData.splice(
                        oldIndex,
                        1
                    )[0];
                    _this.createPageTableData.splice(newIndex, 0, currRow);
                    // 调接口发送更改后的数据给后端并获取新数据
                }
            });
        },
        removeItem(index) {
            this.createPageTableData.splice(index, 1);
        },
        filedChange(val) {
            this.filed = "";
            this.fetchTableData();
        },
        fetchTableData() {
            // 表格数据
            this.createPageTableData = [
                {
                    name: "创建人",
                    required: true
                },
                {
                    name: "bug名称",
                    required: true
                },
                {
                    name: "期望修复日期",
                    required: true
                },
                {
                    name: "bug描述",
                    required: true
                },
                {
                    name: "负责人",
                    required: true
                }
            ];
            // 下拉选项
            this.filedOptions = [
                {
                    label: "创建人",
                    value: "创建人"
                },
                {
                    label: "bug名称",
                    value: "bug名称"
                },
                {
                    label: "期望修复日期",
                    value: "期望修复日期"
                },
                {
                    label: "bug描述",
                    value: "bug描述"
                },
                {
                    label: "负责人",
                    value: "负责人"
                }
            ];
        }
    }
};
</script>

<style lang="scss" scoped>
.margin-top-16 {
    margin: 16px 0;
}
.label {
    margin-right: 16px;
}
// ::v-deep .el-table {
//     th.el-table__cell > .cell {
//         font-weight: 400;
//     }
//     th.el-table__cell {
//         background-color: #f5f6f8;
//     }
//     .el-table__cell {
//         padding: 8px 0;
//     }
// }
</style>
