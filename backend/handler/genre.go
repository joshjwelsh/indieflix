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

func AllGenres(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		genres, err := model.AllGenres(env.DB)
		if err != nil {
			log.Printf("Error querying genres: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		jsonData, err := json.Marshal(genres)
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
		w.Header().Set("Content-Type", "application/json")
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
