<template>
    <el-dialog
        title=""
        :visible.sync="show"
        append-to-body
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :show-close="false"
        id="global-jump-dialog-id"
        class="basic-ui global-jump-dialog"
    >
        <img :src="img" alt="" width="182px" height="136px" />
        <div class="msg">{{ displayMsg }}</div>
        <span slot="footer" class="dialog-footer">
            <el-button
                class="basic-ui"
                size="small"
                type="primary"
                @click="onConfirm"
                >{{ btnText }}</el-button
            >
        </span>
    </el-dialog>
</template>

<script>
import router from "@/router";
export default {
    data() {
        return {
            show: false,
            msg: "",
            code: ""
        };
    },
    computed: {
        btnText() {
            return this.code === 401 ? "我知道了" : "返回首页";
        },
        displayMsg() {
            return this.code === 401
                ? "当前页面token已过期,请退出重新登录哦!"
                : this.msg;
        },
        img() {
            return this.code === 401
                ? require(`@/assets/image/common/jump_dialog/token_expired.png`)
                : require(`@/assets/image/common/jump_dialog/no_auth.png`);
        }
    },
    methods: {
        onConfirm() {
            this.show = false;
            if (this.code === 401) {
                router.push({
                    name: "login"
                });
            } else {
                router.push({
                    name: "myHomePage"
                });
            }
        }
    }
};
</script>

<style lang="scss">
.basic-ui.global-jump-dialog {
    .el-dialog {
        width: 480px;
        border-radius: 4px;
        .el-dialog__header {
            padding: 24px 32px;
            font-size: 16px;
            color: #2f384c;
        }
        .el-dialog__body {
            padding: 0 32px;
            text-align: center;
            .msg {
                margin: 20px 0 32px;
                color: #5c6479;
            }
        }
        .el-dialog__footer {
            padding: 0 32px 32px;
            text-align: center;
        }
        .el-dialog__headerbtn {
            top: 24px;
            right: 32px;
        }
    }
}
</style>
