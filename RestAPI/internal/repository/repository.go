package Repository

import (
	"github.com/go-redis/redis"
)

var Rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})
