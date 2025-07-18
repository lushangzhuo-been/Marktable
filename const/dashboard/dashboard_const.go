package dashboard

// BoardModeEnumVal 定义图表模式枚举值
type BoardModeEnumVal string

const (
	BaseHistogram  BoardModeEnumVal = "base_histogram"
	GroupHistogram BoardModeEnumVal = "group_histogram"
	StackHistogram BoardModeEnumVal = "stack_histogram"
	BaseLine       BoardModeEnumVal = "base_line"
	GroupLine      BoardModeEnumVal = "group_line"
	BasePie        BoardModeEnumVal = "base_pie"
	DonutPie       BoardModeEnumVal = "donut_pie"
)

// BoardModeEnumCn 定义图表模式中文名称
type BoardModeEnumCn string

const (
	BaseHistogramCn  BoardModeEnumCn = "基础柱状图"
	GroupHistogramCn BoardModeEnumCn = "分组柱状图"
	StackHistogramCn BoardModeEnumCn = "堆叠柱状图"
	BaseLineCn       BoardModeEnumCn = "基础折线图"
	GroupLineCn      BoardModeEnumCn = "分组折线图"
	BasePieCn        BoardModeEnumCn = "基础饼图"
	DonutPieCn       BoardModeEnumCn = "环形饼图"
)

// BoardYaxisEnumVal 定义Y轴类型枚举值
type BoardYaxisEnumVal string

const (
	Count BoardYaxisEnumVal = "count"
)

// BoardYaxisEnumCn 定义Y轴类型中文名称
type BoardYaxisEnumCn string

const (
	CountCn BoardYaxisEnumCn = "计数"
)

// FieldConfig 定义字段配置结构
type FieldConfig struct {
	Field   string `json:"filed"`
	LabelCn string `json:"label_cn"`
}

// TmplBoardConstant 包含所有图表相关常量
type TmplBoardConstant struct{}

// GetBoardList 获取图表类型列表
func (t *TmplBoardConstant) GetBoardList() []map[string]interface{} {
	return []map[string]interface{}{
		{"label": BaseHistogramCn, "value": BaseHistogram},
		{"label": GroupHistogramCn, "value": GroupHistogram},
		{"label": StackHistogramCn, "value": StackHistogram},
		{"label": BaseLineCn, "value": BaseLine},
		{"label": GroupLineCn, "value": GroupLine},
		{"label": BasePieCn, "value": BasePie},
		{"label": DonutPieCn, "value": DonutPie},
	}
}

// GetBoardYaxisList 获取Y轴类型列表
func (t *TmplBoardConstant) GetBoardYaxisList() []map[string]interface{} {
	return []map[string]interface{}{
		{"label": CountCn, "value": Count},
	}
}

// GetBoardFieldConfig 获取图表字段配置
func (t *TmplBoardConstant) GetBoardFieldConfig() map[BoardModeEnumVal][]FieldConfig {
	return map[BoardModeEnumVal][]FieldConfig{
		BaseHistogram: {
			{"title", "标题"},
			{"x_axis", "X轴"},
			{"y_axis", "Y轴"},
			{"order", "排序"},
			{"show_empty", "是否显示空数据"},
			{"filter", "筛选"},
		},
		GroupHistogram: {
			{"title", "标题"},
			{"x_axis", "X轴"},
			{"y_axis", "Y轴"},
			{"group_axis", "分组"},
			{"order", "排序"},
			{"show_empty", "是否显示空数据"},
			{"filter", "筛选"},
		},
		StackHistogram: {
			{"title", "标题"},
			{"x_axis", "X轴"},
			{"y_axis", "Y轴"},
			{"group_axis", "分组"},
			{"order", "排序"},
			{"show_empty", "是否显示空数据"},
			{"filter", "筛选"},
		},
		BaseLine: {
			{"title", "标题"},
			{"x_axis", "X轴"},
			{"y_axis", "Y轴"},
			{"order", "排序"},
			{"show_empty", "是否显示空数据"},
			{"filter", "筛选"},
		},
		GroupLine: {
			{"title", "标题"},
			{"x_axis", "X轴"},
			{"y_axis", "Y轴"},
			{"order", "排序"},
			{"show_empty", "是否显示空数据"},
			{"filter", "筛选"},
		},
		BasePie: {
			{"title", "标题"},
			{"x_axis", "X轴"},
			{"y_axis", "Y轴"},
			{"order", "排序"},
			{"show_empty", "是否显示空数据"},
			{"filter", "筛选"},
		},
		DonutPie: {
			{"title", "标题"},
			{"x_axis", "X轴"},
			{"y_axis", "Y轴"},
			{"order", "排序"},
			{"show_empty", "是否显示空数据"},
			{"filter", "筛选"},
		},
	}
}
