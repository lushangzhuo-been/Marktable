<template>
    <el-tooltip
        class="item"
        effect="dark"
        placement="top-start"
        :popper-class="`tips-popper ${popperClass}`"
    >
        <div slot="content" v-html="tips"></div>
        <b class="tips"></b>
    </el-tooltip>
</template>

<script>
export default {
    props: {
        tips: {
            type: String,
            default: '',
        },
        placement: {
            type: String,
            default: 'top',
        },
        popperClass: {
            type: String,
            default: '',
        },
    },
}
</script>

<style lang="scss" scoped>
.tips {
    display: inline-block;
    width: 16px;
    height: 16px;
    margin-left: 8px;
    background-image: url('~@/assets/image/common/tip.png');
    background-size: 100% 100%;
    vertical-align: middle;
}
</style>
