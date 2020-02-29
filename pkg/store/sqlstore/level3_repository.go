package sqlstore

import (
	"strconv"

	"github.com/opencars/koatuu/pkg/model"
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

// func (r *Level3Repository) FindByID() (*model.Level3, error) {
// 	return nil, nil
// }
