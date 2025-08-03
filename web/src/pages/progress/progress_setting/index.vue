<template>
    <!-- <div v-if="!isNoData" style="height: 100%">
        <no-data :isTextShow="true"></no-data>
    </div>
    <div v-else-if="isNoData && isShowSpaceInfo" style="height: 100%">
        <space-info :spaceInfo="curSpace"></space-info>
    </div> -->
    <div>
        <div class="sub-title">
            <img
                :src="require(`@/assets/image/common/progress_setting.png`)"
                alt=""
                width="20px"
                height="20px"
                class="icon"
            />
            流程设置
        </div>
        <el-tabs
            class="basic-ui font"
            size="small"
            v-model="activeName"
            @tab-click="handleClick"
        >
            <el-tab-pane
                v-for="(item, index) in tabsList"
                :label="item.label"
                :name="item.name"
                :key="index"
            ></el-tab-pane>
        </el-tabs>
        <div class="setting-content">
            <div v-if="activeName === 'filed'">
                <filed ref="filed"></filed>
            </div>
            <div v-if="activeName === 'page'">
                <page ref="page"></page>
            </div>
            <div v-if="activeName === 'node'">
                <!-- 旧流程 -->
                <!-- <node ref="node"></node> -->
                <!-- 新流程 -->
                <node-new ref="node"></node-new>
            </div>
        </div>
    </div>
</template>

<script>
import SpaceInfo from "@/pages/common/space_info";
import NoData from "@/pages/common/no_data";
import Filed from "./filed/index.vue";
import Page from "./page/index.vue";
// import Page from "./page_old/index.vue";
import Node from "./node/index.vue";
import NodeNew from "./node_new/index_new.vue";
export default {
    components: {
        Filed,
        Page,
        Node,
        NodeNew,
        NoData,
        SpaceInfo
    },
    computed: {
        progressTreeList() {
            return this.$store.state.progressTree;
        },
        // isNoData() {
        //     return (
        //         (this.progressTreeList && this.progressTreeList.length) || false
        //     )
        // },
        curSpace() {
            return this.$store.state.curSpace || {};
        }
    },
    data() {
        return {
            isShowSpaceInfo: false,
            activeName: "filed",
            tabsList: [
                {
                    label: "字段",
                    name: "filed"
                },
                {
                    label: "界面",
                    name: "page"
                },
                {
                    label: "流程",
                    name: "node"
                }
            ]
        };
    },
    watch: {
        progressTreeList(arr) {
            if (arr && arr.length) {
                if (
                    arr.some(
                        (item) => item.id === Number(this.$route.params.id)
                    )
                ) {
                    if (this.activeName === "filed") {
                        this.$nextTick(() => {
                            this.$refs.filed.fetchFiledConfig();
                            this.$refs.filed.fetchFiledList();
                        });
                    } else if (this.activeName === "page") {
                        this.$nextTick(() => {
                            this.$refs.page.fetchPageConfig();
                        });
                    } else {
                        this.$nextTick(() => {
                            this.$refs.node.fetchNodeConfig();
                        });
                    }
                    this.isShowSpaceInfo = false;
                } else {
                    this.isShowSpaceInfo = true;
                }
            }
        },
        $route() {
            this.isShowSpaceInfo = false;
        }
    },
    methods: {
        handleClick() {}
    }
};
</script>

<style lang="scss" scoped>
::v-deep .el-tabs__item {
    height: 36px;
    line-height: 24px;
}
</style>
