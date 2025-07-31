<template>
    <div>
        节点demo
        <div @click="openDrawer">开打抽屉</div>
        <node-drawer ref="NodeDrawer"></node-drawer>
    </div>
</template>

<script>
import NodeDrawer from "./node_drawer.vue";
export default {
    components: {
        NodeDrawer,
    },
    methods: {
        openDrawer() {
            this.$refs.NodeDrawer.openDrawer()
        },
    },
};
</script>

<style lang="scss" scoped></style>
