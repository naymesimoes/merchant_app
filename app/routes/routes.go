package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to SumUp!")
	fmt.Println("Endpoint Hit: homePage")
}

func homeHell(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to hell!")
	fmt.Println("Endpoint Hit: homePage")
}

func getName(w http.ResponseWriter, r *http.Request) {
	k, _ := r.URL.Query()["name"]
	fmt.Fprintln(w, k)
}

func getName2(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/provisions/")

	fmt.Fprintln(w, id)
}

func HandleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/hell", homeHell)
	http.HandleFunc("/health-check", HealthCheck)
	http.HandleFunc("/get-name-2/", getName2)
	http.HandleFunc("/var/:name", getName)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
