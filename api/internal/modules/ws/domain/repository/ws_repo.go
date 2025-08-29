package repository

import (
	"mark3/internal/modules/ws/domain/entity"
)

// WsRepository 空间仓储接口(领域层定义)
type WsRepository interface {
	CreateWs(ws *entity.WsModel) error
	// ... 其他仓储接口
}

// MemberRepository 空间记忆仓储接口(领域层定义)
type MemberRepository interface {
	CreateMember(member *entity.MemberModel) error
	// ... 其他仓储接口
}

// WsFolderRepository 空间下文件夹和流程排序仓储接口(领域层定义)
type WsFolderRepository interface {
	CreateWsFileAndTmplMergedCoordinate(wsFileAndTmplMergedCoordinate *entity.WsFileAndTmplMergedCoordinateModel) error
	CreateWsFolderAndTmplMerged(wsFileAndTmplMerged *entity.WsFileAndTmplMergedModel) error
	// ... 其他仓储接口
}
