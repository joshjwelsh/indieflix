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

func GetAllGenres(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		offset := r.Context().Value("offset").(int)
		size := r.Context().Value("size").(int)
		current_page := r.Context().Value("current_page").(int)
		genres, err := model.AllGenres(env.DB, offset, size)
		if err != nil {
			log.Printf("Error querying genres: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		response := view.GenresResponseList{
			Results: genres,
			Page:    current_page,
			Size:    len(genres),
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Printf("Error marshaling genres: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(jsonData)
	}
}

func GetGenres(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "id")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("Error converting id to int: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		genre, err := model.SelectGenres(env.DB, id)
		if err != nil {
			log.Printf("Error querying genres: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		jsonData, err := json.Marshal(genre)
		if err != nil {
			log.Printf("Error marshaling genres: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(jsonData)
	}
}

func CreateGenres(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var genre view.GenreRequest
		err := json.NewDecoder(r.Body).Decode(&genre)
		if err != nil {
			log.Printf("Error decoding genre: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		id, err := model.InsertGenres(env.DB, genre.Name)
		if err != nil {
			log.Printf("Error inserting genre: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		jsonData, err := json.Marshal(id)
		if err != nil {
			log.Printf("Error marshaling genre: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(jsonData)
	}
}

func DeleteGenres(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "id")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("Error converting id to int: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		err = model.DeleteGenres(env.DB, id)
		if err != nil {
			log.Printf("Error deleting genre: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
	}
}

func UpdateGenres(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var genre view.GenreRequest
		err := json.NewDecoder(r.Body).Decode(&genre)
		if err != nil {
			log.Printf("Error decoding genre: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		param := chi.URLParam(r, "id")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("Error converting id to int: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		err = model.UpdateGenres(env.DB, id, genre.Name)
		if err != nil {
			log.Printf("Error updating genre: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		source, err := model.SelectGenres(env.DB, id)
		if err != nil {
			log.Printf("Error querying genre: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		jsonData, err := json.Marshal(source)
		if err != nil {
			log.Printf("Error marshaling genre: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(jsonData)
	}
}
