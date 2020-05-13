package main

import (
	"github.com/donng/teemo/router"
	"log"
	"net/http"
)

func main() {
	r := router.InitRouter()

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
	}
	log.Fatal(s.ListenAndServe())
}
