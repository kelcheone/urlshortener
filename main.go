package main

import (
	"flag"
	"log"

	"github.com/kelcheone/urlshortener/api"
)

func main() {
	//get from cli the prot
	listenAddr := flag.String("listen-addr", ":8080", "server listen address")
	flag.Parse()

	log.Fatal(api.NewServer(*listenAddr).Start())
}
