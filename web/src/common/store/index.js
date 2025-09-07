import Vue from "vue";
import Vuex from "vuex";
import api from "../api";
import progress_setting from "./progress_setting";
import api_space from "@/common/api/module/space";
import api_progress_setting from "@/common/api/module/progress_setting";
import api_progress_dashboard from "@/common/api/module/progress_dashboard";
import _ from "lodash";
Vue.use(Vuex);

//不是在生产环境debug为true
const debug = process.env.NODE_ENV !== "production";
const curSpace =
    localStorage.getItem("space") &&
    Object.keys(localStorage.getItem("space")).length &&
    localStorage.getItem("space") !== "undefined"
        ? JSON.parse(localStorage.getItem("space"))
        : {};
//创建Vuex实例对象
const store = new Vuex.Store({
    strict: debug, //在不是生产环境下都开启严格模式
    modules: {
        progress_setting
    },
    state: {
        // 取消请求得数组
        cancelTokenArr: [],
        // token
        token: localStorage.getItem("token") || "",
        // 用户信息
        userInfo: {},
        // 当前空间
        curSpace: curSpace,
        // 空间列表
        workSpaceList: [],
        // 左侧树列表(根据不同空间变化)
        progressTree: [],
        // 流程角色列表
        progressRoleList: [],
        // 字段配置
        filedConfig: {},
        curPrswitchProgressogressNodeData: {},
        // 是否手动切换流程或者流程更多中的选项等
        switchProgress: false,
        workspaceHomePageInfo: {},
        // 删除还是添加流程
        deleteOrAddTree: {
            action: "", // 删除: 'delete', 增加: 'add'
            id: "",
            type: "", // 文件夹还是文件 'tmpl' 或者 'ws_file'
            file_id: "" // 文件夹id
        },
        // progress的dashboard
        dashboard: {
            config: {}
        }
    },
    getters: {
        token: (state) => state.token,
        userInfo: (state) => state.userInfo
    },
    mutations: {
        pushToken(state, cancel) {
            state.cancelTokenArr.push(cancel);
        },
        clearToken(state) {
            state.cancelTokenArr.length &&
                state.cancelTokenArr.forEach((c) => {
                    if (c) {
                        c();
                    }
                });
            state.cancelTokenArr = [];
        },
        setToken(state, token) {
            state.token = token;
            localStorage.setItem("token", token);
        },
        setRefreshToken(state, token) {
            state.refreshToken = token;
            localStorage.setItem("refresh_token", token);
        },
        setCurSpace(state, curSpace) {
            localStorage.setItem("space", JSON.stringify(curSpace));
            state.curSpace = _.cloneDeep(curSpace);
        },
        setWorkSpaceList(state, list) {
            state.workSpaceList = _.cloneDeep(list);
        },
        setProgressTree(state, list) {
            state.progressTree = _.cloneDeep(list);
        },
        delToken(state) {
            state.token = "";
            localStorage.removeItem("token");
        },
        setUserInfo(state, userInfo) {
            state.userInfo = _.cloneDeep(userInfo);
        },
        setProgressRoleList(state, progressRoleList) {
            state.progressRoleList = _.cloneDeep(progressRoleList);
        },
        setFiledConfig(state, filedConfig) {
            state.filedConfig = _.cloneDeep(filedConfig);
        },
        // 存储当前流程模板数据
        setCurProgressNodeData(state, nodeData) {
            state.curProgressNodeData = _.cloneDeep(nodeData);
        },
        // 是否手动切换流程或者流程更多中的选项等
        switchProgress(state, flag) {
            state.switchProgress = flag;
        },
        // 存储空间首页信息
        setWorkspaceHomePageInfo(state, data) {
            state.workspaceHomePageInfo = _.cloneDeep(data);
        },
        deleteOrAddTree(state, params) {
            state.deleteOrAddTree = {
                action: params.action || "", // 删除: 'delete', 增加: 'add'
                id: params.id || "", // 流程或者文件夹的id
                type: params.type || "", // 文件夹还是文件 'tmpl' 或者 'ws_file'
                file_id: params.file_id || "" // 文件夹id
            };
        },
        // 存储dashboard的config
        setDashboardConfig(state, dashboardConfig) {
            state.dashboard.config = _.cloneDeep(dashboardConfig);
        }
    },
    actions: {
        doLogin(context, params) {
            return new Promise((resolve, reject) => {
                api.login(params)
                    .then((res) => {
                        if (
                            res &&
                            res.resultCode === 200 &&
                            res.data &&
                            Object.keys(res.data).length
                        ) {
                            let userData = res.data.user || {};
                            let token = res.data.token || "";
                            let refreshToken = res.data.refresh_token || "";
                            if (token) {
                                context.commit("setToken", token);
                                context.commit("setRefreshToken", refreshToken);
                            }
                            if (userData && Object.keys(userData).length) {
                                context.commit("setUserInfo", userData);
                                resolve(userData);
                            }
                        }
                    })
                    .catch((err) => {
                        reject(err);
                    });
            });
        },
        fetchUserInfo(context) {
            return new Promise((resolve, reject) => {
                api.getUserInfo()
                    .then((res) => {
                        if (
                            res &&
                            res.resultCode === 200 &&
                            res.data &&
                            Object.keys(res.data).length
                        ) {
                            let userData = res.data;
                            if (userData && Object.keys(userData).length) {
                                context.commit("setUserInfo", userData);
                                resolve(userData);
                            }
                        }
                    })
                    .catch((err) => {
                        reject(err);
                    });
            });
        },
        // 获取空间列表
        fetchSpaceList(context) {
            return new Promise((resolve, reject) => {
                api_space
                    .getWorkspaceList()
                    .then((res) => {
                        if (res && res.resultCode === 200) {
                            if (res.data && res.data.length) {
                                context.commit("setWorkSpaceList", res.data);
                                resolve(res.data);
                            } else {
                                context.commit("setWorkSpaceList", []);
                                resolve([]);
                            }
                        } else {
                            context.commit("setWorkSpaceList", []);
                            resolve([]);
                        }
                    })
                    .catch((err) => {
                        reject(err);
                    });
            });
        },
        // 获取流程树列表
        fetchProgressTree(context, params) {
            return new Promise((resolve, reject) => {
                api_progress_setting
                    .getProgressTree(params)
                    .then((res) => {
                        if (res && res.resultCode === 200) {
                            if (res.data && res.data.length) {
                                context.commit("setProgressTree", res.data);
                                resolve(res.data);
                            } else {
                                context.commit("setProgressTree", []);
                                resolve([]);
                            }
                        } else {
                            context.commit("setProgressTree", []);
                            resolve([]);
                        }
                    })
                    .catch((err) => {
                        reject(err);
                    });
            });
        },
        // 获取当前流程的角色列表
        fetchProgressRoleList(context, params) {
            return new Promise((resolve, reject) => {
                api_progress_setting
                    .getRoleList(params)
                    .then((res) => {
                        if (res && res.resultCode === 200 && res.data) {
                            this.roleTableData = res.data.list || [];
                            let list = res.data.list;
                            if (list && list.length) {
                                context.commit("setProgressRoleList", list);
                                resolve({
                                    list: list,
                                    total: res.data.cnt
                                });
                            } else {
                                context.commit("setProgressRoleList", []);
                                resolve({
                                    list: [],
                                    total: 0
                                });
                            }
                        } else {
                            context.commit("setProgressRoleList", []);
                            resolve({
                                list: [],
                                total: 0
                            });
                        }
                    })
                    .catch((err) => {
                        reject(err);
                    });
            });
        },
        // 获取字段config
        fetchFiledConfig(context) {
            return new Promise((resolve, reject) => {
                api_progress_setting
                    .getFieldConfig()
                    .then((res) => {
                        if (res && res.resultCode === 200 && res.data) {
                            context.commit("setFiledConfig", res.data || {});
                            resolve(res.data || {});
                        } else {
                            context.commit("setFiledConfig", {});
                            resolve({});
                        }
                    })
                    .catch((err) => {
                        reject(err);
                    });
            });
        },
        // 获取dashboard的config
        fetchDashboardConfig(context, params) {
            return new Promise((resolve, reject) => {
                api_progress_dashboard
                    .getDashboardConfig(params)
                    .then((res) => {
                        if (
                            res &&
                            res.resultCode === 200 &&
                            res.data &&
                            Object.keys(res.data).length
                        ) {
                            context.commit("setDashboardConfig", res.data);
                            resolve(res.data);
                        } else {
                            context.commit("setDashboardConfig", {});
                            resolve({});
                        }
                    })
                    .catch((err) => {
                        reject(err);
                    });
            });
        }
    }
});
export default store;
