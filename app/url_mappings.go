package app

import (
	"net/http"

	"github.com/amitabhprasad/bookstore-app/bookstore-items-api/controllers"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items", controllers.ItemController.Get).Methods(http.MethodGet)

	router.HandleFunc("/ping", controllers.PingController.Ping)
	router.HandleFunc("/", controllers.PingController.Ping)
}
