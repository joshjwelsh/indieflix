package handler

import (
	"encoding/json"
	"main/env"
	"main/model"
	"main/view"
	"net/http"
	"strconv"
)

func GetAllEntityListEntriesByEntityListId(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list_id, err := strconv.Atoi(r.URL.Query().Get("list_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		entries, err := model.GetAllEntityListEntriesByEntityListId(env.DB, list_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := view.EntityListEntryResponseList{
			Results: entries,
			Page:    1,
			Size:    len(entries),
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	}
}
