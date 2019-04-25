package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running server!")
	go log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path[1:]
	m := map[string]string{"message": message}
	json.NewEncoder(w).Encode(m)

}
