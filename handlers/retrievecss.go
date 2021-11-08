package handlers

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func RetrieveCSSHandler(rw http.ResponseWriter, r *http.Request) {
	handlerLog := log.With().Str("package", "handlers").Str("handler", "retrieveCss").Logger()

	rw.WriteHeader(http.StatusOK)

	handlerLog.Debug().Msg("Thanks for testing retrieve-css!")
	fmt.Fprint(rw, "Thanks for testing retrieve-css!")
}
