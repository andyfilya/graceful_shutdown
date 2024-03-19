package server

import (
	"context"
	"net"
	"net/http"
)

func InitServer(ctx context.Context, port string) *http.Server {
	return &http.Server{
		Addr: ":" + port,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			<-r.Context().Done()
			w.WriteHeader(http.StatusOK)
		}),
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}
}
