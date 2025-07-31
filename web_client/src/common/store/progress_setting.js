import api_progress_setting from "@/common/api/module/progress_setting";
const userModule = {
    namespaced: true, // 启用命名空间
    state: {
        createPageFiled: [],
        detailPageFiled: [],
        detailTabFiled: [],
        detailSubFiled: []
    },
    mutations: {
        setCreatePageFiled(state, arr) {  // 创建页面字段
            state.createPageFiled = arr;
        },
        setDetailPageFiled(state, arr) {   // 详情页面，，详情字段
            state.detailPageFiled = arr;
        },
        setDetailTabFiled(state, arr) {  // 详情页面， tab字段
            state.detailTabFiled = arr;
        },
        setDetailSubFiled(state, arr) {  // 详情页面，，子任务字段
            state.detailSubFiled = arr;
        },
    },
    actions: {
        // 获取创建页面字段
        fetchCreatePageFiled(context, params) {
            return new Promise((resolve, reject) => {
                api_progress_setting.getScreenList(params).then((res) => {
                    if (res && res.resultCode === 200) {
                        if (res.data && res.data.length) {
                            context.commit("setCreatePageFiled", res.data);
                            resolve(res.data);
                        } else {
                            context.commit("setCreatePageFiled", []);
                            resolve([]);
                        }
                    } else {
                        context.commit("setCreatePageFiled", []);
                        resolve([]);
                    }
                }).catch((err) => {
                    reject(err);
                });
            });
        },
        // 获取详情页面字段
        fetchDetailPageFiled(context, params) {
            return new Promise((resolve, reject) => {
                api_progress_setting.getScreenList(params).then((res) => {
                    if (res && res.resultCode === 200) {
                        if (res.data && res.data.length) {
                            context.commit("setDetailPageFiled", res.data);
                            resolve(res.data);
                        } else {
                            context.commit("setDetailPageFiled", []);
                            resolve([]);
                        }
                    } else {
                        context.commit("setDetailPageFiled", []);
                        resolve([]);
                    }
                }).catch((err) => {
                    reject(err);
                });
            });
        },
        // 获取详情tabs数组字段
        fetchDetailTabFiled(context, params) {
            return new Promise((resolve, reject) => {
                api_progress_setting.getTabList(params).then((res) => {
                    if (res && res.resultCode === 200) {
                        if (res.data && res.data.length) {
                            context.commit("setDetailTabFiled", res.data);
                            resolve(res.data);
                        } else {
                            context.commit("setDetailTabFiled", []);
                            resolve([]);
                        }
                    } else {
                        context.commit("setDetailTabFiled", []);
                        resolve([]);
                    }
                }).catch((err) => {
                    reject(err);
                });
            });
        },
        // 获取子任务字段
        fetchDetailSubFiled(context, params) {
            return new Promise((resolve, reject) => {
                api_progress_setting.getSubWorkList(params).then((res) => {
                    if (res && res.resultCode === 200) {
                        if (res.data && res.data.length) {
                            context.commit("setDetailSubFiled", res.data);
                            resolve(res.data);
                        } else {
                            context.commit("setDetailSubFiled", []);
                            resolve([]);
                        }
                    } else {
                        context.commit("setDetailSubFiled", []);
                        resolve([]);
                    }
                }).catch((err) => {
                    reject(err);
                });
            });
        }
    },
    getters: {
        createPageFiled: state => state.createPageFiled,
        detailPageFiled: state => state.detailPageFiled
    }
};

export default userModule;