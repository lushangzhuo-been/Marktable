package message

const (
	ReadSignYes string = "yes"
	ReadSignNo  string = "no"
)

const (
	ModeReassign string = "reassign_mode"
	ModeAtPerson string = "at_person_mode"
	ModeKickOut  string = "kick_out_mode"
)

// TemplateForReassign 重新指派的消息模版
type TemplateForReassign struct {
	IssueTitle string `json:"issue_title"`
	IssueURL   string `json:"issue_url"`
	WsName     string `json:"ws_name"`
	TmplName   string `json:"tmpl_name"`
}

// TemplateForAtPerson @某人的消息模版
type TemplateForAtPerson struct {
	ProgressContent string `json:"progress_content"`
	IssueTitle      string `json:"issue_title"`
	IssueURL        string `json:"issue_url"`
	WsName          string `json:"ws_name"`
	TmplName        string `json:"tmpl_name"`
}

// TemplateForKickOut 将用户踢出空间或者流程的消息模版
type TemplateForKickOut struct {
}
