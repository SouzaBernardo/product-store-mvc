package main

import (
	"github.com/SouzaBernardo/first-go-api/controllers/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	routes.CreateRoutes()
	http.ListenAndServe(":8080", nil)
}
