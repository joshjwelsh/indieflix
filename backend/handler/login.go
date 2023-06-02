package handler

import (
	"encoding/json"
	"main/env"
	"main/model"
	"main/view"
	"net/http"
)

func Login(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := view.LoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		err = model.SelectUsersByUsername(env.DB, request.Username)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		setSessions(env.Session, r, w)
	}
}
