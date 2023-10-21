package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vildan-valeev/perx_test/internal/service"
	"github.com/vildan-valeev/perx_test/internal/transport/dto"
)

func ItemToQueueHandler(ctx context.Context, sc service.Item) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: check header,

		var i dto.ItemToQueueDTO
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = sc.AddItemToQueueService(ctx, &i)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		fmt.Fprintf(w, "Done!")
	})
}
