package koatuu

type Level struct {
	Code   string  `json:"code"`
	Name   string  `json:"name"`
	Type   string  `json:"type,omitempty"`
	Level1 []Level `json:"level1,omitempty"`
	Level2 []Level `json:"level2,omitempty"`
	Level3 []Level `json:"level3,omitempty"`
	Level4 []Level `json:"level4,omitempty"`
}

//func Parse(levels1 []Level) []Level {
//	for _, level := range level1 {
//
//	}
//}
