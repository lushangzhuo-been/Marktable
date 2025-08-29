package repository

import (
	"mark3/internal/modules/tmpl/domain/entity"
)

// TmplRepository 空间仓储接口(领域层定义)
type TmplRepository interface {
	CreateTmpl(ws *entity.TmplModel) error
	// ... 其他仓储接口
}

// FieldRepository 空间仓储接口(领域层定义)
type FieldRepository interface {
	CreateFields(fields []entity.FieldModel) error
	// ... 其他仓储接口
}

// ScreenRepository 空间仓储接口(领域层定义)
type ScreenRepository interface {
	CreateScreens(screens []entity.ScreenModel) error
	// ... 其他仓储接口
}

// StatusRepository 空间仓储接口(领域层定义)
type StatusRepository interface {
	CreateStatus(status *entity.StatusModel) error
	// ... 其他仓储接口
}

// StepRepository 空间仓储接口(领域层定义)
type StepRepository interface {
	CreateSteps(steps []entity.StepModel) error
	// ... 其他仓储接口
}
