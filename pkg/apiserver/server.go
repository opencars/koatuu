package apiserver

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/opencars/koatuu/pkg/handler"
	"github.com/opencars/koatuu/pkg/store"
	"github.com/opencars/koatuu/pkg/version"
)

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Api-Key"})

	cors := handlers.CORS(origins, methods, headers)(s.router)
	cors.ServeHTTP(w, r)
}

func (s *server) findByCode() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		// code

		//
		return nil
	}
}

func (_ *server) Version(w http.ResponseWriter, r *http.Request) error {
	v := struct {
		Version string `json:"version"`
		Go      string `json:"go"`
	}{
		Version: version.Version,
		Go:      runtime.Version(),
	}

	if err := json.NewEncoder(w).Encode(v); err != nil {
		return err
	}

	return nil
}
