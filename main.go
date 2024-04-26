package main

import (
	"fmt"
	"github.com/tsawler/page"
	"github.com/tsawler/toolbox"
	"github.com/tsawler/ws"
	"net/http"
	"time"
)

type application struct {
	ws        *ws.Sockets
	render    *page.Render
	eventChan chan string
}

const port = 8080

func main() {
	render := page.New()
	render.UseCache = false

	app := application{
		ws:        ws.New(),
		render:    render,
		eventChan: make(chan string, 100),
	}

	// start websocket functionality
	fmt.Println("Starting websocket functionality...")
	go app.ws.ListenToWsChannel()

	go app.RandomString()

	// start the web server
	fmt.Printf("Starting web server on port %d...\n", port)

	// create http server
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func (app *application) RandomString() {
	var t toolbox.Tools

	for {
		time.Sleep(3 * time.Second)
		for k, _ := range app.ws.Clients {
			payload := ws.Payload{
				MessageType: ws.JSONMessage,
				Message:     t.RandomString(5),
				Conn:        k,
			}
			app.ws.ClientChan <- payload
		}
	}
}
