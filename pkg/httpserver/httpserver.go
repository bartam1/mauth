package httpserver

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/bartam1/mauth/pkg/logs"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func New(createHandler func(router chi.Router) http.Handler) *http.Server {
	return NewOnAddr(":3005", createHandler) //+os.Getenv("PORT")
}

func NewOnAddr(addr string, createHandler func(router chi.Router) http.Handler) *http.Server {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter)

	rootRouter := chi.NewRouter()
	// we are mounting all APIs under /api path
	rootRouter.Mount("/auth", createHandler(apiRouter))

	logrus.Info("Starting HTTP server")

	srv := &http.Server{
		Addr:    addr,
		Handler: rootRouter,
	}

	return srv
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(log.NewStructuredLogger(logrus.StandardLogger()))
	router.Use(middleware.Recoverer)
	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}

func CatchInterrupt(ctx context.Context, idleConnsClosed chan struct{}, srv *http.Server) {
	sigint := make(chan os.Signal, 1)

	// interrupt signal sent from terminal
	signal.Notify(sigint, os.Interrupt)
	// sigterm signal sent from kubernetes
	signal.Notify(sigint, syscall.SIGTERM)

	<-sigint

	logrus.Info("Interrupt received => Server is shutting down gracefully")

	// We received an interrupt signal, shut down.
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		logrus.Printf("HTTP server Shutdown interrupted: %v", err)
	}
	close(idleConnsClosed)

}
