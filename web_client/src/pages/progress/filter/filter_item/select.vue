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
            popper-class="progress-popover padding12"
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
                        v-for="(listItem, listIndex) in option"
                        :key="listIndex"
                        @click="changStatus(listItem)"
                    >
                        {{ listItem.value }}
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
                <div class="status-block" v-if="frontArr.length">
                    <div
                        v-for="(tagItem, tagIndex) in frontArr"
                        :key="tagIndex"
                    >
                        {{ tagItem.value }}
                        <div
                            class="status-background"
                            :style="{
                                backgroundColor: tagItem.color,
                                borderColor: tagItem.color
                            }"
                        ></div>
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
export default {
    props: {
        option: {
            type: Array,
            default: () => []
        },
        param: {
            type: [Array, String],
            default: () => []
        }
    },
    data() {
        return {
            isEditing: false,
            popoverWidth: 220,
            frontArr: [],
            submitArr: []
        };
    },
    watch: {
        param: {
            handler(arr) {
                let selectArr = [];
                if (arr) {
                    arr.forEach((objItem) => {
                        selectArr.push(objItem.value);
                    });
                    this.submitArr = arr;
                }
                this.selectArr = selectArr;
                this.frontArr = this.getArrFront(this.submitArr);
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
            this.submitArr = [listItem];
            // this.formData[this.formItem.field_key];
            this.frontArr = this.getArrFront(this.submitArr);
            this.$emit("input", this.submitArr);
            this.$refs["status-popover"].doClose();
        },
        popoverShow() {
            this.isEditing = true;
            this.popoverWidth = this.$refs["column-block"].clientWidth;
        },
        afterLeave() {
            this.isEditing = false;
            // this.formData[this.formItem.field_key] = this.submitArr;
            // this.doValidate();
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
        // doValidate() {
        //     if (this.formItem.required === "yes" && !this.submitArr.length) {
        //         this.validateFailed = true;
        //         this.formItem.validateFailed = true;
        //     } else {
        //         this.validateFailed = false;
        //         this.formItem.validateFailed = false;
        //     }
        // },
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
