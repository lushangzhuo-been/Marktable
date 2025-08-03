<template>
    <el-dialog
        title="增加触发器"
        :visible.sync="dialogVisible"
        class="add-trigger-dialog basic-ui"
        append-to-body
        @close="cancel"
    >
        <el-form
            size="small"
            ref="form"
            :model="form"
            :rules="rule"
            label-width="124px"
        >
            <el-form-item prop="desc" label="描述:">
                <el-input
                    class="basic-ui width"
                    v-model="form.number"
                    type="textarea"
                    placeholder="请输入描述"
                ></el-input>
            </el-form-item>
            <el-form-item label="编辑条件:">
                <el-button
                    size="small"
                    class="basic-ui"
                    @click="addFilter"
                    plain
                >
                    添加条件
                </el-button>
            </el-form-item>
            <!-- 筛选设置 -->
            <div>
                <!-- 遍历筛选大组，内分小组 , 结构：大数组 套 小数组， 小数组是每项条件， 每项条件可以直接删除，小数组条件也可以直接删除 -->
                <!-- 大数组遍历小数组 -->
                <div
                    class="filter-list-item"
                    v-for="(groupItem, groupIndex) in filterGroup"
                    :key="groupIndex"
                >
                    <div
                        v-for="(filterItem, filterIndex) in groupItem"
                        :key="filterIndex"
                    >
                        <span @click="addFilter" v-show="filterIndex === 0"
                            >与</span
                        >
                        <span
                            v-show="filterIndex === 0"
                            @click="deleteFilterGroup(filterGroup, groupIndex)"
                            >删</span
                        >
                        <!-- 字段 -->
                        <el-select
                            class="basic-ui p-select field"
                            v-model="filterItem.field"
                            placeholder="请选择条件字段"
                        >
                            <el-option
                                v-for="item in fieldOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                            >
                            </el-option>
                        </el-select>
                        <!-- 关系 -->
                        <el-select
                            class="basic-ui p-select relationship"
                            v-model="filterItem.relationship"
                            placeholder="请选择条件关系"
                        >
                            <el-option
                                v-for="item in relationshipOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                            >
                            </el-option>
                        </el-select>
                        <!-- 关键词 -->
                        <el-input
                            class="basic-ui p-input"
                            v-model="filterItem.word"
                            placeholder="请输入"
                        ></el-input>
                        <span
                            v-show="groupItem.length > 1"
                            @click="deleteFilterItem(groupItem, filterIndex)"
                            >删</span
                        >
                        <span @click="addFilterItem(groupItem)">或</span>
                    </div>
                </div>
            </div>
            <el-form-item prop="type" label="触发事件类型:">
                <el-checkbox-group v-model="form.type">
                    <el-checkbox label="触发通知事件"></el-checkbox>
                    <el-checkbox label="触发重新指派事件"></el-checkbox>
                    <el-checkbox label="转化节点"></el-checkbox>
                </el-checkbox-group>
            </el-form-item>
            <!-- 触发通知事件 -->
            <div>
                <el-form-item prop="emailTo" label="发送邮件给:">
                    <el-checkbox-group v-model="form.noticeEmailTo">
                        <el-checkbox label="创建人"></el-checkbox>
                        <el-checkbox label="处理人"></el-checkbox>
                        <el-checkbox label="其他人"></el-checkbox>
                    </el-checkbox-group>
                </el-form-item>
                <el-form-item prop="emailContent" label="邮件内容:">
                    <el-input
                        class="basic-ui width"
                        v-model="form.emailContent"
                        type="textarea"
                        placeholder="请输入邮件内容"
                    ></el-input>
                </el-form-item>
            </div>
            <!-- 触发重新指派事件 -->
            <div>
                <el-form-item prop="assignTo" label="指派给:">
                    <el-radio-group v-model="form.assignTo">
                        <el-radio :label="'创建人'">创建人</el-radio>
                        <el-radio :label="'其他人'">其他人</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item prop="assignToOther" label="">
                    <el-select
                        class="basic-ui"
                        v-model="form.assignToOther"
                        placeholder="请选择"
                    >
                        <el-option
                            v-for="item in assignToOtherList"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        >
                        </el-option>
                    </el-select>
                </el-form-item>
            </div>
            <!-- 设置节点 -->
            <div>
                <el-form-item prop="assignToOther" label="设置节点:">
                    <el-select
                        class="basic-ui"
                        v-model="form.setNode"
                        placeholder="请选择"
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
            </div>
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
export default {
    data() {
        return {
            dialogVisible: false,
            form: {
                desc: "",
                type: ["触发通知事件"],
                // 触发通知事件
                noticeEmailTo: [],
                emailContent: "",
                // 触发重新指派
                assignTo: "",
                setNode: ""
            },
            rule: {
                desc: [
                    {
                        required: true,
                        message: "请输入描述",
                        trigger: "blur"
                    }
                ],
                type: [
                    {
                        required: true,
                        message: "请选择类型"
                    }
                ],
                emailContent: [
                    {
                        required: true,
                        message: "请输入邮件内容",
                        trigger: "blur"
                    }
                ],
                assignTo: [
                    {
                        required: true,
                        message: "请选择指派人"
                    }
                ],
                setNode: [
                    {
                        required: true,
                        message: "请选择节点"
                    }
                ]
            },
            assignToOtherList: [
                {
                    value: "张飞",
                    label: "张飞"
                },
                {
                    value: "关羽",
                    label: "关羽"
                },
                {
                    value: "黄忠",
                    label: "黄忠"
                },
                {
                    value: "赵子龙",
                    label: "赵子龙"
                },
                {
                    value: "马超",
                    label: "马超"
                }
            ],
            nodeList: [
                {
                    value: "开始",
                    label: "开始"
                },
                {
                    value: "进行中",
                    label: "进行中"
                },
                {
                    value: "结束",
                    label: "结束"
                }
            ],
            filterGroup: [],
            fieldOptions: [
                {
                    label: "字段1",
                    value: "字段1"
                },
                {
                    label: "字段2",
                    value: "字段2"
                }
            ],
            relationshipOptions: [
                {
                    label: "大于",
                    value: "大于"
                },
                {
                    label: "等于",
                    value: "等于"
                },
                {
                    label: "小于",
                    value: "小于"
                },
                {
                    label: "包含",
                    value: "包含"
                },
                {
                    label: "排除",
                    value: "排除"
                }
            ]
        };
    },
    methods: {
        addFilter() {
            let groupItem = [
                {
                    field: "",
                    relationship: "",
                    word: ""
                }
            ];
            this.filterGroup.push(groupItem);
        },
        deleteFilterGroup(group, groupIndex) {
            // 删除大组下的小组
            group.splice(groupIndex, 1);
        },
        addFilterItem(group) {
            // 给小组条件添加小项
            let obj = {
                field: "",
                relationship: "",
                word: ""
            };
            group.push(obj);
        },
        deleteFilterItem(group, itemIndex) {
            // 删除小组里的小项
            group.splice(itemIndex, 1);
        },
        openDialog() {
            this.dialogVisible = true;
            this.$nextTick(() => {
                this.$refs.form.clearValidate();
            });
        },
        cancel() {
            this.$refs.form.clearValidate();
            this.form = {
                name: "",
                desciption: ""
            };
            this.dialogVisible = false;
        },
        onConfirm() {
            this.$refs.form.validate((flag) => {
                if (flag) {
                    // this.$emit('on-confirm', this.form)
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.add-trigger-dialog {
    .el-dialog {
        width: 820px;
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
// 筛选条件
.filter-list-item {
    padding-left: 124px;
    border-bottom: 1px solid #2f384c;
}
</style>
<style lang="scss" scoped>
.add-trigger-dialog {
    .width100 {
        width: 100%;
    }
}
</style>
