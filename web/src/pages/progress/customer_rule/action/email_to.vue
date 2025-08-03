<template>
    <div
        class="column-block"
        ref="ColumnBlock"
        :class="{ isEditing: isEditing }"
    >
        <el-popover
            ref="AddUserAndRole"
            placement="bottom-start"
            popper-class="add-user-and-role-popove"
            trigger="click"
            :width="350"
            @show="popoverShow"
            @after-leave="afterLeave"
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

                <!-- 任务角色 -->
                <div v-show="activeName === '任务角色'" class="role-content">
                    <el-checkbox-group
                        v-if="roleListShow.length"
                        v-model="checkArr"
                    >
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
                    <div v-if="!roleListShow.length" class="list-no-data">
                        暂无数据
                    </div>
                </div>
            </div>
            <!-- 标签化 -->
            <div class="detail" slot="reference">
                <!-- 名字列表 -->
                <div class="mem-list-content">
                    <div
                        ref="listCon"
                        class="tag-list"
                        :class="{ 'show-num': showNum }"
                    >
                        <div
                            class="tag-item"
                            v-for="(tagItem, tagIndex) in frontArr"
                            :key="tagIndex"
                        >
                            <el-avatar
                                v-show="tagItem.avatar !== 'role'"
                                class="progress-avatar"
                                icon="el-icon-user-solid"
                                size="small"
                                :src="getAvatar(tagItem.avatar)"
                                :style="
                                    getAvatarStack(tagItem.avatar, tagIndex)
                                "
                            ></el-avatar>
                            {{ tagItem.full_name }}
                        </div>
                    </div>
                    <!-- 数字气泡 -->
                    <el-tooltip
                        v-show="showNum && behandArr.length"
                        class="item"
                        effect="dark"
                        placement="top"
                    >
                        <div slot="content">
                            <div
                                v-for="(tagItem, tagIndex) in behandArr"
                                :key="tagIndex"
                                class="member-list"
                            >
                                <el-avatar
                                    v-show="tagItem.avatar !== 'role'"
                                    class="progress-avatar"
                                    icon="el-icon-user-solid"
                                    size="small"
                                    :src="getAvatar(tagItem.avatar)"
                                    :style="
                                        getAvatarStack(tagItem.avatar, tagIndex)
                                    "
                                ></el-avatar>
                                {{ tagItem.full_name }}
                            </div>
                        </div>
                        <b ref="num-box" class="num-box"
                            >+{{ behandArr.length }}</b
                        >
                    </el-tooltip>
                </div>
                <div class="default-text" v-if="!frontArr.length">
                    {{ placeholder }}
                </div>
                <b class="triangle"></b>
            </div>
        </el-popover>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import apiSetting from "@/common/api/module/progress_setting";
import { imgHost } from "@/assets/tool/const";
import { FieldType } from "@/assets/tool/const";
export default {
    props: {
        placeholder: {
            type: String,
            default: "请选择"
        },
        userList: {
            type: String,
            default: ""
        },
        userInfoList: {
            type: Array,
            default: () => []
        },
        roleCheckList: {
            type: String,
            default: ""
        }
    },
    data() {
        return {
            isEditing: false,
            activeName: "",
            tabsList: [
                {
                    label: "成员",
                    name: "成员"
                },
                {
                    label: "任务角色",
                    name: "任务角色"
                }
            ],
            searchInput: "", // 关键词
            // 成员的相关数据
            selectArr: [], // 已选择用户id信息
            peopleList: [], // 成员列表origin
            showList: [], // 成员列表show
            personList: [], // 已选择用户对象信息
            // 任务角色相关数据
            checkArr: [], // 已选择角色
            roleList: [], // 角色列表
            roleListShow: [], // 角色列表
            integratedArr: "",
            frontArr: [],
            behandArr: [],
            labelIndex: 0,
            showNum: true
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
            handler(str) {
                this.selectArr = str.split(",");
            },
            deep: true,
            immediate: true
        },
        selectArr: {
            handler(arr) {
                // 变形数据与checkArr整合回显
                this.inputShow();
                this.$emit("userlist-change", arr);
            },
            deep: true
        },
        checkArr: {
            handler(arr) {
                // 变形数据与selectArr整合回显
                this.inputShow();
                this.$emit("rolelist-change", arr);
            },
            deep: true
        },
        userInfoList: {
            handler(arr) {
                if (arr && arr.length) {
                    this.personList = _.cloneDeep(arr);
                    this.inputShow();
                }
            },
            deep: true,
            immediate: true
        },
        roleCheckList: {
            handler(str) {
                this.checkArr = str.split(",");
                this.inputShow();
            },
            deep: true
        }
    },
    mounted() {
        this.getPeopleList();
        this.getRuleRoleList();
    },
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
        popoverShow() {
            this.isEditing = true;
            this.popoverWidth = this.$refs.ColumnBlock.clientWidth;
            this.activeName = "成员";
            this.searchInput = "";
            this.$nextTick(() => {
                if (this.$refs.SearchInput) {
                    this.$refs.SearchInput.focus();
                }
            });
            // this.getPeopleList();
        },
        afterLeave() {
            this.isEditing = false;
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
                let numId = parseInt(id);
                if (_.find(this.showList, { id: numId })) {
                    _.remove(this.showList, { id: numId });
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
        },
        inputShow() {
            let roleArr = [];
            this.checkArr.forEach((field_key) => {
                let obj = {
                    full_name: _.find(this.roleList, { field_key: field_key })
                        .name,
                    avatar: "role"
                };
                roleArr.push(obj);
            });
            let integratedArr = [...this.personList, ...roleArr];
            this.integratedArr = integratedArr;
            this.frontArr = this.integratedArr;
            this.$nextTick(() => {
                this.getTagInit();
            });
        },

        getTagInit() {
            const listCon = this.$refs.listCon;
            if (listCon) {
                const labels = listCon.querySelectorAll(".tag-item");
                let labelIndex = 0; // 渲染到第几个
                const listConBottom = listCon.getBoundingClientRect().bottom; // 容器底部距视口顶部距离
                this.showNum = true;
                this.$nextTick(() => {
                    for (let i = 0; i < labels.length; i++) {
                        const _top = labels[i].getBoundingClientRect().top;
                        if (_top >= listConBottom) {
                            // 如果有标签顶部距离超过容器底部则表示超出容器隐藏
                            this.showNum = true;
                            labelIndex = i;
                            this.getShowLabel(labelIndex);
                            break;
                        } else {
                            this.getAllLabel();
                            this.showNum = false;
                        }
                    }
                });
            }
        },
        getShowLabel(labelIndex) {
            this.labelIndex = labelIndex;
            this.frontArr = this.getArrFront(this.integratedArr);
            this.behandArr = this.getArrBehand(this.integratedArr);
        },
        getAllLabel() {
            this.frontArr = this.integratedArr;
            this.behandArr = [];
        },
        getArrFront(arr) {
            let deepClone = _.cloneDeep(arr);
            let front = deepClone.splice(0, this.labelIndex);
            return front;
        },
        getArrBehand(arr) {
            let deepClone = _.cloneDeep(arr);
            let behand = deepClone.splice(this.labelIndex);
            return behand;
        }
    }
};
</script>

<style lang="scss" scoped>
@import "@/pages/progress/add_progress/style.scss";
.column-block {
    .detail {
        display: flex;
        height: 32px;
        padding: 0 10px;
        white-space: wrap;
        .mem-list-content {
            display: flex;
            // white-space: nowrap;
            .tag-list {
                display: inline-block;
                height: 32px;
                &.show-num {
                    // width: calc(100% - 32px);
                    // white-space: nowrap;
                }
            }
        }
    }
}
.default-text {
    color: #c0c4cc;
}

.tag-item {
    display: inline-block;
    margin-right: 8px;
    overflow: hidden;
    text-wrap: nowrap;
    text-overflow: ellipsis;
    max-width: 148px;
}

.avatar-box {
    display: inline-block;
    min-width: 24px;
    height: 24px;
    border-radius: 50%;
    background-size: 100% 100%;
    vertical-align: middle;
}
.num-box {
    box-sizing: border-box;
    display: inline-block;
    min-width: 22px;
    height: 22px;
    line-height: 22px;
    border-radius: 11px;
    padding: 0 2px;
    background-color: #f8f8f8;
    text-align: center;
    position: relative;
    top: 4px;
    vertical-align: middle;
    font-weight: 400;
    font-size: 12px;
    color: #2f384c;
}
.list-item {
    padding: 0 12px;
    height: 40px;
    line-height: 40px;
    &:hover {
        background-color: #f1f9ff;
    }
}
.popover-tag-item {
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
.member-list {
    margin: 6px 0;
    &:first-child {
        margin-top: 0;
    }
    &:last-child {
        margin-bottom: 0;
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
