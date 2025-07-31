<template>
    <div
        class="column-block"
        ref="column-block"
        :class="{ isEditing: isEditing }"
    >
        <el-popover
            ref="status-popover"
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
                        v-for="(tagItem, tagIndex) in selectedArrEdit"
                        :key="tagIndex"
                        @click="removeTagItem(tagItem)"
                    >
                        <div class="name-content">
                            {{ tagItem.title }}
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
                        {{ listItem.title }}
                    </div>
                    <div v-show="!personOption.length" class="list-no-data">
                        暂无数据
                    </div>
                </div>
            </div>
            <div slot="reference" class="detail">
                <div v-if="selectArr.length" class="tag-list">
                    <div
                        class="tag-item"
                        v-for="(tagItem, tagIndex) in selectArr"
                        :key="tagIndex"
                    >
                        {{ tagItem.title }}
                    </div>
                </div>
                <div class="default-text" v-else>请选择关联的流程数据</div>
                <b class="triangle"></b>
            </div>
        </el-popover>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
import { imgHost } from "@/assets/tool/const";
import api_common from "@/common/api";
export default {
    props: {
        configData: {
            type: Object,
            default: () => {}
        },
        param: {
            type: [Array, String],
            default: () => []
        }
    },
    data() {
        return {
            isEditing: false,
            selectArr: [],
            selectedArrEdit: [],
            selectText: "",
            submitArr: [],
            popoverWidth: 220,
            personOption: [],
            searchInput: "",
            curConfigItem: {}
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
        configData: {
            handler(obj) {
                if (obj && Object.keys(obj).length) {
                    if (obj.filterConfig && obj.filterConfig.length) {
                        obj.filterConfig.forEach((item) => {
                            if (item.field_key === obj.filterItem.field_key) {
                                this.curConfigItem = _.cloneDeep(item);
                            }
                        });
                    }
                }
            },
            deep: true,
            immediate: true
        },
        param: {
            handler(arr) {
                this.selectArr = arr && arr.length ? arr : [];
                this.selectedArrEdit = _.cloneDeep(this.selectArr);
            },
            immediate: true
        },
        searchInput: _.debounce(function () {
            this.getProgressDataList();
        }, 500)
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
            this.popoverWidth = this.$refs["column-block"].clientWidth;
            // 调查询人的接口
            this.getProgressDataList();
        },
        getProgressDataList() {
            let params = {
                ws_id: this.curSpace.id,
                ex: this.selectedArrEdit.length
                    ? this.selectedArrEdit.map((item) => item._id).join(",")
                    : "",
                key: this.searchInput,
                page_size: 50,
                page: 1,
                tmpl_id: this.curConfigItem.related_tmpl_id
            };
            api_common.getOtherTmplDataList(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.personOption = res.data.list || [];
                } else {
                    this.personOption = [];
                }
            });
        },
        afterLeave() {
            this.isEditing = false;
        },
        getPopoverTagListAvatar() {
            return {
                position: "relative",
                top: "-7px"
            };
        },
        // 选择下拉选项时
        confirmPeople(people) {
            this.selectedArrEdit = [people];
            this.selectArr = _.cloneDeep(this.selectedArrEdit);
            this.getProgressDataList();
            this.$emit("input", this.selectArr);
        },
        // 移除选中的
        removeTagItem() {
            this.selectedArrEdit = [];
            this.selectArr = _.cloneDeep(this.selectedArrEdit);
            this.getProgressDataList();
            this.$emit("input", this.selectArr);
        }
    }
};
</script>

<style lang="scss" scoped>
@import "./style.scss";
.column-block {
    width: 220px !important;
}
.default-text {
    color: #c0c4cc;
    overflow: hidden;
}
.column-block .detail {
    padding: 0 20px 0 10px !important;
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
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    &:hover {
        background-color: #f1f9ff;
    }
}
.tag-list {
    height: 100%;
}
.tag-item {
    display: inline-block;
    overflow: hidden;
    text-wrap: nowrap;
    text-overflow: ellipsis;
    max-width: 100%;
}
.had-select-tag {
    padding: 4px;
    margin-top: 4px;
}
.popover-tag-item {
    // box-sizing: border-box;
    display: inline-block;
    padding: 0 8px;
    height: 20px;
    line-height: 19px;
    font-size: 12px;
    color: #82889e;
    background-color: #e2e7f0;
    /* margin: 4px; */
    border-radius: 10px;
    max-width: calc(100% - 16px);
    .name-content {
        display: inline-block;
        max-width: calc(100% - 16px);
        margin-right: 4px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }
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
