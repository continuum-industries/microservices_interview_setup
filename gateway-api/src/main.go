package main

import (
	"flag"

	"github.com/continuum-industries/microservices_interview/gateway-api/server"
)

func main() {
	flag.Parse()
	server.StartServer()
}
