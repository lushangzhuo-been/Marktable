<template>
    <div
        class="column-block"
        ref="ColumnBlock"
        :class="{ isEditing: isEditing }"
    >
        <el-popover
            ref="status-popover"
            popper-class="progress-popover"
            placement="bottom"
            :width="popoverWidth"
            trigger="click"
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
                <!-- 所选tag -->
                <div v-if="personList.length" class="had-select-tag">
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
                        ></el-avatar>
                        {{ listItem.full_name }}
                    </div>
                    <div v-show="!showList.length" class="list-no-data">
                        暂无数据
                    </div>
                </div>
            </div>
            <!-- 标签化 -->
            <div slot="reference">
                <!-- 名字列表 -->
                <div
                    class="detail"
                    :class="{
                        'validate-failed': validateFailed,
                    }"
                    v-if="frontArr.length"
                >
                    <div
                        class="tag-item"
                        v-for="(tagItem, tagIndex) in frontArr"
                        :key="tagIndex"
                    >
                        <el-avatar
                            class="progress-avatar"
                            icon="el-icon-user-solid"
                            size="small"
                            :src="getAvatar(tagItem.avatar)"
                            :style="getAvatarStack(tagItem.avatar, tagIndex)"
                        ></el-avatar>
                        {{ tagItem.full_name }}
                    </div>
                    <!-- 数字气泡 -->
                    <el-tooltip class="item" effect="dark" placement="top">
                        <div slot="content">
                            <div
                                v-for="(tagItem, tagIndex) in behandArr"
                                :key="tagIndex"
                                class="num-list"
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
                        <b v-show="behandArr.length" class="num-box"
                            >+{{ behandArr.length }}</b
                        >
                    </el-tooltip>
                </div>
                <div
                    class="detail default-text"
                    :class="{
                        'validate-failed': validateFailed,
                    }"
                    v-else
                >
                    请选择(多人)
                </div>
                <b class="triangle"></b>
            </div>
        </el-popover>
        <div v-show="validateFailed" class="validate-desc">此项为必填</div>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import { imgHost } from "@/assets/tool/const";
export default {
    props: {
        formItem: {
            type: Object,
            default: () => {},
        },
        formData: {
            type: Object,
            default: () => {},
        },
        validateOrder: {
            type: Number,
            default: 0,
        },
    },
    data() {
        return {
            isEditing: false,
            selectArr: [],
            vModelArr: [],
            frontArr: [],
            behandArr: [],
            popoverWidth: 220,
            validateFailed: false,
            showList: [],
            personList: [],
            peopleList: [],
            searchInput: "",
        };
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        },
    },
    watch: {
        formData: {
            handler(formData) {
                let selectArr = [];
                let vModelArr = [];
                if (
                    formData[this.formItem.field_key] &&
                    formData[this.formItem.field_key].length
                ) {
                    vModelArr = formData[this.formItem.field_key];

                    formData[this.formItem.field_key].forEach((objItem) => {
                        selectArr.push(objItem.id);
                    });
                }
                this.selectArr = selectArr;
                this.personList = vModelArr;
                this.frontArr = this.getArrFront(vModelArr);
                this.behandArr = this.getArrBehand(vModelArr);
                this.$emit(
                    "edit-form-item",
                    JSON.stringify(this.selectArr),
                    this.formItem.field_key,
                    this.formItem.mode
                );
            },
            // immediate: true,
            deep: true,
        },
        validateOrder: {
            handler(order) {
                this.doValidate();
            },
        },
        searchInput: _.debounce(function () {
            this.getPeopleList();
        }, 500),
    },
    mounted() {},
    methods: {
        doValidate() {
            if (this.formItem.required === "yes" && !this.selectArr.length) {
                this.validateFailed = true;
                this.formItem.validateFailed = true;
            } else {
                this.validateFailed = false;
                this.formItem.validateFailed = false;
            }
        },
        getPeopleList() {
            let params = {
                ws_id: this.curSpace.id,
                ex: this.selectArr.join(","),
                key: this.searchInput,
                page_size: 50,
                page: 1,
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
        popoverShow() {
            this.searchInput = "";
            this.isEditing = true;
            this.popoverWidth = this.$refs.ColumnBlock.clientWidth;
            this.$nextTick(() => {
                if (this.$refs.SearchInput) {
                    this.$refs.SearchInput.focus();
                }
            });
            this.getPeopleList();
        },
        afterLeave() {
            this.isEditing = false;
            this.$emit(
                "edit-form-item",
                JSON.stringify(this.selectArr),
                this.formItem.field_key,
                this.formItem.mode
            );
            this.doValidate();
        },
        getArrFront(arr) {
            let deepClone = _.cloneDeep(arr);
            let front = deepClone.splice(0, 1);
            return front;
        },
        getArrBehand(arr) {
            let deepClone = _.cloneDeep(arr);
            let behand = deepClone.splice(1);
            return behand;
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
                top: "-2px",
            };
        },
        getPopoverTagListAvatar() {
            return {
                position: "relative",
                top: "-7px",
            };
        },
        removeTagItem(people) {
            // 移除提交参数
            _.remove(this.selectArr, function (id) {
                return id === people.id;
            });
            // 移除页面绑定
            _.remove(this.personList, function (persion) {
                return persion.id === people.id;
            });
            this.getPeopleList();
            this.doValidate();
        },
        confirmPeople(people) {
            // 列表选择
            this.selectArr.push(people.id);
            this.removeHadSelect();
            this.getPeopleList();
            if (_.find(this.peopleList, { id: people.id })) {
                this.personList.push(
                    _.find(this.peopleList, { id: people.id })
                );
            }
            this.frontArr = this.getArrFront(this.personList);
            this.behandArr = this.getArrBehand(this.personList);
            this.doValidate();
        },
        removeHadSelect() {
            this.selectArr.forEach((id) => {
                if (_.find(this.showList, { id: id })) {
                    _.remove(this.showList, { id: id });
                }
            });
            this.frontArr = this.getArrFront(this.personList);
            this.behandArr = this.getArrBehand(this.personList);
        },
    },
};
</script>

<style lang="scss" scoped>
@import "../style.scss";
.default-text {
    color: #c0c4cc;
}

.tag-item {
    display: inline-block;
    margin-right: 8px;
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
    top: -2px;
    vertical-align: middle;
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
.num-list {
    margin: 4px 0;
}
</style>
