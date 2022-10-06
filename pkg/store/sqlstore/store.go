package sqlstore

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/opencars/koatuu/pkg/config"
	"github.com/opencars/koatuu/pkg/domain"
)

// Store is an implementation of store.Store interface based on SQL.
type Store struct {
	db *sqlx.DB

	level1Repository *Level1Repository
	level2Repository *Level2Repository
	level3Repository *Level3Repository
	level4Repository *Level4Repository
}

func (s *Store) Level1() domain.Level1Repository {
	if s.level1Repository != nil {
		return s.level1Repository
	}

	s.level1Repository = &Level1Repository{
		store: s,
	}

	return s.level1Repository
}

func (s *Store) Level2() domain.Level2Repository {
	if s.level2Repository != nil {
		return s.level2Repository
	}

	s.level2Repository = &Level2Repository{
		store: s,
	}

	return s.level2Repository
}

func (s *Store) Level3() domain.Level3Repository {
	if s.level3Repository != nil {
		return s.level3Repository
	}

	s.level3Repository = &Level3Repository{
		store: s,
	}

	return s.level3Repository
}

func (s *Store) Level4() domain.Level4Repository {
	if s.level4Repository != nil {
		return s.level4Repository
	}

	s.level4Repository = &Level4Repository{
		store: s,
	}

	return s.level4Repository
}

// New returns new instance of Store.
func New(settings *config.Database) (*Store, error) {
	info := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		settings.Host,
		settings.Port,
		settings.User,
		settings.Name,
		settings.SSLMode,
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
