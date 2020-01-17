package server

import (
	"context"
	"fmt"
	"net/http"
)

//Store struct
type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

//Server with store
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return
		}
		fmt.Fprint(w, data)
	}
}
