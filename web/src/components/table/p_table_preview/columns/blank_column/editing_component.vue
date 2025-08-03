<template>
    <div class="column-block" :class="{ isEditing: isEditing }">
        <!-- 只有详情没有编辑 -->
        <div class="detail">
            <!-- 创建时间 、更新时间 -->
            <div
                class="padding"
                v-if="
                    item.field_key === 'created_at' ||
                    item.field_key === 'updated_at'
                "
            >
                {{ time || emptySpace() }}
            </div>
            <div
                class="status"
                :style="{
                    backgroundColor: statusColor,
                    borderColor: statusColor
                }"
                v-if="item.field_key === 'status'"
            >
                {{ status || emptySpace() }}
            </div>
            <div class="padding" v-if="item.field_key === 'creator'">
                <div
                    class="tag-item"
                    v-show="frontArr.length"
                    v-for="(tagItem, tagIndex) in frontArr"
                    :key="tagIndex"
                >
                    <user-message :userMessage="tagItem">
                        <span>
                            <el-avatar
                                class="progress-avatar"
                                icon="el-icon-user-solid"
                                size="small"
                                :src="getAvatar(tagItem.avatar)"
                            ></el-avatar>
                            {{ tagItem.full_name }}
                        </span>
                    </user-message>
                </div>
                <div v-show="!frontArr.length">
                    {{ emptySpace() }}
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import { imgHost } from "@/assets/tool/const";
import { emptySpace } from "@/assets/tool/func";
import UserMessage from "@/components/user_message_tip";
export default {
    components: {
        UserMessage
    },
    props: {
        item: {
            type: Object,
            default: () => {}
        },
        scope: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            isEditing: false,
            time: "",
            status: "",
            statusColor: "",
            creator: {},
            frontArr: []
        };
    },
    watch: {
        scope: {
            handler(scope) {
                if (
                    this.item.field_key === "created_at" ||
                    this.item.field_key === "updated_at"
                ) {
                    this.time = scope.row[this.item.field_key] || "";
                } else if (this.item.field_key === "status") {
                    this.status = scope.row[this.item.field_key] || {};
                    this.statusColor =
                        scope.row[this.item.field_key + "_color"] || "";
                } else if (this.item.field_key === "creator")
                    if (scope.row[this.item.field_key]) {
                        // this.creator = scope.row[this.item.field_key] || {};
                        this.frontArr = this.getArrFront(
                            scope.row[this.item.field_key]
                        );
                    }
            },
            immediate: true
        }
    },
    mounted() {},
    methods: {
        emptySpace(param) {
            return emptySpace(param);
        },
        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        },
        getArrFront(arr) {
            let deepClone = [_.cloneDeep(arr)];
            let front = deepClone.splice(0, 1);
            return front;
        }
    }
};
</script>

<style lang="scss" scoped>
.status {
    text-align: center;
    color: #fff;
}
.column-block {
    box-sizing: border-box;
    height: 40px;
    line-height: 40px;
    border: 1px solid rgba(0, 0, 0, 0);
    // padding: 0 10px;
    &.isEditing {
        border: 1px solid var(--PRIMARY_COLOR);
    }
}

.detail {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    height: 40px;
    .padding {
        padding: 0 10px;
    }
}
.tag-item {
    display: inline-block;
    width: 100%;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
}

::v-deep .progress-input.el-input {
    .el-input__inner {
        padding: 0;
        border: none;
        height: 38px;
        line-height: 38px;
        background-color: rgba(0, 0, 0, 0);
    }
}
</style>
