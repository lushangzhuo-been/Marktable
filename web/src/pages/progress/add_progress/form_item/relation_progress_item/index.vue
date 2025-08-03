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
                <!-- 已选中流程 -->
                <div class="had-select-tag" v-show="selectedArrEdit.length">
                    <div
                        class="popover-tag-item"
                        v-for="(tagItem, tagIndex) in selectedArrEdit"
                        :key="tagIndex"
                    >
                        <span class="name-content">
                            {{ tagItem.title }}
                        </span>
                        <!-- 移除icon -->
                        <i
                            @click="removeTagItem(tagItem)"
                            class="el-tag__close el-icon-close tag-delete"
                        ></i>
                    </div>
                </div>
                <!-- 下拉列表 -->
                <div class="list-content">
                    <div
                        class="list-item multiple-people-list-item"
                        v-for="(listItem, listIndex) in progressDataList"
                        :key="listIndex"
                        @click="confirmPeople(listItem)"
                    >
                        {{ listItem.title }}
                    </div>
                    <div v-show="!progressDataList.length" class="list-no-data">
                        暂无数据
                    </div>
                </div>
            </div>
            <!-- 标签化 -->
            <div
                class="detail"
                :class="{
                    'validate-failed': validateFailed
                }"
                slot="reference"
            >
                <span v-if="selectedArrEdit && selectedArrEdit.length">
                    {{ selectedArrEdit[0].title }}
                </span>
                <span v-else class="default-text">请选择关联流程</span>
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
import api_common from "@/common/api";
export default {
    props: {
        formItem: {
            type: Object,
            default: () => {}
        },
        formData: {
            type: Object,
            default: () => {}
        },
        validateOrder: {
            type: Number,
            default: 0
        }
    },
    data() {
        return {
            isEditing: false,
            popoverWidth: 220,
            validateFailed: false,
            selectArr: [],
            selectedArrEdit: [],
            progressDataList: [],
            searchInput: ""
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
            handler(data) {
                let curValue = data[this.formItem.field_key];
                this.selectArr = curValue && curValue.length ? curValue : [];
                // this.selectedArrEdit = _.cloneDeep(this.selectArr);
            },
            deep: true
        },
        validateOrder: {
            handler() {
                this.doValidate();
            }
        },
        searchInput: _.debounce(function () {
            this.getProgressDataList();
        }, 500)
    },
    methods: {
        popoverShow() {
            this.isEditing = true;
            this.popoverWidth = this.$refs.ColumnBlock.clientWidth;
            this.$nextTick(() => {
                if (this.$refs.SearchInput) {
                    this.$refs.SearchInput.focus();
                }
            });
            this.getProgressDataList();
        },
        getProgressDataList() {
            let params = {
                ws_id: this.curSpace.id,
                ex: this.selectedArrEdit.length
                    ? this.selectedArrEdit.map((item) => item._id).join(",")
                    : "",
                key: this.searchInput,
                page_size: 20,
                page: 1,
                tmpl_id: this.formItem.related_tmpl_id // 关联流程列的流程
            };
            api_common.getOtherTmplDataList(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.progressDataList = res.data.list || [];
                } else {
                    this.progressDataList = [];
                }
            });
        },
        // 关闭popover
        afterLeave() {
            this.isEditing = false;
            this.selectArr = _.cloneDeep(this.selectedArrEdit);
            this.formData[this.formItem.field_key] = this.selectArr.map(
                (item) => item._id
            );
            this.doValidate();
        },
        // 选择下拉选项时
        confirmPeople(people) {
            this.selectedArrEdit = [people];
            this.getProgressDataList();
        },
        // 移除选中的
        removeTagItem() {
            this.selectedArrEdit = [];
            this.getProgressDataList();
        },
        doValidate() {
            if (this.formItem.required === "yes" && !this.selectArr.length) {
                this.validateFailed = true;
                this.formItem.validateFailed = true;
            } else {
                this.validateFailed = false;
                this.formItem.validateFailed = false;
            }
        }
    }
};
</script>

<style lang="scss" scoped>
@import "../../style.scss";
.detail {
    padding: 0 20px 0 10px !important;
}
.default-text {
    color: #c0c4cc;
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
.tag-item {
    display: inline-block;
    margin-right: 8px;
    overflow: hidden;
    text-wrap: nowrap;
    text-overflow: ellipsis;
    max-width: 148px;
}
.popover-tag-item {
    box-sizing: border-box;
    display: inline-block;
    padding: 0 8px;
    height: 20px;
    line-height: 19px;
    font-size: 12px;
    color: #82889e;
    background-color: #e2e7f0;
    /* margin: 4px; */
    border-radius: 10px;
    max-width: 100%;
    .name-content {
        display: inline-block;
        max-width: calc(100% - 16px);
        /* padding: 0 4px; */
        margin-right: 4px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }
}
.had-select-tag {
    padding: 4px;
    margin-top: 4px;
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
