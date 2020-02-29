package sqlstore

import (
	"database/sql"

	"github.com/opencars/koatuu/pkg/model"
)

type Level3Repository struct {
	store *Store
}

func (r *Level3Repository) Create(level *model.ThirdLevel) error {
	_, err := r.store.db.Exec(
		`INSERT INTO level3 (
			id, name, kind, level2_id, level1_id
		) VALUES (
			$1, $2, $3, $4, $5
		) ON CONFLICT DO NOTHING`,
		level.ID, level.Name, level.Kind, level.SecondLevelID, level.FirstLevelID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Level3Repository) FindByID(id string) (*model.ThirdLevel, error) {
	var level model.ThirdLevel

	err := r.store.db.Get(&level,
		`SELECT id, name, kind FROM level3 WHERE id = $1`,
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
