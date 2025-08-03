<template>
    <div class="home">
        <left-menu ref="LeftMenu"></left-menu>
        <div class="main">
            <router-view @refresh-unread="refreshUnread"></router-view>
        </div>
    </div>
</template>

<script>
import LeftMenu from "./left_menu.vue";
export default {
    components: {
        LeftMenu,
    },
    methods: {
        refreshUnread() {
            this.$refs.LeftMenu.getUnReadNum();
        },
    },
};
</script>

<style lang="scss" scoped>
.home {
    height: 100vh;
    width: 100vw;
    display: flex;
    background-color: var(--color2);
    .left-menu {
        height: 100%;
        width: 256px;
        background-color: #f5f6f8;
    }
    .main {
        margin: 20px 0 4px 24px;
        padding: 0 16px 20px 0;
        height: calc(100% - 44px);
        width: calc(100% - 24px - 256px);
        overflow-y: auto;
        .title {
            font-size: 24px;
            padding-bottom: 20px;
        }
    }
}
</style>
