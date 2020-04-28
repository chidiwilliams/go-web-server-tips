package main

import (
	"log"

	"github.com/chidiwilliams/go-web-server-tips/server"
)

func main() {
	srv := server.Server()
	log.Println("Server listening on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
