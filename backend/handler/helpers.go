package handler

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/google/uuid"
)

func setSessions(session *scs.SessionManager, r *http.Request, w http.ResponseWriter) {
	uuid := uuid.NewString()
	session.Put(r.Context(), uuid, "")
	w.WriteHeader(http.StatusOK)
	http.SetCookie(w, &http.Cookie{
		Name:  "session_token",
		Value: uuid,
	})
}

func paginate(r *http.Request) (int, int, int) {
	offset := r.Context().Value("offset").(int)
	size := r.Context().Value("size").(int)
	current_page := r.Context().Value("current_page").(int)
	return offset, size, current_page
}
