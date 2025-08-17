<template>
    <div>
        <el-popover
            ref="AddUserAndRole"
            placement="bottom-start"
            popper-class="add-user-and-role-popove"
            trigger="click"
            :width="240"
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
                <div v-show="activeName === '成员'">
                    <!-- 所选tag -->
                    <div class="had-select-tag">
                        <!-- 遍历personList -->
                        <div
                            class="popover-tag-item"
                            v-for="(tagItem, tagIndex) in personList"
                            :key="tagIndex"
                        >
                            <el-avatar
                                class="progress-avatar"
                                icon="el-icon-user-solid"
                                size="small"
                                :src="getAvatar(tagItem.avatar)"
                                :style="getPopoverTagListAvatar(tagItem.avatar)"
                            ></el-avatar>
                            <div class="name-content">
                                {{ tagItem.full_name }}
                            </div>
                            <i
                                @click="removeTagItem(tagItem)"
                                class="el-tag__close el-icon-close tag-delete"
                            ></i>
                        </div>
                    </div>
                    <!-- 列表 -->
                    <div class="list-content">
                        <div
                            class="list-item multiple-people-list-item"
                            v-for="(listItem, listIndex) in showList"
                            :key="listIndex"
                            @click="confirmPeople(listItem)"
                        >
                            <el-avatar
                                class="progress-avatar"
                                icon="el-icon-user-solid"
                                size="small"
                                :src="getAvatar(listItem.avatar)"
                            ></el-avatar>
                            {{ listItem.full_name }}
                        </div>
                        <div v-show="!showList.length" class="list-no-data">
                            暂无数据
                        </div>
                    </div>
                </div>

                <!-- 任务属性 -->
                <div v-show="activeName === '任务属性'">
                    <el-checkbox-group v-model="checkArr">
                        <div
                            v-for="(roleItem, roleIndex) in roleListShow"
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
                                        getType(roleItem.mode, roleItem.expr)
                                    "
                                ></b>
                                {{ roleItem.name }}
                            </el-checkbox>
                        </div>
                    </el-checkbox-group>
                    <div v-show="!roleListShow.length" class="list-no-data">
                        暂无数据
                    </div>
                </div>
            </div>
            <div slot="reference" class="title-block">
                <span>可编辑字段的用户：</span>
                <span class="btn-block" :class="{ editing: editing }">
                    <b class="operation-box add"></b>
                    添加用户
                </span>
            </div>
        </el-popover>
        <div>
            <!-- 遍历出的列表 -->
            <!-- 所选角色 -->
            <div
                class="popover-tag-item"
                v-for="(tagItem, tagIndex) in personList"
                :key="tagIndex.id"
            >
                <el-avatar
                    class="progress-avatar"
                    icon="el-icon-user-solid"
                    size="small"
                    :src="getAvatar(tagItem.avatar)"
                    :style="getPopoverTagListAvatar(tagItem.avatar)"
                ></el-avatar>
                <div class="name-content">
                    {{ tagItem.full_name }}
                </div>
                <i
                    @click="removeTagItem(tagItem)"
                    class="el-tag__close el-icon-close tag-delete"
                ></i>
            </div>
            <!-- 所选用户 -->
            <div
                class="popover-tag-item"
                v-for="roleItem in checkArr"
                :key="roleItem"
            >
                <div class="role-content">
                    <!-- {{ roleItem }} -->
                    {{ getRoleName(roleItem) }}
                </div>
                <i
                    @click="removeRoleItem(roleItem)"
                    class="el-tag__close el-icon-close tag-delete"
                ></i>
            </div>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import apiSetting from "@/common/api/module/progress_setting";
import { imgHost } from "@/assets/tool/const";
import UserMessage from "@/components/user_message_tip";
import { FieldType } from "@/assets/tool/const";
export default {
    components: {
        UserMessage
    },
    props: {
        rowInfo: {
            type: Object,
            default: () => {}
        },
        userList: {
            type: Array,
            default: () => []
        },
        userInfoList: {
            type: Array,
            default: () => []
        },
        roleCheckList: {
            type: Array,
            default: () => []
        }
    },
    data() {
        return {
            activeName: "",
            tabsList: [
                {
                    label: "成员",
                    name: "成员"
                },
                {
                    label: "任务属性",
                    name: "任务属性"
                }
            ],
            editing: false,
            searchInput: "", // 关键词
            // 成员的相关数据
            selectArr: [], // 已选择用户id信息
            peopleList: [], // 成员列表origin
            showList: [], // 成员列表show
            personList: [], // 已选择用户对象信息
            // 任务属性相关数据
            checkArr: [], // 已选择角色
            roleList: [], // 角色列表
            roleListShow: [] // 角色列表
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
    watch: {
        searchInput: _.debounce(function () {
            if (this.activeName === "成员") {
                this.getPeopleList();
            } else {
                this.roleListShow = this.getOptionsList();
            }
        }, 500),
        userList: {
            handler(arr) {
                this.selectArr = arr;
            },
            deep: true,
            immediate: true
        },
        selectArr: {
            handler(arr) {
                this.$emit("userlist-change", arr);
            },
            deep: true,
            immediate: true
        },
        checkArr: {
            handler(arr) {
                this.$emit("rolelist-change", arr);
            },
            deep: true,
            immediate: true
        },
        userInfoList: {
            handler(arr) {
                this.personList = arr;
            },
            deep: true,
            immediate: true
        },
        roleCheckList: {
            handler(arr) {
                this.checkArr = _.cloneDeep(arr);
            },
            deep: true,
            immediate: true
        }
    },
    mounted() {
        this.getRuleRoleList();
    },
    updated() {},
    methods: {
        getOptionsList() {
            let tmpList = _.cloneDeep(this.roleList);
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
        tabChange() {
            this.searchInput = "";
            if (this.activeName === "成员") {
                this.getPeopleList();
            } else {
                this.roleListShow = this.getOptionsList();
            }
        },
        confirmPeople(people) {
            // 添加逻辑
            this.personList.push(people);
            this.selectArr.push(people.id);
            this.getPeopleList();
            // 移除列表刷新逻辑
            this.removeHadSelect();
        },
        removeTagItem(people) {
            // 移除页面绑定
            _.remove(this.personList, function (persion) {
                return persion.id === people.id;
            });
            // 移除提交参数
            _.remove(this.selectArr, function (id) {
                return id === people.id;
            });
            this.getPeopleList();
        },
        removeRoleItem(name) {
            let clone = _.cloneDeep(this.checkArr);
            this.checkArr = _.remove(clone, function (checkName) {
                return checkName !== name;
            });
        },
        popoverShow() {
            this.activeName = "成员";
            this.searchInput = "";
            this.getPeopleList();
            // this.getRuleRoleList();
        },
        getPeopleList() {
            let params = {
                ws_id: this.curSpace.id,
                ex: this.selectArr.join(","),
                key: this.searchInput,
                page_size: 50,
                page: 1
            };
            api.getPeopleList(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.peopleList = res.data;
                    this.showList = _.cloneDeep(this.peopleList);
                    this.removeHadSelect();
                } else {
                    this.peopleList = [];
                    this.showList = [];
                }
            });
        },
        getRuleRoleList() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id
                // id: this.rowInfo.id,
            };
            apiSetting.getRuleRoleList(params).then((res) => {
                if (res.resultCode === 200) {
                    this.roleList = res.data || [];
                    this.roleListShow = res.data || [];
                } else {
                    this.roleList = [];
                    this.roleListShow = [];
                }
            });
        },
        getRoleName(field) {
            if (_.find(this.roleList, { field_key: field })) {
                let name = _.find(this.roleList, { field_key: field }).name;
                return name;
            } else {
                return "已被删除";
            }
        },
        removeHadSelect() {
            this.selectArr.forEach((id) => {
                if (_.find(this.showList, { id: id })) {
                    _.remove(this.showList, { id: id });
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
            return {
                position: "relative",
                top: "-7px"
            };
        }
    }
};
</script>

<style lang="scss" scoped>
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
}
.title-block {
    margin-bottom: 6px;
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
    padding: 0 12px;
    height: 40px;
    line-height: 40px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    &:hover {
        background-color: #f1f9ff;
    }
}
.popover-tag-item {
    box-sizing: border-box;
    display: inline-block;
    padding: 0 6px 0 0;
    height: 20px;
    line-height: 20px;
    font-size: 12px;
    color: #82889e;
    background-color: #e2e7f0;
    margin: 4px;
    border-radius: 10px;
    max-width: calc(100% - 8px);
    .name-content {
        display: inline-block;
        max-width: calc(100% - 40px);
        padding: 0 4px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
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
    line-height: 12px;
    color: #ffffff;
    background-color: #a6abbc;
    border-radius: 50%;
    vertical-align: super;
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
</style>
<style lang="scss">
.el-popper.add-user-and-role-popove {
    box-sizing: border-box;
    margin-top: 0px;
    padding: 4px 0;
    .popper__arrow {
        display: none;
    }
}
</style>
