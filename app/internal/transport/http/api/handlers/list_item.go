package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"perx/internal/service"
)

func ListItemHandler(ctx context.Context, sc service.Item) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		items, err := sc.ListItemService(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	})
}
