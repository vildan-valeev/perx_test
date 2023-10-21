package handlers

import (
	"fmt"
	"net/http"
)

func ListItemHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /")
	})
}
