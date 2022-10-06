package http

import "github.com/opencars/koatuu/pkg/handler"

func (s *server) configureRoutes() {
	router := s.router.PathPrefix("/api/v1/").Subrouter()
	router.Use(
		AuthorizationMiddleware(),
	)

	// GET /api/v1/koatuu/{id}.
	router.Handle("/koatuu/{id:[0-9]{10}}", s.findByCode()).Methods("GET")

	// GET /api/v1/koatuu/version.
	router.Handle("/koatuu/version", handler.Handler(s.Version)).Methods("GET")
}
