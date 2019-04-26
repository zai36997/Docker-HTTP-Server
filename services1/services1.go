package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// message := r.URL.Path[1:]
	// m := map[string]string{"message": message}
	// json.NewEncoder(w).Encode(m)

	message := r.URL.Path[1:]
	connectService2("http://services2:8000/" + message)
	m := map[string]string{"message": message}
	json.NewEncoder(w).Encode(m)
}

func connectService2(url string) {
	//client := &http.Client{}
	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// body, err := ioutil.ReadAll(res.Body)
	// defer res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(body))

	req, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
