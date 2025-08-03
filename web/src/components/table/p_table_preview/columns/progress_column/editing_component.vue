<template>
    <div class="column-block flex-content">
        <!-- 右侧显示数字 -->
        <el-slider
            :disabled="true"
            class="silder-block"
            :show-tooltip="false"
            v-model="slider"
        >
        </el-slider>
        <div class="num-block">{{ slider + "%" }}</div>
    </div>
</template>

<script>
export default {
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
            slider: 0
        };
    },
    watch: {
        scope: {
            handler(scope) {
                this.slider = scope.row[this.item.field_key] || 0;
            },
            deep: true,
            immediate: true
        }
    },
    mounted() {},
    methods: {
        sliderChange(slider) {
            this.$emit(
                "edit-form-item",
                this.slider,
                this.item.field_key,
                this.scope.row._id,
                this.item.mode
            );
        }
    }
};
</script>

<style lang="scss" scoped>
.column-block {
    box-sizing: border-box;
    height: 40px;
    line-height: 40px;
    border: 1px solid rgba(0, 0, 0, 0);
    padding-left: 10px;
}
::v-deep.flex-content {
    width: 100%;
    display: flex;
    .el-slider.silder-block {
        width: calc(100% - 50px);
        .el-slider__runway {
            // 背景条
            box-sizing: border-box;
            height: 10px;
            margin: 15px 0;
            background-color: rgba(0, 0, 0, 0);
            border: 1px solid #10a862;
            border-radius: 5px;
            .el-slider__bar {
                // 实际条
                box-sizing: border-box;
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
        line-height: 40px;
    }
}
</style>
