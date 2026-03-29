package main

import (
	"log"
	"net/http"
	"time"

	repo "github.com/Davidmuthee12/notes-rest-API/internal/adapters/postgresql/sqlc"
	"github.com/Davidmuthee12/notes-rest-API/internal/notes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

type application struct {
	config config
	db *pgx.Conn
}

type config struct {
	addr string
	db dbConfig
}

type dbConfig struct {
	dsn string
}

// MOUNT

func (app *application) mount() http.Handler{
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID) // important for rate limiting
	r.Use(middleware.RealIP) // important for rate limiting 
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	notesService := notes.NewService(repo.New(app.db))
	notesHandler := notes.NewHandler(notesService)
	r.Get("/notes", notesHandler.ListNotes)

	
	r.Post("/notes", notesHandler.CreateNote)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 30,
		IdleTimeout: time.Minute,
	}

	log.Printf("Server is running at addr %s", app.config.addr)
	return srv.ListenAndServe()
}
