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
	LevelRepository
}

type Level3Repository interface {
	LevelRepository
}

type Level4Repository interface {
	LevelRepository
}
