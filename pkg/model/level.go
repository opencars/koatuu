package model

type Level struct {
	FirstLevel  string `json:"level1"`
	SecondLevel string `json:"level2"`
	ThirdLevel  string `json:"level3"`
	ForthLevel  string `json:"level4"`
	Category    string `json:"category"`
	Name        string `json:"name"`
}

type Area struct {
	FirstLevel  LevelInfo `json:"level1"`
	SecondLevel LevelInfo `json:"level2"`
	ThirdLevel  LevelInfo `json:"level3"`
	ForthLevel  LevelInfo `json:"level4"`
}

type LevelInfo struct {
	Name string
	Kind string
}
