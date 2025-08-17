<template>
    <div class="list-content">
        <sub-task-item
            v-for="(item, index) in subTaskConfig"
            :key="index"
            :detailId="detailId"
            :subTaskInfo="item"
            :subTmplAuth="subTmplAuth"
            @refresh-sub-task-count="refreshSubTaskCount"
        ></sub-task-item>
    </div>
</template>

<script>
import subTaskItem from "./sub_task_item.vue";
import api from "@/common/api/module/progress";
export default {
    components: {
        subTaskItem
    },
    props: {
        subTaskConfig: {
            type: Array,
            default: () => []
        },
        detailId: {
            type: String,
            default: ""
        }
    },
    data() {
        return {
            subTmplAuth: false
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
    methods: {
        refreshSubTaskCount() {
            this.$emit("refresh-sub-task-count");
        },
        // 获取子任务的权限
        getSubTmplAuth() {
            let params = {
                ws_id: this.curSpace.id,
                tmpl_id: this.curProgress,
                id: this.detailId,
                auth_mode: "sub_tmpl"
            };
            api.getUserAuth(params).then((res) => {
                if (res && res.resultCode === 200) {
                    this.subTmplAuth = res.data;
                } else {
                    this.subTmplAuth = false;
                }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.list-content {
    padding-bottom: 20px;
}
</style>
