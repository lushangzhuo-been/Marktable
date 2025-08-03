<template>
    <el-tooltip
        effect="light"
        :popper-class="`basic-ui user-card-message ${myClass}`"
        :placement="position"
        :visible-arrow="false"
    >
        <div slot="content" class="card-content">
            <!-- 头像 -->
            <div class="avatar-block">
                <div class="content">
                    <el-avatar
                        class="user-msg-card-avatar"
                        icon="el-icon-user-solid"
                        :src="getAvatar(userMessage.avatar)"
                    ></el-avatar>
                    <div class="name">
                        {{ userMessage.full_name }}
                    </div>
                </div>
            </div>
            <div class="msg-block">
                <!-- 卡片信息
                {{ userMessage }} -->
                <div class="msg-item">
                    <div class="item-title">用户名：</div>
                    <div class="item-content">
                        {{ userMessage.username | emptySpace }}
                    </div>
                </div>
                <div class="msg-item">
                    <div class="item-title">昵称：</div>
                    <div class="item-content">
                        {{ userMessage.full_name | emptySpace }}
                    </div>
                </div>
                <div class="msg-item">
                    <div class="item-title">邮箱：</div>
                    <div
                        class="item-content"
                        :class="{ 'email-to': userMessage.email }"
                        @click="emailTo(userMessage.email)"
                    >
                        {{ userMessage.email | emptySpace }}
                    </div>
                    <b
                        v-if="userMessage.email"
                        @click="copyEmail(userMessage.email)"
                        class="operation-item-box copy"
                    ></b>
                </div>
                <div class="msg-item">
                    <div class="item-title">电话：</div>
                    <div class="item-content">
                        {{ userMessage.phone | emptySpace }}
                    </div>
                </div>
            </div>
        </div>
        <slot></slot>
    </el-tooltip>
</template>

<script>
import { imgHost } from "@/assets/tool/const";
import { emptySpace } from "@/assets/tool/func";
export default {
    name: "ellipsisTooltip",
    filters: {
        emptySpace
    },
    props: {
        userMessage: {
            type: Object,
            default: () => {}
        },
        position: {
            type: String,
            default: "bottom"
        },
        myClass: {
            type: String,
            default: ""
        }
    },
    computed: {},

    data() {
        return {};
    },
    watch: {},
    mounted() {},
    beforeDestroy() {},
    methods: {
        emailTo(email) {
            if (email) {
                const mailtoLink = `mailto:${email}`;
                window.location.href = mailtoLink;
            }
        },
        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        },
        copyEmail(email) {
            let tempTextarea = document.createElement("textarea");
            // 设置textarea的value为当前网址
            tempTextarea.value = email;
            // 将textarea添加到DOM中
            document.body.appendChild(tempTextarea);
            // 选中textarea中的文本
            tempTextarea.select();
            // 复制选中的文本
            document.execCommand("copy");
            // 移除临时的textarea元素
            document.body.removeChild(tempTextarea);
            this.$message({
                showClose: true,
                message: "复制成功",
                type: "success"
            });
        }
    }
};
</script>
<style lang="scss" scope>
.card-content {
    background-color: #ffffff;
    border-radius: 8px;
    .avatar-block {
        display: flex;
        justify-content: center;
        background-color: #f0f0f0;
        border-radius: 8px 8px 0 0;
        .content {
            height: 76px;
            box-sizing: border-box;
            display: inline-block;
            line-height: 52px;
            padding: 10px 0;
            margin: 0 auto;
            white-space: nowrap;
            border-radius: 0 0 8px 8px;
            .user-msg-card-avatar {
                width: 52px;
                height: 52px;
            }
            .name {
                margin-left: 14px;
                display: inline-block;
                line-height: 52px;
                vertical-align: top;
                font-size: 16px;
                color: #171e31;
                max-width: 120px;
                overflow: hidden;
                text-overflow: ellipsis;
            }
        }
    }
    .msg-block {
        box-sizing: border-box;
        padding: 16px 20px;
        height: 160px;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        .msg-item {
            white-space: nowrap;
            vertical-align: bottom;
            .item-title {
                display: inline-block;
                width: 56px;
                text-align: right;
                font-family: MiSans, MiSans;
                font-size: 14px;
                color: #82889e;
                vertical-align: bottom;
            }
            .item-content {
                display: inline-block;
                max-width: 180px;
                overflow: hidden;
                white-space: nowrap;
                text-overflow: ellipsis;
                font-size: 14px;
                vertical-align: bottom;
                &.email-to {
                    cursor: pointer;
                    &:hover {
                        color: #1890ff;
                    }
                }
            }
        }
    }
}
.operation-item-box {
    display: inline-block;
    width: 20px;
    height: 20px;
    background-size: 100% 100%;
    vertical-align: middle;
    cursor: pointer;
    &.copy {
        margin-left: 6px;
        background-image: url(@/assets/image/progress_column/copy.svg);
        &:hover {
            background-image: url(@/assets/image/progress_column/copy-active.svg);
        }
    }
}
</style>
<style lang="scss">
.el-tooltip__popper.user-card-message.is-light {
    box-sizing: border-box;
    border: none;
    min-width: 240px;
    max-width: 300px;
    padding: 0;
    border-radius: 8px;
    box-shadow: 0px 0px 10px 1px rgba(0, 0, 0, 0.2);
}
</style>
