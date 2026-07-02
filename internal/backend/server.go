package backend

import (
	"net"
	"net/http"
	"net/url"
	"strings"

	openapidocs "1-task/docs/openapi"
	"1-task/internal/storage/postgres"
)

type Server struct {
	repo *postgres.Repository
}

func NewServer(repo *postgres.Repository) *Server {
	return &Server{repo: repo}
}

func (s *Server) routes() http.Handler {
	return HandlerWithOptions(s, ChiServerOptions{
		ErrorHandlerFunc: func(w http.ResponseWriter, _ *http.Request, err error) {
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		},
	})
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, withLocalhostCORS(s.routes()))
}

func withLocalhostCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		allowOrigin := isLocalhostOrigin(origin)

		if allowOrigin {
			headers := w.Header()
			headers.Set("Access-Control-Allow-Origin", origin)
			headers.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		}

		next.ServeHTTP(w, r)
	})
}

func isLocalhostOrigin(origin string) bool {
	parsed, err := url.Parse(origin)
	if err != nil {
		return false
	}

	hostname := strings.ToLower(parsed.Hostname())
	if hostname == "localhost" {
		return true
	}

	if ip := net.ParseIP(hostname); ip != nil {
		return ip.IsLoopback()
	}

	return false
}

func (s *Server) docsRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/openapi.yaml", serveOpenAPIYAML)
	mux.HandleFunc("/openapi.json", serveOpenAPIJSON)
	mux.HandleFunc("/", serveSwaggerUI)
	return mux
}

func (s *Server) RunSwaggerUI(addr string) error {
	return http.ListenAndServe(addr, s.docsRoutes())
}

func serveOpenAPIYAML(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/yaml; charset=utf-8")
	_, _ = w.Write(openapidocs.OpenAPIYAML)
}

func serveOpenAPIJSON(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write(openapidocs.OpenAPIJSON)
}

func serveSwaggerUI(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/swagger" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(swaggerUIHTML))
}

const swaggerUIHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Umbrella Cooldown API</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
  <style>
    body { margin: 0; background: #fafafa; }
  </style>
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.onload = function() {
      window.ui = SwaggerUIBundle({
        url: "/openapi.yaml",
        dom_id: "#swagger-ui"
      });
    };
  </script>
</body>
</html>
`
