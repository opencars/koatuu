package service

import (
	"context"

	"github.com/opencars/schema"
	"github.com/opencars/seedwork"

	"github.com/opencars/koatuu/pkg/domain"
	"github.com/opencars/koatuu/pkg/domain/command"
	"github.com/opencars/koatuu/pkg/domain/model"
)

type CustomerService struct {
	s domain.Store
	p schema.Producer
}

func NewCustomerService(s domain.Store, p schema.Producer) *CustomerService {
	return &CustomerService{
		s: s,
		p: p,
	}
}

func (s *CustomerService) Decode(ctx context.Context, c *command.Decode) (*model.Result, error) {
	if err := seedwork.ProcessCommand(c); err != nil {
		return nil, err
	}

	result, err := decode(ctx, s.s, c.Code)
	if err != nil {
		return nil, err
	}

	if err := s.p.Produce(ctx, c.Event()); err != nil {
		return nil, err
	}

	return result, nil
}
