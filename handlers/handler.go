package handlers

import "net/http"

type PikesiesHandler = func(w http.ResponseWriter, r *http.Request)
