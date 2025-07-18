package tmpl

const (
	NodeModeHead   string = "head"
	NodeModeMiddle string = "middle"
	NodeModeEnd    string = "end"
)

var NodeModeMap = map[string]string{
	NodeModeHead:   "首节点",
	NodeModeMiddle: "中间节点",
	NodeModeEnd:    "末节点",
}
