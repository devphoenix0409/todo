package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/go-chi/chi"
	"github.com/rs/cors"

	"github.com/99designs/gqlgen/_examples/todo"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081", "http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(todo.NewExecutableSchema(todo.New()))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		// send this panic somewhere
		log.Print(err)
		debug.PrintStack()
		return errors.New("user message on panic")
	})

	router.Handle("/", playground.Handler("Todo", "/query"))
	router.Handle("/query", srv)
	log.Fatal(http.ListenAndServe(":8081", router))
}
