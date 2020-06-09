package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../controllers"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to SumUp!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleMerchant(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getMerchantById(w, r)
	case "POST":
		controllers.CreateMerchant(w, r)
	}
}

func getMerchantById(w http.ResponseWriter, r *http.Request) {
	// k, _ := r.URL.Query()["name"]
	// fmt.Fprintln(w, k)

	query := r.URL.Query()

	filters, present := query["merchant_id"]

	if !present || len(filters) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("filters not present")
		return
	}

	merchant := controllers.GetMerchantById(filters[0])

	if merchant == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not Found")
		return
	}

	responseBody, _ := json.Marshal(merchant)

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
	fmt.Fprintln(w, "OK")
}

func HandleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/merchants", handleMerchant)
