package tmpl

type GetFileList struct {
	CommonReq
	IssueId          string `form:"issue_id" json:"issue_id" binding:"required"`
	GroupId          int    `form:"group_id" json:"group_id" binding:"required"`
	IsCurrentVersion int    `form:"is_current_version" json:"is_current_version"`
}

type UploadFile struct {
	CommonReq
	IssueId          string `form:"issue_id" json:"issue_id" binding:"required"`
	GroupId          int    `form:"group_id" json:"group_id" binding:"required"`
	IsCurrentVersion int    `form:"is_current_version" json:"is_current_version" binding:"required"`
}

type UpdateFileIsCurrentVersion struct {
	CommonReq
	Id               string `form:"id" json:"id" binding:"required"`
	IssueId          string `form:"issue_id" json:"issue_id" binding:"required"`
	GroupId          int    `form:"group_id" json:"group_id" binding:"required"`
	IsCurrentVersion int    `form:"is_current_version" json:"is_current_version" binding:"required"`
}

type DownloadFile struct {
	CommonReq
	IssueId          string `form:"issue_id" json:"issue_id" binding:"required"`
	Id               string `form:"id" json:"id" binding:"required"`
	DownloadFileType string `form:"download_file_type" json:"download_file_type" binding:"required"`
}

type DeleteFile struct {
	CommonReq
	IssueId string `form:"issue_id" json:"issue_id" binding:"required"`
	GroupId int    `form:"group_id" json:"group_id" binding:"required"`
	Id      string `form:"id" json:"id" binding:"required"`
}

type UploadImageForHtml struct {
	CommonReq
}
