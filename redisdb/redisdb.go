package redisdb

import (
	"encoding/json"
	"fmt"
	"github.com/eleijonmarck/codeshopping/cart"
	"github.com/garyburd/redigo/redis"
	"net/url"
	"time"
)

type cartRepository struct {
	db   string
	Pool *redis.Pool
}

func (r *cartRepository) Store(cart *cart.Cart) error {

	c := r.Pool.Get()
	defer c.Close()
	serialized, err := json.Marshal(&cart)
	if err != nil {
		// error handle
		panic(err)
	}
	_, err2 := c.Do("SET", cart.CartID, string(serialized))
	if err2 != nil {
		panic(err2)
	}
	defer c.Close()
	return err
}

func (r *cartRepository) Find(key string) (*cart.Cart, error) {
	c := r.Pool.Get()
	defer c.Close()
	values, err := redis.Bytes(c.Do("GET", key))
	carty := cart.Cart{}
	err2 := json.Unmarshal(values, &carty)
	if err2 != nil {
		//
		fmt.Println("Error occured in redisgodb Find with error %s", err2)
	}
	return &carty, err
}

func (r *cartRepository) FindAll() []*cart.Cart {
	c := r.Pool.Get()
	// TODO: return slices of bytes and return them
	keys, err := redis.Strings(c.Do("KEYS", "test*"))
	fmt.Printf("Got the keys %s", keys)
	if err != nil {
		// handle it
		fmt.Printf("Couldnt get the keys %s", keys)
	}
	var result = make([]*cart.Cart, len(keys))
	carts, err2 := redis.ByteSlices(c.Do("MGET", keys))
	fmt.Println(carts)
	if err2 != nil {
		//
		panic(err)
	}
	for i := 0; i < len(carts); i++ {
		json.Unmarshal(carts[i], result[i])
	}
	fmt.Println("FindAll() returning the result %s", result)
	return result
}

// NewRedisPool returns a new redis pool to be able to use
func NewRedisPool(addr string, maxIdle, maxActive int) (*redis.Pool, error) {
	url, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	return &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", url.Host)
			if err != nil {
				return nil, err
			}

			if url.User != nil {
				password, _ := url.User.Password()
				_, err := c.Do("AUTH", password)
				if err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, _ time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}, nil
}

// NewCartRepository creates a repository for storage of the carts
func NewCartRepository(db string, pool *redis.Pool) (cart.Repository, error) {
	r := &cartRepository{
		db:   db,
		Pool: pool,
	}
	return r, nil
}
