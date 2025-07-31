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
            重置邮箱
        </div>
        <el-form
            ref="resetEmailForm"
            :model="resetEmailForm"
            :rules="rules"
            label-width="88px"
            size="small"
        >
            <el-form-item prop="oldEmail" label="当前邮箱:">
                <el-input
                    class="basic-ui width224"
                    suffix-icon="pwd-eye-icon"
                    v-model="resetEmailForm.oldEmail"
                    type="text"
                    placeholder="请输入当前邮箱"
                    :disabled="true"
                >
                </el-input>
            </el-form-item>
            <el-form-item prop="newEmail" label="新邮箱:">
                <el-input
                    class="basic-ui width224"
                    suffix-icon="pwd-eye-icon"
                    v-model="resetEmailForm.newEmail"
                    type="text"
                    placeholder="请输入新邮箱"
                >
                </el-input>
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
            </el-form-item>
            <el-form-item v-if="codeShow" prop="code" label="验证码:">
                <el-input
                    class="basic-ui width224"
                    suffix-icon="pwd-eye-icon"
                    v-model="resetEmailForm.code"
                    type="text"
                    placeholder="请输入验证码"
                >
                </el-input>
            </el-form-item>
            <el-form-item>
                <el-button
                    class="basic-ui width100"
                    type="primary"
                    size="small"
                    @click="updateEmail"
                    >确定</el-button
                >
                <!-- :disabled="btnDisabled" -->
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
import api from "@/common/api/module/personal_center";
export default {
    computed: {
        btnDisabled() {
            if (this.resetEmailForm.newEmail && this.resetEmailForm.code) {
                return false;
            } else {
                return true;
            }
        }
    },
    data() {
        return {
            isCodeBtnShow: false,
            btnContent: "发送验证码",
            codeShow: false,
            totalSecond: 60, // 总秒数
            second: 60, // 倒计时的秒数
            timer: null, // 定时器 id
            resetEmailForm: {
                oldEmail: "",
                newEmail: "",
                code: ""
            },
            rules: {
                newEmail: [
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
                code: [
                    {
                        required: true,
                        trigger: "blur",
                        message: "请输入验证码"
                    }
                ]
            },
            userInfo: {}
        };
    },

    watch: {
        "resetEmailForm.newEmail"(val) {
            let parttern =
                /^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9_\.\-])+\.)+([a-zA-Z0-9]{2,4})+$/;
            if (val && parttern.test(val)) {
                this.isCodeBtnShow = true;
            } else {
                this.isCodeBtnShow = false;
            }
        }
    },
    mounted() {
        this.fetchPersonalUserInfo();
    },
    methods: {
        fetchPersonalUserInfo() {
            this.$store.dispatch("fetchUserInfo").then((res) => {
                this.userInfo = res;
                this.resetEmailForm.oldEmail = res.email;
            });
        },
        // 获取验证码
        getValideCode() {
            let params = {
                email: this.resetEmailForm.newEmail
            };
            api.resetEmailForCode(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.codeShow = true;
                    this.timer = setInterval(() => {
                        this.second--;
                        if (this.second <= 0) {
                            clearInterval(this.timer);
                            this.timer = null; // 重置定时器 id
                            this.second = this.totalSecond; // 归位
                        }
                    }, 1000);
                } else {
                    this.codeShow = false;
                }
            });
        },
        changeEmail() {
            let params = {
                email: this.resetEmailForm.newEmail,
                code: this.resetEmailForm.code
            };
            api.updateEmail(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.$message({
                        showClose: true,
                        message: "重置邮箱成功！",
                        type: "success"
                    });
                }
            });
        },

        updateEmail() {
            this.$refs.resetEmailForm.validate((flag) => {
                if (flag) {
                    this.changeEmail();
                }
            });
        }
    },
    destroyed() {
        clearInterval(this.timer);
    }
};
</script>

<style lang="scss" scoped>
.margin-left {
    margin-left: 8px;
}
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
