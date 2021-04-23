package main

import (
	"context"
	"net/http"

	"github.com/bartam1/mauth/internal/application/port"
	"github.com/bartam1/mauth/internal/application/service"
	"github.com/bartam1/mauth/internal/handlers/httphandler"
	repo "github.com/bartam1/mauth/internal/repositories"
	"github.com/bartam1/mauth/pkg/httpserver"
	"github.com/bartam1/mauth/pkg/logs"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	logs.Init()

	ctx := context.Background()

	repo, err := repo.NewMem()
	if err != nil {
		logrus.Panicf("Db error: ", err)
	}
	service := service.New(&repo)
	handl := httphandler.New(&service)

	srv := httpserver.New(func(router chi.Router) http.Handler {
		return port.HandlerFromMux(handl, router)
	})

	//Let server shutdown gracefully
	idleConnsClosed := make(chan struct{})
	go httpserver.CatchInterrupt(ctx, idleConnsClosed, srv)

	if err = srv.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Panicf("Cannot bind address: ", err)
	}
	<-idleConnsClosed

}
