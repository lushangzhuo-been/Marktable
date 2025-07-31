<template>
    <div class="count-block" :id="item.i">
        <countTo
            class="number"
            :style="{ fontSize: this.numSize + 'px' }"
            :startVal="startVal"
            :endVal="endVal"
            :duration="duration"
            :separator="separator"
            ref="CountToPower"
        >
        </countTo>
        <div class="title" :style="{ fontSize: this.titleSize + 'px' }">
            目标收入
        </div>
    </div>
</template>

<script>
import countTo from "vue-count-to";
export default {
    components: { countTo },
    props: {
        item: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            startVal: 0,
            endVal: 234100,
            duration: 1400, // 持续时间
            separator: ",", // 分隔符
            numSize: 42,
            titleSize: 22
        };
    },
    mounted() {
        // 选择你需要监听的元素
        const element = document.getElementById(this.item.i);

        // 创建一个ResizeObserver实例并定义回调函数
        const resizeObserver = new ResizeObserver((entries) => {
            for (const entry of entries) {
                const width = entry.contentRect.width;
                const height = entry.contentRect.height;
                this.setSize(width, height);
            }
        });
        // 开始监听元素的尺寸变化
        resizeObserver.observe(element);
    },
    methods: {
        setSize(width, height) {
            if (width < height) {
                this.numSize = Math.ceil(width / 10);
                this.titleSize = Math.ceil(width / 14);
            } else {
                this.numSize = Math.ceil(height / 10);
                this.titleSize = Math.ceil(height / 14);
            }
        }
    }
};
</script>

<style lang="scss" scoped>
.count-block {
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;

    .number {
        font-size: 42px;
        text-align: center;
    }

    .title {
        font-size: 22px;
        text-align: center;
    }
}
</style>
