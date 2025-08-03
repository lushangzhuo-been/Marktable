import Vue from "vue";
import store from "./common/store";
import App from "./App.vue";
import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";
import router from "./router/index.js";
import "./assets/css/main.scss";
import { JumpDialogBox } from "@/assets/tool/login";
import "@/assets/directive/overflow";
import "@/assets/directive/remove_aria_hidden";
import _ from "lodash";
Vue.config.productionTip = false;

const echarts = require("echarts");
Vue.prototype.$echarts = echarts;
Vue.prototype.$JumpDialog = JumpDialogBox.install;
// Vue.prototype.$Reg = regBox.install;
Vue.use(ElementUI);
new Vue({
    router,
    store,
    render: (h) => h(App)
}).$mount("#app");
