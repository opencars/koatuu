package sqlstore

import (
	"database/sql"

	"github.com/opencars/koatuu/pkg/model"
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

func (r *Level1Repository) FindByID(id string) (*model.Kek, error) {
	var level model.Kek

	err := r.store.db.Get(&level,
		`SELECT id, name FROM level1 WHERE id = $1`,
		id,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &level, nil
}
