<template>
    <el-tooltip popper-class="basic-ui" effect="dark" placement="top">
        <div slot="content" class="card-tooltip-content">
            {{ fieldItem.name + "：" + relationInfo.title }}
        </div>
        <div class="detail">
            {{ relationInfo.title }}
        </div>
    </el-tooltip>
</template>

<script>
import _ from "lodash";
export default {
    props: {
        fieldItem: {
            type: Object,
            default: () => {},
        },
        detailInfo: {
            type: Array,
            default: () => [],
        },
    },
    data() {
        return {
            relationInfo: "",
        };
    },
    watch: {
        detailInfo: {
            handler(info) {
                if (info && info.length) {
                    this.relationInfo = _.cloneDeep(info[0]);
                } else {
                    this.relationInfo = [];
                }
            },
            immediate: true,
            deep: true,
        },
    },
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
}
</style>
