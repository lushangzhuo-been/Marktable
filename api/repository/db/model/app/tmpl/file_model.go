package tmpl

import (
	"database/sql/driver"
	"fmt"
)

// 文件转换状态枚举
type TransformedFileStatus string

const (
	Transforming TransformedFileStatus = "transforming"
	Succeed      TransformedFileStatus = "succeed"
	Failed       TransformedFileStatus = "failed"
	NoTransform  TransformedFileStatus = "no-transform"
)

type FileModel struct {
	Id                      int                   `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId                    int                   `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId                  int                   `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	IssueId                 string                `json:"issue_id" gorm:"column:issue_id;index;comment:issue_id;type:varchar(100);"`
	GroupId                 int                   `json:"group_id" gorm:"column:group_id;index;comment:分组ID;type:int(11);"`
	IsCurrentVersion        int                   `json:"is_current_version" gorm:"column:is_current_version;index;comment:是否为当前版本0-否，1是;type:int(11);"`
	VersionCode             int                   `json:"version_code" gorm:"column:version_code;index;comment:版本号;type:int(11);"`
	OriginalName            string                `json:"original_name" gorm:"column:original_name;comment:文件原始名;type:text;"`
	RelativePath            string                `json:"relative_path" gorm:"column:relative_path;comment:文件存储地址;type:text;"`
	TransformedOriginalName string                `json:"transformed_original_name" gorm:"column:transformed_original_name;comment:文件原始名;type:text;"`
	TransformedRelativePath string                `json:"transformed_relative_path" gorm:"column:transformed_relative_path;comment:文件存储地址;type:text;"`
	FileSize                int                   `json:"file_size" gorm:"column:file_size;comment:文件大小;type:int(11);"`
	TransformedStatus       TransformedFileStatus `json:"transformed_status" gorm:"column:transformed_status;comment:上传文件状态;type:ENUM('transforming','succeed','failed', 'no-transform');"`
	Userid                  int                   `json:"userid" gorm:"column:userid;comment:用户ID;type:int(50);"`
	CreatedAt               string                `json:"-" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
}

func (FileModel) TableName() string {
	return "user_tmpl_file"
}

// 验证枚举值的有效性
func (s TransformedFileStatus) IsValid() bool {
	switch s {
	case Transforming, Succeed, Failed, NoTransform:
		return true
	}
	return false
}

// 实现 Scanner 接口，用于从数据库读取
func (s *TransformedFileStatus) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	if str, ok := value.([]byte); ok {
		*s = TransformedFileStatus(str)
		return nil
	}
	return fmt.Errorf("failed to scan UserStatus")
}

// 实现 Valuer 接口，用于写入数据库
func (s TransformedFileStatus) Value() (driver.Value, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf("invalid UserStatus: %s", s)
	}
	return string(s), nil
}
