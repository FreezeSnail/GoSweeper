package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/cors"

	"github.com/freezesnail/goSweeper/app/mines/handlers"
	mines "github.com/freezesnail/goSweeper/app/mines/minesGame"
	"github.com/go-chi/chi"
)

func main() {

	l := log.New(os.Stdout, "Mines-api", log.LstdFlags)

	r := chi.NewRouter()

	mh := handlers.NewMines(l, make(map[int]*mines.Game))
	//r.Use(mh.CorsMiddleware)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*", "http://localhost:5000"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/Games", mh.ListGames)
	r.Get("/{id:[0-9]+}/Map", mh.GetBoard)

	r.Post("/{id:[0-9]+}/open", mh.Open)
	r.Post("/NewGame", mh.NewGame)

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

}
