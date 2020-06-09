package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"../repositories"
	"../schemas"
)

func GetMerchantById(id string) (merchant *schemas.Merchant) {
	idInt, _ := strconv.Atoi(id)
	merchant = repositories.GetMerchantById(idInt)

	if merchant == nil {
		return nil
	}

	return
}

func CreateMerchant(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var merchant schemas.Merchant

	err = json.Unmarshal(body, &merchant)
	if err != nil {
		fmt.Fprintf(w, "Error Unmarshalling body: %v", err)
		http.Error(w, "can't unmarshal body", http.StatusBadRequest)
		return
	}

	result, err := repositories.InsertMerchant(merchant)
	if err != nil {
		fmt.Fprintf(w, "Error inserting merchant in database: %v", err)
		http.Error(w, "can't insert merchant in database", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, result)
}
