package model

import (
	"errors"
	"strconv"
)

// SecondLevelKind ...
type SecondLevelKind string

const (
	// міста обласного значення;
	SecondLevelKindDiscrictCity SecondLevelKind = "DISTRICT_CITY"
	// райони Автономної Республіки Крим, області;
	SecondLevelKindDiscrict SecondLevelKind = "DISTRICT"
	// райони міст, що мають спеціальний статус.
	SecondLevelKindSpecialCityRegion SecondLevelKind = "SPECIAL_CITY_REGION"
)

type SecondLevel struct {
	ID           string          `json:"id" db:"id"`
	Name         string          `json:"name" db:"name"`
	Kind         SecondLevelKind `json:"kind" db:"kind"`
	FirstLevelID string          `json:"level1_id,omitempty" db:"level1_id"`
}

// NewSecondLevel creates new instance of NewSecondLevel.
func NewSecondLevel(id, name string) (*SecondLevel, error) {
	if len(id) != 10 {
		return nil, errors.New("koatuu code is not a valid")
	}

	kindAsInt, err := strconv.ParseInt(id[2:3], 10, 64)
	if err != nil {
		return nil, err
	}

	return &SecondLevel{
		ID:           id[:5],
		Name:         name,
		Kind:         ParseSecondLevelKind(kindAsInt),
		FirstLevelID: id[:2],
	}, nil
}

// ParseSecondLevelKind returns SecondLevelKind from i, where 0 < i < 4.
func ParseSecondLevelKind(i int64) SecondLevelKind {
	switch i {
	case 1:
		return SecondLevelKindDiscrictCity
	case 2:
		return SecondLevelKindDiscrict
	case 3:
		return SecondLevelKindSpecialCityRegion
	}

	return SecondLevelKind("?")
}
