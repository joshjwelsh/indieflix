package router

import (
	"main/env"
	"main/handler"
	m "main/middleware"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouterTree(env *env.Env) *chi.Mux {
	r := chi.NewRouter()

	r.Use(m.Response)
	r.Use(middleware.Timeout(time.Second * 60))

	r.Route("/genres", func(r chi.Router) {
		r.Use(m.Authenticate(env.Session))
		r.With(m.Paginate).Get("/", handler.GetAllGenres(env))
		r.Get("/{id}", handler.GetGenres(env))
		r.Post("/", handler.CreateGenres(env))
		r.Delete("/{id}", handler.DeleteGenres(env))
		r.Put("/{id}", handler.UpdateGenres(env))
	})

	r.Route("/login", func(r chi.Router) {
		r.Post("/", handler.Login(env))
	})

	r.Route("/movies", func(r chi.Router) {
		r.Use(m.Authenticate(env.Session))
		r.With(m.Paginate).Get("/", handler.GetAllMovies(env))
		r.Get("/{id}", handler.GetMovies(env))
	})

	r.Route("/register", func(r chi.Router) {
		r.Post("/", handler.CreateUsers(env))
	})

	r.Route("/sources", func(r chi.Router) {
		r.Use(m.Authenticate(env.Session))
		r.With(m.Paginate).Get("/", handler.GetAllSources(env))
		r.Get("/{id}", handler.GetSources(env))
		r.Post("/", handler.CreateSources(env))
		r.Delete("/{id}", handler.DeleteSources(env))
		r.Put("/{id}", handler.UpdateSources(env))
	})

	r.Route("/users", func(r chi.Router) {
		r.Use(m.Authenticate(env.Session))
		r.With(m.Paginate).Get("/", handler.GetAllUsers(env))
		r.Get("/{id}", handler.GetUsers(env))
		r.Delete("/{id}", handler.DeleteUsers(env))
		r.Put("/{id}", handler.UpdateUsers(env))
	})

	return r
}
