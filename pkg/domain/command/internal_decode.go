package command

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/opencars/seedwork"
)

type InternalDecode struct {
	Code string
}

func (c *InternalDecode) Prepare() {}

func (c *InternalDecode) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(
			&c.Code,
			validation.Required.Error(seedwork.Required),
		),
	)
}
