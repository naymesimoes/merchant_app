package controllers

import (
	"strconv"

	"../repositories"
	"../schemas"
)

func GetMerchantById(id string) (merchant schemas.Merchant) {
	idInt, _ := strconv.Atoi(id)
	merchant = repositories.GetMerchantById(idInt)

	return
}
