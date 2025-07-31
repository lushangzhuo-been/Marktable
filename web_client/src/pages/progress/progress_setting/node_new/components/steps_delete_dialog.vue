<template>
    <el-dialog
        title="删除"
        :visible.sync="dialogVisible"
        class="basic-ui"
        append-to-body
        @close="cancel"
    >
        <div slot="title" class="dialog-title">
            <i class="el-icon-warning"></i>
            <span>删除</span>
        </div>
        <div class="name">
            <!-- 确定要删除步骤流转"{{ curRow.name || "--" }}" -> "{{
                curCol.label || "--"
            }}"吗？ -->
            确定要删除步骤流转:&nbsp;&nbsp;

            <tip-one :text="curRow.name">
                <span
                    class="label-name"
                    :style="{
                        // color: `${curRow.color || '#000'}`,
                        color: '#171e31',
                        backgroundColor: rgbToRgba(
                            `${curRow.color || '#fff'}`,
                            0.3
                        )
                    }"
                    >{{ curRow.name || "--" }}
                </span>
            </tip-one>
            <img
                :src="require('@/assets/image/common/right_arrow.png')"
                alt=""
                width="16px"
                height="16px"
                class="arrow"
            />
            <tip-one :text="curCol.label">
                <span
                    class="label-name"
                    :style="{
                        // color: `${curCol.color || '#000'}`,
                        color: '#171e31',
                        backgroundColor: rgbToRgba(
                            `${curCol.color || '#fff'}`,
                            0.3
                        )
                    }"
                    >{{ curCol.label || "--" }}
                </span>
            </tip-one>
            &nbsp;&nbsp;吗?
        </div>
        <span slot="footer" class="dialog-footer">
            <el-button size="small" class="basic-ui" @click="cancel"
                >取消</el-button
            >
            <el-button
                size="small"
                class="basic-ui"
                type="primary"
                @click="onConfirm"
            >
                确定
            </el-button>
        </span>
    </el-dialog>
</template>

<script>
import _ from "lodash";
import api from "@/common/api/module/progress_setting";
import TipOne from "@/pages/common/tooltip_one_line.vue";
import { rgbToRgba } from "@/assets/tool/func";
export default {
    components: {
        TipOne
    },
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        }
    },
    data() {
        return {
            curRow: {},
            curCol: {},
            dialogVisible: false
        };
    },
    methods: {
        rgbToRgba(color, opacity) {
            return rgbToRgba(color, opacity);
        },
        openDialog(val, row, col) {
            this.curRow = row;
            this.curCol = _.cloneDeep(col);
            this.dialogVisible = true;
        },
        cancel() {
            this.dialogVisible = false;
            this.$set(
                this.curRow[`row_${this.curRow.id}_${this.curCol.id}`],
                "checked",
                true
            );
        },
        onConfirm() {
            //确认取消的时候
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id
            };
            if (this.curRow.steps) {
                this.curRow.steps.forEach((item) => {
                    if (
                        item.start_status_id === this.curRow.id &&
                        item.end_status_id === this.curCol.id
                    ) {
                        params.id = item.id;
                    }
                });
            }
            api.deleteSteps(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.$emit("on-confirm");
                    this.$message({
                        showClose: true,
                        message: "删除成功",
                        type: "success"
                    });
                } else {
                    this.$set(
                        this.curRow[`row_${this.curRow.id}_${this.curCol.id}`],
                        "checked",
                        true
                    );
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.width100 {
    width: 100%;
}
.dialog-title {
    display: flex;
    align-items: center;
}
.name {
    display: flex;
    align-items: center;
    img.arrow {
        margin: 0 12px;
    }
    .label-name {
        box-sizing: border-box;
        display: inline-block;
        height: 24px;
        line-height: 24px;
        border-radius: 4px;
        padding: 0 10px;
        font-size: 14px;
        font-weight: 400;
        color: #171e31;
        max-width: 118px;
        text-overflow: ellipsis;
        overflow: hidden;
        white-space: nowrap;
    }
}
</style>
