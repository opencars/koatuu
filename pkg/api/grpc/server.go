package grpc

import (
	"context"

	"github.com/opencars/grpc/pkg/koatuu"

	"github.com/opencars/koatuu/pkg/domain/command"
)

type handler struct {
	koatuu.UnimplementedServiceServer
	api *API
}

func (h *handler) Decode(ctx context.Context, req *koatuu.DecodeRequest) (*koatuu.DecodeResultList, error) {
	items := make([]command.Item, 0, len(req.Codes))
	for _, code := range req.Codes {
		items = append(items, command.Item{
			Code: code,
		})
	}

	c := command.InternalBulkDecode{
		Items: items,
	}

	result, err := h.api.svc.BulkDecode(ctx, &c)
	if err != nil {
		return nil, handleErr(err)
	}

	dto := koatuu.DecodeResultList{
		Items: make([]*koatuu.DecodeResultItem, 0, len(result.Results)),
	}

	for i := range result.Results {
		dto.Items = append(dto.Items, toResultItem(&result.Results[i]))
	}

	return &dto, nil
}
