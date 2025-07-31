<template>
    <div class="create-page-index">
        <div class="create-page__left">
            <div class="title">预览效果</div>
            <left-preview ref="leftPreview" :modules="modules"></left-preview>
        </div>
        <div class="create-page__right">
            <right-setting
                :modules="modules"
                @filed-change="filedChange"
            ></right-setting>
        </div>
    </div>
</template>
<script>
import RightSetting from "./right_setting.vue";
import LeftPreview from "./left_preview.vue";
export default {
    props: {
        modules: {
            type: String,
            default: ""
        }
    },
    components: {
        LeftPreview,
        RightSetting
    },
    data() {
        return {};
    },
    methods: {
        filedChange() {
            this.$refs.leftPreview.fetchCreatePage();
        }
    }
};
</script>
<style lang="scss" scoped>
.create-page-index {
    height: 800px;
    background-color: #fafafa;
    display: flex;
    .create-page__left {
        position: relative;
        box-sizing: border-box;
        width: 60%;
        padding: 20px 24px;
        display: flex;
        align-items: center;
        justify-content: center;
        .title {
            position: absolute;
            top: 20px;
            left: 24px;
            font-size: 14px;
            color: #5c6479;
        }
    }
    .create-page__right {
        width: 40%;
        padding: 20px 24px;
        background: #ffffff;
        box-shadow: -3px 0px 10px 1px rgba(0, 0, 0, 0.1);
    }
}
</style>
