<template>
    <div id="app" v-cloak>
        <router-view></router-view>
        <!-- 添加全局公用el-tooltip -->
        <el-tooltip
            popper-class="overflow-tip"
            placement="top"
            ref="tooltip"
            :content="tooltipContent"
        ></el-tooltip>
        <!-- at-tip -->
        <!-- <el-tooltip placement="top" ref="AtTooltip" :content="tooltipContent">
            <div slot="content">测试user-tip</div>
        </el-tooltip> -->
    </div>
</template>

<script>
export default {
    name: "App",
    data() {
        return {
            // 为el-tooltip设置content值
            tooltipContent: "",
        };
    },
};
</script>

<style lang="scss">
#app {
    font-family: MiSans, MiSans;
    color: var(--FONT_COLOR);
    font-size: 14px;
    height: 100vh;
    width: 100vw;
}
[v-cloak] {
    display: none !important;
}
/**
 * App.vue中
 * 网址、@功能
 **/
.mention-link,
.mention-at {
    display: inline-block;
    box-sizing: border-box;
    color: #1890ff;
    margin: 2px;
    background-color: #e3f0fc;
    border-radius: 4px;
    padding: 0 4px;
    height: 20px;
    line-height: 18px;
    font-size: 14px;
}
.el-tooltip__popper.overflow-tip {
    max-width: 480px;
}
</style>
