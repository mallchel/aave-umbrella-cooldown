package backend

import (
	"net/http"

	"1-task/internal/storage/postgres"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	repo *postgres.Repository
}

func NewServer(repo *postgres.Repository) *Server {
	return &Server{repo: repo}
}

func (s *Server) routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/healthz", s.handleHealth)
	r.Get("/withdraw-flows", s.handleListWithdrawFlows)
	return r
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s.routes())
}
