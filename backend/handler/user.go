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

func GetAllUsers(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		offset, size, current_page := paginate(r)

		users, err := model.AllUsers(env.DB, offset, size)

		if err != nil {
			log.Printf("Error querying users: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		response := view.UsersResponseList{
			Results: users,
			Page:    current_page,
			Size:    len(users),
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Printf("Error marshaling users: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(jsonData)
	}
}

func GetUsers(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "id")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("Error converting id to int: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		user, err := model.SelectUsers(env.DB, id)
		if err != nil {
			log.Printf("Error querying users: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		jsonData, err := json.Marshal(user)
		if err != nil {
			log.Printf("Error marshaling users: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(jsonData)
	}
}

func CreateUsers(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := view.UserRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			log.Printf("Error decoding user: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		id, err := model.InsertUsers(env.DB, request.Username, request.Password, request.Email)
		if err != nil {
			log.Printf("Error inserting user: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		user, err := model.SelectUsers(env.DB, id)
		if err != nil {
			log.Printf("Error querying user: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		jsonData, err := json.Marshal(user)
		if err != nil {
			log.Printf("Error marshaling user: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(jsonData)
	}
}

func DeleteUsers(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "id")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("Error converting id to int: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		err = model.DeleteUsers(env.DB, id)
		if err != nil {
			log.Printf("Error deleting user: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.WriteHeader(200)
	}
}

func UpdateUsers(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "id")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("Error converting id to int: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		request := view.UserRequest{}
		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			log.Printf("Error decoding user: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		err = model.UpdateUsers(env.DB, id, request.Username, request.Password, request.Email)
		if err != nil {
			log.Printf("Error updating user: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		user, err := model.SelectUsers(env.DB, id)
		if err != nil {
			log.Printf("Error querying user: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		jsonData, err := json.Marshal(user)
		if err != nil {
			log.Printf("Error marshaling user: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(jsonData)
	}
}
