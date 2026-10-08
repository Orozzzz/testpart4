package di

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.uber.org/fx"

	"project03/internal/datasource/repository"
	"project03/internal/domain/service"
	"project03/internal/domain/strategy"
	"project03/internal/web/handlers"
)

func BuildContainer() *fx.App {
	return fx.New(
		fx.Provide(repository.NewGameRepository),
		fx.Provide(service.NewService),
		fx.Provide(handlers.NewGameHandler),
		fx.Provide(strategy.NewMinimaxStrategy),
		fx.Invoke(registerRoutes),
		fx.Invoke(startServer),
	)
}

func registerRoutes(handler *handlers.GameHandler) {
	http.HandleFunc("POST /game", handler.CreateGame)
	http.HandleFunc("POST /game/{id}", handler.MakeMove)
	http.HandleFunc("GET /game/{id}", handler.GetGame)
}

func startServer(lc fx.Lifecycle) {
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Server starting on :8080")
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("Server error: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down server...")
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			return server.Shutdown(ctx)
		},
	})
}
