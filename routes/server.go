package routes

import (
	"net/http"
	"gosql/controllers"
)
func LoadRoutes(){
	http.HandleFunc("/", controllers.Index)
}
