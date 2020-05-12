package routes

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API is up and running!")
	fmt.Println("Endpoint Hit: healthCheck")
}
