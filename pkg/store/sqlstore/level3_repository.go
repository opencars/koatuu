package sqlstore

import (
	"github.com/opencars/koatuu/pkg/model"
)

type Level2Repository struct {
	store *Store
}

func (r *Level2Repository) Create(level2 model.Level2) error {
	_, err := r.store.db.NamedExec(
		`INSERT INTO level2 (
			id, name, kind, level1_id
		) VALUES (
			:id, :name, :kind, :level1_id
		)`,
		level2,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Level2Repository) FindByID() (*model.Level2, error) {
	return nil, nil
}
