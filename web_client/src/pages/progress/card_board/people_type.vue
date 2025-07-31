<template>
    <el-tooltip popper-class="basic-ui" effect="dark" placement="top">
        <div slot="content" class="card-tooltip-content">
            {{ fieldItem.name + "：" }}
            <span v-for="(item, index) in personList" :key="index"
                >{{ item.full_name }}
                <span v-show="index !== personList.length - 1">、</span>
            </span>
        </div>
        <div class="detail">
            <div
                ref="listCon"
                class="tag-list"
                :class="{ 'show-num': showNum }"
            >
                <div
                    class="tag-item"
                    v-for="(tagItem, tagIndex) in frontArr"
                    :key="tagIndex"
                >
                    <el-avatar
                        class="progress-avatar"
                        icon="el-icon-user-solid"
                        size="small"
                        :src="getAvatar(tagItem.avatar)"
                        :style="getAvatarStack(tagItem.avatar, tagIndex)"
                    ></el-avatar>
                    {{ tagItem.full_name }}
                </div>
            </div>
            <b v-show="showNum" class="num-box" ref="num-box"
                >+{{ behandArr.length }}</b
            >
        </div>
    </el-tooltip>
</template>

<script>
import _ from "lodash";
import { imgHost } from "@/assets/tool/const";
export default {
    components: {},
    props: {
        fieldItem: {
            type: Object,
            default: () => {}
        },
        detailInfo: {
            type: Array,
            default: () => []
        }
    },
    data() {
        return {
            personList: [],
            frontArr: [],
            behandArr: [],
            labelIndex: 1,
            showNum: false
        };
    },
    watch: {
        detailInfo: {
            handler(info) {
                this.personList = _.cloneDeep(info);
                if (this.personList && this.personList.length > 1) {
                    this.showNum = true;
                } else {
                    this.showNum = false;
                }
                this.frontArr = this.getArrFront(this.personList);
                this.behandArr = this.getArrBehand(this.personList);
            },
            immediate: true,
            deep: true
        }
    },
    methods: {
        getArrFront(arr) {
            let deepClone = _.cloneDeep(arr);
            let front = deepClone.splice(0, this.labelIndex);
            return front;
        },
        getArrBehand(arr) {
            let deepClone = _.cloneDeep(arr);
            let behand = deepClone.splice(this.labelIndex);
            return behand;
        },
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
.detail {
    display: flex;
    height: 24px;
    .tag-list {
        display: inline-block;
        height: 24px;
        max-width: calc(100% - 30px);
        &.show-num {
            white-space: nowrap;
        }
        .tag-item {
            display: inline-block;
            padding-right: 8px;
            overflow: hidden;
            text-wrap: nowrap;
            text-overflow: ellipsis;
            max-width: 148px;
            height: 24px;
            line-height: 24px;
        }
    }
}
.progress-avatar.el-avatar {
    width: 20px;
    height: 20px;
    line-height: 1;
    vertical-align: middle;
    margin-right: 4px;
}
.num-box {
    box-sizing: border-box;
    display: inline-block;
    min-width: 22px;
    height: 22px;
    line-height: 22px;
    border-radius: 11px;
    padding: 0 2px;
    background-color: #f8f8f8;
    text-align: center;
    vertical-align: middle;
    font-weight: 400;
    font-size: 12px;
    color: #2f384c;
}
</style>
