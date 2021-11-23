package handlers

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func HealthHandler(rw http.ResponseWriter, r *http.Request) {
	handlerLog := log.With().Str("package", "handlers").Str("handler", "health").Logger()

	s := http.StatusOK
	rw.WriteHeader(s)
	st := http.StatusText(s)
	handlerLog.Debug().Int("code", s).Msg(st)
	fmt.Fprint(rw, st)
}
