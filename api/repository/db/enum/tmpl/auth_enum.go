package tmpl

const (
	AuthSee      = "see"
	AuthExport   = "export"
	AuthCreate   = "create"
	AuthEdit     = "edit"
	AuthDelete   = "delete"
	AuthProgress = "progress"
	AuthFile     = "file"
	AuthSubTmpl  = "sub_tmpl"
)

var AuthMap = map[string]string{
	AuthSee:      "查看",
	AuthExport:   "导出",
	AuthCreate:   "创建",
	AuthEdit:     "编辑",
	AuthDelete:   "删除",
	AuthProgress: "进展",
	AuthFile:     "附件",
	AuthSubTmpl:  "子任务",
}
