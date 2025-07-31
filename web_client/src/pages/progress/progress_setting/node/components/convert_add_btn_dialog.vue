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
            label-width="84px"
        >
            <el-form-item prop="name" label="转换名称:">
                <el-input
                    class="width"
                    v-model="form.name"
                    maxlength="25"
                    show-word-limit
                    placeholder="请输入转换名称"
                ></el-input>
            </el-form-item>
            <el-form-item prop="target_node_id" label="目标节点:">
                <el-select
                    class="width100 basic-ui"
                    v-model="form.target_node_id"
                    placeholder="请选择目标节点"
                >
                    <el-option
                        v-for="item in nodeList"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    >
                    </el-option>
                </el-select>
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
import _ from "lodash";
export default {
    computed: {
        curSpace() {
            return this.$store.state.curSpace || {};
        }
    },
    data() {
        return {
            dialogVisible: false,
            form: {
                name: "",
                target_node_id: ""
            },
            rule: {
                name: [
                    {
                        required: true,
                        message: "请输入转换名称"
                    }
                ],
                target_node_id: [
                    {
                        required: true,
                        message: "请选择目标节点"
                    }
                ]
            },
            nodeList: [],
            title: "",
            curType: "",
            curRow: {
                node: {}
            },
            curBtnItem: {}
        };
    },

    methods: {
        openDialog(item, row, type, nodeList) {
            this.curBtnItem = _.cloneDeep(item);
            this.curRow = _.cloneDeep(row);
            this.curType = type;
            this.nodeList = this.formatNodeList(nodeList);
            this.dialogVisible = true;
            if (type === "add") {
                this.title = "添加转换按钮";
            } else {
                this.title = "编辑转换按钮";
                this.form.name = item.name;
                this.form.target_node_id = item.target_node_id;
            }
            this.$nextTick(() => {
                this.$refs.form.clearValidate();
            });
        },
        // 格式化目标节点下拉选项
        formatNodeList(list) {
            let tmpArr = [];
            if (list && list.length) {
                list.forEach((item) => {
                    if (item.node && Object.keys(item.node).length) {
                        let node = item.node;
                        tmpArr.push({
                            label: node.name,
                            value: node.id,
                            key: node.id
                        });
                    }
                });
            }
            return tmpArr;
        },
        cancel() {
            this.$refs.form.clearValidate();
            this.dialogVisible = false;
            this.form = {
                name: "",
                target_node_id: ""
            };
        },
        onConfirm() {
            this.$refs.form.validate((flag) => {
                if (flag) {
                    let mapping = {
                        add: "createBtn",
                        edit: "updateBtn"
                    };
                    let params = {
                        ws_id: this.curSpace.id,
                        tmpl_id: this.$route.params.id,
                        node_id: this.curRow.node.id,
                        ...this.form
                    };
                    if (this.curType === "edit") {
                        params.id = this.curBtnItem.id;
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
.width100.el-select {
    width: 100%;
}
.add-trigger-dialog {
    .width100 {
        width: 100%;
    }
}
</style>
