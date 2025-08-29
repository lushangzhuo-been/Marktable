<template>
    <div
        class="no-data"
        v-loading="loading"
        :element-loading-text="loadingMessage"
        element-loading-background="rgba(255, 255, 255)"
    >
        <div class="box" v-show="imgShow">
            <img
                :src="require(`@/assets/image/common/no_data_1.png`)"
                alt=""
                width="220px"
                height="132px"
            />
            <div v-html="text"></div>
            <div v-if="isTextShow" class="add" v-html="mainText"></div>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        isTextShow: {
            type: Boolean,
            default: false
        },
        text: {
            type: String,
            default: "暂无数据"
        },
        mainText: {
            type: String,
            default: "如需添加，请点击左侧【+添加】按钮"
        },
        imgShow: {
            isTextShow: {
                type: Boolean,
                default: false
            }
        },
        loading: {
            isTextShow: {
                type: Boolean,
                default: true
            }
        },
        loadingMessage: {
            type: String,
            default: "加载中"
        }
    }
};
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
