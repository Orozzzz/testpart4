package di

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"

	"project03/internal/datasource"
	"project03/internal/datasource/repository"
	"project03/internal/domain/service"
	"project03/internal/domain/strategy"
	"project03/internal/web/handlers"
)

func BuildContainer() *fx.App {
	return fx.New(
		fx.Provide(
			NewPostgresPool,
			repository.NewGameRepository,
			strategy.NewMinimaxStrategy,
			service.NewService,
			handlers.NewGameHandler,
			http.NewServeMux,
		),
		fx.Invoke(registerRoutes),
		fx.Invoke(startServer),
	)
}

func NewPostgresPool(lc fx.Lifecycle) (*pgxpool.Pool, error) {
	pool, err := datasource.NewPostgresPool()
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			pool.Close()
			return nil
		},
	})

	return pool, nil
}

func registerRoutes(mux *http.ServeMux, handler *handlers.GameHandler) {
	mux.HandleFunc("POST /game", handler.CreateGame)
	mux.HandleFunc("POST /game/{id}", handler.MakeMove)
	mux.HandleFunc("GET /game/{id}", handler.GetGame)
}

func startServer(lc fx.Lifecycle, mux *http.ServeMux) {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Server starting on :8080")

			go func() {
				if err := server.ListenAndServe(); err != nil &&
					err != http.ErrServerClosed {
					log.Printf("Server error: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			return server.Shutdown(shutdownCtx)
		},
	})
}