package model

import (
	"errors"
	"strconv"
)

type ThirdLevelKind string

const (
	// міста районного значення;
	ThirdLevelKindRegionCity ThirdLevelKind = "REGION_CITY"
	// райони в містах обласного значення;
	ThirdLevelKindDistrictCityRegion ThirdLevelKind = "DISTRICT_CITY_REGION"
	// селища міського типу, що входять до складу міськради;
	ThirdLevelKindCityUrbanSettlement ThirdLevelKind = "CITY_URBAN_SETTLEMENT"
	// селища міського типу, що входять до складу райради;
	ThirdLevelKindRegionUrbanSettlement ThirdLevelKind = "REGION_URBAN_SETTLEMENT"
	// селища міського типу, що входять до складу райради в місті;
	ThirdLevelKindCityRregionUrbanSettlement ThirdLevelKind = "CITY_REGION_URBAN_SETTLEMENT"
	// міста, що входять до складу міськради;
	ThirdLevelKindCity ThirdLevelKind = "CITY"
	// сільради, що входять до складу райради;
	ThirdLevelKindRegionSettlement ThirdLevelKind = "REGION_SETTLEMENT"
	// сільради, села, що входять до складу райради міста, міськради.
	ThirdLevelKindCitySettlement ThirdLevelKind = "CITY_SETTLEMENT"
)

type ThirdLevel struct {
	ID            string         `json:"id" db:"id"`
	Name          string         `json:"name" db:"name"`
	Kind          ThirdLevelKind `json:"kind" db:"kind"`
	FirstLevelID  string         `json:"level1_id,omitempty" db:"level1_id"`
	SecondLevelID string         `json:"level2_id,omitempty" db:"level2_id"`
}

// NewSecondLevel creates new instance of NewSecondLevel.
func NewThirdLevel(id, name string) (*ThirdLevel, error) {
	if len(id) != 10 {
		return nil, errors.New("koatuu code is not a valid")
	}

	kindAsInt, err := strconv.ParseInt(id[5:6], 10, 64)
	if err != nil {
		return nil, err
	}

	return &ThirdLevel{
		ID:            id[:8],
		Name:          name,
		Kind:          ParseThirdLevelKind(kindAsInt),
		FirstLevelID:  id[:2],
		SecondLevelID: id[:5],
	}, nil
}

// ParseThirdLevelKind returns SecondLevelKind from i, where 0 < i < 10, i != 2.
func ParseThirdLevelKind(i int64) ThirdLevelKind {
	switch i {
	case 1:
		return ThirdLevelKindRegionCity
	case 3:
		return ThirdLevelKindDistrictCityRegion
	case 4:
		return ThirdLevelKindCityUrbanSettlement
	case 5:
		return ThirdLevelKindRegionUrbanSettlement
	case 6:
		return ThirdLevelKindCityRregionUrbanSettlement
	case 7:
		return ThirdLevelKindCity
	case 8:
		return ThirdLevelKindRegionSettlement
	case 9:
		return ThirdLevelKindCitySettlement
	}

	return ThirdLevelKind("?")
}
