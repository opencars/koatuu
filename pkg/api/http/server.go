package http

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/opencars/koatuu/pkg/domain"
	"github.com/opencars/koatuu/pkg/domain/command"
	"github.com/opencars/koatuu/pkg/handler"
	"github.com/opencars/koatuu/pkg/version"
)

type server struct {
	router *mux.Router

	svc domain.CustomerService
}

func newServer(svc domain.CustomerService) *server {
	srv := server{
		router: mux.NewRouter(),
		svc:    svc,
	}

	srv.configureRoutes()

	return &srv
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

		c := command.Decode{
			UserID:  UserIDFromContext(r.Context()),
			TokenID: TokenIDFromContext(r.Context()),
			Code:    mux.Vars(r)["id"],
		}

		result, err := s.svc.Decode(r.Context(), &c)
		if err != nil {
			return err
		}

		if err := json.NewEncoder(w).Encode(result); err != nil {
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
