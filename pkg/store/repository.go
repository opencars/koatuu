package store

import (
	"github.com/opencars/koatuu/pkg/model"
)

type Level1Repository interface {
	Create(model.Level) error
	FindByID() (*model.Level, error)
}

type Level2Repository interface {
	Create(model.Level) error
	FindByID() (*model.Level, error)
}

type Level3Repository interface {
	Create(model.Level) error
	FindByID() (*model.Level, error)
}

type Level4Repository interface {
	Create(model.Level) error
	// FindByID() (*model.Level4, error)
}

type AreaRepository interface {
	Create(model.Level) error
	FindByID() (*model.Area, error)
}
