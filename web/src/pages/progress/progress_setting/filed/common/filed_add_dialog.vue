<template>
    <el-dialog
        title="新增字段"
        :visible.sync="dialogVisible"
        class="filed-add-dialog basic-ui"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        append-to-body
        @close="cancel"
    >
        <el-form
            ref="addFiledForm"
            :model="addFiledForm"
            :rules="rule"
            :label-width="getLabelWidth"
            size="small"
        >
            <el-form-item prop="mode" label="类型:">
                <el-select
                    v-model="addFiledForm.mode"
                    class="basic-ui width100"
                    :class="
                        addFiledForm.mode ? 'type-radio' : 'type-radio-empty'
                    "
                    placeholder="请选择类型"
                    size="small"
                    @change="$refs.addFiledForm.clearValidate()"
                >
                    <i slot="prefix" v-show="addFiledForm.mode">
                        <img
                            style="position: relative; top: 5px; left: 6px"
                            :src="typeIconMapping[addFiledForm.mode]"
                            alt=""
                            width="20px"
                            height="20px"
                        />
                    </i>
                    <el-option
                        v-for="item in typeOptions"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    >
                        <span style="display: flex; align-items: center">
                            <img
                                :src="item.icon"
                                alt=""
                                width="20px"
                                height="20px"
                            />
                            <span style="margin-left: 8px">{{
                                item.label
                            }}</span>
                        </span>
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item prop="name" label="名称:">
                <el-input
                    v-model="addFiledForm.name"
                    type="text"
                    class="basic-ui"
                    size="small"
                    maxlength="25"
                    show-word-limit
                    placeholder="请输入名称"
                ></el-input>
            </el-form-item>
            <!-- 关联字段下的模板(关联流程) -->
            <el-form-item
                v-if="addFiledForm.mode === 'related_com'"
                prop="progressValue"
                label="流程:"
            >
                <el-select
                    v-model="addFiledForm.progressValue"
                    class="basic-ui width100"
                    placeholder="请选择/输入关联流程"
                    popper-class="filed-add-dialog-related-popper"
                    size="small"
                    filterable
                    :rexmote-method="loadOptions"
                    :loading="isLoading"
                >
                    <el-option
                        v-for="item in progressOptions"
                        :key="item.id"
                        :label="item.name"
                        :value="item.id"
                    >
                        <span v-overflow class="related-label">{{
                            item.name
                        }}</span>
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="描述:">
                <el-input
                    class="basic-ui width fixed-height-dialog"
                    v-model="addFiledForm.desc"
                    type="textarea"
                    :autosize="{ minRows: 2, maxRows: 6 }"
                    size="small"
                    placeholder="请输入描述"
                ></el-input>
            </el-form-item>
            <!-- 下拉单选 -->
            <template v-if="addFiledForm.mode === 'select_com'">
                <draggable v-model="typeDrapDownList">
                    <el-form-item
                        class="form-item-radio"
                        v-for="(item, index) in typeDrapDownList"
                        :key="index"
                        @mouseenter.native="mouseenter(item)"
                        @mouseleave.native="mouseleave(item)"
                    >
                        <span slot="label" style="position: relative">
                            <img
                                v-show="item.isShow"
                                :src="
                                    require('@/assets/image/common/icon_dragtable_move.png')
                                "
                                alt=""
                                width="24px"
                                height="24px"
                                style="
                                    position: absolute;
                                    left: -26px;
                                    top: 4px;
                                "
                            />
                            <!-- <span>{{ item.label }}</span> -->
                        </span>
                        <el-input
                            class="basic-ui width drapdown"
                            v-model="item.value"
                            type="text"
                            size="small"
                            placeholder="请输入选项"
                        >
                            <template slot="suffix">
                                <span
                                    style="
                                        display: flex;
                                        align-items: center;
                                        position: relative;
                                        top: 4px;
                                        left: -4px;
                                    "
                                >
                                    <el-color-picker
                                        :ref="`colorPicker${index}`"
                                        class="basic-ui"
                                        v-model="item.color"
                                        popper-class="radioColorPicker"
                                        :predefine="COLOR_PICKER_LIST"
                                        @active-change="
                                            (val) =>
                                                colorActiveChange(
                                                    val,
                                                    item,
                                                    index
                                                )
                                        "
                                    ></el-color-picker>
                                    <img
                                        class="remove-item"
                                        :src="
                                            require('@/assets/image/common/delete.svg')
                                        "
                                        alt=""
                                        width="16px"
                                        height="16px"
                                        @click="
                                            removeRadio(
                                                'radio',
                                                typeDrapDownList,
                                                index
                                            )
                                        "
                                    />
                                </span>
                            </template>
                        </el-input>
                    </el-form-item>
                </draggable>
                <el-form-item>
                    <span
                        class="add-item"
                        @click="addRadio('select', typeDrapDownList)"
                    >
                        <b class="el-icon-circle-plus-outline"></b>
                        <span>添加选项</span>
                    </span>
                </el-form-item>
            </template>
            <!-- 下拉多选 -->
            <template v-if="addFiledForm.mode === 'multi_select_com'">
                <draggable v-model="typeLabelList">
                    <el-form-item
                        class="form-item-radio"
                        v-for="(item, index) in typeLabelList"
                        :key="index"
                        @mouseenter.native="mouseenter(item)"
                        @mouseleave.native="mouseleave(item)"
                    >
                        <span slot="label" style="position: relative">
                            <img
                                v-show="item.isShow"
                                :src="
                                    require('@/assets/image/common/icon_dragtable_move.png')
                                "
                                alt=""
                                width="24px"
                                height="24px"
                                style="
                                    position: absolute;
                                    left: -26px;
                                    top: 4px;
                                "
                            />
                            <!-- <span>{{ item.label }}</span> -->
                        </span>
                        <el-input
                            class="basic-ui width drapdown"
                            v-model="item.value"
                            type="text"
                            size="small"
                            placeholder="请输入选项"
                        >
                            <template slot="suffix">
                                <span
                                    style="
                                        display: flex;
                                        align-items: center;
                                        position: relative;
                                        top: 4px;
                                        left: -4px;
                                    "
                                >
                                    <el-color-picker
                                        :ref="`colorPicker${index}`"
                                        class="basic-ui"
                                        v-model="item.color"
                                        popper-class="radioColorPicker"
                                        :predefine="COLOR_PICKER_LIST"
                                        @active-change="
                                            (val) =>
                                                colorActiveChange(
                                                    val,
                                                    item,
                                                    index
                                                )
                                        "
                                    ></el-color-picker>
                                    <img
                                        class="remove-item"
                                        :src="
                                            require('@/assets/image/common/delete.svg')
                                        "
                                        alt=""
                                        width="16px"
                                        height="16px"
                                        @click="
                                            removeRadio(
                                                'radio',
                                                typeLabelList,
                                                index
                                            )
                                        "
                                    />
                                </span>
                            </template>
                        </el-input>
                    </el-form-item>
                </draggable>
                <el-form-item>
                    <span
                        class="add-item"
                        @click="addRadio('select', typeLabelList)"
                    >
                        <b class="el-icon-circle-plus-outline"></b>
                        <span>添加选项</span>
                    </span>
                </el-form-item>
            </template>
            <!-- 进度条 -->
            <!-- <template v-if="addFiledForm.mode === 'progress_com'">
                <draggable v-model="typeProgressList">
                    <el-form-item
                        v-for="(item, index) in typeProgressList"
                        :key="index"
                    >
                        <div style="display: inline-block">
                            <span>当状态为</span>
                            <el-select
                                v-model="item.status"
                                clearable
                                class="basic-ui"
                                size="small"
                                style="width: 100px; margin: 0 6px"
                                placeholder="请选择"
                            >
                                <el-option
                                    v-for="item in typeProgressListOptions"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.value"
                                >
                                </el-option> </el-select
                            >时，进度显示
                            <el-input
                                class="basic-ui progress"
                                placeholder="请输入"
                                size="small"
                                v-model="item.progress"
                                style="width: 104px; margin: 0 6px"
                            >
                                <template slot="append">%</template>
                            </el-input>
                        </div>
                        <img
                            class="remove-item position-ab"
                            :src="require('@/assets/image/common/delete.svg')"
                            alt=""
                            width="16px"
                            height="16px"
                            @click="
                                removeRadio(
                                    'radio',
                                    addFiledForm.typeRadio,
                                    index
                                )
                            "
                        />
                    </el-form-item>
                </draggable>
                <el-form-item>
                    <span
                        class="add-item"
                        @click="addRadio('progress', typeProgressList)"
                    >
                        <b class="el-icon-circle-plus-outline"></b>
                        <span>添加选项</span>
                    </span>
                </el-form-item>
            </template> -->
            <el-form-item
                v-if="addFiledForm.mode === 'number_com'"
                label="单位:"
            >
                <div>
                    <el-radio
                        v-model="numberRadio"
                        v-for="item in numberUnitList"
                        :key="item.value"
                        :label="item.value"
                        >{{ item.label }}</el-radio
                    >
                </div>
                <el-input
                    class="basic-ui"
                    size="small"
                    v-if="numberRadio === 'self'"
                    v-model="numberRadioUnit"
                    placeholder="请输入"
                ></el-input>
            </el-form-item>
            <!-- 时间 -->
            <el-form-item v-if="addFiledForm.mode === 'time_com'" label="">
                <!-- <div>
                    <el-radio
                        v-model="timeRadio"
                        v-for="item in timeRadioOptions"
                        :key="item.value"
                        :label="item.value"
                        >{{ item.label }}</el-radio
                    >
                </div> -->
                <span class="example">示意：2023-05-31 14:34:17</span>
            </el-form-item>
            <!-- 日期 -->
            <el-form-item v-if="addFiledForm.mode === 'date_com'" label="">
                <!-- <div>
                    <el-radio
                        v-model="timeRadio"
                        v-for="item in timeRadioOptions"
                        :key="item.value"
                        :label="item.value"
                        >{{ item.label }}</el-radio
                    >
                </div> -->
                <span class="example">示意：2023-05-23</span>
            </el-form-item>
            <!-- 单多人 -->
            <el-form-item
                v-if="addFiledForm.mode === 'person_com'"
                label="单/多人:"
            >
                <div>
                    <el-radio
                        v-model="personRadio"
                        v-for="item in personRadioOptions"
                        :key="item.value"
                        :label="item.value"
                        >{{ item.label }}</el-radio
                    >
                </div>
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
            <el-button size="small" @click="cancel">取 消</el-button>
            <el-button
                class="basic-ui"
                size="small"
                type="primary"
                @click="onConfirm"
                >确 定</el-button
            >
        </span>
    </el-dialog>
</template>

<script>
import api from "@/common/api/module/progress_setting";
import api_common from "@/common/api";
import draggable from "vuedraggable";
import dataHandle from "../data_handle";
import _ from "lodash";
import {
    PRIMARY_COLOR_LIST,
    COLOR_PICKER_LIST
} from "@/assets/tool/color_list";
export default {
    components: {
        draggable
    },
    computed: {
        getLabelWidth() {
            return this.addFiledForm.mode === "person_com" ? "72px" : "60px";
        },
        curSpace() {
            return this.$store.state.curSpace || {};
        }
    },
    data() {
        return {
            COLOR_PICKER_LIST: COLOR_PICKER_LIST,
            dialogVisible: false,
            // 类型iconyings
            typeIconMapping: _.cloneDeep(dataHandle.typeIconMapping),
            // 下拉多选
            typeDrapDownList: _.cloneDeep(dataHandle.typeDrapDownList),
            // 标签
            typeLabelList: _.cloneDeep(dataHandle.typeLabelList),
            // 进度条(进展)
            // typeProgressList: _.cloneDeep(dataHandle.typeProgressList), //数组
            // typeProgressListOptions: _.cloneDeep(
            //     dataHandle.typeProgressListOptions
            // ), //下拉选项
            // 类型
            typeOptions: [],
            // 数字
            numberUnitList: [],
            numberRadio: "None",
            numberRadioUnit: "",
            // 人员
            personRadioOptions: [],
            personRadio: "single",
            // 时间
            // timeRadioOptions: [],
            // timeRadio: "Ymd",
            // 表单默认字段
            addFiledForm: {
                mode: "",
                name: "",
                desc: "",
                progressValue: "" // 关联流程
            },
            progressOptions: [],
            rule: {
                mode: [
                    {
                        required: true,
                        message: "请选择类型",
                        trigger: "change"
                    }
                ],
                name: [
                    {
                        required: true,
                        message: "请输入名称",
                        trigger: "blur"
                    }
                ],
                desc: [
                    {
                        required: true,
                        message: "请输入描述",
                        trigger: "blur"
                    }
                ],
                progressValue: [
                    {
                        required: true,
                        message: "请选择关联流程",
                        trigger: ["blur", "change"]
                    }
                ]
            },
            timer: null,
            isLoading: false
        };
    },
    watch: {
        "addFiledForm.mode"(val) {
            if (val === "related_com") {
                this.fetchOtherTmplList();
            }
        }
    },
    beforeDestroy() {
        if (this.timer) {
            clearTimeout(this.timer);
            this.timer = null;
        }
    },
    methods: {
        // 该方法生效，需要添加remote属性
        loadOptions(key) {
            this.isLoading = true;
            let that = this;
            // 在这里实现远程方法，返回一个 Promise 对象
            this.timer = setTimeout(() => {
                that.fetchOtherTmplList(key);
                that.isLoading = false;
            }, 500);
        },
        fetchOtherTmplList(key) {
            let params = {
                tmpl_id: this.$route.params.id,
                ws_id: this.curSpace.id,
                key: key || ""
            };
            api_common.getOtherTmplList(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.progressOptions = res.data;
                } else {
                    this.progressOptions = [];
                }
            });
        },
        mouseenter(item) {
            this.$set(item, "isShow", true);
        },
        mouseleave(item) {
            this.$set(item, "isShow", false);
        },
        start() {},
        end() {},
        colorActiveChange(val, item, index) {
            this.$set(item, "color", val);
            let ref = this.$refs[`colorPicker${index}`];
            if (ref && ref.length) {
                ref[0].hide();
            }
        },
        // getTimeEample(radio) {
        //     let mapping = {
        //         Ymd: "示意：2023-05-23",
        //         YmdHis: "示意：2023-05-31 14:34:17"
        //     };
        //     return mapping[radio];
        // },
        addRadio(type, arr) {
            let mapping = {
                select: {
                    label: "",
                    value: "",
                    color: PRIMARY_COLOR_LIST.PRIMARY_COLOR_RGB
                },
                progress: {
                    status: "",
                    progress: ""
                }
            };
            if (!arr || !type || !mapping[type]) return;
            // 新增逻辑
            if (type !== "progress") {
                if (arr && arr.length) {
                    let noLists = arr.map((item) => item.item_no);
                    let maxNo = Math.max(...noLists);
                    mapping[type].label = `选项${maxNo + 1}`;
                    mapping[type].item_no = maxNo + 1;
                } else {
                    mapping[type].label = "选项1";
                    mapping[type].item_no = 1;
                }
            }
            arr.push(mapping[type]);
            this.$refs.addFiledForm.clearValidate();
        },
        removeRadio(type, arr, index) {
            if (!arr || !arr.length) return;
            arr.splice(index, 1);
        },

        openDialog(obj) {
            this.dialogVisible = true;
            // this.fetchFiledConfig()
            this.typeOptions = _.cloneDeep(obj.typeOptions);
            this.numberUnitList = _.cloneDeep(obj.numberUnitList);
            // this.timeRadioOptions = _.cloneDeep(obj.timeRadioOptions);
            this.personRadioOptions = _.cloneDeep(obj.personRadioOptions);
            this.$nextTick(() => {
                this.$refs.addFiledForm.clearValidate();
            });
        },
        fetchFiledConfig() {
            this.$store.dispatch("fetchFiledConfig").then((res) => {
                if (res && Object.keys(res).length) {
                    let {
                        mode_config,
                        number_com_config,
                        // time_com_config,
                        person_com_config
                    } = res;
                    this.setTypeConfig(mode_config);
                    this.setNumberConfig(number_com_config);
                    // this.setTimeConfig(time_com_config);
                    this.setPersonConfig(person_com_config);
                } else {
                    this.typeOptions = [];
                    this.numberUnitList = [];
                    // this.timeRadioOptions = [];
                    this.personRadio = [];
                }
            });
        },
        setTypeConfig(arr) {
            if (arr && arr.length) {
                let defaultConfig = _.cloneDeep(dataHandle.typeIconMapping);
                arr.forEach((item) => {
                    for (let key in defaultConfig) {
                        if (item.value === key) {
                            item.icon = defaultConfig[key];
                        }
                    }
                });
                this.typeOptions = _.cloneDeep(arr);
            } else {
                this.typeOptions = [];
            }
        },
        setNumberConfig(arr) {
            if (arr && arr.length) {
                arr.unshift({
                    label: "自定义",
                    value: "self"
                });
                this.numberUnitList = arr;
            } else {
                this.numberUnitList = [];
            }
        },
        // setTimeConfig(arr) {
        //     this.timeRadioOptions = arr && arr.length ? arr : [];
        // },
        setPersonConfig(arr) {
            this.personRadioOptions = arr && arr.length ? arr : [];
        },
        cancel() {
            this.dialogVisible = false;
            // 表单默认字段
            this.addFiledForm = {
                mode: "",
                name: "",
                desc: "",
                progressValue: ""
            };
            this.typeOptions = [];
            // 数字
            this.numberUnitList = [];
            this.numberRadio = "None";
            this.numberRadioUnit = "";
            // 人员
            this.personRadioOptions = [];
            this.personRadio = "single";
            // 时间
            // this.timeRadioOptions = [];
            // this.timeRadio = "Ymd";
            // 关联流程
            // this.progressValue = "";
            //   类型icon映射
            this.typeIconMapping = _.cloneDeep(dataHandle.typeIconMapping);
            // 下拉
            this.typeDrapDownList = _.cloneDeep(dataHandle.typeDrapDownList);
            // 标签
            this.typeLabelList = _.cloneDeep(dataHandle.typeLabelList);
            // 进度条
            // this.typeProgressList = _.cloneDeep(dataHandle.typeProgressList);
            // this.typeProgressListOptions = _.cloneDeep(
            //     dataHandle.typeProgressListOptions
            // );

            // this.$refs.addFiledForm.clearValidate()
        },
        onConfirm() {
            this.$refs.addFiledForm.validate((flag) => {
                let select = ["multi_select_com", "select_com"];
                let radio = ["person_com", "time_com", "number_com"];
                let mapping = {
                    multi_select_com: this.typeLabelList,
                    select_com: this.typeDrapDownList,
                    // progress_com: this.typeProgressList,
                    person_com: this.personRadio,
                    // time_com: this.timeRadio,
                    number_com: this.numberRadio
                };
                if (flag) {
                    let params = {
                        ws_id: this.curSpace.id,
                        tmpl_id: this.$route.params.id,
                        ...this.addFiledForm
                    };
                    // 单/多选
                    if (select.indexOf(this.addFiledForm.mode) > -1) {
                        params.enum_values = JSON.stringify(
                            mapping[this.addFiledForm.mode]
                        );
                    }
                    // 数字/人员/时间
                    if (radio.indexOf(this.addFiledForm.mode) > -1) {
                        if (
                            this.addFiledForm.mode === "number_com" &&
                            mapping[this.addFiledForm.mode] === "self"
                        ) {
                            params.expr = this.numberRadioUnit;
                        } else {
                            params.expr = mapping[this.addFiledForm.mode];
                        }
                    }
                    // 关联关系
                    if (this.addFiledForm.mode === "related_com") {
                        params.related_tmpl_id = params.progressValue;
                    } else {
                        delete params.progressValue;
                    }
                    api.createField(params).then((res) => {
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

<style lang="scss">
.filed-add-dialog-related-popper {
    width: 396px !important;
    .el-select-dropdown__empty {
        margin: 10px 0 16px 0;
    }
}
.filed-add-dialog {
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
        // .el-icon-close:before {
        //     content: url('~@/assets/image/common/close.png');
        // }
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
.related-label {
    display: inline-block;
    width: 100%;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}
.filed-add-dialog {
    .example {
        color: var(--GRAY_FONT_COLOR);
        font-size: 12px;
    }
    .width100 {
        width: 100%;
    }
    .pointer {
        cursor: pointer;
    }
    .add-item {
        cursor: pointer;
        display: flex;
        align-items: center;
        width: 100px;
        .el-icon-circle-plus-outline {
            color: var(--PRIMARY_COLOR);
            margin-right: 8px;
            font-size: 18px;
        }
        span {
            color: var(--PRIMARY_COLOR);
        }
    }
    .remove-item {
        cursor: pointer;
        margin-left: 8px;
        vertical-align: middle;
    }
    .position-ab {
        position: absolute;
        top: 9px;
        right: 2px;
    }

    .drapdown {
        ::v-deep .el-input-group__append {
            padding: 0 8px 0 16px;
            background-color: #fff;
            width: 62px;
        }
        // .type-radio.el-color-picker {
        //     height: 16px;
        //     ::v-deep .el-color-picker__trigger {
        //         height: 16px;
        //         width: 16px;
        //         padding: 0;
        //         border: none;
        //         .el-color-picker__color {
        //             border: none;
        //             .el-color-picker__color-inner {
        //                 border-radius: 2px;
        //             }
        //         }
        //     }
        // }
        .el-input--suffix .el-input__inner {
            padding-right: 54px;
        }
    }
    .progress ::v-deep .el-input-group__append {
        padding: 0 8px;
    }
    .type-radio {
        ::v-deep .el-input--prefix .el-input__inner {
            padding-left: 36px !important;
        }
    }
    .type-radio-empty {
        ::v-deep .el-input--prefix .el-input__inner {
            padding-left: 15px !important;
        }
    }
}
</style>
