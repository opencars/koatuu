package apiserver

import (
	"log"
	"net/http"

	"github.com/opencars/koatuu/pkg/config"
	"github.com/opencars/koatuu/pkg/store/sqlstore"
)

// Start starts the server with postgres store.
func Start(addr string, settings *config.Settings) error {
	store, err := sqlstore.New(&settings.DB)
	if err != nil {
		return err
	}

	srv := newServer(store)

	log.Printf("Listening on %s...\n", addr)
	return http.ListenAndServe(addr, srv)
}
