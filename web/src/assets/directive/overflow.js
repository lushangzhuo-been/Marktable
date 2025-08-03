import Vue from "vue";
/* eslint-disable */

function getStyle(obj, attr) {
    if (obj.currentStyle) {
        return obj.currentStyle[attr];
    } else {
        return getComputedStyle(obj)[attr];
    }
}

Vue.directive("overflow", {
    inserted(el, binding, vnode, oldVnode) {
        // 获取app.vue根节点下的 ref为tooltip的el-tooltip组件
        const tooltip = vnode.context.$root._vnode.child.$refs.tooltip;
        // 给添加指令的el对象添加鼠标移入事件
        el.__vueOverflowTooltipMouseenter__ = function () {
            // 获取el.childNodes长度判断是否触发Range对象计算宽度
            const childNodesLength = el.childNodes.length;
            if (childNodesLength) {
                // 使用 range 代替 scrollWidth 来判断文本是否溢出，这样做是为了解决潜在的 bug。
                const range = document.createRange();
                // 设置 range 的起点
                range.setStart(el, 0);
                // 设置 range 的终点，因为起终点都在同一个节点上，所以设置终点偏移量以选中节点的内容
                range.setEnd(el, childNodesLength);
                // 获取节点的内容的宽度
                const rangeWidth = range.getBoundingClientRect().maxWidth;
                // 获取el盒模型
                const boxSizing = getStyle(el, "boxSizing");
                // 获取el左右pandding
                const padding =
                    boxSizing === "border-box"
                        ? 0
                        : parseInt(getStyle(el, "paddingLeft"), 10) +
                        parseInt(getStyle(el, "paddingRight"), 10);

                if (
                    rangeWidth + padding > el.offsetWidth ||
                    el.scrollWidth > el.offsetWidth
                ) {
                    // 改变app.vue根节点的data里tooltipContent 值（给el-tooltip组件的content设置值）
                    tooltip.$parent.$data.tooltipContent =
                        el.innerText || el.textContent;
                    // 关联el
                    tooltip.referenceElm = el;
                    // 隐藏之前打开的popper
                    tooltip.$refs.popper &&
                        (tooltip.$refs.popper.style.display = "none");
                    tooltip.doDestroy();
                    tooltip.setExpectedState(true);
                    tooltip.handleShowPopper();
                }
            }
        };
        // 给添加指令的dom对象添加鼠标移出事件
        el.__vueOverflowTooltipMouseleave__ = function () {
            tooltip.setExpectedState(false);
            tooltip.handleClosePopper();
        };

        // 绑定事件
        el.addEventListener("mouseenter", el.__vueOverflowTooltipMouseenter__);
        el.addEventListener("mouseleave", el.__vueOverflowTooltipMouseleave__);
    },
    update: function () { },
    componentUpdated: function () { },
    unbind: function (el) {
        el.removeEventListener(
            "mouseenter",
            el.__vueOverflowTooltipMouseenter__
        );
        el.removeEventListener(
            "mouseleave",
            el.__vueOverflowTooltipMouseleave__
        );
        delete el.__vueOverflowTooltipMouseenter__;
        delete el.__vueOverflowTooltipMouseleave__;
    },
});
