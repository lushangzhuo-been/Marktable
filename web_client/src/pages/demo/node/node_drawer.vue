<template>
    <div>
        <el-drawer
            :visible.sync="drawer"
            direction="rtl"
            :size="'60%'"
            :before-close="handleClose"
            :append-to-body="true"
        >
            <div>节点抽屉</div>
            <!-- 顶部切换按钮 -->
            <el-tabs v-model="tabName">
                <el-tab-pane label="转换界面" name="转换界面"></el-tab-pane>
                <el-tab-pane label="触发器" name="触发器"></el-tab-pane>
                <el-tab-pane label="条件" name="条件"></el-tab-pane>
            </el-tabs>
            <!-- 内容 -->
            <div v-if="tabName === '转换界面'">
                <!-- 选择器 与 转换意见 -->
                <div>
                    <div class="field_suggest">
                        <div class="field">
                            选择字段：
                            <el-select
                                class="basic-ui"
                                v-model="field"
                                placeholder="请选择字段"
                            >
                                <el-option
                                    v-for="item in fieldOptions"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.value"
                                >
                                </el-option>
                            </el-select>
                        </div>
                        <div class="suggest">
                            转换意见
                            <el-switch
                                v-model="suggest"
                                active-color="#13ce66"
                                inactive-color="#e6e9f0"
                            >
                            </el-switch>
                        </div>
                    </div>
                    <div></div>
                </div>
            </div>
            <div v-if="tabName === '触发器'">
                <!-- 添加触发器 -->
                <div class="trigger_condition">
                    <el-button class="basic-ui" size="small" type="primary"
                        >增加触发器</el-button
                    >
                </div>
            </div>
            <div v-if="tabName === '条件'">
                <div class="trigger_condition">
                    <el-button class="basic-ui" size="small" type="primary"
                        >增加条件</el-button
                    >
                </div>
            </div>
        </el-drawer>
    </div>
</template>

<script>
export default {
    components: {},
    props: {},
    data() {
        return {
            drawer: false,
            tabName: "转换界面",
            fieldOptions: [
                {
                    value: "选项1",
                    label: "黄金糕"
                },
                {
                    value: "选项2",
                    label: "双皮奶"
                },
                {
                    value: "选项3",
                    label: "蚵仔煎"
                },
                {
                    value: "选项4",
                    label: "龙须面"
                },
                {
                    value: "选项5",
                    label: "北京烤鸭"
                }
            ],
            field: "",
            suggest: false
        };
    },
    watch: {},
    methods: {
        openDrawer() {
            this.drawer = true;
        },
        handleClose(done) {
            this.$confirm("确认关闭？")
                .then((_) => {
                    done();
                })
                .catch((_) => {});
        }
    }
};
</script>

<style lang="scss" scoped>
.field_suggest {
    display: flex;
    justify-content: space-between;
}
.trigger_condition {
    display: flex;
    justify-content: right;
}
</style>
