package server

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/Jan-Ka/pikesies-srv/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func RunSingleHandlerServer(ctx context.Context, wg *sync.WaitGroup, addr string, pattern string, handler handlers.PikesiesHandler) {
	serverLog := log.With().Str("addr", addr).Str("pattern", pattern).Logger()

	mux := http.NewServeMux()
	mux.HandleFunc(pattern, handler)
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()

		shutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(shutCtx); err != nil {
			log.Error().Msgf("error shutting down handler %s\n", err)
		}

		serverLog.Info().Msg("handler is now closed")
		wg.Done()
	}()

	serverLog.Info().Msg("handler is starting")

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		serverLog.Error().Msgf("error starting up handler %s\n", err)
	}

	serverLog.Info().Msg("handler is closing")
}

func RunServer(ctx context.Context, wg *sync.WaitGroup, addr string, router *mux.Router) {
	serverLog := log.With().Str("addr", addr).Logger()

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		<-ctx.Done()

		shutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(shutCtx); err != nil {
			log.Error().Msgf("error shutting down server %s\n", err)
		}

		serverLog.Info().Msg("server is now closed")
		wg.Done()
	}()

	serverLog.Info().Msg("server is starting")

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		serverLog.Error().Msgf("error starting up server %s\n", err)
	}

	serverLog.Info().Msg("server is closing")
}
