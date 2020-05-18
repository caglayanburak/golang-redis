package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
)

var client *redis.Client

func homeLink(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'key' is missing")
        return
	}
	
	key := keys[0]
	fmt.Println(key)

	val, err := client.Get("Ghost:" + key).Result()
if err != nil {
    fmt.Println(err)
}

fmt.Println(val)
	fmt.Fprintf(w, val)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
