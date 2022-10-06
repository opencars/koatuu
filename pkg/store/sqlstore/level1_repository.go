package sqlstore

import (
	"database/sql"
	"errors"

	"github.com/opencars/koatuu/pkg/domain/model"
)

type Level1Repository struct {
	store *Store
}

func (r *Level1Repository) Create(level model.Level) error {
	_, err := r.store.db.Exec(
		`INSERT INTO level1 (
			id, name
		) VALUES (
			$1, $2
		) ON CONFLICT DO NOTHING`,
		level.FirstLevel[:2], level.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Level1Repository) FindByID(id string) (*model.Level1or4, error) {
	var level model.Level1or4

	err := r.store.db.Get(&level,
		`SELECT id, name FROM level1 WHERE id = $1`,
		id,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrLocationNotFound
		}

		return nil, err
	}

	return &level, nil
}
