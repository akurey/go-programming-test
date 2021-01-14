package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/akurey/go-programming-test/in_memory"
	"github.com/akurey/go-programming-test/market"
)

const defaultPort = "8080"

func main() {
	// Set configuration variables
	var (
		addr     = envString("PORT", defaultPort)
		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
	)
	flag.Parse()
	// Create a logger
	logger := log.New(os.Stdout, "", log.Lshortfile)
	// Create services
	is := in_memory.NewInventoryService()
	ps := in_memory.NewProductService()
	os := in_memory.NewOrderService(ps, is)
	// Create Market's HTTP handler
	handler := market.MakeHandler(os, is, ps)
	logger.Printf("starting http service at %s", *httpAddr)
	if err := http.ListenAndServe(*httpAddr, handler); err != nil {
		logger.Fatal(err)
	}
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
