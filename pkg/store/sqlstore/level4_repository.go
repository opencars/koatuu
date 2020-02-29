package sqlstore

import (
	"database/sql"

	"github.com/opencars/koatuu/pkg/model"
	"github.com/opencars/wanted/pkg/store"
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

func (r *Level4Repository) FindByID(id string) (*model.Kek, error) {
	var level model.Kek

	err := r.store.db.Get(&level,
		`SELECT id, name FROM level3 WHERE id = $1`,
		id,
	)

	if err == sql.ErrNoRows {
		return nil, store.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &level, nil
}
