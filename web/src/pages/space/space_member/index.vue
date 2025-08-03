<template>
    <div class="space-member">
        <div class="sub-title">
            <img
                :src="
                    require(`@/assets/image/field_type_icon/people_multiple.svg`)
                "
                alt=""
                width="20px"
                height="20px"
                class="icon"
            />
            空间成员
        </div>
        <div class="btns">
            <el-input
                class="basic-ui"
                placeholder="输入关键词搜索"
                v-model="seachText"
                size="small"
                clearable
                suffix-icon="el-icon-search"
            ></el-input>
            <div>
                <!-- <el-button
                    type="primary"
                    size="small"
                    plain
                    @click="importRember"
                >
                    <i class="el-icon-download"></i>
                    导入
                </el-button> -->
                <el-button
                    class="basic-ui"
                    type="primary"
                    size="small"
                    @click="addRember"
                >
                    +添加
                </el-button>
            </div>
        </div>
        <el-table
            class="space-member-table"
            :data="remberTableData"
            style="width: 100%"
        >
            <el-table-column
                prop="userid"
                show-overflow-tooltip
                :min-width="110"
                label="用户名"
                width="200px"
            >
                <template slot-scope="scope">
                    {{ scope.row.username || "--" }}
                </template>
            </el-table-column>
            <el-table-column prop="full_name" label="昵称" :min-width="140">
                <template slot-scope="scope">
                    <div class="full-name-flex">
                        <el-avatar
                            icon="el-icon-user-solid"
                            size="small"
                            :src="getAvatar(scope.row)"
                        ></el-avatar>
                        <!-- <tip-one :text="scope.row.full_name"> -->
                        <span v-overflow class="full-name">
                            {{ scope.row.full_name || "--" }}</span
                        >
                        <!-- </tip-one> -->
                    </div>
                </template>
            </el-table-column>
            <el-table-column
                prop="phone"
                show-overflow-tooltip
                label="电话"
                :min-width="110"
            >
                <template slot-scope="scope">
                    {{ scope.row.phone || "--" }}
                </template>
            </el-table-column>
            <el-table-column
                prop="email"
                show-overflow-tooltip
                label="邮箱"
                :min-width="210"
            >
                <template slot-scope="scope">
                    {{ scope.row.email || "--" }}
                </template>
            </el-table-column>
            <el-table-column
                prop="status"
                show-overflow-tooltip
                label="状态"
                :width="80"
            >
                <template slot-scope="scope">
                    <div :class="statusClass(scope.row)">
                        {{ scope.row.status === "ok" ? "开启" : "禁用" }}
                    </div>
                </template>
            </el-table-column>
            <el-table-column
                prop="role"
                show-overflow-tooltip
                label="角色"
                :width="80"
            >
                <template slot-scope="scope">
                    {{ formartRole(scope.row.role) }}
                </template>
            </el-table-column>
            <el-table-column
                prop="action"
                show-overflow-tooltip
                label="操作"
                width="120"
            >
                <template slot-scope="scope">
                    <el-button
                        @click="editRember(scope.$index, scope.row)"
                        type="text"
                        size="small"
                    >
                        编辑
                    </el-button>
                    <el-button
                        @click.native.prevent="
                            deleteRember(scope.$index, scope.row)
                        "
                        type="text"
                        size="small"
                    >
                        删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
        <pagination
            v-show="total > 10"
            :page="page"
            :total="total"
            @pageSizeChange="pageSizeChange"
            @curPageChange="curPageChange"
        ></pagination>
        <!-- 导入弹窗 -->
        <template>
            <import-space-rember
                ref="importSpaceRemberDialog"
                @on-confirm="onConfirmImportRember"
            ></import-space-rember>
        </template>
        <!-- 添加弹窗 -->
        <template>
            <add-space-rember
                ref="addSpaceRemberDialog"
                @on-confirm="onConfirmAddRember"
            ></add-space-rember>
        </template>
        <!-- 编辑弹窗 -->
        <template>
            <edit-space-rember
                ref="editSpaceRemberDialog"
                @on-confirm="onConfirmEditRember"
            ></edit-space-rember>
        </template>
        <!-- 删除弹窗 -->
        <template>
            <delete-space-rember
                ref="deleteSpaceRemberDialog"
                @on-confirm="onConfirmDeleteRember"
            ></delete-space-rember>
        </template>
    </div>
</template>

<script>
import api from "@/common/api/module/space";
import _ from "lodash";
import { imgHost } from "@/assets/tool/const";
import Pagination from "@/pages/common/pagination.vue";
import importSpaceRember from "./components/import_rember_dialog.vue";
import addSpaceRember from "./components/add_rember_dialog.vue";
import editSpaceRember from "./components/edit_rember_dialog.vue";
import deleteSpaceRember from "./components/delete_rember_dialog.vue";
import { PRIMARY_COLOR_LIST } from "@/assets/tool/color_list";
import TipOne from "@/pages/common/tooltip_one_line.vue";
export default {
    components: {
        Pagination,
        addSpaceRember,
        editSpaceRember,
        deleteSpaceRember,
        importSpaceRember,
        TipOne
    },
    computed: {
        formartRole() {
            return (val) => {
                if (!val) return "--";
                return val === "admin" ? "管理员" : "普通";
            };
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        statusClass() {
            return (row) => {
                if (!row.status) return "";
                return row.status === "ok" ? "status ok" : "status forbid";
            };
        },
        getAvatar() {
            return (row) => {
                if (row && row.avatar) {
                    return `${imgHost}${row.avatar}`;
                }
                return require(`@/assets/image/common/default_avatar_big.png`);
            };
        }
    },
    data() {
        return {
            PRIMARY_COLOR: PRIMARY_COLOR_LIST.PRIMARY_COLOR,
            seachText: "",
            remberTableData: [],
            page: 1,
            size: 10,
            total: 0
        };
    },
    watch: {
        curSpace: {
            handler() {
                this.seachText = "";
                this.fetchMemberList();
            },
            deep: true
        },
        seachText: _.debounce(function () {
            this.fetchMemberList();
        }, 500)
    },

    mounted() {
        this.fetchMemberConfig();

        // 第一次进入时调接口，刷新不调接口
        // if (window.name === '') {
        this.fetchMemberList();
        // window.name = 'space_member'
        // }
    },
    beforeDestroy() {
        window.name = "";
    },
    methods: {
        pageSizeChange(size) {
            this.size = size;
            this.page = 1;
            this.fetchMemberList();
        },
        curPageChange(page) {
            this.page = page;
            this.fetchMemberList();
        },
        fetchMemberConfig() {},
        fetchMemberList() {
            let params = {
                ws_id: this.curSpace.id,
                page_size: this.size,
                page_num: this.page,
                key: this.seachText
            };
            api.getWorkspaceMemberList(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    this.remberTableData = res.data.list || [];
                    this.total = res.data.cnt || 0;
                } else {
                    this.remberTableData = [];
                    this.total = 0;
                }
            });
        },
        // 导入
        importRember() {
            this.$refs.importSpaceRemberDialog.openDialog();
        },
        onConfirmImportRember() {},
        // 添加
        addRember() {
            this.$refs.addSpaceRemberDialog.openDialog();
        },
        onConfirmAddRember() {
            this.$refs.addSpaceRemberDialog.cancel();
            this.fetchMemberList();
        },
        // 编辑
        editRember(index, row) {
            this.$refs.editSpaceRemberDialog.openDialog(row);
        },
        onConfirmEditRember() {
            this.$refs.editSpaceRemberDialog.cancel();
            this.fetchMemberList();
        },
        // 删除
        deleteRember(index, row) {
            this.$refs.deleteSpaceRemberDialog.openDialog(row);
        },
        onConfirmDeleteRember() {
            this.$refs.deleteSpaceRemberDialog.cancel();
            this.fetchMemberList();
        }
    }
};
</script>

<style lang="scss" scoped>
.space-member {
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
    .status {
        display: inline-block;
        height: 24px;
        line-height: 24px;
        padding: 0 10px;
        font-size: 12px;
        border-radius: 12px;
    }
    .status.ok {
        color: var(--PRIMARY_COLOR);
        background-color: var(--PRIMARY_COLOR_RGBA10);
    }
    .status.forbid {
        color: var(--STATUS_CLOSE);
        background-color: var(--STATUS_CLOSE_RGBA10);
    }
    .full-name-flex {
        display: flex;
        align-items: center;
        ::v-deep .el-avatar--small {
            width: 20px;
            height: 20px;
            line-height: 20px;
        }
        .full-name {
            margin-left: 8px;
            width: calc(100% - 48px);
            overflow: hidden;
            white-space: nowrap;
            text-overflow: ellipsis;
        }
    }
}
</style>
