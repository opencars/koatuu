package command

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/opencars/seedwork"
)

type InternalBulkDecode struct {
	Codes []string
}

func (c *InternalBulkDecode) Prepare() {}

func (c *InternalBulkDecode) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(
			&c.Codes,
			validation.Required.Error(seedwork.Required),
		),
	)
}
