package handler

import (
	"encoding/json"
	"main/env"
	"main/model"
	"net/http"
	"strconv"
)

func GetAllCommentsByMovie(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movie_id, err := strconv.Atoi(r.URL.Query().Get("movie_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		comments, err := model.GetAllCommentsByMovie(env.DB, movie_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(comments)
	}
}

func GetAllCommentsByUser(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		comments, err := model.GetAllCommentsByUser(env.DB, user_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(comments)
	}
}
