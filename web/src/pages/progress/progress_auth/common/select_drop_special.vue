<template>
    <div>
        <!-- <div class="desc">{{ curItem.desc }}</div> -->
        <div>
            <div class="note">以下成员拥有此操作权限</div>
            <!-- 遍历出的列表 -->
            <!-- 所选角色 -->
            <div v-if="hasAuth" class="has-auth">
                <div v-if="checkedRoleArr.length" class="auth-line-style">
                    <span class="label">空间角色:</span>
                    <div class="value-arr">
                        <div
                            class="popover-tag-item"
                            v-for="item in checkedRoleArr"
                            :key="item.value"
                        >
                            <div class="role-content">
                                {{ item.label }}
                            </div>
                            <i
                                @click="removeTagItem(item, '空间角色')"
                                class="el-tag__close el-icon-close tag-delete"
                            ></i>
                        </div>
                    </div>
                </div>
                <!-- 所选任务属性 -->
                <div v-if="checkedTaskArr.length" class="auth-line-style">
                    <span class="label">任务属性:</span>
                    <div class="value-arr">
                        <div
                            class="popover-tag-item"
                            v-for="(tagItem, tagIndex) in checkedTaskArr"
                            :key="tagIndex.id"
                        >
                            <div class="role-content">
                                {{ tagItem.name }}
                            </div>
                            <i
                                @click="removeTagItem(tagItem, '任务属性')"
                                class="el-tag__close el-icon-close tag-delete"
                            ></i>
                        </div>
                    </div>
                </div>
                <!-- 所选用户 -->
                <div v-if="selectedPersonArr.length" class="auth-line-style">
                    <span class="label">成员:</span>
                    <div class="value-arr">
                        <div
                            class="popover-tag-item"
                            v-for="roleItem in selectedPersonArr"
                            :key="roleItem.id"
                        >
                            <el-avatar
                                icon="el-icon-user-solid"
                                size="small"
                                :src="getAvatar(roleItem.avatar)"
                                :style="
                                    getPopoverTagListAvatar(roleItem.avatar)
                                "
                            ></el-avatar>
                            <div class="role-content">
                                {{ roleItem.full_name }}
                            </div>
                            <i
                                @click="removePepoleItem(roleItem)"
                                class="el-tag__close el-icon-close tag-delete"
                            ></i>
                        </div>
                    </div>
                </div>
            </div>
            <div v-else class="no-auth">暂无操作权限</div>
        </div>
        <el-popover
            ref="authPopper"
            placement="bottom-start"
            popper-class="add-user-and-role-popove"
            trigger="click"
            :width="262"
            @show="popoverShow"
        >
            <!-- 遍历item.list  做下拉内容 -->
            <div>
                <!-- 搜索框 -->
                <div class="search-input">
                    <el-input
                        ref="SearchInput"
                        class="basic-ui progress-search-input"
                        v-model="searchInput"
                        size="small"
                        placeholder="搜索"
                        prefix-icon="el-icon-search"
                    ></el-input>
                </div>
                <!-- 成员 角色 tab -->
                <el-tabs
                    class="basic-ui font add-user-and-role"
                    size="small"
                    v-model="activeName"
                    @tab-click="tabChange"
                >
                    <el-tab-pane
                        v-for="(item, index) in tabsList"
                        :label="item.label"
                        :name="item.name"
                        :key="index"
                    ></el-tab-pane>
                </el-tabs>
                <!-- 用户 -->
                <div v-if="activeName === '成员'">
                    <!-- 所选tag -->
                    <div class="had-select-tag">
                        <!-- 遍历选中的人员-->
                        <div
                            class="popover-tag-item"
                            v-for="(tagItem, tagIndex) in selectedPersonArr"
                            :key="tagIndex"
                        >
                            <el-avatar
                                icon="el-icon-user-solid"
                                size="small"
                                :src="getAvatar(tagItem.avatar)"
                                :style="getPopoverTagListAvatar(tagItem.avatar)"
                            ></el-avatar>
                            <div class="role-content">
                                {{ tagItem.full_name }}
                            </div>
                            <i
                                @click="removePepoleItem(tagItem)"
                                class="el-tag__close el-icon-close tag-delete"
                            ></i>
                        </div>
                    </div>
                    <!-- 列表 -->
                    <div class="list-content">
                        <template
                            v-if="showPersonList && showPersonList.length"
                        >
                            <div
                                class="list-item multiple-people-list-item"
                                v-for="(listItem, listIndex) in showPersonList"
                                :key="listIndex"
                                @click="selectPerson(listItem)"
                            >
                                <el-avatar
                                    icon="el-icon-user-solid"
                                    size="small"
                                    :src="getAvatar(listItem.avatar)"
                                ></el-avatar>
                                <span v-overflow class="full-name">
                                    {{ listItem.full_name }}</span
                                >
                            </div>
                        </template>
                        <div v-else class="list-no-data">暂无数据</div>
                    </div>
                </div>

                <!-- 任务属性 -->
                <div class="list-content top4" v-if="activeName === '任务属性'">
                    <template
                        v-if="specialPersonList && specialPersonList.length"
                    >
                        <el-checkbox-group
                            v-model="checkedTask"
                            @change="checkedTaskChange"
                        >
                            <div
                                v-for="(
                                    roleItem, roleIndex
                                ) in specialPersonList"
                                :key="roleIndex"
                                class="role-list-item"
                            >
                                <el-checkbox
                                    :label="roleItem.field_key"
                                    class="role-checkbox"
                                >
                                    <b
                                        class="type-box"
                                        :style="
                                            getType(
                                                roleItem.mode,
                                                roleItem.expr
                                            )
                                        "
                                    ></b>
                                    {{ roleItem.name }}
                                </el-checkbox>
                            </div>
                        </el-checkbox-group>
                    </template>
                    <div v-else class="list-no-data">暂无数据</div>
                </div>
                <!-- 空间角色 -->
                <div class="list-content top4" v-if="activeName === '空间角色'">
                    <template v-if="spaceRoleList && spaceRoleList.length">
                        <el-checkbox-group
                            v-model="checkedRole"
                            @change="checkedRoleChange"
                        >
                            <div
                                v-for="item in spaceRoleList"
                                :key="item.value"
                                class="role-list-item"
                            >
                                <el-checkbox
                                    :label="item.value"
                                    class="role-checkbox"
                                >
                                    {{ item.label }}
                                </el-checkbox>
                            </div>
                        </el-checkbox-group>
                    </template>
                    <div v-else class="list-no-data">暂无数据</div>
                </div>
            </div>
            <div slot="reference">
                <el-input
                    ref="SearchInput"
                    style="width: 420px"
                    size="small"
                    :placeholder="placeholder"
                    suffix-icon="el-icon-caret-bottom"
                ></el-input>
            </div>
        </el-popover>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import apiSetting from "@/common/api/module/progress_setting";
import { mapState } from "vuex";
import { imgHost } from "@/assets/tool/const";
import UserMessage from "@/components/user_message_tip";
import { FieldType } from "@/assets/tool/const";
export default {
    components: {
        UserMessage
    },
    props: {
        curItem: {
            type: Object,
            default: () => {}
        }
        // userList: {
        //     type: Array,
        //     default: () => []
        // }
        // userInfoList: {
        //     type: Array,
        //     default: () => []
        // },
        // roleCheckList: {
        //     type: Array,
        //     default: () => []
        // }
    },
    data() {
        return {
            placeholder: "请选择空间角色、任务属性、成员",
            // 人员
            selectedPersonArr: [], // 选中的人员
            showPersonList: [], // 待选择的人员列表
            activeName: "空间角色",
            tabsList: [
                // {
                //     label: "空间角色",
                //     name: "空间角色"
                // },
                // {
                //     label: "任务属性",
                //     name: "任务属性"
                // },
                // {
                //     label: "成员",
                //     name: "成员"
                // }
            ],
            editing: false,
            searchInput: "", // 关键词
            // 成员的相关数据
            // selectedPersonArr: [], // 已选择用户id信息
            peopleList: [], // 成员列表备份
            showList: [], // 成员列表show
            personList: [], // 已选择用户对象信息
            // 任务属性相关数据
            checkedRole: [], // 下拉绑定的已选择的空间角色
            checkedRoleArr: [], // 获取的已选择的空间角色
            checkedTask: [],
            checkedTaskArr: [], // 已选择的控件属性
            tmpRoleList: [], // 角色列表
            specialPersonList: [], // 任务属性列表
            spaceRoleList: [],
            tmpSpaceRoleList: []
        };
    },
    computed: {
        roleList() {
            return this.$store.state.progress_setting.roleList;
        },
        hasAuth() {
            return (
                this.checkedRoleArr.length ||
                this.checkedTaskArr.length ||
                this.selectedPersonArr.length
            );
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        }
    },
    watch: {
        // $route() {
        //     this.activeName = "成员";
        //     this.searchInput = "";
        //     if (this.curItem && Object.keys(this.curItem).length) {
        //         this.switchAuth();
        //         this.getSpecialPersonList(); // 获取任务属性下拉
        //     }
        // },
        roleList: {
            handler(newVal) {
                if (!newVal || !newVal.length) {
                    this.tmpSpaceRoleList = [];
                    this.spaceRoleList = [];
                    return;
                }
                this.spaceRoleList = _.cloneDeep(newVal);
                this.tmpSpaceRoleList = _.cloneDeep(newVal);
                // 安全操作
            },
            immediate: true
        },
        searchInput: _.debounce(function () {
            if (this.activeName === "成员") {
                this.getShowPeopleList();
            } else if (this.activeName === "任务属性") {
                this.specialPersonList = this.getTmpSpecialPersonList();
            } else if (this.activeName === "空间角色") {
                this.spaceRoleList = this.getTmpSpaceRoleList();
            }
        }, 500),
        curItem: {
            handler(item) {
                this.activeName = "空间角色";
                this.searchInput = "";
                this.tabsList = [];
                if (item && Object.keys(item).length) {
                    if (
                        item.value === "delete" ||
                        item.value === "edit" ||
                        item.value === "file" ||
                        item.value === "sub_tmpl" ||
                        item.value === "progress"
                    ) {
                        this.placeholder = "请选择空间角色、任务属性、成员";
                        this.tabsList = [
                            {
                                label: "空间角色",
                                name: "空间角色"
                            },
                            {
                                label: "任务属性",
                                name: "任务属性"
                            },
                            {
                                label: "成员",
                                name: "成员"
                            }
                        ];
                    } else {
                        this.placeholder = "请选择空间角色、成员";
                        this.tabsList = [
                            {
                                label: "空间角色",
                                name: "空间角色"
                            },
                            {
                                label: "成员",
                                name: "成员"
                            }
                        ];
                    }
                    this.switchAuth();
                    this.getSpecialPersonList(); // 获取任务属性下拉， 每个流程都不一样，
                }
            },
            deep: true,
            immediate: true
        }
    },
    mounted() {
        this.getSpecialPersonList();
    },
    updated() {},
    methods: {
        switchAuth() {
            // 切换左边的权限重新获取现有的权限
            this.fetchCurAuthDetail();
        },
        // 获取当前权限
        fetchCurAuthDetail() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                mode: this.curItem.value
            };
            apiSetting.fetchAuthDetail(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    // 获取详情 && 下拉回填
                    if (res.data.ws_roles && res.data.ws_roles.length) {
                        // 空间角色
                        this.checkedRoleArr = this.spaceRoleList.filter(
                            (item) =>
                                res.data.ws_roles
                                    .split(",")
                                    .find((sub) => sub === item.value)
                        );
                        this.checkedRole = res.data.ws_roles.split(",");
                    } else {
                        this.checkedRoleArr = [];
                        this.checkedRole = [];
                    }

                    if (
                        res.data.issue_role_list &&
                        res.data.issue_role_list.length
                    ) {
                        //  任务属性
                        this.checkedTaskArr = this.specialPersonList.filter(
                            (item) =>
                                res.data.issue_role_list
                                    .split(",")
                                    .find((sub) => sub === item.field_key)
                        );
                        this.checkedTask = res.data.issue_role_list.split(",");
                    } else {
                        this.checkedTaskArr = [];
                        this.checkedTask = [];
                    }

                    if (
                        res.data.user_list_info &&
                        res.data.user_list_info.length
                    ) {
                        //  人员
                        this.selectedPersonArrTmp = res.data.user_list_info; // 暂存
                        this.selectedPersonArr = _.cloneDeep(
                            this.selectedPersonArrTmp
                        ); // 显示的选中的
                    } else {
                        this.selectedPersonArr = [];
                        this.selectedPersonArrTmp = [];
                    }
                } else {
                    // 空间角色
                    this.checkedRoleArr = [];
                    this.checkedRole = [];
                    //  任务属性
                    this.checkedTaskArr = [];
                    this.checkedTask = [];
                    this.selectedPersonArr = [];
                    this.selectedPersonArrTmp = [];
                }
                this.getShowPeopleList(); // 重新获取人员下拉
                this.$nextTick(() => {
                    this.$refs.authPopper.updatePopper();
                });
            });
        },
        /** 人员处理 start */
        // 获取人员下拉
        getShowPeopleList() {
            let params = {
                ws_id: this.curSpace.id,
                ex: this.selectedPersonArr.map((item) => item.id).join(","),
                key: this.searchInput,
                page_size: 50,
                page: 1
            };
            api.getPeopleList(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    res.data.length
                ) {
                    this.showPersonList = this.resArr(
                        res.data,
                        this.selectedPersonArr
                    );
                } else {
                    this.peopleList = [];
                    this.showPersonList = [];
                }
            });
        },

        // 从list中去掉已经选中的，并作为下拉
        resArr(arr1, arr2) {
            return arr1.filter((v) => arr2.every((val) => val.id != v.id));
        },
        // 从人员下拉中选择
        selectPerson(people) {
            this.selectedPersonArrTmp.push(people);
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                mode: this.curItem.value,
                ws_roles: this.checkedRole.join(","),
                user_list: this.selectedPersonArrTmp
                    .map((item) => item.id)
                    .join(","),
                issue_role_list: this.checkedTask.join(",")
            };
            // 选中后发送接口，获取选中的数据
            apiSetting.updateSelectedAuth(params).then((res) => {
                // 选中人员后，传递给后端成功后，添加到展示框
                if (res && res.resultCode === 200) {
                    this.selectedPersonArr = _.cloneDeep(
                        this.selectedPersonArrTmp
                    );
                    this.getShowPeopleList();
                } else {
                    this.$message({
                        showClose: true,
                        message: "添加人员失败",
                        type: "warning"
                    });
                }
                this.fetchCurAuthDetail();
            });
        },
        // 移除已经选中的人员
        removePepoleItem(people) {
            this.selectedPersonArrTmp = this.selectedPersonArr.filter(
                (item) => item.id !== people.id
            );
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                mode: this.curItem.value,
                ws_roles: this.checkedRole.join(","),
                user_list: this.selectedPersonArrTmp
                    .map((item) => item.id)
                    .join(","),
                issue_role_list: this.checkedTask.join(",")
            };
            // 选中后发送接口，获取选中的数据
            apiSetting.updateSelectedAuth(params).then((res) => {
                // 选中人员后，传递给后端成功后，添加到展示框
                if (res && res.resultCode === 200) {
                    this.selectedPersonArr = _.cloneDeep(
                        this.selectedPersonArrTmp
                    );
                    this.getShowPeopleList();
                } else {
                    this.$message({
                        showClose: true,
                        message: "添加人员失败",
                        type: "warning"
                    });
                }
                this.fetchCurAuthDetail();
            });
        },
        /** 人员处理 end */
        // 选中的空间角色变化
        checkedRoleChange() {
            this.updateInfo();
        },
        // 选中的任务属性变化
        checkedTaskChange(val) {
            this.updateInfo();
        },
        initUpdateParams() {
            return {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id,
                mode: this.curItem.value,
                ws_roles: this.checkedRole.join(","),
                user_list: this.selectedPersonArrTmp
                    .map((item) => item.id)
                    .join(","),
                issue_role_list: this.checkedTask.join(",")
            };
        },
        // 更新选中的结果并刷新数据
        updateInfo() {
            let params = this.initUpdateParams();
            // 选中后发送接口，获取选中的数据
            apiSetting.updateSelectedAuth(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.fetchCurAuthDetail();
                }
            });
        },
        getTmpSpecialPersonList() {
            let tmpList = _.cloneDeep(this.tmpRoleList);
            if (this.searchInput.trim()) {
                return _.cloneDeep(
                    tmpList.filter((item) => {
                        return (
                            item.name
                                .toLocaleUpperCase()
                                .indexOf(
                                    this.searchInput.trim().toLocaleUpperCase()
                                ) > -1
                        );
                    })
                );
            } else {
                return tmpList;
            }
        },
        getTmpSpaceRoleList() {
            let tmpList = _.cloneDeep(this.tmpSpaceRoleList);
            if (this.searchInput.trim()) {
                return _.cloneDeep(
                    tmpList.filter((item) => {
                        return (
                            item.label
                                .toLocaleUpperCase()
                                .indexOf(
                                    this.searchInput.trim().toLocaleUpperCase()
                                ) > -1
                        );
                    })
                );
            } else {
                return tmpList;
            }
        },
        tabChange() {
            this.searchInput = "";

            if (this.activeName === "成员") {
                this.getShowPeopleList();
            } else if (this.activeName === "任务属性") {
                this.specialPersonList = this.getTmpSpecialPersonList();
            } else if (this.activeName === "空间角色") {
                this.spaceRoleList = this.getTmpSpaceRoleList();
            }
        },
        // 移除人物属性和角色
        removeTagItem(item, type) {
            if (type === "任务属性") {
                let tmp = this.checkedTaskArr.filter(
                    (sub) => sub.field_key !== item.field_key
                );
                let params = this.initUpdateParams();
                params.issue_role_list = tmp
                    .map((item) => item.field_key)
                    .join(",");
                // 选中后发送接口，获取选中的数据
                apiSetting.updateSelectedAuth(params).then((res) => {
                    if (res && res.resultCode === 200) {
                        this.checkedTaskArr = _.cloneDeep(tmp);
                        this.checkedTask = tmp
                            .map((item) => item.field_key)
                            .join(",");
                        this.fetchCurAuthDetail();
                    } else {
                        this.$message({
                            showClose: true,
                            message: "移除任务属性失败",
                            type: "warning"
                        });
                    }
                });
            } else if (type === "空间角色") {
                let tmp = this.checkedRoleArr.filter(
                    (sub) => sub.value !== item.value
                );
                let params = this.initUpdateParams();
                params.ws_roles = tmp.map((item) => item.value).join(",");
                // 选中后发送接口，获取选中的数据
                apiSetting.updateSelectedAuth(params).then((res) => {
                    if (res && res.resultCode === 200) {
                        this.checkedRoleArr = _.cloneDeep(tmp);
                        this.checkedRole = tmp
                            .map((item) => item.value)
                            .join(",");
                        this.fetchCurAuthDetail();
                    } else {
                        this.$message({
                            showClose: true,
                            message: "移除空间角色失败",
                            type: "warning"
                        });
                    }
                });
            }
        },
        // 打开popover选择权限
        popoverShow() {
            // this.activeName = "空间角色";
            this.searchInput = "";
            this.getShowPeopleList(); // 获取人员接口
            // this.getSpecialPersonList(); // 获取任务属性接口(特殊配置的人员)
        },
        // 获取任务属性
        getSpecialPersonList() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id
            };
            apiSetting.getRuleRoleList(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.tmpRoleList = res.data || [];
                    this.specialPersonList = res.data || [];
                } else {
                    this.tmpRoleList = [];
                    this.specialPersonList = [];
                }
            });
        },

        getType(type, expr) {
            if (type) {
                return {
                    "background-image": `url(${FieldType[type + "_" + expr]})`
                };
            }
        },
        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        },
        getAvatarStack() {
            return {
                position: "relative",
                top: "-2px"
            };
        },
        getPopoverTagListAvatar() {
            return {};
        }
    }
};
</script>

<style lang="scss" scoped>
.list-no-data {
    height: 40px;
    line-height: 40px;
    text-align: center;
    color: #82889e;
}
::v-deep.basic-ui.el-tabs.el-tabs--top {
    margin-top: 2px;
    background-color: #f5f5f5;
    .el-tabs__header {
        margin-bottom: 0;
        .el-tabs__nav-scroll {
            padding: 0 16px;
            .el-tabs__item {
                color: #82889e;
                font-weight: 400;
                &.is-active {
                    color: #1890ff;
                }
            }
        }
    }
}
.list-content {
    max-height: 320px;
    overflow-y: auto;
    &.top4 {
        margin-top: 4px;
    }
}
.btn-block {
    color: #1890ff;
    cursor: pointer;
    .operation-box {
        display: inline-block;
        width: 18px;
        height: 18px;
        background-size: 100% 100%;
        vertical-align: text-bottom;
        position: relative;
        top: -1px;
        &.add {
            background-image: url(@/assets/image/progress/add.png);
        }
    }
}
.list-item {
    box-sizing: border-box;
    display: flex;
    align-items: center;
    padding: 0 12px;
    height: 40px;

    &:hover {
        background-color: #f1f9ff;
    }
    .full-name {
        display: inline-block;
        width: calc(100% - 32px);
        margin-left: 4px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }
}
.note {
    font-weight: 700;
    margin: 16px 0 20px;
}
.auth-line-style {
    display: flex;
    align-items: flex-start;
    margin-bottom: 12px;
    &:last-child {
        margin-bottom: 0;
    }
    .label {
        display: inline-block;
        width: 64px;
        /* height: 32px; */
        /* line-height: 32px; */
        margin-right: 10px;
        text-align: right;
        color: #82889e;
    }
    .value-arr {
        display: inline-block;
        width: calc(100% - 74px);
    }
}
.popover-tag-item {
    box-sizing: border-box;
    display: inline-block;
    /* align-items: center; */
    padding: 0 6px 0 0;
    height: 20px;
    line-height: 20px;
    font-size: 12px;
    color: #82889e;
    background-color: #e2e7f0;
    margin: 0 4px;
    border-radius: 10px;
    /* max-width: calc(100% - 8px); */
    ::v-deep .el-avatar--small {
        // position: relative;
        // top: -7px;
        width: 20px;
        height: 20px;
        line-height: 20px;
    }

    .role-content {
        display: inline-block;
        max-width: calc(100% - 20px);
        padding: 0 4px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }
}
.had-select-tag {
    padding: 4px;
}
.el-tag__close.el-icon-close.tag-delete {
    width: 12px;
    height: 12px;
    text-align: center;
    line-height: 14px;
    color: #ffffff;
    background-color: #a6abbc;
    border-radius: 50%;
    vertical-align: super;
    cursor: pointer;
    &:hover {
        background-color: #82889e;
    }
}
::v-deep.role-list-item {
    height: 40px;
    line-height: 40px;
    padding: 0 16px;
    .role-checkbox {
        display: flex;
        flex-direction: row-reverse;
        justify-content: space-between;
        height: 100%;
        .el-checkbox__input {
            line-height: inherit;
        }
        .el-checkbox__label {
            line-height: 40px;
            padding: 0;
        }
    }
    &:hover {
        background-color: #f5faff;
    }
    &.check-all {
        border-bottom: 1px solid #f0f1f5;
    }
}
.type-box {
    display: inline-block;
    width: 20px;
    height: 20px;
    background-size: 100% 100%;
    vertical-align: middle;
    position: relative;
    top: -2px;
}
::v-deep .el-input--small.progress-search-input .el-input__inner {
    border: none;
}
.has-auth {
    padding: 16px;
    background-color: #f5f6f8;
    margin-bottom: 16px;
    border-radius: 4px;
}
.no-auth {
    box-sizing: border-box;
    height: 100px;
    line-height: 100px;
    text-align: center;
    color: #a6abbc;
    background-color: #f5f6f8;
    margin-bottom: 16px;
    border-radius: 4px;
}
</style>

<style lang="scss">
.el-popper.add-user-and-role-popove {
    box-sizing: border-box;
    margin-top: 2px;
    padding: 4px 0;
    .popper__arrow {
        display: none;
    }
}
</style>
