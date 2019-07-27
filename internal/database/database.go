package database

import (
	"fmt"
	"github.com/opencars/koatuu/internal/model"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Database interface makes handler testable.
type Adapter interface {
	Healthy() bool
	Select(
		model interface{},
		limit,
		condition string,
		params ...interface{},
	) error
}

func CreateSchema(db *pg.DB) error {
	tables := []interface{}{
		(*model.Level1Territory)(nil),
		(*model.Level2Territory)(nil),
		(*model.Level3Territory)(nil),
		(*model.Level4Territory)(nil),
	}

	for _, table := range tables {
		err := db.CreateTable(table, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func DB() (*pg.DB, error) {
	host := "localhost"
	port := "5432"

	if os.Getenv("DATABASE_HOST") != "" {
		host = os.Getenv("DATABASE_HOST")
	}

	if os.Getenv("DATABASE_PORT") != "" {
		port = os.Getenv("DATABASE_PORT")
	}

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		User:     "postgres",
		Password: "postgres",
		Database: "territories",
	})

	err := CreateSchema(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Must(db *pg.DB, err error) *pg.DB {
	if err != nil {
		panic(err)
	}

	return db
}
