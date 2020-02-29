package sqlstore

import (
	"database/sql"

	"github.com/opencars/koatuu/pkg/model"
)

type Level2Repository struct {
	store *Store
}

func (r *Level2Repository) Create(level *model.SecondLevel) error {
	_, err := r.store.db.Exec(
		`INSERT INTO level2 (
			id, name, kind, level1_id
		) VALUES (
			$1, $2, $3, $4
		) ON CONFLICT DO NOTHING`,
		level.ID, level.Name, level.Kind, level.FirstLevelID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Level2Repository) FindByID(id string) (*model.SecondLevel, error) {
	var level model.SecondLevel

	err := r.store.db.Get(&level,
		`SELECT id, name, kind FROM level2 WHERE id = $1`,
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
