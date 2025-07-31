import http from "../http";
export default {
    // 规则列表
    getRuleList(params) {
        let url = "/rule/page";
        return http.get(url, params);
    },
    // 日志列表
    getLogList(params) {
        let url = "/rule/rule_log";
        return http.get(url, params);
    },
    // 获取单日志详情
    getLogDetail(params) {
        let url = "/rule/action_log";
        return http.get(url, params);
    },
    // 规则配置
    getRuleConfig(params) {
        let url = "/rule/config";
        return http.get(url, params);
    },
    // 创建规则
    ruleCreate(params) {
        let url = "/rule/create";
        return http.post(url, params);
    },
    // 查询单规则详情
    getRuleDetail(params) {
        let url = "/rule/detail";
        return http.get(url, params);
    },
    // 修改单规则
    ruleUpdate(params) {
        let url = "/rule/update";
        return http.post(url, params);
    },
    // 删除单规则
    ruleDelete(params) {
        let url = "/rule/delete";
        return http.post(url, params);
    },
    // 单规则禁用
    ruleEnable(params) {
        let url = "/rule/enable";
        return http.post(url, params);
    }
};
