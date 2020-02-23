package model

type Level struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Head struct {
	Level1 []Level1 `json:"level1"`
}

type Level1 struct {
	Level
	Type   string   `json:"type"`
	Level2 []Level2 `json:"level2"`
}

type Level2 struct {
	Level
	Type   string   `json:"type"`
	Level3 []Level3 `json:"level3"`
}

type Level3 struct {
	Level
	Type   string   `json:"type"`
	Level4 []Level4 `json:"level4"`
}

type Level4 struct {
	Level
	Type string `json:"type"`
}
