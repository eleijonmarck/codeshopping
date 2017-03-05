package cart

import (
	"encoding/json"
	"net/http"

	"fmt"
	"log"
)

// Carts return all carts in db, If used combined with POST method you are
// able to create a cart
func Carts(cr Repository) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch method := r.Method; method {

		case "POST":
			key := r.PostFormValue("CartID")

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

		default:
			allCarts := cr.FindAll()
			byteCart, _ := json.Marshal(&allCarts)
			if err2 := json.NewEncoder(w).Encode(ret{Carts: byteCart}); err2 != nil {
				w.Write([]byte(fmt.Sprintf(`{"error marshal": "%s"}`, err2.Error())))
			}
			w.Write(byteCart)
		}
	})
}
