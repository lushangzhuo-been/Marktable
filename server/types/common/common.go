package common

type BasePageReq struct {
	PageNum  int `form:"page_num" json:"page_num"`
	PageSize int `form:"page_size" json:"page_size"`
}

type BasePageResp struct {
	List interface{} `json:"list"`
	Cnt  int64       `json:"cnt"`
}

type BaseConfig struct {
	Value string `json:"value"`
	Label string `json:"label"`
}
