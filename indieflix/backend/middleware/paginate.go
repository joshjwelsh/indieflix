package middleware

import (
	"context"
	"net/http"
	"strconv"
)

func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		curr_page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || curr_page < 1 {
			curr_page = 1
		}
		page_size, err := strconv.Atoi(r.URL.Query().Get("page_size"))
		if err != nil || page_size < 1 {
			page_size = 10
		}
		page_offset := (curr_page - 1) * page_size
		ctx := context.WithValue(r.Context(), "offset", page_offset)
		ctx = context.WithValue(ctx, "size", page_size)
		ctx = context.WithValue(ctx, "current_page", curr_page)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
