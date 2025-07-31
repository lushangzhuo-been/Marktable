<template>
    <div>
        <div ref="listCon" class="tag-list">
            <div
                class="tag-item"
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
                            :style="getAvatarStack(tagItem.avatar, tagIndex)"
                        ></el-avatar>
                        {{ tagItem.full_name }}
                    </span>
                </user-message>
            </div>
        </div>
    </div>
</template>

<script>
import { imgHost } from "@/assets/tool/const";
import UserMessage from "@/components/user_message_tip";
export default {
    props: {
        userInfo: {
            type: Array,
            default: () => []
        }
    },
    components: {
        UserMessage
    },
    data() {
        return {
            frontArr: []
        };
    },
    watch: {
        userInfo: {
            handler(arr) {
                this.frontArr = arr || [];
            },
            immediate: true,
            deep: true
        }
    },
    methods: {
        getAvatar(src) {
            if (src) {
                return `${imgHost}${src}`;
            }
            return require(`@/assets/image/common/default_avatar_big.png`);
        },
        getAvatarStack() {
            return {
                position: "relative",
                top: "-3px"
            };
        }
    }
};
</script>

<style lang="scss" scoped></style>
