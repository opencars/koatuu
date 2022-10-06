package domain

import (
	"context"

	"github.com/opencars/koatuu/pkg/domain/command"
	"github.com/opencars/koatuu/pkg/domain/model"
)

type LevelRepository interface {
	Create(model.Level) error
	FindByID(id string) (*model.Level1or4, error)
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

type Store interface {
	Level1() Level1Repository
	Level2() Level2Repository
	Level3() Level3Repository
	Level4() Level4Repository
}

type CustomerService interface {
	Decode(context.Context, *command.Decode) (*model.Result, error)
}

type InternalService interface {
	Decode(context.Context, *command.InternalDecode) (*model.Result, error)
	BulkDecode(context.Context, *command.InternalBulkDecode) (*model.BulkResult, error)
}
