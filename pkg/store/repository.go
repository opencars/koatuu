package store

import (
	"github.com/opencars/koatuu/pkg/model"
)

type LevelRepository interface {
	Create(model.Level) error
	FindByID(id string) (*model.Kek, error)
}

type Level1Repository interface {
	LevelRepository
}

type Level2Repository interface {
	Create(level *model.SecondLevel) error
	FindByID(id string) (*model.SecondLevel, error)
}

type Level3Repository interface {
	Create(level *model.ThirdLevel) error
	FindByID(id string) (*model.ThirdLevel, error)
}

type Level4Repository interface {
	LevelRepository
}
