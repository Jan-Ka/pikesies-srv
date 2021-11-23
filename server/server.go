package server

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/Jan-Ka/pikesies-srv/config"
	"github.com/Jan-Ka/pikesies-srv/handlers"
	gHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func RunSingleHandlerServer(ctx context.Context, wg *sync.WaitGroup, addr string, pattern string, handler handlers.PikesiesHandler) {
	r := mux.NewRouter()
	r.HandleFunc(pattern, handler)

	RunServer(ctx, wg, addr, r)
}

func RunServer(ctx context.Context, wg *sync.WaitGroup, addr string, router *mux.Router) {
	serverLog := log.With().Str("addr", addr).Logger()
	cm := config.GetConfigManager()

	router.Use(gHandlers.RecoveryHandler())

	server := &http.Server{
		Addr: addr,

		WriteTimeout: cm.Config.WriteTimeout,
		ReadTimeout:  cm.Config.ReadTimeout,
		IdleTimeout:  cm.Config.IdleTimeout,

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
