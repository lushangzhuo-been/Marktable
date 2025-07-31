<template>
    <div
        class="column-block"
        ref="ColumnBlock"
        :class="{ isEditing: isEditing }"
    >
        <el-popover
            ref="DropPopover"
            :width="popoverWidth"
            popper-class="progress-popover"
            placement="bottom"
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
                <div class="had-select-tag">
                    <!-- 遍历personList -->
                    <div
                        class="popover-tag-item"
                        v-for="(tagItem, tagIndex) in frontArr"
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
                        v-for="(listItem, listIndex) in personOption"
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
                    <div v-show="!personOption.length" class="list-no-data">
                        暂无数据
                    </div>
                </div>
            </div>
            <div
                slot="reference"
                class="detail"
                :class="{
                    'validate-failed': validateFailed,
                }"
            >
                <div v-if="frontArr.length">
                    <div
                        class="tag-item"
                        v-for="(tagItem, tagIndex) in frontArr"
                        :key="tagIndex"
                    >
                        <el-avatar
                            class="progress-avatar single-avatar"
                            icon="el-icon-user-solid"
                            size="small"
                            :src="getAvatar(tagItem.avatar)"
                            :style="geInputAvatar()"
                        ></el-avatar>
                        {{ tagItem.full_name }}
                    </div>
                </div>
                <div
                    class="default-text"
                    :class="{
                        'validate-failed': validateFailed,
                    }"
                    v-else
                >
                    请选择(单人)
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
            frontArr: [],
            popoverWidth: 220,
            validateFailed: false,
            selectArr: [],
            personList: [],
            personOption: [],
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
            handler(data) {
                let selectArr = [];
                this.personList = _.cloneDeep(data[this.formItem.field_key]);
                this.personList.forEach((objItem) => {
                    selectArr.push(objItem.id);
                });
                // this.personList = _.cloneDeep(data[this.formItem.field_key]);
                this.selectArr = selectArr;
                this.frontArr = this.getArrFront(this.personList);
            },
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
        getArrFront(arr) {
            let deepClone = _.cloneDeep(arr);
            let front = deepClone.splice(0, 1);
            return front;
        },
        popoverShow() {
            this.isEditing = true;
            this.popoverWidth = this.$refs.ColumnBlock.clientWidth;
            this.$nextTick(() => {
                if (this.$refs.SearchInput) {
                    this.$refs.SearchInput.focus();
                }
            });
            // 调查询人的接口
            this.getPeopleList();
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
                    this.personOption = res.data;
                } else {
                    this.personOption = [];
                }
            });
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
            this.personList = [];
            this.frontArr = this.getArrFront(this.personList);
            this.formData[this.formItem.field_key] = this.personList;
        },
        confirmPeople(people) {
            // 列表选择
            this.selectArr = [people.id];
            this.removeHadSelect();
            this.getPeopleList();
            this.personList = [people];
            this.frontArr = this.getArrFront(this.personList);
            this.formData[this.formItem.field_key] = this.personList;
        },
        removeHadSelect() {
            this.selectArr.forEach((id) => {
                if (_.find(this.showList, { id: id })) {
                    _.remove(this.showList, { id: id });
                }
            });
        },
        afterLeave() {
            this.isEditing = false;
            this.formData[this.formItem.field_key] = this.personList;
            this.doValidate();
        },
        doValidate() {
            if (this.formItem.required === "yes" && !this.personList.length) {
                this.validateFailed = true;
                this.formItem.validateFailed = true;
            } else {
                this.validateFailed = false;
                this.formItem.validateFailed = false;
            }
        },
        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        },
        getPopoverTagListAvatar() {
            return {
                position: "relative",
                top: "-6px",
            };
        },
        geInputAvatar() {
            return {
                position: "relative",
                top: "-2px",
            };
        },
    },
};
</script>

<style lang="scss" scoped>
@import "../../style.scss";
.default-text {
    color: #c0c4cc;
}

.avatar-box {
    display: inline-block;
    min-width: 24px;
    height: 24px;
    border-radius: 50%;
    background-size: 100% 100%;
    vertical-align: middle;
}
.list-item {
    padding: 0 12px;
    height: 40px;
    line-height: 40px;
    &:hover {
        background-color: #f1f9ff;
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
</style>
