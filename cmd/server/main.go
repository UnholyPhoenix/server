package main

import (
	"flag"
	"fmt"
	"net/http"
	"server/pkg/calculator"
)

func startServer(address string, handler http.Handler) {
	//Load css
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Starting server on http://" + address)
	http.ListenAndServe(address, handler)
}

func main() {
	var addr = flag.String("addr", "0.0.0.0:8080", "Interface and port to listen on")

	// parse the flags
	flag.Parse()

	kv := calculator.NewMemoryStorage()

	service := calculator.New()

	endpoint := calculator.MakeEndpoint(service, kv)

	handler := calculator.NewHTTPHandler(endpoint)

	startServer(*addr, handler)

}
