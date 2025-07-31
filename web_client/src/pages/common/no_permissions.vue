<template>
    <div class="no-permissions">
        <div class="box">
            <img
                :src="require(`@/assets/image/common/no_permissions.png`)"
                alt=""
                width="220px"
                height="132px"
            />
            <div>暂无权限</div>
            <div class="add">
                暂无权限 请联系管理员<span class="admin">{{ getAdmin }}</span
                >添加权限
            </div>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        admin: {
            type: String,
            default: ""
        }
    },
    computed: {
        getAdmin() {
            if (!this.admin) return "";
            if (typeof this.admin === String) {
                return `: ${this.admin}`;
            } else {
                return `: ${this.admin.join(",")}`;
            }
        }
    }
};
</script>

<style lang="scss" scoped>
.no-permissions {
    position: relative;
    height: 100%;
    width: 100%;
    .box {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        text-align: center;
        color: #a6abbc;
        font-size: 14px;
        .add {
            margin-top: 8px;
        }
    }
}
</style>
