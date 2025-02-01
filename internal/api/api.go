package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/haydenisler/api/internal/middleware"
)

type api struct{}

func NewAPI(ctx context.Context) *api {
	return &api{}
}

func (a *api) Server(port int) *http.Server {
	stack := middleware.CreateStack(
		middleware.Logging,
	)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: stack(a.Routes()),
	}
}

func (a *api) Routes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /", a.getDebugHandler)

	return r
}
