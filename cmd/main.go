package main

import (
	"github.com/whoophy/gateway-service/handler"
	"log"
	"net/http"
)

func main() {
	r := handler.Handler()
	port := "3000"
	http.ListenAndServe(":"+port, r)
	log.Printf("Starting up on http://localhost:%s", port)
}
