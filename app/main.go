package main

import (
	"./repositories"
	"./routes"
)

func main() {
	repositories.Conection()
	defer repositories.CloseConection()

	routes.HandleRequest()
}
