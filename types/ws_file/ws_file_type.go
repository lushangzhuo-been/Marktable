package ws_file

type WsFileInfoReq struct {
	WsId int `form:"ws_id" json:"ws_id" binding:"required"`
	Id   int `form:"id" json:"id" binding:"required"`
}

type AddWsFileReq struct {
	WsId int    `form:"ws_id" json:"ws_id" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
	Desc string `form:"desc" json:"desc"`
}

type WsFileDeleteReq struct {
	Id   int `form:"id" json:"id" binding:"required"`
	WsId int `form:"ws_id" json:"ws_id" binding:"required"`
}

type UpdateWsFileName struct {
	WsFileInfoReq
	Name string `form:"name" json:"name" binding:"required"`
	Desc string `form:"desc" json:"desc"`
}

type UpdateWsFileCoordinate struct {
	WsFileInfoReq
	Coordinate string `form:"coordinate" json:"coordinate" binding:"required"`
}

type UpdateWsFileAndTmplMergedCoordinate struct {
	WsId       int    `form:"ws_id" json:"ws_id" binding:"required"`
	Coordinate string `form:"coordinate" json:"coordinate" binding:"required"`
}

type TmplMoveInWsFile struct {
	WsFileInfoReq
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type TmplWsFileMoveReq struct {
	WsId     int    `form:"ws_id" json:"ws_id" binding:"required"`
	FromId   int    `form:"from_id" json:"from_id" binding:"required"`
	FromType string `form:"from_type" json:"from_type" binding:"required"`
	ToId     int    `form:"to_id" json:"to_id" binding:"required"`
	ToType   string `form:"to_type" json:"to_type" binding:"required"`
	Order    string `form:"order" json:"order" binding:"required"`
}

type IsAdminAndIsMember struct {
	IsAdmin  string `json:"is_admin"`
	IsMember string `json:"is_member"`
}

type TmplVo struct {
	MergedId      int    `json:"merged_id"`
	Id            int    `json:"id"`
	WsId          int    `json:"ws_id"`
	WsFileId      int    `json:"ws_file_id"`
	IsAdmin       string `json:"is_admin"`
	IsMember      string `json:"is_member"`
	ImplTableName string `json:"impl_table_name"`
	Name          string `json:"name"`
	Desc          string `json:"desc"`
	Mode          string `json:"mode"`
	Coordinate    string `json:"coordinate"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type WsFileAndTmplMergedVo struct {
	MergedId      int      `json:"merged_id"`
	Id            int      `json:"id"`
	WsId          int      `json:"ws_id"`
	WsFileId      int      `json:"ws_file_id"`
	IsAdmin       string   `json:"is_admin"`
	IsMember      string   `json:"is_member"`
	ImplTableName string   `json:"impl_table_name"`
	Name          string   `json:"name"`
	Desc          string   `json:"desc"`
	Mode          string   `json:"mode"`
	Children      []TmplVo `json:"children"`
	Coordinate    string   `json:"coordinate"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
}
