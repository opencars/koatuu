package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/opencars/seedwork"
	"github.com/opencars/seedwork/logger"

	"github.com/opencars/koatuu/pkg/domain"
	"github.com/opencars/koatuu/pkg/domain/command"
	"github.com/opencars/koatuu/pkg/domain/model"
)

type InternalService struct {
	s domain.Store
}

func NewInternalService(s domain.Store) *InternalService {
	return &InternalService{
		s: s,
	}
}

func (s *InternalService) Decode(ctx context.Context, c *command.InternalDecode) (*model.Result, error) {
	if err := seedwork.ProcessCommand(c); err != nil {
		return nil, err
	}

	return decode(ctx, s.s, c.Code)
}

func (s *InternalService) BulkDecode(ctx context.Context, c *command.InternalBulkDecode) (*model.BulkResult, error) {
	if err := seedwork.ProcessCommand(c); err != nil {
		return nil, err
	}

	bulkResult := model.BulkResult{
		Results: make([]model.WrappeResult, 0),
	}

	for _, code := range c.Codes {
		result, err := decode(ctx, s.s, code)
		if err != nil {
			err := handleErr(err)
			bulkResult.Results = append(bulkResult.Results, model.WrappeResult{Error: err})
			continue
		}

		bulkResult.Results = append(bulkResult.Results, model.WrappeResult{
			Result: result,
		})

	}

	return &bulkResult, nil
}

func handleErr(err error) *model.DecodingError {
	if err != nil {
		logger.Errorf("handleErr: %s", err)
	}

	var e seedwork.Error
	if errors.As(err, &e) {
		return &model.DecodingError{
			Messages: []string{e.Error()},
		}
	}

	var vErr seedwork.ValidationError
	if errors.As(err, &vErr) {
		errMessage := make([]string, 0)
		for k, vv := range vErr.Messages {
			for _, v := range vv {
				errMessage = append(errMessage, fmt.Sprintf("%s.%s", k, v))
			}
		}

		return &model.DecodingError{
			Messages: errMessage,
		}
	}

	return &model.DecodingError{
		Messages: []string{err.Error()},
	}
}
