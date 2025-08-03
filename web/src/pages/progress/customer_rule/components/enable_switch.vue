<template>
    <el-switch v-model="enable" @change="enableChange"> </el-switch>
</template>

<script>
export default {
    props: {
        column: {
            type: Object,
            default: () => {}
        },
        row: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            enable: false
        };
    },
    computed: {},
    watch: {
        row: {
            handler(row) {
                this.enable = row[this.column.prop] === "yes" ? true : false;
            },
            immediate: true
        }
    },
    methods: {
        enableChange(val) {
            this.$emit("enable-change", val, this.column, this.row);
        }
    }
};
</script>

<style lang="scss" scoped></style>
