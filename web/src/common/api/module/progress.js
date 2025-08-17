import http from "../http";
export default {
    // 成辉流程相关接口
    // 获取流程视图，col配置。
    getProgressColConfig(params) {
        let url = "/tmpl/app/get_screen";
        return http.get(url, params);
    },
    // 流程列表返回
    getProgressListData(params) {
        let url = "/tmpl/app/get_list_data";
        return http.get(url, params);
    },
    // 流程导出
    exportProgress(params) {
        let url = "/tmpl/app/get_list_data";
        return http.downloadFile(url, params);
    },
    // 获取人员列表
    getPersonList(params) {
        let url = "/ws/get_user_list";
        return http.get(url, params);
    },
    // 获取推荐人员列表
    getPeopleList(params) {
        let url = "/ws/app/get_ws_user_list";
        return http.get(url, params);
    },
    // 创建流程项
    createProgressItem(params) {
        let url = "/tmpl/app/create_action";
        return http.post(url, params);
    },
    // 获取单个流程详情数据
    getFormData(params) {
        let url = "/tmpl/app/get_data";
        return http.get(url, params);
    },
    // 获取状态下拉
    getStatusMenusList(params) {
        let url = "/tmpl/app/get_step_list";
        return http.get(url, params);
    },
    // 获取全部状态下拉
    getAllStatusList(params) {
        let url = "/tmpl/app/get_status_list";
        return http.get(url, params);
    },
    // 获取流程下一个节点操作，是否需要填表
    getBtnScreen(params) {
        let url = "/tmpl/app/get_step_screen";
        return http.get(url, params);
    },
    // 操作流程进入下一个节点
    switchStepAction(params) {
        let url = "/tmpl/app/switch_step_action";
        return http.post(url, params);
    },
    // 编辑流程信息
    updateProgress(params) {
        let url = "/tmpl/app/update_action";
        return http.post(url, params);
    },
    // 获取进展详情
    getEvolveDetail(params) {
        let url = "/tmpl/app/get_list_progress";
        return http.get(url, params);
    },
    // 更新进展详情
    addEvolve(params) {
        let url = "/tmpl/app/add_progress";
        return http.post(url, params);
    },
    // 编辑进展
    updateEvolve(params) {
        let url = "/tmpl/app/update_progress";
        return http.post(url, params);
    },
    // 删除进展
    deleteEvolve(params) {
        let url = "/tmpl/app/delete_progress";
        return http.post(url, params);
    },
    // // 获取显隐配置
    // getColHideShow(params) {
    //     let url = "/tmpl/app/get_list_module_screen_config";
    //     return http.get(url, params);
    // },
    // // 配置显隐
    // setColHideShow(params) {
    //     let url = "/tmpl/app/set_list_module_screen_config";
    //     return http.post(url, params);
    // },
    // 获取过滤配置
    getFilterConfig(params) {
        // let url = "tmpl/app/get_filter_config";
        let url = "tmpl/app/get_config";
        return http.get(url, params);
    },
    // 获取视图
    getView(params) {
        let url = "tmpl/app/get_user_views";
        return http.get(url, params);
    },
    // 查询单个视图
    getViewInfo(params) {
        let url = "tmpl/app/get_view_info";
        return http.get(url, params);
    },
    // 视图重命名
    viewRename(params) {
        let url = "tmpl/app/rename_view";
        return http.post(url, params);
    },
    // 固定视图
    pinView(params) {
        let url = "tmpl/app/pin_view";
        return http.post(url, params);
    },
    // 解除固定试图
    unPinView(params) {
        let url = "tmpl/app/unpin_view";
        return http.post(url, params);
    },
    // 排序视图
    coordinate(params) {
        let url = "tmpl/app/set_view_coordinate";
        return http.post(url, params);
    },
    // 新增视图
    createView(params) {
        let url = "tmpl/app/create_view";
        return http.post(url, params);
    },
    // 更新视图
    updateViewFilter(params) {
        let url = "tmpl/app/update_view_filter";
        return http.post(url, params);
    },
    // 更新排序
    updateSortOrder(params) {
        let url = "tmpl/app/update_view_sort_order";
        return http.post(url, params);
    },
    // 卡片更新分组
    updateViewCardGroup(params) {
        let url = "tmpl/app/update_view_card_group";
        return http.post(url, params);
    },
    // 更新列表显示隐藏
    updateViewColumns(params) {
        let url = "tmpl/app/update_view_columns";
        return http.post(url, params);
    },
    // 删除视图
    deleteView(params) {
        let url = "tmpl/app/delete_view";
        return http.post(url, params);
    },
    // 获取流程信息
    getProgressInfo(params) {
        let url = "tmpl/app/get_tmpl_info";
        return http.get(url, params);
    },
    // 详情日志
    getProgressLog(params) {
        let url = "tmpl/app/get_list_log";
        return http.get(url, params);
    },
    // 删除流程项
    deleteProgressItems(params) {
        let url = "tmpl/app/delete_action";
        return http.post(url, params);
    },
    // 获取附件权限接口
    getFileAuth(params) {
        let url = "tmpl/app/get_file_right";
        return http.get(url, params);
    },
    // 获取附件列表
    getFileList(params) {
        let url = "tmpl/app/get_file_list";
        return http.get(url, params);
    },
    // 附件上传
    uploadFile(params, uploadChange = "") {
        let url = "tmpl/app/upload_file";
        return http.uplodFile(url, params, uploadChange);
    },
    // 附件下载
    downloadFile(params) {
        let url = "tmpl/app/download_file";
        return http.downloadFile(url, params);
    },
    downloadTypeCheck(params) {
        let url = "tmpl/app/get_upload_ext_config";
        return http.get(url, params);
    },
    // 附件删除
    deleteFile(params) {
        let url = "tmpl/app/delete_file";
        return http.post(url, params);
    },
    // 预览文件弹窗-将历史文件设置为当前文件
    setCurrentFile(params) {
        let url = "tmpl/app/update_file_is_current_version";
        return http.post(url, params);
    },
    // 预览文件弹窗-选中预览文件
    previewFile(params) {
        let url = "tmpl/app/download_file";
        return http.previewFile(url, params);
    },
    // 消息通知
    // 最新
    getMessageList(params) {
        let url = "message/get_message_list";
        return http.get(url, params);
    },
    // 全部标记已读
    markAllReadRead(params) {
        let url = "message/read_all_message";
        return http.post(url, params);
    },
    // 确认已读
    hadReadMessage(params) {
        let url = "message/read_message";
        return http.post(url, params);
    },
    // 获取未读得条数
    getUnReadNum(params) {
        let url = "message/get_un_read_count";
        return http.get(url, params);
    },
    // 用户侧获取查看、编辑、导出、进展、新增等权限
    getUserAuth(params) {
        let url = "tmpl/app/get_user_auth";
        return http.get(url, params);
    },
    // 卡片看板分组列表
    getCardGroupByList(params) {
        let url = "tmpl/field/enumeration";
        return http.get(url, params);
    },
    // 卡片全状态及可切换配置
    getAllStatusConfig(params) {
        let url = "tmpl/status/next";
        return http.get(url, params);
    },
    // 创建关联子任务
    createSubAction(params) {
        let url = "tmpl/app/create_sub_action ";
        return http.post(url, params);
    },
    // 删除关联子任务
    deleteSubAction(params) {
        let url = "tmpl/app/delete_sub_action ";
        return http.post(url, params);
    },
    // 子任务数量查询
    getSubTaskCount(params) {
        let url = "tmpl/app/get_sub_list_count";
        return http.get(url, params);
    }
};
