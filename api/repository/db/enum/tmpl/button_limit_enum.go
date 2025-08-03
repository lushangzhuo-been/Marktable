package tmpl

const (
	ButtonLimitModePermission string = "permission"
	ButtonLimitModeSub        string = "sub"
)

var ButtonLimitModeMap = map[string]string{
	ButtonLimitModePermission: "权限验证",
	ButtonLimitModeSub:        "子任务阻止",
}
