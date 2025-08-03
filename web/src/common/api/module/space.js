import http from '../http'
export default {
    // 获取空间主页信息
    getCurWorkspaceInfo(params) {
        let url = '/ws/app/get_ws_info'
        return http.get(url, params)
    },
    // 获取空间列表
    getWorkspaceList(params) {
        // let url = '/ws/my_list'
        let url = '/ws/app/get_my_ws_list'
        return http.get(url, params)
    },
    // 创建空间
    creatWorkspace(params) {
        let url = '/ws/create'
        return http.post(url, params)
    },
    // 修改空间
    updateWorkspace(params) {
        let url = '/ws/update'
        return http.post(url, params)
    },
    // 删除空间
    deleteWorkspace(params) {
        let url = '/ws/delete'
        return http.post(url, params)
    },
    // 获取空间信息
    getWorkspaceInfo(params) {
        let url = '/ws/info'
        return http.get(url, params)
    },
    // 获取空间成员配置
    getWorkspaceMemberConfig(params) {
        let url = '/ws/member/config'
        return http.get(url, params)
    },
    // 获取空间成员列表
    getWorkspaceMemberList(params) {
        let url = '/ws/member/list'
        return http.get(url, params)
    },
    // 新增空间成员
    addWorkspaceMember(params) {
        let url = '/ws/member/create'
        return http.post(url, params)
    },
    // 编辑空间成员
    updateWorkspaceMember(params) {
        let url = '/ws/member/update'
        return http.post(url, params)
    },
    // 删除空间成员
    delWorkspaceMember(params) {
        let url = '/ws/member/delete'
        return http.post(url, params)
    },
}