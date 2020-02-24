package sqlstore

import (
	"github.com/opencars/koatuu/pkg/model"
)

type Level1Repository struct {
	store *Store
}

func (r *Level1Repository) Create(level1 model.Level1) error {
	_, err := r.store.db.NamedExec(
		`INSERT INTO level1 (
			id, name
		) VALUES (
			:id, :name
		)`,
		level1,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Level1Repository) FindByID() (*model.Level1, error) {
	return nil, nil
}
