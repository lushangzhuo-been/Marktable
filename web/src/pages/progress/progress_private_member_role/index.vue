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
                :src="
                    require(`@/assets/image/field_type_icon/people_multiple.svg`)
                "
                alt=""
                width="20px"
                height="20px"
                class="icon"
            />
            角色&成员
        </div>
        <el-tabs class="basic-ui" v-model="activeName" @tab-click="handleClick">
            <el-tab-pane
                v-for="(item, index) in tabsList"
                :label="item.label"
                :name="item.name"
                :key="index"
            ></el-tab-pane>
        </el-tabs>
        <div class="setting-content">
            <div v-if="activeName === 'role'">
                <role ref="role"></role>
            </div>
            <div v-if="activeName === 'member'">
                <member ref="member"></member>
            </div>
        </div>
    </div>
</template>

<script>
import NoData from "@/pages/common/no_data";
import SpaceInfo from "@/pages/common/space_info";
import Role from "./role/index.vue";
import Member from "./member/index.vue";
export default {
    components: {
        Role,
        Member,
        NoData,
        SpaceInfo
    },
    computed: {
        progressTreeList() {
            return this.$store.state.progressTree;
        },
        isNoData() {
            return (
                (this.progressTreeList && this.progressTreeList.length) || false
            );
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        }
    },
    data() {
        return {
            isShowSpaceInfo: false,
            activeName: "role",
            tabsList: [
                {
                    label: "角色",
                    name: "role"
                },
                {
                    label: "成员",
                    name: "member"
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
                    if (this.activeName === "role") {
                        this.$nextTick(() => {
                            this.$refs.role.seachText = "";
                            this.$refs.role.fetchRoleList();
                        });
                    } else {
                        this.$nextTick(() => {
                            this.$refs.member.seachText = "";
                            this.$refs.member.fetchMemberList();
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
        handleClick() {},
        getId(arr, id) {
            arr.some((item) => {
                if (item.id === id) {
                    return true;
                } else if (item.children && item.children.length) {
                    this.getId(item.children, id);
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
::v-deep .el-tabs__item {
    height: 36px;
    line-height: 24px;
}
</style>
