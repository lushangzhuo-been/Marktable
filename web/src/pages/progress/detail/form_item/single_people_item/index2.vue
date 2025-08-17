<template>
    <div
        class="column-block form-width"
        ref="ColumnBlock"
        :class="{
            isEditing: isEditing,
            'editing-effect': formItem.can_modify === 'yes',
            'col-required': formItem.required === 'yes'
        }"
    >
        <div class="detail" v-if="!isEditing" @click="checkScope">
            <!-- 名字列表 -->
            <template v-if="selectArr && selectArr.length">
                <div
                    ref="listCon"
                    class="tag-list"
                    :class="{ 'show-num': showNum }"
                >
                    <div
                        class="tag-item"
                        v-for="(tagItem, tagIndex) in frontArr"
                        :key="tagIndex"
                        v-show="frontArr && frontArr.length"
                    >
                        <user-message :userMessage="tagItem">
                            <span>
                                <el-avatar
                                    class="progress-avatar"
                                    icon="el-icon-user-solid"
                                    size="small"
                                    :src="getAvatar(tagItem.avatar)"
                                    :style="
                                        getAvatarStack(tagItem.avatar, tagIndex)
                                    "
                                ></el-avatar>
                                {{ tagItem.full_name }}
                            </span>
                        </user-message>
                    </div>
                </div>
                <!-- 数字气泡 -->
                <el-tooltip
                    v-show="showNum"
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
                            <user-message
                                :userMessage="tagItem"
                                position="left"
                            >
                                <span>
                                    <el-avatar
                                        class="progress-avatar"
                                        icon="el-icon-user-solid"
                                        size="small"
                                        :src="getAvatar(tagItem.avatar)"
                                        :style="
                                            getAvatarStack(
                                                tagItem.avatar,
                                                tagIndex
                                            )
                                        "
                                    ></el-avatar>
                                    {{ tagItem.full_name }}
                                </span>
                            </user-message>
                        </div>
                    </div>
                    <b class="num-box" ref="num-box">+{{ behandArr.length }}</b>
                </el-tooltip>
            </template>
            <div v-else>
                {{ emptySpace() }}
            </div>
        </div>
        <div v-if="isEditing">
            <el-popover
                ref="DropPopover"
                :width="popoverWidth"
                popper-class="progress-popover"
                placement="bottom-start"
                trigger="click"
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
                                :style="getPopoverTagListAvatar()"
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
                                :style="getAvatar(listItem.avatar)"
                            ></el-avatar>
                            {{ listItem.full_name }}
                        </div>
                        <div v-show="!showList.length" class="list-no-data">
                            暂无数据
                        </div>
                    </div>
                </div>
                <div class="detail" slot="reference">
                    <!-- 名字列表 -->
                    <template v-if="selectArr && selectArr.length">
                        <div
                            ref="listCon"
                            class="tag-list"
                            :class="{ 'show-num': showNum }"
                        >
                            <div
                                class="tag-item"
                                v-show="frontArr && frontArr.length"
                                v-for="(tagItem, tagIndex) in frontArr"
                                :key="tagIndex"
                            >
                                <el-avatar
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
                            v-show="showNum"
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
                                    <user-message
                                        :userMessage="tagItem"
                                        position="left"
                                    >
                                        <span>
                                            <el-avatar
                                                class="progress-avatar"
                                                icon="el-icon-user-solid"
                                                size="small"
                                                :src="getAvatar(tagItem.avatar)"
                                                :style="
                                                    getAvatarStack(
                                                        tagItem.avatar,
                                                        tagIndex
                                                    )
                                                "
                                            ></el-avatar>
                                            {{ tagItem.full_name }}
                                        </span>
                                    </user-message>
                                </div>
                            </div>
                            <b class="num-box" ref="num-box"
                                >+{{ behandArr.length }}</b
                            >
                        </el-tooltip>
                    </template>
                    <div v-if="!selectArr.length"></div>
                </div>
            </el-popover>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import { imgHost } from "@/assets/tool/const";
import { emptySpace } from "@/assets/tool/func";
import UserMessage from "@/components/user_message_tip";
export default {
    components: {
        UserMessage
    },
    props: {
        formItem: {
            type: Object,
            default: () => {}
        },
        formData: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            isEditing: false,
            selectArr: [],
            frontArr: [],
            behandArr: [],
            popoverWidth: 220,
            searchInput: "",
            personList: [],
            showList: [],
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
        formData: {
            handler(formData) {
                this.personList = [];
                // 有值，只取第一个
                let selectArr = []; // 双向绑定id数组
                if (
                    formData[this.formItem.field_key] &&
                    formData[this.formItem.field_key].length
                ) {
                    if (formData[this.formItem.field_key]) {
                        this.personList = _.cloneDeep(
                            formData[this.formItem.field_key]
                        );
                        this.personList.forEach((objItem) => {
                            selectArr.push(objItem.id);
                        });
                    } else {
                        this.personList = [];
                    }
                }
                this.selectArr = selectArr;
                // this.frontArr = this.getArrFront(this.personList);
                // this.behandArr = this.getArrBehand(this.personList);
                this.frontArr = this.personList;
                this.$nextTick(() => {
                    this.getTagInit();
                });
            },
            deep: true
        },
        searchInput: _.debounce(function () {
            this.getPeopleList();
        }, 500)
    },
    mounted() {},
    methods: {
        getShowLabel(labelIndex) {
            this.labelIndex = labelIndex;
            this.frontArr = this.getArrFront(this.personList);
            this.behandArr = this.getArrBehand(this.personList);
        },
        getAllLabel() {
            this.frontArr = this.personList;
            this.behandArr = [];
        },
        getTagInit() {
            const listCon = this.$refs.listCon;
            if (listCon) {
                const labels = listCon.querySelectorAll(".tag-item");
                let labelIndex = 0; // 渲染到第几个
                const listConBottom = listCon.getBoundingClientRect().bottom; // 容器底部距视口顶部距离
                this.showNum = false;
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
            }
        },
        emptySpace(param) {
            return emptySpace(param);
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
        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        },
        getAvatarStack(src, index) {
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
        getArrFront(arr) {
            let deepClone = _.cloneDeep(arr);
            let front = deepClone.splice(0, this.labelIndex);
            return front;
        },
        getArrBehand(arr) {
            let deepClone = _.cloneDeep(arr);
            let behand = deepClone.splice(this.labelIndex);
            return behand;
        },
        checkScope() {
            this.fetAuthEdit();
        },
        fetAuthEdit() {
            // 获取进展权限
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: this.formData._id,
                auth_mode: "edit",
                field_key: this.formItem.field_key
            };
            api.getUserAuth(params).then((res) => {
                if (res && res.resultCode === 200) {
                    if (res.data) {
                        this.isEditing = !this.isEditing;
                        this.$set(this.formItem, "isEditing", this.isEditing);
                        if (this.isEditing) {
                            // this.popoverWidth = this.$refs.ColumnBlock.clientWidth;
                            this.popoverWidth = 220;
                            setTimeout(() => {
                                this.$refs.DropPopover.doShow();
                                this.$nextTick(() => {
                                    if (this.$refs.SearchInput) {
                                        this.$refs.SearchInput.focus();
                                    }
                                });
                                this.getPeopleList();
                            }, 20);
                        } else {
                            this.afterLeave();
                        }
                    } else {
                        this.isEditing = false;
                    }
                } else {
                    this.isEditing = false;
                }
            });
        },
        afterLeave() {
            this.isEditing = false;
            this.$set(this.formItem, "isEditing", this.isEditing);
            // 校验
            if (
                this.formItem.required === "yes" &&
                this.formData[this.formItem.field_key] &&
                this.formData[this.formItem.field_key].length &&
                !this.selectArr.length
            ) {
                this.$message({
                    showClose: true,
                    message: "此为必填项",
                    type: "warning"
                });
                let selectArr = []; // 双向绑定id数组
                this.personList = _.cloneDeep(
                    this.formData[this.formItem.field_key]
                );
                if (this.personList) {
                    this.personList.forEach((objItem) => {
                        selectArr.push(objItem.id);
                    });
                }
                this.selectArr = selectArr;
                // this.frontArr = this.getArrFront(this.personList);
                // this.behandArr = this.getArrBehand(this.personList);
                this.frontArr = this.personList;
                this.$nextTick(() => {
                    this.getTagInit();
                });
            } else if (
                (!this.formData[this.formItem.field_key] &&
                    !this.selectArr.length) ||
                (this.formData[this.formItem.field_key] &&
                    !this.formData[this.formItem.field_key].length &&
                    !this.selectArr.length) ||
                JSON.stringify(this.formData[this.formItem.field_key]) ===
                    JSON.stringify(this.personList)
            ) {
                // 前后相同
            } else {
                this.$emit(
                    "edit-form-item",
                    JSON.stringify(this.selectArr),
                    this.formItem.field_key,
                    this.formItem.mode
                );
            }
        },
        removeTagItem(people) {
            if (
                this.formItem.required === "yes" &&
                this.selectArr.length === 1
            ) {
                this.$message({
                    showClose: true,
                    message: "此为必填项",
                    type: "warning"
                });
                return;
            }
            // 移除提交参数
            _.remove(this.selectArr, function (id) {
                return id === people.id;
            });
            // 移除页面绑定
            _.remove(this.personList, function (persion) {
                return persion.id === people.id;
            });
            this.frontArr = this.personList;
            this.$nextTick(() => {
                this.getTagInit();
            });
            // this.frontArr = this.getArrFront(this.personList);
            // this.behandArr = this.getArrBehand(this.personList);
            // this.getPeopleList();
        },
        confirmPeople(people) {
            // 列表选择
            this.personList = [people];
            // this.frontArr = this.getArrFront(this.personList);
            // this.behandArr = this.getArrBehand(this.personList);
            this.frontArr = this.personList;
            this.$nextTick(() => {
                this.getTagInit();
            });
            this.selectArr = [people.id];
            this.removeHadSelect();
            this.getPeopleList();
        },
        removeHadSelect() {
            this.selectArr.forEach((id) => {
                if (_.find(this.showList, { id: id })) {
                    _.remove(this.showList, { id: id });
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
@import "../../style.scss";
.detail {
    display: flex;
    height: 32px;
    padding: 0 10px;
    .tag-list {
        display: inline-block;
        height: 32px;
    }
}
.tag-item {
    display: inline-block;
    margin-right: 8px;
    overflow: hidden;
    text-wrap: nowrap;
    text-overflow: ellipsis;
    max-width: 148px;
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
</style>
