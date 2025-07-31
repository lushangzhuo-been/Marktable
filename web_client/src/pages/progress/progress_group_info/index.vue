<template>
    <div>
        <div class="sub-title">
            <img
                :src="require(`@/assets/image/common/progress_update.png`)"
                alt=""
                width="20px"
                height="20px"
                class="icon"
            />
            基本信息修改
            <img
                v-if="!isEdit"
                class="edit"
                :src="require(`@/assets/image/common/edit-active.svg`)"
                alt=""
                @click="editfolderInfoForm"
            />
        </div>
        <el-form
            ref="folderInfoForm"
            :model="folderInfoForm"
            :rules="rule"
            class="detail paddingb basic-ui"
            label-width="60px"
            v-if="isEdit"
            size="small"
        >
            <el-form-item label="名称:" prop="name">
                <el-input
                    type="text"
                    class="basic-ui width372"
                    maxlength="25"
                    show-word-limit
                    placeholder="请输入流程名称"
                    size="small"
                    v-model="folderInfoForm.name"
                ></el-input>
            </el-form-item>
            <el-form-item label="描述:">
                <el-input
                    class="basic-ui fixed-height"
                    type="textarea"
                    size="small"
                    :autosize="{ minRows: 2, maxRows: 6 }"
                    v-model="folderInfoForm.desc"
                ></el-input>
            </el-form-item>
            <el-form-item>
                <el-button size="small" @click="cancelEdit">取 消</el-button>
                <el-button
                    class="basic-ui"
                    size="small"
                    type="primary"
                    @click="onConfirmEdit"
                    >确 定</el-button
                >
            </el-form-item>
        </el-form>
        <el-descriptions :column="1" v-else class="detail">
            <el-descriptions-item
                v-for="(item, index) in folderInfoList"
                :key="index"
                :label="item.label"
            >
                <span v-if="item.label === '类型'">
                    {{ formatTypeValue(item.value) }}
                </span>
                <span v-else>
                    {{ item.value || "--" }}
                </span>
            </el-descriptions-item>
        </el-descriptions>
    </div>
</template>

<script>
import SpaceInfo from "@/pages/common/space_info";
import dataHandle from "../data_handle";
import _ from "lodash";
import NoData from "@/pages/common/no_data";
import api from "@/common/api/module/progress_setting";
export default {
    components: {
        NoData,
        SpaceInfo
    },
    data() {
        return {
            isShowSpaceInfo: false,
            curRes: {},
            isEdit: false,
            folderInfoForm: {
                name: "",
                desc: ""
            },
            folderInfoList: dataHandle.defaultProgressInfo,
            folderInfoListTmp: [],
            rule: {
                name: [
                    {
                        required: true,
                        trigger: "blur",
                        message: "请输入流程名称"
                    }
                ],
                desc: [
                    {
                        required: true,
                        trigger: "blur",
                        message: "请输入流程描述"
                    }
                ]
            }
        };
    },
    computed: {
        formatTypeValue() {
            return (val) => {
                if (!val) return "--";
                return val === "private" ? "私有流程" : "公有流程";
            };
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        progressTreeList() {
            return this.$store.state.progressTree;
        },
        isNoData() {
            return (
                (this.progressTreeList && this.progressTreeList.length) || false
            );
        }
    },
    watch: {
        $route(to, from) {
            this.isEdit = false;
            this.isShowSpaceInfo = false;
            this.fetchfolderInfoForm();
        },
        progressTreeList(arr) {
            if (arr && arr.length) {
                if (
                    arr.some(
                        (item) => item.id === Number(this.$route.params.id)
                    )
                ) {
                    this.isEdit = false;
                    this.isShowSpaceInfo = false;
                    this.fetchfolderInfoForm();
                } else {
                    this.isShowSpaceInfo = true;
                }
            }
        }
    },
    mounted() {
        this.$nextTick(() => {
            if (!this.isNoData) return;
            this.fetchfolderInfoForm();
        });
    },
    methods: {
        fetchfolderInfoForm() {
            let params = {
                id: this.$route.params.id,
                ws_id: this.curSpace.id
            };
            api.getFolderDetail(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    this.setFormData(res.data);
                    this.curRes = _.cloneDeep(res.data);
                    this.folderInfoList = dataHandle.defaultfolderInfoForm(
                        res.data
                    );
                } else {
                    this.folderInfoList = dataHandle.defaultfolderInfoForm({});
                    this.folderInfoForm = {
                        name: "",
                        desc: ""
                    };
                }
            });

            // 接口成功和失败
        },
        setFormData(res) {
            for (let key in this.folderInfoForm) {
                this.folderInfoForm[key] = res[key] || "";
            }
        },
        editfolderInfoForm() {
            this.isEdit = true;
        },
        cancelEdit() {
            this.setFormData(this.curRes);
            this.$refs.folderInfoForm.clearValidate();
            this.isEdit = false;
        },

        onConfirmEdit() {
            // 调接口提交到后端
            this.$refs.folderInfoForm.validate((flag) => {
                if (flag) {
                    let params = {
                        id: this.$route.params.id,
                        ws_id: this.curSpace.id,
                        ...this.folderInfoForm
                    };
                    api.updateFolderDetail(params).then((res) => {
                        if (res && res.resultCode === 200) {
                            this.isEdit = false;
                            this.fetchfolderInfoForm();
                            this.$store.dispatch("fetchProgressTree", {
                                ws_id: this.curSpace.id
                            });
                        }
                    });
                }
            });
        },
        getPrivate(val) {
            let map = {
                private: "私有流程",
                public: "公有流程"
            };
            return map[val] || "--";
        }
    }
};
</script>

<style lang="scss" scoped>
.width200 {
    width: 60%;
}
.edit {
    position: relative;
    top: 2px;
    height: 16px;
    width: 16px;
    margin-left: 8px;
    cursor: pointer;
}
.detail.paddingb {
    padding-bottom: 24px;
}
.width372 {
    width: 372px;
}
.width642 {
    width: 642px;
}
.detail {
    background: #f5f6f8;
    padding: 24px 24px 4px;
    border-radius: 4px;
    // opacity: 0.5;
    ::v-deep .el-form-item {
        margin-bottom: 20px;
        &:last-child {
            margin-bottom: 0;
        }
        .el-form-item__label {
            color: #828898;
        }
        .el-form-item__content {
            color: #2a3048;
        }
    }
    .value {
        margin-left: 16px;
    }
    .el-textarea {
        width: 80%;
    }

    ::v-deep .el-descriptions__body {
        background-color: #f5f6f8;
        .el-descriptions-item__label:not(.is-bordered-label) {
            display: inline-block;
            margin-right: 16px;
            width: 62px;
            text-align: right;
            color: #828898;
        }
        :not(.is-bordered) .el-descriptions-item__cell {
            padding-bottom: 20px;
        }
    }
}
</style>
