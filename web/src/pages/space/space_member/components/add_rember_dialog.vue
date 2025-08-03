<template>
    <el-dialog
        title="添加成员"
        :visible.sync="dialogVisible"
        class="add-rember-dialog basic-ui"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        @close="cancel"
    >
        <el-form
            ref="spaceRemberForm"
            class="basic-ui"
            :model="spaceRemberForm"
            :rules="rule"
            label-width="60px"
            size="small"
        >
            <el-form-item prop="username" label="名称:">
                <el-select
                    v-model="spaceRemberForm.username"
                    clearable
                    size="small"
                    multiple
                    popper-class="space-member-name"
                    class="basic-ui width100 user-tag"
                    placeholder="请输入用户名或昵称或邮箱"
                    @visible-change="visibleChange"
                >
                    <div class="customer-user-search">
                        <el-input
                            ref="SearchInput"
                            class="user-search-input"
                            prefix-icon="el-icon-search"
                            v-model="userKey"
                            placeholder="请输入"
                        ></el-input>
                        <div
                            class="user-list"
                            :class="{ empty: !nameOptions.length }"
                        >
                            <el-checkbox-group
                                class="user-list-checkbox-group"
                                v-model="spaceRemberForm.username"
                            >
                                <el-option
                                    v-for="item in nameOptions"
                                    :key="item.value"
                                    :label="item.text"
                                    :value="item.value"
                                >
                                    <span class="format-role-option">
                                        <span class="username">
                                            {{ item.label }}
                                        </span>
                                        <span class="name">
                                            <el-avatar
                                                icon="el-icon-user-solid"
                                                size="small"
                                                :src="getAvatar(item.icon)"
                                            ></el-avatar>
                                            <span>{{ item.text }}</span>
                                        </span>
                                        <span class="email">{{
                                            item.email
                                        }}</span>
                                        <el-checkbox
                                            @click.native.prevent
                                            :label="item.value"
                                        ></el-checkbox>
                                    </span>
                                </el-option>
                            </el-checkbox-group>
                        </div>
                    </div>
                </el-select>
            </el-form-item>
            <el-form-item prop="role" label="角色:">
                <el-select
                    v-model="spaceRemberForm.role"
                    clearable
                    size="small"
                    class="basic-ui width100"
                    placeholder="请选择角色"
                >
                    <el-option
                        v-for="item in roleOptions"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    >
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item prop="status" label="状态:">
                <!-- <el-switch
                    v-model="spaceRemberForm.status"
                    :active-color="PRIMARY_COLOR"
                    inactive-color="#E6E9F0"
                    size="small"
                >
                </el-switch> -->
                <el-radio-group
                    class="basic-ui right"
                    v-model="spaceRemberForm.status"
                >
                    <el-radio label="ok">开启</el-radio>
                    <el-radio label="forbid">禁用</el-radio>
                </el-radio-group>
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
            <el-button class="basic-ui" size="small" @click="cancel"
                >取 消</el-button
            >
            <el-button
                class="basic-ui"
                size="small"
                type="primary"
                @click="onConfirm"
                >确 定</el-button
            >
        </span>
    </el-dialog>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/space";
import api_common from "@/common/api";
import { PRIMARY_COLOR_LIST } from "@/assets/tool/color_list";
import { imgHost } from "@/assets/tool/const";
export default {
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        getAvatar() {
            return (avatar) => {
                if (avatar) {
                    return `${imgHost}${avatar}`;
                }
                return require(`@/assets/image/common/default_avatar_big.png`);
            };
        },
    },
    data() {
        return {
            PRIMARY_COLOR: PRIMARY_COLOR_LIST.PRIMARY_COLOR,
            dialogVisible: false,
            spaceRemberForm: {
                username: [],
                role: "",
                status: "ok",
            },
            roleOptions: [
                {
                    label: "管理员",
                    value: "admin",
                },
                {
                    label: "普通",
                    value: "general",
                },
            ],
            nameOptions: [],
            rule: {
                userid: [
                    {
                        required: true,
                        message: "请输入账号",
                        // trigger: "blur"
                    },
                ],
                username: [
                    {
                        required: true,
                        message: "请输入名称",
                        trigger: "change",
                    },
                ],
                role: [
                    {
                        required: true,
                        message: "请选择角色",
                        trigger: "change",
                    },
                ],
            },
            userKey: "",
        };
    },
    watch: {
        userKey: _.debounce(function (key) {
            if (key) {
                this.fetchAllUserList(key);
            } else {
                this.nameOptions = [];
            }
        }, 500),
    },
    methods: {
        visibleChange(show) {
            if (show) {
                this.$nextTick(() => {
                    setTimeout(() => {
                        this.$refs.SearchInput.focus();
                    }, 300);
                });
            }
        },
        openDialog() {
            this.dialogVisible = true;
            this.$nextTick(() => {
                this.$refs.spaceRemberForm.clearValidate();
            });
        },
        fetchAllUserList(val) {
            let params = {
                key: val,
                page_size: 10,
                page: 1,
            };
            api_common.getCurWorkSpaceMemberList(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.nameOptions = this.nameListFormat(res.data.list || []);
                } else {
                    this.nameOptions = [];
                }
            });
        },
        nameListFormat(arr) {
            if (arr && arr.length) {
                return arr.map((item) => {
                    return {
                        ...item,
                        label: item.username,
                        value: item.id,
                        text: item.full_name,
                        icon: item.avatar,
                    };
                });
            }
            return [];
        },
        cancel() {
            this.dialogVisible = false;
            this.spaceRemberForm = {
                username: [],
                role: "",
                status: "ok",
            };
            this.$nextTick(() => {
                this.$refs.spaceRemberForm.clearValidate();
            });
        },
        onConfirm() {
            this.$refs.spaceRemberForm.validate((flag) => {
                if (flag) {
                    // this.$emit('on-confirm', this.spaceRemberForm)
                    let params = {
                        ws_id: this.curSpace.id,
                        userid_list: this.spaceRemberForm.username.join(),
                        status: this.spaceRemberForm.status,
                        role: this.spaceRemberForm.role,
                    };
                    api.addWorkspaceMember(params).then((res) => {
                        if (res && res.resultCode === 200) {
                            this.$emit("on-confirm", this.spaceRemberForm);
                        }
                    });
                }
            });
        },
    },
};
</script>

<style lang="scss">
.customer-user-search {
    position: relative;
    min-height: 32px;
    .user-search-input {
        padding: 4px 0;
        border-bottom: 1px solid #e6e9f0;
        .el-input__inner {
            border: 1px solid rgba(0, 0, 0, 0);
        }
    }
    .user-list {
        padding: 8px;
        max-height: 320px;
        overflow-y: auto;
        &.empty {
            padding: 0;
        }
        // .el-select-dropdown {
        //     .el-select-dropdown__item {
        //         padding: 0 px;
        //     }
        // }
    }
}
.add-rember-dialog.basic-ui {
    .el-dialog {
        width: 600px;
        border-radius: 4px;
        .el-dialog__header {
            padding: 24px 32px;
            font-size: 16px;
            color: #2f384c;
        }
        .el-dialog__body {
            padding: 0 32px;
        }
        .el-dialog__footer {
            padding: 0 32px 32px;
        }
        .el-dialog__headerbtn {
            top: 24px;
            right: 32px;
        }
        .el-form-item {
            margin-bottom: 24px;
        }
        .el-dialog__footer {
            text-align: right;
        }
    }
}
.space-member-name {
    .el-scrollbar {
        display: block !important;
        margin: 0 !important;
        .el-select-dropdown__wrap {
            max-height: 900px;
        }
    }
    .el-checkbox-group .el-checkbox__label {
        display: none;
    }
}
.space-member-name.el-select-dropdown {
    .el-scrollbar {
        .el-select-dropdown__item {
            padding: 0 16px;
            &.selected {
                color: var(--FONT_COLOR);
                font-weight: 400;
            }
        }
    }
    .el-select-dropdown__empty {
        padding: 20px 0;
    }
    .popper__arrow {
        display: none;
    }
    &.el-popper {
        margin-top: 4px;
    }
}
.format-role-option {
    display: flex;
    justify-content: space-between;
    align-items: center;
    .username {
        width: 120px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }
    .name {
        flex: 1;
        padding: 0 12px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        .el-avatar {
            vertical-align: middle;
            margin-right: 10px;
            height: 20px;
            width: 20px;
            line-height: 20px;
            border-radius: 50%;
        }
    }
    .email {
        width: 160px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }
}
.user-tag {
    padding: 4px 0;
    .el-select__tags {
        .el-tag {
            height: 24px;
            background: #e9f0f8;
            border-radius: 4px;
            font-size: 12px;
            color: #171e31;
            .el-tag__close {
                background-color: #cdd5e6 !important;
            }
        }
    }
}
</style>
<style lang="scss" scoped>
.add-rember-dialog {
    .width100 {
        width: 100%;
    }
}
</style>
