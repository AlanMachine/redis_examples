package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	err = client.Set("say", "Hello, World!", 0).Err()
	if err != nil {
		log.Fatalln(err)
	}

	val, err := client.Get("say").Result()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(val)
}
