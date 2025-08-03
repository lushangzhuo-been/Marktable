<template>
    <div
        class="column-block flex-content"
        :class="{
            isEditing: isEditing,
            'editing-effect': formItem.can_modify === 'yes',
            'col-required': formItem.required === 'yes'
        }"
    >
        <!-- 右侧显示数字 -->
        <el-slider
            :disabled="!(formItem.can_modify === 'yes')"
            class="silder-block"
            :show-tooltip="false"
            v-model="slider"
            @change="sliderChange"
        >
        </el-slider>
        <div class="num-block">{{ slider + "%" }}</div>
    </div>
</template>

<script>
import _ from "lodash";
import { kSeq } from "@/assets/tool/func";
export default {
    props: {
        formItem: {
            type: Object,
            default: () => {}
        },
        formData: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            isEditing: false,
            slider: 0
        };
    },
    watch: {
        formData: {
            handler(formData) {
                this.slider =
                    _.cloneDeep(formData[this.formItem.field_key]) || 0;
            }
        }
    },
    mounted() {},
    methods: {
        sliderChange() {
            this.$emit(
                "edit-form-item",
                this.slider,
                this.formItem.field_key,
                this.formItem.mode
            );
        }
    }
};
</script>

<style lang="scss" scoped>
.column-block {
    box-sizing: border-box;
    // height: 40px;
    // line-height: 40px;
    height: 32px;
    line-height: 32px;
    border: 1px solid rgba(0, 0, 0, 0);
    padding: 0 10px;
    border-radius: 4px;
    &:hover {
        background: #f5f6f8;
        border-radius: 4px 4px 4px 4px;
    }
    &.isEditing {
        border: 1px solid var(--PRIMARY_COLOR);
        background: #ffffff;
        box-shadow: 0px 4px 16px 1px rgba(0, 0, 0, 0.1);
        border-radius: 4px;
        border: 1px solid #e6e9f0;
    }
}
::v-deep.flex-content {
    // width: 100%;
    position: relative;
    display: flex;
    width: 60%;
    min-width: 220px;
    .el-slider.silder-block {
        width: calc(100% - 50px);
        .el-slider__runway {
            // 背景条
            box-sizing: border-box;
            height: 10px;
            margin: 10px 0;
            background-color: rgba(0, 0, 0, 0);
            border: 1px solid #10a862;
            border-radius: 5px;
            .el-slider__bar {
                // 实际条
                background-color: #10a862;
                height: 9px;
                border-top-left-radius: 5px;
                border-bottom-left-radius: 5px;
            }
        }
        .el-slider__button-wrapper {
            top: -14px;
            z-index: 3;
            .el-tooltip.el-slider__button {
                box-sizing: border-box;
                width: 12px;
                height: 12px;
                border: 2px solid #10a862;
                background-color: #10a862;
            }
        }
    }
    .num-block {
        width: 50px;
        text-align: center;
        color: #10a862;
        font-size: 12px;
    }
}
</style>
