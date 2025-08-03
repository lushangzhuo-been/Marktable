<template>
    <el-tooltip popper-class="basic-ui" effect="dark" placement="top">
        <div slot="content" class="card-tooltip-content">
            {{ fieldItem.name + "：" + text }}
        </div>
        <div class="detail">
            {{ text }}
        </div>
    </el-tooltip>
</template>

<script>
import _ from "lodash";
export default {
    props: {
        fieldItem: {
            type: Object,
            default: () => {}
        },
        detailInfo: {
            type: String,
            default: ""
        }
    },
    data() {
        return {
            text: ""
        };
    },
    watch: {
        detailInfo: {
            handler(info) {
                this.text = _.cloneDeep(info);
            },
            immediate: true,
            deep: true
        }
    }
};
</script>

<style lang="scss" scoped>
.detail {
    box-sizing: border-box;
    display: inline-block;
    height: 24px;
    line-height: 24px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
}
</style>
