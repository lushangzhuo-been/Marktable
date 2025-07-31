package tmpl

type GetUserViewsReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type GetViewInfoReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Id     int `form:"id" json:"id" binding:"required"`
}

type CreateViewReq struct {
	WsId      int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId    int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Mode      string `form:"mode" json:"mode" binding:"required"`
	Name      string `form:"name" json:"name" binding:"required"`
	Filter    string `form:"filter" json:"filter"`
	Lor       string `form:"lor" json:"lor"`
	SortOrder string `form:"sort_order" json:"sort_order"`
	Columns   string `form:"columns" json:"columns"`
	CardGroup string `form:"card_group" json:"card_group"`
}

type RenameViewReq struct {
	WsId   int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Id     int    `form:"id" json:"id" binding:"required"`
	Name   string `form:"name" json:"name" binding:"required"`
}

type UpdateViewFilterReq struct {
	WsId   int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Id     int    `form:"id" json:"id" binding:"required"`
	Filter string `form:"filter" json:"filter"`
	Lor    string `form:"lor" json:"lor"`
}

type UpdateViewSortOrderReq struct {
	WsId      int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId    int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Id        int    `form:"id" json:"id" binding:"required"`
	SortOrder string `form:"sort_order" json:"sort_order"`
}

type UpdateViewColumnsReq struct {
	WsId    int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId  int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Id      int    `form:"id" json:"id" binding:"required"`
	Columns string `form:"columns" json:"columns" binding:"required"`
}

type UpdateViewCardGroupReq struct {
	WsId      int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId    int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Id        int    `form:"id" json:"id" binding:"required"`
	CardGroup string `form:"card_group" json:"card_group"`
}

type DeleteViewReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Id     int `form:"id" json:"id" binding:"required"`
}

type PinViewReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Id     int `form:"id" json:"id" binding:"required"`
}

type UnPinViewReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Id     int `form:"id" json:"id" binding:"required"`
}

type SetViewCoordinateReq struct {
	WsId       int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId     int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Coordinate string `form:"coordinate" json:"coordinate" binding:"required"`
}

type ViewFilterItem struct {
	FieldKey string `json:"field_key"`
	Op       string `json:"op"`
	Value    string `json:"value"`
}

type ViewSortOrderItem struct {
	FieldKey string `json:"field_key"`
	Sort     string `json:"sort"`
}

type ViewColumnsItem struct {
	FieldKey string `json:"field_key"`
	Visible  string `json:"visible"`
}
