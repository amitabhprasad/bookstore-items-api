package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/amitabhprasad/bookstore-app/bookstore-items-api/client/elasticsearch"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	fmt.Println("initializing elastic search")
	elasticsearch.Init()
	fmt.Println("Starting server....")
	mapUrls()
	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8084",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  time.Second * 60,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
