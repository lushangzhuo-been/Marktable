<template>
    <div>
        <div class="sub-title">
            <img
                :src="require(`@/assets/image/common/lock.svg`)"
                alt=""
                width="20px"
                height="20px"
                class="icon"
            />
            修改密码
        </div>
        <el-form
            ref="modifyPwdForm"
            :model="modifyPwdForm"
            :rules="rules"
            label-width="88px"
            size="small"
        >
            <el-form-item prop="oldPwd" label="当前密码:">
                <el-input
                    class="basic-ui width224"
                    ref="oldPwdType"
                    suffix-icon="pwd-eye-icon"
                    v-model="modifyPwdForm.oldPwd"
                    :type="oldPwdType"
                    placeholder="请输入当前密码"
                >
                    <el-image
                        slot="suffix"
                        class="img-sty"
                        :src="
                            getIconUrl(
                                oldPwdType === 'text'
                                    ? 'pwd_open_eye'
                                    : 'pwd_close_eye'
                            )
                        "
                        fit="scale-down"
                        @click.stop="changeEye(oldPwdType, 'oldPwdType')"
                    />
                </el-input>
            </el-form-item>
            <el-form-item prop="newPwd" label="新密码:">
                <el-popover
                    placement="bottom-start"
                    popper-class="pwd-popover"
                    width="228"
                    v-model="isFocus"
                    trigger="manual"
                    :visible-arrow="false"
                >
                    <div class="verify-item">
                        <img
                            :src="getRuleImg(index1)"
                            alt=""
                        />密码长度限制6～18位
                    </div>
                    <div class="verify-item">
                        <img :src="getRuleImg(index2)" alt="" />必须包含数字
                    </div>
                    <div class="verify-item">
                        <img :src="getRuleImg(index3)" alt="" />必须包含小写字母
                    </div>
                    <div class="verify-item">
                        <img :src="getRuleImg(index4)" alt="" />必须包含大写字母
                    </div>
                    <el-input
                        slot="reference"
                        class="basic-ui width224"
                        ref="newPwdType"
                        suffix-icon="pwd-eye-icon"
                        v-model="modifyPwdForm.newPwd"
                        :type="newPwdType"
                        placeholder="请输入新密码"
                        @focus="focus"
                        @blur="blur"
                    >
                        <el-image
                            slot="suffix"
                            class="img-sty"
                            :src="
                                getIconUrl(
                                    newPwdType === 'text'
                                        ? 'pwd_open_eye'
                                        : 'pwd_close_eye'
                                )
                            "
                            fit="scale-down"
                            @click.stop="changeEye(newPwdType, 'newPwdType')"
                        />
                    </el-input>
                </el-popover>
            </el-form-item>
            <el-form-item prop="checkNewPwd" label="确认密码:">
                <el-input
                    class="basic-ui width224"
                    ref="checkNewPwdType"
                    suffix-icon="pwd-eye-icon"
                    v-model="modifyPwdForm.checkNewPwd"
                    :type="checkNewPwdType"
                    placeholder="请确认新密码"
                >
                    <el-image
                        slot="suffix"
                        class="img-sty"
                        :src="
                            getIconUrl(
                                checkNewPwdType === 'text'
                                    ? 'pwd_open_eye'
                                    : 'pwd_close_eye'
                            )
                        "
                        fit="scale-down"
                        @click.stop="
                            changeEye(checkNewPwdType, 'checkNewPwdType')
                        "
                    />
                </el-input>
            </el-form-item>
            <el-form-item>
                <!-- <el-button class="basic-ui width100" size="small" @click="reset"
                    >重置</el-button
                > -->
                <el-button
                    class="basic-ui width100"
                    type="primary"
                    size="small"
                    @click="updatePwd"
                    >确定</el-button
                >
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
import api from "@/common/api/module/personal_center";
export default {
    computed: {
        getIconUrl() {
            return function (name) {
                // return require(`@/assets/image/common/${name}.png`)
                if (name === "pwd_open_eye") {
                    return require(`@/assets/image/common/could-see.svg`);
                } else {
                    return require(`@/assets/image/progress/hide-show.png`);
                }
            };
        },
        getRuleImg() {
            return function (index) {
                return require(`@/assets/image/common/${
                    index ? "verify_success" : "verify_error"
                }.png`);
            };
        }
    },
    data() {
        var validatePass = (rule, value, callback) => {
            if (value === "") {
                callback(new Error("请输入密码"));
            } else {
                const pattern2 = /^(?=.*?\d)/; //必须包含数字
                const pattern3 = /^(?=.*?[a-z])/; //必须包含小写写字母
                const pattern4 = /^(?=.*?[A-Z])/; //必须包含大写字母
                if (value.length > 5 && value.length < 19) {
                    //必须6-18个字符
                    this.index1 = true;
                } else {
                    this.index1 = false;
                }
                if (!value.match(pattern2)) {
                    this.index2 = false;
                } else {
                    this.index2 = true;
                }
                if (!value.match(pattern3)) {
                    this.index3 = false;
                } else {
                    this.index3 = true;
                }
                if (!value.match(pattern4)) {
                    this.index4 = false;
                } else {
                    this.index4 = true;
                }
                if (
                    !this.index1 ||
                    !this.index2 ||
                    !this.index3 ||
                    !this.index4
                ) {
                    callback(
                        new Error("密码必须包含数字、大小写字母、长度6-18位")
                    );
                } else {
                    if (this.modifyPwdForm.checkPwd !== "") {
                        this.$refs.modifyPwdForm.validateField("checkPwd");
                    }
                    callback();
                }
            }
        };
        var validatePass2 = (rule, value, callback) => {
            if (value === "") {
                callback(new Error("请确认新密码"));
            } else if (value !== this.modifyPwdForm.newPwd) {
                callback(new Error("两次输入密码不一致!"));
            } else {
                callback();
            }
        };
        return {
            isFocus: false,
            index1: false,
            index2: false,
            index3: false,
            index4: false,
            modifyPwdForm: {
                oldPwd: "",
                newPwd: "",
                checkNewPwd: ""
            },
            oldPwdType: "password",
            newPwdType: "password",
            checkNewPwdType: "password",
            rules: {
                oldPwd: [
                    {
                        required: true,
                        trigger: "blur",
                        message: "请输入当前密码"
                    }
                ],
                newPwd: [
                    {
                        required: true,
                        trigger: "blur",
                        validator: validatePass
                    }
                ],
                checkNewPwd: [
                    {
                        required: true,
                        trigger: "blur",
                        validator: validatePass2
                    }
                ]
            }
        };
    },
    mounted() {
        this.reset();
    },
    watch: {
        "modifyPwdForm.newPwd"(pwd) {
            const pattern2 = /^(?=.*?\d)/; //必须包含数字
            const pattern3 = /^(?=.*?[a-z])/; //必须包含小写写字母
            const pattern4 = /^(?=.*?[A-Z])/; //必须包含大写字母

            if (pwd.length > 5 && pwd.length < 19) {
                //必须6-18个字符
                this.index1 = true;
            } else {
                this.index1 = false;
            }
            if (!pwd.match(pattern2)) {
                this.index2 = false;
            } else {
                this.index2 = true;
            }
            if (!pwd.match(pattern3)) {
                this.index3 = false;
            } else {
                this.index3 = true;
            }
            if (!pwd.match(pattern4)) {
                this.index4 = false;
            } else {
                this.index4 = true;
            }

            if (this.index1 && this.index2 && this.index3 && this.index4) {
                this.isFocus = false;
            } else {
                this.isFocus = true;
            }
        }
    },
    methods: {
        changePwd() {
            let params = {
                old_password: this.modifyPwdForm.oldPwd,
                new_password: this.modifyPwdForm.newPwd
            };
            api.updateUserPwd(params)
                .then((res) => {
                    if (res && res.resultCode === 200) {
                        this.$message({
                            showClose: true,
                            message: "密码修改成功,请重新登录",
                            type: "success"
                        });
                        setTimeout(() => {
                            this.$router.push({ path: "/login" });
                        }, 1000);
                    }
                })
                .catch((err) => {
                    this.isFocus = false;
                });
        },
        focus() {
            // if (this.index1) {
            //     this.isFocus = false
            if (this.index1 && this.index2 && this.index3 && this.index4) {
                this.isFocus = false;
            } else {
                this.isFocus = true;
            }
        },
        blur() {
            this.isFocus = false;
        },
        change() {
            if (this.index1 && this.index2 && this.index3 && this.index4) {
                this.$refs.newPwdType.focus();
                this.isFocus = false;
            } else {
                this.isFocus = true;
            }
        },
        pwdChange(val) {},
        changeEye(type, text) {
            this[`${text}`] = type === "text" ? "password" : "text";
            // this.$refs.text.focus()
        },
        updatePwd() {
            this.$refs.modifyPwdForm.validate((flag) => {
                if (flag) {
                    if (
                        this.index1 &&
                        this.index2 &&
                        this.index3 &&
                        this.index4
                    ) {
                        this.isFocus = false;
                        this.changePwd();
                    } else {
                        this.isFocus = true;
                    }
                    // 调接口
                }
            });
        },
        reset() {
            this.modifyPwdForm = {
                oldPwd: "",
                newPwd: "",
                checkNewPwd: ""
            };
            this.$refs.modifyPwdForm.clearValidate();
            this.$refs.newPwdType.blur();
            this.$nextTick(() => {
                this.isFocus = false;
            });
        }
    }
};
</script>

<style lang="scss" scoped>
// .content-title {
//     font-size: 20px;
//     font-family: MiSans, MiSans;
//     font-weight: 500;
//     margin-bottom: 30px;
// }
.width224 {
    width: 224px;
}
::v-deep .el-form-item__label {
    color: #171e31;
}
.img-sty ::v-deep .el-image__inner {
    vertical-align: middle;
    cursor: pointer;
    width: 20px;
    height: 20px;
    background-size: 100% 100%;
}
</style>
