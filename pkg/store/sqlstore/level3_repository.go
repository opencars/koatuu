package sqlstore

import (
	"database/sql"
	"strconv"

	"github.com/opencars/koatuu/pkg/model"
	"github.com/opencars/wanted/pkg/store"
)

type Level3Repository struct {
	store *Store
}

func (r *Level3Repository) Create(level model.Level) error {
	code, _ := strconv.Atoi(string(level.FirstLevel[5]))

	_, err := r.store.db.Exec(
		`INSERT INTO level3 (
			id, name, kind, level2_id, level1_id
		) VALUES (
			$1, $2, $3, $4, $5
		) ON CONFLICT DO NOTHING`,
		level.ThirdLevel[:8], level.Name, code, level.SecondLevel[:5], level.FirstLevel[:2],
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Level3Repository) FindByID(id string) (*model.Kek, error) {
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
