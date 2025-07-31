<template>
    <div class="no-data">
        <div class="box">
            空间信息待设计
            <div>空间名称： {{ spaceInfo.name }}</div>
            <div>空间描述： {{ spaceInfo.desc }}</div>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        spaceInfo: {
            type: Object,
            default: () => {},
        },
    },
}
</script>

<style lang="scss" scoped>
.no-data {
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
