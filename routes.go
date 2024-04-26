package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", app.showHome)

	mux.Get("/ws", http.HandlerFunc(app.ws.SocketEndPoint))

	return mux
}
