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
	headers := handlers.AllowedHeaders([]string{"Api-Key", "X-Api-Key"})

	cors := handlers.CORS(origins, methods, headers)(s.router)
	cors.ServeHTTP(w, r)
}

func (s *server) findByCode() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")
		id := mux.Vars(r)["id"]

		res := Result{}
		var err error

		res.Level1, err = s.store.Level1().FindByID(id[:2])
		if err != nil {
			return err
		}

		if res.Level1 != nil {
			res.Name = res.Level1.Name
		}

		if id[2:5] != "000" {
			res.Level2, err = s.store.Level2().FindByID(id[:5])
			if err != nil {
				return err
			}

			if res.Level2 != nil {
				res.Name += ", " + res.Level2.Name
			}
		}

		if id[5:8] != "000" {
			res.Level3, err = s.store.Level3().FindByID(id[:8])
			if err != nil {
				return err
			}

			if res.Level3 != nil {
				res.Name += ", " + res.Level3.Name
			}
		}

		if id[8:] != "00" {
			res.Level4, err = s.store.Level4().FindByID(id)
			if err != nil {
				return err
			}

			if res.Level4 != nil {
				res.Name += ", " + res.Level4.Name
			}
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			return err
		}

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
