package main

import (
	"flag"
	"log"

	"github.com/opencars/koatuu/internal/database"
	"github.com/opencars/koatuu/internal/model"
	"github.com/opencars/koatuu/pkg/parser"
)

var path = flag.String("path", "data/koatuu.json", "Path to data file")

func main() {
	flag.Parse()

	data, err := parser.Parse(*path)
	if err != nil {
		log.Fatal(err)
	}

	db := database.Must(database.DB())
	defer db.Close()

	for _, lvl1 := range data.Level1 {
		lvl := model.Level1Territory{
			Code: lvl1.Code,
			Name: lvl1.Name,
		}

		if err := db.Insert(&lvl); err != nil {
			panic(err)
		}
		for _, lvl2 := range lvl1.Level2 {
			lvl := model.Level2Territory{
				Code:     lvl2.Code,
				Name:     lvl2.Name,
				ParentID: lvl.ID,
			}
			if err := db.Insert(&lvl); err != nil {
				panic(err)
			}
			for _, lvl3 := range lvl2.Level3 {
				lvl := model.Level3Territory{
					Code:     lvl3.Code,
					Name:     lvl3.Name,
					Type:     lvl3.Type,
					ParentID: lvl.ID,
				}
				if err := db.Insert(&lvl); err != nil {
					panic(err)
				}
				for _, lvl4 := range lvl3.Level4 {
					lvl := model.Level4Territory{
						Code:     lvl4.Code,
						Name:     lvl4.Name,
						Type:     lvl4.Type,
						ParentID: lvl.ID,
					}
					if err := db.Insert(&lvl); err != nil {
						panic(err)
					}
				}
			}
		}
	}

}
