package contextlearn

import (
	"fmt"
	"net/http"
	"context"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
		
	}
}
