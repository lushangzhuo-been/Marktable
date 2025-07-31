<template>
    <div
        class="column-block"
        :class="{
            isEditing: isEditing,
            'editing-effect': item.can_modify === 'yes',
            'col-required': item.required === 'yes'
        }"
    >
        <div v-if="!isEditing" class="row">
            <div
                class="title-block detail"
                v-overflow
                @click="checkScope(scope)"
            >
                <span class="detail-text" @click.stop="openDetail(scope)">
                    {{ text }}
                </span>
            </div>
            <div class="btn-block" @click="openDetail(scope)">
                <b
                    class="evolve-box"
                    :class="
                        scope.row.progress_num ? 'have-evolve' : 'no-evolve'
                    "
                ></b>
                <div v-show="scope.row.progress_num" class="evolve-num">
                    {{
                        scope.row.progress_num < 100
                            ? scope.row.progress_num
                            : "99+"
                    }}
                </div>
            </div>
        </div>
        <div v-if="isEditing">
            <el-input
                class="basic-ui progress-input"
                ref="main-text-input"
                v-model.trim="text"
                placeholder="请输入内容"
                @blur="blur"
            ></el-input>
        </div>
    </div>
</template>

<script>
export default {
    components: {},
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
            text: ""
        };
    },
    watch: {
        scope: {
            handler(scope) {
                this.text = scope.row[this.item.field_key] || "";
            },
            immediate: true
        }
    },
    mounted() {},
    methods: {
        checkScope() {
            if (this.item.can_modify === "no") {
                return;
            }
            this.isEditing = true;
            this.$nextTick(() => {
                this.$refs["main-text-input"].focus();
            });
        },
        blur() {
            this.isEditing = false;
            // 必填校验 与 可编辑校验
            if (
                this.item.required === "yes" &&
                this.scope.row[this.item.field_key] &&
                !this.text
            ) {
                this.$message({
                    showClose: true,
                    message: "此为必填项",
                    type: "warning"
                });
                this.text = this.scope.row[this.item.field_key];
            } else if (
                this.scope.row[this.item.field_key] == this.text ||
                (!this.scope.row[this.item.field_key] && !this.text)
            ) {
                //
            } else {
                this.$emit(
                    "edit-form-item",
                    this.text,
                    this.item.field_key,
                    this.scope.row._id,
                    this.item.mode
                );
            }
        },
        openDetail(scope) {
            this.$emit("open-detail", scope, "tmpl_detail");
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

    &.isEditing {
        border: 1px solid var(--PRIMARY_COLOR);
    }
}

.row {
    display: flex;
    height: 40px;
    line-height: 40px;

    .title-block {
        flex: 1;
    }

    .btn-block {
        width: 40px;
        border-left: 1px solid #ebeef5;
        cursor: pointer;
        position: relative;
        .evolve-num {
            position: absolute;
            top: 56%;
            right: 14%;
            padding: 0 2px;
            min-width: 12px;
            height: 12px;
            line-height: 12px;
            background: #8dc8ff;
            border-radius: 5px;
            font-weight: 500;
            font-size: 8px;
            color: #ffffff;
            text-align: center;
        }
    }
}

.detail {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}

.detail-text {
    color: var(--PRIMARY_COLOR);
    cursor: pointer;
    &:hover {
        text-decoration: underline;
    }
}

::v-deep .progress-input.el-input {
    .el-input__inner {
        padding: 0;
        border: none;
        height: 38px;
        line-height: 38px;
        background-color: rgba(0, 0, 0, 0);
        padding-right: 10px;
    }
}
.evolve-box {
    display: inline-block;
    width: 20px;
    height: 20px;
    margin: 10px;
    background-size: 100% 100%;
    &.have-evolve {
        background-image: url("~@/assets/image/progress_detail/have-evolve.svg");
    }
    &.no-evolve {
        background-image: url("~@/assets/image/progress_detail/no-evolve.svg");
    }
}
</style>
