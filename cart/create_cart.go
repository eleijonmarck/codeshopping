package cart

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"
)

// CreateCart will create a item in the store
func CreateCart(cr Repository) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing key in query string", http.StatusBadRequest)
			return
		}

		newCart := New(key)
		err := cr.Store(newCart)
		if err != nil {
			fmt.Printf("Store of %s could not be completed due to %v\n", key, err)
			return
		}
		log.Printf("CreateCart")

		jsonCart, err := json.Marshal(newCart)
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonCart)
	})
}
