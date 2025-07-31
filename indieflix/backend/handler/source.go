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

func GetAllSources(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		offset := r.Context().Value("offset").(int)
		size := r.Context().Value("size").(int)
		current_page := r.Context().Value("current_page").(int)
		sources, err := model.AllSources(env.DB, offset, size)
		if err != nil {
			log.Printf("Error querying sources: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		response := view.SourcesResponseList{
			Results: sources,
			Page:    current_page,
			Size:    len(sources),
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Printf("Error marshaling sources: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		// Write the JSON data to the response writer
		w.Write(jsonData)
	}
}

func GetSources(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "id")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("Error converting id to int: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		source, err := model.SelectSources(env.DB, id)
		if err != nil {
			log.Printf("Error querying sources: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		jsonData, err := json.Marshal(source)
		if err != nil {
			log.Printf("Error marshaling sources: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(jsonData)
	}
}

func CreateSources(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sourceRequest := view.SourceRequest{}
		err := json.NewDecoder(r.Body).Decode(&sourceRequest)
		if err != nil {
			log.Printf("Error decoding request body: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		id, err := model.InsertSources(env.DB, sourceRequest.Name, sourceRequest.Website)
		if err != nil {
			log.Printf("Error inserting source: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		source, err := model.SelectSources(env.DB, id)
		if err != nil {
			log.Printf("Error querying sources: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		jsonData, err := json.Marshal(source)
		if err != nil {
			log.Printf("Error marshaling sources: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(jsonData)
	}

}

func DeleteSources(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "id")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("Error converting id to int: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		err = model.DeleteSources(env.DB, id)
		if err != nil {
			log.Printf("Error deleting source: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.WriteHeader(200)
	}

}

func UpdateSources(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "id")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("Error converting id to int: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		sourceRequest := view.SourceRequest{}
		err = json.NewDecoder(r.Body).Decode(&sourceRequest)
		if err != nil {
			log.Printf("Error decoding request body: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		err = model.UpdateSources(env.DB, id, sourceRequest.Name, sourceRequest.Website)
		if err != nil {
			log.Printf("Error updating source: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		source, err := model.SelectSources(env.DB, id)
		if err != nil {
			log.Printf("Error querying sources: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		jsonData, err := json.Marshal(source)
		if err != nil {
			log.Printf("Error marshaling sources: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(jsonData)
	}

}
