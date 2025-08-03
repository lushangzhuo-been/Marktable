<template>
    <div>
        <el-popover
            ref="Sort"
            placement="bottom-start"
            popper-class="operation-popover"
            trigger="click"
            @show="popoverShow"
            @hide="popoverHide"
        >
            <!-- 遍历item.list  做下拉内容 -->
            <div class="popover-content">
                <el-checkbox
                    :indeterminate="isIndeterminate"
                    v-model="checkAll"
                    @change="handleCheckAllChange"
                    class="popover-checkbox-item-block check-all"
                    >全选</el-checkbox
                >
                <div class="check-all-checkbox-bottom-line"></div>
                <el-checkbox-group
                    class="add-rule-operation-checkbox"
                    v-model="checkedtrigger"
                    @change="handleCheckedTriggerChange"
                >
                    <el-checkbox
                        v-for="(item, index) in triggersList"
                        :label="item.value"
                        :key="index"
                        class="popover-checkbox-item-block"
                    >
                        <b class="type-icon" :class="item.value"></b>
                        {{ item.label }}</el-checkbox
                    >
                </el-checkbox-group>
            </div>
            <span
                slot="reference"
                class="btn-icon"
                :class="{ editing: editing, effect: effectiveParamsNum }"
            >
                是否启用
                <span v-if="effectiveParamsNum">
                    {{ `(${showCheckedtrigger})` }}
                    <i
                        @click="removeAll"
                        class="el-tag__close el-icon-close tag-delete"
                    ></i>
                </span>
            </span>
        </el-popover>
    </div>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress_rule";
export default {
    components: {},
    props: {},
    data() {
        return {
            editing: false,
            isIndeterminate: false,
            checkAll: false,
            checkedtrigger: [],
            triggersList: []
        };
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        },
        effectiveParamsNum() {
            return this.checkedtrigger.length;
        },
        showCheckedtrigger() {
            if (!this.checkedtrigger.length) {
                return "";
            } else if (this.checkedtrigger.length === 1) {
                if (this.checkedtrigger[0] === "yes") {
                    return "是";
                } else {
                    return "否";
                }
            } else {
                return "全部";
            }
        }
    },
    watch: {
        checkedtrigger: {
            handler(arr) {
                let param = "";
                if (arr) {
                    param = arr.join(",");
                }
                this.$emit("input", param);
            }
        }
    },
    mounted() {},
    methods: {
        removeAll() {
            this.isIndeterminate = false;
            this.checkAll = false;
            this.checkedtrigger = [];
        },
        handleCheckAllChange(val) {
            if (val) {
                let arr = [];
                this.triggersList.forEach((item) => {
                    arr.push(item.value);
                });
                this.checkedtrigger = arr;
            } else {
                this.checkedtrigger = [];
            }
            this.isIndeterminate = false;
        },
        handleCheckedTriggerChange(value) {
            let checkedCount = value.length;
            this.checkAll = checkedCount === this.triggersList.length;
            this.isIndeterminate =
                this.triggersList.length > 0 &&
                checkedCount > 0 &&
                checkedCount < this.triggersList.length;
        },
        getConfig() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress
            };
            api.getRuleConfig(params).then((res) => {
                if (res.resultCode === 200) {
                    let data = res.data;
                    this.triggersList = data.rule_enable || [];
                }
            });
        },
        closePopver() {
            if (this.$refs.Sort) {
                // 重置参数
                this.$refs.Sort.doClose();
            }
        },
        popoverShow() {
            this.editing = true;
            this.getConfig();
        },
        popoverHide() {
            this.editing = false;
        }
    }
};
</script>

<style lang="scss" scoped>
.btn-icon {
    cursor: pointer;
    color: #171e17;
    padding: 4px 8px;
    border-radius: 4px;
    &:hover {
        // color: var(--PRIMARY_COLOR);
        color: #171e17;
        background-color: rgb(202, 228, 253);
        padding: 4px 8px;
        border-radius: 4px;
    }
    &.editing,
    &.effect {
        color: #171e17;
        background-color: rgb(202, 228, 253);
        padding: 4px 8px;
        border-radius: 4px;
    }
}
.type-icon {
    display: inline-block;
    width: 18px;
    height: 18px;
    background-size: 100% 100%;
    vertical-align: text-top;
    &.yes {
        background-image: url(@/assets/image/progress/add_rule/update.svg);
    }
    &.no {
        background-image: url(@/assets/image/progress/add_rule/email.svg);
    }
}
.el-tag__close.el-icon-close.tag-delete {
    width: 12px;
    height: 12px;
    text-align: center;
    line-height: 12px;
    color: #ffffff;
    background-color: #a6abbc;
    border-radius: 50%;
    margin-left: 2px;
    &:hover {
        background-color: #82889e;
    }
    &::before {
        content: "\e6db";
        font-size: 8px;
        position: relative;
        top: -2px;
    }
}
</style>
<style lang="scss">
.el-popper.operation-popover {
    box-sizing: border-box;
    margin-top: 0px;
    padding: 0;
    .popper__arrow {
        display: none;
    }
}
</style>
