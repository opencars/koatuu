package apiserver

func (s *server) configureRouter() {
	s.router.Handle("/api/v1/koatuu/{code:[0-9]{10}}", s.findByCode()).Methods("GET", "OPTIONS")
}
