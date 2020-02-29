package sqlstore

import (
	"github.com/opencars/koatuu/pkg/model"
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

// func (r *Level4Repository) FindByID() (*model.Level4, error) {
// 	return nil, nil
// }
