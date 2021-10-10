package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/freezesnail/goSweeper/app/mines/handlers"
	mines "github.com/freezesnail/goSweeper/app/mines/minesGame"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "Mines-api", log.LstdFlags)

	r := mux.NewRouter()

	mh := handlers.NewMines(l, make(map[int]*mines.Game))

	getR := r.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/Games", mh.ListGames)

	postR := r.Methods(http.MethodPost).Subrouter()
	//	postR.HandleFunc("/{id:[0-9]+}/open", mh.Open)
	postR.HandleFunc("/NewGame", mh.NewGame)

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	s := http.Server{
		Addr:         "8090",
		Handler:      ch(r),
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive

	}

	go func() {
		l.Println("Starting server on port 8090")

		err := http.ListenAndServe(":8090", r)
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}
