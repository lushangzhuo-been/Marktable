<template>
    <el-tooltip
        :content="curText(text)"
        effect="dark"
        :popper-class="`basic-ui ${myClass}`"
        :placement="position"
        :disabled="overflow"
    >
        <slot></slot>
    </el-tooltip>
</template>

<script>
export default {
    name: 'ellipsisTooltip',
    props: {
        text: [Number, String],
        position: {
            type: String,
            default: 'top',
        },
        myClass: {
            type: String,
            default: '',
        },
    },
    computed: {
        curText() {
            return (text) => {
                if (text === 0) {
                    return '0'
                }
                return text && !isNaN(text) ? text.toString() : text
            }
        },
    },

    data() {
        return {
            overflow: false,
            screenWidth: document.body.clientWidth,
            timer: null,
        }
    },
    watch: {
        text: {
            handler() {
                this.setOverflow()
            },
            immediate: true,
        },
        screenWidth: {
            handler() {
                if (!this.timer) {
                    let that = this
                    that.timer = true
                    setTimeout(() => {
                        that.setOverflow()
                        that.timer = false
                    }, 400)
                }
            },
        },
    },
    mounted() {
        const that = this
        window.addEventListener('resize', () => {
            return (() => {
                window.screenWidth = document.body.clientWidth
                that.screenWidth = window.screenWidth
            })()
        })
    },
    methods: {
        setOverflow() {
            this.$nextTick(() => {
                const node = this.$slots.default[0].elm
                this.overflow = node.clientHeight >= node.scrollHeight
            })
        },
    },
    beforeDestroy() {
        if (this.timer) {
            clearTimeout(this.timer)
            this.timer = null
        }
    },
}
</script>
