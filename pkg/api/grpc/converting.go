package grpc

import (
	"github.com/opencars/grpc/pkg/common"
	"github.com/opencars/grpc/pkg/koatuu"

	"github.com/opencars/koatuu/pkg/domain/model"
)

func toResultItem(item *model.WrappeResult) *koatuu.DecodeResultItem {
	if item.Error != nil {
		return &koatuu.DecodeResultItem{
			Error: &common.Error{
				Messages: item.Error.Messages,
			},
		}
	}

	return &koatuu.DecodeResultItem{
		Summary: item.Result.Summary,
		Parts:   item.Result.Parts,
	}
}
