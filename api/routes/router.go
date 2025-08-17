package routes

import (
	"mark3/api/v1/app"
	dashboard "mark3/api/v1/app/dashboard"
	appTmpl "mark3/api/v1/app/tmpl"
	ws2 "mark3/api/v1/app/ws"
	"mark3/api/v1/message"
	"mark3/api/v1/rule"
	"mark3/api/v1/schedule"
	"mark3/api/v1/tmpl"
	"mark3/api/v1/user"
	"mark3/api/v1/ws"
	"mark3/api/v1/ws_file"
	"mark3/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.CorsMiddleware())

	//加载静态资源
	r.Static("/static", "./views")

	/**
	 * 用户模块接口
	 */
	userApi := new(user.UserApi)
	r.POST("/user/register", userApi.Register)
	r.POST("/user/login", userApi.Login)
	r.POST("/user/register/email/code", userApi.RegisterEmailCode)
	r.POST("/user/register/v2", userApi.RegisterV2)
	r.POST("/user/reset_pwd/email/code", userApi.ResetPwdEmailCode)
	r.POST("/user/reset_pwd", userApi.ResetPwd)
	r.GET("/user/info", middleware.AuthMiddleware(), userApi.Info)
	r.POST("/user/upload_avatar", middleware.AuthMiddleware(), userApi.UploadAvatar)
	r.POST("/user/update", middleware.AuthMiddleware(), userApi.Update)
	r.POST("/user/update_pwd", middleware.AuthMiddleware(), userApi.UpdatePwd)
	r.POST("/user/update_email", middleware.AuthMiddleware(), userApi.UpdateEmail)

	/**
	 * 消息模块接口
	 */
	messageApi := new(message.MessageApi)
	r.GET("/message/get_un_read_count", middleware.AuthMiddleware(), messageApi.GetUnReadCount)
	r.GET("/message/get_message_list", middleware.AuthMiddleware(), messageApi.GetMessageList)
	r.POST("/message/read_message", middleware.AuthMiddleware(), messageApi.ReadMessage)
	r.POST("/message/read_all_message", middleware.AuthMiddleware(), messageApi.ReadAllMessage)

	/**
	 * 空间模块接口
	 */
	wsApi := new(ws.WsApi)
	r.POST("/ws/create", middleware.AuthMiddleware(), wsApi.Create)
	r.POST("/ws/update", middleware.AuthMiddleware(), wsApi.Update)
	r.POST("/ws/delete", middleware.AuthMiddleware(), wsApi.Delete)
	r.GET("/ws/info", middleware.AuthMiddleware(), wsApi.Info)

	wsCommonApi := new(ws.CommonApi)
	r.GET("/ws/get_user_list", middleware.AuthMiddleware(), wsCommonApi.GetUserList)

	wsMemberApi := new(ws.MemberApi)
	r.POST("/ws/member/create", middleware.AuthMiddleware(), wsMemberApi.Create)
	r.POST("/ws/member/update", middleware.AuthMiddleware(), wsMemberApi.Update)
	r.POST("/ws/member/delete", middleware.AuthMiddleware(), wsMemberApi.Delete)
	r.GET("/ws/member/list", middleware.AuthMiddleware(), wsMemberApi.List)
	r.GET("/ws/member/config", middleware.AuthMiddleware(), wsMemberApi.Config)

	/**
	文件夹接口
	*/
	WsFileApi := new(ws_file.WsFileApi)
	r.POST("/ws/ws_file/create", middleware.AuthMiddleware(), WsFileApi.Create)
	r.GET("/ws/ws_file/info", middleware.AuthMiddleware(), WsFileApi.Info)
	r.POST("/ws/ws_file/update_ws_file_name", middleware.AuthMiddleware(), WsFileApi.UpdateWsFileName)
	r.POST("/ws/ws_file/update_ws_fileCoordinate", middleware.AuthMiddleware(), WsFileApi.UpdateWsFileCoordinate)
	r.POST("/ws/ws_file/delete", middleware.AuthMiddleware(), WsFileApi.Delete)
	r.POST("/ws/ws_file/tmpl_ws_move_file", middleware.AuthMiddleware(), WsFileApi.TmplWsFileMove)
	r.GET("/ws/ws_file/list", middleware.AuthMiddleware(), WsFileApi.List)

	/**
	 * 模板模块接口
	 */
	tmplApi := new(tmpl.TmplApi)
	r.GET("/tmpl/config", middleware.AuthMiddleware(), tmplApi.Config)
	r.GET("/tmpl/list", middleware.AuthMiddleware(), tmplApi.List)
	r.GET("/tmpl/info", middleware.AuthMiddleware(), tmplApi.Info)
	r.POST("/tmpl/create", middleware.AuthMiddleware(), tmplApi.Create)
	r.POST("/tmpl/update", middleware.AuthMiddleware(), tmplApi.Update)
	r.POST("/tmpl/delete", middleware.AuthMiddleware(), tmplApi.Delete)

	tmplRoleApi := new(tmpl.RoleApi)
	r.GET("/tmpl/role/config", middleware.AuthMiddleware(), tmplRoleApi.Config)
	r.POST("/tmpl/role/create", middleware.AuthMiddleware(), tmplRoleApi.Create)
	r.POST("/tmpl/role/update", middleware.AuthMiddleware(), tmplRoleApi.Update)
	r.POST("/tmpl/role/delete", middleware.AuthMiddleware(), tmplRoleApi.Delete)
	r.GET("/tmpl/role/list", middleware.AuthMiddleware(), tmplRoleApi.List)

	tmplAuthApi := new(tmpl.AuthApi)
	r.GET("/tmpl/auth/config", middleware.AuthMiddleware(), tmplAuthApi.Config)
	r.POST("/tmpl/auth/update", middleware.AuthMiddleware(), tmplAuthApi.Update)
	r.GET("/tmpl/auth/Info", middleware.AuthMiddleware(), tmplAuthApi.Detail)

	tmplCommonApi := new(tmpl.CommonApi)
	r.GET("/tmpl/get_user_list", middleware.AuthMiddleware(), tmplCommonApi.GetUserList)

	tmplMemberApi := new(tmpl.MemberApi)
	r.POST("/tmpl/member/create", middleware.AuthMiddleware(), tmplMemberApi.Create)
	r.POST("/tmpl/member/update", middleware.AuthMiddleware(), tmplMemberApi.Update)
	r.GET("/tmpl/member/list", middleware.AuthMiddleware(), tmplMemberApi.List)
	r.POST("/tmpl/member/delete", middleware.AuthMiddleware(), tmplMemberApi.Delete)

	tmplFieldApi := new(tmpl.FieldApi)
	r.GET("/tmpl/field/config", middleware.AuthMiddleware(), tmplFieldApi.Config)
	r.GET("/tmpl/field/list", middleware.AuthMiddleware(), tmplFieldApi.List)
	r.GET("/tmpl/field/info", middleware.AuthMiddleware(), tmplFieldApi.Info)
	r.POST("/tmpl/field/create", middleware.AuthMiddleware(), tmplFieldApi.Create)
	r.POST("/tmpl/field/update", middleware.AuthMiddleware(), tmplFieldApi.Update)
	r.GET("/tmpl/field/get_all_person_com", middleware.AuthMiddleware(), tmplFieldApi.GetAllPersonCom)
	r.GET("/tmpl/field/get_read_only_rule", middleware.AuthMiddleware(), tmplFieldApi.GetReadOnlyRule)
	r.POST("/tmpl/field/update_read_only_rule", middleware.AuthMiddleware(), tmplFieldApi.UpdateReadOnlyRule)
	r.POST("/tmpl/field/delete", middleware.AuthMiddleware(), tmplFieldApi.Delete)
	r.GET("/tmpl/field/enumeration", middleware.AuthMiddleware(), tmplFieldApi.Enumeration)

	tmplScreenApi := new(tmpl.ScreenApi)
	r.GET("/tmpl/screen/config", middleware.AuthMiddleware(), tmplScreenApi.Config)
	r.GET("/tmpl/screen/list", middleware.AuthMiddleware(), tmplScreenApi.List)
	r.POST("/tmpl/screen/create", middleware.AuthMiddleware(), tmplScreenApi.Create)
	r.POST("/tmpl/screen/set", middleware.AuthMiddleware(), tmplScreenApi.Set)
	r.POST("/tmpl/screen/move", middleware.AuthMiddleware(), tmplScreenApi.Move)
	r.POST("/tmpl/screen/delete", middleware.AuthMiddleware(), tmplScreenApi.Delete)
	r.GET("/tmpl/screen/surplus", middleware.AuthMiddleware(), tmplScreenApi.Surplus)

	tmplNodeApi := new(tmpl.NodeApi)
	r.GET("/tmpl/node/config", middleware.AuthMiddleware(), tmplNodeApi.Config)
	r.GET("/tmpl/node/list", middleware.AuthMiddleware(), tmplNodeApi.List)
	r.GET("/tmpl/node/info", middleware.AuthMiddleware(), tmplNodeApi.Info)
	r.POST("/tmpl/node/create", middleware.AuthMiddleware(), tmplNodeApi.Create)
	r.POST("/tmpl/node/update", middleware.AuthMiddleware(), tmplNodeApi.Update)
	r.GET("/tmpl/node/get_all_status", middleware.AuthMiddleware(), tmplNodeApi.GetAllStatus)
	r.POST("/tmpl/node/delete", middleware.AuthMiddleware(), tmplNodeApi.Delete)

	tmplButtonApi := new(tmpl.ButtonApi)
	r.POST("/tmpl/button/create", middleware.AuthMiddleware(), tmplButtonApi.Create)
	r.POST("/tmpl/button/update", middleware.AuthMiddleware(), tmplButtonApi.Update)
	r.POST("/tmpl/button/delete", middleware.AuthMiddleware(), tmplButtonApi.Delete)
	r.GET("/tmpl/button/info", middleware.AuthMiddleware(), tmplButtonApi.Info)

	tmplButtonScreenApi := new(tmpl.ButtonScreenApi)
	r.GET("/tmpl/button/screen/config", middleware.AuthMiddleware(), tmplButtonScreenApi.Config)
	r.GET("/tmpl/button/screen/list", middleware.AuthMiddleware(), tmplButtonScreenApi.List)
	r.POST("/tmpl/button/screen/create", middleware.AuthMiddleware(), tmplButtonScreenApi.Create)
	r.POST("/tmpl/button/screen/set_required", middleware.AuthMiddleware(), tmplButtonScreenApi.SetRequired)
	r.POST("/tmpl/button/screen/delete", middleware.AuthMiddleware(), tmplButtonScreenApi.Delete)
	r.POST("/tmpl/button/screen/move", middleware.AuthMiddleware(), tmplButtonScreenApi.Move)
	r.GET("/tmpl/button/screen/surplus", middleware.AuthMiddleware(), tmplButtonScreenApi.Surplus)

	tmplButtonLimitApi := new(tmpl.ButtonLimitApi)
	r.GET("/tmpl/button/limit/info", middleware.AuthMiddleware(), tmplButtonLimitApi.Info)
	r.GET("/tmpl/button/limit/list", middleware.AuthMiddleware(), tmplButtonLimitApi.List)
	r.POST("/tmpl/button/limit/create", middleware.AuthMiddleware(), tmplButtonLimitApi.Create)
	r.POST("/tmpl/button/limit/update", middleware.AuthMiddleware(), tmplButtonLimitApi.Update)
	r.POST("/tmpl/button/limit/delete", middleware.AuthMiddleware(), tmplButtonLimitApi.Delete)

	tmplStatusApi := new(tmpl.StatusApi)
	r.GET("/tmpl/status/overall_view", middleware.AuthMiddleware(), tmplStatusApi.OverallView)
	r.GET("/tmpl/status/get_all", middleware.AuthMiddleware(), tmplStatusApi.GetAll)
	r.POST("/tmpl/status/create", middleware.AuthMiddleware(), tmplStatusApi.Create)
	r.POST("/tmpl/status/rename", middleware.AuthMiddleware(), tmplStatusApi.Rename)
	r.POST("/tmpl/status/set_first", middleware.AuthMiddleware(), tmplStatusApi.SetFirst)
	r.POST("/tmpl/status/move", middleware.AuthMiddleware(), tmplStatusApi.Move)
	r.POST("/tmpl/status/delete", middleware.AuthMiddleware(), tmplStatusApi.Delete)
	r.GET("/tmpl/status/next", middleware.AuthMiddleware(), tmplStatusApi.StatusNext)

	tmplStepApi := new(tmpl.StepApi)
	r.POST("/tmpl/step/create", middleware.AuthMiddleware(), tmplStepApi.Create)
	r.POST("/tmpl/step/delete", middleware.AuthMiddleware(), tmplStepApi.Delete)
	r.GET("/tmpl/step/info", middleware.AuthMiddleware(), tmplStepApi.Info)

	tmplStepScreenApi := new(tmpl.StepScreenApi)
	r.GET("/tmpl/step/screen/config", middleware.AuthMiddleware(), tmplStepScreenApi.Config)
	r.GET("/tmpl/step/screen/list", middleware.AuthMiddleware(), tmplStepScreenApi.List)
	r.POST("/tmpl/step/screen/create", middleware.AuthMiddleware(), tmplStepScreenApi.Create)
	r.POST("/tmpl/step/screen/set_required", middleware.AuthMiddleware(), tmplStepScreenApi.SetRequired)
	r.POST("/tmpl/step/screen/delete", middleware.AuthMiddleware(), tmplStepScreenApi.Delete)
	r.POST("/tmpl/step/screen/move", middleware.AuthMiddleware(), tmplStepScreenApi.Move)
	r.GET("/tmpl/step/screen/surplus", middleware.AuthMiddleware(), tmplStepScreenApi.Surplus)

	tmplStepLimiterApi := new(tmpl.StepLimiterApi)
	r.GET("/tmpl/step/limit/info", middleware.AuthMiddleware(), tmplStepLimiterApi.Info)
	r.GET("/tmpl/step/limit/list", middleware.AuthMiddleware(), tmplStepLimiterApi.List)
	r.POST("/tmpl/step/limit/create", middleware.AuthMiddleware(), tmplStepLimiterApi.Create)
	r.POST("/tmpl/step/limit/update", middleware.AuthMiddleware(), tmplStepLimiterApi.Update)
	r.POST("/tmpl/step/limit/delete", middleware.AuthMiddleware(), tmplStepLimiterApi.Delete)

	//模块标签
	tmplTabApi := new(tmpl.TabApi)
	r.GET("/tmpl/tab/config", middleware.AuthMiddleware(), tmplTabApi.TabConfig)
	r.GET("/tmpl/tab/list", middleware.AuthMiddleware(), tmplTabApi.TmplTabList)
	r.POST("/tmpl/tab/create", middleware.AuthMiddleware(), tmplTabApi.TmplTabCreate)
	r.POST("/tmpl/tab/delete", middleware.AuthMiddleware(), tmplTabApi.TmplTabDelete)

	//子任务配置
	tmplSubConfigApi := new(tmpl.TmplSubConfigApi)
	r.GET("/tmpl/sub/config/check", middleware.AuthMiddleware(), tmplSubConfigApi.TmplSubConfigCheck)
	r.GET("/tmpl/sub/config/list", middleware.AuthMiddleware(), tmplSubConfigApi.TmplSubConfigList)
	r.POST("/tmpl/sub/config/create", middleware.AuthMiddleware(), tmplSubConfigApi.TmplSubConfigCreate)
	r.POST("/tmpl/sub/config/delete", middleware.AuthMiddleware(), tmplSubConfigApi.TmplSubConfigDelete)

	/*
	 * 以下都是用户应用侧接口
	 */
	wsAppApi := new(ws2.WsApi)
	r.GET("/ws/app/get_my_ws_list", middleware.AuthMiddleware(), wsAppApi.GetMyWsList)
	r.GET("/ws/app/get_ws_info", middleware.AuthMiddleware(), wsAppApi.GetWsInfo)

	commonAppApi := new(app.CommonApi)
	//获取用户最近访问的模板
	r.GET("/ws/app/get_user_recently_visited_log", middleware.AuthMiddleware(), commonAppApi.GetUserRecentlyVisitedLog)
	//获取空间下用户列表，为用户组件提供接口数据
	r.GET("/ws/app/get_ws_user_list", middleware.AuthMiddleware(), commonAppApi.GetWsUserList)

	tmplAppApi := new(appTmpl.TmplApi)
	r.GET("/tmpl/app/get_tmpl_list", middleware.AuthMiddleware(), tmplAppApi.GetTmplList)
	r.GET("/tmpl/app/get_tmpl_list_other", middleware.AuthMiddleware(), tmplAppApi.GetTmplOtherList)
	r.GET("/tmpl/app/get_tmpl_info", middleware.AuthMiddleware(), tmplAppApi.GetTmplInfo)
	r.GET("/tmpl/app/get_config", middleware.AuthMiddleware(), tmplAppApi.GetConfig)
	r.GET("/tmpl/app/get_status_list", middleware.AuthMiddleware(), tmplAppApi.GetStatusList)
	r.GET("/tmpl/app/get_screen", middleware.AuthMiddleware(), tmplAppApi.GetScreen)
	r.POST("/tmpl/app/create_action", middleware.AuthMiddleware(), tmplAppApi.CreateAction)
	r.POST("/tmpl/app/create_sub_action", middleware.AuthMiddleware(), tmplAppApi.CreateSubAction)
	r.GET("/tmpl/app/get_user_auth", middleware.AuthMiddleware(), tmplAppApi.GetUserAuth)
	r.GET("/tmpl/app/get_list_data", middleware.AuthMiddleware(), tmplAppApi.GetListData)
	r.GET("/tmpl/app/get_list_data_select", middleware.AuthMiddleware(), tmplAppApi.GetListDataSelect)
	r.GET("/tmpl/app/get_data", middleware.AuthMiddleware(), tmplAppApi.GetData)
	r.GET("/tmpl/app/get_file_right", middleware.AuthMiddleware(), tmplAppApi.GetFileRight)
	r.POST("/tmpl/app/update_action", middleware.AuthMiddleware(), tmplAppApi.UpdateAction)
	r.POST("/tmpl/app/delete_action", middleware.AuthMiddleware(), tmplAppApi.DeleteAction)
	r.POST("/tmpl/app/delete_sub_action", middleware.AuthMiddleware(), tmplAppApi.DeleteSubAction)
	r.GET("/tmpl/app/get_step_list", middleware.AuthMiddleware(), tmplAppApi.GetStepList)
	r.GET("/tmpl/app/get_step_screen", middleware.AuthMiddleware(), tmplAppApi.GetStepScreen)
	r.POST("/tmpl/app/switch_step_action", middleware.AuthMiddleware(), tmplAppApi.SwitchStepAction)
	r.POST("/tmpl/app/add_progress", middleware.AuthMiddleware(), tmplAppApi.AddProgress)
	r.POST("/tmpl/app/update_progress", middleware.AuthMiddleware(), tmplAppApi.UpdateProgress)
	r.POST("/tmpl/app/delete_progress", middleware.AuthMiddleware(), tmplAppApi.DeleteProgress)
	r.GET("/tmpl/app/get_list_progress", middleware.AuthMiddleware(), tmplAppApi.GetProgressList)
	r.GET("/tmpl/app/get_list_log", middleware.AuthMiddleware(), tmplAppApi.GetLogList)
	r.GET("/tmpl/app/get_sub_list_count", middleware.AuthMiddleware(), tmplAppApi.GetSubListCount)

	fileAppApi := new(appTmpl.UploadApi)
	r.GET("/tmpl/app/get_upload_ext_config", middleware.AuthMiddleware(), fileAppApi.GetUploadExtConfig)
	r.GET("/tmpl/app/get_file_list", middleware.AuthMiddleware(), fileAppApi.GetFileList)
	r.POST("/tmpl/app/upload_file", middleware.AuthMiddleware(), fileAppApi.UploadFile)
	r.GET("/tmpl/app/download_file", middleware.AuthMiddleware(), fileAppApi.DownloadFile)
	r.POST("/tmpl/app/delete_file", middleware.AuthMiddleware(), fileAppApi.DeleteFile)
	r.POST("/tmpl/app/update_file_is_current_version", middleware.AuthMiddleware(), fileAppApi.UpdateFileIsCurrentVersion)
	r.POST("/tmpl/app/upload_image_for_html", middleware.AuthMiddleware(), fileAppApi.UploadImageForHtml)

	viewAppApi := new(appTmpl.ViewApi)
	r.GET("/tmpl/app/get_user_views", middleware.AuthMiddleware(), viewAppApi.GetUserViews)
	r.GET("/tmpl/app/get_view_info", middleware.AuthMiddleware(), viewAppApi.GetViewInfo)
	r.POST("/tmpl/app/create_view", middleware.AuthMiddleware(), viewAppApi.CreateView)
	r.POST("/tmpl/app/rename_view", middleware.AuthMiddleware(), viewAppApi.RenameView)
	r.POST("/tmpl/app/update_view_filter", middleware.AuthMiddleware(), viewAppApi.UpdateViewFilter)
	r.POST("/tmpl/app/update_view_sort_order", middleware.AuthMiddleware(), viewAppApi.UpdateViewSortOrder)
	r.POST("/tmpl/app/update_view_columns", middleware.AuthMiddleware(), viewAppApi.UpdateViewColumns)
	r.POST("/tmpl/app/update_view_card_group", middleware.AuthMiddleware(), viewAppApi.UpdateViewCardGroup)
	r.POST("/tmpl/app/delete_view", middleware.AuthMiddleware(), viewAppApi.DeleteView)
	r.POST("/tmpl/app/pin_view", middleware.AuthMiddleware(), viewAppApi.PinView)
	r.POST("/tmpl/app/unpin_view", middleware.AuthMiddleware(), viewAppApi.UnPinView)
	r.POST("/tmpl/app/set_view_coordinate", middleware.AuthMiddleware(), viewAppApi.SetViewCoordinate)

	//定时任务
	taskApi := new(schedule.TaskApi)
	r.GET("/task/status", taskApi.GetSchedule)
	r.GET("/task/status/set", taskApi.SetSchedule)
	r.GET("/task/status/update", taskApi.AddOrUpdateSchedule)
	r.GET("/task/rule/do", taskApi.DoRuleCron)

	//规则
	ruleApi := new(rule.RuleApi)
	r.POST("/rule/create", middleware.AuthMiddleware(), ruleApi.Create)
	r.POST("/rule/delete", middleware.AuthMiddleware(), ruleApi.Delete)
	r.POST("/rule/update", middleware.AuthMiddleware(), ruleApi.Update)
	r.GET("/rule/detail", middleware.AuthMiddleware(), ruleApi.Detail)
	r.GET("/rule/page", middleware.AuthMiddleware(), ruleApi.Page)
	r.POST("/rule/enable", middleware.AuthMiddleware(), ruleApi.Enable)
	r.GET("/rule/rule_log", middleware.AuthMiddleware(), ruleApi.RuleLog)
	r.GET("/rule/action_log", middleware.AuthMiddleware(), ruleApi.ActionLog)
	r.GET("/rule/config", middleware.AuthMiddleware(), ruleApi.Config)

	//dashboard 接口
	dashBoardStepApi := new(dashboard.DashBoardApi)
	r.GET("/board/config", middleware.AuthMiddleware(), dashBoardStepApi.GetBoardConfig)
	r.POST("/board/add", middleware.AuthMiddleware(), dashBoardStepApi.AddBoard)
	r.POST("/board/modify", middleware.AuthMiddleware(), dashBoardStepApi.ModifyBoard)
	r.POST("/board/delete", middleware.AuthMiddleware(), dashBoardStepApi.DeleteBoard)
	r.POST("/board/copy", middleware.AuthMiddleware(), dashBoardStepApi.CopyBoard)
	r.GET("/board/detail", middleware.AuthMiddleware(), dashBoardStepApi.GetBoardByID)
	r.GET("/board/list", middleware.AuthMiddleware(), dashBoardStepApi.GetBoardList)
	r.POST("/board/statistics", middleware.AuthMiddleware(), dashBoardStepApi.GetBoardStatistics)
	r.POST("/board/statistics/preview", middleware.AuthMiddleware(), dashBoardStepApi.GetBoardStatisticsPreview)
	r.POST("/board/modify/location", middleware.AuthMiddleware(), dashBoardStepApi.ModifyLocation)

	return r

}
