package handler

import (
	"encoding/json"
	"log"
	"main/env"
	"main/model"
	"net/http"
	"strconv"
	"strings"
)

func GetAllSources(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		sources, err := model.AllSources(env.DB)
		if err != nil {
			log.Printf("Error querying sources: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		// Marshal the slice of SourceResponse
		jsonData, err := json.Marshal(sources)
		if err != nil {
			log.Printf("Error marshaling sources: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		// Write the JSON data to the response writer
		w.Write(jsonData)
	}
}

func GetSource(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := strings.TrimPrefix(r.URL.Path, "/sources/")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("Error converting id to int: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
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
