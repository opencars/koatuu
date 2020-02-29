package apiserver

import (
	"github.com/opencars/koatuu/pkg/model"
)

type Result struct {
	Level1 *model.Kek `json:"level1,omitempty"`
	Level2 *model.Kek `json:"level2,omitempty"`
	Level3 *model.Kek `json:"level3,omitempty"`
	Level4 *model.Kek `json:"level4,omitempty"`
	Name   string     `json:"name,omitempty"`
}
