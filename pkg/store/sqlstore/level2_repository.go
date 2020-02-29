package sqlstore

import (
	"strconv"

	"github.com/opencars/koatuu/pkg/model"
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
		level.SecondLevel[:5], level.Name, code, level.SecondLevel[:2],
	)

	if err != nil {
		return err
	}

	return nil
}

// func (r *Level2Repository) FindByID() (*model.Level2, error) {
// 	return nil, nil
// }
