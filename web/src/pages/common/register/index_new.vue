<template>
    <div class="register-page">
        <div class="register-box">
            <div class="name">
                <!-- <img
                    :src="require('@/assets/image/common/a_logo_2.png')"
                    alt=""
                    :width="46"
                    :height="50"
                />
                <span>项目管理</span> -->
                <!-- <img
                    :src="require('@/assets/image/common/a_logo.png')"
                    alt=""
                    :width="48"
                    :height="48"
                />
                <span>Marktable</span> -->
                <b class="logo-name"></b>
            </div>
            <p>
                Marktable
                是一款高度灵活的项目管理平台，通过多维自定义配置（如流程、权限、视图、报表），支持企业快速搭建适配不同业务场景的数字化工作流，涵盖研发管理、客户成功、项目协作等领域，助力团队实现跨部门高效协同与流程标准化
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
                <el-form-item prop="username">
                    <div slot="label">用户名</div>
                    <el-input
                        class="basic-ui"
                        prefix-icon="username-icon"
                        v-model="registerForm.username"
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
                <!-- <el-form-item prop="checkPwd">
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
                </el-form-item> -->
                <el-form-item prop="email">
                    <div slot="label">邮箱</div>
                    <div
                        style="
                            position: relative;
                            display: flex;
                            align-items: center;
                            justify-content: space-between;
                        "
                    >
                        <el-input
                            class="basic-ui width68"
                            prefix-icon="email-icon"
                            v-model="registerForm.email"
                            type="text"
                            placeholder="请输入邮箱"
                        ></el-input>
                        <el-button
                            class="basic-ui margin-left"
                            type="primary"
                            size="small"
                            :disabled="!isCodeBtnShow || second !== totalSecond"
                            @click="getValideCode"
                            >{{
                                second === totalSecond
                                    ? "获取验证码"
                                    : second + `秒后重新发送`
                            }}</el-button
                        >
                    </div>
                    <!-- <el-button
                        style="
                            position: absolute;
                            right: -96px;
                            top: 0;
                            width: 92px;
                        "
                        class="basic-ui"
                        type="primary"
                        size="small"
                        disabled
                        @click="login"
                        >60s</el-button
                    > -->
                </el-form-item>
                <el-form-item prop="code">
                    <div slot="label">验证码</div>
                    <el-input
                        class="basic-ui"
                        prefix-icon="pwd-icon"
                        v-model="registerForm.code"
                        type="text"
                        placeholder="请输入验证码"
                    ></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button
                        class="basic-ui width100"
                        type="primary"
                        size="small"
                        @click="registerUsername"
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
                    if (this.registerForm.checkPwd !== "") {
                        this.$refs.registerForm.validateField("checkPwd");
                    }
                    callback();
                }
            }
        };
        // var validatePass2 = (rule, value, callback) => {
        //     if (value === "") {
        //         callback(new Error("请再次输入密码"));
        //     } else if (value !== this.registerForm.pwd) {
        //         callback(new Error("两次输入密码不一致!"));
        //     } else {
        //         callback();
        //     }
        // };
        return {
            isCodeBtnShow: false,
            totalSecond: 60, // 总秒数
            second: 60, // 倒计时的秒数
            timer: null, // 定时器 id
            isFocus: false,
            index1: false,
            index2: false,
            index3: false,
            index4: false,
            title: "注册",
            registerPwdType: "password",
            // registerCheckPwdType: "password",
            registerForm: {
                username: "",
                pwd: "",
                // checkPwd: "",
                email: "",
                code: ""
            },
            registerRules: {
                username: [
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
                        trigger: ["blur"],
                        validator: validatePass
                    }
                ],
                email: [
                    {
                        required: true,
                        trigger: "change",
                        message: "请输入新邮箱"
                    },
                    {
                        pattern:
                            /^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9_\.\-])+\.)+([a-zA-Z0-9]{2,4})+$/,
                        message: "请输入正确的邮箱格式",
                        trigger: "blur"
                    }
                ],
                // checkPwd: [
                //     {
                //         required: true,
                //         trigger: ["change", "blur"],
                //         validator: validatePass2
                //     }
                // ],

                code: [
                    {
                        required: true,
                        trigger: "blur",
                        message: "请输入验证码"
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
        },
        "registerForm.email"(val) {
            let parttern =
                /^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9_\.\-])+\.)+([a-zA-Z0-9]{2,4})+$/;
            if (val && parttern.test(val)) {
                this.isCodeBtnShow = true;
            } else {
                this.isCodeBtnShow = false;
            }
        }
    },
    mounted() {},
    methods: {
        getValideCode() {
            api.getCodeForRegister({ email: this.registerForm.email }).then(
                (res) => {
                    if (res && res.resultCode === 200) {
                        this.$message({
                            showClose: true,
                            message: "验证码已发送至邮箱中，请查收！",
                            type: "success"
                        });
                        this.timer = setInterval(() => {
                            this.second--;
                            if (this.second <= 0) {
                                clearInterval(this.timer);
                                this.timer = null; // 重置定时器 id
                                this.second = this.totalSecond; // 归位
                            }
                        }, 1000);
                    } else {
                        this.timer = null;
                        this.second = this.totalSecond = 60; // 归位
                    }
                }
            );
        },
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
            this[text] = type === "text" ? "password" : "text";
        },

        registerUsername() {
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
                            username: this.registerForm.username,
                            password: this.registerForm.pwd,
                            email: this.registerForm.email,
                            code: this.registerForm.code
                        };
                        api.registerV2(params).then((res) => {
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
    .margin-left {
        margin-left: 8px;
    }
    .width100 {
        width: 100%;
    }
    .width68 {
        width: 68%;
    }
    .register-box {
        position: relative;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        width: 548px;
        text-align: center;
        box-sizing: border-box;
        padding: 40px 112px;
        border-radius: 8px;
        background-color: #fff;
        .name {
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 40px;
            margin-bottom: 20px;
            height: 48px;
            font-family: SF Compact Display Semibold;
            img {
                margin-right: 12px;
            }
            .logo-name {
                display: inline-block;
                // width: 1066px;
                // height: 224px;
                width: 210px;
                height: 44px;
                background-image: url("@/assets/image/logo_name.png");
                background-size: 100% 100%;
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
            float: inherit;
            text-align: left;
            line-height: 1;
            color: #40485f;
            div {
                margin-bottom: -10px;
            }
        }
        ::v-deep
            .el-form-item.is-required:not(.is-no-asterisk)
            > .el-form-item__label:before {
            display: none;
        }
        ::v-deep .el-input__icon.username-icon:after {
            position: relative;
            top: -1px;
            content: "";
            width: 20px;
            height: 20px;
            background-image: url("~@/assets/image/field_type_icon/people_single.svg");
            background-size: 100% 100%;
        }
        ::v-deep .el-input__icon.email-icon:after {
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
}
</style>
