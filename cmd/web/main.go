package main

import (
	"log"
	"net/http"

	"github.com/nelbermora/go-web-socket/internal/handlers"
)

func main() {
	log.Println("starting listener routine")
	httpHandler := routes()

	log.Println("starting listener routine")
	go handlers.ListenToWsChannel()

	log.Println("starting in port 9090")

	_ = http.ListenAndServe(":9090", httpHandler)
}
