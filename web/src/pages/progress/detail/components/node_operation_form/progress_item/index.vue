<template>
    <div class="column-block flex-content">
        <el-slider
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
            isEditing: true,
            slider: 0
        };
    },
    watch: {
        formData: {
            handler(formData) {
                this.slider = formData[this.formItem.field_key] || 0;
            },
            immediate: true
        }
    },
    mounted() {},
    methods: {
        sliderChange(slider) {
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
    line-height: 32px;
    border: 1px solid rgba(0, 0, 0, 0);
    padding: 0;
    .validate-desc {
        color: #f56c6c;
        font-size: 12px;
        height: 12px;
        line-height: 12px;
    }
}
::v-deep .progress-input.el-input {
    &.number.hidden-arrow {
        input::-webkit-outer-spin-button,
        input::-webkit-inner-spin-button {
            -webkit-appearance: none;
        }
        input[type="number"] {
            -moz-appearance: textfield;
        }
    }
    &.validate-failed {
        .el-input__inner {
            border: 1px solid #fd5050;
        }
    }
}
::v-deep.flex-content {
    width: 100%;
    display: flex;
    padding-left: 12px;
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
