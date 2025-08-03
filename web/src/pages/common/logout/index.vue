<template>
    <div class="register-page">
        <!-- <div class="logo">
            <img
                :src="require('@/assets/image/common/a_logo.png')"
                alt=""
                :width="32"
                :height="32"
            />
            Marktable
        </div> -->
        <div class="register-box">
            <div class="name">
                <!-- <img
                    :src="require('@/assets/image/common/a_logo_2.png')"
                    alt=""
                    :width="46"
                    :height="50"
                />
                <span>项目管理</span> -->
                <img
                    :src="require('@/assets/image/common/a_logo.png')"
                    alt=""
                    :width="32"
                    :height="32"
                />
                Marktable
            </div>
            <p>
                Marktable系统简单的介绍Marktable系统简单的介绍Marktable系统简单的介绍
            </p>
            <div class="title">注册</div>
            <!-- <el-steps
                    :space="200"
                    :active="activeStep"
                    align-center
                    finish-status="success"
                >
                    <el-step
                        v-for="(item, index) in registerSteps"
                        :key="index"
                        :title="item.title"
                    ></el-step>
                </el-steps> -->
            <el-form
                class="basic-ui"
                ref="registerForm"
                size="small"
                :model="registerForm"
                :rules="registerRules"
            >
                <el-form-item prop="account">
                    <div slot="label">用户名</div>
                    <el-input
                        class="basic-ui"
                        prefix-icon="account-icon"
                        v-model="registerForm.account"
                        type="text"
                        placeholder="请输入用户名"
                    ></el-input>
                </el-form-item>
                <el-form-item prop="pwd">
                    <div slot="label">密码</div>
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
                            <img
                                :src="getRuleImg(index3)"
                                alt=""
                            />必须包含小写字母
                        </div>
                        <div class="verify-item">
                            <img
                                :src="getRuleImg(index4)"
                                alt=""
                            />必须包含大写字母
                        </div>
                        <el-input
                            slot="reference"
                            class="basic-ui"
                            ref="pwdInput"
                            prefix-icon="pwd-icon"
                            suffix-icon="pwd-eye-icon"
                            v-model="registerForm.pwd"
                            :type="registerPwdType"
                            placeholder="请输入密码"
                            @focus="focus"
                            @blur="blur"
                        >
                            <el-image
                                slot="suffix"
                                class="img-sty"
                                :src="
                                    getIconUrlRegister(
                                        registerPwdType === 'text'
                                            ? 'pwd_open_eye'
                                            : 'pwd_close_eye'
                                    )
                                "
                                fit="scale-down"
                                @click="
                                    changeEye(
                                        registerPwdType,
                                        'registerPwdType'
                                    )
                                "
                            />
                        </el-input>
                    </el-popover>
                </el-form-item>
                <el-form-item prop="checkPwd">
                    <div slot="label">确认密码</div>
                    <el-input
                        class="basic-ui"
                        ref="checkPwdInput"
                        prefix-icon="pwd-icon"
                        suffix-icon="pwd-eye-icon"
                        v-model="registerForm.checkPwd"
                        :type="registerCheckPwdType"
                        placeholder="请再次输入密码"
                    >
                        <el-image
                            slot="suffix"
                            class="img-sty"
                            :src="
                                getIconUrlRegister(
                                    registerCheckPwdType === 'text'
                                        ? 'pwd_open_eye'
                                        : 'pwd_close_eye'
                                )
                            "
                            fit="scale-down"
                            @click="
                                changeEye(
                                    registerCheckPwdType,
                                    'registerCheckPwdType'
                                )
                            "
                        />
                    </el-input>
                </el-form-item>
                <el-form-item>
                    <el-button
                        class="basic-ui width100"
                        type="primary"
                        size="small"
                        @click="register"
                        >注册</el-button
                    >
                </el-form-item>
            </el-form>
            <div class="click-to-login">
                <span class="pwd">
                    <!-- 忘记密码？ -->
                </span>
                <span>
                    <span class="font">已有账号？</span>
                    <span class="reg" @click="onbtnClick">点击登录</span>
                </span>
            </div>
        </div>
    </div>
</template>

<script>
import api from "@/common/api";
export default {
    data() {
        var validatePass = (rule, value, callback) => {
            if (value === "") {
                callback(new Error("请输入密码"));
            } else {
                if (this.registerForm.checkPwd !== "") {
                    this.$refs.registerForm.validateField("checkPwd");
                }
                callback();
            }
        };
        var validatePass2 = (rule, value, callback) => {
            if (value === "") {
                callback(new Error("请再次输入密码"));
            } else if (value !== this.registerForm.pwd) {
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
            title: "注册",
            registerPwdType: "password",
            registerCheckPwdType: "password",
            registerForm: {
                account: "",
                pwd: "",
                checkPwd: ""
            },
            registerRules: {
                account: [
                    {
                        required: true,
                        min: 4,
                        max: 18,
                        trigger: "blur",
                        message: "请输入4到18位用户名"
                    }
                    // {
                    //     min: 4,
                    //     max: 18,
                    //     message: '长度在4到18个字符',
                    //     trigger: 'blur',
                    // },
                ],

                pwd: [
                    {
                        required: true,
                        trigger: ["change", "blur"],
                        validator: validatePass
                    }
                ],
                checkPwd: [
                    {
                        required: true,
                        trigger: ["change", "blur"],
                        validator: validatePass2
                    }
                ]
            },
            registerSteps: [
                {
                    title: "邮箱验证"
                },
                {
                    title: "设置密码"
                },
                {
                    title: "设置完成"
                }
            ],
            activeStep: 1
        };
    },
    computed: {
        getIconUrlRegister() {
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
    watch: {
        "registerForm.pwd"(pwd) {
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
    mounted() {},
    methods: {
        // 点击注册/登录按钮
        onbtnClick() {
            this.$router.push({ name: "login" });
        },
        focus() {
            if (this.index1 && this.index2 && this.index3 && this.index4) {
                this.isFocus = false;
            } else {
                this.isFocus = true;
            }
        },
        blur() {
            this.isFocus = false;
        },

        changeEye(type, text) {
            this[`${text}`] = type === "text" ? "password" : "text";
            // this.$refs.text.focus()
        },

        register() {
            this.$refs.registerForm.validate((flag) => {
                if (flag) {
                    if (
                        this.index1 &&
                        this.index2 &&
                        this.index3 &&
                        this.index4
                    ) {
                        this.isFocus = false;
                        let params = {
                            username: this.registerForm.account,
                            password: this.registerForm.pwd
                        };
                        api.register(params).then((res) => {
                            if (res && res.resultCode === 200) {
                                this.$message({
                                    showClose: true,
                                    message: "注册成功，请登录",
                                    type: "success"
                                });
                                this.$router.push({ name: "login" });
                            }
                        });
                    } else {
                        this.isFocus = true;
                    }
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.register-page {
    height: 100%;
    background-color: #f2f3f5;
    .logo {
        display: flex;
        align-items: center;
        padding-left: 16px;
        height: 66px;
        font-family: DingTalk JinBuTi-Regular;
        font-size: 24px;
        img {
            margin-right: 8px;
        }
    }
}
.width100 {
    width: 100%;
}
.logo {
    position: absolute;
    top: 0;
    left: 0;
}
.register-box {
    position: relative;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 500px;
    text-align: center;
    box-sizing: border-box;
    padding: 60px 90px;
    border-radius: 8px;
    background-color: #fff;
    .name {
        display: flex;
        justify-content: center;
        align-items: center;
        font-size: 40px;
        margin-bottom: 20px;
        font-family: DingTalk JinBuTi-Regular;
        img {
            margin-right: 12px;
        }
    }
    .title {
        font-size: 24px;
        font-family: MiSans, MiSans;
        font-weight: 500;
        color: #2f384c;
        margin-bottom: 16px;
    }
    p {
        font-size: 12px;
        color: var(--GRAY_FONT_COLOR);
        margin: 0 0 20px;
        text-align: left;
    }
    .el-form-item {
        margin-bottom: 16px;
    }
    ::v-deep .el-form-item--small .el-form-item__content {
        line-height: 1;
        margin-bottom: 0;
    }
    ::v-deep .el-form-item--small .el-form-item__label {
        line-height: 1;
        margin-bottom: 8px;
        color: #40485f;
    }
    ::v-deep
        .el-form-item.is-required:not(.is-no-asterisk)
        > .el-form-item__label:before {
        display: none;
    }
    ::v-deep .el-input__icon.account-icon:after {
        position: relative;
        top: -1px;
        content: "";
        width: 20px;
        height: 20px;
        background-image: url("~@/assets/image/common/email.png");
        background-size: 100% 100%;
    }
    ::v-deep .el-input__icon.pwd-icon:after {
        position: relative;
        top: -1px;
        content: "";
        width: 20px;
        height: 20px;
        background-image: url("~@/assets/image/common/lock.png");
        background-size: 100% 100%;
    }
    .el-input {
        cursor: pointer;
    }
    .img-sty {
        line-height: 32px;
        ::v-deep .el-image__inner {
            position: relative;
            top: -1px;
            vertical-align: middle;
            cursor: pointer;
            width: 20px;
            height: 20px;
            background-size: 100% 100%;
        }
    }
    .click-to-login {
        display: flex;
        justify-content: space-between;
        height: 16px;
        margin-top: 16px;
        font-size: 12px;
        text-align: right;
        .font {
            color: var(--GRAY_FONT_COLOR);
        }
        .reg,
        .pwd {
            color: var(--PRIMARY_COLOR);
            cursor: pointer;
        }
    }
}
</style>
