package apiserver

import "github.com/opencars/koatuu/pkg/handler"

func (s *server) configureRouter() {
	s.router.Handle("/api/v1/koatuu/{id:[0-9]{10}}", s.findByCode()).Methods("GET", "OPTIONS")
	s.router.Handle("/api/v1/koatuu/version", handler.Handler(s.Version)).Methods("GET", "OPTIONS")
}
