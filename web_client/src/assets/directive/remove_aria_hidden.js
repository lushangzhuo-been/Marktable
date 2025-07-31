import Vue from "vue";
/* eslint-disable */

Vue.directive("removeAriaHidden", {
    bind(el, binding) {
        // let ariaEls = el.querySelectorAll(
        //     ".el-radio__original",
        //     "el-popover el-popper"
        // );
        // let ariaEls = el.querySelectorAll(
        //     ".el-radio__original",
        //     ".el-popover el-popper"
        // );
        let ariaEls = el.getElementsByClassName("el-popover");
        // ariaEls.forEach((item) => {
        //     item.removeAttribute("aria-hidden");
        // });
        Array.prototype.forEach.call(ariaEls, function (element, index) {
            // 处理每个元素
            element.removeAttribute("aria-hidden");
        });
    }
});
