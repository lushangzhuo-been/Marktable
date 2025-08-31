<template>
    <el-table-column
        :align="item.align"
        :fixed="item.fixed"
        :min-width="getWidth(item.mode, item.field_key)"
        :width="getWidth(item.mode, item.field_key)"
        :prop="item.columnName"
    >
        <template slot="header">
            <div>
                <b class="type-box-img"></b>
                {{ item.label }}
                <el-tooltip
                    popper-class="basic-ui"
                    effect="dark"
                    placement="top"
                >
                    <div slot="content">
                        {{ item.desc }}
                    </div>
                    <b v-show="item.desc" class="tip-box"></b>
                </el-tooltip>
            </div>
        </template>
        <template slot-scope="scope">
            <el-popover
                :ref="`progress_operate_${scope.row._id}`"
                placement="bottom-end"
                trigger="click"
                popper-class="progress-operate-col"
                :visible-arrow="false"
                @show="(val) => poppoverShow(scope.row, val)"
                @hide="(val) => poppoverHide(scope.row, val)"
            >
                <div class="pop-item" @click="copyItemLink(scope.row)">
                    <img
                        :src="
                            require(`@/assets/image/progress_column/copy.svg`)
                        "
                        alt=""
                        width="18px"
                        height="18px"
                    />
                    复制链接
                </div>
                <div
                    class="pop-item"
                    :class="hasDeleteAuth ? '' : 'disabled'"
                    @click="deleteRow(scope.row)"
                >
                    <img
                        :src="require(`@/assets/image/common/delete.svg`)"
                        alt=""
                        width="18px"
                        height="18px"
                    />
                    删除
                </div>
                <b
                    slot="reference"
                    :class="{
                        'more-icon': true,
                        actived: scope.row.isActived
                    }"
                    width="18px"
                    height="18px"
                >
                </b>
            </el-popover>
        </template>
    </el-table-column>
</template>

<script>
import { baseMixin } from "@/mixin.js";
import api from "@/common/api/module/progress";
export default {
    mixins: [baseMixin],
    props: {
        item: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            rowObj: {},
            hasDeleteAuth: false,
            isActived: false
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
    watch: {},
    mounted() {},
    methods: {
        poppoverShow(row) {
            this.$set(row, "isActived", true);
            this.fetAuthDelete(row);
        },
        poppoverHide(row) {
            this.$set(row, "isActived", false);
        },
        fetAuthDelete(row) {
            // 获取进展权限
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: row._id,
                auth_mode: "delete"
            };
            api.getUserAuth(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.hasDeleteAuth = res.data;
                } else {
                    this.hasDeleteAuth = false;
                }
            });
        },

        // 复制连接
        copyItemLink(row) {
            let jumpSite = window.location.origin;
            jumpSite += `/#/task/detail/${this.curSpace.id}/${this.curProgress}/${row._id}`;
            let tempTextarea = document.createElement("textarea");
            // 设置textarea的value为当前网址
            tempTextarea.value = jumpSite;
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
                message: "链接已复制",
                type: "success"
            });
            document.body.click();
        },
        // 删除行
        deleteRow(row) {
            if (!this.hasDeleteAuth) return;
            this.$emit("delete-row", row);
        },
        getType() {
            return {
                "background-image": "url(@/assets/image/common/edit.svg)"
            };
        }
    }
};
</script>

<style lang="scss" scoped>
.type-box-img {
    display: inline-block;
    width: 18px;
    height: 18px;
    vertical-align: middle;
    background-size: 100% 100%;
    background-image: url(@/assets/image/common/edit.svg);
}
.more-icon {
    display: inline-block;
    width: 18px;
    height: 18px;
    cursor: pointer;
    vertical-align: middle;
    background-size: 100% 100%;
    background-image: url(@/assets/image/common/operation_more_default.png);
    &.actived {
        background-image: url(@/assets/image/common/operation_more_active.png);
    }
    &:hover {
        background-image: url(@/assets/image/common/operation_more_active.png);
    }
}
</style>
