package main

import (
	"net/http"
)

func (app *application) showHome(w http.ResponseWriter, _ *http.Request) {
	_ = app.render.Show(w, "index.page.gohtml", nil)
}
