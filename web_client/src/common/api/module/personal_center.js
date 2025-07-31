import http from '../http'
export default {

    // 更新用户密码
    updateUserPwd(params) {
        let url = '/user/update_pwd'
        return http.post(url, params)
    },
    // 更新用户信息
    updateUserInfo(params) {
        let url = '/user/update'
        return http.post(url, params)
    },
    // 更新用户头像
    updateUserAvatar(params) {
        let url = '/user/upload_avatar'
        return http.uplodFile(url, params)
    },
    // 重置邮箱验证码
    resetEmailForCode(params) {
        let url = '/user/register/email/code'
        return http.post(url, params)
    },
    // 更新邮箱
    updateEmail(params) {
        let url = "/user/update_email";
        return http.post(url, params);

    }
}