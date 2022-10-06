package model

import "regexp"

var (
	IsKOATUU = regexp.MustCompile(`^[0-9]{10}$`)
)

type Level struct {
	FirstLevel  string `json:"level1"`
	SecondLevel string `json:"level2"`
	ThirdLevel  string `json:"level3"`
	ForthLevel  string `json:"level4"`
	Category    string `json:"category"`
	Name        string `json:"name"`
}

type Level1or4 struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
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

type Result struct {
	Level1  *Level1or4   `json:"level1,omitempty"`
	Level2  *SecondLevel `json:"level2,omitempty"`
	Level3  *ThirdLevel  `json:"level3,omitempty"`
	Level4  *Level1or4   `json:"level4,omitempty"`
	Parts   []string     `json:"parts,omitempty"`
	Summary string       `json:"name,omitempty"`
}

type WrappeResult struct {
	Result *Result
	Error  *DecodingError
}

type DecodingError struct {
	Messages []string `json:"messages"`
}

type BulkResult struct {
	Results []WrappeResult
}
