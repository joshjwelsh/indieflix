package handler

import (
	"encoding/json"
	"main/env"
	"main/model"
	"net/http"
	"strconv"
)

func GetAllMovieGenresByMovieId(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movie_id, err := strconv.Atoi(r.URL.Query().Get("movie_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		genres, err := model.GetAllMovieGenresByMovieId(env.DB, movie_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonData, err := json.Marshal(genres)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	}
}
