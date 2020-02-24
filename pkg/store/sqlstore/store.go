package sqlstore

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/opencars/koatuu/pkg/config"
	"github.com/opencars/koatuu/pkg/store"
)

// Store is an implementation of store.Store interface based on SQL.
type Store struct {
	db                 *sqlx.DB
	level1Repository   *Level1Repository
}

func (s *Store) Level1() store.Level1Repository {
	if s.level1Repository != nil {
		return s.level1Repository
	}

	s.level1Repository = &Level1Repository{
		store: s,
	}

	return s.level1Repository
}

// New returns new instance of Store.
func New(settings *config.Database) (*Store, error) {
	info := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		settings.Host,
		settings.Port,
		settings.User,
		settings.Name,
		settings.Password,
	)

	db, err := sqlx.Connect("postgres", info)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}
