<template>
    <el-table-column
        show-overflow-tooltip
        :min-width="column.minWidth"
        :width="column.width"
        :label="column.name"
    >
        <!-- 单人信息列 -->
        <template slot-scope="scope">
            <div v-if="scope.row[column.prop] && scope.row[column.prop].length">
                <user-message :userMessage="scope.row[column.prop][0]">
                    <span>
                        <el-avatar
                            class="progress-avatar"
                            icon="el-icon-user-solid"
                            size="small"
                            :src="getAvatar(scope.row[column.prop][0].avatar)"
                            :style="
                                getAvatarStack(scope.row[column.prop][0].avatar)
                            "
                        ></el-avatar>
                        {{ scope.row[column.prop][0].full_name }}
                    </span>
                </user-message>
            </div>
            <div v-else>--</div>
        </template>
    </el-table-column>
</template>

<script>
import { imgHost } from "@/assets/tool/const";
import UserMessage from "@/components/user_message_tip";
export default {
    props: {
        column: {
            type: Object,
            default: () => {}
        }
    },
    components: {
        UserMessage
    },
    data() {
        return {};
    },
    computed: {},
    watch: {},
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
                top: "-2px"
            };
        }
    }
};
</script>

<style lang="scss" scoped>
.status {
    display: inline-block;
    width: 8px;
    height: 8px;
    border-radius: 50%;
    margin-right: 2px;
    &.yes {
        background-color: #1890ff;
    }
    &.no {
        background-color: #ffa414;
    }
}
</style>
