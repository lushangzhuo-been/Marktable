<template>
    <el-tooltip popper-class="basic-ui" effect="dark" placement="top">
        <div slot="content" class="card-tooltip-content">
            {{ fieldItem.name + "：" + linkInfo.name || linkInfo.url }}
        </div>
        <div class="detail" v-overflow>
            {{ linkInfo.name || linkInfo.url }}
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
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            linkInfo: ""
        };
    },
    watch: {
        detailInfo: {
            handler(info) {
                this.linkInfo = _.cloneDeep(info);
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
    height: 24px;
    line-height: 24px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    background: #f8f8f8;
    border-radius: 4px;
    padding: 0 8px;
    color: #1890ff;
}
</style>
