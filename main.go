package main

import (
	"net/http"
	"gosql/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}


