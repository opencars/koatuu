package service

import (
	"context"

	"github.com/opencars/koatuu/pkg/domain"
	"github.com/opencars/koatuu/pkg/domain/model"
)

func decode(ctx context.Context, s domain.Store, code string) (*model.Result, error) {
	var result model.Result
	var err error

	result.Level1, err = s.Level1().FindByID(code[:2])
	if err != nil {
		return nil, err
	}

	if result.Level1 != nil {
		result.Parts = append(result.Parts, result.Level1.Name)
	}

	if code[2:5] != "000" {
		result.Level2, err = s.Level2().FindByID(code[:5])
		if err != nil {
			return nil, err
		}

		if result.Level2 != nil {
			result.Parts = append(result.Parts, result.Level2.Name)
		}
	}

	if code[5:8] != "000" {
		result.Level3, err = s.Level3().FindByID(code[:8])
		if err != nil {
			return nil, err
		}

		if result.Level3 != nil {
			result.Parts = append(result.Parts, result.Level3.Name)
		}
	}

	if code[8:] != "00" {
		result.Level4, err = s.Level4().FindByID(code)
		if err != nil {
			return nil, err
		}

		if result.Level4 != nil {
			result.Parts = append(result.Parts, result.Level4.Name)
		}
	}

	return &result, nil
}
