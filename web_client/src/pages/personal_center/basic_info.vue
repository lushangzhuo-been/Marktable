<template>
    <div>
        <div class="sub-title">
            <img
                :src="
                    require(`@/assets/image/field_type_icon/people_single.svg`)
                "
                alt=""
                width="20px"
                height="20px"
                class="icon"
            />
            基本信息
            <img
                v-if="!isEdit"
                class="edit"
                :src="require(`@/assets/image/common/edit-active.svg`)"
                alt=""
                @click="edit"
            />
        </div>
        <el-form
            ref="form"
            size="small"
            class="basic-ui"
            :class="isEdit ? '' : 'not-edit'"
            :model="form"
            :rules="rules"
            :label-width="labelWidth"
        >
            <el-form-item label="头像:" class="avatar">
                <el-upload
                    class="avatar-uploader"
                    action="#"
                    :show-file-list="false"
                    :http-request="uploadAvatar"
                    :on-success="uploadSuccess"
                >
                    <!-- <img :src="getAvatar(form.avatar)" class="avatar" /> -->
                    <el-avatar
                        icon="el-icon-user-solid"
                        size="small"
                        class="avatar"
                        :src="getAvatar(form.avatar)"
                    ></el-avatar>
                    <div class="div-hover">
                        <span class="img-font">更改</span>
                    </div>
                </el-upload>
            </el-form-item>
            <el-form-item v-if="!isEdit" prop="username" label="用户名:">
                <span>{{ form.username || "--" }}</span>
            </el-form-item>
            <el-form-item prop="nike" label="昵称:">
                <span v-if="!isEdit">{{ form.nike || "--" }}</span>
                <el-input
                    v-else
                    v-model="form.nike"
                    class="basic-ui width224"
                    type="text"
                    size="small"
                    maxlength="25"
                    show-word-limit
                    placeholder="请输入昵称"
                ></el-input>
            </el-form-item>
            <!-- <el-form-item prop="email" label="邮箱:">
                <span v-if="!isEdit">{{ form.email || "--" }}</span>
                <el-input
                    v-else
                    class="basic-ui width224"
                    type="text"
                    placeholder="请输入邮箱"
                    show-word-limit
                    v-model="form.email"
                ></el-input>
            </el-form-item> -->
            <el-form-item prop="phone" label="电话:">
                <span v-if="!isEdit">{{ form.phone || "--" }}</span>
                <el-input
                    v-else
                    class="basic-ui width224"
                    type="text"
                    placeholder="请输入电话"
                    show-word-limit
                    v-model="form.phone"
                ></el-input>
            </el-form-item>
            <el-form-item v-if="isEdit">
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
                    @click="updateNike"
                    >确定</el-button
                >
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/personal_center";
import { imgHost } from "@/assets/tool/const";
export default {
    computed: {
        userInfo() {
            return this.$store.getters.userInfo;
        },
        getAvatar() {
            return (avatar) => {
                if (avatar) {
                    return `${imgHost}${avatar}`;
                }
                return require(`@/assets/image/common/default_avatar_big.png`);
            };
        },
        labelWidth() {
            return this.isEdit ? "64px" : "66px";
        }
    },
    data() {
        // 电话号码校验
        var validateMobilePhone = (rule, value, callback) => {
            if (value === "") {
                callback(new Error("请输入电话"));
            } else {
                if (value !== "") {
                    var reg = /^1[3456789]\d{9}$/;
                    if (!reg.test(value)) {
                        callback(new Error("请输入有效的电话号码"));
                    }
                }
                callback();
            }
        };
        // 邮箱校验
        var validateEmail = (rule, value, callback) => {
            if (value === "") {
                callback(new Error("请输入邮箱"));
            } else {
                if (value !== "") {
                    var reg =
                        /^[A-Za-z0-9\u4e00-\u9fa5._-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/;
                    // var reg =
                    // /^([a-zA-Z0-9]+[_|/_|/.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|/_|/.]?)*[a-zA-Z0-9]+[/.][a-zA-Z]{2,3}$/;
                    if (!reg.test(value)) {
                        callback(new Error("请输入有效的邮箱"));
                    }
                }
                callback();
            }
        };
        return {
            form: {
                avatar: "",
                nike: "",
                phone: "",
                email: "",
                username: ""
            },
            infoList: [
                {
                    label: "头像",
                    value: ""
                },
                {
                    label: "昵称",
                    value: "胡思睿"
                }
            ],
            rules: {},
            rules2: {
                nike: [
                    {
                        required: true,
                        trigger: "blur",
                        message: "请输入昵称"
                    }
                ],
                phone: [
                    {
                        required: true,
                        trigger: "blur",
                        message: "请输入电话"
                    },
                    { validator: validateMobilePhone, trigger: "blur" }
                ],
                email: [
                    {
                        required: true,
                        trigger: "blur",
                        message: "请输入邮箱"
                    },
                    { validator: validateEmail, trigger: "blur" }
                ]
            },
            isEdit: false
        };
    },
    mounted() {
        this.fetchPersonalUserInfo();
    },
    methods: {
        uploadAvatar(data) {
            var formData = new FormData();
            formData.append("avatar", data.file);
            api.updateUserAvatar(formData).then((res) => {
                if (res && res.resultCode === 200) {
                    this.fetchPersonalUserInfo();
                }
            });
        },
        uploadSuccess() {
            this.fetchPersonalUserInfo();
        },
        fetchPersonalUserInfo() {
            this.$store.dispatch("fetchUserInfo").then((res) => {
                this.setFormInfo(res);
            });
        },
        setFormInfo(res) {
            for (let key in this.form) {
                if (key === "nike") {
                    this.form[key] = res["full_name"] || "";
                } else {
                    this.form[key] = res[key] || "";
                }
            }
        },
        edit() {
            this.rules = _.cloneDeep(this.rules2);
            this.isEdit = true;
            this.$nextTick(() => {
                this.$refs.form.clearValidate();
            });
        },
        updateNike() {
            this.$refs.form.validate((flag) => {
                if (flag) {
                    let params = {
                        full_name: this.form.nike,
                        email: this.form.email,
                        phone: this.form.phone
                    };
                    if (this.form.nike) {
                        api.updateUserInfo(params).then((res) => {
                            if (res && res.resultCode === 200) {
                                this.$store.dispatch("fetchUserInfo");
                                this.isEdit = false;
                                this.rules = {};
                            }
                        });
                    }
                }
            });
        },
        cancel() {
            this.setFormInfo(this.userInfo);
            this.$refs.form.clearValidate();
            this.isEdit = false;
            this.rules = {};
        }
    }
};
</script>

<style lang="scss" scoped>
.label {
    display: inline-block;
    width: 80px;
}
.value {
    display: inline-block;
}
.width224 {
    width: 224px;
}

.div-hover {
    position: absolute;
    height: 64px;
    width: 64px;
    top: 0;
    opacity: 0.6;
    display: none;
    background-color: #8e919a;
    .img-font {
        color: #fff;
        bottom: 0;
        position: relative;
        top: 40%;
    }
}
.basic-ui.not-edit.el-form .el-form-item {
    margin-bottom: 10px;
}
.basic-ui.el-form .el-form-item .el-form-item__content {
    color: #2a3048 !important;
}
.avatar-uploader ::v-deep .el-upload {
    width: 64px;
    height: 64px;
    border-radius: 50%;
    cursor: pointer;
    position: relative;
    overflow: hidden;
}
.avatar-uploader ::v-deep .el-upload:hover {
    .div-hover {
        display: block;
    }
}
.avatar-uploader ::v-deep .el-icon-plus:before {
    display: none;
}
.avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 64px;
    height: 64px;
    line-height: 64px;
    text-align: center;
}
.avatar {
    width: 64px;
    height: 64px;
    display: block;
}
::v-deep .el-form-item__label {
    padding: 0 16px 0 0;
}
::v-deep .el-form-item.avatar {
    .el-form-item__label {
        position: relative;
        top: 16px;
    }
    .el-form-item__content {
        line-height: 1;
    }
}

.el-upload:hover {
    span.font {
        display: inline-block;
    }
}
.edit {
    position: relative;
    top: 2px;
    height: 16px;
    width: 16px;
    margin-left: 8px;
    cursor: pointer;
}
</style>
