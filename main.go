package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	"net/http"
)

// ReturnPublicIP parses received remote address
// of client from request's headers and outputs it.
func ReturnPublicIP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, fmt.Sprintf("%s\n", req.Header.Get("X-Forwarded-For")))
}

func main() {

	// Require an IP and port for this service.
	ip := flag.String("ip", "0.0.0.0", "Define this service's IP to operate on.")
	port := flag.String("port", "27300", "Define the port this service is supposed to listen on.")
	flag.Parse()

	// Register above handler on root page.
	http.HandleFunc("/", ReturnPublicIP)

	// Start an HTTP server on supplied port.
	log.Printf("[whats-my-ip] Will listen on %s:%s\n", *ip, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", *ip, *port), nil))
}
