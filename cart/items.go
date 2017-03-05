package cart

import (
	"encoding/json"
	"net/http"

	"log"
)

// Items returns the items that belong to the id of the cart
func Items(cr Repository) http.Handler {
	type ret struct {
		Items map[string]*CartItem `json:"items"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			// error
			return
		}
		defer r.Body.Close()
		items := cr.FindAll()

		jsonitems, err := json.Marshal(items)
		if err != nil {
			// error handling
		}
		log.Printf("cart item called")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonitems)
	})
}
