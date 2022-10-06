package command

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/opencars/seedwork"

	"github.com/opencars/koatuu/pkg/domain/model"
)

type Item struct {
	Code string
}

func (c *Item) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(
			&c.Code,
			validation.Required.Error(seedwork.Required),
			validation.Match(model.IsKOATUU).Error(seedwork.Invalid),
		),
	)
}

type InternalBulkDecode struct {
	Items []Item
}

func (c *InternalBulkDecode) Prepare() {}

func (c *InternalBulkDecode) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(
			&c.Items,
			validation.Required.Error(seedwork.Required),
		),
	)
}
