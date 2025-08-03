import Vue from "vue";
// import Login from './login.vue'
// import Register from './register.vue'
import JumpDialog from "@/pages/common/jump_dialog.vue";
import _ from "lodash";

const JumpDialogBox = Vue.extend(JumpDialog);
// const regBox = Vue.extend(Register)

// LoginBox.install = function () {
//     let instance = new LoginBox({
//     }).$mount()
//     Vue.nextTick(() => {
//         instance.show = true
//     })
// }
JumpDialogBox.install = function (options) {
    if (options === undefined || options === null) {
        options = { isOrganization: false };
    } else {
        options = _.cloneDeep(options);
    }
    let instance = new JumpDialogBox({
        data: options
    }).$mount();
    Vue.nextTick(() => {
        instance.show = true;
    });
};
// export { LoginBox, regBox }

export { JumpDialogBox };
