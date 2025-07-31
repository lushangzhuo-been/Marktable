<template>
    <div class="create-page-index">
        <!-- <n-table class="table-content" :col="tableCol" :data="limiterData">
            <template slot="status" slot-scope="{ column }">
                <status-column :column="column"></status-column>
            </template>
            <template slot="status" slot-scope="{ column }">
                <status-column :column="column"></status-column>
            </template>
            <template slot="operation" slot-scope="{ column }">
                <el-table-column
                    show-overflow-tooltip
                    :width="200"
                    :label="column.name"
                >
                    <template slot-scope="scope">
                        <div>
                            <el-button
                                type="text"
                                size="small"
                                @click="openLimitEdit(scope.row)"
                            >
                                编辑
                            </el-button>
                            <el-button
                                type="text"
                                size="small"
                                @click="deleteLimit(scope.row)"
                            >
                                删除
                            </el-button>
                        </div>
                    </template>
                </el-table-column>
            </template>
        </n-table> -->
        <p-table
            ref="ProgressTable"
            :data="tableData"
            :col="tableCol"
        ></p-table>
    </div>
</template>
<script>
import NTable from "@/pages/common/tables/common_table/table";
import PTable from "@/components/table/p_table/progress_table";
export default {
    props: {
        modules: {
            type: String,
            default: ""
        },
        tableCol: {
            type: Array,
            default: () => []
        },
        dataCol: {
            type: Array,
            default: () => []
        }
    },
    components: {
        // NTable,
        PTable
    },
    data() {
        return {
            tableData: []
        };
    },
    methods: {}
};
</script>
<style lang="scss" scoped>
.create-page-index {
    height: 800px;
    background-color: #fafafa;
    display: flex;
    .create-page__left {
        position: relative;
        box-sizing: border-box;
        width: 60%;
        padding: 20px 24px;
        display: flex;
        align-items: center;
        justify-content: center;
        .title {
            position: absolute;
            top: 20px;
            left: 24px;
            font-size: 14px;
            color: #5c6479;
        }
    }
    .create-page__right {
        width: 40%;
        padding: 20px 24px;
        background: #ffffff;
        box-shadow: -3px 0px 10px 1px rgba(0, 0, 0, 0.1);
    }
}
</style>
