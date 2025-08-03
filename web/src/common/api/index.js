/****   http.js   ****/
import http from "./http";
const api = {
    // 登录
    login(params) {
        let url = "/user/login";
        return http.post(url, params);
    },
    // 注册账号时获取验证码
    getCodeForRegister(params) {
        let url = "/user/register/email/code";
        return http.post(url, params);
    },
    // 注册账号
    register(params) {
        let url = "/user/register";
        return http.post(url, params);
    },
    registerV2(params) {
        let url = "/user/register/v2";
        return http.post(url, params);
    },
    // 重置密码时获取验证码
    getCodeForResetPwd(params) {
        let url = "/user/reset_pwd/email/code";
        return http.post(url, params);
    },
    // 重置密码
    resetPwd(params) {
        let url = "user/reset_pwd";
        return http.post(url, params);
    },
    // 退出登录 {
    logout() {
        let url = "/user/logout";
        return http.get(url);
    },
    // 获取用户信息
    getUserInfo(params) {
        let url = "/user/info";
        return http.get(url, params);
    },
    // 获取全量成员列表
    getUserList(params) {
        let url = "/user/list";
        return http.get(url, params);
    },
    // 获取当前空间成员
    getCurWorkSpaceMemberList(params) {
        let url = '/ws/get_user_list'
        return http.get(url, params)
    },
    // 获取当前流程成员
    getCurProgressMemberList(params) {
        let url = '/tmpl/get_user_list'
        return http.get(url, params)
    },
    // 获取最近访问
    getRecentItem() {
        let url = '/ws/app/get_user_recently_visited_log'
        return http.get(url)
    },
    // html富文本编辑器上传图片
    uploadHtmlImg(params) {
        let url = '/tmpl/app/upload_image_for_html'
        return http.uplodFile(url, params)
    },
    // 获取当前空间内当前流程之外的流程列表
    getOtherTmplList(params) {
        let url = '/tmpl/app/get_tmpl_list_other'
        return http.get(url, params)
    },
    // 获取获取关联流程的表格数据下拉
    getOtherTmplDataList(params) {
        let url = '/tmpl/app/get_list_data_select'
        return http.get(url, params)
    }
};
//导出
export default api;
