<template>
    <div class="login-page">
        <!-- <div class="logo">
            <img
                :src="require('@/assets/image/common/a_logo.png')"
                alt=""
                :width="32"
                :height="32"
            />
            Marktable
        </div> -->
        <div class="login-box">
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
                MarkTable
                是一款高度灵活的项目管理平台，通过多维自定义配置（如流程、权限、视图、报表），支持企业快速搭建适配不同业务场景的数字化工作流，涵盖研发管理、客户成功、项目协作等领域，助力团队实现跨部门高效协同与流程标准化
            </p>
            <div class="title">登录</div>
            <el-form
                class="basic-ui"
                ref="logonForm"
                size="small"
                :model="logonForm"
                :rules="rules"
            >
                <el-form-item prop="username">
                    <div slot="label">用户名</div>
                    <el-input
                        class="basic-ui"
                        prefix-icon="username-icon"
                        v-model="logonForm.username"
                        type="text"
                        placeholder="请输入用户名"
                    ></el-input>
                </el-form-item>
                <el-form-item prop="pwd">
                    <div slot="label">密码</div>
                    <el-input
                        class="basic-ui"
                        ref="pwdInput"
                        prefix-icon="pwd-icon"
                        suffix-icon="pwd-eye-icon"
                        v-model="logonForm.pwd"
                        :type="pwdType"
                        placeholder="请输入登录密码"
                    >
                        <el-image
                            slot="suffix"
                            class="img-sty"
                            :src="getIconUrl"
                            fit="scale-down"
                            @click.stop="changeEye"
                        />
                    </el-input>
                </el-form-item>
                <el-form-item>
                    <el-button
                        class="basic-ui width100"
                        type="primary"
                        size="small"
                        @click="login"
                        >登录</el-button
                    >
                </el-form-item>
            </el-form>
            <div class="click-to-register">
                <span class="pwd" @click="onResetPwd"> 忘记密码 </span>
                <span>
                    <span class="font">没有账号？</span>
                    <span class="reg" @click="onbtnClick">点击注册</span>
                </span>
            </div>
        </div>
    </div>
</template>

<script>
import api from "@/common/api";
export default {
    data() {
        return {
            title: "登录页",
            logonForm: {
                username: "",
                pwd: ""
            },
            pwdType: "password",
            rules: {
                username: [
                    {
                        required: true,
                        trigger: "blur",
                        message: "请输入用户名"
                    }
                ],
                pwd: [
                    {
                        required: true,
                        trigger: "blur",
                        message: "密码不能为空"
                    }
                ]
            }
        };
    },
    computed: {
        getIconUrl() {
            if (this.pwdType === "text") {
                return require(`@/assets/image/common/could-see.svg`);
            } else {
                return require(`@/assets/image/progress/hide-show.png`);
            }
        }
    },
    mounted() {},
    methods: {
        onbtnClick() {
            this.$router.push({ name: "register" });
        },
        onResetPwd() {
            this.$router.push({ name: "resetPwd" });
        },
        userLogin() {
            let params = {
                username: this.logonForm.username,
                password: this.logonForm.pwd
            };
            this.$store.dispatch("doLogin", params).then((res) => {
                let path = this.$route.query;
                if (Object.keys(path).length) {
                    this.$router.push({ path: path.redirect });
                } else {
                    this.$router.push({ name: "home" });
                }
            });
        },
        changeEye() {
            this.pwdType = this.pwdType === "text" ? "password" : "text";
            this.$refs.pwdInput.focus();
        },
        login() {
            this.$refs.logonForm.validate((flag) => {
                if (flag) {
                    this.userLogin();
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.login-page {
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
.test-account-block {
    font-size: 14px;
    color: #82889e;
    height: 32px;
    line-height: 32px;
    background-color: #f9fbff;
    border-radius: 4px;
    border: 1px solid #cdd5e6;
    display: flex;
    padding: 0 12px;
    margin-bottom: 16px;
    .account {
        margin-right: 24px;
    }
    .passwod {
    }
}
.login-box {
    position: relative;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 500px;
    text-align: center;
    box-sizing: border-box;
    padding: 40px 90px;
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
        height: 32.5px;
        font-size: 24px;
        font-family: MiSans, MiSans;
        font-weight: 500;
        color: #2f384c;
        margin-bottom: 16px;
    }
    p {
        font-size: 12px;
        color: #82889e;
        margin: 0 0 24px;
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
    ::v-deep .el-input__icon.username-icon:after {
        position: relative;
        top: -1px;
        content: "";
        width: 20px;
        height: 20px;
        background-image: url("~@/assets/image/field_type_icon/people_single.svg");
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
        // content: url('~@/assets/image/common/pwd.png');
        // position: relative;
        // top: 4px;
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
            width: 20px;
            height: 20px;
            background-size: 100% 100%;
        }
    }
    .click-to-register {
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
