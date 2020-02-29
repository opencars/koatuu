package sqlstore

import (
	"database/sql"
	"strconv"

	"github.com/opencars/koatuu/pkg/model"
	"github.com/opencars/wanted/pkg/store"
)

type Level2Repository struct {
	store *Store
}

func (r *Level2Repository) Create(level model.Level) error {
	code, _ := strconv.Atoi(string(level.SecondLevel[2]))

	_, err := r.store.db.Exec(
		`INSERT INTO level2 (
			id, name, kind, level1_id
		) VALUES (
			$1, $2, $3, $4
		) ON CONFLICT DO NOTHING`,
		level.SecondLevel[:5], level.Name, code, level.FirstLevel[:2],
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Level2Repository) FindByID(id string) (*model.Kek, error) {
	var level model.Kek

	err := r.store.db.Get(&level,
		`SELECT id, name FROM level2 WHERE id = $1`,
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
