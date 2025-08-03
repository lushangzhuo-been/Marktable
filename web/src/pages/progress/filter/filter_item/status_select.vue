<template>
    <div
        class="column-block select-single"
        ref="column-block"
        :class="{
            isEditing: isEditing,
            empty: !frontArr.length
        }"
    >
        <el-popover
            ref="status-popover"
            popper-class="progress-status-filter-poppover progress-popover padding12"
            placement="bottom"
            :width="popoverWidth"
            trigger="click"
            @show="popoverShow"
            @after-leave="afterLeave"
        >
            <div class="popover-content">
                <div class="list">
                    <div
                        class="list-item"
                        v-for="(listItem, listIndex) in statusOptions"
                        :key="listIndex"
                        @click="changStatus(listItem)"
                    >
                        {{ listItem.name }}
                        <div
                            class="status-background"
                            :style="{
                                backgroundColor: listItem.color,
                                borderColor: listItem.color
                            }"
                        ></div>
                    </div>
                </div>
            </div>
            <div slot="reference">
                <div class="status-block" v-if="submitValue">
                    <div>
                        {{ submitValue }}
                    </div>
                </div>
                <div class="status-block default-text" v-else>请选择</div>
                <b class="triangle"></b>
            </div>
        </el-popover>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress";
export default {
    props: {
        param: {
            type: [Object, String],
            default: () => {}
        }
    },
    data() {
        return {
            isEditing: false,
            popoverWidth: 220,
            frontArr: [],
            submitValue: "",
            statusOptions: []
        };
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        }
    },
    watch: {
        param: {
            // 只有value 没有颜色
            handler(obj) {
                if (obj && Object.keys(obj).length) {
                    this.submitValue = obj.name;
                } else {
                    this.submitValue = "";
                }
            },
            immediate: true
        }
    },
    mounted() {},
    methods: {
        checkScope() {
            this.isEditing = !this.isEditing;
            this.popoverWidth = this.$refs["column-block"].clientWidth;
            this.$nextTick(() => {
                setTimeout(() => {
                    this.$refs["status-popover"].doShow();
                }, 20);
            });
        },
        changStatus(listItem) {
            let singleObj;
            if (_.find(this.statusOptions, { id: listItem.id })) {
                singleObj = _.find(this.statusOptions, {
                    id: listItem.id
                });
            }
            this.submitValue = singleObj.name;
            this.$emit("input", singleObj);
            this.$refs["status-popover"].doClose();
        },
        popoverShow() {
            this.isEditing = true;
            this.popoverWidth = this.$refs["column-block"].clientWidth; // 获取全部状态下拉
            this.fetchAllStatusList();
        },
        fetchAllStatusList() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress
            };
            api.getAllStatusList(params).then((res) => {
                if (res && res.resultCode === 200 && res.data) {
                    this.statusOptions = res.data;
                } else {
                    this.statusOptions = [];
                }
            });
        },
        afterLeave() {
            this.isEditing = false;
        },
        getBorderColor(color) {
            return {
                "border-color": `rgba(${color},.2)`
            };
        },
        getArrFront(arr) {
            let deepClone = _.cloneDeep(arr);
            let front = deepClone.splice(0, 1);
            return front;
        }
    }
};
</script>

<style lang="scss" scoped>
@import "./style.scss";

.status-block {
    box-sizing: border-box;
    position: relative;
    padding: 0 10px;
    text-align: center;
    height: 32px;
    border: 1px solid #dcdfe6;
    background-color: #fff;
    border-radius: 4px;
    .status-background {
        position: absolute;
        height: 100%;
        width: 100%;
        top: 0;
        left: 0;
        opacity: 0.2;
        border-radius: 4px;
    }
    &.default-text {
        text-align: left;
        color: #c0c4cc;
    }
}
.list-item {
    position: relative;
    margin-bottom: 4px;
    border-radius: 4px;
    text-align: center;
    &:last-child {
        margin-bottom: 0px;
    }
    &:hover {
        opacity: 0.9;
    }
    .status-background {
        position: absolute;
        height: 100%;
        width: 100%;
        top: 0;
        left: 0;
        opacity: 0.2;
        border-radius: 4px;
    }
}
</style>
<style lang="scss">
.progress-status-filter-poppover {
    box-sizing: border-box;
    .popover-content {
        max-height: 300px;
        overflow-y: auto;
        padding-right: 12px;
        margin-right: -12px;
    }
}
</style>
