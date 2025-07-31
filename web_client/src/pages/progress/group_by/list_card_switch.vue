<template>
    <div>
        <div class="switch-block">
            <!-- 列表 -->
            <div
                class="type-block list-block"
                :class="{ active: active === 'list' }"
                @click="changeType('list')"
            >
                <b class="type-box list-box"></b>
            </div>
            <!-- 卡片 -->
            <div
                class="type-block card-block"
                :class="{ active: active === 'card' }"
                @click="changeType('card')"
            >
                <b class="type-box card-box"></b>
            </div>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
export default {
    components: {},
    props: {
        value: {
            type: String,
            default: ""
        }
    },
    data() {
        return {
            active: "list"
        };
    },
    computed: {},
    watch: {
        value: {
            handler(type) {
                this.active = type;
            },
            immediate: true,
            deep: true
        }
    },
    mounted() {},
    methods: {
        changeType(type) {
            this.active = type;
            this.$emit("input", type);
        }
    }
};
</script>

<style lang="scss" scoped>
.switch-block {
    box-sizing: border-box;
    border: 1px solid #cdd5e6;
    width: 72px;
    height: 28px;
    display: flex;
    border-radius: 4px;
    .type-block {
        width: 50%;
        display: flex;
        justify-content: center;
        cursor: pointer;
        &.active {
            background-color: #cae4fe;
            border-radius: 4px;
        }
        .type-box {
            display: inline-block;
            width: 18px;
            height: 18px;
            margin: 5px 0;
            background-size: 100% 100%;
            &.list-box {
                background-image: url(@/assets/image/progress/list_view.svg);
            }
            &.card-box {
                background-image: url(@/assets/image/progress/card_view.svg);
            }
        }
    }
}
</style>
