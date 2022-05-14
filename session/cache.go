package cognitosession

import (
    "context"
    "github.com/go-redis/redis/v8"
    "fmt"
	"os"
	"encoding/json"
)

var redis_sample string = "sss"
var ctx = context.Background()


//https://github.com/go-redis/redis


type RedisStruct struct {
	Addr string
	Password string
	DB int
}

func RedisClient() *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr:	  os.Getenv("REDIS_ENDPOINT"),
		Password: "", // no password set
		DB:		  0,  // use default DB
	})
	return rdb
}


func Redis_Save(key string, value string)  {
	rdb := RedisClient()
	json_value, error := json.Marshal(value)
	if error != nil{
		panic(error)
	}
	err := rdb.Set(ctx, key, json_value, 0).Err()
	if err != nil {
		panic(err)
	}
	
}

func Redis_Get(key string)  {
	rdb := RedisClient()

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
