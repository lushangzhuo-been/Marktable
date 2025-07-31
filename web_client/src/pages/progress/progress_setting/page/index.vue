<template>
    <div>
        <el-button-group>
            <el-button
                plain
                v-for="(item, index) in btnsList"
                :key="index"
                class="basic-ui"
                :class="item.active ? 'active-btn' : ''"
                @click="clickBtn(item)"
            >
                {{ item.label }}
            </el-button>
        </el-button-group>
        <div class="page-content">
            <create-page
                v-if="curBtn === 'create'"
                :modules="curBtn"
            ></create-page>
            <list-page v-if="curBtn === 'list'" :modules="curBtn"></list-page>
            <detail-page
                v-if="curBtn === 'detail'"
                :modules="curBtn"
            ></detail-page>
        </div>
    </div>
</template>

<script>
import CreatePage from "./create_page/index.vue";
import ListPage from "./list_page/index.vue";
import DetailPage from "./detail_page/index.vue";
import api from "@/common/api/module/progress_setting";
import _ from "lodash";
export default {
    components: {
        CreatePage,
        ListPage,
        DetailPage
    },
    data() {
        return {
            curBtn: "create",
            btnsList: [
                {
                    label: "创建页面",
                    value: "create",
                    active: true
                },
                // {
                //     label: "列表页面",
                //     value: "list",
                //     active: false
                // },
                {
                    label: "详情页面",
                    value: "detail",
                    active: false
                }
            ]
        };
    },
    methods: {
        fetchPageConfig() {
            api.getScreenConfig().then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    let moduleConfig = res.data.module_config;
                    if (moduleConfig && moduleConfig.length) {
                        moduleConfig.forEach((item) => {
                            item.active =
                                item.value === "create" ? true : false;
                            if (item.value === "create") {
                                this.curBtn = "create";
                            }
                        });
                        // this.btnsList = _.cloneDeep(moduleConfig)
                    }
                }
            });
        },
        clickBtn(curItem) {
            this.curBtn = curItem.value;
            this.curObj = curItem;
            this.btnsList.forEach((item) => {
                this.$set(item, "active", item.label === curItem.label);
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.el-button {
    border: 0;
    color: var(--GRAY_FONT_COLOR);
    border-radius: 4px !important;
}
.el-button.active-btn {
    color: var(--PRIMARY_COLOR);
    background-color: #eff8ff;
}
.page-content {
    margin-top: 16px;
    ::v-deep .el-table {
        th.el-table__cell > .cell {
            font-weight: 400;
        }
        th.el-table__cell {
            background-color: #f5f6f8;
        }
        .el-table__cell {
            box-sizing: border-box;
            padding: 0;
            height: 40px;
            line-height: 39px;
        }
    }
}
</style>
