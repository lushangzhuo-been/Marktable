<template>
    <el-dialog
        :title="title"
        :visible.sync="dialogVisible"
        class="add-trigger-dialog basic-ui"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        @close="cancel"
    >
        <el-form
            size="small"
            ref="form"
            :model="form"
            :rules="rule"
            label-width="82px"
        >
            <el-form-item prop="type" label="类型:">
                <el-radio-group v-model="form.mode">
                    <el-radio
                        v-for="(item, index) in radioList"
                        :key="index"
                        :label="item.value"
                    >
                        {{ item.label }}
                    </el-radio>
                    <!-- <el-radio :label="'首节点'">首节点</el-radio>
                    <el-radio :label="'中间节点'">中间节点</el-radio>
                    <el-radio :label="'末节点'">末节点</el-radio> -->
                </el-radio-group>
            </el-form-item>
            <el-form-item prop="name" label="名称:">
                <el-input
                    class="basic-ui width"
                    v-model="form.name"
                    maxlength="25"
                    show-word-limit
                    placeholder="请输入名称"
                ></el-input>
            </el-form-item>

            <el-form-item prop="status_text" label="状态:">
                <el-input
                    class="basic-ui width"
                    v-model="form.status_text"
                    maxlength="25"
                    show-word-limit
                    placeholder="请输入状态"
                ></el-input>
            </el-form-item>
            <el-form-item label="状态颜色:">
                <el-color-picker
                    ref="statusColorPicker"
                    style="top: 4px"
                    :class="['basic-ui', !form.status_color ? 'border' : '']"
                    v-model="form.status_color"
                    popper-class="radioColorPicker"
                    :predefine="COLOR_PICKER_LIST"
                    @active-change="colorActiveChange"
                ></el-color-picker>
            </el-form-item>
        </el-form>
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
import api from "@/common/api/module/progress_setting";
import {
    PRIMARY_COLOR_LIST,
    COLOR_PICKER_LIST
} from "@/assets/tool/color_list";
import _ from "lodash";
export default {
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        }
    },
    data() {
        return {
            curType: "",
            curRow: {},
            title: "",
            radioList: [
                {
                    value: "head",
                    label: "首节点"
                },
                {
                    value: "middle",
                    label: "中间节点"
                },
                {
                    value: "end",
                    label: "末节点"
                }
            ],
            dialogVisible: false,
            COLOR_PICKER_LIST: COLOR_PICKER_LIST,
            form: {
                name: "",
                mode: "middle",
                status_text: "",
                status_color: PRIMARY_COLOR_LIST.PRIMARY_COLOR_RGB
            },
            rule: {
                mode: [
                    {
                        required: true,
                        message: "请选择类型"
                    }
                ],
                name: [
                    {
                        required: true,
                        message: "请输入名称"
                    }
                ],
                status_text: [
                    {
                        required: true,
                        message: "请输入状态"
                    }
                ]
            }
        };
    },
    watch: {
        dialogVisible: {
            handler(flag) {
                if (!flag) {
                    this.form = {
                        name: "",
                        mode: "middle",
                        status_text: "",
                        status_color: PRIMARY_COLOR_LIST.PRIMARY_COLOR_RGB
                    };
                }
            }
        }
    },
    methods: {
        colorActiveChange(val) {
            this.$set(this.form, "status_color", val);
            let ref = this.$refs["statusColorPicker"];
            if (!ref) return;
            ref.hide();
        },
        openDialog(type, list, row = {}) {
            this.curType = type;
            this.curRow = _.cloneDeep(row);
            // 设置节点radio配置
            // this.setModeConfig(list)
            this.dialogVisible = true;
            if (type === "add") {
                this.title = "新增步骤";
                this.form = {
                    name: "",
                    mode: "middle",
                    status_text: "",
                    status_color: PRIMARY_COLOR_LIST.PRIMARY_COLOR_RGB
                };
            } else {
                this.title = "编辑步骤";
                this.fetchNodeDetail(row);
            }
            this.$nextTick(() => {
                this.$refs.form.clearValidate();
            });
        },
        setModeConfig(list) {
            if (list && list.length) {
                this.radioList = _.cloneDeep(list);
            } else {
                this.radioList = [];
            }
        },
        fetchNodeDetail(row) {
            if (!row || !row.node) return;
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.$route.params.id
            };
            params.id = row && row.node ? row.node.id : "";
            api.getNodeDetail(params).then((res) => {
                if (
                    res &&
                    res.resultCode === 200 &&
                    res.data &&
                    Object.keys(res.data).length
                ) {
                    this.setFormValue(res.data);
                } else {
                    this.form = {
                        name: "",
                        mode: "middle",
                        status_text: "",
                        status_color: row.node.status_color || null
                    };
                }
            });
        },
        setFormValue(obj) {
            for (let key in this.form) {
                this.form[key] = obj[key] || "";
            }
        },
        cancel() {
            this.form = {
                name: "",
                mode: "middle",
                status_text: "",
                status_color: PRIMARY_COLOR_LIST.PRIMARY_COLOR_RGB
            };
            this.$refs.form.clearValidate();
            this.dialogVisible = false;
        },
        onConfirm() {
            this.$refs.form.validate((flag) => {
                let mapping = {
                    add: "createNode",
                    edit: "updateNode"
                };
                let status_color = this.form.status_color || "";
                if (flag) {
                    let params = {
                        ws_id: this.curSpace.id,
                        tmpl_id: this.$route.params.id,
                        ...this.form,
                        status_color
                    };
                    if (this.curType === "edit") {
                        params.id = this.curRow.node.id;
                    }
                    api[mapping[this.curType]](params).then((res) => {
                        if (res && res.resultCode === 200) {
                            this.$emit("on-confirm");
                        }
                    });
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.add-trigger-dialog {
    .el-dialog {
        width: 520px;
        border-radius: 4px;
        .el-dialog__header {
            padding: 24px 32px;
            font-size: 16px;
            color: #2f384c;
        }
        .el-dialog__body {
            padding: 0 32px;
        }
        .el-dialog__footer {
            padding: 0 32px 32px;
        }
        .el-dialog__headerbtn {
            top: 24px;
            right: 32px;
        }
        .el-form-item {
            margin-bottom: 24px;
        }
        .el-dialog__footer {
            text-align: right;
        }
    }
}
</style>
<style lang="scss" scoped>
.add-trigger-dialog {
    .width100 {
        width: 100%;
    }
}
</style>
