package apiserver

import (
	"net/http"

	"github.com/opencars/koatuu/pkg/config"
	"github.com/opencars/koatuu/pkg/store/sqlstore"
)

// Start starts the server with postgres store.
func Start(settings *config.Settings) error {
	store, err := sqlstore.New(&settings.DB)
	if err != nil {
		return err
	}

	srv := newServer(store)

	return http.ListenAndServe(":8080", srv)
}
