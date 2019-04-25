package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc("/", Handler2)
	// fmt.Println("Running server!")
	// log.Fatal(http.ListenAndServe(":8081", nil))
	http.HandleFunc("/", Handler2)
	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func Handler2(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path[1:]
	json.NewEncoder(w).Encode(map[string]string{"message2": message})

}

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	// Request /hello over port 8080 via the GET method
// 	r, err := http.Get("http://localhost:8081/hello")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Read the response body
// 	defer r.Body.Close()
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Print the response body to stdout
// 	fmt.Printf("%s\n", body)
// }
