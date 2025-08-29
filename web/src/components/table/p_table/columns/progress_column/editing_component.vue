<template>
    <div
        class="column-block flex-content"
        :class="{
            isEditing: isEditing,
            'editing-effect': item.can_modify === 'yes',
            'col-required': item.required === 'yes'
        }"
    >
        <!-- 右侧显示数字 -->
        <el-slider
            :disabled="!(item.can_modify === 'yes')"
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
import api from "@/common/api/module/progress";
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
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        },
        curProgress() {
            return this.$route.params.id;
        }
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
            this.fetAuthEdit();
        },
        handleCheck(a, b, c) {},
        fetAuthEdit() {
            // 获取进展权限
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: this.scope.row._id,
                auth_mode: "edit",
                field_key: this.item.field_key
            };
            api.getUserAuth(params).then((res) => {
                if (res && res.resultCode === 200) {
                    if (res.data) {
                        this.$emit(
                            "edit-form-item",
                            this.slider,
                            this.item.field_key,
                            this.scope.row._id,
                            this.item.mode
                        );
                    } else {
                        this.$message({
                            type: "warning",
                            message: "暂无编辑权限"
                        });
                        this.slider = this.scope.row[this.item.field_key] || 0;
                    }
                }
            });
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
