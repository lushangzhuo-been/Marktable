<template>
    <div class="progress-private-role">
        <div class="btns">
            <el-input
                class="basic-ui"
                placeholder="输入关键词搜索"
                size="small"
                v-model="seachText"
                clearable
                suffix-icon="el-icon-search"
            ></el-input>
            <div>
                <el-button type="primary" size="small" @click="addRole">
                    +添加角色
                </el-button>
            </div>
        </div>
        <el-table
            class="progress-private-role-table"
            :data="roleTableData"
            style="width: 100%"
        >
            <el-table-column
                prop="name"
                width="200"
                show-overflow-tooltip
                label="名称"
            >
            </el-table-column>
            <el-table-column prop="desc" show-overflow-tooltip label="角色说明">
            </el-table-column>
            <el-table-column
                prop="action"
                show-overflow-tooltip
                label="操作"
                width="120"
            >
                <template slot-scope="scope">
                    <el-button
                        @click="editRole(scope.$index, scope.row)"
                        type="text"
                        size="small"
                        :disabled="btnDisabled(scope.row.sign)"
                    >
                        编辑
                    </el-button>
                    <el-button
                        @click.native.prevent="
                            deleteRole(scope.$index, scope.row)
                        "
                        type="text"
                        size="small"
                        :disabled="btnDisabled(scope.row.sign)"
                    >
                        删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
        <pagination
            v-show="total > 10"
            :total="total"
            @pageSizeChange="pageSizeChange"
            @curPageChange="curPageChange"
        ></pagination>
        <!-- 添加角色弹窗 -->
        <template>
            <add-progress-role
                ref="addProgressRoleDialog"
                @on-confirm="onConfirmaddRole"
            ></add-progress-role>
        </template>
        <!-- 编辑角色弹窗 -->
        <template>
            <edit-progress-role
                ref="editProgressRoleDialog"
                @on-confirm="onConfirmEditRole"
            ></edit-progress-role>
        </template>
        <!-- 删除角色弹窗 -->
        <template>
            <delete-progress-role
                ref="deleteProgressRoleDialog"
                @on-confirm="onConfirmDeleteRole"
            ></delete-progress-role>
        </template>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress_setting";
import Pagination from "@/pages/common/pagination.vue";
import addProgressRole from "./components/add_role_dialog.vue";
import editProgressRole from "./components/edit_role_dialog.vue";
import deleteProgressRole from "./components/delete_role_dialog.vue";
export default {
    components: {
        Pagination,
        addProgressRole,
        editProgressRole,
        deleteProgressRole
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        btnDisabled() {
            return (sign) => {
                return sign === "admin" ? true : false;
            };
        }
    },
    data() {
        return {
            seachText: "",
            roleTableData: [],
            page: 1,
            size: 10,
            total: 0
        };
    },
    watch: {
        $route() {
            this.seachText = "";
            this.fetchRoleList();
        },
        seachText: _.debounce(function () {
            this.fetchRoleList();
        }, 500)
    },
    mounted() {
        this.fetchRoleList();
    },
    methods: {
        pageSizeChange(size) {
            this.size = size;
            this.page = 1;
            this.fetchRoleList();
        },
        curPageChange(page) {
            this.page = page;
            this.fetchRoleList();
        },
        fetchRoleList() {
            this.$store
                .dispatch("fetchProgressRoleList", {
                    ws_id: this.curSpace.id,
                    tmpl_id: this.$route.params.id,
                    key: this.seachText,
                    page_num: this.page,
                    page_size: this.size
                })
                .then((res) => {
                    this.roleTableData = res.list;
                    this.total = res.total;
                });
        },
        // 添加角色弹窗显示
        addRole() {
            this.$refs.addProgressRoleDialog.openDialog();
        },
        // 确认添加成功后刷新并关闭弹窗
        onConfirmaddRole() {
            this.fetchRoleList();
            this.$refs.addProgressRoleDialog.cancel();
        },
        // 编辑角色弹窗显示
        editRole(index, row) {
            this.$refs.editProgressRoleDialog.openDialog(row);
        },
        // 确认编辑成功后刷新并关闭弹窗
        onConfirmEditRole() {
            this.fetchRoleList();
            this.$refs.editProgressRoleDialog.cancel();
        },
        // 删除角色弹窗显示
        deleteRole(index, row) {
            this.$refs.deleteProgressRoleDialog.openDialog(row);
        },
        // 确认删除成功后刷新并关闭弹窗
        onConfirmDeleteRole() {
            this.$refs.deleteProgressRoleDialog.cancel();
            this.fetchRoleList();
        }
    }
};
</script>

<style lang="scss" scoped>
.progress-private-role {
    .btns {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 16px;
        ::v-deep .el-input {
            width: 224px;
        }
    }
    ::v-deep .el-table {
        th.el-table__cell > .cell {
            font-weight: 400;
        }
        th.el-table__cell {
            background-color: #f5f6f8;
        }
        .el-table__cell {
            padding: 8px 0;
        }
    }
}
</style>
