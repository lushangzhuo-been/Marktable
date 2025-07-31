import http from "../http";
import httpNew from "../httpNew";
export default {
    // dashboard
    // 获取config
    getDashboardConfig(params) {
        let url = "/board/config";
        return httpNew.get(url, params);
    },
    // 获取dashboard详情
    getDashboardDetail(params) {
        let url = "/board/detail";
        return httpNew.get(url, params);
    },
    // 获取dashboard列表
    getDashboardList(params) {
        let url = "/board/list";
        return httpNew.get(url, params);
    },
    // 获取dashbord后，加载列表的每一个图
    getChart(params) {
        let url = '/board/statistics'
        return httpNew.post(url, params);
    },
    // 新增或编辑时预览图
    getPrviewChart(params) {
        let url = '/board/statistics/preview'
        return httpNew.post(url, params);
    },
    // 刷新位置
    refresLocation(params) {
        let url = '/board/modify/location'
        return httpNew.post(url, params);
    },
    // 新增
    addDashboardItem(params) {
        let url = "/board/add";
        return httpNew.post(url, params);
    },
    // 复制
    copyDashboardItem(params) {
        let url = "/board/copy";
        return httpNew.post(url, params);
    },
    // 编辑弹窗先获取编辑数据
    getModifyDashboardItem(params) {
        let url = "/board/detail";
        return httpNew.get(url, params);
    },
    // 编辑
    modifyDashboardItem(params) {
        let url = "/board/modify";
        return httpNew.post(url, params);
    },
    // 删除
    deleteDashboardItem(params) {
        let url = "/board/delete";
        return httpNew.post(url, params);
    },

};
