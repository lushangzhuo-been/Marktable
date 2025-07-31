<template>
    <div class="progress-private-member">
        <!-- <div class="sub-title">空间</div> -->
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
                <!-- <el-button
                    type="primary"
                    size="small"
                    plain
                    @click="importRember"
                >
                    <i class="el-icon-download"></i>
                    导入成员
                </el-button> -->
                <el-button type="primary" size="small" @click="addRember">
                    +添加成员
                </el-button>
            </div>
        </div>
        <el-table
            class="progress-private-member-table"
            :data="remberTableData"
            style="width: 100%"
        >
            <el-table-column
                prop="userid"
                show-overflow-tooltip
                label="用户名"
                :min-width="110"
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
                :show-overflow-tooltip="false"
                label="邮箱"
                :min-width="250"
            >
                <template slot-scope="scope">
                    <div class="email-flex">
                        <tip-one :text="scope.row.email">
                            <span class="email">{{
                                scope.row.email || "--"
                            }}</span>
                        </tip-one>
                        <b
                            v-if="scope.row.email"
                            @click="copyEmail(scope.row.email)"
                            class="operation-item-box copy"
                        ></b>
                    </div>
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
                prop="role_name"
                min-width="110"
                show-overflow-tooltip
                label="角色"
                :width="100"
            >
                <template slot-scope="scope">
                    {{ scope.row.role_name || "--" }}
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
                        @click="editRember(scope.row)"
                        type="text"
                        size="small"
                    >
                        编辑
                    </el-button>
                    <el-button
                        @click.native.prevent="deleteRember(scope.row)"
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
            :total="total"
            @pageSizeChange="pageSizeChange"
            @curPageChange="curPageChange"
        ></pagination>
        <!-- 添加弹窗 -->
        <template>
            <add-progress-rember
                ref="addProgressRemberDialog"
                @on-confirm="onConfirmAddRember"
            ></add-progress-rember>
        </template>
        <!-- 编辑弹窗 -->
        <template>
            <edit-progress-rember
                ref="editProgressRemberDialog"
                @on-confirm="onConfirmEditRember"
            ></edit-progress-rember>
        </template>
        <!-- 删除弹窗 -->
        <template>
            <delete-progress-rember
                ref="deleteProgressRemberDialog"
                @on-confirm="onConfirmDeleteRember"
            ></delete-progress-rember>
        </template>
    </div>
</template>

<script>
import _ from "lodash";
import { imgHost } from "@/assets/tool/const";
import Pagination from "@/pages/common/pagination.vue";
import api from "@/common/api/module/progress_setting";
import addProgressRember from "./components/add_rember_dialog.vue";
import editProgressRember from "./components/edit_rember_dialog.vue";
import deleteProgressRember from "./components/delete_rember_dialog.vue";
import { PRIMARY_COLOR_LIST } from "@/assets/tool/color_list";
import TipOne from "@/pages/common/tooltip_one_line.vue";
export default {
    components: {
        Pagination,
        addProgressRember,
        editProgressRember,
        deleteProgressRember,
        TipOne
    },
    computed: {
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
        $route() {
            this.seachText = "";
            this.fetchMemberList();
            this.$store.dispatch("fetchProgressRoleList", {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id
            });
        },
        seachText: _.debounce(function () {
            this.fetchMemberList();
        }, 500)
    },
    mounted() {
        this.fetchMemberList();
    },
    methods: {
        copyEmail(email) {
            let tempTextarea = document.createElement("textarea");
            // 设置textarea的value为当前网址
            tempTextarea.value = email;
            // 将textarea添加到DOM中
            document.body.appendChild(tempTextarea);
            // 选中textarea中的文本
            tempTextarea.select();
            // 复制选中的文本
            document.execCommand("copy");
            // 移除临时的textarea元素
            document.body.removeChild(tempTextarea);
            this.$message({
                showClose: true,
                message: "复制成功",
                type: "success"
            });
        },
        pageSizeChange(size) {
            this.size = size;
            this.page = 1;
            this.fetchMemberList();
        },
        curPageChange(page) {
            this.page = page;
            this.fetchMemberList();
        },
        fetchMemberList() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                key: this.seachText,
                page_size: this.size,
                page_num: this.page
            };
            api.getRemberList(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.remberTableData = res.data.list || [];
                    this.total = res.data.cnt || 0;
                } else {
                    this.remberTableData = [];
                    this.total = 0;
                }
            });
        },
        // fetchRoleList() {
        //     let params = {
        //         ws_id: this.curSpace.id,
        //         tmpl_id: this.$route.params.id,
        //     }
        //     api.getRoleList(params).then((res) => {
        //         if (res && res.resultCode === 200 && res.data) {
        //             this.roleTableData = res.data.list || []
        //         } else {
        //             this.roleTableData = []
        //         }
        //     })
        // },
        switchChange() {},
        // 添加
        addRember() {
            this.$refs.addProgressRemberDialog.openDialog();
        },
        onConfirmAddRember() {
            this.fetchMemberList();
            this.$refs.addProgressRemberDialog.cancel();
        },
        // 编辑
        editRember(row) {
            this.$refs.editProgressRemberDialog.openDialog(row);
        },
        onConfirmEditRember() {
            this.fetchMemberList();
            this.$refs.editProgressRemberDialog.cancel();
        },
        // 删除
        deleteRember(row) {
            this.fetchMemberList();
            this.$refs.deleteProgressRemberDialog.openDialog(row);
        },
        onConfirmDeleteRember() {
            // this.remberTableData.splice(this.curIndex, 1)
            this.fetchMemberList();
            this.$refs.deleteProgressRemberDialog.cancel();
        }
    }
};
</script>

<style lang="scss" scoped>
.progress-private-member {
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
    .email-flex {
        display: flex;
        align-items: center;
        .email {
            max-width: calc(100% - 36px);
            overflow: hidden;
            white-space: nowrap;
            text-overflow: ellipsis;
        }
    }

    // 邮箱复制
    .operation-item-box {
        display: inline-block;
        width: 20px;
        height: 20px;
        background-size: 100% 100%;
        vertical-align: middle;
        cursor: pointer;
        &.copy {
            margin-left: 6px;
            background-image: url(@/assets/image/progress_column/copy.svg);
            &:hover {
                background-image: url(@/assets/image/progress_column/copy-active.svg);
            }
        }
    }
}
</style>
