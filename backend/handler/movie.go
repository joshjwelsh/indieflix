package handler

import (
	"encoding/json"
	"log"
	"main/env"
	"main/model"
	"main/view"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetAllMovies(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		offset := r.Context().Value("offset").(int)
		size := r.Context().Value("size").(int)
		current_page := r.Context().Value("current_page").(int)
		movies, err := model.AllMovies(env.DB, offset, size)
		if err != nil {
			log.Printf("Error querying movies: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		response := view.MoviesResponseList{
			Results: movies,
			Page:    current_page,
			Size:    len(movies),
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Printf("Error marshaling movies: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(jsonData)
	}
}

func GetMovies(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "id")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("Error converting id to int: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		movie, err := model.SelectMovies(env.DB, id)
		if err != nil {
			log.Printf("Error querying movies: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		jsonData, err := json.Marshal(movie)
		if err != nil {
			log.Printf("Error marshaling movies: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(jsonData)
	}
}
