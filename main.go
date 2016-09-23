package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root"))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
