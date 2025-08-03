<template>
    <div class="card-content">
        <div
            class="card-item"
            v-for="(item, index) in cardDataArr"
            :key="index"
        >
            <div class="card-line">{{ item.name }}</div>
            <div class="card-line">
                <div class="full-name-flex">
                    <el-avatar
                        icon="el-icon-user-solid"
                        size="small"
                        :src="getAvatar()"
                    ></el-avatar>
                    <!-- <tip-one :text="scope.row.full_name"> -->
                    <span v-overflow class="full-name">
                        {{ item.full_name || "--" }}</span
                    >
                </div>
            </div>
            <div class="card-line">{{ item.date }}</div>
            <div class="card-line">
                <span class="label bg1">标签1</span>
                <span class="label bg2">标签2</span>
                <span class="label">+2</span>
            </div>
        </div>
    </div>
</template>
<script>
import { imgHost } from "@/assets/tool/const";
export default {
    data() {
        return {
            cardDataArr: [
                {
                    name: "文档详情编辑后，保存不了",
                    full_name: "潘发发",
                    date: "2023-10-12",
                    label: ["标签1", "标签2", "标签3", "标签4"]
                },
                {
                    name: "文档详情编辑后，保存不了",
                    full_name: "潘发发",
                    date: "2023-10-12",
                    label: ["标签1", "标签2", "标签3", "标签4"]
                },
                {
                    name: "文档详情编辑后，保存不了",
                    full_name: "潘发发",
                    date: "2023-10-12",
                    label: ["标签1", "标签2", "标签3", "标签4"]
                }
            ]
        };
    },
    methods: {
        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        }
    }
};
</script>
<style lang="scss" scoped>
.card-content {
    display: flex;
    .card-item {
        margin-right: 20px;

        width: calc((100% - 40px) / 3);
        padding: 16px;
        background: #ffffff;
        border-radius: 4px 4px 4px 4px;
        border: 1px solid #e6e9f0;
        &:last-child {
            margin-right: 0;
        }
    }
    .card-line {
        margin-bottom: 16px;
        &:last-child {
            margin-bottom: 0;
        }
    }
    .full-name-flex {
        display: flex;
        align-items: center;
        ::v-deep .el-avatar--small {
            width: 24px;
            height: 24px;
            line-height: 24px;
        }
        .full-name {
            margin-left: 8px;
            width: calc(100% - 48px);
            overflow: hidden;
            white-space: nowrap;
            text-overflow: ellipsis;
        }
    }
    .label {
        display: inline-block;
        padding: 0 8px;
        height: 24px;
        line-height: 24px;
        margin-right: 4px;
        border-radius: 4px 4px 4px 4px;
    }
    .bg1 {
        color: #b14300;
        background-color: #fff4e3;
    }
    .bg2 {
        color: #0c6c50;
        background-color: #e6f6f1;
    }
}
</style>
