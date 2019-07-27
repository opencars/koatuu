package model

type LevelTerritory struct {
	ID   int32  `sql:"id,pk"`
	Code string `sql:"code,notnull,unique,type:varchar(11)"`
	Name string `sql:"name,notnull"`
	Type string `sql:"type"`
}

type Level1Territory struct {
	ID   int32  `sql:"id,pk"`
	Code string `sql:"code,notnull,unique,type:varchar(11)"`
	Name string `sql:"name,notnull"`

	tableName struct{} `sql:"level1"`
}

type Level2Territory struct {
	ID   int32  `sql:"id,pk"`
	Code string `sql:"code,notnull,unique,type:varchar(11)"`
	Name string `sql:"name,notnull,type:varchar(255)"`

	ParentID  int32    `sql:"level1_id"`
	tableName struct{} `sql:"level2"`
}

type Level3Territory struct {
	ID   int32  `sql:"id,pk"`
	Code string `sql:"code,notnull,unique,type:varchar(11)"`
	Name string `sql:"name,notnull,type:varchar(255)"`
	Type string `sql:"type",type:varchar(1)`

	ParentID  int32    `sql:"level2_id"`
	tableName struct{} `sql:"level3"`
}

type Level4Territory struct {
	ID   int32  `sql:"id,pk"`
	Code string `sql:"code,notnull,unique,type:varchar(11)"`
	Name string `sql:"name,notnull,type:varchar(255)"`
	Type string `sql:"type",type:varchar(1)`

	ParentID  int32    `sql:"level3_id"`
	tableName struct{} `sql:"level4"`
}
