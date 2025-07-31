<template>
    <el-tooltip popper-class="basic-ui" effect="dark" placement="top">
        <div slot="content" class="card-tooltip-content">
            {{ fieldItem.name + "：" + slider + "%" }}
        </div>
        <div class="detail flex-content">
            <el-slider
                disabled
                class="silder-block"
                :show-tooltip="false"
                v-model="slider"
            >
            </el-slider>
            <div class="num-block">{{ slider + "%" }}</div>
        </div>
    </el-tooltip>
</template>

<script>
import _ from "lodash";
export default {
    props: {
        fieldItem: {
            type: Object,
            default: () => {}
        },
        detailInfo: {
            type: Number,
            default: 0
        }
    },
    data() {
        return {
            slider: 0
        };
    },
    watch: {
        detailInfo: {
            handler(info) {
                this.slider = _.cloneDeep(info);
            },
            immediate: true,
            deep: true
        }
    }
};
</script>

<style lang="scss" scoped>
.detail {
    box-sizing: border-box;
    height: 24px;
    line-height: 24px;
    background: #f8f8f8;
    border-radius: 4px;
    padding: 0 0 0 8px;
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
            margin: 8px 0;
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
            &.disabled {
                .el-slider__button-wrapper {
                    &:hover {
                        cursor: default;
                    }
                    .el-tooltip.el-slider__button {
                        cursor: default;
                    }
                }
            }
            .el-slider__button-wrapper {
                height: 24px;
                width: 24px;
                top: -8px;
                z-index: 3;
                .el-tooltip.el-slider__button {
                    z-index: 10;
                    box-sizing: border-box;
                    width: 12px;
                    height: 12px;
                    border: 2px solid #10a862;
                    background-color: #10a862;
                }
            }
        }
    }
    .num-block {
        width: 50px;
        text-align: center;
        color: #10a862;
        font-size: 12px;
        line-height: 24px;
    }
}
</style>
