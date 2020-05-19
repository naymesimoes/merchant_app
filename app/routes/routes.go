package routes

import (
	"fmt"
	"log"
	"net/http"

	"../controllers"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to SumUp!")
	fmt.Println("Endpoint Hit: homePage")
}

func getMerchantById(w http.ResponseWriter, r *http.Request) {
	// k, _ := r.URL.Query()["name"]
	// fmt.Fprintln(w, k)

	query := r.URL.Query()

	filters, present := query["merchant_id"]

	if !present || len(filters) == 0 {
		fmt.Println("filters not present")
	}

	merchant := controllers.GetMerchantById(filters[0])

	w.WriteHeader(200)
	fmt.Fprintln(w, merchant)
}

func HandleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/merchants", getMerchantById)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
