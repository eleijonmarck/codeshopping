package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/eleijonmarck/codeshopping/cart"
	"github.com/eleijonmarck/codeshopping/redisdb"
)

const (
	defaultPort           = "8080"
	defaultRedisURL       = "http://127.0.0.1:6379"
	defaultRedisDBPort    = "6379"
	defaultDBName         = "codeshoppingDB"
	defaultRedisMaxIdle   = 3
	defaultRedisMaxActive = 32
)

func main() {

	// Setup repositories
	var (
		carts cart.Repository
	)

	// Create the logger used by the server
	logger := log.New(os.Stdout, "", 0)

	// Create new Redis Pool
	pool, err := redisdb.NewRedisPool(
		envString("REDISCLOUD_URL", defaultRedisURL),
		envInt("REDIS_MAX_IDLE", defaultRedisMaxIdle),
		envInt("REDIS_MAX_ACTIVE", defaultRedisMaxActive),
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer pool.Close()
	carts, _ = redisdb.NewCartRepository(defaultDBName, pool)

	// creates a http.ServeMux, register handlers to execute in response to routes
	mux := http.NewServeMux()

	// api
	mux.Handle("/carts", cart.Carts(carts))
	mux.Handle("/carts/", cart.GetCart(carts))
	mux.Handle("/", http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Pong")) })))

	// test storage
	storeTestData(carts)

	// start of server
	fmt.Printf("starting server at port %s\n", defaultPort)
	http.ListenAndServe(":8080", mux)
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func envInt(env string, fallback int) int {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	eInt, _ := strconv.Atoi(e)
	return eInt
}

func storeTestData(r cart.Repository) {
	test1 := cart.New("test1")
	if err := r.Store(test1); err != nil {
		panic(err)
	}
	log.Print("stored test1")

	test2 := cart.New("test2")
	if err2 := r.Store(test2); err2 != nil {
		panic(err2)
	}
	log.Print("stored test2")
}
