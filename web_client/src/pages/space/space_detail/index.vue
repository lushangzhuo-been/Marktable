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

            空间详情
            <img
                v-if="!isEdit"
                class="edit-img"
                :src="require(`@/assets/image/common/edit-active.svg`)"
                alt=""
                @click="edit"
            />
        </div>
        <!-- <div class="detail"> -->
        <el-form
            v-if="isEdit"
            ref="spaceDetailForm"
            :model="spaceDetailForm"
            :rules="spaceDetailRule"
            label-width="82px"
            class="detail paddingb"
            size="small"
        >
            <el-form-item prop="name" label="名称:">
                <el-input
                    class="basic-ui width372"
                    placeholder="请输入空间名称"
                    size="small"
                    maxlength="25"
                    show-word-limit
                    v-model="spaceDetailForm.name"
                ></el-input>
            </el-form-item>
            <el-form-item prop="desc" label="描述:">
                <el-input
                    class="basic-ui fixed-height"
                    type="textarea"
                    size="small"
                    placeholder="请输入空间描述"
                    :clearable="true"
                    v-model="spaceDetailForm.desc"
                ></el-input>
            </el-form-item>
            <!-- <el-form-item label="创建人:">
                <el-input
                    disabled
                    class="basic-ui width372"
                    size="small"
                    placeholder="请输入创建人"
                    v-model="spaceDetailForm.creator_name"
                ></el-input>
            </el-form-item>
            <el-form-item label="创建时间:">
                <el-input
                    class="basic-ui width372"
                    disabled
                    size="small"
                    placeholder="请输入创建时间"
                    v-model="spaceDetailForm.created_at"
                ></el-input>
            </el-form-item>
            <el-form-item label="管理员:">
                <el-select
                    class="basic-ui width372"
                    placeholder="请输入管理员"
                    disabled
                    multiple
                    size="small"
                    v-model="spaceDetailForm.admin"
                ></el-select>
            </el-form-item> -->
            <el-form-item>
                <el-button
                    class="basic-ui width100"
                    size="small"
                    @click="cancel"
                    >取消</el-button
                >
                <el-button
                    class="basic-ui width100"
                    type="primary"
                    size="small"
                    @click="updateDetail"
                    >确定</el-button
                >
            </el-form-item>
        </el-form>
        <el-descriptions :column="1" v-else class="detail">
            <el-descriptions-item
                v-for="(item, index) in spaceDetailList"
                :key="index"
                :label="item.label"
            >
                <span v-if="item.label === '管理员'">
                    <span v-if="item.value && item.value.length" class="admin">
                        <span
                            v-for="(sub, i) in item.value"
                            :key="i"
                            class="item"
                        >
                            <el-avatar
                                icon="el-icon-user-solid"
                                size="small"
                                :src="getAvatar(sub)"
                            ></el-avatar>
                            <span>{{ sub.full_name }}</span>
                            <span v-if="i < item.value.length - 1">,</span>
                        </span>
                    </span>
                    <span v-else>--</span>
                </span>
                <span v-else>
                    {{ item.value || "--" }}
                </span>
            </el-descriptions-item>
        </el-descriptions>
    </div>
</template>

<script>
import { imgHost } from "@/assets/tool/const";
import dataHandle from "../data_handle";
import api from "@/common/api/module/space";
import _ from "lodash";
export default {
    computed: {
        getAdminFormat() {
            return (val) => {
                return val || "--";
            };
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        getAvatar() {
            return (item) => {
                if (item && item.avatar) {
                    return `${imgHost}${item.avatar}`;
                }
                return require(`@/assets/image/common/default_avatar_big.png`);
            };
        }
    },
    data() {
        return {
            curRes: {},
            isEdit: false,
            spaceDetailForm: {
                name: "",
                desc: "",
                creator_name: "",
                created_at: "",
                admin: []
            },
            spaceDetailList: dataHandle.defaultSpaceDetail,
            spaceDetailListTmp: [],
            spaceDetailRule: {
                name: [
                    {
                        required: true,
                        message: "空间名不能为空",
                        trigger: "blur"
                    }
                ],
                desc: [
                    {
                        required: true,
                        message: "空间描述不能为空",
                        trigger: "blur"
                    }
                ]
            }
        };
    },
    watch: {
        curSpace: {
            handler() {
                this.isEdit = false;
                this.fetchSpaceDetail(1);
            },
            deep: true
        }
    },
    mounted() {
        // 第一次进入时调接口，刷新不调接口
        this.fetchSpaceDetail(1);
    },
    methods: {
        fetchSpaceDetail() {
            let params = {
                id: this.curSpace.id
            };
            api.getWorkspaceInfo(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    this.setFormData(res.data);
                    this.curRes = _.cloneDeep(res.data);
                    this.spaceDetailList = dataHandle.spaceDetailForm(res.data);
                    // if (num === 1) {
                    //     this.$store.dispatch('fetchSpaceList').then(() => {
                    //         let curSp = _.cloneDeep(this.curSpace)
                    //         curSp.name = this.spaceDetailForm.name
                    //         this.$store.commit('setCurSpace', curSp)
                    //         this.isEdit = false
                    //     })
                    // }
                } else {
                    this.spaceDetailList = dataHandle.defaultSpaceDetail;
                    this.spaceDetailForm = {
                        name: "",
                        desc: "",
                        creator_name: "",
                        created_at: "",
                        admin: []
                    };
                }
            });
            // let res = {
            //     name: '数字化平台小组',
            //     desc: '此空间为数字化平台小组的内部空间，主要用于项目管理等事务。',
            //     creator_name: '414262042@qq.com',
            //     created_at: '2023-08-11 11:35:50',
            //     admin: 'Neil Guo',
            // }
            // 接口成功和失败
        },
        setFormData(res) {
            for (let key in this.spaceDetailForm) {
                if (key === "admin") {
                    this.spaceDetailForm[key] =
                        res[key] && res[key].length
                            ? res[key].map((item) => item.username)
                            : "";
                } else {
                    this.spaceDetailForm[key] = res[key] || "";
                }
            }
        },
        edit() {
            this.isEdit = true;
        },
        updateDetail() {
            this.$refs.spaceDetailForm.validate((flag) => {
                if (flag) {
                    let params = {
                        id: this.curSpace.id,
                        name: this.spaceDetailForm.name,
                        desc: this.spaceDetailForm.desc
                    };
                    api.updateWorkspace(params).then((res) => {
                        if (res.resultCode === 200) {
                            this.fetchSpaceDetail();
                            this.$store.dispatch("fetchSpaceList").then(() => {
                                let curSp = _.cloneDeep(this.curSpace);
                                curSp.name = this.spaceDetailForm.name;
                                this.$store.commit("setCurSpace", curSp);
                                this.isEdit = false;
                            });
                        }
                    });
                }
            });
        },
        cancel() {
            this.setFormData(this.curRes);
            this.$refs.spaceDetailForm.clearValidate();
            this.isEdit = false;
        }
    }
};
</script>

<style lang="scss" scoped>
.edit-img {
    cursor: pointer;
    height: 16px;
    width: 16px;
    position: relative;
    top: 2px;
    margin-left: 8px;
}
.width372 {
    width: 372px;
}
.width642 {
    width: 642px;
}
.admin {
    display: flex;
    flex-wrap: wrap;
    .item {
        display: flex;
        align-items: center;
        margin-bottom: 6px;
        margin: 0 20px 10px 0;
        .el-avatar {
            margin-right: 8px;
            height: 20px;
            width: 20px;
            line-height: 20px;
        }
    }
}
.detail.paddingb {
    padding-bottom: 24px;
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
            width: 63px;
            text-align: right;
            color: #828898;
        }
        :not(.is-bordered) .el-descriptions-item__cell {
            padding-bottom: 20px;
        }
    }
}
</style>
