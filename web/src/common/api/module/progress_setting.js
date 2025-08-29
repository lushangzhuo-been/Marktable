import http from "../http";
export default {
    // <-- 流程配置及信息 -->
    // 获取左侧流程树
    // getProgressTree(params) {   // 旧接口
    //     // let url = '/tmpl/list'
    //     let url = "/tmpl/app/get_tmpl_list";
    //     return http.get(url, params);
    // },
    getProgressTree(params) { // 新接口
        let url = '/ws/ws_file/list'
        // let url = "/tmpl/app/get_tmpl_list";
        return http.get(url, params);
    },
    // <-- 流程start -->
    getProgressTreeDetail(params) {
        let url = "/tmpl/info";
        return http.get(url, params);
    },
    // 修改流程信息
    updateProgressTreeDetail(params) {
        let url = "/tmpl/update";
        return http.post(url, params);
    },
    // 添加流程
    addProgress(params) {
        let url = "/tmpl/create";
        return http.post(url, params);
    },
    // 删除流程
    deleteProgressTreeDetail(params) {
        let url = "/tmpl/delete";
        return http.post(url, params);
    },
    // <-- 流程 end -->
    // <-- 文件夹  start -->
    addFolder(params) {
        let url = "/ws/ws_file/create";
        return http.post(url, params);
    },
    // 获取文件夹信息
    getFolderDetail(params) {
        let url = "/ws/ws_file/info";
        return http.get(url, params);
    },
    // 修改文件夹信息
    updateFolderDetail(params) {
        let url = "/ws/ws_file/update_ws_file_name";
        return http.post(url, params);
    },
    // 删除文件夹
    deleteFolder(params) {
        let url = "/ws/ws_file/delete";
        return http.post(url, params);
    },
    // <-- 文件夹  end -->
    // 流程树排序(拖动)
    moveProgressTree(params) {
        let url = "ws/ws_file/tmpl_ws_move_file"
        return http.post(url, params);
    },
    // <-- 角色 -->
    // 角色配置
    getRoleConfig() {
        let url = "/tmpl/role/config";
        return http.get(url);
    },
    // 创建角色
    createRole(params) {
        let url = "/tmpl/role/create";
        return http.post(url, params);
    },
    //  修改角色
    updateRole(params) {
        let url = "/tmpl/role/update";
        return http.post(url, params);
    },
    //  删除角色
    delRole(params) {
        let url = "/tmpl/role/delete";
        return http.post(url, params);
    },
    // 获取角色列表
    getRoleList(params) {
        let url = "/tmpl/role/list";
        return http.get(url, params);
    },

    // <-- 成员 -->
    // 获取成员列表
    getRemberList(params) {
        let url = "/tmpl/member/list";
        return http.get(url, params);
    },
    // 添加成员
    addRember(params) {
        let url = "/tmpl/member/create";
        return http.post(url, params);
    },
    // 修改成员
    updateRember(params) {
        let url = "/tmpl/member/update";
        return http.post(url, params);
    },
    // 删除成员
    delRember(params) {
        let url = "/tmpl/member/delete";
        return http.post(url, params);
    },

    // <-- 字段 -->
    // 创建字段
    createField(params) {
        let url = "/tmpl/field/create";
        return http.post(url, params);
    },
    // 字段配置
    getFieldConfig() {
        let url = "/tmpl/field/config";
        return http.get(url);
    },
    // 修改字段
    updateField(params) {
        let url = "/tmpl/field/update";
        return http.post(url, params);
    },
    // 删除字段
    deleteField(params) {
        let url = "/tmpl/field/delete";
        return http.post(url, params);
    },
    getFieldList(params) {
        let url = "/tmpl/field/list";
        return http.get(url, params);
    },
    getFieldInfo(params) {
        let url = "/tmpl/field/info";
        return http.get(url, params);
    },

    // <-- 界面 -->
    // 获取界面配置
    getScreenConfig() {
        let url = "/tmpl/screen/config";
        return http.get(url);
    },
    // 获取界面列表(传入不同的module)
    getScreenList(params) {
        let url = "/tmpl/screen/list";
        return http.get(url, params);
    },
    // 获取剩余界面列表
    getRestScreenList(params) {
        let url = "/tmpl/screen/surplus";
        return http.get(url, params);
    },
    // 添加界面
    addScreen(params) {
        let url = "/tmpl/screen/create";
        return http.post(url, params);
    },
    // 设置是否必填
    setIsRequired(params) {
        let url = "/tmpl/screen/set";
        return http.post(url, params);
    },
    // 删除界面
    delScreen(params) {
        let url = "/tmpl/screen/delete";
        return http.post(url, params);
    },
    // 移动界面字段
    moveScreenField(params) {
        let url = "/tmpl/screen/move";
        return http.post(url, params);
    },
    // 详情页面tab
    //  tab config
    getTabConfig(params) {
        let url = "/tmpl/tab/config";
        return http.get(url, params);
    },
    //  tab list
    getTabList(params) {
        let url = "/tmpl/tab/list";
        return http.get(url, params);
    },
    //  tab create
    tabCreate(params) {
        let url = "/tmpl/tab/create";
        return http.post(url, params);
    },
    //  tab delete
    tabDelete(params) {
        let url = "/tmpl/tab/delete";
        return http.post(url, params);
    },
    // 详情页面子任务
    //  sub config check
    getSubWorkConfig(params) {
        let url = "/tmpl/sub/config/check";
        return http.get(url, params);
    },
    //  sub list
    getSubWorkList(params) {
        let url = "/tmpl/sub/config/list";
        return http.get(url, params);
    },
    //  sub create
    subWorkCreate(params) {
        let url = "/tmpl/sub/config/create";
        return http.post(url, params);
    },
    //  sub delete
    subWorkDelete(params) {
        let url = "/tmpl/sub/config/delete";
        return http.post(url, params);
    },
    //  <-- 节点 start -->
    // 获取节点配置
    getNodeConfig() {
        let url = "/tmpl/node/config";
        return http.get(url);
    },
    // 获取节点列表
    getNodeList(params) {
        let url = "/tmpl/node/list";
        return http.get(url, params);
    },
    // 获取节点详情
    getNodeDetail(params) {
        let url = "/tmpl/node/info";
        return http.get(url, params);
    },
    // 创建节点
    createNode(params) {
        let url = "/tmpl/node/create";
        return http.post(url, params);
    },
    // 更新节点
    updateNode(params) {
        let url = "/tmpl/node/update";
        return http.post(url, params);
    },
    // 删除流程
    deleteNode(params) {
        let url = "/tmpl/node/delete";
        return http.post(url, params);
    },
    // 创建转化按钮
    createBtn(params) {
        let url = "/tmpl/step/create";
        return http.post(url, params);
    },
    // 修改转化按钮
    updateBtn(params) {
        let url = "/tmpl/step/update";
        return http.post(url, params);
    },
    // 删除转化按钮
    deleteBtn(params) {
        let url = "/tmpl/step/delete";
        return http.post(url, params);
    },
    // 获取按钮详情
    getBtnDetail(params) {
        let url = "/tmpl/step/info";
        return http.get(url, params);
    },
    // <-- 节点抽屉 -->
    // 转化界面-获取配置(TS-TransferScreen)
    getTSConfig() {
        let url = "/tmpl/step/screen/config";
        return http.get(url);
    },
    // 转化界面-获取字段列表
    getTSList(params) {
        let url = "/tmpl/step/screen/list";
        return http.get(url, params);
    },
    // 转化界面-获取字段下拉列表
    getTSOptions(params) {
        let url = "/tmpl/step/screen/surplus";
        return http.get(url, params);
    },
    // 转化界面-添加字段
    createTSField(params) {
        let url = "/tmpl/step/screen/create";
        return http.post(url, params);
    },
    // 转化界面-设置是否必填
    setTSRequired(params) {
        let url = "/tmpl/step/screen/set_required";
        return http.post(url, params);
    },
    // 转化界面-删除字段
    delTSField(params) {
        let url = "/tmpl/step/screen/delete";
        return http.post(url, params);
    },
    //转化界面-移动界面字段
    moveTSField(params) {
        let url = "/tmpl/step/screen/move";
        return http.post(url, params);
    },
    // 只读规则角色列表
    getRuleRoleList(params) {
        let url = "/tmpl/field/get_all_person_com";
        return http.get(url, params);
    },
    // 只读规则状态列表
    getRuleStatusList(params) {
        let url = "/tmpl/status/get_all";
        return http.get(url, params);
    },
    // 更新只读规则
    updateRuleRoleList(params) {
        let url = "/tmpl/field/update_read_only_rule";
        return http.post(url, params);
    },
    // 获取详情
    getReadOnlyRule(params) {
        let url = "/tmpl/field/get_read_only_rule";
        return http.get(url, params);
    },
    // 创建限制器
    createLimit(params) {
        let url = "/tmpl/step/limit/create";
        return http.post(url, params);
    },
    // 查询限制器列表
    getLimitList(params) {
        let url = "/tmpl/step/limit/list";
        return http.get(url, params);
    },
    // 查询单限制器详情
    getLimitInfo(params) {
        let url = "/tmpl/step/limit/info";
        return http.get(url, params);
    },
    // 编辑限制器
    updateLimitInfo(params) {
        let url = "/tmpl/step/limit/update";
        return http.post(url, params);
    },
    // 删除限制器
    deleteLimitInfo(params) {
        let url = "/tmpl/step/limit/delete";
        return http.post(url, params);
    },
    //  <-- 节点 end -->
    //  <-- 新节点(流程) start -->
    // 获取流程表格
    getOverviewTable(params) {
        let url = "/tmpl/status/overall_view";
        return http.get(url, params);
    },
    // 移动流程表格行
    moveOverviewTable(params) {
        let url = "/tmpl/status/move";
        return http.post(url, params);
    },
    // 新建工作状态
    createWorkStatus(params) {
        let url = "/tmpl/status/create";
        return http.post(url, params);
    },
    // 更改状态名称
    editWorkName(params) {
        let url = "/tmpl/status/rename";
        return http.post(url, params);
    },
    // 设置为初始状态
    setFirst(params) {
        let url = "/tmpl/status/set_first"
        return http.post(url, params);
    },
    // 删除状态
    deleteStatus(params) {
        let url = "/tmpl/status/delete"
        return http.post(url, params);
    },
    // 新建步骤
    addSteps(params) {
        let url = "tmpl/step/create"
        return http.post(url, params);
    },
    // 删除步骤
    deleteSteps(params) {
        let url = "tmpl/step/delete"
        return http.post(url, params);
    },
    //  <-- 新节点(流程) end -->
    // 获取权限配置
    fetchAuthConfig(params) {
        let url = "tmpl/auth/config"
        return http.get(url, params);
    },
    // 获取当前权限的详情
    fetchAuthDetail(params) {
        let url = "tmpl/auth/Info"
        return http.get(url, params);
    },
    // 保存选中的权限角色等
    updateSelectedAuth(params) {
        let url = "tmpl/auth/update"
        return http.post(url, params);
    }
};
