package sqlstore

import (
	"database/sql"
	"errors"

	"github.com/opencars/koatuu/pkg/domain/model"
)

type Level4Repository struct {
	store *Store
}

func (r *Level4Repository) Create(level model.Level) error {
	_, err := r.store.db.Exec(
		`INSERT INTO level4 (
			id, name, level3_id, level2_id, level1_id
		) VALUES (
			$1, $2, $3, $4, $5
		) ON CONFLICT DO NOTHING`,
		level.ForthLevel, level.Name, level.ThirdLevel[:8], level.SecondLevel[:5], level.FirstLevel[:2],
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Level4Repository) FindByID(id string) (*model.Level1or4, error) {
	var level model.Level1or4

	err := r.store.db.Get(&level,
		`SELECT id, name FROM level3 WHERE id = $1`,
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
