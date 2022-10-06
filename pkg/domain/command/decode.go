package command

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/opencars/schema"
	"github.com/opencars/schema/koatuu"
	"github.com/opencars/seedwork"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/opencars/koatuu/pkg/domain/model"
)

type Decode struct {
	UserID  string
	TokenID string
	Code    string
}

func (c *Decode) Prepare() {}

func (c *Decode) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(
			&c.UserID,
			validation.Required.Error(seedwork.Required),
		),
		validation.Field(
			&c.TokenID,
			validation.Required.Error(seedwork.Required),
		),
		validation.Field(
			&c.Code,
			validation.Required.Error(seedwork.Required),
			validation.Match(model.IsKOATUU).Error(seedwork.Invalid),
		),
	)
}

func (c *Decode) Event() schema.Producable {
	msg := koatuu.DecodingCalled{
		UserId:   c.UserID,
		TokenId:  c.TokenID,
		Code:     c.Code,
		CalledAt: timestamppb.New(time.Now().UTC()),
	}

	return schema.New(&source, &msg).Message(
		schema.WithSubject(schema.KoatuuCustomerActions),
	)
}
