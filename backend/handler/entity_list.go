package handler

import (
	"encoding/json"
	"main/env"
	"main/model"
	"main/view"
	"net/http"
	"strconv"
)

func GetAllEntityListsByUserId(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		offset := r.Context().Value("offset").(int)
		size := r.Context().Value("size").(int)
		current_page := r.Context().Value("current_page").(int)
		entity_lists, err := model.GetAllEntityListsByUserId(env.DB, user_id, offset, size)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := view.EntityListResponseList{
			Page:    current_page,
			Size:    size,
			Results: entity_lists,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	}
}
