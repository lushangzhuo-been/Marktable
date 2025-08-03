<template>
    <div class="list-content">
        <sub-task-item
            v-for="(item, index) in subTaskConfig"
            :key="index"
            :detailId="detailId"
            :subTaskInfo="item"
            @refresh-sub-task-count="refreshSubTaskCount"
        ></sub-task-item>
    </div>
</template>

<script>
import subTaskItem from "./sub_task_item.vue";
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
    methods: {
        refreshSubTaskCount() {
            this.$emit("refresh-sub-task-count");
        }
    }
};
</script>

<style lang="scss" scoped>
.list-content {
    padding-bottom: 20px;
}
</style>
